package safety_factor_base

import (
	"errors"
	"log"
	"math/big"
	"time"
)

type SFModule interface {
	QueryProfitFromCorruption() (float64, error)
	QueryCostOfCorruption() (float64, error)
	SafeQuerySafetyFactorInfo() (SFModuleResponse, error)
}

type SFModuleConfig struct {
	SFLowerBound float64
	SFUpperBound float64
}

type SFModuleNodeConfig struct {
	// Define the fields as needed
}

// define constant 1_000_000_000 to represent 1 billion big int
const PRECISION = 1_000_000_000

type SFModuleResponse struct {
	SF        *big.Int
	PfC       *float64
	CoC       *float64
	Timestamp int64
	Error     string
}

var DefaultSFModuleConfig = SFModuleNodeConfig{}

type BaseSFModule struct {
	Name       string
	ID         string
	NodeConfig SFModuleNodeConfig
	Config     *SFModuleConfig
	Module     SFModule // Add reference to interface
}

func NewBaseSFModule(name, id string, module SFModule) *BaseSFModule {
	m := &BaseSFModule{
		Name:       name,
		ID:         id,
		NodeConfig: DefaultSFModuleConfig,
		Module:     module, // Initialize interface reference
	}

	config, err := m.querySafetyFactorConfig()
	if err == nil {
		m.Config = config
	}

	return m
}

func (m *BaseSFModule) querySafetyFactorConfig() (*SFModuleConfig, error) {
	// return mock data
	// Todo: replace with actual on-chain query
	return &SFModuleConfig{
		SFLowerBound: 0.2,
		SFUpperBound: 0.3,
	}, nil
}

func (m *BaseSFModule) SafeQueryProfitFromCorruption() (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Failed to query profit from corruption: %v", r)
		}
	}()

	return m.Module.QueryProfitFromCorruption() // Call the interface method
}

func (m *BaseSFModule) SafeQueryCostOfCorruption() (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Failed to query cost of corruption: %v", r)
		}
	}()

	return m.Module.QueryCostOfCorruption() // Call the interface method
}

func (m *BaseSFModule) SafeQuerySafetyFactorInfo() (SFModuleResponse, error) {
	PfC, err1 := m.SafeQueryProfitFromCorruption()
	CoC, err2 := m.SafeQueryCostOfCorruption()

	if err1 != nil || err2 != nil {
		errorMessage := "PfC or CoC query failed"
		log.Printf(errorMessage)
		return SFModuleResponse{
			SF:        nil,
			PfC:       nil,
			CoC:       nil,
			Timestamp: time.Now().Unix(),
			Error:     errorMessage,
		}, errors.New(errorMessage)
	}

	if CoC == 0 {
		errorMessage := "CoC is zero, cannot divide"
		log.Printf(errorMessage)
		return SFModuleResponse{
			SF:        nil,
			PfC:       &PfC,
			CoC:       &CoC,
			Timestamp: time.Now().Unix(),
			Error:     errorMessage,
		}, errors.New(errorMessage)
	}

	sf := (CoC - PfC) / CoC
	sfInt := big.NewInt(int64(sf * PRECISION))

	return SFModuleResponse{
		SF:        sfInt,
		PfC:       &PfC,
		CoC:       &CoC,
		Timestamp: time.Now().Unix(),
		Error:     "",
	}, nil
}
