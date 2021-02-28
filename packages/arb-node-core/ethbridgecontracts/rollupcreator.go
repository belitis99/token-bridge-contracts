// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"deliverMessageToInbox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b610d958061007d6000396000f3fe6080604052600436106100975760003560e01c806302bbfad11461009c5780633dbcc8d1146100e3578063413b35bd146100f8578063715018a61461013f5780638da5cb5b146101565780639e5d4c4914610187578063ab5d894314610298578063c29372de146102ad578063cee3d728146102e0578063d9dd67ab1461031b578063e45b7ce614610345578063f2fde38b14610380575b600080fd5b6100d1600480360360608110156100b257600080fd5b5060ff813516906001600160a01b0360208201351690604001356103b3565b60408051918252519081900360200190f35b3480156100ef57600080fd5b506100d16104c9565b34801561010457600080fd5b5061012b6004803603602081101561011b57600080fd5b50356001600160a01b03166104cf565b604080519115158252519081900360200190f35b34801561014b57600080fd5b506101546104f0565b005b34801561016257600080fd5b5061016b61058a565b604080516001600160a01b039092168252519081900360200190f35b34801561019357600080fd5b50610217600480360360608110156101aa57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156101d957600080fd5b8201836020820111156101eb57600080fd5b803590602001918460018302840111600160201b8311171561020c57600080fd5b509092509050610599565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561025c578181015183820152602001610244565b50505050905090810190601f1680156102895780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156102a457600080fd5b5061016b6106a2565b3480156102b957600080fd5b5061012b600480360360208110156102d057600080fd5b50356001600160a01b03166106b1565b3480156102ec57600080fd5b506101546004803603604081101561030357600080fd5b506001600160a01b03813516906020013515156106d3565b34801561032757600080fd5b506100d16004803603602081101561033e57600080fd5b503561090f565b34801561035157600080fd5b506101546004803603604081101561036857600080fd5b506001600160a01b038135169060200135151561092d565b34801561038c57600080fd5b50610154600480360360208110156103a357600080fd5b50356001600160a01b0316610b68565b3360009081526001602081905260408220015460ff1661040b576040805162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b604482015290519081900360640190fd5b600654600061041f86864342863a8a610c58565b905060008215610447576006600184038154811061043957fe5b906000526020600020015490505b60066104538284610cc9565b8154600181018355600092835260209283902001556040805133815260ff8a16928101929092526001600160a01b038816828201526060820187905251829185917f23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf79181900360800190a3509095945050505050565b60065490565b6001600160a01b031660009081526002602052604090206001015460ff1690565b6104f8610cf5565b6001600160a01b031661050961058a565b6001600160a01b031614610552576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610d40833981519152908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b3360009081526002602052604081206001015460609060ff166105f5576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b600580546001600160a01b0319811633179091556040516001600160a01b0391821691881690879087908790808383808284376040519201945060009350909150508083038185875af1925050503d806000811461066f576040519150601f19603f3d011682016040523d82523d6000602084013e610674565b606091505b50600580546001600160a01b0319166001600160a01b03949094169390931790925597909650945050505050565b6005546001600160a01b031681565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b6106db610cf5565b6001600160a01b03166106ec61058a565b6001600160a01b031614610735576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b6001600160a01b0382166000908152600260205260409020600181015460ff1680801561075f5750825b80610771575080158015610771575082155b1561077d57505061090b565b821561080c57604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319169091179055610908565b60048054600019810190811061081e57fe5b6000918252602090912001548254600480546001600160a01b0390931692909190811061084757fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460026000600485600001548154811061088f57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560048054806108bf57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff191690555b50505b5050565b6006818154811061091c57fe5b600091825260209091200154905081565b610935610cf5565b6001600160a01b031661094661058a565b6001600160a01b03161461098f576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b6001600160a01b03821660009081526001602081905260409091209081015460ff168080156109bb5750825b806109cd5750801580156109cd575082155b156109d957505061090b565b8215610a6757604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055610908565b600380546000198101908110610a7957fe5b6000918252602090912001548254600380546001600160a01b03909316929091908110610aa257fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610aea57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556003805480610b1a57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff1916905550505050565b610b70610cf5565b6001600160a01b0316610b8161058a565b6001600160a01b031614610bca576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b6001600160a01b038116610c0f5760405162461bcd60e51b8152600401808060200182810382526026815260200180610cfa6026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020610d4083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6001600160601b0319166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212208b6882e0b90e769a155451f0978d5937ce19dc35c6925c62ea997bdaa630206564736f6c634300060c0033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCallerSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeCaller) AllowedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeSession) AllowedInboxes(inbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedInboxes(&_Bridge.CallOpts, inbox)
}

// AllowedInboxes is a free data retrieval call binding the contract method 0xc29372de.
//
// Solidity: function allowedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedInboxes(inbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedInboxes(&_Bridge.CallOpts, inbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, outbox)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCaller) InboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "inboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.InboxAccs(&_Bridge.CallOpts, arg0)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCallerSession) InboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.InboxAccs(&_Bridge.CallOpts, arg0)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Bridge *BridgeCaller) MessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "messageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Bridge *BridgeSession) MessageCount() (*big.Int, error) {
	return _Bridge.Contract.MessageCount(&_Bridge.CallOpts)
}

