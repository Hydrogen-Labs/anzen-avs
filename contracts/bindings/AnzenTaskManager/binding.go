// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAnzenTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IAnzenTaskManagerOraclePullTask is an auto generated low-level Go binding around an user-defined struct.
type IAnzenTaskManagerOraclePullTask struct {
	OracleIndex               uint32
	ProposedSafetyFactor      *big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IAnzenTaskManagerOraclePullTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IAnzenTaskManagerOraclePullTaskResponse struct {
	ReferenceTaskIndex uint32
	SafetyFactor       *big.Int
}

// IAnzenTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IAnzenTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractAnzenTaskManagerMetaData contains all meta data concerning the ContractAnzenTaskManager contract.
var ContractAnzenTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewOraclePullTask\",\"inputs\":[{\"name\":\"oracleIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proposedSafetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToOracleTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIAnzenTaskManager.OraclePullTask\",\"components\":[{\"name\":\"oracleIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proposedSafetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIAnzenTaskManager.OraclePullTaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"safetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewOraclePullTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"oraclePullTask\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAnzenTaskManager.OraclePullTask\",\"components\":[{\"name\":\"oracleIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proposedSafetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OraclePullTaskResponded\",\"inputs\":[{\"name\":\"oraclePullTaskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAnzenTaskManager.OraclePullTaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"safetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAnzenTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200580f3803806200580f8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e0516101005161551f620002f0600039600081816102720152818161057e0152611877015260008181610547015261247c015260008181610413015261265e01526000818161043a0152818161283401526129f601526000818161046101528181610ded01528181612147015281816122df0152612519015261551f6000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80636830483511610125578063b98d0908116100ad578063f2fde38b1161007c578063f2fde38b14610569578063f5c9899d1461057c578063f63c5bab146105a2578063f8c8765e146105aa578063fabc1cbc146105bd57600080fd5b8063b98d090814610501578063cefdc1d41461050e578063d07dcfda1461052f578063df5cf7231461054257600080fd5b806372d18e8d116100f457806372d18e8d146104ac5780637afa1eed146104ba578063886f1195146104cd5780638b00ce7c146104e05780638da5cb5b146104f057600080fd5b806368304835146104355780636d14a9871461045c5780636efb463614610483578063715018a6146104a457600080fd5b8063416c7e5e116101a85780635ac86ab7116101775780635ac86ab7146103905780635c155662146103c35780635c975abb146103e35780635decc3f5146103eb5780635df459461461040e57600080fd5b8063416c7e5e146103425780634f739f741461035557806350f8cf7c14610375578063595c6a671461038857600080fd5b8063245a7bfc116101e4578063245a7bfc146102a95780632cb223d5146102d45780632d89f6fc146103025780633563b0d11461032257600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d575b600080fd5b6102296102243660046140d5565b6105d0565b005b6102296102393660046140f2565b61068c565b61025161024c366004614270565b6107cb565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b60cd546102bc906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102f46102e23660046142de565b60cb6020526000908152604090205481565b604051908152602001610264565b6102f46103103660046142de565b60ca6020526000908152604090205481565b6103356103303660046142fb565b610955565b6040516102649190614456565b61022961035036600461447e565b610deb565b6103686103633660046144e3565b610f60565b60405161026491906145e7565b610229610383366004614946565b611686565b610229611b05565b6103b361039e3660046149dd565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6103d66103d13660046149fa565b611bcc565b6040516102649190614aa6565b6066546102f4565b6103b36103f93660046142de565b60cc6020526000908152604090205460ff1681565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b610496610491366004614aea565b611d94565b604051610264929190614baa565b610229612cab565b60c95463ffffffff16610294565b60ce546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b60c9546102949063ffffffff1681565b6033546001600160a01b03166102bc565b6097546103b39060ff1681565b61052161051c366004614bf3565b612cbf565b604051610264929190614c35565b61022961053d366004614c56565b612e51565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102296105773660046140d5565b612fd7565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b610294606481565b6102296105b8366004614cc8565b61304d565b6102296105cb3660046140f2565b61319e565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610623573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106479190614d24565b6001600160a01b0316336001600160a01b0316146106805760405162461bcd60e51b815260040161067790614d41565b60405180910390fd5b610689816132fa565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f89190614d8b565b6107145760405162461bcd60e51b815260040161067790614da8565b6066548181161461078d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610677565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061081357610813614df0565b60200201518951600160200201518a6020015160006002811061083857610838614df0565b60200201518b6020015160016002811061085457610854614df0565b602090810291909101518c518d8301516040516108b19a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108d49190614e06565b90506109476108ed6108e688846133f1565b8690613488565b6108f561351c565b61093d61092e85610928604080518082018252600080825260209182015281518083019092526001825260029082015290565b906133f1565b6109378c6135dc565b90613488565b886201d4c061366c565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610997573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bb9190614d24565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a219190614d24565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a879190614d24565b9050600086516001600160401b03811115610aa457610aa461410b565b604051908082528060200260200182016040528015610ad757816020015b6060815260200190600190039081610ac25790505b50905060005b8751811015610ddf576000888281518110610afa57610afa614df0565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b5b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b839190810190614e28565b905080516001600160401b03811115610b9e57610b9e61410b565b604051908082528060200260200182016040528015610be957816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bbc5790505b50848481518110610bfc57610bfc614df0565b602002602001018190525060005b8151811015610dc9576040518060600160405280876001600160a01b03166347b314e8858581518110610c3f57610c3f614df0565b60200260200101516040518263ffffffff1660e01b8152600401610c6591815260200190565b602060405180830381865afa158015610c82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca69190614d24565b6001600160a01b03168152602001838381518110610cc657610cc6614df0565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610cf457610cf4614df0565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d749190614eb8565b6001600160601b0316815250858581518110610d9257610d92614df0565b60200260200101518281518110610dab57610dab614df0565b60200260200101819052508080610dc190614ef7565b915050610c0a565b5050508080610dd790614ef7565b915050610add565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6d9190614d24565b6001600160a01b0316336001600160a01b031614610f195760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610677565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610f8b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fcb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fef9190614d24565b905061101c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061104c908b9089908990600401614f12565b600060405180830381865afa158015611069573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110919190810190614f5c565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110c3908b908b908b90600401615013565b600060405180830381865afa1580156110e0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111089190810190614f5c565b6040820152856001600160401b038111156111255761112561410b565b60405190808252806020026020018201604052801561115857816020015b60608152602001906001900390816111435790505b50606082015260005b60ff8116871115611597576000856001600160401b038111156111865761118661410b565b6040519080825280602002602001820160405280156111af578160200160208202803683370190505b5083606001518360ff16815181106111c9576111c9614df0565b602002602001018190525060005b868110156114975760008c6001600160a01b03166304ec63518a8a8581811061120257611202614df0565b905060200201358e8860000151868151811061122057611220614df0565b60200260200101516040518463ffffffff1660e01b815260040161125d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561127a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129e919061503c565b90506001600160c01b0381166113425760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610677565b8a8a8560ff1681811061135757611357614df0565b6001600160c01b03841692013560f81c9190911c60019081161415905061148457856001600160a01b031663dd9846b98a8a8581811061139957611399614df0565b905060200201358d8d8860ff168181106113b5576113b5614df0565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561140b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142f9190615065565b85606001518560ff168151811061144857611448614df0565b6020026020010151848151811061146157611461614df0565b63ffffffff909216602092830291909101909101528261148081614ef7565b9350505b508061148f81614ef7565b9150506111d7565b506000816001600160401b038111156114b2576114b261410b565b6040519080825280602002602001820160405280156114db578160200160208202803683370190505b50905060005b8281101561155c5784606001518460ff168151811061150257611502614df0565b6020026020010151818151811061151b5761151b614df0565b602002602001015182828151811061153557611535614df0565b63ffffffff909216602092830291909101909101528061155481614ef7565b9150506114e1565b508084606001518460ff168151811061157757611577614df0565b60200260200101819052505050808061158f90615082565b915050611161565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115fc9190614d24565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061162f908b908b908e906004016150a2565b600060405180830381865afa15801561164c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116749190810190614f5c565b60208301525098975050505050505050565b60cd546001600160a01b031633146116e05760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610677565b60006116f260608501604086016142de565b905036600061170460608701876150cc565b9092509050600061171b60a08801608089016142de565b905060ca600061172e60208901896142de565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161175a9190615112565b60405160208183030381529060405280519060200120146117e35760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610677565b600060cb816117f560208a018a6142de565b63ffffffff1663ffffffff16815260200190815260200160002054146118725760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610677565b61189c7f0000000000000000000000000000000000000000000000000000000000000000856151cc565b63ffffffff164363ffffffff16111561190d5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610677565b6000866040516020016119209190615212565b6040516020818303038152906040528051906020012090506000806119488387878a8c611d94565b9150915060005b85811015611a47578460ff168360200151828151811061197157611971614df0565b60200260200101516119839190615220565b6001600160601b03166064846000015183815181106119a4576119a4614df0565b60200260200101516001600160601b03166119bf919061524f565b1015611a35576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610677565b80611a3f81614ef7565b91505061194f565b5060408051808201825263ffffffff43168152602080820184905291519091611a74918c9184910161526e565b6040516020818303038152906040528051906020012060cb60008c6000016020810190611aa191906142de565b63ffffffff1663ffffffff168152602001908152602001600020819055507f7ca1f1c5447c165db410c5518491c7d3591539160a4d3d093688ba26c4d40f3c8a82604051611af092919061526e565b60405180910390a15050505050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611b4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b719190614d8b565b611b8d5760405162461bcd60e51b815260040161067790614da8565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611bfe92919061529a565b600060405180830381865afa158015611c1b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c439190810190614f5c565b9050600084516001600160401b03811115611c6057611c6061410b565b604051908082528060200260200182016040528015611c89578160200160208202803683370190505b50905060005b8551811015611d8a57866001600160a01b03166304ec6351878381518110611cb957611cb9614df0565b602002602001015187868581518110611cd457611cd4614df0565b60200260200101516040518463ffffffff1660e01b8152600401611d119392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d52919061503c565b6001600160c01b0316828281518110611d6d57611d6d614df0565b602090810291909101015280611d8281614ef7565b915050611c8f565b5095945050505050565b6040805180820190915260608082526020820152600084611e0b5760405162461bcd60e51b815260206004820152603760248201526000805160206154ca83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610677565b60408301515185148015611e23575060a08301515185145b8015611e33575060c08301515185145b8015611e43575060e08301515185145b611ead5760405162461bcd60e51b815260206004820152604160248201526000805160206154ca83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610677565b82515160208401515114611f255760405162461bcd60e51b8152602060048201526044602482018190526000805160206154ca833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610677565b4363ffffffff168463ffffffff1610611f945760405162461bcd60e51b815260206004820152603c60248201526000805160206154ca83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610677565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611fd557611fd561410b565b604051908082528060200260200182016040528015611ffe578160200160208202803683370190505b506020820152866001600160401b0381111561201c5761201c61410b565b604051908082528060200260200182016040528015612045578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156120795761207961410b565b6040519080825280602002602001820160405280156120a2578160200160208202803683370190505b5081526020860151516001600160401b038111156120c2576120c261410b565b6040519080825280602002602001820160405280156120eb578160200160208202803683370190505b50816020018190525060006121bd8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612194573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b891906152ee565b613890565b905060005b87602001515181101561245857612207886020015182815181106121e8576121e8614df0565b6020026020010151805160009081526020918201519091526040902090565b8360200151828151811061221d5761221d614df0565b602090810291909101015280156122dd57602083015161223e60018361530b565b8151811061224e5761224e614df0565b602002602001015160001c8360200151828151811061226f5761226f614df0565b602002602001015160001c116122dd576040805162461bcd60e51b81526020600482015260248101919091526000805160206154ca83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610677565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061232257612322614df0565b60200260200101518b8b60000151858151811061234157612341614df0565b60200260200101516040518463ffffffff1660e01b815260040161237e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561239b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123bf919061503c565b6001600160c01b0316836000015182815181106123de576123de614df0565b6020026020010181815250506124446108e6612418848660000151858151811061240a5761240a614df0565b602002602001015116613923565b8a60200151848151811061242e5761242e614df0565b602002602001015161394e90919063ffffffff16565b94508061245081614ef7565b9150506121c2565b505061246383613a32565b60975490935060ff1660008161247a5760006124fc565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124fc9190615322565b905060005b8a811015612b7a57821561265c578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061255857612558614df0565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612598573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125bc9190615322565b6125c6919061533b565b1161265c5760405162461bcd60e51b815260206004820152606660248201526000805160206154ca83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610677565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061269d5761269d614df0565b9050013560f81c60f81b60f81c8c8c60a0015185815181106126c1576126c1614df0565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561271d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127419190615353565b6001600160401b0319166127648a6040015183815181106121e8576121e8614df0565b67ffffffffffffffff1916146128005760405162461bcd60e51b815260206004820152606160248201526000805160206154ca83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610677565b6128308960400151828151811061281957612819614df0565b60200260200101518761348890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061287357612873614df0565b9050013560f81c60f81b60f81c8c8c60c00151858151811061289757612897614df0565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156128f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129179190614eb8565b8560200151828151811061292d5761292d614df0565b6001600160601b0390921660209283029190910182015285015180518290811061295957612959614df0565b60200260200101518560000151828151811061297757612977614df0565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612b65576129ef866000015182815181106129c1576129c1614df0565b60200260200101518f8f868181106129db576129db614df0565b600192013560f81c9290921c811614919050565b15612b53577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612a3557612a35614df0565b9050013560f81c60f81b60f81c8e89602001518581518110612a5957612a59614df0565b60200260200101518f60e001518881518110612a7757612a77614df0565b60200260200101518781518110612a9057612a90614df0565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612af4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b189190614eb8565b8751805185908110612b2c57612b2c614df0565b60200260200101818151612b40919061537e565b6001600160601b03169052506001909101905b80612b5d81614ef7565b91505061299b565b50508080612b7290614ef7565b915050612501565b505050600080612b948c868a606001518b608001516107cb565b9150915081612c055760405162461bcd60e51b815260206004820152604360248201526000805160206154ca83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610677565b80612c665760405162461bcd60e51b815260206004820152603960248201526000805160206154ca83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610677565b50506000878260200151604051602001612c819291906153a6565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612cb3613acd565b612cbd6000613b27565b565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612cfa57612cfa614df0565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612d36908890869060040161529a565b600060405180830381865afa158015612d53573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612d7b9190810190614f5c565b600081518110612d8d57612d8d614df0565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612df9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e1d919061503c565b6001600160c01b031690506000612e3382613b79565b905081612e418a838a610955565b9550955050505050935093915050565b60ce546001600160a01b03163314612eb55760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610677565b6040805160a081018252600060208083019190915260608083015263ffffffff888116835243811683850152861660808301528251601f8501829004820281018201909352838352909190849084908190840183828082843760009201919091525050505060608201526020808201869052604051612f36918391016153ee565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f6dd4fcaaac535ef63b8629e190e6551084203c3740f0ad282bee1e9af6e4909090612f999084906153ee565b60405180910390a260c954612fb59063ffffffff1660016151cc565b60c9805463ffffffff191663ffffffff92909216919091179055505050505050565b612fdf613acd565b6001600160a01b0381166130445760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610677565b61068981613b27565b600054610100900460ff161580801561306d5750600054600160ff909116105b806130875750303b158015613087575060005460ff166001145b6130ea5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610677565b6000805460ff19166001179055801561310d576000805461ff0019166101001790555b613118856000613c45565b61312184613b27565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce8054928516929091169190911790558015613197576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156131f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132159190614d24565b6001600160a01b0316336001600160a01b0316146132455760405162461bcd60e51b815260040161067790614d41565b6066541981196066541916146132c35760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610677565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107c0565b6001600160a01b0381166133885760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610677565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261340d613fe6565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561344057613442565bfe5b50806134805760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610677565b505092915050565b60408051808201909152600080825260208201526134a4614004565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156134405750806134805760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610677565b613524614022565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061360c6000805160206154aa83398151915286614e06565b90505b61361881613d2f565b90935091506000805160206154aa833981519152828309831415613652576040805180820190915290815260208101919091529392505050565b6000805160206154aa83398151915260018208905061360f565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061369e614047565b60005b60028110156138635760006136b782600661524f565b90508482600281106136cb576136cb614df0565b602002015151836136dd83600061533b565b600c81106136ed576136ed614df0565b602002015284826002811061370457613704614df0565b6020020151602001518382600161371b919061533b565b600c811061372b5761372b614df0565b602002015283826002811061374257613742614df0565b602002015151518361375583600261533b565b600c811061376557613765614df0565b602002015283826002811061377c5761377c614df0565b602002015151600160200201518361379583600361533b565b600c81106137a5576137a5614df0565b60200201528382600281106137bc576137bc614df0565b6020020151602001516000600281106137d7576137d7614df0565b6020020151836137e883600461533b565b600c81106137f8576137f8614df0565b602002015283826002811061380f5761380f614df0565b60200201516020015160016002811061382a5761382a614df0565b60200201518361383b83600561533b565b600c811061384b5761384b614df0565b6020020152508061385b81614ef7565b9150506136a1565b5061386c614066565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008061389c84613db1565b9050808360ff166001901b1161391a5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610677565b90505b92915050565b6000805b821561391d5761393860018461530b565b909216918061394681615487565b915050613927565b60408051808201909152600080825260208201526102008261ffff16106139aa5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610677565b8161ffff16600114156139be57508161391d565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613a2757600161ffff871660ff83161c81161415613a0a57613a078484613488565b93505b613a148384613488565b92506201fffe600192831b1691016139da565b509195945050505050565b60408051808201909152600080825260208201528151158015613a5757506020820151155b15613a75575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206154aa8339815191528460200151613aa89190614e06565b613ac0906000805160206154aa83398151915261530b565b905292915050565b919050565b6033546001600160a01b03163314612cbd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610677565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613b8784613923565b61ffff166001600160401b03811115613ba257613ba261410b565b6040519080825280601f01601f191660200182016040528015613bcc576020820181803683370190505b5090506000805b825182108015613be4575061010081105b15613c3b576001811b935085841615613c2b578060f81b838381518110613c0d57613c0d614df0565b60200101906001600160f81b031916908160001a9053508160010191505b613c3481614ef7565b9050613bd3565b5090949350505050565b6065546001600160a01b0316158015613c6657506001600160a01b03821615155b613ce85760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610677565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613d2b826132fa565b5050565b600080806000805160206154aa83398151915260036000805160206154aa833981519152866000805160206154aa833981519152888909090890506000613da5827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206154aa833981519152613f3e565b91959194509092505050565b600061010082511115613e3a5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610677565b8151613e4857506000919050565b60008083600081518110613e5e57613e5e614df0565b0160200151600160f89190911c81901b92505b8451811015613f3557848181518110613e8c57613e8c614df0565b0160200151600160f89190911c1b9150828211613f215760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610677565b91811791613f2e81614ef7565b9050613e71565b50909392505050565b600080613f49614066565b613f51614084565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613440575082613fdb5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610677565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806140356140a2565b81526020016140426140a2565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461068957600080fd5b6000602082840312156140e757600080fd5b813561391a816140c0565b60006020828403121561410457600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156141435761414361410b565b60405290565b60405161010081016001600160401b03811182821017156141435761414361410b565b604051601f8201601f191681016001600160401b03811182821017156141945761419461410b565b604052919050565b6000604082840312156141ae57600080fd5b6141b6614121565b9050813581526020820135602082015292915050565b600082601f8301126141dd57600080fd5b604051604081018181106001600160401b03821117156141ff576141ff61410b565b806040525080604084018581111561421657600080fd5b845b81811015613a27578035835260209283019201614218565b60006080828403121561424257600080fd5b61424a614121565b905061425683836141cc565b815261426583604084016141cc565b602082015292915050565b600080600080610120858703121561428757600080fd5b84359350614298866020870161419c565b92506142a78660608701614230565b91506142b68660e0870161419c565b905092959194509250565b63ffffffff8116811461068957600080fd5b8035613ac8816142c1565b6000602082840312156142f057600080fd5b813561391a816142c1565b60008060006060848603121561431057600080fd5b833561431b816140c0565b92506020848101356001600160401b038082111561433857600080fd5b818701915087601f83011261434c57600080fd5b81358181111561435e5761435e61410b565b614370601f8201601f1916850161416c565b9150808252888482850101111561438657600080fd5b80848401858401376000848284010152508094505050506143a9604085016142d3565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614448578385038a52825180518087529087019087870190845b8181101561443357835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016143ef565b50509a87019a955050918501916001016143d1565b509298975050505050505050565b60208152600061446960208301846143b2565b9392505050565b801515811461068957600080fd5b60006020828403121561449057600080fd5b813561391a81614470565b60008083601f8401126144ad57600080fd5b5081356001600160401b038111156144c457600080fd5b6020830191508360208285010111156144dc57600080fd5b9250929050565b600080600080600080608087890312156144fc57600080fd5b8635614507816140c0565b95506020870135614517816142c1565b945060408701356001600160401b038082111561453357600080fd5b61453f8a838b0161449b565b9096509450606089013591508082111561455857600080fd5b818901915089601f83011261456c57600080fd5b81358181111561457b57600080fd5b8a60208260051b850101111561459057600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156145dc57815163ffffffff16875295820195908201906001016145ba565b509495945050505050565b60006020808352835160808285015261460360a08501826145a6565b905081850151601f198086840301604087015261462083836145a6565b9250604087015191508086840301606087015261463d83836145a6565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561469457848783030184526146828287516145a6565b95880195938801939150600101614668565b509998505050505050505050565b60006001600160401b038211156146bb576146bb61410b565b5060051b60200190565b600082601f8301126146d657600080fd5b813560206146eb6146e6836146a2565b61416c565b82815260059290921b8401810191818101908684111561470a57600080fd5b8286015b8481101561472e578035614721816142c1565b835291830191830161470e565b509695505050505050565b600082601f83011261474a57600080fd5b8135602061475a6146e6836146a2565b82815260069290921b8401810191818101908684111561477957600080fd5b8286015b8481101561472e5761478f888261419c565b83529183019160400161477d565b600082601f8301126147ae57600080fd5b813560206147be6146e6836146a2565b82815260059290921b840181019181810190868411156147dd57600080fd5b8286015b8481101561472e5780356001600160401b038111156148005760008081fd5b61480e8986838b01016146c5565b8452509183019183016147e1565b6000610180828403121561482f57600080fd5b614837614149565b905081356001600160401b038082111561485057600080fd5b61485c858386016146c5565b8352602084013591508082111561487257600080fd5b61487e85838601614739565b6020840152604084013591508082111561489757600080fd5b6148a385838601614739565b60408401526148b58560608601614230565b60608401526148c78560e0860161419c565b60808401526101208401359150808211156148e157600080fd5b6148ed858386016146c5565b60a084015261014084013591508082111561490757600080fd5b614913858386016146c5565b60c084015261016084013591508082111561492d57600080fd5b5061493a8482850161479d565b60e08301525092915050565b6000806000838503608081121561495c57600080fd5b84356001600160401b038082111561497357600080fd5b9086019060a0828903121561498757600080fd5b8195506040601f198401121561499c57600080fd5b60208701945060608701359250808311156149b657600080fd5b50506149c48682870161481c565b9150509250925092565b60ff8116811461068957600080fd5b6000602082840312156149ef57600080fd5b813561391a816149ce565b600080600060608486031215614a0f57600080fd5b8335614a1a816140c0565b92506020848101356001600160401b03811115614a3657600080fd5b8501601f81018713614a4757600080fd5b8035614a556146e6826146a2565b81815260059190911b82018301908381019089831115614a7457600080fd5b928401925b82841015614a9257833582529284019290840190614a79565b80965050505050506143a9604085016142d3565b6020808252825182820181905260009190848201906040850190845b81811015614ade57835183529284019291840191600101614ac2565b50909695505050505050565b600080600080600060808688031215614b0257600080fd5b8535945060208601356001600160401b0380821115614b2057600080fd5b614b2c89838a0161449b565b909650945060408801359150614b41826142c1565b90925060608701359080821115614b5757600080fd5b50614b648882890161481c565b9150509295509295909350565b600081518084526020808501945080840160005b838110156145dc5781516001600160601b031687529582019590820190600101614b85565b6040815260008351604080840152614bc56080840182614b71565b90506020850151603f19848303016060850152614be28282614b71565b925050508260208301529392505050565b600080600060608486031215614c0857600080fd5b8335614c13816140c0565b9250602084013591506040840135614c2a816142c1565b809150509250925092565b828152604060208201526000614c4e60408301846143b2565b949350505050565b600080600080600060808688031215614c6e57600080fd5b8535614c79816142c1565b9450602086013593506040860135614c90816142c1565b925060608601356001600160401b03811115614cab57600080fd5b614cb78882890161449b565b969995985093965092949392505050565b60008060008060808587031215614cde57600080fd5b8435614ce9816140c0565b93506020850135614cf9816140c0565b92506040850135614d09816140c0565b91506060850135614d19816140c0565b939692955090935050565b600060208284031215614d3657600080fd5b815161391a816140c0565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614d9d57600080fd5b815161391a81614470565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614e2357634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614e3b57600080fd5b82516001600160401b03811115614e5157600080fd5b8301601f81018513614e6257600080fd5b8051614e706146e6826146a2565b81815260059190911b82018301908381019087831115614e8f57600080fd5b928401925b82841015614ead57835182529284019290840190614e94565b979650505050505050565b600060208284031215614eca57600080fd5b81516001600160601b038116811461391a57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415614f0b57614f0b614ee1565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b03831115614f3f57600080fd5b8260051b8085606085013760009201606001918252509392505050565b60006020808385031215614f6f57600080fd5b82516001600160401b03811115614f8557600080fd5b8301601f81018513614f9657600080fd5b8051614fa46146e6826146a2565b81815260059190911b82018301908381019087831115614fc357600080fd5b928401925b82841015614ead578351614fdb816142c1565b82529284019290840190614fc8565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615033604083018486614fea565b95945050505050565b60006020828403121561504e57600080fd5b81516001600160c01b038116811461391a57600080fd5b60006020828403121561507757600080fd5b815161391a816142c1565b600060ff821660ff81141561509957615099614ee1565b60010192915050565b6040815260006150b6604083018587614fea565b905063ffffffff83166020830152949350505050565b6000808335601e198436030181126150e357600080fd5b8301803591506001600160401b038211156150fd57600080fd5b6020019150368190038213156144dc57600080fd5b6020815260008235615123816142c1565b63ffffffff8082166020850152602085013560408501526040850135915061514a826142c1565b1660608381019190915283013536849003601e1901811261516a57600080fd5b830180356001600160401b0381111561518257600080fd5b80360385131561519157600080fd5b60a060808501526151a960c085018260208501614fea565b9150506151b8608085016142d3565b63ffffffff811660a0850152509392505050565b600063ffffffff8083168185168083038211156151eb576151eb614ee1565b01949350505050565b80356151ff816142c1565b63ffffffff168252602090810135910152565b6040810161391d82846151f4565b60006001600160601b038083168185168183048111821515161561524657615246614ee1565b02949350505050565b600081600019048311821515161561526957615269614ee1565b500290565b6080810161527c82856151f4565b63ffffffff8351166040830152602083015160608301529392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156152e1578451835293830193918301916001016152c5565b5090979650505050505050565b60006020828403121561530057600080fd5b815161391a816149ce565b60008282101561531d5761531d614ee1565b500390565b60006020828403121561533457600080fd5b5051919050565b6000821982111561534e5761534e614ee1565b500190565b60006020828403121561536557600080fd5b815167ffffffffffffffff198116811461391a57600080fd5b60006001600160601b038381169083168181101561539e5761539e614ee1565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156153e1578151855293820193908201906001016153c5565b5092979650505050505050565b6000602080835263ffffffff808551168285015281850151604085015280604086015116606085015250606084015160a0608085015280518060c086015260005b8181101561544b5782810184015186820160e00152830161542f565b8181111561545d57600060e083880101525b50608086015163ffffffff811660a08701529250601f01601f19169390930160e001949350505050565b600061ffff8083168181141561549f5761549f614ee1565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220f1fc04cbbd7fa9c07eeed01761e87ab6e131f5fd2181a1c6e2131fccee7b9d6f64736f6c634300080c0033",
}

// ContractAnzenTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAnzenTaskManagerMetaData.ABI instead.
var ContractAnzenTaskManagerABI = ContractAnzenTaskManagerMetaData.ABI

// ContractAnzenTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAnzenTaskManagerMetaData.Bin instead.
var ContractAnzenTaskManagerBin = ContractAnzenTaskManagerMetaData.Bin

// DeployContractAnzenTaskManager deploys a new Ethereum contract, binding an instance of ContractAnzenTaskManager to it.
func DeployContractAnzenTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractAnzenTaskManager, error) {
	parsed, err := ContractAnzenTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAnzenTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAnzenTaskManager{ContractAnzenTaskManagerCaller: ContractAnzenTaskManagerCaller{contract: contract}, ContractAnzenTaskManagerTransactor: ContractAnzenTaskManagerTransactor{contract: contract}, ContractAnzenTaskManagerFilterer: ContractAnzenTaskManagerFilterer{contract: contract}}, nil
}

// ContractAnzenTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractAnzenTaskManager struct {
	ContractAnzenTaskManagerCaller     // Read-only binding to the contract
	ContractAnzenTaskManagerTransactor // Write-only binding to the contract
	ContractAnzenTaskManagerFilterer   // Log filterer for contract events
}

// ContractAnzenTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAnzenTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAnzenTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAnzenTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAnzenTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAnzenTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAnzenTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAnzenTaskManagerSession struct {
	Contract     *ContractAnzenTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractAnzenTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAnzenTaskManagerCallerSession struct {
	Contract *ContractAnzenTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractAnzenTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAnzenTaskManagerTransactorSession struct {
	Contract     *ContractAnzenTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractAnzenTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAnzenTaskManagerRaw struct {
	Contract *ContractAnzenTaskManager // Generic contract binding to access the raw methods on
}

// ContractAnzenTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAnzenTaskManagerCallerRaw struct {
	Contract *ContractAnzenTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAnzenTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAnzenTaskManagerTransactorRaw struct {
	Contract *ContractAnzenTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAnzenTaskManager creates a new instance of ContractAnzenTaskManager, bound to a specific deployed contract.
func NewContractAnzenTaskManager(address common.Address, backend bind.ContractBackend) (*ContractAnzenTaskManager, error) {
	contract, err := bindContractAnzenTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManager{ContractAnzenTaskManagerCaller: ContractAnzenTaskManagerCaller{contract: contract}, ContractAnzenTaskManagerTransactor: ContractAnzenTaskManagerTransactor{contract: contract}, ContractAnzenTaskManagerFilterer: ContractAnzenTaskManagerFilterer{contract: contract}}, nil
}

// NewContractAnzenTaskManagerCaller creates a new read-only instance of ContractAnzenTaskManager, bound to a specific deployed contract.
func NewContractAnzenTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractAnzenTaskManagerCaller, error) {
	contract, err := bindContractAnzenTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerCaller{contract: contract}, nil
}

// NewContractAnzenTaskManagerTransactor creates a new write-only instance of ContractAnzenTaskManager, bound to a specific deployed contract.
func NewContractAnzenTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAnzenTaskManagerTransactor, error) {
	contract, err := bindContractAnzenTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerTransactor{contract: contract}, nil
}

// NewContractAnzenTaskManagerFilterer creates a new log filterer instance of ContractAnzenTaskManager, bound to a specific deployed contract.
func NewContractAnzenTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAnzenTaskManagerFilterer, error) {
	contract, err := bindContractAnzenTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerFilterer{contract: contract}, nil
}

