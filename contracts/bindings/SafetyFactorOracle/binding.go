// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSafetyFactorOracle

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

// SafetyFactorSnapshot is an auto generated low-level Go binding around an user-defined struct.
type SafetyFactorSnapshot struct {
	SafetyFactor *big.Int
	Timestamp    *big.Int
}

// ContractSafetyFactorOracleMetaData contains all meta data concerning the ContractSafetyFactorOracle contract.
var ContractSafetyFactorOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeProtocols\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addProtocol\",\"inputs\":[{\"name\":\"_protocolIdToAdd\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_reservesManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anzenGov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"anzenTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsReservesManagerFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fallBackSafetyFactorPoster\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDisputeStatus\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSafetyFactor\",\"inputs\":[{\"name\":\"_protocolId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"safetyFactor\",\"type\":\"tuple\",\"internalType\":\"structSafetyFactorSnapshot\",\"components\":[{\"name\":\"safetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_anzenTaskManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_anzenGov\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fallBackSafetyFactorPoster\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingDispute\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolToReservesManager\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeProtocol\",\"inputs\":[{\"name\":\"_protocolIdToRemove\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safetyFactorSnapshots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"safetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDisputeStatus\",\"inputs\":[{\"name\":\"_status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSafetyFactor\",\"inputs\":[{\"name\":\"_protocolId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_safetyFactor\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFallbackPoster\",\"inputs\":[{\"name\":\"newFallBackSafetyFactorPoster\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DisputeStatusSet\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolAdded\",\"inputs\":[{\"name\":\"protocolId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"reservesManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolRemoved\",\"inputs\":[{\"name\":\"protocolId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SFUpdated\",\"inputs\":[{\"name\":\"protocolId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newSF\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50610a6d806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806378e3cdaf11610097578063b44eb60911610066578063b44eb6091461025a578063bb6cb0dd1461026d578063c0c53b8b146102a9578063d6ae724f146102bc57600080fd5b806378e3cdaf1461020d57806380bd9f33146102205780639312afa814610234578063a3e453111461024757600080fd5b80633cac45df116100d35780633cac45df146101a1578063557d0f64146101d45780635bf9c96f146101e75780636b8731a4146101fa57600080fd5b8063111ab8911461010557806327fdf8a614610138578063378982461461014d57806338422fc41461018e575b600080fd5b6101186101133660046108bc565b6102ce565b604080518251815260209283015192810192909252015b60405180910390f35b61014b6101463660046108de565b610350565b005b61017661015b3660046108bc565b6005602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161012f565b61014b61019c366004610908565b6104eb565b6101c46101af3660046108bc565b60066020526000908152604090205460ff1681565b604051901515815260200161012f565b61014b6101e23660046108bc565b610543565b600254610176906001600160a01b031681565b600154610176906001600160a01b031681565b600354610176906001600160a01b031681565b6003546101c490600160a01b900460ff1681565b61014b610242366004610941565b610605565b61014b61025536600461095c565b610651565b600054610176906001600160a01b031681565b61029461027b3660046108bc565b6004602052600090815260409020805460019091015482565b6040805192835260208301919091520161012f565b61014b6102b736600461098f565b61075a565b600354600160a01b900460ff166101c4565b604080518082019091526000808252602082015263ffffffff821660009081526006602052604090205460ff166103205760405162461bcd60e51b8152600401610317906109d2565b60405180910390fd5b5063ffffffff16600090815260046020908152604091829020825180840190935280548352600101549082015290565b6000546001600160a01b0316331480156103745750600354600160a01b900460ff16155b8061039d57506002546001600160a01b03163314801561039d5750600354600160a01b900460ff165b6104445760405162461bcd60e51b815260206004820152606660248201527f6f6e6c79416e7a656e5461736b4d616e616765724f7246616c6c4261636b496660448201527f446973707574653a206e6f742066726f6d20616e7a656e207461736b206d616e60648201527f61676572206f722066616c6c206261636b2073616665747920666163746f72206084820152653837b9ba32b960d11b60a482015260c401610317565b63ffffffff821660009081526006602052604090205460ff166104795760405162461bcd60e51b8152600401610317906109d2565b60408051808201825282815242602080830191825263ffffffff861660008181526004835285902093518455915160019093019290925582519081529081018390527fa44739527f61ae356bd84c77fef41b10448dfa00b4003185bd685d0262002ee991015b60405180910390a15050565b60038054821515600160a01b0260ff60a01b199091161790556040517fdf50452fa5e81a42df35c9251524558f2706512ff3f7d76f748faebc80bf3d119061053890831515815260200190565b60405180910390a150565b6001546001600160a01b0316331461056d5760405162461bcd60e51b815260040161031790610a02565b63ffffffff811660009081526006602052604090205460ff166105a25760405162461bcd60e51b8152600401610317906109d2565b63ffffffff81166000818152600660209081526040808320805460ff19169055600582529182902080546001600160a01b031916905590519182527fce1aef06406bb91542575b2d3cf322f105d9882c25386a7b14e32dfd5f69809f9101610538565b6001546001600160a01b0316331461062f5760405162461bcd60e51b815260040161031790610a02565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b0316331461067b5760405162461bcd60e51b815260040161031790610a02565b63ffffffff821660009081526006602052604090205460ff16156106e15760405162461bcd60e51b815260206004820152601a60248201527f50726f746f636f6c20697320616c7265616479206163746976650000000000006044820152606401610317565b63ffffffff82166000818152600660209081526040808320805460ff19166001179055600582529182902080546001600160a01b0386166001600160a01b031990911681179091558251938452908301527fa2a1b0869a1e50f59e1537abb0b64453a3568f48b5b9ed7add5d30fc1e83421891016104df565b600754610100900460ff161580801561077a5750600754600160ff909116105b806107945750303b158015610794575060075460ff166001145b6107f75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610317565b6007805460ff19166001179055801561081a576007805461ff0019166101001790555b600080546001600160a01b038087166001600160a01b031992831617909255600180548684169083161790556002805492851692909116919091179055801561089d576007805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b803563ffffffff811681146108b757600080fd5b919050565b6000602082840312156108ce57600080fd5b6108d7826108a3565b9392505050565b600080604083850312156108f157600080fd5b6108fa836108a3565b946020939093013593505050565b60006020828403121561091a57600080fd5b813580151581146108d757600080fd5b80356001600160a01b03811681146108b757600080fd5b60006020828403121561095357600080fd5b6108d78261092a565b6000806040838503121561096f57600080fd5b610978836108a3565b91506109866020840161092a565b90509250929050565b6000806000606084860312156109a457600080fd5b6109ad8461092a565b92506109bb6020850161092a565b91506109c96040850161092a565b90509250925092565b60208082526016908201527550726f746f636f6c206973206e6f742061637469766560501b604082015260600190565b6020808252818101527f6f6e6c79416e7a656e476f763a206e6f742066726f6d20616e7a656e20676f7660408201526060019056fea264697066735822122037768ce1de646861d3de1dcb8a494028bad26480a4a970147e7070afbdcb5b7864736f6c63430008180033",
}

// ContractSafetyFactorOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSafetyFactorOracleMetaData.ABI instead.
var ContractSafetyFactorOracleABI = ContractSafetyFactorOracleMetaData.ABI

// ContractSafetyFactorOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSafetyFactorOracleMetaData.Bin instead.
var ContractSafetyFactorOracleBin = ContractSafetyFactorOracleMetaData.Bin

// DeployContractSafetyFactorOracle deploys a new Ethereum contract, binding an instance of ContractSafetyFactorOracle to it.
func DeployContractSafetyFactorOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractSafetyFactorOracle, error) {
	parsed, err := ContractSafetyFactorOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSafetyFactorOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSafetyFactorOracle{ContractSafetyFactorOracleCaller: ContractSafetyFactorOracleCaller{contract: contract}, ContractSafetyFactorOracleTransactor: ContractSafetyFactorOracleTransactor{contract: contract}, ContractSafetyFactorOracleFilterer: ContractSafetyFactorOracleFilterer{contract: contract}}, nil
}

// ContractSafetyFactorOracle is an auto generated Go binding around an Ethereum contract.
type ContractSafetyFactorOracle struct {
	ContractSafetyFactorOracleCaller     // Read-only binding to the contract
	ContractSafetyFactorOracleTransactor // Write-only binding to the contract
	ContractSafetyFactorOracleFilterer   // Log filterer for contract events
}

// ContractSafetyFactorOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSafetyFactorOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSafetyFactorOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSafetyFactorOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSafetyFactorOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSafetyFactorOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSafetyFactorOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSafetyFactorOracleSession struct {
	Contract     *ContractSafetyFactorOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractSafetyFactorOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSafetyFactorOracleCallerSession struct {
	Contract *ContractSafetyFactorOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractSafetyFactorOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSafetyFactorOracleTransactorSession struct {
	Contract     *ContractSafetyFactorOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractSafetyFactorOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSafetyFactorOracleRaw struct {
	Contract *ContractSafetyFactorOracle // Generic contract binding to access the raw methods on
}

// ContractSafetyFactorOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSafetyFactorOracleCallerRaw struct {
	Contract *ContractSafetyFactorOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSafetyFactorOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSafetyFactorOracleTransactorRaw struct {
	Contract *ContractSafetyFactorOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSafetyFactorOracle creates a new instance of ContractSafetyFactorOracle, bound to a specific deployed contract.
func NewContractSafetyFactorOracle(address common.Address, backend bind.ContractBackend) (*ContractSafetyFactorOracle, error) {
	contract, err := bindContractSafetyFactorOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracle{ContractSafetyFactorOracleCaller: ContractSafetyFactorOracleCaller{contract: contract}, ContractSafetyFactorOracleTransactor: ContractSafetyFactorOracleTransactor{contract: contract}, ContractSafetyFactorOracleFilterer: ContractSafetyFactorOracleFilterer{contract: contract}}, nil
}

// NewContractSafetyFactorOracleCaller creates a new read-only instance of ContractSafetyFactorOracle, bound to a specific deployed contract.
func NewContractSafetyFactorOracleCaller(address common.Address, caller bind.ContractCaller) (*ContractSafetyFactorOracleCaller, error) {
	contract, err := bindContractSafetyFactorOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleCaller{contract: contract}, nil
}

// NewContractSafetyFactorOracleTransactor creates a new write-only instance of ContractSafetyFactorOracle, bound to a specific deployed contract.
func NewContractSafetyFactorOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSafetyFactorOracleTransactor, error) {
	contract, err := bindContractSafetyFactorOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleTransactor{contract: contract}, nil
}

// NewContractSafetyFactorOracleFilterer creates a new log filterer instance of ContractSafetyFactorOracle, bound to a specific deployed contract.
func NewContractSafetyFactorOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSafetyFactorOracleFilterer, error) {
	contract, err := bindContractSafetyFactorOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleFilterer{contract: contract}, nil
}

// bindContractSafetyFactorOracle binds a generic wrapper to an already deployed contract.
func bindContractSafetyFactorOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSafetyFactorOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSafetyFactorOracle.Contract.ContractSafetyFactorOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.ContractSafetyFactorOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.ContractSafetyFactorOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSafetyFactorOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.contract.Transact(opts, method, params...)
}

// ActiveProtocols is a free data retrieval call binding the contract method 0x3cac45df.
//
// Solidity: function activeProtocols(uint32 ) view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) ActiveProtocols(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "activeProtocols", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveProtocols is a free data retrieval call binding the contract method 0x3cac45df.
//
// Solidity: function activeProtocols(uint32 ) view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) ActiveProtocols(arg0 uint32) (bool, error) {
	return _ContractSafetyFactorOracle.Contract.ActiveProtocols(&_ContractSafetyFactorOracle.CallOpts, arg0)
}

// ActiveProtocols is a free data retrieval call binding the contract method 0x3cac45df.
//
// Solidity: function activeProtocols(uint32 ) view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) ActiveProtocols(arg0 uint32) (bool, error) {
	return _ContractSafetyFactorOracle.Contract.ActiveProtocols(&_ContractSafetyFactorOracle.CallOpts, arg0)
}

// AnzenGov is a free data retrieval call binding the contract method 0x6b8731a4.
//
// Solidity: function anzenGov() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) AnzenGov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "anzenGov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnzenGov is a free data retrieval call binding the contract method 0x6b8731a4.
//
// Solidity: function anzenGov() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) AnzenGov() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.AnzenGov(&_ContractSafetyFactorOracle.CallOpts)
}

// AnzenGov is a free data retrieval call binding the contract method 0x6b8731a4.
//
// Solidity: function anzenGov() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) AnzenGov() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.AnzenGov(&_ContractSafetyFactorOracle.CallOpts)
}

// AnzenTaskManager is a free data retrieval call binding the contract method 0xb44eb609.
//
// Solidity: function anzenTaskManager() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) AnzenTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "anzenTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnzenTaskManager is a free data retrieval call binding the contract method 0xb44eb609.
//
// Solidity: function anzenTaskManager() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) AnzenTaskManager() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.AnzenTaskManager(&_ContractSafetyFactorOracle.CallOpts)
}

// AnzenTaskManager is a free data retrieval call binding the contract method 0xb44eb609.
//
// Solidity: function anzenTaskManager() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) AnzenTaskManager() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.AnzenTaskManager(&_ContractSafetyFactorOracle.CallOpts)
}

// AvsReservesManagerFactory is a free data retrieval call binding the contract method 0x78e3cdaf.
//
// Solidity: function avsReservesManagerFactory() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) AvsReservesManagerFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "avsReservesManagerFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsReservesManagerFactory is a free data retrieval call binding the contract method 0x78e3cdaf.
//
// Solidity: function avsReservesManagerFactory() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) AvsReservesManagerFactory() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.AvsReservesManagerFactory(&_ContractSafetyFactorOracle.CallOpts)
}

// AvsReservesManagerFactory is a free data retrieval call binding the contract method 0x78e3cdaf.
//
// Solidity: function avsReservesManagerFactory() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) AvsReservesManagerFactory() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.AvsReservesManagerFactory(&_ContractSafetyFactorOracle.CallOpts)
}

