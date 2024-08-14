package operator

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	"anzen-avs/metrics"
	safetyfactormock "anzen-avs/safety-factor/mocks"
	"anzen-avs/tests"
)

const MOCK_OPERATOR_BLS_PRIVATE_KEY = "69"

// hash of bls_public_key (hardcoded for sk=69)
var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}

// Name starts with Integration test because we don't want it to run with go test ./...
// since this starts a chain and takes longer to run
// TODO(samlaf): buggy test, fix it
func IntegrationTestOperatorRegistration(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	anvilCmd := tests.StartAnvilChainAndDeployContracts()
	defer anvilCmd.Process.Kill()
	operator, _, err := createMockOperator(mockCtrl)
	assert.Nil(t, err)
	err = operator.RegisterOperatorWithEigenlayer()
	assert.Nil(t, err)
}

func createMockOperator(mockCtrl *gomock.Controller) (*Operator, *safetyfactormock.MockSafetyFactorServicer, error) {
	logger := sdklogging.NewNoopLogger()
	reg := prometheus.NewRegistry()
	noopMetrics := metrics.NewNoopMetrics()
	mockSafetyFactorService := safetyfactormock.NewMockSafetyFactorServicer(mockCtrl)

	blsPrivateKey, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	if err != nil {
		return nil, nil, err
	}
	operatorKeypair := bls.NewKeyPair(blsPrivateKey)

	operator := &Operator{
		logger:                       logger,
		blsKeypair:                   operatorKeypair,
		metricsReg:                   reg,
		metrics:                      noopMetrics,
		safetyFactorService:          mockSafetyFactorService,
		newOraclePullTaskCreatedChan: make(chan *anzentaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated),
		operatorId:                   MOCK_OPERATOR_ID,
	}
	return operator, mockSafetyFactorService, nil
}
