// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAnzenServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractAnzenServiceManagerMetaData contains all meta data concerning the ContractAnzenServiceManager contract.
var ContractAnzenServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_anzenTaskManager\",\"type\":\"address\",\"internalType\":\"contractIAnzenTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"anzenTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAnzenTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001ae938038062001ae983398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e05161010051611872620002776000396000818161019b0152610913015260008181610652015281816107a40152818161083b01528181610c3201528181610dac0152610e4b01526000818161047d0152818161050c0152818161058c015281816109c101528181610a5701528181610b6d0152610d0701526000818161031501526103eb01526000818161010c01528181610a1501528181610ab30152610b3201526118726000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a364f4da14610170578063a98fb35514610183578063b44eb60914610196578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a6146101445780638da5cb5b1461014c575b600080fd5b6100cc6100c736600461112a565b6101d8565b005b6100e16100dc3660046111b4565b610458565b6040516100ee91906111d8565b60405180910390f35b6100cc6101053660046111b4565b610908565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109a2565b6033546001600160a01b031661012c565b6100cc61016b3660046112da565b6109b6565b6100cc61017e3660046111b4565b610a4c565b6100cc610191366004611385565b610b13565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6100e1610b67565b6100cc6101d33660046111b4565b610f14565b6101e0610f8a565b60005b818110156103d3578282828181106101fd576101fd6113d6565b905060200281019061020f91906113ec565b6102209060408101906020016111b4565b6001600160a01b03166323b872dd3330868686818110610242576102426113d6565b905060200281019061025491906113ec565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf919061141c565b508282828181106102e2576102e26113d6565b90506020028101906102f491906113ec565b6103059060408101906020016111b4565b6001600160a01b031663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000858585818110610346576103466113d6565b905060200281019061035891906113ec565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca919061141c565b506001016101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042290859085906004016114d7565b600060405180830381600087803b15801561043c57600080fd5b505af1158015610450573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e891906115e6565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610553573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057791906115ff565b90506001600160c01b038116158061061157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060c9190611628565b60ff16155b1561062d57505060408051600081526020810190915292915050565b6000610641826001600160c01b0316610fe4565b90506000805b825181101561070d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f5848381518110610691576106916113d6565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f991906115e6565b6107039083611661565b9150600101610647565b5060008167ffffffffffffffff81111561072957610729611225565b604051908082528060200260200182016040528015610752578160200160208202803683370190505b5090506000805b84518110156108fb576000858281518110610776576107766113d6565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080f91906115e6565b905060005b818110156108f0576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610889573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ad9190611674565b600001518686815181106108c3576108c36113d6565b6001600160a01b0390921660209283029190910190910152846108e5816116d3565b955050600101610814565b505050600101610759565b5090979650505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461099f5760405162461bcd60e51b815260206004820152603160248201527f6f6e6c79416e7a656e5461736b4d616e616765723a206e6f742066726f6d2061604482015270373d32b7103a30b9b59036b0b730b3b2b960791b60648201526084015b60405180910390fd5b50565b6109aa610f8a565b6109b460006110a7565b565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109fe5760405162461bcd60e51b8152600401610996906116ec565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042290859085906004016117aa565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a945760405162461bcd60e51b8152600401610996906116ec565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610af857600080fd5b505af1158015610b0c573d6000803e3d6000fd5b5050505050565b610b1b610f8a565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610ade9084906004016117f5565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bed9190611628565b60ff16905080600003610c0e57505060408051600081526020810190915290565b6000805b82811015610cb957604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610c81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca591906115e6565b610caf9083611661565b9150600101610c12565b5060008167ffffffffffffffff811115610cd557610cd5611225565b604051908082528060200260200182016040528015610cfe578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d879190611628565b60ff16811015610f0a57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610dfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1f91906115e6565b905060005b81811015610f00576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610e99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ebd9190611674565b60000151858581518110610ed357610ed36113d6565b6001600160a01b039092166020928302919091019091015283610ef5816116d3565b945050600101610e24565b5050600101610d05565b5090949350505050565b610f1c610f8a565b6001600160a01b038116610f815760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610996565b61099f816110a7565b6033546001600160a01b031633146109b45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610996565b6060600080610ff2846110f9565b61ffff1667ffffffffffffffff81111561100e5761100e611225565b6040519080825280601f01601f191660200182016040528015611038576020820181803683370190505b5090506000805b825182108015611050575061010081105b15610f0a576001811b935085841615611097578060f81b838381518110611079576110796113d6565b60200101906001600160f81b031916908160001a9053508160010191505b6110a0816116d3565b905061103f565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156111245761110e600184611808565b909216918061111c8161181b565b9150506110fd565b92915050565b6000806020838503121561113d57600080fd5b823567ffffffffffffffff8082111561115557600080fd5b818501915085601f83011261116957600080fd5b81358181111561117857600080fd5b8660208260051b850101111561118d57600080fd5b60209290920196919550909350505050565b6001600160a01b038116811461099f57600080fd5b6000602082840312156111c657600080fd5b81356111d18161119f565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112195783516001600160a01b0316835292840192918401916001016111f4565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561125e5761125e611225565b60405290565b600067ffffffffffffffff8084111561127f5761127f611225565b604051601f8501601f19908116603f011681019082821181831017156112a7576112a7611225565b816040528093508581528686860111156112c057600080fd5b858560208301376000602087830101525050509392505050565b600080604083850312156112ed57600080fd5b82356112f88161119f565b9150602083013567ffffffffffffffff8082111561131557600080fd5b908401906060828703121561132957600080fd5b61133161123b565b82358281111561134057600080fd5b83019150601f8201871361135357600080fd5b61136287833560208501611264565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561139757600080fd5b813567ffffffffffffffff8111156113ae57600080fd5b8201601f810184136113bf57600080fd5b6113ce84823560208401611264565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261140257600080fd5b9190910192915050565b80356114178161119f565b919050565b60006020828403121561142e57600080fd5b815180151581146111d157600080fd5b6bffffffffffffffffffffffff8116811461099f57600080fd5b8183526000602080850194508260005b858110156114b857813561147b8161119f565b6001600160a01b03168752818301356114938161143e565b6bffffffffffffffffffffffff16878401526040968701969190910190600101611468565b509495945050505050565b803563ffffffff8116811461141757600080fd5b60208082528181018390526000906040808401600586901b8501820187855b888110156115d857878303603f190184528135368b9003609e1901811261151c57600080fd5b8a0160a0813536839003601e1901811261153557600080fd5b8201888101903567ffffffffffffffff81111561155157600080fd5b8060061b360382131561156357600080fd5b8287526115738388018284611458565b9250505061158288830161140c565b6001600160a01b031688860152818701358786015260606115a48184016114c3565b63ffffffff169086015260806115bb8382016114c3565b63ffffffff169501949094525092850192908501906001016114f6565b509098975050505050505050565b6000602082840312156115f857600080fd5b5051919050565b60006020828403121561161157600080fd5b81516001600160c01b03811681146111d157600080fd5b60006020828403121561163a57600080fd5b815160ff811681146111d157600080fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156111245761112461164b565b60006040828403121561168657600080fd5b6040516040810181811067ffffffffffffffff821117156116a9576116a9611225565b60405282516116b78161119f565b815260208301516116c78161143e565b60208201529392505050565b6000600182016116e5576116e561164b565b5060010190565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b8181101561178a5760208185018101518683018201520161176e565b506000602082860101526020601f19601f83011685010191505092915050565b60018060a01b03831681526040602082015260008251606060408401526117d460a0840182611764565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006111d16020830184611764565b818103818111156111245761112461164b565b600061ffff8083168181036118325761183261164b565b600101939250505056fea2646970667358221220accb73d0490d9d6e4216a247a409c69b5508de6a7fb01ba69701333ed276b9d764736f6c63430008180033",
}

// ContractAnzenServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAnzenServiceManagerMetaData.ABI instead.
var ContractAnzenServiceManagerABI = ContractAnzenServiceManagerMetaData.ABI

// ContractAnzenServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAnzenServiceManagerMetaData.Bin instead.
var ContractAnzenServiceManagerBin = ContractAnzenServiceManagerMetaData.Bin

// DeployContractAnzenServiceManager deploys a new Ethereum contract, binding an instance of ContractAnzenServiceManager to it.
func DeployContractAnzenServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _anzenTaskManager common.Address) (common.Address, *types.Transaction, *ContractAnzenServiceManager, error) {
	parsed, err := ContractAnzenServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAnzenServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _anzenTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAnzenServiceManager{ContractAnzenServiceManagerCaller: ContractAnzenServiceManagerCaller{contract: contract}, ContractAnzenServiceManagerTransactor: ContractAnzenServiceManagerTransactor{contract: contract}, ContractAnzenServiceManagerFilterer: ContractAnzenServiceManagerFilterer{contract: contract}}, nil
}

// ContractAnzenServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractAnzenServiceManager struct {
	ContractAnzenServiceManagerCaller     // Read-only binding to the contract
	ContractAnzenServiceManagerTransactor // Write-only binding to the contract
	ContractAnzenServiceManagerFilterer   // Log filterer for contract events
}

// ContractAnzenServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAnzenServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAnzenServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAnzenServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAnzenServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAnzenServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAnzenServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAnzenServiceManagerSession struct {
	Contract     *ContractAnzenServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractAnzenServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAnzenServiceManagerCallerSession struct {
	Contract *ContractAnzenServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ContractAnzenServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAnzenServiceManagerTransactorSession struct {
	Contract     *ContractAnzenServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractAnzenServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAnzenServiceManagerRaw struct {
	Contract *ContractAnzenServiceManager // Generic contract binding to access the raw methods on
}

// ContractAnzenServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAnzenServiceManagerCallerRaw struct {
	Contract *ContractAnzenServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAnzenServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAnzenServiceManagerTransactorRaw struct {
	Contract *ContractAnzenServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAnzenServiceManager creates a new instance of ContractAnzenServiceManager, bound to a specific deployed contract.
func NewContractAnzenServiceManager(address common.Address, backend bind.ContractBackend) (*ContractAnzenServiceManager, error) {
	contract, err := bindContractAnzenServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenServiceManager{ContractAnzenServiceManagerCaller: ContractAnzenServiceManagerCaller{contract: contract}, ContractAnzenServiceManagerTransactor: ContractAnzenServiceManagerTransactor{contract: contract}, ContractAnzenServiceManagerFilterer: ContractAnzenServiceManagerFilterer{contract: contract}}, nil
}

// NewContractAnzenServiceManagerCaller creates a new read-only instance of ContractAnzenServiceManager, bound to a specific deployed contract.
func NewContractAnzenServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractAnzenServiceManagerCaller, error) {
	contract, err := bindContractAnzenServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenServiceManagerCaller{contract: contract}, nil
}

// NewContractAnzenServiceManagerTransactor creates a new write-only instance of ContractAnzenServiceManager, bound to a specific deployed contract.
func NewContractAnzenServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAnzenServiceManagerTransactor, error) {
	contract, err := bindContractAnzenServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenServiceManagerTransactor{contract: contract}, nil
}

// NewContractAnzenServiceManagerFilterer creates a new log filterer instance of ContractAnzenServiceManager, bound to a specific deployed contract.
func NewContractAnzenServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAnzenServiceManagerFilterer, error) {
	contract, err := bindContractAnzenServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenServiceManagerFilterer{contract: contract}, nil
}

