package challenger

import (
	"math/big"
	"testing"

	chtypes "anzen-avs/challenger/types"
	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	chainiomocks "anzen-avs/core/chainio/mocks"

	mockethclient "github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const TASK_INDEX = 1
const BLOCK_NUMBER = uint32(100)

func TestCallChallengeModule(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	challenger, _, _, _, _, err := createMockChallenger(mockCtrl)
	assert.Nil(t, err)

	challenger.tasks[TASK_INDEX] = anzentaskmanager.IAnzenTaskManagerOraclePullTask{
		OracleIndex: uint32(big.NewInt(1).Uint64()),
	}

	challenger.taskResponses[TASK_INDEX] = chtypes.TaskResponseData{
		TaskResponse: anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse{
			ReferenceTaskIndex: TASK_INDEX,
			SafetyFactor:       big.NewInt(1),
		},
	}

	// mockAvsWriter.EXPECT().RaiseChallenge(
	// 	context.Background(),
	// 	challenger.tasks[TASK_INDEX],
	// 	challenger.taskResponses[TASK_INDEX].TaskResponse,
	// 	challenger.taskResponses[TASK_INDEX].TaskResponseMetadata,
	// 	challenger.taskResponses[TASK_INDEX].NonSigningOperatorPubKeys,
	// ).Return(mocks.MockRaiseAndResolveChallengeCall(BLOCK_NUMBER, TASK_INDEX), nil)

	// err = challenger.callChallengeModule(TASK_INDEX)
	// assert.Nil(t, err)
}

func createMockChallenger(mockCtrl *gomock.Controller) (*Challenger, *chainiomocks.MockAvsWriterer, *chainiomocks.MockAvsReaderer, *chainiomocks.MockAvsSubscriberer, *mockethclient.MockEthClient, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockAvsReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
	mockAvsSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
	mockEthClient := mockethclient.NewMockEthClient(mockCtrl)
	challenger := &Challenger{
		logger:             logger,
		avsWriter:          mockAvsWriter,
		avsReader:          mockAvsReader,
		ethClient:          mockEthClient,
		avsSubscriber:      mockAvsSubscriber,
		tasks:              make(map[uint32]anzentaskmanager.IAnzenTaskManagerOraclePullTask),
		taskResponses:      make(map[uint32]chtypes.TaskResponseData),
		taskResponseChan:   make(chan *anzentaskmanager.ContractAnzenTaskManagerOraclePullTaskResponded),
		newTaskCreatedChan: make(chan *anzentaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated),
	}
	return challenger, mockAvsWriter, mockAvsReader, mockAvsSubscriber, mockEthClient, nil
}
