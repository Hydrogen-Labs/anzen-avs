package aggregator

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"anzen-avs/aggregator/types"
	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	"anzen-avs/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

func TestProcessSignedOracleTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var NUMBER_TO_SQUARE = uint32(3)

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
	aggregator, _, mockBlsAggServ, _, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	signedTaskResponse, err := createMockSignedOracleTaskResponse(MockTask{
		TaskNum:        TASK_INDEX,
		BlockNumber:    BLOCK_NUMBER,
		NumberToSquare: NUMBER_TO_SQUARE,
	}, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	signedTaskResponseDigest, err := core.GetPullOracleTaskResponseDigest(&signedTaskResponse.OraclePullTaskResponse)
	assert.Nil(t, err)

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedTaskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId)
	err = aggregator.ProcessSignedOraclePullTaskResponse(signedTaskResponse, nil)
	assert.Nil(t, err)
}

// mocks an operator signing on a task response
func createMockSignedOracleTaskResponse(mockTask MockTask, keypair bls.KeyPair) (*SignedOraclePullTaskResponse, error) {
	taskResponse := &anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse{
		ReferenceTaskIndex: mockTask.TaskNum,
		SafetyFactor:       big.NewInt(1_000_000_000),
	}
	taskResponseHash, err := core.GetPullOracleTaskResponseDigest(taskResponse)
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(taskResponseHash)
	signedTaskResponse := &SignedOraclePullTaskResponse{
		OraclePullTaskResponse: *taskResponse,
		BlsSignature:           *blsSignature,
		OperatorId:             MOCK_OPERATOR_ID,
	}
	return signedTaskResponse, nil
}
