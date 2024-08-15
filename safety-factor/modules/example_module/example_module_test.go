package example_sf_module

import (
	"log"
	"math/big"
	"testing"

	safetyfactorbase "anzen-avs/safety-factor/safety-factor-base"

	"github.com/stretchr/testify/assert"
)

func TestExampleModule_QueryProfitFromCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", safetyfactorbase.ExampleModuleID)
	profit, err := module.QueryProfitFromCorruption()

	assert.NoError(t, err, "Expected no error for QueryProfitFromCorruption")
	assert.Equal(t, 150.0, profit, "Expected profit from corruption to be 150.0")
}

func TestExampleModule_QueryCostOfCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", safetyfactorbase.ExampleModuleID)
	cost, err := module.QueryCostOfCorruption()

	assert.NoError(t, err, "Expected no error for QueryCostOfCorruption")
	assert.Equal(t, 250.0, cost, "Expected cost of corruption to be 250.0")
}

func TestExampleModule_SafeQueryProfitFromCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", safetyfactorbase.ExampleModuleID)
	profit, err := module.SafeQueryProfitFromCorruption()

	assert.NoError(t, err, "Expected no error for SafeQueryProfitFromCorruption")
	assert.Equal(t, 150.0, profit, "Expected profit from corruption to be 150.0")
}

func TestExampleModule_SafeQueryCostOfCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", safetyfactorbase.ExampleModuleID)
	cost, err := module.SafeQueryCostOfCorruption()

	assert.NoError(t, err, "Expected no error for SafeQueryCostOfCorruption")
	assert.Equal(t, 250.0, cost, "Expected cost of corruption to be 250.0")
}

func TestExampleModule_SafeQuerySafetyFactorInfo(t *testing.T) {
	module := NewExampleModule("exampleName", safetyfactorbase.ExampleModuleID)
	response, err := module.SafeQuerySafetyFactorInfo()

	expectedSf := big.NewInt(int64((250.0 - 150.0) / 250.0 * 1_000_000_000))

	log.Printf("response: %v", response)

	assert.NoError(t, err, "Expected no error for SafeQuerySafetyFactorInfo")
	assert.NotNil(t, response.SF, "Expected SF to be not nil")
	assert.Equal(t, 150.0, *response.PfC, "Expected PfC to be 150.0")
	assert.Equal(t, 250.0, *response.CoC, "Expected CoC to be 250.0")
	assert.Equal(t, *expectedSf, *response.SF, "Expected SF to be 0.4")
	assert.Empty(t, response.Error, "Expected no error message")
}
