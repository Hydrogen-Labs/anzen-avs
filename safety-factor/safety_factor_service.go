package safety_factor

import (
	"github.com/Layr-Labs/eigensdk-go/logging"

	example_sf_module "anzen-avs/safety-factor/modules/example_module"
	safety_factor_base "anzen-avs/safety-factor/safety-factor-base"
	"errors"
)

// interface
type SafetyFactorServicer interface {
	GetModuleByOracleIndex(index int) (safety_factor_base.SFModule, error)
	GetSafetyFactorInfoByOracleIndex(index int) (*safety_factor_base.SFModuleResponse, error)
}

type SafetyFactorService struct {
	logger logging.Logger

	modules map[int]safety_factor_base.SFModule
}

func NewSafetyFactorService(logger logging.Logger) *SafetyFactorService {
	service := &SafetyFactorService{
		modules: make(map[int]safety_factor_base.SFModule),
		logger:  logger,
	}
	service.registerModules()
	return service
}

func (s *SafetyFactorService) registerModules() {
	s.modules[0] = example_sf_module.NewExampleModule("exampleName", "exampleID")
}

func (s *SafetyFactorService) GetModuleByOracleIndex(index int) (safety_factor_base.SFModule, error) {
	module, exists := s.modules[index]
	if !exists {
		return nil, errors.New("module not found for the given oracle index")
	}
	return module, nil
}

func (s *SafetyFactorService) GetSafetyFactorInfoByOracleIndex(index int) (*safety_factor_base.SFModuleResponse, error) {
	module, err := s.GetModuleByOracleIndex(index)
	if err != nil {
		return nil, err
	}

	response, err := module.SafeQuerySafetyFactorInfo()
	if err != nil {
		return nil, err
	}

	return &response, nil
}
