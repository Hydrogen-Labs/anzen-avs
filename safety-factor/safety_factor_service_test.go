package safety_factor

import (
	"testing"

	"go.uber.org/mock/gomock"

	safetyfactormocks "anzen-avs/safety-factor/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNewSafetyFactorService(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	service := safetyfactormocks.NewMockSafetyFactorServicer(mockCtrl)

	assert.NotNil(t, service, "Expected non-nil SafetyFactorService instance")
}

// func TestSafetyFactorService_RegisterModules(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	service := safetyfactormocks.NewMockSafetyFactorServicer(mockCtrl)

// 	// Ensure the modules are registered correctly
// 	exampleModule, err := service.GetModuleByOracleIndex(safety_factor_base.ExampleModuleID)
// 	assert.NoError(t, err, "Expected no error for getting module by index 0")
// 	assert.NotNil(t, exampleModule, "Expected non-nil module for index 0")
// 	assert.IsType(t, &example_sf_module.ExampleModule{}, exampleModule, "Expected ExampleModule type for index 0")

// }

// func TestSafetyFactorService_GetModuleByOracleIndex_NotFound(t *testing.T) {

// 	mockCtrl := gomock.NewController(t)
// 	service := safetyfactormocks.NewMockSafetyFactorServicer(mockCtrl)

// 	_, err := service.GetModuleByOracleIndex(292) // An index that is not registered
// 	assert.Error(t, err, "Expected error for getting module by non-existent index")
// 	assert.Equal(t, "module not found for the given oracle index", err.Error(), "Expected 'module not found' error message")
// }

// func TestExampleModule_QueryProfitFromCorruption(t *testing.T) {
// 	module := example_sf_module.NewExampleModule("exampleName", safety_factor_base.ExampleModuleID)
// 	profit, err := module.QueryProfitFromCorruption()

// 	assert.NoError(t, err, "Expected no error for QueryProfitFromCorruption")
// 	assert.Equal(t, 150.0, profit, "Expected profit from corruption to be 150.0")
// }

// func TestExampleModule_QueryCostOfCorruption(t *testing.T) {
// 	module := example_sf_module.NewExampleModule("exampleName", safety_factor_base.ExampleModuleID)
// 	cost, err := module.QueryCostOfCorruption()

// 	assert.NoError(t, err, "Expected no error for QueryCostOfCorruption")
// 	assert.Equal(t, 250.0, cost, "Expected cost of corruption to be 250.0")
// }

// func TestSafetyFactorService_GetSafetyFactorInfoByOracleIndex(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	service := safetyfactormocks.NewMockSafetyFactorServicer(mockCtrl)

// 	expectedSf := big.NewInt(int64((250.0 - 150.0) / 250.0 * 1_000_000_000))

// 	response, err := service.GetSafetyFactorInfoByOracleIndex(safety_factor_base.ExampleModuleID)
// 	assert.NoError(t, err, "Expected no error for GetSafetyFactorInfoByOracleIndex")
// 	assert.NotNil(t, response, "Expected non-nil response for GetSafetyFactorInfoByOracleIndex")
// 	assert.Equal(t, 150.0, *response.PfC, "Expected PfC to be 150.0")
// 	assert.Equal(t, 250.0, *response.CoC, "Expected CoC to be 250.0")
// 	assert.Equal(t, *expectedSf, *response.SF, "Expected SF to be 0.4")
// }