// FallBackSafetyFactorPoster is a free data retrieval call binding the contract method 0x5bf9c96f.
//
// Solidity: function fallBackSafetyFactorPoster() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) FallBackSafetyFactorPoster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "fallBackSafetyFactorPoster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FallBackSafetyFactorPoster is a free data retrieval call binding the contract method 0x5bf9c96f.
//
// Solidity: function fallBackSafetyFactorPoster() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) FallBackSafetyFactorPoster() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.FallBackSafetyFactorPoster(&_ContractSafetyFactorOracle.CallOpts)
}

// FallBackSafetyFactorPoster is a free data retrieval call binding the contract method 0x5bf9c96f.
//
// Solidity: function fallBackSafetyFactorPoster() view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) FallBackSafetyFactorPoster() (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.FallBackSafetyFactorPoster(&_ContractSafetyFactorOracle.CallOpts)
}

// GetDisputeStatus is a free data retrieval call binding the contract method 0xd6ae724f.
//
// Solidity: function getDisputeStatus() view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) GetDisputeStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "getDisputeStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetDisputeStatus is a free data retrieval call binding the contract method 0xd6ae724f.
//
// Solidity: function getDisputeStatus() view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) GetDisputeStatus() (bool, error) {
	return _ContractSafetyFactorOracle.Contract.GetDisputeStatus(&_ContractSafetyFactorOracle.CallOpts)
}

// GetDisputeStatus is a free data retrieval call binding the contract method 0xd6ae724f.
//
// Solidity: function getDisputeStatus() view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) GetDisputeStatus() (bool, error) {
	return _ContractSafetyFactorOracle.Contract.GetDisputeStatus(&_ContractSafetyFactorOracle.CallOpts)
}

// GetSafetyFactor is a free data retrieval call binding the contract method 0x111ab891.
//
// Solidity: function getSafetyFactor(uint32 _protocolId) view returns((int256,uint256) safetyFactor)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) GetSafetyFactor(opts *bind.CallOpts, _protocolId uint32) (SafetyFactorSnapshot, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "getSafetyFactor", _protocolId)

	if err != nil {
		return *new(SafetyFactorSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(SafetyFactorSnapshot)).(*SafetyFactorSnapshot)

	return out0, err

}

// GetSafetyFactor is a free data retrieval call binding the contract method 0x111ab891.
//
// Solidity: function getSafetyFactor(uint32 _protocolId) view returns((int256,uint256) safetyFactor)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) GetSafetyFactor(_protocolId uint32) (SafetyFactorSnapshot, error) {
	return _ContractSafetyFactorOracle.Contract.GetSafetyFactor(&_ContractSafetyFactorOracle.CallOpts, _protocolId)
}

// GetSafetyFactor is a free data retrieval call binding the contract method 0x111ab891.
//
// Solidity: function getSafetyFactor(uint32 _protocolId) view returns((int256,uint256) safetyFactor)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) GetSafetyFactor(_protocolId uint32) (SafetyFactorSnapshot, error) {
	return _ContractSafetyFactorOracle.Contract.GetSafetyFactor(&_ContractSafetyFactorOracle.CallOpts, _protocolId)
}

// PendingDispute is a free data retrieval call binding the contract method 0x80bd9f33.
//
// Solidity: function pendingDispute() view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) PendingDispute(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "pendingDispute")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingDispute is a free data retrieval call binding the contract method 0x80bd9f33.
//
// Solidity: function pendingDispute() view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) PendingDispute() (bool, error) {
	return _ContractSafetyFactorOracle.Contract.PendingDispute(&_ContractSafetyFactorOracle.CallOpts)
}

// PendingDispute is a free data retrieval call binding the contract method 0x80bd9f33.
//
// Solidity: function pendingDispute() view returns(bool)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) PendingDispute() (bool, error) {
	return _ContractSafetyFactorOracle.Contract.PendingDispute(&_ContractSafetyFactorOracle.CallOpts)
}

// ProtocolToReservesManager is a free data retrieval call binding the contract method 0x37898246.
//
// Solidity: function protocolToReservesManager(uint32 ) view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) ProtocolToReservesManager(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "protocolToReservesManager", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolToReservesManager is a free data retrieval call binding the contract method 0x37898246.
//
// Solidity: function protocolToReservesManager(uint32 ) view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) ProtocolToReservesManager(arg0 uint32) (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.ProtocolToReservesManager(&_ContractSafetyFactorOracle.CallOpts, arg0)
}

// ProtocolToReservesManager is a free data retrieval call binding the contract method 0x37898246.
//
// Solidity: function protocolToReservesManager(uint32 ) view returns(address)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) ProtocolToReservesManager(arg0 uint32) (common.Address, error) {
	return _ContractSafetyFactorOracle.Contract.ProtocolToReservesManager(&_ContractSafetyFactorOracle.CallOpts, arg0)
}