// bindContractAnzenTaskManager binds a generic wrapper to an already deployed contract.
func bindContractAnzenTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAnzenTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAnzenTaskManager.Contract.ContractAnzenTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.ContractAnzenTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.ContractAnzenTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAnzenTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractAnzenTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractAnzenTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractAnzenTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractAnzenTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Aggregator(&_ContractAnzenTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Aggregator(&_ContractAnzenTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractAnzenTaskManager.Contract.AllTaskHashes(&_ContractAnzenTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractAnzenTaskManager.Contract.AllTaskHashes(&_ContractAnzenTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractAnzenTaskManager.Contract.AllTaskResponses(&_ContractAnzenTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractAnzenTaskManager.Contract.AllTaskResponses(&_ContractAnzenTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.BlsApkRegistry(&_ContractAnzenTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.BlsApkRegistry(&_ContractAnzenTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractAnzenTaskManager.Contract.CheckSignatures(&_ContractAnzenTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractAnzenTaskManager.Contract.CheckSignatures(&_ContractAnzenTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Delegation(&_ContractAnzenTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Delegation(&_ContractAnzenTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Generator() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Generator(&_ContractAnzenTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Generator(&_ContractAnzenTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractAnzenTaskManager.Contract.GetCheckSignaturesIndices(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractAnzenTaskManager.Contract.GetCheckSignaturesIndices(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractAnzenTaskManager.Contract.GetOperatorState(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractAnzenTaskManager.Contract.GetOperatorState(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractAnzenTaskManager.Contract.GetOperatorState0(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractAnzenTaskManager.Contract.GetOperatorState0(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractAnzenTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractAnzenTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractAnzenTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractAnzenTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractAnzenTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.LatestTaskNum(&_ContractAnzenTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.LatestTaskNum(&_ContractAnzenTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Owner() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Owner(&_ContractAnzenTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.Owner(&_ContractAnzenTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractAnzenTaskManager.Contract.Paused(&_ContractAnzenTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAnzenTaskManager.Contract.Paused(&_ContractAnzenTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractAnzenTaskManager.Contract.Paused0(&_ContractAnzenTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractAnzenTaskManager.Contract.Paused0(&_ContractAnzenTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.PauserRegistry(&_ContractAnzenTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.PauserRegistry(&_ContractAnzenTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.RegistryCoordinator(&_ContractAnzenTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.RegistryCoordinator(&_ContractAnzenTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.StakeRegistry(&_ContractAnzenTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractAnzenTaskManager.Contract.StakeRegistry(&_ContractAnzenTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractAnzenTaskManager.Contract.StaleStakesForbidden(&_ContractAnzenTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractAnzenTaskManager.Contract.StaleStakesForbidden(&_ContractAnzenTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.TaskNumber(&_ContractAnzenTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractAnzenTaskManager.Contract.TaskNumber(&_ContractAnzenTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAnzenTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractAnzenTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractAnzenTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractAnzenTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractAnzenTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractAnzenTaskManager.Contract.TrySignatureAndApkVerification(&_ContractAnzenTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractAnzenTaskManager.Contract.TrySignatureAndApkVerification(&_ContractAnzenTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewOraclePullTask is a paid mutator transaction binding the contract method 0xd07dcfda.
//
// Solidity: function createNewOraclePullTask(uint32 oracleIndex, int256 proposedSafetyFactor, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) CreateNewOraclePullTask(opts *bind.TransactOpts, oracleIndex uint32, proposedSafetyFactor *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "createNewOraclePullTask", oracleIndex, proposedSafetyFactor, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewOraclePullTask is a paid mutator transaction binding the contract method 0xd07dcfda.
//
// Solidity: function createNewOraclePullTask(uint32 oracleIndex, int256 proposedSafetyFactor, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) CreateNewOraclePullTask(oracleIndex uint32, proposedSafetyFactor *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.CreateNewOraclePullTask(&_ContractAnzenTaskManager.TransactOpts, oracleIndex, proposedSafetyFactor, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewOraclePullTask is a paid mutator transaction binding the contract method 0xd07dcfda.
//
// Solidity: function createNewOraclePullTask(uint32 oracleIndex, int256 proposedSafetyFactor, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) CreateNewOraclePullTask(oracleIndex uint32, proposedSafetyFactor *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.CreateNewOraclePullTask(&_ContractAnzenTaskManager.TransactOpts, oracleIndex, proposedSafetyFactor, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.Initialize(&_ContractAnzenTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.Initialize(&_ContractAnzenTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.Pause(&_ContractAnzenTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.Pause(&_ContractAnzenTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.PauseAll(&_ContractAnzenTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.PauseAll(&_ContractAnzenTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.RenounceOwnership(&_ContractAnzenTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.RenounceOwnership(&_ContractAnzenTaskManager.TransactOpts)
}

// RespondToOracleTask is a paid mutator transaction binding the contract method 0x50f8cf7c.
//
// Solidity: function respondToOracleTask((uint32,int256,uint32,bytes,uint32) task, (uint32,int256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) RespondToOracleTask(opts *bind.TransactOpts, task IAnzenTaskManagerOraclePullTask, taskResponse IAnzenTaskManagerOraclePullTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "respondToOracleTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToOracleTask is a paid mutator transaction binding the contract method 0x50f8cf7c.
//
// Solidity: function respondToOracleTask((uint32,int256,uint32,bytes,uint32) task, (uint32,int256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) RespondToOracleTask(task IAnzenTaskManagerOraclePullTask, taskResponse IAnzenTaskManagerOraclePullTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.RespondToOracleTask(&_ContractAnzenTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToOracleTask is a paid mutator transaction binding the contract method 0x50f8cf7c.
//
// Solidity: function respondToOracleTask((uint32,int256,uint32,bytes,uint32) task, (uint32,int256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) RespondToOracleTask(task IAnzenTaskManagerOraclePullTask, taskResponse IAnzenTaskManagerOraclePullTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.RespondToOracleTask(&_ContractAnzenTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.SetPauserRegistry(&_ContractAnzenTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.SetPauserRegistry(&_ContractAnzenTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.SetStaleStakesForbidden(&_ContractAnzenTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.SetStaleStakesForbidden(&_ContractAnzenTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.TransferOwnership(&_ContractAnzenTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.TransferOwnership(&_ContractAnzenTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.Unpause(&_ContractAnzenTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAnzenTaskManager.Contract.Unpause(&_ContractAnzenTaskManager.TransactOpts, newPausedStatus)
}

// ContractAnzenTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerInitializedIterator struct {
	Event *ContractAnzenTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerInitialized represents a Initialized event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAnzenTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerInitializedIterator{contract: _ContractAnzenTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerInitialized)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractAnzenTaskManagerInitialized, error) {
	event := new(ContractAnzenTaskManagerInitialized)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator is returned from FilterNewOraclePullTaskCreated and is used to iterate over the raw logs and unpacked data for NewOraclePullTaskCreated events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator struct {
	Event *ContractAnzenTaskManagerNewOraclePullTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerNewOraclePullTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerNewOraclePullTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerNewOraclePullTaskCreated represents a NewOraclePullTaskCreated event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerNewOraclePullTaskCreated struct {
	TaskIndex      uint32
	OraclePullTask IAnzenTaskManagerOraclePullTask
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewOraclePullTaskCreated is a free log retrieval operation binding the contract event 0x6dd4fcaaac535ef63b8629e190e6551084203c3740f0ad282bee1e9af6e49090.
//
// Solidity: event NewOraclePullTaskCreated(uint32 indexed taskIndex, (uint32,int256,uint32,bytes,uint32) oraclePullTask)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterNewOraclePullTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "NewOraclePullTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerNewOraclePullTaskCreatedIterator{contract: _ContractAnzenTaskManager.contract, event: "NewOraclePullTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewOraclePullTaskCreated is a free log subscription operation binding the contract event 0x6dd4fcaaac535ef63b8629e190e6551084203c3740f0ad282bee1e9af6e49090.
//
// Solidity: event NewOraclePullTaskCreated(uint32 indexed taskIndex, (uint32,int256,uint32,bytes,uint32) oraclePullTask)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchNewOraclePullTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerNewOraclePullTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "NewOraclePullTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerNewOraclePullTaskCreated)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "NewOraclePullTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOraclePullTaskCreated is a log parse operation binding the contract event 0x6dd4fcaaac535ef63b8629e190e6551084203c3740f0ad282bee1e9af6e49090.
//
// Solidity: event NewOraclePullTaskCreated(uint32 indexed taskIndex, (uint32,int256,uint32,bytes,uint32) oraclePullTask)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseNewOraclePullTaskCreated(log types.Log) (*ContractAnzenTaskManagerNewOraclePullTaskCreated, error) {
	event := new(ContractAnzenTaskManagerNewOraclePullTaskCreated)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "NewOraclePullTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerOraclePullTaskRespondedIterator is returned from FilterOraclePullTaskResponded and is used to iterate over the raw logs and unpacked data for OraclePullTaskResponded events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerOraclePullTaskRespondedIterator struct {
	Event *ContractAnzenTaskManagerOraclePullTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerOraclePullTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerOraclePullTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerOraclePullTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerOraclePullTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerOraclePullTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerOraclePullTaskResponded represents a OraclePullTaskResponded event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerOraclePullTaskResponded struct {
	OraclePullTaskResponse IAnzenTaskManagerOraclePullTaskResponse
	TaskResponseMetadata   IAnzenTaskManagerTaskResponseMetadata
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOraclePullTaskResponded is a free log retrieval operation binding the contract event 0x7ca1f1c5447c165db410c5518491c7d3591539160a4d3d093688ba26c4d40f3c.
//
// Solidity: event OraclePullTaskResponded((uint32,int256) oraclePullTaskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterOraclePullTaskResponded(opts *bind.FilterOpts) (*ContractAnzenTaskManagerOraclePullTaskRespondedIterator, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "OraclePullTaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerOraclePullTaskRespondedIterator{contract: _ContractAnzenTaskManager.contract, event: "OraclePullTaskResponded", logs: logs, sub: sub}, nil
}

// WatchOraclePullTaskResponded is a free log subscription operation binding the contract event 0x7ca1f1c5447c165db410c5518491c7d3591539160a4d3d093688ba26c4d40f3c.
//
// Solidity: event OraclePullTaskResponded((uint32,int256) oraclePullTaskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchOraclePullTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerOraclePullTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "OraclePullTaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerOraclePullTaskResponded)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "OraclePullTaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOraclePullTaskResponded is a log parse operation binding the contract event 0x7ca1f1c5447c165db410c5518491c7d3591539160a4d3d093688ba26c4d40f3c.
//
// Solidity: event OraclePullTaskResponded((uint32,int256) oraclePullTaskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseOraclePullTaskResponded(log types.Log) (*ContractAnzenTaskManagerOraclePullTaskResponded, error) {
	event := new(ContractAnzenTaskManagerOraclePullTaskResponded)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "OraclePullTaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerOwnershipTransferredIterator struct {
	Event *ContractAnzenTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAnzenTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerOwnershipTransferredIterator{contract: _ContractAnzenTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerOwnershipTransferred)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAnzenTaskManagerOwnershipTransferred, error) {
	event := new(ContractAnzenTaskManagerOwnershipTransferred)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerPausedIterator struct {
	Event *ContractAnzenTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerPaused represents a Paused event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAnzenTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerPausedIterator{contract: _ContractAnzenTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerPaused)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParsePaused(log types.Log) (*ContractAnzenTaskManagerPaused, error) {
	event := new(ContractAnzenTaskManagerPaused)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerPauserRegistrySetIterator struct {
	Event *ContractAnzenTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractAnzenTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerPauserRegistrySetIterator{contract: _ContractAnzenTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerPauserRegistrySet)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractAnzenTaskManagerPauserRegistrySet, error) {
	event := new(ContractAnzenTaskManagerPauserRegistrySet)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractAnzenTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractAnzenTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractAnzenTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractAnzenTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractAnzenTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractAnzenTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerTaskChallengedSuccessfully)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractAnzenTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractAnzenTaskManagerTaskChallengedSuccessfully)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractAnzenTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractAnzenTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractAnzenTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractAnzenTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerTaskCompletedIterator struct {
	Event *ContractAnzenTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractAnzenTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerTaskCompletedIterator{contract: _ContractAnzenTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerTaskCompleted)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractAnzenTaskManagerTaskCompleted, error) {
	event := new(ContractAnzenTaskManagerTaskCompleted)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerUnpausedIterator struct {
	Event *ContractAnzenTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAnzenTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAnzenTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAnzenTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenTaskManagerUnpaused represents a Unpaused event raised by the ContractAnzenTaskManager contract.
type ContractAnzenTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAnzenTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenTaskManagerUnpausedIterator{contract: _ContractAnzenTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAnzenTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAnzenTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenTaskManagerUnpaused)
				if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAnzenTaskManager *ContractAnzenTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractAnzenTaskManagerUnpaused, error) {
	event := new(ContractAnzenTaskManagerUnpaused)
	if err := _ContractAnzenTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
