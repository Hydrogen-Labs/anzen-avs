package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	anzenservicemanager "anzen-avs/contracts/bindings/AnzenServiceManager"
	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	erc20mock "anzen-avs/contracts/bindings/ERC20Mock"
	safetyfactororacle "anzen-avs/contracts/bindings/SafetyFactorOracle"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

type AvsManagersBindings struct {
	TaskManager        *anzentaskmanager.ContractAnzenTaskManager
	ServiceManager     *anzenservicemanager.ContractAnzenServiceManager
	SafetyFactorOracle *safetyfactororacle.ContractSafetyFactorOracle
	ethClient          eth.Client
	logger             logging.Logger
}

func NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, safetyFactorOracleAddr gethcommon.Address, ethclient eth.Client, logger logging.Logger) (*AvsManagersBindings, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(registryCoordinatorAddr, ethclient)
	if err != nil {
		return nil, err
	}

	serviceManagerAddr, err := contractRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	contractServiceManager, err := anzenservicemanager.NewContractAnzenServiceManager(serviceManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}

	taskManagerAddr, err := contractServiceManager.AnzenTaskManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}

	contractTaskManager, err := anzentaskmanager.NewContractAnzenTaskManager(taskManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IAnzenTaskManager contract", "err", err)
		return nil, err
	}

	contractSafetyFactorOracle, err := safetyfactororacle.NewContractSafetyFactorOracle(safetyFactorOracleAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch SafetyFactorOracle contract", "err", err)
		return nil, err
	}

	return &AvsManagersBindings{
		ServiceManager:     contractServiceManager,
		TaskManager:        contractTaskManager,
		SafetyFactorOracle: contractSafetyFactorOracle,
		ethClient:          ethclient,
		logger:             logger,
	}, nil
}

func (b *AvsManagersBindings) GetErc20Mock(tokenAddr common.Address) (*erc20mock.ContractERC20Mock, error) {
	contractErc20Mock, err := erc20mock.NewContractERC20Mock(tokenAddr, b.ethClient)
	if err != nil {
		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return contractErc20Mock, nil
}
