package chainio

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	cstaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	"anzen-avs/core/config"
)

type AvsSubscriberer interface {
	SubcribeToNewOraclePullTasks(newOraclePullTaskLogs chan *cstaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated) event.Subscription
}

// Subscribers use a ws connection instead of http connection like Readers
// kind of stupid that the geth client doesn't have a unified interface for both...
// it takes a single url, so the bindings, even though they have watcher functions, those can't be used
// with the http connection... seems very very stupid. Am I missing something?
type AvsSubscriber struct {
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

func BuildAvsSubscriberFromConfig(config *config.Config) (*AvsSubscriber, error) {
	return BuildAvsSubscriber(
		config.IncredibleSquaringRegistryCoordinatorAddr,
		config.OperatorStateRetrieverAddr,
		config.EthWsClient,
		config.Logger,
	)
}

func BuildAvsSubscriber(registryCoordinatorAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.Client, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, blsOperatorStateRetrieverAddr, ethclient, logger)
	if err != nil {
		logger.Errorf("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return NewAvsSubscriber(avsContractBindings, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
	}
}

func (s *AvsSubscriber) SubcribeToNewOraclePullTasks(newOraclePullTaskLogs chan *cstaskmanager.ContractAnzenTaskManagerNewOraclePullTaskCreated) event.Subscription {
	sub, err := s.AvsContractBindings.TaskManager.WatchNewOraclePullTaskCreated(
		&bind.WatchOpts{}, newOraclePullTaskLogs, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new OraclePullTask events", "err", err)
	}
	s.logger.Infof("Subscribed to new OraclePullTask events")
	return sub
}
