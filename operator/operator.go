package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"anzen-avs/aggregator"
	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	"anzen-avs/core"
	"anzen-avs/core/chainio"
	"anzen-avs/metrics"
	safety_factor "anzen-avs/safety-factor"
	"anzen-avs/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const AVS_NAME = "anzen"
const SEM_VER = "0.0.1"

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.Client
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
	// receive new oracle pull tasks in this chan (typically from listening to onchain event)	// receive new oracle pull tasks in this chan (typically from listening to onchain event)
	newOraclePullTaskCreatedChan chan *anzentaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated
	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	anzenServiceManagerAddr common.Address

	safetyFactorService safety_factor.SafetyFactorServicer
}

// TODO(samlaf): config is a mess right now, since the chainio client constructors
//
//	take the config in core (which is shared with aggregator and challenger)
func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	var logLevel logging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient eth.Client
	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewClient(c.EthWsUrl)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}
	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                 c.EthRpcUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   c.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		c.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, common.HexToAddress(c.OperatorAddress))

	avsWriter, err := chainio.BuildAvsWriter(
		txMgr, common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		common.HexToAddress(c.SafetyFactorOracleAddr),
		ethRpcClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		common.HexToAddress(c.SafetyFactorOracleAddr),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriber(common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		common.HexToAddress(c.SafetyFactorOracleAddr),
		ethWsClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	safetyFactorService := safety_factor.NewSafetyFactorService(logger, avsReader)

	operator := &Operator{
		config:                       c,
		logger:                       logger,
		metricsReg:                   reg,
		metrics:                      avsAndEigenMetrics,
		nodeApi:                      nodeApi,
		ethClient:                    ethRpcClient,
		avsWriter:                    avsWriter,
		avsReader:                    avsReader,
		avsSubscriber:                avsSubscriber,
		eigenlayerReader:             sdkClients.ElChainReader,
		eigenlayerWriter:             sdkClients.ElChainWriter,
		blsKeypair:                   blsKeyPair,
		operatorAddr:                 common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr:   c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:          aggregatorRpcClient,
		newOraclePullTaskCreatedChan: make(chan *anzentaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated),
		anzenServiceManagerAddr:      common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		operatorId:                   [32]byte{0}, // this is set below
		safetyFactorService:          safetyFactorService,
	}

	if c.RegisterOperatorOnStartup {
		depositAmount := big.NewInt(int64(c.DepositAmount))
		operator.registerOperatorOnStartup(operatorEcdsaPrivateKey, common.HexToAddress(c.TokenStrategyAddr), depositAmount)
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.operatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		o.logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
		// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
		return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}

	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}
	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub := o.avsSubscriber.SubcribeToNewOraclePullTasks(o.newOraclePullTaskCreatedChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.logger.Fatal("Error in metrics server", "err", err)
		case err := <-sub.Err():
			o.logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub = o.avsSubscriber.SubcribeToNewOraclePullTasks(o.newOraclePullTaskCreatedChan)

		case newOraclePullTaskCreatedLog := <-o.newOraclePullTaskCreatedChan:
			// TODO Record metrics for new oracle pull task
			o.logger.Info("Received new oracle pull task", "taskIndex", newOraclePullTaskCreatedLog.OraclePullTask)
			taskResponse, err := o.ProcessNewOraclePullTaskLog(newOraclePullTaskCreatedLog)
			if err != nil {
				continue
			}
			signedPullOracleTaskResponse, err := o.SignOraclePullTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedOraclePullTaskReponseToAggregator(signedPullOracleTaskResponse)
			o.logger.Info("Sending task response to aggregator", "taskResponse", taskResponse)

		}

	}
}

// Takes a newOraclePullTaskSolutionProposedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewOraclePullTaskLog(newOraclePullTaskLog *anzentaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated) (*anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse, error) {
	o.logger.Info("Received new oracle pull task solution proposed", "task", newOraclePullTaskLog)

	oracleIndex := newOraclePullTaskLog.OraclePullTask.OracleIndex

	safetyFactorInfo, err := o.safetyFactorService.GetSafetyFactorInfoByOracleIndex(int(oracleIndex))
	if err != nil {
		o.logger.Error("Error getting safety factor info", "err", err)
		return nil, err
	}

	// Check if the proposed solution is equal to what we expect
	if safetyFactorInfo.SF.Cmp(newOraclePullTaskLog.OraclePullTask.ProposedSafetyFactor) != 0 {
		o.logger.Error("Proposed solution does not match expected solution", "expected", safetyFactorInfo.SF, "proposed", newOraclePullTaskLog.OraclePullTask.ProposedSafetyFactor)

		return nil, fmt.Errorf("proposed solution does not match expected solution")
	}

	// Check if the safety factor info is stale
	isStale, err := o.safetyFactorService.IsSafetyFactorInfoStale(int32(oracleIndex))
	if err != nil {
		o.logger.Error("Error checking if safety factor info is stale", "err", err)
		return nil, err
	}

	if !isStale {
		o.logger.Error("Safety factor info is not stale", "safetyFactorInfo", safetyFactorInfo)
		return nil, fmt.Errorf("safety factor info is not stale")
	}

	taskResponse := &anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse{
		ReferenceTaskIndex: newOraclePullTaskLog.TaskIndex,
		SafetyFactor:       safetyFactorInfo.SF,
	}

	return taskResponse, nil
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.

func (o *Operator) SignOraclePullTaskResponse(oraclePullTaskResponse *anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse) (*aggregator.SignedOraclePullTaskResponse, error) {
	oraclePullTaskResponseHash, err := core.GetPullOracleTaskResponseDigest(oraclePullTaskResponse)
	if err != nil {
		o.logger.Error("Error getting oracle pull task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}
	blsSignature := o.blsKeypair.SignMessage(oraclePullTaskResponseHash)
	signedOraclePullTaskResponse := &aggregator.SignedOraclePullTaskResponse{
		OraclePullTaskResponse: *oraclePullTaskResponse,
		BlsSignature:           *blsSignature,
		OperatorId:             o.operatorId,
	}
	o.logger.Debug("Signed oracle pull task response", "signedOraclePullTaskResponse", signedOraclePullTaskResponse)
	return signedOraclePullTaskResponse, nil
}