// MessageCount is a free data retrieval call binding the contract method 0x3dbcc8d1.
//
// Solidity: function messageCount() view returns(uint256)
func (_Bridge *BridgeCallerSession) MessageCount() (*big.Int, error) {
	return _Bridge.Contract.MessageCount(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeTransactor) DeliverMessageToInbox(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deliverMessageToInbox", kind, sender, messageDataHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeSession) DeliverMessageToInbox(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DeliverMessageToInbox(&_Bridge.TransactOpts, kind, sender, messageDataHash)
}

// DeliverMessageToInbox is a paid mutator transaction binding the contract method 0x02bbfad1.
//
// Solidity: function deliverMessageToInbox(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeTransactorSession) DeliverMessageToInbox(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.DeliverMessageToInbox(&_Bridge.TransactOpts, kind, sender, messageDataHash)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactor) ExecuteCall(opts *bind.TransactOpts, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeCall", destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, destAddr, amount, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address destAddr, uint256 amount, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactorSession) ExecuteCall(destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, destAddr, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setInbox", inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, outbox, enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Bridge contract.
type BridgeMessageDeliveredIterator struct {
	Event *BridgeMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMessageDelivered)
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
		it.Event = new(BridgeMessageDelivered)
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
func (it *BridgeMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMessageDelivered represents a MessageDelivered event raised by the Bridge contract.
type BridgeMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_Bridge *BridgeFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeMessageDeliveredIterator{contract: _Bridge.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_Bridge *BridgeFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMessageDelivered)
				if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf7.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash)
func (_Bridge *BridgeFilterer) ParseMessageDelivered(log types.Log) (*BridgeMessageDelivered, error) {
	event := new(BridgeMessageDelivered)
	if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonesABI is the input ABI used to generate the binding from.
const ClonesABI = "[]"

// ClonesBin is the compiled bytecode used for deploying new contracts.
var ClonesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d30fcdeea53525f7e5c478e6c95608c9a572aacd8e5c00a80157e00e2b5ebdb564736f6c634300060c0033"

// DeployClones deploys a new Ethereum contract, binding an instance of Clones to it.
func DeployClones(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clones, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClonesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// Clones is an auto generated Go binding around an Ethereum contract.
type Clones struct {
	ClonesCaller     // Read-only binding to the contract
	ClonesTransactor // Write-only binding to the contract
	ClonesFilterer   // Log filterer for contract events
}

// ClonesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonesSession struct {
	Contract     *Clones           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonesCallerSession struct {
	Contract *ClonesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClonesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonesTransactorSession struct {
	Contract     *ClonesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonesRaw struct {
	Contract *Clones // Generic contract binding to access the raw methods on
}

// ClonesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonesCallerRaw struct {
	Contract *ClonesCaller // Generic read-only contract binding to access the raw methods on
}

// ClonesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonesTransactorRaw struct {
	Contract *ClonesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClones creates a new instance of Clones, bound to a specific deployed contract.
func NewClones(address common.Address, backend bind.ContractBackend) (*Clones, error) {
	contract, err := bindClones(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// NewClonesCaller creates a new read-only instance of Clones, bound to a specific deployed contract.
func NewClonesCaller(address common.Address, caller bind.ContractCaller) (*ClonesCaller, error) {
	contract, err := bindClones(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesCaller{contract: contract}, nil
}

// NewClonesTransactor creates a new write-only instance of Clones, bound to a specific deployed contract.
func NewClonesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonesTransactor, error) {
	contract, err := bindClones(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesTransactor{contract: contract}, nil
}

// NewClonesFilterer creates a new log filterer instance of Clones, bound to a specific deployed contract.
func NewClonesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonesFilterer, error) {
	contract, err := bindClones(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonesFilterer{contract: contract}, nil
}

// bindClones binds a generic wrapper to an already deployed contract.
func bindClones(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.ClonesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transact(opts, method, params...)
}

// IInboxABI is the input ABI used to generate the binding from.
const IInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IInbox is an auto generated Go binding around an Ethereum contract.
type IInbox struct {
	IInboxCaller     // Read-only binding to the contract
	IInboxTransactor // Write-only binding to the contract
	IInboxFilterer   // Log filterer for contract events
}

// IInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInboxSession struct {
	Contract     *IInbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInboxCallerSession struct {
	Contract *IInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInboxTransactorSession struct {
	Contract     *IInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInboxRaw struct {
	Contract *IInbox // Generic contract binding to access the raw methods on
}

// IInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInboxCallerRaw struct {
	Contract *IInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInboxTransactorRaw struct {
	Contract *IInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInbox creates a new instance of IInbox, bound to a specific deployed contract.
func NewIInbox(address common.Address, backend bind.ContractBackend) (*IInbox, error) {
	contract, err := bindIInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInbox{IInboxCaller: IInboxCaller{contract: contract}, IInboxTransactor: IInboxTransactor{contract: contract}, IInboxFilterer: IInboxFilterer{contract: contract}}, nil
}

// NewIInboxCaller creates a new read-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxCaller(address common.Address, caller bind.ContractCaller) (*IInboxCaller, error) {
	contract, err := bindIInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxCaller{contract: contract}, nil
}

// NewIInboxTransactor creates a new write-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IInboxTransactor, error) {
	contract, err := bindIInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxTransactor{contract: contract}, nil
}

// NewIInboxFilterer creates a new log filterer instance of IInbox, bound to a specific deployed contract.
func NewIInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IInboxFilterer, error) {
	contract, err := bindIInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInboxFilterer{contract: contract}, nil
}

// bindIInbox binds a generic wrapper to an already deployed contract.
func bindIInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.IInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IInbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxSession) Bridge() (common.Address, error) {
	return _IInbox.Contract.Bridge(&_IInbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxCallerSession) Bridge() (common.Address, error) {
	return _IInbox.Contract.Bridge(&_IInbox.CallOpts)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_IInbox *IInboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_IInbox *IInboxSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.DeployL2ContractPair(&_IInbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_IInbox *IInboxTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.DeployL2ContractPair(&_IInbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns()
func (_IInbox *IInboxTransactor) DepositEth(opts *bind.TransactOpts, destAddr common.Address) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "depositEth", destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns()
func (_IInbox *IInboxSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth(&_IInbox.TransactOpts, destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns()
func (_IInbox *IInboxTransactorSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth(&_IInbox.TransactOpts, destAddr)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns()
func (_IInbox *IInboxTransactor) SendContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendContractTransaction", maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns()
func (_IInbox *IInboxSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns()
func (_IInbox *IInboxTransactorSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns()
func (_IInbox *IInboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL1FundedContractTransaction", maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns()
func (_IInbox *IInboxSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns()
func (_IInbox *IInboxTransactorSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedContractTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns()
func (_IInbox *IInboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns()
func (_IInbox *IInboxSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns()
func (_IInbox *IInboxTransactorSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_IInbox *IInboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_IInbox *IInboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_IInbox *IInboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns()
func (_IInbox *IInboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns()
func (_IInbox *IInboxSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns()
func (_IInbox *IInboxTransactorSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendUnsignedTransaction(&_IInbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// IInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the IInbox contract.
type IInboxInboxMessageDeliveredIterator struct {
	Event *IInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDelivered)
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
		it.Event = new(IInboxInboxMessageDelivered)
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
func (it *IInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the IInbox contract.
type IInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredIterator{contract: _IInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDelivered)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*IInboxInboxMessageDelivered, error) {
	event := new(IInboxInboxMessageDelivered)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *IInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredFromOriginIterator{contract: _IInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDeliveredFromOrigin)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*IInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(IInboxInboxMessageDeliveredFromOrigin)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxABI is the input ABI used to generate the binding from.
const InboxABI = "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractData\",\"type\":\"bytes\"}],\"name\":\"deployL2ContractPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InboxBin is the compiled bytecode used for deploying new contracts.
var InboxBin = "0x608060405234801561001057600080fd5b50604051610afb380380610afb8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610a96806100656000396000f3fe6080604052600436106100765760003560e01c80631fe927cf1461007b5780635075788b146100f85780635e9167581461019b57806367ef3ab8146102255780636f5dfdca146102b45780638a631aa614610342578063ad9d4ba3146103de578063b75436bb14610404578063e78cea921461047f575b600080fd5b34801561008757600080fd5b506100f66004803603602081101561009e57600080fd5b810190602081018135600160201b8111156100b857600080fd5b8201836020820111156100ca57600080fd5b803590602001918460018302840111600160201b831117156100eb57600080fd5b5090925090506104b0565b005b34801561010457600080fd5b506100f6600480360360c081101561011b57600080fd5b8135916020810135916040820135916001600160a01b03606082013516916080820135919081019060c0810160a0820135600160201b81111561015d57600080fd5b82018360208201111561016f57600080fd5b803590602001918460018302840111600160201b8311171561019057600080fd5b50909250905061054f565b6100f6600480360360808110156101b157600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b8111156101e757600080fd5b8201836020820111156101f957600080fd5b803590602001918460018302840111600160201b8311171561021a57600080fd5b5090925090506105d2565b6100f6600480360360a081101561023b57600080fd5b8135916020810135916040820135916001600160a01b036060820135169181019060a081016080820135600160201b81111561027657600080fd5b82018360208201111561028857600080fd5b803590602001918460018302840111600160201b831117156102a957600080fd5b50909250905061064b565b3480156102c057600080fd5b506100f6600480360360808110156102d757600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561030457600080fd5b82018360208201111561031657600080fd5b803590602001918460018302840111600160201b8311171561033757600080fd5b5090925090506106cd565b34801561034e57600080fd5b506100f6600480360360a081101561036557600080fd5b8135916020810135916001600160a01b036040830135169160608101359181019060a081016080820135600160201b8111156103a057600080fd5b8201836020820111156103b257600080fd5b803590602001918460018302840111600160201b831117156103d357600080fd5b509092509050610770565b6100f6600480360360208110156103f457600080fd5b50356001600160a01b03166107e2565b34801561041057600080fd5b506100f66004803603602081101561042757600080fd5b810190602081018135600160201b81111561044157600080fd5b82018360208201111561045357600080fd5b803590602001918460018302840111600160201b8311171561047457600080fd5b50909250905061083a565b34801561048b57600080fd5b506104946108ba565b604080516001600160a01b039092168252519081900360200190f35b3332146104f2576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600061051c600333858560405180838380828437604051920182900390912093506108c992505050565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a2505050565b6105c960033360008a8a8a8a60601b60601c6001600160a01b03168a8a8a604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610960565b50505050505050565b610644600733600188888860601b60601c6001600160a01b0316348989604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610960565b5050505050565b6106c560073360008989898960601b60601c6001600160a01b0316348a8a604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610960565b505050505050565b6106d633610a04565b610724576040805162461bcd60e51b815260206004820152601a6024820152791b5d5cdd0818994818d85b1b195908189e4818dbdb9d1c9858dd60321b604482015290519081900360640190fd5b6106446005338787878787604051602001808681526020018581526020018481526020018383808284378083019250505095505050505050604051602081830303815290604052610960565b6106c5600333600189898960601b60601c6001600160a01b0316898989604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610960565b60408051600160f81b602082015260006021820181905260418201526001600160a01b0383166061820152346081808301919091528251808303909101815260a1909101909152610837906007908390610960565b50565b6000610864600333858560405180838380828437604051920182900390912093506108c992505050565b905080600080516020610a41833981519152848460405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2505050565b6000546001600160a01b031681565b60008054604080516302bbfad160e01b815260ff871660048201526001600160a01b03868116602483015260448201869052915191909216916302bbfad191349160648082019260209290919082900301818588803b15801561092b57600080fd5b505af115801561093f573d6000803e3d6000fd5b50505050506040513d602081101561095657600080fd5b5051949350505050565b6000610974848484805190602001206108c9565b905080600080516020610a41833981519152836040518080602001828103825283818151815260200191508051906020019080838360005b838110156109c45781810151838201526020016109ac565b50505050905090810190601f1680156109f15780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610a3857508115155b94935050505056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba2646970667358221220c69f0f2def64d34c144be3e1b5bf9c8d8644a74e6c602ebc3e7570d6b1afa55a64736f6c634300060c0033"

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InboxBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCallerSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactor) DeployL2ContractPair(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "deployL2ContractPair", maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DeployL2ContractPair is a paid mutator transaction binding the contract method 0x6f5dfdca.
//
// Solidity: function deployL2ContractPair(uint256 maxGas, uint256 gasPriceBid, uint256 payment, bytes contractData) returns()
func (_Inbox *InboxTransactorSession) DeployL2ContractPair(maxGas *big.Int, gasPriceBid *big.Int, payment *big.Int, contractData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.DeployL2ContractPair(&_Inbox.TransactOpts, maxGas, gasPriceBid, payment, contractData)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns()
func (_Inbox *InboxTransactor) DepositEth(opts *bind.TransactOpts, destAddr common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth", destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns()
func (_Inbox *InboxSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, destAddr)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address destAddr) payable returns()
func (_Inbox *InboxTransactorSession) DepositEth(destAddr common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, destAddr)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns()
func (_Inbox *InboxTransactor) SendContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendContractTransaction", maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns()
func (_Inbox *InboxSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, uint256 amount, bytes data) returns()
func (_Inbox *InboxTransactorSession) SendContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, amount, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns()
func (_Inbox *InboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedContractTransaction", maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns()
func (_Inbox *InboxSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 maxGas, uint256 gasPriceBid, address destAddr, bytes data) payable returns()
func (_Inbox *InboxTransactorSession) SendL1FundedContractTransaction(maxGas *big.Int, gasPriceBid *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns()
func (_Inbox *InboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns()
func (_Inbox *InboxSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, bytes data) payable returns()
func (_Inbox *InboxTransactorSession) SendL1FundedUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns()
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns()
func (_Inbox *InboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendUnsignedTransaction", maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns()
func (_Inbox *InboxSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 maxGas, uint256 gasPriceBid, uint256 nonce, address destAddr, uint256 amount, bytes data) returns()
func (_Inbox *InboxTransactorSession) SendUnsignedTransaction(maxGas *big.Int, gasPriceBid *big.Int, nonce *big.Int, destAddr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, maxGas, gasPriceBid, nonce, destAddr, amount, data)
}

// InboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the Inbox contract.
type InboxInboxMessageDeliveredIterator struct {
	Event *InboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDelivered)
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
		it.Event = new(InboxInboxMessageDelivered)
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
func (it *InboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the Inbox contract.
type InboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredIterator{contract: _Inbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxInboxMessageDelivered, error) {
	event := new(InboxInboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxInboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxABI is the input ABI used to generate the binding from.
const OutboxABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numInBatch\",\"type\":\"uint256\"}],\"name\":\"OutboxEntryCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Contract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdSuccessfully\",\"type\":\"bool\"}],\"name\":\"executeBuddyContractReceipt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outboxIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"calldataForL1\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"}],\"name\":\"processOutgoingMessages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OutboxBin is the compiled bytecode used for deploying new contracts.
var OutboxBin = "0x608060405234801561001057600080fd5b5060405161168f38038061168f8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b038085166001600160a01b0319928316179092556001805492841692909116919091179055604051610078906100bd565b604051809103906000f080158015610094573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b0392909216919091179055506100ca9050565b61059a806110f583390190565b61101c806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630c72684714610067578063465477901461012757806372fea01f1461014157806380648b0214610238578063b0f3053714610127578063d933556a1461025c575b600080fd5b6101256004803603604081101561007d57600080fd5b810190602081018135600160201b81111561009757600080fd5b8201836020820111156100a957600080fd5b803590602001918460018302840111600160201b831117156100ca57600080fd5b919390929091602081019035600160201b8111156100e757600080fd5b8201836020820111156100f957600080fd5b803590602001918460208302840111600160201b8311171561011a57600080fd5b5090925090506102e8565b005b61012f6103cf565b60408051918252519081900360200190f35b610125600480360361012081101561015857600080fd5b81359190810190604081016020820135600160201b81111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460208302840111600160201b831117156101ac57600080fd5b919390928235926001600160a01b03602082013581169360408301359091169260608301359260808101359260a082013592909160e081019060c00135600160201b8111156101fa57600080fd5b82018360208201111561020c57600080fd5b803590602001918460018302840111600160201b8311171561022d57600080fd5b5090925090506103de565b6102406103cf565b604080516001600160a01b039092168252519081900360200190f35b610125600480360360a081101561027257600080fd5b81359190810190604081016020820135600160201b81111561029357600080fd5b8201836020820111156102a557600080fd5b803590602001918460208302840111600160201b831117156102c657600080fd5b91935091508035906001600160a01b0360208201351690604001351515610597565b6000546001600160a01b03163314610335576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b806000805b828110156103c6576103a487838888888681811061035457fe5b9050602002013586019261036a93929190610fbe565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061065e92505050565b8484828181106103b057fe5b602002919091013592909201915060010161033a565b50505050505050565b6004546001600160a01b031690565b6000808860601b60601c6001600160a01b03168860601b60601c6001600160a01b031688888888886040516020018089815260200188815260200187815260200186815260200185815260200184815260200183838082843780830192505050985050505050505050506040516020818303038152906040528051906020012090506104a18c8c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508691506107db9050565b60048054600580546001600160a01b038c81166001600160a01b03198516179094556001600160801b03898116600160801b9081028c83166001600160801b0319851617831617909355604080516020601f8a0181900481028201810190925288815295909416948183169493909204169161053c918c918a91908a908a9081908401838280828437600092019190915250610a6192505050565b600480546001600160a01b039094166001600160a01b031990941693909317909255600580546001600160801b03938416600160801b029284166001600160801b031990911617909216179055505050505050505050505050565b6040805160056020808301919091526001600160a01b0385168284015283151560f81b606083015282516041818403018152606183018085528151918301919091206081928902808501840190955288825293610614938b938b928b92839201908490808284376000920191909152508992508691506107db9050565b604080518315156024808301919091528251808303909101815260449091019091526020810180516001600160e01b03166323d82c5d60e21b1790526103c6908490600090610c7e565b80516000908290829061066d57fe5b01602001516001600160f81b03191614156107d857600061068f826001610cc9565b9050600061069e836021610cc9565b905060006106ad846041610cc9565b6002549091506000906106c8906001600160a01b0316610d22565b60015460408051633422b05160e11b81526001600160a01b039283166004820152602481018690526044810187905290519293509083169163684560a29160648082019260009290919082900301818387803b15801561072757600080fd5b505af115801561073b573d6000803e3d6000fd5b5050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810180546001600160a01b0386166001600160a01b0319909116179055604080518281526020810187905280820188905290519193508792507fe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131919081900360600190a250505050505b50565b61010083511115610824576040805162461bcd60e51b815260206004820152600e60248201526d50524f4f465f544f4f5f4c4f4e4760901b604482015290519081900360640190fd5b825160020a821061086f576040805162461bcd60e51b815260206004820152601060248201526f1410551217d393d517d352539253505360821b604482015290519081900360640190fd5b60006108a28484846040516020018082815260200191505060405160208183030381529060405280519060200120610e4e565b90506000600386815481106108b357fe5b6000918252602090912001546001600160a01b0316905080610908576040805162461bcd60e51b815260206004820152600960248201526809c9ebe9eaaa8849eb60bb1b604482015290519081900360640190fd5b8451604080516020808201889052818301939093528151808203830181526060820183528051908401206084820186905260a48083018290528351808403909101815260c490920190925291820180516001600160e01b0316630ad0379b60e01b1790529061097b908390600090610c7e565b816001600160a01b0316635780e4e76040518163ffffffff1660e01b815260040160206040518083038186803b1580156109b457600080fd5b505afa1580156109c8573d6000803e3d6000fd5b505050506040513d60208110156109de57600080fd5b50516103c6576040805160048152602481019091526020810180516001600160e01b031663083197ef60e41b179052610a1b908390600090610c7e565b600060038881548110610a2a57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050565b600154604051639e5d4c4960e01b81526001600160a01b03858116600483019081526024830186905260606044840181815286516064860152865160009692959490921693639e5d4c49938a938a938a93909160849091019060208501908083838e5b83811015610adc578181015183820152602001610ac4565b50505050905090810190601f168015610b095780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610b2a57600080fd5b505af1158015610b3e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610b6757600080fd5b815160208301805160405192949293830192919084600160201b821115610b8d57600080fd5b908301906020820185811115610ba257600080fd5b8251600160201b811182820188101715610bbb57600080fd5b82525081516020918201929091019080838360005b83811015610be8578181015183820152602001610bd0565b50505050905090810190601f168015610c155780820380516001836020036101000a031916815260200191505b506040525050509150915081610c7757805115610c355780518082602001fd5b6040805162461bcd60e51b81526020600482015260126024820152711094925111d157d0d0531317d1905253115160721b604482015290519081900360640190fd5b5050505050565b600480546001600160a01b031981169091556001600160a01b0316610ca4848484610a61565b600480546001600160a01b0319166001600160a01b0392909216919091179055505050565b60008160200183511015610d19576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b158015610d5d57600080fd5b505afa158015610d71573d6000803e3d6000fd5b505050506040513d6020811015610d8757600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b602082015290610e355760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610dfa578181015183820152602001610de2565b50505050905090810190601f168015610e275780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610e48826001600160a01b0316610f1c565b92915050565b8251600090610100811115610e6257600080fd5b8260005b82811015610f125760028606610ebf57868181518110610e8257fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150610f04565b81878281518110610ecc57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101610e66565b5095945050505050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610fb9576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b919050565b60008085851115610fcd578182fd5b83861115610fd9578182fd5b505082019391909203915056fea26469706673582212207a6f9fa675d15467bc4ffc2121ab858ac426793d7dbbdee246af9719f33f7c3564736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff1916600117905561056d8061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c578063684560a2146100b15780636f791d29146100e357806383197ef0146100ff5780639db9af8114610107578063ebf0c71714610124575b600080fd5b61007a61012c565b60408051918252519081900360200190f35b6100af600480360360408110156100a257600080fd5b5080359060200135610132565b005b6100af600480360360608110156100c757600080fd5b506001600160a01b038135169060208101359060400135610205565b6100eb6102b5565b604080519115158252519081900360200190f35b6100af6102be565b6100eb6004803603602081101561011d57600080fd5b50356102d1565b61007a6102e6565b60025481565b61013a6102ec565b60008181526003602052604090205460ff161561018e576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b60015482146101cf576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b6000818152600360205260409020805460ff19166001179055600280546000190190819055610201576102013361047a565b5050565b60015415610249576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b81610286576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b600080546001600160a01b0390941661010002610100600160a81b031990941693909317909255600155600255565b60005460ff1690565b6102c66102ec565b6102cf3361047a565b565b60036020526000908152604090205460ff1681565b60015481565b60005461010090046001600160a01b0316331461033e576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f42524944474560a81b604482015290519081900360640190fd5b60006001600160a01b0316600060019054906101000a90046001600160a01b03166001600160a01b031663ab5d89436040518163ffffffff1660e01b815260040160206040518083038186803b15801561039757600080fd5b505afa1580156103ab573d6000803e3d6000fd5b505050506040513d60208110156103c157600080fd5b505160408051634032458160e11b815290516001600160a01b03909216916380648b0291600480820192602092909190829003018186803b15801561040557600080fd5b505afa158015610419573d6000803e3d6000fd5b505050506040513d602081101561042f57600080fd5b50516001600160a01b0316146102cf576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f53595354454d60a81b604482015290519081900360640190fd5b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff161561052a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104ef5781810151838201526020016104d7565b50505050905090810190601f16801561051c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea26469706673582212203b9da5963c6361e0f514db3f5dae9eca4924dbd009dfdf6859b88da66ef9b44664736f6c634300060c0033"

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend, _rollup common.Address, _bridge common.Address) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxBin), backend, _rollup, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// ExecuteBuddyContractReceipt is a paid mutator transaction binding the contract method 0xd933556a.
//
// Solidity: function executeBuddyContractReceipt(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Contract, bool createdSuccessfully) returns()
func (_Outbox *OutboxTransactor) ExecuteBuddyContractReceipt(opts *bind.TransactOpts, outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Contract common.Address, createdSuccessfully bool) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeBuddyContractReceipt", outboxIndex, proof, index, l2Contract, createdSuccessfully)
}

// ExecuteBuddyContractReceipt is a paid mutator transaction binding the contract method 0xd933556a.
//
// Solidity: function executeBuddyContractReceipt(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Contract, bool createdSuccessfully) returns()
func (_Outbox *OutboxSession) ExecuteBuddyContractReceipt(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Contract common.Address, createdSuccessfully bool) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteBuddyContractReceipt(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Contract, createdSuccessfully)
}

// ExecuteBuddyContractReceipt is a paid mutator transaction binding the contract method 0xd933556a.
//
// Solidity: function executeBuddyContractReceipt(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Contract, bool createdSuccessfully) returns()
func (_Outbox *OutboxTransactorSession) ExecuteBuddyContractReceipt(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Contract common.Address, createdSuccessfully bool) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteBuddyContractReceipt(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Contract, createdSuccessfully)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x72fea01f.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", outboxIndex, proof, index, l2Sender, destAddr, l2Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x72fea01f.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Sender, destAddr, l2Block, l2Timestamp, amount, calldataForL1)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x72fea01f.
//
// Solidity: function executeTransaction(uint256 outboxIndex, bytes32[] proof, uint256 index, address l2Sender, address destAddr, uint256 l2Block, uint256 l2Timestamp, uint256 amount, bytes calldataForL1) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(outboxIndex *big.Int, proof [][32]byte, index *big.Int, l2Sender common.Address, destAddr common.Address, l2Block *big.Int, l2Timestamp *big.Int, amount *big.Int, calldataForL1 []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, outboxIndex, proof, index, l2Sender, destAddr, l2Block, l2Timestamp, amount, calldataForL1)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactor) ProcessOutgoingMessages(opts *bind.TransactOpts, sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "processOutgoingMessages", sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// ProcessOutgoingMessages is a paid mutator transaction binding the contract method 0x0c726847.
//
// Solidity: function processOutgoingMessages(bytes sendsData, uint256[] sendLengths) returns()
func (_Outbox *OutboxTransactorSession) ProcessOutgoingMessages(sendsData []byte, sendLengths []*big.Int) (*types.Transaction, error) {
	return _Outbox.Contract.ProcessOutgoingMessages(&_Outbox.TransactOpts, sendsData, sendLengths)
}

// OutboxOutboxEntryCreatedIterator is returned from FilterOutboxEntryCreated and is used to iterate over the raw logs and unpacked data for OutboxEntryCreated events raised by the Outbox contract.
type OutboxOutboxEntryCreatedIterator struct {
	Event *OutboxOutboxEntryCreated // Event containing the contract specifics and raw log

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
func (it *OutboxOutboxEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxOutboxEntryCreated)
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
		it.Event = new(OutboxOutboxEntryCreated)
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
func (it *OutboxOutboxEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxOutboxEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxOutboxEntryCreated represents a OutboxEntryCreated event raised by the Outbox contract.
type OutboxOutboxEntryCreated struct {
	BatchNum    *big.Int
	OutboxIndex *big.Int
	OutputRoot  [32]byte
	NumInBatch  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOutboxEntryCreated is a free log retrieval operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) FilterOutboxEntryCreated(opts *bind.FilterOpts, batchNum []*big.Int) (*OutboxOutboxEntryCreatedIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &OutboxOutboxEntryCreatedIterator{contract: _Outbox.contract, event: "OutboxEntryCreated", logs: logs, sub: sub}, nil
}

// WatchOutboxEntryCreated is a free log subscription operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) WatchOutboxEntryCreated(opts *bind.WatchOpts, sink chan<- *OutboxOutboxEntryCreated, batchNum []*big.Int) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "OutboxEntryCreated", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxOutboxEntryCreated)
				if err := _Outbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
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

// ParseOutboxEntryCreated is a log parse operation binding the contract event 0xe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131.
//
// Solidity: event OutboxEntryCreated(uint256 indexed batchNum, uint256 outboxIndex, bytes32 outputRoot, uint256 numInBatch)
func (_Outbox *OutboxFilterer) ParseOutboxEntryCreated(log types.Log) (*OutboxOutboxEntryCreated, error) {
	event := new(OutboxOutboxEntryCreated)
	if err := _Outbox.contract.UnpackLog(event, "OutboxEntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxEntryABI is the input ABI used to generate the binding from.
const OutboxEntryABI = "[{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_numInBatch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"spendOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spentOutput\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OutboxEntryBin is the compiled bytecode used for deploying new contracts.
var OutboxEntryBin = "0x608060405234801561001057600080fd5b506000805460ff1916600117905561056d8061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c578063684560a2146100b15780636f791d29146100e357806383197ef0146100ff5780639db9af8114610107578063ebf0c71714610124575b600080fd5b61007a61012c565b60408051918252519081900360200190f35b6100af600480360360408110156100a257600080fd5b5080359060200135610132565b005b6100af600480360360608110156100c757600080fd5b506001600160a01b038135169060208101359060400135610205565b6100eb6102b5565b604080519115158252519081900360200190f35b6100af6102be565b6100eb6004803603602081101561011d57600080fd5b50356102d1565b61007a6102e6565b60025481565b61013a6102ec565b60008181526003602052604090205460ff161561018e576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b60015482146101cf576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b6000818152600360205260409020805460ff19166001179055600280546000190190819055610201576102013361047a565b5050565b60015415610249576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b81610286576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b600080546001600160a01b0390941661010002610100600160a81b031990941693909317909255600155600255565b60005460ff1690565b6102c66102ec565b6102cf3361047a565b565b60036020526000908152604090205460ff1681565b60015481565b60005461010090046001600160a01b0316331461033e576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f42524944474560a81b604482015290519081900360640190fd5b60006001600160a01b0316600060019054906101000a90046001600160a01b03166001600160a01b031663ab5d89436040518163ffffffff1660e01b815260040160206040518083038186803b15801561039757600080fd5b505afa1580156103ab573d6000803e3d6000fd5b505050506040513d60208110156103c157600080fd5b505160408051634032458160e11b815290516001600160a01b03909216916380648b0291600480820192602092909190829003018186803b15801561040557600080fd5b505afa158015610419573d6000803e3d6000fd5b505050506040513d602081101561042f57600080fd5b50516001600160a01b0316146102cf576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f53595354454d60a81b604482015290519081900360640190fd5b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff161561052a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104ef5781810151838201526020016104d7565b50505050905090810190601f16801561051c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea26469706673582212203b9da5963c6361e0f514db3f5dae9eca4924dbd009dfdf6859b88da66ef9b44664736f6c634300060c0033"

// DeployOutboxEntry deploys a new Ethereum contract, binding an instance of OutboxEntry to it.
func DeployOutboxEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OutboxEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// OutboxEntry is an auto generated Go binding around an Ethereum contract.
type OutboxEntry struct {
	OutboxEntryCaller     // Read-only binding to the contract
	OutboxEntryTransactor // Write-only binding to the contract
	OutboxEntryFilterer   // Log filterer for contract events
}

// OutboxEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxEntrySession struct {
	Contract     *OutboxEntry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxEntryCallerSession struct {
	Contract *OutboxEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OutboxEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxEntryTransactorSession struct {
	Contract     *OutboxEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OutboxEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxEntryRaw struct {
	Contract *OutboxEntry // Generic contract binding to access the raw methods on
}

// OutboxEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxEntryCallerRaw struct {
	Contract *OutboxEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxEntryTransactorRaw struct {
	Contract *OutboxEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxEntry creates a new instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntry(address common.Address, backend bind.ContractBackend) (*OutboxEntry, error) {
	contract, err := bindOutboxEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxEntry{OutboxEntryCaller: OutboxEntryCaller{contract: contract}, OutboxEntryTransactor: OutboxEntryTransactor{contract: contract}, OutboxEntryFilterer: OutboxEntryFilterer{contract: contract}}, nil
}

// NewOutboxEntryCaller creates a new read-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryCaller(address common.Address, caller bind.ContractCaller) (*OutboxEntryCaller, error) {
	contract, err := bindOutboxEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryCaller{contract: contract}, nil
}

// NewOutboxEntryTransactor creates a new write-only instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxEntryTransactor, error) {
	contract, err := bindOutboxEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryTransactor{contract: contract}, nil
}

// NewOutboxEntryFilterer creates a new log filterer instance of OutboxEntry, bound to a specific deployed contract.
func NewOutboxEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxEntryFilterer, error) {
	contract, err := bindOutboxEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxEntryFilterer{contract: contract}, nil
}

// bindOutboxEntry binds a generic wrapper to an already deployed contract.
func bindOutboxEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.OutboxEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.OutboxEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxEntry *OutboxEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxEntry *OutboxEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxEntry.Contract.contract.Transact(opts, method, params...)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntrySession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) IsMaster() (bool, error) {
	return _OutboxEntry.Contract.IsMaster(&_OutboxEntry.CallOpts)
}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntryCaller) NumRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "numRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntrySession) NumRemaining() (*big.Int, error) {
	return _OutboxEntry.Contract.NumRemaining(&_OutboxEntry.CallOpts)
}

// NumRemaining is a free data retrieval call binding the contract method 0x5780e4e7.
//
// Solidity: function numRemaining() view returns(uint256)
func (_OutboxEntry *OutboxEntryCallerSession) NumRemaining() (*big.Int, error) {
	return _OutboxEntry.Contract.NumRemaining(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntrySession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_OutboxEntry *OutboxEntryCallerSession) Root() ([32]byte, error) {
	return _OutboxEntry.Contract.Root(&_OutboxEntry.CallOpts)
}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCaller) SpentOutput(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _OutboxEntry.contract.Call(opts, &out, "spentOutput", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntrySession) SpentOutput(arg0 [32]byte) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// SpentOutput is a free data retrieval call binding the contract method 0x9db9af81.
//
// Solidity: function spentOutput(bytes32 ) view returns(bool)
func (_OutboxEntry *OutboxEntryCallerSession) SpentOutput(arg0 [32]byte) (bool, error) {
	return _OutboxEntry.Contract.SpentOutput(&_OutboxEntry.CallOpts, arg0)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntryTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntrySession) Destroy() (*types.Transaction, error) {
	return _OutboxEntry.Contract.Destroy(&_OutboxEntry.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Destroy() (*types.Transaction, error) {
	return _OutboxEntry.Contract.Destroy(&_OutboxEntry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x684560a2.
//
// Solidity: function initialize(address _bridge, bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntryTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, _root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "initialize", _bridge, _root, _numInBatch)
}

// Initialize is a paid mutator transaction binding the contract method 0x684560a2.
//
// Solidity: function initialize(address _bridge, bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntrySession) Initialize(_bridge common.Address, _root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _bridge, _root, _numInBatch)
}

// Initialize is a paid mutator transaction binding the contract method 0x684560a2.
//
// Solidity: function initialize(address _bridge, bytes32 _root, uint256 _numInBatch) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) Initialize(_bridge common.Address, _root [32]byte, _numInBatch *big.Int) (*types.Transaction, error) {
	return _OutboxEntry.Contract.Initialize(&_OutboxEntry.TransactOpts, _bridge, _root, _numInBatch)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns()
func (_OutboxEntry *OutboxEntryTransactor) SpendOutput(opts *bind.TransactOpts, _root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.contract.Transact(opts, "spendOutput", _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns()
func (_OutboxEntry *OutboxEntrySession) SpendOutput(_root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// SpendOutput is a paid mutator transaction binding the contract method 0x57d61c0b.
//
// Solidity: function spendOutput(bytes32 _root, bytes32 _id) returns()
func (_OutboxEntry *OutboxEntryTransactorSession) SpendOutput(_root [32]byte, _id [32]byte) (*types.Transaction, error) {
	return _OutboxEntry.Contract.SpendOutput(&_OutboxEntry.TransactOpts, _root, _id)
}

// RollupCreatorABI is the input ABI used to generate the binding from.
const RollupCreatorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"createRollup\",\"outputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"_rollupTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"}],\"name\":\"setTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupCreatorBin is the compiled bytecode used for deploying new contracts.
var RollupCreatorBin = "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6159658061007d6000396000f3fe60806040523480156200001157600080fd5b50600436106200005e5760003560e01c8063715018a6146200006357806384b99970146200006f5780638da5cb5b1462000139578063d92208241462000143578063f2fde38b146200017e575b600080fd5b6200006d620001a7565b005b6200011d60048036036101008110156200008857600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c08301359091169190810190610100810160e0820135600160201b811115620000dc57600080fd5b820183602082011115620000ef57600080fd5b803590602001918460018302840111600160201b831117156200011157600080fd5b50909250905062000248565b604080516001600160a01b039092168252519081900360200190f35b6200011d620002e2565b6200006d600480360360608110156200015b57600080fd5b506001600160a01b038135811691602081013582169160409091013516620002f1565b6200006d600480360360208110156200019657600080fd5b50356001600160a01b031662000398565b620001b162000491565b6001600160a01b0316620001c4620002e2565b6001600160a01b0316146200020f576040805162461bcd60e51b81526020600482018190526024820152600080516020620058f0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b039091169060008051602062005910833981519152908390a3600080546001600160a01b0319169055565b6000620002d46040518061010001604052808c81526020018b81526020018a8152602001898152602001888152602001876001600160a01b03168152602001866001600160a01b0316815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505091525062000495565b9a9950505050505050505050565b6000546001600160a01b031690565b620002fb62000491565b6001600160a01b03166200030e620002e2565b6001600160a01b03161462000359576040805162461bcd60e51b81526020600482018190526024820152600080516020620058f0833981519152604482015290519081900360640190fd5b600180546001600160a01b039485166001600160a01b031991821617909155600280549385169382169390931790925560038054919093169116179055565b620003a262000491565b6001600160a01b0316620003b5620002e2565b6001600160a01b03161462000400576040805162461bcd60e51b81526020600482018190526024820152600080516020620058f0833981519152604482015290519081900360640190fd5b6001600160a01b038116620004475760405162461bcd60e51b8152600401808060200182810382526026815260200180620058ca6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216916000805160206200591083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b6000620004a162000a4d565b604051620004af9062000a82565b604051809103906000f080158015620004cc573d6000803e3d6000fd5b506001600160a01b03908116808352600154604051921691620004ef9062000a90565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f08015801562000533573d6000803e3d6000fd5b506001600160a01b031660a0820152604051620005509062000a9e565b604051809103906000f0801580156200056d573d6000803e3d6000fd5b506001600160a01b0316602082018190526040516200058c9062000aac565b6001600160a01b03909116815260405190819003602001906000f080158015620005ba573d6000803e3d6000fd5b506001600160a01b0316604080830191909152602082015160a08301519151909190620005e79062000aba565b6001600160a01b03928316815291166020820152604080519182900301906000f0801580156200061b573d6000803e3d6000fd5b506001600160a01b0390811660608301526020820151604080840151815163722dbe7360e11b81529084166004820152600160248201529051919092169163e45b7ce691604480830192600092919082900301818387803b1580156200068057600080fd5b505af115801562000695573d6000803e3d6000fd5b505050508060a001518160200151604051620006b19062000ac8565b6001600160a01b03928316815291166020820152604080519182900301906000f080158015620006e5573d6000803e3d6000fd5b506001600160a01b039081166080830152602082015160a08301516040805163f2fde38b60e01b8152918416600483015251919092169163f2fde38b91602480830192600092919082900301818387803b1580156200074357600080fd5b505af115801562000758573d6000803e3d6000fd5b5050505080600001516001600160a01b031663f2fde38b8260a001516040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015620007b457600080fd5b505af1158015620007c9573d6000803e3d6000fd5b505050508060a001516001600160a01b031663fdaf5797846000015185602001518660400151876060015188608001518960a001518a60c001518b60e001516040518060c001604052808c600001516001600160a01b03166001600160a01b031681526020018c602001516001600160a01b03166001600160a01b031681526020018c608001516001600160a01b03166001600160a01b031681526020018c606001516001600160a01b03166001600160a01b03168152602001600260009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600360009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152506040518a63ffffffff1660e01b8152600401808a8152602001898152602001888152602001878152602001868152602001856001600160a01b03168152602001846001600160a01b031681526020018060200183600660200280838360005b83811015620009545781810151838201526020016200093a565b50505050905001828103825284818151815260200191508051906020019080838360005b838110156200099257818101518382015260200162000978565b50505050905090810190601f168015620009c05780820380516001836020036101000a031916815260200191505b509a5050505050505050505050600060405180830381600087803b158015620009e857600080fd5b505af1158015620009fd573d6000803e3d6000fd5b50505060a0820151604080516001600160a01b039092168252517f84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c92509081900360200190a160a0015192915050565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b6108e28062000ad783390190565b610c7480620013b983390190565b610e12806200202d83390190565b610afb8062002e3f83390190565b610901806200393a83390190565b61168f806200423b8339019056fe608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6108658061007d6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e3578063f2fde38b1461021e578063f3b7dead14610251575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610284565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610316565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103b0565b34801561011d57600080fd5b506100a361047d565b6100d46004803603606081101561013c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460018302840111600160201b831117156101a257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048c945050505050565b3480156101ef57600080fd5b506100d46004803603604081101561020657600080fd5b506001600160a01b03813581169160200135166105c5565b34801561022a57600080fd5b506100d46004803603602081101561024157600080fd5b50356001600160a01b0316610676565b34801561025d57600080fd5b506100a36004803603602081101561027457600080fd5b50356001600160a01b0316610766565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b606091505b5091509150816102f757600080fd5b80806020019051602081101561030c57600080fd5b5051949350505050565b61031e6107c5565b6001600160a01b031661032f61047d565b6001600160a01b031614610378576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610810833981519152908390a3600080546001600160a01b0319169055565b6103b86107c5565b6001600160a01b03166103c961047d565b6001600160a01b031614610412576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b505af1158015610475573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b6104946107c5565b6001600160a01b03166104a561047d565b6001600160a01b0316146104ee576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561055b578181015183820152602001610543565b50505050905090810190601f1680156105885780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b5050505050505050565b6105cd6107c5565b6001600160a01b03166105de61047d565b6001600160a01b031614610627576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561046157600080fd5b61067e6107c5565b6001600160a01b031661068f61047d565b6001600160a01b0316146106d8576040805162461bcd60e51b815260206004820181905260248201526000805160206107f0833981519152604482015290519081900360640190fd5b6001600160a01b03811661071d5760405162461bcd60e51b81526004018080602001828103825260268152602001806107ca6026913960400191505060405180910390fd5b600080546040516001600160a01b038085169392169160008051602061081083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e3576040519150601f19603f3d011682016040523d82523d6000602084013e6102e8565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122036b5070566a1002642d9947563bea5cf538675280563c6ec32ab459f386727a264736f6c634300060c0033608060405260405162000c7438038062000c74833981810160405260608110156200002957600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200005557600080fd5b9083019060208201858111156200006b57600080fd5b82516401000000008111828201881017156200008657600080fd5b82525081516020918201929091019080838360005b83811015620000b55781810151838201526020016200009b565b50505050905090810190601f168015620000e35780820380516001836020036101000a031916815260200191505b5060405250849150829050620000f98262000137565b8051156200011a57620001188282620001ae60201b620003821760201c565b505b50620001239050565b6200012e82620001dd565b505050620003bf565b6200014d816200020160201b620003ae1760201c565b6200018a5760405162461bcd60e51b815260040180806020018281038252603681526020018062000c186036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6060620001d6838360405180606001604052806027815260200162000bf16027913962000207565b9392505050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b3b151590565b6060620002148462000201565b620002515760405162461bcd60e51b815260040180806020018281038252602681526020018062000c4e6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b60208310620002915780518252601f19909201916020918201910162000270565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114620002f3576040519150601f19603f3d011682016040523d82523d6000602084013e620002f8565b606091505b5090925090506200030b82828662000315565b9695505050505050565b6060831562000326575081620001d6565b825115620003375782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200038357818101518382015260200162000369565b50505050905090810190601f168015620003b15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b61082280620003cf6000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101165780638f28397014610147578063f851a4401461017a5761005d565b3661005d5761005b61018f565b005b61005b61018f565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101a9565b61005b600480360360408110156100ae57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460018302840111600160201b8311171561010b57600080fd5b5090925090506101e3565b34801561012257600080fd5b5061012b610260565b604080516001600160a01b039092168252519081900360200190f35b34801561015357600080fd5b5061005b6004803603602081101561016a57600080fd5b50356001600160a01b031661029d565b34801561018657600080fd5b5061012b610357565b6101976103b4565b6101a76101a2610414565b610427565b565b6101b161044b565b6001600160a01b0316336001600160a01b031614156101d8576101d38161045e565b6101e0565b6101e061018f565b50565b6101eb61044b565b6001600160a01b0316336001600160a01b031614156102535761020d8361045e565b61024d8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038292505050565b5061025b565b61025b61018f565b505050565b600061026a61044b565b6001600160a01b0316336001600160a01b031614156102925761028b610414565b905061029a565b61029a61018f565b90565b6102a561044b565b6001600160a01b0316336001600160a01b031614156101d8576001600160a01b0381166103035760405162461bcd60e51b815260040180806020018281038252603a8152602001806106ce603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61032c61044b565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d38161049e565b600061036161044b565b6001600160a01b0316336001600160a01b031614156102925761028b61044b565b60606103a78383604051806060016040528060278152602001610728602791396104b0565b9392505050565b3b151590565b6103bc61044b565b6001600160a01b0316336001600160a01b0316141561040c5760405162461bcd60e51b81526004018080602001828103825260428152602001806107ab6042913960600191505060405180910390fd5b6101a76101a7565b6000805160206107088339815191525490565b3660008037600080366000845af43d6000803e808015610446573d6000f35b3d6000fd5b6000805160206106ae8339815191525490565b610467816105b3565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6000805160206106ae83398151915255565b60606104bb846103ae565b6104f65760405162461bcd60e51b81526004018080602001828103825260268152602001806107856026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106105345780518252601f199092019160209182019101610515565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610594576040519150601f19603f3d011682016040523d82523d6000602084013e610599565b606091505b50915091506105a9828286610609565b9695505050505050565b6105bc816103ae565b6105f75760405162461bcd60e51b815260040180806020018281038252603681526020018061074f6036913960400191505060405180910390fd5b60008051602061070883398151915255565b606083156106185750816103a7565b8251156106285782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561067257818101518382015260200161065a565b50505050905090810190601f16801561069f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfeb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a264697066735822122081df10ca698e92e05c31196844f9e9c309ede61eb7e45463098c415a459a179564736f6c634300060c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e7472616374608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b610d958061007d6000396000f3fe6080604052600436106100975760003560e01c806302bbfad11461009c5780633dbcc8d1146100e3578063413b35bd146100f8578063715018a61461013f5780638da5cb5b146101565780639e5d4c4914610187578063ab5d894314610298578063c29372de146102ad578063cee3d728146102e0578063d9dd67ab1461031b578063e45b7ce614610345578063f2fde38b14610380575b600080fd5b6100d1600480360360608110156100b257600080fd5b5060ff813516906001600160a01b0360208201351690604001356103b3565b60408051918252519081900360200190f35b3480156100ef57600080fd5b506100d16104c9565b34801561010457600080fd5b5061012b6004803603602081101561011b57600080fd5b50356001600160a01b03166104cf565b604080519115158252519081900360200190f35b34801561014b57600080fd5b506101546104f0565b005b34801561016257600080fd5b5061016b61058a565b604080516001600160a01b039092168252519081900360200190f35b34801561019357600080fd5b50610217600480360360608110156101aa57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156101d957600080fd5b8201836020820111156101eb57600080fd5b803590602001918460018302840111600160201b8311171561020c57600080fd5b509092509050610599565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561025c578181015183820152602001610244565b50505050905090810190601f1680156102895780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156102a457600080fd5b5061016b6106a2565b3480156102b957600080fd5b5061012b600480360360208110156102d057600080fd5b50356001600160a01b03166106b1565b3480156102ec57600080fd5b506101546004803603604081101561030357600080fd5b506001600160a01b03813516906020013515156106d3565b34801561032757600080fd5b506100d16004803603602081101561033e57600080fd5b503561090f565b34801561035157600080fd5b506101546004803603604081101561036857600080fd5b506001600160a01b038135169060200135151561092d565b34801561038c57600080fd5b50610154600480360360208110156103a357600080fd5b50356001600160a01b0316610b68565b3360009081526001602081905260408220015460ff1661040b576040805162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b604482015290519081900360640190fd5b600654600061041f86864342863a8a610c58565b905060008215610447576006600184038154811061043957fe5b906000526020600020015490505b60066104538284610cc9565b8154600181018355600092835260209283902001556040805133815260ff8a16928101929092526001600160a01b038816828201526060820187905251829185917f23be8e12e420b5da9fb98d8102572f640fb3c11a0085060472dfc0ed194b3cf79181900360800190a3509095945050505050565b60065490565b6001600160a01b031660009081526002602052604090206001015460ff1690565b6104f8610cf5565b6001600160a01b031661050961058a565b6001600160a01b031614610552576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020610d40833981519152908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b3360009081526002602052604081206001015460609060ff166105f5576040805162461bcd60e51b815260206004820152600f60248201526e09c9ea8be8ca49e9abe9eaaa8849eb608b1b604482015290519081900360640190fd5b600580546001600160a01b0319811633179091556040516001600160a01b0391821691881690879087908790808383808284376040519201945060009350909150508083038185875af1925050503d806000811461066f576040519150601f19603f3d011682016040523d82523d6000602084013e610674565b606091505b50600580546001600160a01b0319166001600160a01b03949094169390931790925597909650945050505050565b6005546001600160a01b031681565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b6106db610cf5565b6001600160a01b03166106ec61058a565b6001600160a01b031614610735576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b6001600160a01b0382166000908152600260205260409020600181015460ff1680801561075f5750825b80610771575080158015610771575082155b1561077d57505061090b565b821561080c57604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319169091179055610908565b60048054600019810190811061081e57fe5b6000918252602090912001548254600480546001600160a01b0390931692909190811061084757fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460026000600485600001548154811061088f57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205560048054806108bf57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff191690555b50505b5050565b6006818154811061091c57fe5b600091825260209091200154905081565b610935610cf5565b6001600160a01b031661094661058a565b6001600160a01b03161461098f576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b6001600160a01b03821660009081526001602081905260409091209081015460ff168080156109bb5750825b806109cd5750801580156109cd575082155b156109d957505061090b565b8215610a6757604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055610908565b600380546000198101908110610a7957fe5b6000918252602090912001548254600380546001600160a01b03909316929091908110610aa257fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600160006003856000015481548110610aea57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556003805480610b1a57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff1916905550505050565b610b70610cf5565b6001600160a01b0316610b8161058a565b6001600160a01b031614610bca576040805162461bcd60e51b81526020600482018190526024820152600080516020610d20833981519152604482015290519081900360640190fd5b6001600160a01b038116610c0f5760405162461bcd60e51b8152600401808060200182810382526026815260200180610cfa6026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020610d4083398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805160f89890981b6001600160f81b0319166020808a019190915260609790971b6001600160601b0319166021890152603588019590955260558701939093526075860191909152609585015260b5808501919091528151808503909101815260d59093019052815191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212208b6882e0b90e769a155451f0978d5937ce19dc35c6925c62ea997bdaa630206564736f6c634300060c0033608060405234801561001057600080fd5b50604051610afb380380610afb8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610a96806100656000396000f3fe6080604052600436106100765760003560e01c80631fe927cf1461007b5780635075788b146100f85780635e9167581461019b57806367ef3ab8146102255780636f5dfdca146102b45780638a631aa614610342578063ad9d4ba3146103de578063b75436bb14610404578063e78cea921461047f575b600080fd5b34801561008757600080fd5b506100f66004803603602081101561009e57600080fd5b810190602081018135600160201b8111156100b857600080fd5b8201836020820111156100ca57600080fd5b803590602001918460018302840111600160201b831117156100eb57600080fd5b5090925090506104b0565b005b34801561010457600080fd5b506100f6600480360360c081101561011b57600080fd5b8135916020810135916040820135916001600160a01b03606082013516916080820135919081019060c0810160a0820135600160201b81111561015d57600080fd5b82018360208201111561016f57600080fd5b803590602001918460018302840111600160201b8311171561019057600080fd5b50909250905061054f565b6100f6600480360360808110156101b157600080fd5b8135916020810135916001600160a01b036040830135169190810190608081016060820135600160201b8111156101e757600080fd5b8201836020820111156101f957600080fd5b803590602001918460018302840111600160201b8311171561021a57600080fd5b5090925090506105d2565b6100f6600480360360a081101561023b57600080fd5b8135916020810135916040820135916001600160a01b036060820135169181019060a081016080820135600160201b81111561027657600080fd5b82018360208201111561028857600080fd5b803590602001918460018302840111600160201b831117156102a957600080fd5b50909250905061064b565b3480156102c057600080fd5b506100f6600480360360808110156102d757600080fd5b81359160208101359160408201359190810190608081016060820135600160201b81111561030457600080fd5b82018360208201111561031657600080fd5b803590602001918460018302840111600160201b8311171561033757600080fd5b5090925090506106cd565b34801561034e57600080fd5b506100f6600480360360a081101561036557600080fd5b8135916020810135916001600160a01b036040830135169160608101359181019060a081016080820135600160201b8111156103a057600080fd5b8201836020820111156103b257600080fd5b803590602001918460018302840111600160201b831117156103d357600080fd5b509092509050610770565b6100f6600480360360208110156103f457600080fd5b50356001600160a01b03166107e2565b34801561041057600080fd5b506100f66004803603602081101561042757600080fd5b810190602081018135600160201b81111561044157600080fd5b82018360208201111561045357600080fd5b803590602001918460018302840111600160201b8311171561047457600080fd5b50909250905061083a565b34801561048b57600080fd5b506104946108ba565b604080516001600160a01b039092168252519081900360200190f35b3332146104f2576040805162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b604482015290519081900360640190fd5b600061051c600333858560405180838380828437604051920182900390912093506108c992505050565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a2505050565b6105c960033360008a8a8a8a60601b60601c6001600160a01b03168a8a8a604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610960565b50505050505050565b610644600733600188888860601b60601c6001600160a01b0316348989604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610960565b5050505050565b6106c560073360008989898960601b60601c6001600160a01b0316348a8a604051602001808960ff1660f81b81526001018881526020018781526020018681526020018581526020018481526020018383808284378083019250505098505050505050505050604051602081830303815290604052610960565b505050505050565b6106d633610a04565b610724576040805162461bcd60e51b815260206004820152601a6024820152791b5d5cdd0818994818d85b1b195908189e4818dbdb9d1c9858dd60321b604482015290519081900360640190fd5b6106446005338787878787604051602001808681526020018581526020018481526020018383808284378083019250505095505050505050604051602081830303815290604052610960565b6106c5600333600189898960601b60601c6001600160a01b0316898989604051602001808860ff1660f81b815260010187815260200186815260200185815260200184815260200183838082843780830192505050975050505050505050604051602081830303815290604052610960565b60408051600160f81b602082015260006021820181905260418201526001600160a01b0383166061820152346081808301919091528251808303909101815260a1909101909152610837906007908390610960565b50565b6000610864600333858560405180838380828437604051920182900390912093506108c992505050565b905080600080516020610a41833981519152848460405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2505050565b6000546001600160a01b031681565b60008054604080516302bbfad160e01b815260ff871660048201526001600160a01b03868116602483015260448201869052915191909216916302bbfad191349160648082019260209290919082900301818588803b15801561092b57600080fd5b505af115801561093f573d6000803e3d6000fd5b50505050506040513d602081101561095657600080fd5b5051949350505050565b6000610974848484805190602001206108c9565b905080600080516020610a41833981519152836040518080602001828103825283818151815260200191508051906020019080838360005b838110156109c45781810151838201526020016109ac565b50505050905090810190601f1680156109f15780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610a3857508115155b94935050505056feff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60ba2646970667358221220c69f0f2def64d34c144be3e1b5bf9c8d8644a74e6c602ebc3e7570d6b1afa55a64736f6c634300060c0033608060405234801561001057600080fd5b506040516109013803806109018339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b039384166001600160a01b031991821617909155600180549390921692169190911790556108878061007a6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806316b9109b1461006757806330a826b41461008657806364126c7c146100a35780638b8ca199146100cf578063b0f2af2914610107578063f03c04a5146101a6575b600080fd5b6100846004803603602081101561007d57600080fd5b50356101d2565b005b6100846004803603602081101561009c57600080fd5b503561025e565b610084600480360360408110156100b957600080fd5b50803590602001356001600160a01b03166102e5565b610084600480360360808110156100e557600080fd5b50803590602081013590604081013590606001356001600160a01b0316610518565b610084600480360360e081101561011d57600080fd5b8135916020810135916040820135916060810135916001600160a01b03608083013581169260a08101359091169181019060e0810160c0820135600160201b81111561016857600080fd5b82018360208201111561017a57600080fd5b803590602001918460018302840111600160201b8311171561019b57600080fd5b5090925090506105c2565b610084600480360360408110156101bc57600080fd5b506001600160a01b03813516906020013561071f565b6001546001600160a01b0316331461021f576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f81b60208083019190915260218083018590528351808403909101815260419092019092528051910120610259816107c4565b505050565b6001546001600160a01b031633146102ab576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600160f91b60208083019190915260218083018590528351808403909101815260419092019092528051910120610259816107c4565b6001546001600160a01b03163314610332576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60015460408051634f0f4aa960e01b81526004810185905290516001600160a01b03909216916000918391634f0f4aa991602480820192602092909190829003018186803b15801561038357600080fd5b505afa158015610397573d6000803e3d6000fd5b505050506040513d60208110156103ad57600080fd5b5051604080516348b4573960e11b81526001600160a01b038681166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b1580156103fc57600080fd5b505afa158015610410573d6000803e3d6000fd5b505050506040513d602081101561042657600080fd5b5051610466576040805162461bcd60e51b815260206004820152600a6024820152691393d517d4d51052d15160b21b604482015290519081900360640190fd5b816001600160a01b0316632b2af0ab856040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156104aa57600080fd5b505afa1580156104be573d6000803e3d6000fd5b505060408051600160fa1b602080830191909152602182018990526001600160a01b038816604180840191909152835180840390910181526061909201909252805191012091506105109050816107c4565b505050505050565b6001546001600160a01b03163314610565576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6040805160006020808301919091526021820187905260418201869052436061830152608182018590526001600160a01b03841660a1808401919091528351808403909101815260c19092019092528051910120610510816107c4565b6001546001600160a01b0316331461060f576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b6000888888888860601b60601c6001600160a01b03168860601b60601c6001600160a01b03168888604051602001808981526020018881526020018781526020018681526020018581526020018481526020018383808284376040805191909301818103601f190182528084528151602092830120600080546302bbfad160e01b8452600480850152336024850152604484018390529551919f506001600160a01b039095169d506302bbfad19c5060648083019c50929a50985088900301955086945090925089915050803b1580156106e857600080fd5b505af11580156106fc573d6000803e3d6000fd5b505050506040513d602081101561071257600080fd5b5050505050505050505050565b6001546001600160a01b0316331461076c576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b60408051600360f81b6020808301919091526001600160a01b0385166021830152604182018490524360618084019190915283518084039091018152608190920190925280519101206107be816107c4565b50505050565b60008054604080516302bbfad160e01b8152600860048201523360248201526044810185905290516001600160a01b03909216916302bbfad19160648082019260209290919082900301818787803b15801561081f57600080fd5b505af1158015610833573d6000803e3d6000fd5b505050506040513d602081101561084957600080fd5b50519291505056fea26469706673582212204078352c3f13578080f26a6077875d561d2a858b59bdee7e9acba04e442372a564736f6c634300060c0033608060405234801561001057600080fd5b5060405161168f38038061168f8339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b038085166001600160a01b0319928316179092556001805492841692909116919091179055604051610078906100bd565b604051809103906000f080158015610094573d6000803e3d6000fd5b50600280546001600160a01b0319166001600160a01b0392909216919091179055506100ca9050565b61059a806110f583390190565b61101c806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630c72684714610067578063465477901461012757806372fea01f1461014157806380648b0214610238578063b0f3053714610127578063d933556a1461025c575b600080fd5b6101256004803603604081101561007d57600080fd5b810190602081018135600160201b81111561009757600080fd5b8201836020820111156100a957600080fd5b803590602001918460018302840111600160201b831117156100ca57600080fd5b919390929091602081019035600160201b8111156100e757600080fd5b8201836020820111156100f957600080fd5b803590602001918460208302840111600160201b8311171561011a57600080fd5b5090925090506102e8565b005b61012f6103cf565b60408051918252519081900360200190f35b610125600480360361012081101561015857600080fd5b81359190810190604081016020820135600160201b81111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460208302840111600160201b831117156101ac57600080fd5b919390928235926001600160a01b03602082013581169360408301359091169260608301359260808101359260a082013592909160e081019060c00135600160201b8111156101fa57600080fd5b82018360208201111561020c57600080fd5b803590602001918460018302840111600160201b8311171561022d57600080fd5b5090925090506103de565b6102406103cf565b604080516001600160a01b039092168252519081900360200190f35b610125600480360360a081101561027257600080fd5b81359190810190604081016020820135600160201b81111561029357600080fd5b8201836020820111156102a557600080fd5b803590602001918460208302840111600160201b831117156102c657600080fd5b91935091508035906001600160a01b0360208201351690604001351515610597565b6000546001600160a01b03163314610335576040805162461bcd60e51b815260206004820152600b60248201526a04f4e4c595f524f4c4c55560ac1b604482015290519081900360640190fd5b806000805b828110156103c6576103a487838888888681811061035457fe5b9050602002013586019261036a93929190610fbe565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061065e92505050565b8484828181106103b057fe5b602002919091013592909201915060010161033a565b50505050505050565b6004546001600160a01b031690565b6000808860601b60601c6001600160a01b03168860601b60601c6001600160a01b031688888888886040516020018089815260200188815260200187815260200186815260200185815260200184815260200183838082843780830192505050985050505050505050506040516020818303038152906040528051906020012090506104a18c8c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508691506107db9050565b60048054600580546001600160a01b038c81166001600160a01b03198516179094556001600160801b03898116600160801b9081028c83166001600160801b0319851617831617909355604080516020601f8a0181900481028201810190925288815295909416948183169493909204169161053c918c918a91908a908a9081908401838280828437600092019190915250610a6192505050565b600480546001600160a01b039094166001600160a01b031990941693909317909255600580546001600160801b03938416600160801b029284166001600160801b031990911617909216179055505050505050505050505050565b6040805160056020808301919091526001600160a01b0385168284015283151560f81b606083015282516041818403018152606183018085528151918301919091206081928902808501840190955288825293610614938b938b928b92839201908490808284376000920191909152508992508691506107db9050565b604080518315156024808301919091528251808303909101815260449091019091526020810180516001600160e01b03166323d82c5d60e21b1790526103c6908490600090610c7e565b80516000908290829061066d57fe5b01602001516001600160f81b03191614156107d857600061068f826001610cc9565b9050600061069e836021610cc9565b905060006106ad846041610cc9565b6002549091506000906106c8906001600160a01b0316610d22565b60015460408051633422b05160e11b81526001600160a01b039283166004820152602481018690526044810187905290519293509083169163684560a29160648082019260009290919082900301818387803b15801561072757600080fd5b505af115801561073b573d6000803e3d6000fd5b5050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810180546001600160a01b0386166001600160a01b0319909116179055604080518281526020810187905280820188905290519193508792507fe5ccc8d7080a4904b2f4e42d91e8f06b13fe6cb2181ad1fe14644e856b44c131919081900360600190a250505050505b50565b61010083511115610824576040805162461bcd60e51b815260206004820152600e60248201526d50524f4f465f544f4f5f4c4f4e4760901b604482015290519081900360640190fd5b825160020a821061086f576040805162461bcd60e51b815260206004820152601060248201526f1410551217d393d517d352539253505360821b604482015290519081900360640190fd5b60006108a28484846040516020018082815260200191505060405160208183030381529060405280519060200120610e4e565b90506000600386815481106108b357fe5b6000918252602090912001546001600160a01b0316905080610908576040805162461bcd60e51b815260206004820152600960248201526809c9ebe9eaaa8849eb60bb1b604482015290519081900360640190fd5b8451604080516020808201889052818301939093528151808203830181526060820183528051908401206084820186905260a48083018290528351808403909101815260c490920190925291820180516001600160e01b0316630ad0379b60e01b1790529061097b908390600090610c7e565b816001600160a01b0316635780e4e76040518163ffffffff1660e01b815260040160206040518083038186803b1580156109b457600080fd5b505afa1580156109c8573d6000803e3d6000fd5b505050506040513d60208110156109de57600080fd5b50516103c6576040805160048152602481019091526020810180516001600160e01b031663083197ef60e41b179052610a1b908390600090610c7e565b600060038881548110610a2a57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050565b600154604051639e5d4c4960e01b81526001600160a01b03858116600483019081526024830186905260606044840181815286516064860152865160009692959490921693639e5d4c49938a938a938a93909160849091019060208501908083838e5b83811015610adc578181015183820152602001610ac4565b50505050905090810190601f168015610b095780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610b2a57600080fd5b505af1158015610b3e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610b6757600080fd5b815160208301805160405192949293830192919084600160201b821115610b8d57600080fd5b908301906020820185811115610ba257600080fd5b8251600160201b811182820188101715610bbb57600080fd5b82525081516020918201929091019080838360005b83811015610be8578181015183820152602001610bd0565b50505050905090810190601f168015610c155780820380516001836020036101000a031916815260200191505b506040525050509150915081610c7757805115610c355780518082602001fd5b6040805162461bcd60e51b81526020600482015260126024820152711094925111d157d0d0531317d1905253115160721b604482015290519081900360640190fd5b5050505050565b600480546001600160a01b031981169091556001600160a01b0316610ca4848484610a61565b600480546001600160a01b0319166001600160a01b0392909216919091179055505050565b60008160200183511015610d19576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000816001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b158015610d5d57600080fd5b505afa158015610d71573d6000803e3d6000fd5b505050506040513d6020811015610d8757600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b602082015290610e355760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610dfa578181015183820152602001610de2565b50505050905090810190601f168015610e275780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610e48826001600160a01b0316610f1c565b92915050565b8251600090610100811115610e6257600080fd5b8260005b82811015610f125760028606610ebf57868181518110610e8257fe5b6020026020010151826040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150610f04565b81878281518110610ecc57fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091505b600286049550600101610e66565b5095945050505050565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528260601b60148201526e5af43d82803e903d91602b57fd5bf360881b60288201526037816000f09150506001600160a01b038116610fb9576040805162461bcd60e51b8152602060048201526016602482015275115490cc4c4d8dce8818dc99585d194819985a5b195960521b604482015290519081900360640190fd5b919050565b60008085851115610fcd578182fd5b83861115610fd9578182fd5b505082019391909203915056fea26469706673582212207a6f9fa675d15467bc4ffc2121ab858ac426793d7dbbdee246af9719f33f7c3564736f6c634300060c0033608060405234801561001057600080fd5b506000805460ff1916600117905561056d8061002d6000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c80635780e4e71461007257806357d61c0b1461008c578063684560a2146100b15780636f791d29146100e357806383197ef0146100ff5780639db9af8114610107578063ebf0c71714610124575b600080fd5b61007a61012c565b60408051918252519081900360200190f35b6100af600480360360408110156100a257600080fd5b5080359060200135610132565b005b6100af600480360360608110156100c757600080fd5b506001600160a01b038135169060208101359060400135610205565b6100eb6102b5565b604080519115158252519081900360200190f35b6100af6102be565b6100eb6004803603602081101561011d57600080fd5b50356102d1565b61007a6102e6565b60025481565b61013a6102ec565b60008181526003602052604090205460ff161561018e576040805162461bcd60e51b815260206004820152600d60248201526c1053149150511657d4d4115395609a1b604482015290519081900360640190fd5b60015482146101cf576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b6000818152600360205260409020805460ff19166001179055600280546000190190819055610201576102013361047a565b5050565b60015415610249576040805162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015290519081900360640190fd5b81610286576040805162461bcd60e51b815260206004820152600860248201526710905117d493d3d560c21b604482015290519081900360640190fd5b600080546001600160a01b0390941661010002610100600160a81b031990941693909317909255600155600255565b60005460ff1690565b6102c66102ec565b6102cf3361047a565b565b60036020526000908152604090205460ff1681565b60015481565b60005461010090046001600160a01b0316331461033e576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f42524944474560a81b604482015290519081900360640190fd5b60006001600160a01b0316600060019054906101000a90046001600160a01b03166001600160a01b031663ab5d89436040518163ffffffff1660e01b815260040160206040518083038186803b15801561039757600080fd5b505afa1580156103ab573d6000803e3d6000fd5b505050506040513d60208110156103c157600080fd5b505160408051634032458160e11b815290516001600160a01b03909216916380648b0291600480820192602092909190829003018186803b15801561040557600080fd5b505afa158015610419573d6000803e3d6000fd5b505050506040513d602081101561042f57600080fd5b50516001600160a01b0316146102cf576040805162461bcd60e51b815260206004820152600b60248201526a4f4e4c595f53595354454d60a81b604482015290519081900360640190fd5b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff161561052a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104ef5781810151838201526020016104d7565b50505050905090810190601f16801561051c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50806001600160a01b0316fffea26469706673582212203b9da5963c6361e0f514db3f5dae9eca4924dbd009dfdf6859b88da66ef9b44664736f6c634300060c00334f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a264697066735822122001d76a8044e6cd91fa16d46e06a8f3a05ad7bbb44b5af4c4554d19566baf7c7164736f6c634300060c0033"

// DeployRollupCreator deploys a new Ethereum contract, binding an instance of RollupCreator to it.
func DeployRollupCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// RollupCreator is an auto generated Go binding around an Ethereum contract.
type RollupCreator struct {
	RollupCreatorCaller     // Read-only binding to the contract
	RollupCreatorTransactor // Write-only binding to the contract
	RollupCreatorFilterer   // Log filterer for contract events
}

// RollupCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorSession struct {
	Contract     *RollupCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorCallerSession struct {
	Contract *RollupCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RollupCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorTransactorSession struct {
	Contract     *RollupCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RollupCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorRaw struct {
	Contract *RollupCreator // Generic contract binding to access the raw methods on
}

// RollupCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorCallerRaw struct {
	Contract *RollupCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorTransactorRaw struct {
	Contract *RollupCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreator creates a new instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreator(address common.Address, backend bind.ContractBackend) (*RollupCreator, error) {
	contract, err := bindRollupCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// NewRollupCreatorCaller creates a new read-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorCaller, error) {
	contract, err := bindRollupCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorCaller{contract: contract}, nil
}

// NewRollupCreatorTransactor creates a new write-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorTransactor, error) {
	contract, err := bindRollupCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTransactor{contract: contract}, nil
}

// NewRollupCreatorFilterer creates a new log filterer instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorFilterer, error) {
	contract, err := bindRollupCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorFilterer{contract: contract}, nil
}

// bindRollupCreator binds a generic wrapper to an already deployed contract.
func bindRollupCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.RollupCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorSession) Owner() (common.Address, error) {
	return _RollupCreator.Contract.Owner(&_RollupCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) Owner() (common.Address, error) {
	return _RollupCreator.Contract.Owner(&_RollupCreator.CallOpts)
}

// CreateRollup is a paid mutator transaction binding the contract method 0x84b99970.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactor) CreateRollup(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "createRollup", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0x84b99970.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorSession) CreateRollup(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0x84b99970.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactorSession) CreateRollup(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _extraConfig)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreator.Contract.RenounceOwnership(&_RollupCreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreator.Contract.RenounceOwnership(&_RollupCreator.TransactOpts)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xd9220824.
//
// Solidity: function setTemplates(address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreator *RollupCreatorTransactor) SetTemplates(opts *bind.TransactOpts, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "setTemplates", _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xd9220824.
//
// Solidity: function setTemplates(address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreator *RollupCreatorSession) SetTemplates(_rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.SetTemplates(&_RollupCreator.TransactOpts, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xd9220824.
//
// Solidity: function setTemplates(address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreator *RollupCreatorTransactorSession) SetTemplates(_rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.SetTemplates(&_RollupCreator.TransactOpts, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.TransferOwnership(&_RollupCreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.TransferOwnership(&_RollupCreator.TransactOpts, newOwner)
}

// RollupCreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupCreator contract.
type RollupCreatorOwnershipTransferredIterator struct {
	Event *RollupCreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupCreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorOwnershipTransferred)
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
		it.Event = new(RollupCreatorOwnershipTransferred)
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
func (it *RollupCreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorOwnershipTransferred represents a OwnershipTransferred event raised by the RollupCreator contract.
type RollupCreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupCreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorOwnershipTransferredIterator{contract: _RollupCreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupCreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorOwnershipTransferred)
				if err := _RollupCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RollupCreator *RollupCreatorFilterer) ParseOwnershipTransferred(log types.Log) (*RollupCreatorOwnershipTransferred, error) {
	event := new(RollupCreatorOwnershipTransferred)
	if err := _RollupCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreator contract.
type RollupCreatorRollupCreatedIterator struct {
	Event *RollupCreatorRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorRollupCreated)
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
		it.Event = new(RollupCreatorRollupCreated)
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
func (it *RollupCreatorRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorRollupCreated represents a RollupCreated event raised by the RollupCreator contract.
type RollupCreatorRollupCreated struct {
	RollupAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupCreatorRollupCreatedIterator, error) {

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorRollupCreatedIterator{contract: _RollupCreator.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorRollupCreated) (event.Subscription, error) {

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorRollupCreated)
				if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x84c162f1396badc29f9c932c79d7495db699b615e2c0da163ae26bd5dbe71d7c.
//
// Solidity: event RollupCreated(address rollupAddress)
func (_RollupCreator *RollupCreatorFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorRollupCreated, error) {
	event := new(RollupCreatorRollupCreated)
	if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
