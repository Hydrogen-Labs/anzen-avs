package types

import (
	"errors"

	cstaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IAnzenTaskManagerOraclePullTaskResponse
	TaskResponseMetadata      cstaskmanager.IAnzenTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
