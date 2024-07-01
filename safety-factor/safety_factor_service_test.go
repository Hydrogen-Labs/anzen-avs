package safety_factor

import (
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	"math/big"
	"testing"

	example_sf_module "anzen-avs/safety-factor/modules/example_module"

	"github.com/stretchr/testify/assert"
)

func TestNewSafetyFactorService(t *testing.T) {
	logger := sdklogging.NewNoopLogger()
	service := NewSafetyFactorService(logger)

	assert.NotNil(t, service, "Expected non-nil SafetyFactorService instance")
}

func TestSafetyFactorService_RegisterModules(t *testing.T) {
	logger := sdklogging.NewNoopLogger()
	service := NewSafetyFactorService(logger)

	// Ensure the modules are registered correctly
	exampleModule, err := service.GetModuleByOracleIndex(0)
	assert.NoError(t, err, "Expected no error for getting module by index 0")
	assert.NotNil(t, exampleModule, "Expected non-nil module for index 0")
	assert.IsType(t, &example_sf_module.ExampleModule{}, exampleModule, "Expected ExampleModule type for index 0")

}

func TestSafetyFactorService_GetModuleByOracleIndex_NotFound(t *testing.T) {
	logger := sdklogging.NewNoopLogger()
	service := NewSafetyFactorService(logger)

	_, err := service.GetModuleByOracleIndex(999) // An index that is not registered
	assert.Error(t, err, "Expected error for getting module by non-existent index")
	assert.Equal(t, "module not found for the given oracle index", err.Error(), "Expected 'module not found' error message")
}

func TestExampleModule_QueryProfitFromCorruption(t *testing.T) {
	module := example_sf_module.NewExampleModule("exampleName", "exampleID")
	profit, err := module.QueryProfitFromCorruption()

	assert.NoError(t, err, "Expected no error for QueryProfitFromCorruption")
	assert.Equal(t, 150.0, profit, "Expected profit from corruption to be 150.0")
}

func TestExampleModule_QueryCostOfCorruption(t *testing.T) {
	module := example_sf_module.NewExampleModule("exampleName", "exampleID")
	cost, err := module.QueryCostOfCorruption()

	assert.NoError(t, err, "Expected no error for QueryCostOfCorruption")
	assert.Equal(t, 250.0, cost, "Expected cost of corruption to be 250.0")
}

func TestSafetyFactorService_GetSafetyFactorInfoByOracleIndex(t *testing.T) {
	logger := sdklogging.NewNoopLogger()
	service := NewSafetyFactorService(logger)

	expectedSf := big.NewInt(int64((250.0 - 150.0) / 250.0 * 1_000_000_000))

	response, err := service.GetSafetyFactorInfoByOracleIndex(0)
	assert.NoError(t, err, "Expected no error for GetSafetyFactorInfoByOracleIndex")
	assert.NotNil(t, response, "Expected non-nil response for GetSafetyFactorInfoByOracleIndex")
	assert.Equal(t, 150.0, *response.PfC, "Expected PfC to be 150.0")
	assert.Equal(t, 250.0, *response.CoC, "Expected CoC to be 250.0")
	assert.Equal(t, *expectedSf, *response.SF, "Expected SF to be 0.4")
}
