package safety_factor

import (
	"anzen-avs/core/chainio"
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"

	anzen_sf_module "anzen-avs/safety-factor/modules/anzen_module"
	example_sf_module "anzen-avs/safety-factor/modules/example_module"
	safetyfactorbase "anzen-avs/safety-factor/safety-factor-base"

	"errors"
)

// interface
type SafetyFactorServicer interface {
	GetModuleByOracleIndex(index int) (safetyfactorbase.SFModule, error)
	GetSafetyFactorInfoByOracleIndex(index int) (*safetyfactorbase.SFModuleResponse, error)
	IsSafetyFactorInfoStale(moduleID int32) (bool, error)
}

type SafetyFactorService struct {
	avsReader chainio.AvsReaderer
	logger    logging.Logger

	modules map[int]safetyfactorbase.SFModule
}

func NewSafetyFactorService(logger logging.Logger, avsReader *chainio.AvsReader) *SafetyFactorService {
	service := &SafetyFactorService{
		avsReader: avsReader,
		modules:   make(map[int]safetyfactorbase.SFModule),
		logger:    logger,
	}
	service.registerModules()
	return service
}

func (s *SafetyFactorService) IsSafetyFactorInfoStale(moduleID int32) (bool, error) {
	latestSafetyFactorOffChainInfo, err := s.GetSafetyFactorInfoByOracleIndex(int(moduleID))
	if err != nil {
		s.logger.Error("Aggregator failed to get safety factor info", "err", err)
		return false, err
	}
	currentSafetyFactorOnChainInfo, err := s.avsReader.GetSafetyFactorByIndex(context.Background(), uint32(moduleID))
	s.logger.Info("Current safety factor info", "currentSafetyFactorInfo", currentSafetyFactorOnChainInfo)
	if err != nil {
		s.logger.Error("Aggregator failed to get safety factor info", "err", err)
		return false, err
	}

	// Get the current time as a big.Int
	timeNow := big.NewInt(0).SetUint64(uint64(time.Now().Unix()))

	// // One minute in nanoseconds
	oneMinute := big.NewInt(int64(60))
	// TODO let this be configurable

	// Check if the safety factor has deviated
	isDeviated := latestSafetyFactorOffChainInfo.SF.Cmp(currentSafetyFactorOnChainInfo.SafetyFactor) != 0

	// Check if the safety factor is stale by one minute
	isStale := big.NewInt(0).Sub(timeNow, currentSafetyFactorOnChainInfo.Timestamp).Cmp(oneMinute) > 0

	isDeviatedOrStale := isDeviated || isStale
	return isDeviatedOrStale, nil
}

func (s *SafetyFactorService) registerModules() {
	s.modules[safetyfactorbase.ExampleModuleID] = example_sf_module.NewExampleModule("exampleName", safetyfactorbase.ExampleModuleID)
	s.modules[safetyfactorbase.AnzenModuleID] = anzen_sf_module.NewAnzenModule()
}

func (s *SafetyFactorService) GetModuleByOracleIndex(index int) (safetyfactorbase.SFModule, error) {
	module, exists := s.modules[index]
	if !exists {
		return nil, errors.New("module not found for the given oracle index")
	}
	return module, nil
}

func (s *SafetyFactorService) GetSafetyFactorInfoByOracleIndex(index int) (*safetyfactorbase.SFModuleResponse, error) {
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