// SafetyFactorSnapshots is a free data retrieval call binding the contract method 0xbb6cb0dd.
//
// Solidity: function safetyFactorSnapshots(uint32 ) view returns(int256 safetyFactor, uint256 timestamp)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCaller) SafetyFactorSnapshots(opts *bind.CallOpts, arg0 uint32) (struct {
	SafetyFactor *big.Int
	Timestamp    *big.Int
}, error) {
	var out []interface{}
	err := _ContractSafetyFactorOracle.contract.Call(opts, &out, "safetyFactorSnapshots", arg0)

	outstruct := new(struct {
		SafetyFactor *big.Int
		Timestamp    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SafetyFactor = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SafetyFactorSnapshots is a free data retrieval call binding the contract method 0xbb6cb0dd.
//
// Solidity: function safetyFactorSnapshots(uint32 ) view returns(int256 safetyFactor, uint256 timestamp)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) SafetyFactorSnapshots(arg0 uint32) (struct {
	SafetyFactor *big.Int
	Timestamp    *big.Int
}, error) {
	return _ContractSafetyFactorOracle.Contract.SafetyFactorSnapshots(&_ContractSafetyFactorOracle.CallOpts, arg0)
}

// SafetyFactorSnapshots is a free data retrieval call binding the contract method 0xbb6cb0dd.
//
// Solidity: function safetyFactorSnapshots(uint32 ) view returns(int256 safetyFactor, uint256 timestamp)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleCallerSession) SafetyFactorSnapshots(arg0 uint32) (struct {
	SafetyFactor *big.Int
	Timestamp    *big.Int
}, error) {
	return _ContractSafetyFactorOracle.Contract.SafetyFactorSnapshots(&_ContractSafetyFactorOracle.CallOpts, arg0)
}

// AddProtocol is a paid mutator transaction binding the contract method 0xa3e45311.
//
// Solidity: function addProtocol(uint32 _protocolIdToAdd, address _reservesManager) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactor) AddProtocol(opts *bind.TransactOpts, _protocolIdToAdd uint32, _reservesManager common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.contract.Transact(opts, "addProtocol", _protocolIdToAdd, _reservesManager)
}

// AddProtocol is a paid mutator transaction binding the contract method 0xa3e45311.
//
// Solidity: function addProtocol(uint32 _protocolIdToAdd, address _reservesManager) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) AddProtocol(_protocolIdToAdd uint32, _reservesManager common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.AddProtocol(&_ContractSafetyFactorOracle.TransactOpts, _protocolIdToAdd, _reservesManager)
}

// AddProtocol is a paid mutator transaction binding the contract method 0xa3e45311.
//
// Solidity: function addProtocol(uint32 _protocolIdToAdd, address _reservesManager) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorSession) AddProtocol(_protocolIdToAdd uint32, _reservesManager common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.AddProtocol(&_ContractSafetyFactorOracle.TransactOpts, _protocolIdToAdd, _reservesManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _anzenTaskManager, address _anzenGov, address _fallBackSafetyFactorPoster) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactor) Initialize(opts *bind.TransactOpts, _anzenTaskManager common.Address, _anzenGov common.Address, _fallBackSafetyFactorPoster common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.contract.Transact(opts, "initialize", _anzenTaskManager, _anzenGov, _fallBackSafetyFactorPoster)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _anzenTaskManager, address _anzenGov, address _fallBackSafetyFactorPoster) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) Initialize(_anzenTaskManager common.Address, _anzenGov common.Address, _fallBackSafetyFactorPoster common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.Initialize(&_ContractSafetyFactorOracle.TransactOpts, _anzenTaskManager, _anzenGov, _fallBackSafetyFactorPoster)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _anzenTaskManager, address _anzenGov, address _fallBackSafetyFactorPoster) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorSession) Initialize(_anzenTaskManager common.Address, _anzenGov common.Address, _fallBackSafetyFactorPoster common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.Initialize(&_ContractSafetyFactorOracle.TransactOpts, _anzenTaskManager, _anzenGov, _fallBackSafetyFactorPoster)
}

// RemoveProtocol is a paid mutator transaction binding the contract method 0x557d0f64.
//
// Solidity: function removeProtocol(uint32 _protocolIdToRemove) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactor) RemoveProtocol(opts *bind.TransactOpts, _protocolIdToRemove uint32) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.contract.Transact(opts, "removeProtocol", _protocolIdToRemove)
}

// RemoveProtocol is a paid mutator transaction binding the contract method 0x557d0f64.
//
// Solidity: function removeProtocol(uint32 _protocolIdToRemove) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) RemoveProtocol(_protocolIdToRemove uint32) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.RemoveProtocol(&_ContractSafetyFactorOracle.TransactOpts, _protocolIdToRemove)
}

