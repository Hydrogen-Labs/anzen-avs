package aggregator

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	safetyfactormocks "anzen-avs/safety-factor/mocks"
	safety_factor_base "anzen-avs/safety-factor/safety-factor-base"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	blsaggservmock "github.com/Layr-Labs/eigensdk-go/services/mocks/blsagg"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"anzen-avs/aggregator/mocks"
	"anzen-avs/aggregator/types"
	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	chainiomocks "anzen-avs/core/chainio/mocks"
)

var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}
var MOCK_OPERATOR_STAKE = big.NewInt(100)
var MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING = "50"

type MockTask struct {
	TaskNum        uint32
	BlockNumber    uint32
	NumberToSquare uint32
}

func TestSendNewTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	MOCK_OPERATOR_G1PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
	MOCK_OPERATOR_G2PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()

	operatorPubkeyDict := map[sdktypes.OperatorId]types.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
				G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
			},
			OperatorAddr: common.Address{},
		},
	}

	aggregator, mockAvsWriterer, mockBlsAggService, mockSafetyFactorServicer, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var ORACLE_INDEX = uint32(0)
	var ORACLE_INDEX_BIG_INT = big.NewInt(int64(ORACLE_INDEX))
	var PROPSED_SAFETY_FACTOR = big.NewInt(400_000_000)

	mockSafetyFactorServicer.EXPECT().GetSafetyFactorInfoByOracleIndex(int(ORACLE_INDEX)).Return(&safety_factor_base.SFModuleResponse{
		SF: PROPSED_SAFETY_FACTOR,
	}, nil)

	mockAvsWriterer.EXPECT().SendNewOraclePullTask(
		context.Background(), ORACLE_INDEX_BIG_INT, PROPSED_SAFETY_FACTOR, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS,
	).Return(mocks.MocksSendNewOraclePullTaskCall(BLOCK_NUMBER, TASK_INDEX, ORACLE_INDEX, *PROPSED_SAFETY_FACTOR))

	// 100 blocks, each takes 12 seconds. We hardcode for now since aggregator also hardcodes this value
	taskTimeToExpiry := 100 * 12 * time.Second
	// make sure that initializeNewTask was called on the blsAggService
	// maybe there's a better way to do this? There's a saying "don't mock 3rd party code"
	// see https://hynek.me/articles/what-to-mock-in-5-mins/

	mockBlsAggService.EXPECT().InitializeNewTask(TASK_INDEX, BLOCK_NUMBER, types.QUORUM_NUMBERS, sdktypes.QuorumThresholdPercentages{types.QUORUM_THRESHOLD_NUMERATOR}, taskTimeToExpiry)

	err = aggregator.sendNewOraclePullTask(ORACLE_INDEX_BIG_INT)

	assert.Nil(t, err)
}

func createMockAggregator(
	mockCtrl *gomock.Controller, operatorPubkeyDict map[sdktypes.OperatorId]types.OperatorInfo,
) (*Aggregator, *chainiomocks.MockAvsWriterer, *blsaggservmock.MockBlsAggregationService, *safetyfactormocks.MockSafetyFactorServicer, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockAvsReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
	mockBlsAggregationService := blsaggservmock.NewMockBlsAggregationService(mockCtrl)
	mockSafetyFactorService := safetyfactormocks.NewMockSafetyFactorServicer(mockCtrl)

	aggregator := &Aggregator{
		logger:                logger,
		avsWriter:             mockAvsWriter,
		avsReader:             mockAvsReader,
		blsAggregationService: mockBlsAggregationService,
		oracleTasks:           make(map[types.TaskIndex]anzentaskmanager.IAnzenTaskManagerOraclePullTask),
		oracleTaskReponses:    make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse),
		safetyFactorService:   mockSafetyFactorService,
	}
	return aggregator, mockAvsWriter, mockBlsAggregationService, mockSafetyFactorService, nil
}

// just a mock ethclient to pass to bindings
// so that we can access abi methods
func createMockEthClient() *backends.SimulatedBackend {
	genesisAlloc := map[common.Address]gethcore.GenesisAccount{}
	blockGasLimit := uint64(1000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	return client
}
