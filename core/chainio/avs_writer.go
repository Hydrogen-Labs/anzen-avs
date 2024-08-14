package chainio

import (
	"context"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	anzentaskmanager "anzen-avs/contracts/bindings/AnzenTaskManager"
	"anzen-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewOraclePullTask(
		ctx context.Context,
		oracleIndex *big.Int,
		proposedSafetyFactor *big.Int,
		quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
		quorumNumbers sdktypes.QuorumNums,
	) (anzentaskmanager.IAnzenTaskManagerOraclePullTask, uint32, error)
	SendAggregatedOracleResponse(ctx context.Context,
		task anzentaskmanager.IAnzenTaskManagerOraclePullTask,
		taskResponse anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse,
		nonSignerStakesAndSignature anzentaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.AnzenRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.SafetyFactorOracleAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, safetyFactorOracleAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, safetyFactorOracleAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

func (w *AvsWriter) SendNewOraclePullTask(
	ctx context.Context,
	oracleIndex *big.Int,
	proposedSafetyFactor *big.Int,
	quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
	quorumNumbers sdktypes.QuorumNums,
) (anzentaskmanager.IAnzenTaskManagerOraclePullTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return anzentaskmanager.IAnzenTaskManagerOraclePullTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateNewOraclePullTask(txOpts, uint32(oracleIndex.Uint64()), proposedSafetyFactor, uint32(quorumThresholdPercentage), quorumNumbers.UnderlyingType())
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewOraclePullTask tx")
		return anzentaskmanager.IAnzenTaskManagerOraclePullTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewOraclePullTask tx")
		return anzentaskmanager.IAnzenTaskManagerOraclePullTask{}, 0, err
	}
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractAnzenTaskManagerFilterer.ParseNewOraclePullTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new oracle pull task created event", "err", err)
		return anzentaskmanager.IAnzenTaskManagerOraclePullTask{}, 0, err
	}
	return newTaskCreatedEvent.OraclePullTask, newTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) SendAggregatedOracleResponse(
	ctx context.Context, task anzentaskmanager.IAnzenTaskManagerOraclePullTask,
	taskResponse anzentaskmanager.IAnzenTaskManagerOraclePullTaskResponse,
	nonSignerStakesAndSignature anzentaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RespondToOracleTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting respondToTask tx")
		return nil, err
	}
	return receipt, nil
}