// bindContractAnzenServiceManager binds a generic wrapper to an already deployed contract.
func bindContractAnzenServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAnzenServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAnzenServiceManager.Contract.ContractAnzenServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.ContractAnzenServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.ContractAnzenServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAnzenServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AnzenTaskManager is a free data retrieval call binding the contract method 0xb44eb609.
//
// Solidity: function anzenTaskManager() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCaller) AnzenTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenServiceManager.contract.Call(opts, &out, "anzenTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnzenTaskManager is a free data retrieval call binding the contract method 0xb44eb609.
//
// Solidity: function anzenTaskManager() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) AnzenTaskManager() (common.Address, error) {
	return _ContractAnzenServiceManager.Contract.AnzenTaskManager(&_ContractAnzenServiceManager.CallOpts)
}

// AnzenTaskManager is a free data retrieval call binding the contract method 0xb44eb609.
//
// Solidity: function anzenTaskManager() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCallerSession) AnzenTaskManager() (common.Address, error) {
	return _ContractAnzenServiceManager.Contract.AnzenTaskManager(&_ContractAnzenServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractAnzenServiceManager.Contract.AvsDirectory(&_ContractAnzenServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractAnzenServiceManager.Contract.AvsDirectory(&_ContractAnzenServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAnzenServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractAnzenServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractAnzenServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractAnzenServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractAnzenServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractAnzenServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractAnzenServiceManager.Contract.GetRestakeableStrategies(&_ContractAnzenServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractAnzenServiceManager.Contract.GetRestakeableStrategies(&_ContractAnzenServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAnzenServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) Owner() (common.Address, error) {
	return _ContractAnzenServiceManager.Contract.Owner(&_ContractAnzenServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractAnzenServiceManager.Contract.Owner(&_ContractAnzenServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractAnzenServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractAnzenServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.FreezeOperator(&_ContractAnzenServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.FreezeOperator(&_ContractAnzenServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.PayForRange(&_ContractAnzenServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.PayForRange(&_ContractAnzenServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.RegisterOperatorToAVS(&_ContractAnzenServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.RegisterOperatorToAVS(&_ContractAnzenServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.RenounceOwnership(&_ContractAnzenServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.RenounceOwnership(&_ContractAnzenServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.TransferOwnership(&_ContractAnzenServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.TransferOwnership(&_ContractAnzenServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.UpdateAVSMetadataURI(&_ContractAnzenServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractAnzenServiceManager.Contract.UpdateAVSMetadataURI(&_ContractAnzenServiceManager.TransactOpts, _metadataURI)
}

// ContractAnzenServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAnzenServiceManager contract.
type ContractAnzenServiceManagerInitializedIterator struct {
	Event *ContractAnzenServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractAnzenServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenServiceManagerInitialized)
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
		it.Event = new(ContractAnzenServiceManagerInitialized)
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
func (it *ContractAnzenServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenServiceManagerInitialized represents a Initialized event raised by the ContractAnzenServiceManager contract.
type ContractAnzenServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAnzenServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractAnzenServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAnzenServiceManagerInitializedIterator{contract: _ContractAnzenServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAnzenServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAnzenServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenServiceManagerInitialized)
				if err := _ContractAnzenServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractAnzenServiceManagerInitialized, error) {
	event := new(ContractAnzenServiceManagerInitialized)
	if err := _ContractAnzenServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAnzenServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAnzenServiceManager contract.
type ContractAnzenServiceManagerOwnershipTransferredIterator struct {
	Event *ContractAnzenServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAnzenServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAnzenServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractAnzenServiceManagerOwnershipTransferred)
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
func (it *ContractAnzenServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAnzenServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAnzenServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAnzenServiceManager contract.
type ContractAnzenServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAnzenServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAnzenServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAnzenServiceManagerOwnershipTransferredIterator{contract: _ContractAnzenServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAnzenServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAnzenServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAnzenServiceManagerOwnershipTransferred)
				if err := _ContractAnzenServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAnzenServiceManager *ContractAnzenServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAnzenServiceManagerOwnershipTransferred, error) {
	event := new(ContractAnzenServiceManagerOwnershipTransferred)
	if err := _ContractAnzenServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
