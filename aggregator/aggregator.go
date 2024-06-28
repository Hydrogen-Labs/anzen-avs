package aggregator

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"

	"anzen-avs/aggregator/types"
	"anzen-avs/core"
	"anzen-avs/core/chainio"
	"anzen-avs/core/config"
	safety_factor "anzen-avs/safety-factor"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	cstaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "incredible-squaring"
)

// Aggregator sends tasks (numbers to square) onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThresholdPercentage for each quorum
// (currently we only use a single quorum of the ERC20Mock token), it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Aggregator struct {
	logger              logging.Logger
	serverIpPortAddr    string
	avsWriter           chainio.AvsWriterer
	safetyFactorService safety_factor.SafetyFactorServicer
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	oracleTasks           map[types.TaskIndex]cstaskmanager.IAnzenTaskManagerOraclePullTask
	oracleTasksMu         sync.RWMutex
	oracleTaskReponses    map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse
	oracleTaskReponsesMu  sync.RWMutex
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {

	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.IncredibleSquaringRegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
	}
	clients, err := clients.BuildAll(chainioConfig, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)
	safetyFactorService := safety_factor.NewSafetyFactorService(c.Logger)

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		safetyFactorService:   safetyFactorService,
		oracleTasks:           make(map[types.TaskIndex]cstaskmanager.IAnzenTaskManagerOraclePullTask),
		oracleTaskReponses:    make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")
	agg.logger.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	// TODO(soubhik): refactor task generation/sending into a separate function that we can run as goroutine
	ticker := time.NewTicker(10 * time.Second)
	agg.logger.Infof("Aggregator set to send new task every 10 seconds...")
	defer ticker.Stop()
	taskNum := int64(0)
	// ticker doesn't tick immediately, so we send the first task here
	// see https://github.com/golang/go/issues/17601
	_ = agg.sendNewOraclePullTask(big.NewInt(0))
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			agg.sendAggregatedOracleResponseToContract(blsAggServiceResp)
		case <-ticker.C:
			// TODO: create some policy for when to send oracle pull tasks
			err := agg.sendNewOraclePullTask(big.NewInt(0))
			taskNum++
			if err != nil {
				// we log the errors inside sendNewTask() so here we just continue to the next task
				continue
			}
		}
	}
}

func (agg *Aggregator) sendAggregatedOracleResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// TODO: check if blsAggServiceResp contains an err
	if blsAggServiceResp.Err != nil {
		agg.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panicing to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
	}
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []cstaskmanager.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	agg.logger.Info("Threshold reached. Sending aggregated response onchain.",
		"taskIndex", blsAggServiceResp.TaskIndex,
	)
	agg.oracleTasksMu.RLock()
	task := agg.oracleTasks[blsAggServiceResp.TaskIndex]
	agg.oracleTasksMu.RUnlock()
	agg.oracleTaskReponsesMu.RLock()
	taskResponse := agg.oracleTaskReponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	agg.oracleTaskReponsesMu.RUnlock()
	_, err := agg.avsWriter.SendAggregatedOracleResponse(context.Background(), task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		agg.logger.Error("Aggregator failed to respond to task", "err", err)
	}
}

func (agg *Aggregator) sendNewOraclePullTask(oracleIndex *big.Int) error {
	agg.logger.Info("Aggregator sending new oracle pull task", "oracleIndex", oracleIndex)
	// Send number to square to the task manager contract
	proposedSafetyFactorInfo, err := agg.safetyFactorService.GetSafetyFactorInfoByOracleIndex(int(oracleIndex.Int64()))
	if err != nil {
		agg.logger.Error("Aggregator failed to get safety factor info", "err", err)
		return err
	}

	newPullTask, taskIndex, err := agg.avsWriter.SendNewOraclePullTask(context.Background(), oracleIndex, proposedSafetyFactorInfo.SF, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS)

	if err != nil {
		agg.logger.Error("Aggregator failed to send new oracle pull task", "err", err)
		return err
	}

	agg.logger.Info("Aggregator sent new oracle pull task", "task", newPullTask)

	agg.oracleTasksMu.Lock()
	agg.oracleTasks[taskIndex] = newPullTask
	agg.oracleTasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newPullTask.QuorumNumbers))
	for i := range newPullTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newPullTask.QuorumThresholdPercentage)
	}

	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newPullTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}
	agg.blsAggregationService.InitializeNewTask(taskIndex, newPullTask.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)
	return nil
}
