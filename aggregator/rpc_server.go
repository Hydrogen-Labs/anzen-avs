package aggregator

import (
	"context"
	"errors"
	"net/http"
	"net/rpc"

	cstaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	"anzen-avs/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type SignedTaskResponse struct {
	TaskResponse cstaskmanager.IAnzenTaskManagerTaskResponse
	BlsSignature bls.Signature
	OperatorId   types.OperatorId
}

type SignedOraclePullTaskResponse struct {
	OraclePullTaskResponse cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse
	BlsSignature           bls.Signature
	OperatorId             types.OperatorId
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the oracle pull task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedOraclePullTaskResponse(signedOraclePullTaskResponse *SignedOraclePullTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed oracle pull task response: %#v", signedOraclePullTaskResponse)
	taskIndex := signedOraclePullTaskResponse.OraclePullTaskResponse.ReferenceTaskIndex
	taskResponseDigest, err := core.GetPullOracleTaskResponseDigest(&signedOraclePullTaskResponse.OraclePullTaskResponse)
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}
	agg.oracleTaskReponsesMu.Lock()
	if _, ok := agg.oracleTaskReponses[taskIndex]; !ok {
		agg.oracleTaskReponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse)
	}
	if _, ok := agg.oracleTaskReponses[taskIndex][taskResponseDigest]; !ok {
		agg.oracleTaskReponses[taskIndex][taskResponseDigest] = signedOraclePullTaskResponse.OraclePullTaskResponse
	}
	agg.oracleTaskReponsesMu.Unlock()

	err = agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedOraclePullTaskResponse.BlsSignature, signedOraclePullTaskResponse.OperatorId,
	)
	return err
}
