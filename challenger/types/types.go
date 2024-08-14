package types

import (
	"errors"

	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
)

type TaskResponseData struct {
	TaskResponse              anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse
	TaskResponseMetadata      anzentaskmanager.IAnzenTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []anzentaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