// RemoveProtocol is a paid mutator transaction binding the contract method 0x557d0f64.
//
// Solidity: function removeProtocol(uint32 _protocolIdToRemove) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorSession) RemoveProtocol(_protocolIdToRemove uint32) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.RemoveProtocol(&_ContractSafetyFactorOracle.TransactOpts, _protocolIdToRemove)
}

// SetDisputeStatus is a paid mutator transaction binding the contract method 0x38422fc4.
//
// Solidity: function setDisputeStatus(bool _status) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactor) SetDisputeStatus(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.contract.Transact(opts, "setDisputeStatus", _status)
}

// SetDisputeStatus is a paid mutator transaction binding the contract method 0x38422fc4.
//
// Solidity: function setDisputeStatus(bool _status) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) SetDisputeStatus(_status bool) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.SetDisputeStatus(&_ContractSafetyFactorOracle.TransactOpts, _status)
}

// SetDisputeStatus is a paid mutator transaction binding the contract method 0x38422fc4.
//
// Solidity: function setDisputeStatus(bool _status) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorSession) SetDisputeStatus(_status bool) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.SetDisputeStatus(&_ContractSafetyFactorOracle.TransactOpts, _status)
}

// SetSafetyFactor is a paid mutator transaction binding the contract method 0x27fdf8a6.
//
// Solidity: function setSafetyFactor(uint32 _protocolId, int256 _safetyFactor) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactor) SetSafetyFactor(opts *bind.TransactOpts, _protocolId uint32, _safetyFactor *big.Int) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.contract.Transact(opts, "setSafetyFactor", _protocolId, _safetyFactor)
}

// SetSafetyFactor is a paid mutator transaction binding the contract method 0x27fdf8a6.
//
// Solidity: function setSafetyFactor(uint32 _protocolId, int256 _safetyFactor) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) SetSafetyFactor(_protocolId uint32, _safetyFactor *big.Int) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.SetSafetyFactor(&_ContractSafetyFactorOracle.TransactOpts, _protocolId, _safetyFactor)
}

// SetSafetyFactor is a paid mutator transaction binding the contract method 0x27fdf8a6.
//
// Solidity: function setSafetyFactor(uint32 _protocolId, int256 _safetyFactor) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorSession) SetSafetyFactor(_protocolId uint32, _safetyFactor *big.Int) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.SetSafetyFactor(&_ContractSafetyFactorOracle.TransactOpts, _protocolId, _safetyFactor)
}

// UpdateFallbackPoster is a paid mutator transaction binding the contract method 0x9312afa8.
//
// Solidity: function updateFallbackPoster(address newFallBackSafetyFactorPoster) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactor) UpdateFallbackPoster(opts *bind.TransactOpts, newFallBackSafetyFactorPoster common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.contract.Transact(opts, "updateFallbackPoster", newFallBackSafetyFactorPoster)
}

// UpdateFallbackPoster is a paid mutator transaction binding the contract method 0x9312afa8.
//
// Solidity: function updateFallbackPoster(address newFallBackSafetyFactorPoster) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleSession) UpdateFallbackPoster(newFallBackSafetyFactorPoster common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.UpdateFallbackPoster(&_ContractSafetyFactorOracle.TransactOpts, newFallBackSafetyFactorPoster)
}

// UpdateFallbackPoster is a paid mutator transaction binding the contract method 0x9312afa8.
//
// Solidity: function updateFallbackPoster(address newFallBackSafetyFactorPoster) returns()
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleTransactorSession) UpdateFallbackPoster(newFallBackSafetyFactorPoster common.Address) (*types.Transaction, error) {
	return _ContractSafetyFactorOracle.Contract.UpdateFallbackPoster(&_ContractSafetyFactorOracle.TransactOpts, newFallBackSafetyFactorPoster)
}

// ContractSafetyFactorOracleDisputeStatusSetIterator is returned from FilterDisputeStatusSet and is used to iterate over the raw logs and unpacked data for DisputeStatusSet events raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleDisputeStatusSetIterator struct {
	Event *ContractSafetyFactorOracleDisputeStatusSet // Event containing the contract specifics and raw log

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
func (it *ContractSafetyFactorOracleDisputeStatusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSafetyFactorOracleDisputeStatusSet)
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
		it.Event = new(ContractSafetyFactorOracleDisputeStatusSet)
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
func (it *ContractSafetyFactorOracleDisputeStatusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSafetyFactorOracleDisputeStatusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSafetyFactorOracleDisputeStatusSet represents a DisputeStatusSet event raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleDisputeStatusSet struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisputeStatusSet is a free log retrieval operation binding the contract event 0xdf50452fa5e81a42df35c9251524558f2706512ff3f7d76f748faebc80bf3d11.
//
// Solidity: event DisputeStatusSet(bool status)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) FilterDisputeStatusSet(opts *bind.FilterOpts) (*ContractSafetyFactorOracleDisputeStatusSetIterator, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.FilterLogs(opts, "DisputeStatusSet")
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleDisputeStatusSetIterator{contract: _ContractSafetyFactorOracle.contract, event: "DisputeStatusSet", logs: logs, sub: sub}, nil
}

