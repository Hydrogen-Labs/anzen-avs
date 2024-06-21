package example_sf_module

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleModule_QueryProfitFromCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", "exampleID")
	profit, err := module.QueryProfitFromCorruption()

	assert.NoError(t, err, "Expected no error for QueryProfitFromCorruption")
	assert.Equal(t, 150.0, profit, "Expected profit from corruption to be 150.0")
}

func TestExampleModule_QueryCostOfCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", "exampleID")
	cost, err := module.QueryCostOfCorruption()

	assert.NoError(t, err, "Expected no error for QueryCostOfCorruption")
	assert.Equal(t, 250.0, cost, "Expected cost of corruption to be 250.0")
}

func TestExampleModule_SafeQueryProfitFromCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", "exampleID")
	profit, err := module.SafeQueryProfitFromCorruption()

	assert.NoError(t, err, "Expected no error for SafeQueryProfitFromCorruption")
	assert.Equal(t, 150.0, profit, "Expected profit from corruption to be 150.0")
}

func TestExampleModule_SafeQueryCostOfCorruption(t *testing.T) {
	module := NewExampleModule("exampleName", "exampleID")
	cost, err := module.SafeQueryCostOfCorruption()

	assert.NoError(t, err, "Expected no error for SafeQueryCostOfCorruption")
	assert.Equal(t, 250.0, cost, "Expected cost of corruption to be 250.0")
}

func TestExampleModule_SafeQuerySafetyFactorInfo(t *testing.T) {
	module := NewExampleModule("exampleName", "exampleID")
	response, err := module.SafeQuerySafetyFactorInfo()

	log.Printf("response: %v", response)

	assert.NoError(t, err, "Expected no error for SafeQuerySafetyFactorInfo")
	assert.NotNil(t, response.SF, "Expected SF to be not nil")
	assert.Equal(t, 150.0, *response.PfC, "Expected PfC to be 150.0")
	assert.Equal(t, 250.0, *response.CoC, "Expected CoC to be 250.0")
	assert.Equal(t, (250.0-150.0)/250.0, *response.SF, "Expected SF to be 0.4")
	assert.Empty(t, response.Error, "Expected no error message")
}
