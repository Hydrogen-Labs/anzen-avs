package operator

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"anzen-avs/aggregator"
	aggtypes "anzen-avs/aggregator/types"
	cstaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	chainiomocks "anzen-avs/core/chainio/mocks"
	operatormocks "anzen-avs/operator/mocks"
)

func TestOperator(t *testing.T) {
	operator, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

	t.Run("ProcessNewTaskCreatedLog", func(t *testing.T) {
		var ORACLE_INDEX = uint32(0)
		var proposedSafetyFactor = big.NewInt(400_000_000)

		newTaskCreatedLog := &cstaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated{
			TaskIndex: taskIndex,
			OraclePullTask: cstaskmanager.IAnzenTaskManagerOraclePullTask{
				OracleIndex:               ORACLE_INDEX,
				ProposedSafetyFactor:      proposedSafetyFactor,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
				QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
			},
			Raw: types.Log{},
		}

		got, err := operator.ProcessNewOraclePullTaskLog(newTaskCreatedLog)
		assert.Nil(t, err)

		want := &cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse{
			ReferenceTaskIndex: taskIndex,
			SafetyFactor:       proposedSafetyFactor,
		}
		assert.Equal(t, got, want)
	})

	t.Run("Start", func(t *testing.T) {
		var ORACLE_INDEX = uint32(0)
		var proposedSafetyFactor = big.NewInt(400_000_000)

		// new task event
		newTaskCreatedEvent := &cstaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated{
			TaskIndex: taskIndex,
			OraclePullTask: cstaskmanager.IAnzenTaskManagerOraclePullTask{
				OracleIndex:               ORACLE_INDEX,
				ProposedSafetyFactor:      proposedSafetyFactor,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
				QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)
		X, ok := big.NewInt(0).SetString("9021072218505966130870962639715395918774808826906850912982707609409475544839", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("13150337747840177229943570937415137795286280110433164353723356045586207798531", 10)
		assert.True(t, ok)

		signedTaskResponse := &aggregator.SignedOraclePullTaskResponse{
			OraclePullTaskResponse: cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse{
				ReferenceTaskIndex: taskIndex,
				SafetyFactor:       proposedSafetyFactor,
			},
			BlsSignature: bls.Signature{
				G1Point: bls.NewG1Point(X, Y),
			},
			OperatorId: operator.operatorId,
		}
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := operatormocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedOraclePullTaskReponseToAggregator(signedTaskResponse)
		operator.aggregatorRpcClient = mockAggregatorRpcClient

		mockSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
		mockSubscriber.EXPECT().SubcribeToNewOraclePullTasks(operator.newOraclePullTaskCreatedChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}))
		operator.avsSubscriber = mockSubscriber

		mockReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
		mockReader.EXPECT().IsOperatorRegistered(gomock.Any(), operator.operatorAddr).Return(true, nil)
		operator.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			err := operator.Start(ctx)
			assert.Nil(t, err)
		}()
		operator.newOraclePullTaskCreatedChan <- newTaskCreatedEvent
		time.Sleep(1 * time.Second)

		cancel()
	})

}