// WatchDisputeStatusSet is a free log subscription operation binding the contract event 0xdf50452fa5e81a42df35c9251524558f2706512ff3f7d76f748faebc80bf3d11.
//
// Solidity: event DisputeStatusSet(bool status)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) WatchDisputeStatusSet(opts *bind.WatchOpts, sink chan<- *ContractSafetyFactorOracleDisputeStatusSet) (event.Subscription, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.WatchLogs(opts, "DisputeStatusSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSafetyFactorOracleDisputeStatusSet)
				if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "DisputeStatusSet", log); err != nil {
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

// ParseDisputeStatusSet is a log parse operation binding the contract event 0xdf50452fa5e81a42df35c9251524558f2706512ff3f7d76f748faebc80bf3d11.
//
// Solidity: event DisputeStatusSet(bool status)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) ParseDisputeStatusSet(log types.Log) (*ContractSafetyFactorOracleDisputeStatusSet, error) {
	event := new(ContractSafetyFactorOracleDisputeStatusSet)
	if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "DisputeStatusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSafetyFactorOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleInitializedIterator struct {
	Event *ContractSafetyFactorOracleInitialized // Event containing the contract specifics and raw log

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
func (it *ContractSafetyFactorOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSafetyFactorOracleInitialized)
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
		it.Event = new(ContractSafetyFactorOracleInitialized)
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
func (it *ContractSafetyFactorOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSafetyFactorOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSafetyFactorOracleInitialized represents a Initialized event raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSafetyFactorOracleInitializedIterator, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleInitializedIterator{contract: _ContractSafetyFactorOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSafetyFactorOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSafetyFactorOracleInitialized)
				if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) ParseInitialized(log types.Log) (*ContractSafetyFactorOracleInitialized, error) {
	event := new(ContractSafetyFactorOracleInitialized)
	if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSafetyFactorOracleProtocolAddedIterator is returned from FilterProtocolAdded and is used to iterate over the raw logs and unpacked data for ProtocolAdded events raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleProtocolAddedIterator struct {
	Event *ContractSafetyFactorOracleProtocolAdded // Event containing the contract specifics and raw log

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
func (it *ContractSafetyFactorOracleProtocolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSafetyFactorOracleProtocolAdded)
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
		it.Event = new(ContractSafetyFactorOracleProtocolAdded)
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
func (it *ContractSafetyFactorOracleProtocolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSafetyFactorOracleProtocolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSafetyFactorOracleProtocolAdded represents a ProtocolAdded event raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleProtocolAdded struct {
	ProtocolId      uint32
	ReservesManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProtocolAdded is a free log retrieval operation binding the contract event 0xa2a1b0869a1e50f59e1537abb0b64453a3568f48b5b9ed7add5d30fc1e834218.
//
// Solidity: event ProtocolAdded(uint32 protocolId, address reservesManager)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) FilterProtocolAdded(opts *bind.FilterOpts) (*ContractSafetyFactorOracleProtocolAddedIterator, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.FilterLogs(opts, "ProtocolAdded")
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleProtocolAddedIterator{contract: _ContractSafetyFactorOracle.contract, event: "ProtocolAdded", logs: logs, sub: sub}, nil
}

// WatchProtocolAdded is a free log subscription operation binding the contract event 0xa2a1b0869a1e50f59e1537abb0b64453a3568f48b5b9ed7add5d30fc1e834218.
//
// Solidity: event ProtocolAdded(uint32 protocolId, address reservesManager)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) WatchProtocolAdded(opts *bind.WatchOpts, sink chan<- *ContractSafetyFactorOracleProtocolAdded) (event.Subscription, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.WatchLogs(opts, "ProtocolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSafetyFactorOracleProtocolAdded)
				if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "ProtocolAdded", log); err != nil {
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

// ParseProtocolAdded is a log parse operation binding the contract event 0xa2a1b0869a1e50f59e1537abb0b64453a3568f48b5b9ed7add5d30fc1e834218.
//
// Solidity: event ProtocolAdded(uint32 protocolId, address reservesManager)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) ParseProtocolAdded(log types.Log) (*ContractSafetyFactorOracleProtocolAdded, error) {
	event := new(ContractSafetyFactorOracleProtocolAdded)
	if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "ProtocolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSafetyFactorOracleProtocolRemovedIterator is returned from FilterProtocolRemoved and is used to iterate over the raw logs and unpacked data for ProtocolRemoved events raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleProtocolRemovedIterator struct {
	Event *ContractSafetyFactorOracleProtocolRemoved // Event containing the contract specifics and raw log

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
func (it *ContractSafetyFactorOracleProtocolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSafetyFactorOracleProtocolRemoved)
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
		it.Event = new(ContractSafetyFactorOracleProtocolRemoved)
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
func (it *ContractSafetyFactorOracleProtocolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSafetyFactorOracleProtocolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSafetyFactorOracleProtocolRemoved represents a ProtocolRemoved event raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleProtocolRemoved struct {
	ProtocolId uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProtocolRemoved is a free log retrieval operation binding the contract event 0xce1aef06406bb91542575b2d3cf322f105d9882c25386a7b14e32dfd5f69809f.
//
// Solidity: event ProtocolRemoved(uint32 protocolId)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) FilterProtocolRemoved(opts *bind.FilterOpts) (*ContractSafetyFactorOracleProtocolRemovedIterator, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.FilterLogs(opts, "ProtocolRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleProtocolRemovedIterator{contract: _ContractSafetyFactorOracle.contract, event: "ProtocolRemoved", logs: logs, sub: sub}, nil
}

// WatchProtocolRemoved is a free log subscription operation binding the contract event 0xce1aef06406bb91542575b2d3cf322f105d9882c25386a7b14e32dfd5f69809f.
//
// Solidity: event ProtocolRemoved(uint32 protocolId)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) WatchProtocolRemoved(opts *bind.WatchOpts, sink chan<- *ContractSafetyFactorOracleProtocolRemoved) (event.Subscription, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.WatchLogs(opts, "ProtocolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSafetyFactorOracleProtocolRemoved)
				if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "ProtocolRemoved", log); err != nil {
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

// ParseProtocolRemoved is a log parse operation binding the contract event 0xce1aef06406bb91542575b2d3cf322f105d9882c25386a7b14e32dfd5f69809f.
//
// Solidity: event ProtocolRemoved(uint32 protocolId)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) ParseProtocolRemoved(log types.Log) (*ContractSafetyFactorOracleProtocolRemoved, error) {
	event := new(ContractSafetyFactorOracleProtocolRemoved)
	if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "ProtocolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSafetyFactorOracleSFUpdatedIterator is returned from FilterSFUpdated and is used to iterate over the raw logs and unpacked data for SFUpdated events raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleSFUpdatedIterator struct {
	Event *ContractSafetyFactorOracleSFUpdated // Event containing the contract specifics and raw log

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
func (it *ContractSafetyFactorOracleSFUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSafetyFactorOracleSFUpdated)
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
		it.Event = new(ContractSafetyFactorOracleSFUpdated)
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
func (it *ContractSafetyFactorOracleSFUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSafetyFactorOracleSFUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSafetyFactorOracleSFUpdated represents a SFUpdated event raised by the ContractSafetyFactorOracle contract.
type ContractSafetyFactorOracleSFUpdated struct {
	ProtocolId uint32
	NewSF      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSFUpdated is a free log retrieval operation binding the contract event 0xa44739527f61ae356bd84c77fef41b10448dfa00b4003185bd685d0262002ee9.
//
// Solidity: event SFUpdated(uint32 protocolId, int256 newSF)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) FilterSFUpdated(opts *bind.FilterOpts) (*ContractSafetyFactorOracleSFUpdatedIterator, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.FilterLogs(opts, "SFUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractSafetyFactorOracleSFUpdatedIterator{contract: _ContractSafetyFactorOracle.contract, event: "SFUpdated", logs: logs, sub: sub}, nil
}

// WatchSFUpdated is a free log subscription operation binding the contract event 0xa44739527f61ae356bd84c77fef41b10448dfa00b4003185bd685d0262002ee9.
//
// Solidity: event SFUpdated(uint32 protocolId, int256 newSF)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) WatchSFUpdated(opts *bind.WatchOpts, sink chan<- *ContractSafetyFactorOracleSFUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractSafetyFactorOracle.contract.WatchLogs(opts, "SFUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSafetyFactorOracleSFUpdated)
				if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "SFUpdated", log); err != nil {
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

// ParseSFUpdated is a log parse operation binding the contract event 0xa44739527f61ae356bd84c77fef41b10448dfa00b4003185bd685d0262002ee9.
//
// Solidity: event SFUpdated(uint32 protocolId, int256 newSF)
func (_ContractSafetyFactorOracle *ContractSafetyFactorOracleFilterer) ParseSFUpdated(log types.Log) (*ContractSafetyFactorOracleSFUpdated, error) {
	event := new(ContractSafetyFactorOracleSFUpdated)
	if err := _ContractSafetyFactorOracle.contract.UnpackLog(event, "SFUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
