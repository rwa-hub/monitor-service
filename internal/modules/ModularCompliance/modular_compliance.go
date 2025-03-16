// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modularcompliance

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

// ModularcomplianceMetaData contains all meta data concerning the Modularcompliance contract.
var ModularcomplianceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addModule\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bindToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callModuleFunction\",\"inputs\":[{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canTransfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"created\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"destroyed\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getModules\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isModuleBound\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeModule\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferred\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbindToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleAdded\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleInteraction\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleRemoved\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenBound\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenUnbound\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ModularcomplianceABI is the input ABI used to generate the binding from.
// Deprecated: Use ModularcomplianceMetaData.ABI instead.
var ModularcomplianceABI = ModularcomplianceMetaData.ABI

// Modularcompliance is an auto generated Go binding around an Ethereum contract.
type Modularcompliance struct {
	ModularcomplianceCaller     // Read-only binding to the contract
	ModularcomplianceTransactor // Write-only binding to the contract
	ModularcomplianceFilterer   // Log filterer for contract events
}

// ModularcomplianceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModularcomplianceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularcomplianceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModularcomplianceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularcomplianceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModularcomplianceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModularcomplianceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModularcomplianceSession struct {
	Contract     *Modularcompliance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ModularcomplianceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModularcomplianceCallerSession struct {
	Contract *ModularcomplianceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ModularcomplianceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModularcomplianceTransactorSession struct {
	Contract     *ModularcomplianceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ModularcomplianceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModularcomplianceRaw struct {
	Contract *Modularcompliance // Generic contract binding to access the raw methods on
}

// ModularcomplianceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModularcomplianceCallerRaw struct {
	Contract *ModularcomplianceCaller // Generic read-only contract binding to access the raw methods on
}

// ModularcomplianceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModularcomplianceTransactorRaw struct {
	Contract *ModularcomplianceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModularcompliance creates a new instance of Modularcompliance, bound to a specific deployed contract.
func NewModularcompliance(address common.Address, backend bind.ContractBackend) (*Modularcompliance, error) {
	contract, err := bindModularcompliance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Modularcompliance{ModularcomplianceCaller: ModularcomplianceCaller{contract: contract}, ModularcomplianceTransactor: ModularcomplianceTransactor{contract: contract}, ModularcomplianceFilterer: ModularcomplianceFilterer{contract: contract}}, nil
}

// NewModularcomplianceCaller creates a new read-only instance of Modularcompliance, bound to a specific deployed contract.
func NewModularcomplianceCaller(address common.Address, caller bind.ContractCaller) (*ModularcomplianceCaller, error) {
	contract, err := bindModularcompliance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceCaller{contract: contract}, nil
}

// NewModularcomplianceTransactor creates a new write-only instance of Modularcompliance, bound to a specific deployed contract.
func NewModularcomplianceTransactor(address common.Address, transactor bind.ContractTransactor) (*ModularcomplianceTransactor, error) {
	contract, err := bindModularcompliance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceTransactor{contract: contract}, nil
}

// NewModularcomplianceFilterer creates a new log filterer instance of Modularcompliance, bound to a specific deployed contract.
func NewModularcomplianceFilterer(address common.Address, filterer bind.ContractFilterer) (*ModularcomplianceFilterer, error) {
	contract, err := bindModularcompliance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceFilterer{contract: contract}, nil
}

// bindModularcompliance binds a generic wrapper to an already deployed contract.
func bindModularcompliance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModularcomplianceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Modularcompliance *ModularcomplianceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Modularcompliance.Contract.ModularcomplianceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Modularcompliance *ModularcomplianceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modularcompliance.Contract.ModularcomplianceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Modularcompliance *ModularcomplianceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Modularcompliance.Contract.ModularcomplianceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Modularcompliance *ModularcomplianceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Modularcompliance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Modularcompliance *ModularcomplianceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modularcompliance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Modularcompliance *ModularcomplianceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Modularcompliance.Contract.contract.Transact(opts, method, params...)
}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address _from, address _to, uint256 _value) view returns(bool)
func (_Modularcompliance *ModularcomplianceCaller) CanTransfer(opts *bind.CallOpts, _from common.Address, _to common.Address, _value *big.Int) (bool, error) {
	var out []interface{}
	err := _Modularcompliance.contract.Call(opts, &out, "canTransfer", _from, _to, _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address _from, address _to, uint256 _value) view returns(bool)
func (_Modularcompliance *ModularcomplianceSession) CanTransfer(_from common.Address, _to common.Address, _value *big.Int) (bool, error) {
	return _Modularcompliance.Contract.CanTransfer(&_Modularcompliance.CallOpts, _from, _to, _value)
}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address _from, address _to, uint256 _value) view returns(bool)
func (_Modularcompliance *ModularcomplianceCallerSession) CanTransfer(_from common.Address, _to common.Address, _value *big.Int) (bool, error) {
	return _Modularcompliance.Contract.CanTransfer(&_Modularcompliance.CallOpts, _from, _to, _value)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_Modularcompliance *ModularcomplianceCaller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Modularcompliance.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_Modularcompliance *ModularcomplianceSession) GetModules() ([]common.Address, error) {
	return _Modularcompliance.Contract.GetModules(&_Modularcompliance.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_Modularcompliance *ModularcomplianceCallerSession) GetModules() ([]common.Address, error) {
	return _Modularcompliance.Contract.GetModules(&_Modularcompliance.CallOpts)
}

// GetTokenBound is a free data retrieval call binding the contract method 0x6a3edf28.
//
// Solidity: function getTokenBound() view returns(address)
func (_Modularcompliance *ModularcomplianceCaller) GetTokenBound(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Modularcompliance.contract.Call(opts, &out, "getTokenBound")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenBound is a free data retrieval call binding the contract method 0x6a3edf28.
//
// Solidity: function getTokenBound() view returns(address)
func (_Modularcompliance *ModularcomplianceSession) GetTokenBound() (common.Address, error) {
	return _Modularcompliance.Contract.GetTokenBound(&_Modularcompliance.CallOpts)
}

// GetTokenBound is a free data retrieval call binding the contract method 0x6a3edf28.
//
// Solidity: function getTokenBound() view returns(address)
func (_Modularcompliance *ModularcomplianceCallerSession) GetTokenBound() (common.Address, error) {
	return _Modularcompliance.Contract.GetTokenBound(&_Modularcompliance.CallOpts)
}

// IsModuleBound is a free data retrieval call binding the contract method 0xa446d49f.
//
// Solidity: function isModuleBound(address _module) view returns(bool)
func (_Modularcompliance *ModularcomplianceCaller) IsModuleBound(opts *bind.CallOpts, _module common.Address) (bool, error) {
	var out []interface{}
	err := _Modularcompliance.contract.Call(opts, &out, "isModuleBound", _module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleBound is a free data retrieval call binding the contract method 0xa446d49f.
//
// Solidity: function isModuleBound(address _module) view returns(bool)
func (_Modularcompliance *ModularcomplianceSession) IsModuleBound(_module common.Address) (bool, error) {
	return _Modularcompliance.Contract.IsModuleBound(&_Modularcompliance.CallOpts, _module)
}

// IsModuleBound is a free data retrieval call binding the contract method 0xa446d49f.
//
// Solidity: function isModuleBound(address _module) view returns(bool)
func (_Modularcompliance *ModularcomplianceCallerSession) IsModuleBound(_module common.Address) (bool, error) {
	return _Modularcompliance.Contract.IsModuleBound(&_Modularcompliance.CallOpts, _module)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Modularcompliance *ModularcomplianceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Modularcompliance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Modularcompliance *ModularcomplianceSession) Owner() (common.Address, error) {
	return _Modularcompliance.Contract.Owner(&_Modularcompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Modularcompliance *ModularcomplianceCallerSession) Owner() (common.Address, error) {
	return _Modularcompliance.Contract.Owner(&_Modularcompliance.CallOpts)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address _module) returns()
func (_Modularcompliance *ModularcomplianceTransactor) AddModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "addModule", _module)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address _module) returns()
func (_Modularcompliance *ModularcomplianceSession) AddModule(_module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.AddModule(&_Modularcompliance.TransactOpts, _module)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address _module) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) AddModule(_module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.AddModule(&_Modularcompliance.TransactOpts, _module)
}

// BindToken is a paid mutator transaction binding the contract method 0x3ff5aa02.
//
// Solidity: function bindToken(address _token) returns()
func (_Modularcompliance *ModularcomplianceTransactor) BindToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "bindToken", _token)
}

// BindToken is a paid mutator transaction binding the contract method 0x3ff5aa02.
//
// Solidity: function bindToken(address _token) returns()
func (_Modularcompliance *ModularcomplianceSession) BindToken(_token common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.BindToken(&_Modularcompliance.TransactOpts, _token)
}

// BindToken is a paid mutator transaction binding the contract method 0x3ff5aa02.
//
// Solidity: function bindToken(address _token) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) BindToken(_token common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.BindToken(&_Modularcompliance.TransactOpts, _token)
}

// CallModuleFunction is a paid mutator transaction binding the contract method 0xefb22d33.
//
// Solidity: function callModuleFunction(bytes callData, address _module) returns()
func (_Modularcompliance *ModularcomplianceTransactor) CallModuleFunction(opts *bind.TransactOpts, callData []byte, _module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "callModuleFunction", callData, _module)
}

// CallModuleFunction is a paid mutator transaction binding the contract method 0xefb22d33.
//
// Solidity: function callModuleFunction(bytes callData, address _module) returns()
func (_Modularcompliance *ModularcomplianceSession) CallModuleFunction(callData []byte, _module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.CallModuleFunction(&_Modularcompliance.TransactOpts, callData, _module)
}

// CallModuleFunction is a paid mutator transaction binding the contract method 0xefb22d33.
//
// Solidity: function callModuleFunction(bytes callData, address _module) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) CallModuleFunction(callData []byte, _module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.CallModuleFunction(&_Modularcompliance.TransactOpts, callData, _module)
}

// Created is a paid mutator transaction binding the contract method 0x5f8dead3.
//
// Solidity: function created(address _to, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceTransactor) Created(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "created", _to, _value)
}

// Created is a paid mutator transaction binding the contract method 0x5f8dead3.
//
// Solidity: function created(address _to, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceSession) Created(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.Contract.Created(&_Modularcompliance.TransactOpts, _to, _value)
}

// Created is a paid mutator transaction binding the contract method 0x5f8dead3.
//
// Solidity: function created(address _to, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) Created(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.Contract.Created(&_Modularcompliance.TransactOpts, _to, _value)
}

// Destroyed is a paid mutator transaction binding the contract method 0x8d2ea772.
//
// Solidity: function destroyed(address _from, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceTransactor) Destroyed(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "destroyed", _from, _value)
}

// Destroyed is a paid mutator transaction binding the contract method 0x8d2ea772.
//
// Solidity: function destroyed(address _from, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceSession) Destroyed(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.Contract.Destroyed(&_Modularcompliance.TransactOpts, _from, _value)
}

// Destroyed is a paid mutator transaction binding the contract method 0x8d2ea772.
//
// Solidity: function destroyed(address _from, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) Destroyed(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.Contract.Destroyed(&_Modularcompliance.TransactOpts, _from, _value)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Modularcompliance *ModularcomplianceTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Modularcompliance *ModularcomplianceSession) Init() (*types.Transaction, error) {
	return _Modularcompliance.Contract.Init(&_Modularcompliance.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) Init() (*types.Transaction, error) {
	return _Modularcompliance.Contract.Init(&_Modularcompliance.TransactOpts)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address _module) returns()
func (_Modularcompliance *ModularcomplianceTransactor) RemoveModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "removeModule", _module)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address _module) returns()
func (_Modularcompliance *ModularcomplianceSession) RemoveModule(_module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.RemoveModule(&_Modularcompliance.TransactOpts, _module)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address _module) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) RemoveModule(_module common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.RemoveModule(&_Modularcompliance.TransactOpts, _module)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Modularcompliance *ModularcomplianceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Modularcompliance *ModularcomplianceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Modularcompliance.Contract.RenounceOwnership(&_Modularcompliance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Modularcompliance.Contract.RenounceOwnership(&_Modularcompliance.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Modularcompliance *ModularcomplianceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Modularcompliance *ModularcomplianceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.TransferOwnership(&_Modularcompliance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.TransferOwnership(&_Modularcompliance.TransactOpts, newOwner)
}

// Transferred is a paid mutator transaction binding the contract method 0x8baf29b4.
//
// Solidity: function transferred(address _from, address _to, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceTransactor) Transferred(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "transferred", _from, _to, _value)
}

// Transferred is a paid mutator transaction binding the contract method 0x8baf29b4.
//
// Solidity: function transferred(address _from, address _to, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceSession) Transferred(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.Contract.Transferred(&_Modularcompliance.TransactOpts, _from, _to, _value)
}

// Transferred is a paid mutator transaction binding the contract method 0x8baf29b4.
//
// Solidity: function transferred(address _from, address _to, uint256 _value) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) Transferred(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Modularcompliance.Contract.Transferred(&_Modularcompliance.TransactOpts, _from, _to, _value)
}

// UnbindToken is a paid mutator transaction binding the contract method 0x40db3b50.
//
// Solidity: function unbindToken(address _token) returns()
func (_Modularcompliance *ModularcomplianceTransactor) UnbindToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Modularcompliance.contract.Transact(opts, "unbindToken", _token)
}

// UnbindToken is a paid mutator transaction binding the contract method 0x40db3b50.
//
// Solidity: function unbindToken(address _token) returns()
func (_Modularcompliance *ModularcomplianceSession) UnbindToken(_token common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.UnbindToken(&_Modularcompliance.TransactOpts, _token)
}

// UnbindToken is a paid mutator transaction binding the contract method 0x40db3b50.
//
// Solidity: function unbindToken(address _token) returns()
func (_Modularcompliance *ModularcomplianceTransactorSession) UnbindToken(_token common.Address) (*types.Transaction, error) {
	return _Modularcompliance.Contract.UnbindToken(&_Modularcompliance.TransactOpts, _token)
}

// ModularcomplianceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Modularcompliance contract.
type ModularcomplianceInitializedIterator struct {
	Event *ModularcomplianceInitialized // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceInitialized)
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
		it.Event = new(ModularcomplianceInitialized)
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
func (it *ModularcomplianceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceInitialized represents a Initialized event raised by the Modularcompliance contract.
type ModularcomplianceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Modularcompliance *ModularcomplianceFilterer) FilterInitialized(opts *bind.FilterOpts) (*ModularcomplianceInitializedIterator, error) {

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceInitializedIterator{contract: _Modularcompliance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Modularcompliance *ModularcomplianceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ModularcomplianceInitialized) (event.Subscription, error) {

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceInitialized)
				if err := _Modularcompliance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Modularcompliance *ModularcomplianceFilterer) ParseInitialized(log types.Log) (*ModularcomplianceInitialized, error) {
	event := new(ModularcomplianceInitialized)
	if err := _Modularcompliance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModularcomplianceModuleAddedIterator is returned from FilterModuleAdded and is used to iterate over the raw logs and unpacked data for ModuleAdded events raised by the Modularcompliance contract.
type ModularcomplianceModuleAddedIterator struct {
	Event *ModularcomplianceModuleAdded // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceModuleAdded)
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
		it.Event = new(ModularcomplianceModuleAdded)
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
func (it *ModularcomplianceModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceModuleAdded represents a ModuleAdded event raised by the Modularcompliance contract.
type ModularcomplianceModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuleAdded is a free log retrieval operation binding the contract event 0xead6a006345da1073a106d5f32372d2d2204f46cb0b4bca8f5ebafcbbed12b8a.
//
// Solidity: event ModuleAdded(address indexed _module)
func (_Modularcompliance *ModularcomplianceFilterer) FilterModuleAdded(opts *bind.FilterOpts, _module []common.Address) (*ModularcomplianceModuleAddedIterator, error) {

	var _moduleRule []interface{}
	for _, _moduleItem := range _module {
		_moduleRule = append(_moduleRule, _moduleItem)
	}

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "ModuleAdded", _moduleRule)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceModuleAddedIterator{contract: _Modularcompliance.contract, event: "ModuleAdded", logs: logs, sub: sub}, nil
}

// WatchModuleAdded is a free log subscription operation binding the contract event 0xead6a006345da1073a106d5f32372d2d2204f46cb0b4bca8f5ebafcbbed12b8a.
//
// Solidity: event ModuleAdded(address indexed _module)
func (_Modularcompliance *ModularcomplianceFilterer) WatchModuleAdded(opts *bind.WatchOpts, sink chan<- *ModularcomplianceModuleAdded, _module []common.Address) (event.Subscription, error) {

	var _moduleRule []interface{}
	for _, _moduleItem := range _module {
		_moduleRule = append(_moduleRule, _moduleItem)
	}

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "ModuleAdded", _moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceModuleAdded)
				if err := _Modularcompliance.contract.UnpackLog(event, "ModuleAdded", log); err != nil {
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

// ParseModuleAdded is a log parse operation binding the contract event 0xead6a006345da1073a106d5f32372d2d2204f46cb0b4bca8f5ebafcbbed12b8a.
//
// Solidity: event ModuleAdded(address indexed _module)
func (_Modularcompliance *ModularcomplianceFilterer) ParseModuleAdded(log types.Log) (*ModularcomplianceModuleAdded, error) {
	event := new(ModularcomplianceModuleAdded)
	if err := _Modularcompliance.contract.UnpackLog(event, "ModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModularcomplianceModuleInteractionIterator is returned from FilterModuleInteraction and is used to iterate over the raw logs and unpacked data for ModuleInteraction events raised by the Modularcompliance contract.
type ModularcomplianceModuleInteractionIterator struct {
	Event *ModularcomplianceModuleInteraction // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceModuleInteractionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceModuleInteraction)
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
		it.Event = new(ModularcomplianceModuleInteraction)
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
func (it *ModularcomplianceModuleInteractionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceModuleInteractionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceModuleInteraction represents a ModuleInteraction event raised by the Modularcompliance contract.
type ModularcomplianceModuleInteraction struct {
	Target   common.Address
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModuleInteraction is a free log retrieval operation binding the contract event 0x20d79de70adcc6e9353d8a9a5646b46dc352710d0a310b1ad1f67faeca7ef891.
//
// Solidity: event ModuleInteraction(address indexed target, bytes4 selector)
func (_Modularcompliance *ModularcomplianceFilterer) FilterModuleInteraction(opts *bind.FilterOpts, target []common.Address) (*ModularcomplianceModuleInteractionIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "ModuleInteraction", targetRule)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceModuleInteractionIterator{contract: _Modularcompliance.contract, event: "ModuleInteraction", logs: logs, sub: sub}, nil
}

// WatchModuleInteraction is a free log subscription operation binding the contract event 0x20d79de70adcc6e9353d8a9a5646b46dc352710d0a310b1ad1f67faeca7ef891.
//
// Solidity: event ModuleInteraction(address indexed target, bytes4 selector)
func (_Modularcompliance *ModularcomplianceFilterer) WatchModuleInteraction(opts *bind.WatchOpts, sink chan<- *ModularcomplianceModuleInteraction, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "ModuleInteraction", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceModuleInteraction)
				if err := _Modularcompliance.contract.UnpackLog(event, "ModuleInteraction", log); err != nil {
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

// ParseModuleInteraction is a log parse operation binding the contract event 0x20d79de70adcc6e9353d8a9a5646b46dc352710d0a310b1ad1f67faeca7ef891.
//
// Solidity: event ModuleInteraction(address indexed target, bytes4 selector)
func (_Modularcompliance *ModularcomplianceFilterer) ParseModuleInteraction(log types.Log) (*ModularcomplianceModuleInteraction, error) {
	event := new(ModularcomplianceModuleInteraction)
	if err := _Modularcompliance.contract.UnpackLog(event, "ModuleInteraction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModularcomplianceModuleRemovedIterator is returned from FilterModuleRemoved and is used to iterate over the raw logs and unpacked data for ModuleRemoved events raised by the Modularcompliance contract.
type ModularcomplianceModuleRemovedIterator struct {
	Event *ModularcomplianceModuleRemoved // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceModuleRemoved)
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
		it.Event = new(ModularcomplianceModuleRemoved)
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
func (it *ModularcomplianceModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceModuleRemoved represents a ModuleRemoved event raised by the Modularcompliance contract.
type ModularcomplianceModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuleRemoved is a free log retrieval operation binding the contract event 0x0a1ee69f55c33d8467c69ca59ce2007a737a88603d75392972520bf67cb513b8.
//
// Solidity: event ModuleRemoved(address indexed _module)
func (_Modularcompliance *ModularcomplianceFilterer) FilterModuleRemoved(opts *bind.FilterOpts, _module []common.Address) (*ModularcomplianceModuleRemovedIterator, error) {

	var _moduleRule []interface{}
	for _, _moduleItem := range _module {
		_moduleRule = append(_moduleRule, _moduleItem)
	}

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "ModuleRemoved", _moduleRule)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceModuleRemovedIterator{contract: _Modularcompliance.contract, event: "ModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchModuleRemoved is a free log subscription operation binding the contract event 0x0a1ee69f55c33d8467c69ca59ce2007a737a88603d75392972520bf67cb513b8.
//
// Solidity: event ModuleRemoved(address indexed _module)
func (_Modularcompliance *ModularcomplianceFilterer) WatchModuleRemoved(opts *bind.WatchOpts, sink chan<- *ModularcomplianceModuleRemoved, _module []common.Address) (event.Subscription, error) {

	var _moduleRule []interface{}
	for _, _moduleItem := range _module {
		_moduleRule = append(_moduleRule, _moduleItem)
	}

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "ModuleRemoved", _moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceModuleRemoved)
				if err := _Modularcompliance.contract.UnpackLog(event, "ModuleRemoved", log); err != nil {
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

// ParseModuleRemoved is a log parse operation binding the contract event 0x0a1ee69f55c33d8467c69ca59ce2007a737a88603d75392972520bf67cb513b8.
//
// Solidity: event ModuleRemoved(address indexed _module)
func (_Modularcompliance *ModularcomplianceFilterer) ParseModuleRemoved(log types.Log) (*ModularcomplianceModuleRemoved, error) {
	event := new(ModularcomplianceModuleRemoved)
	if err := _Modularcompliance.contract.UnpackLog(event, "ModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModularcomplianceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Modularcompliance contract.
type ModularcomplianceOwnershipTransferredIterator struct {
	Event *ModularcomplianceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceOwnershipTransferred)
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
		it.Event = new(ModularcomplianceOwnershipTransferred)
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
func (it *ModularcomplianceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceOwnershipTransferred represents a OwnershipTransferred event raised by the Modularcompliance contract.
type ModularcomplianceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Modularcompliance *ModularcomplianceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModularcomplianceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceOwnershipTransferredIterator{contract: _Modularcompliance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Modularcompliance *ModularcomplianceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModularcomplianceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceOwnershipTransferred)
				if err := _Modularcompliance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Modularcompliance *ModularcomplianceFilterer) ParseOwnershipTransferred(log types.Log) (*ModularcomplianceOwnershipTransferred, error) {
	event := new(ModularcomplianceOwnershipTransferred)
	if err := _Modularcompliance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModularcomplianceTokenBoundIterator is returned from FilterTokenBound and is used to iterate over the raw logs and unpacked data for TokenBound events raised by the Modularcompliance contract.
type ModularcomplianceTokenBoundIterator struct {
	Event *ModularcomplianceTokenBound // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceTokenBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceTokenBound)
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
		it.Event = new(ModularcomplianceTokenBound)
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
func (it *ModularcomplianceTokenBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceTokenBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceTokenBound represents a TokenBound event raised by the Modularcompliance contract.
type ModularcomplianceTokenBound struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenBound is a free log retrieval operation binding the contract event 0x2de35142b19ed5a07796cf30791959c592018f70b1d2d7c460eef8ffe713692b.
//
// Solidity: event TokenBound(address _token)
func (_Modularcompliance *ModularcomplianceFilterer) FilterTokenBound(opts *bind.FilterOpts) (*ModularcomplianceTokenBoundIterator, error) {

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "TokenBound")
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceTokenBoundIterator{contract: _Modularcompliance.contract, event: "TokenBound", logs: logs, sub: sub}, nil
}

// WatchTokenBound is a free log subscription operation binding the contract event 0x2de35142b19ed5a07796cf30791959c592018f70b1d2d7c460eef8ffe713692b.
//
// Solidity: event TokenBound(address _token)
func (_Modularcompliance *ModularcomplianceFilterer) WatchTokenBound(opts *bind.WatchOpts, sink chan<- *ModularcomplianceTokenBound) (event.Subscription, error) {

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "TokenBound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceTokenBound)
				if err := _Modularcompliance.contract.UnpackLog(event, "TokenBound", log); err != nil {
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

// ParseTokenBound is a log parse operation binding the contract event 0x2de35142b19ed5a07796cf30791959c592018f70b1d2d7c460eef8ffe713692b.
//
// Solidity: event TokenBound(address _token)
func (_Modularcompliance *ModularcomplianceFilterer) ParseTokenBound(log types.Log) (*ModularcomplianceTokenBound, error) {
	event := new(ModularcomplianceTokenBound)
	if err := _Modularcompliance.contract.UnpackLog(event, "TokenBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModularcomplianceTokenUnboundIterator is returned from FilterTokenUnbound and is used to iterate over the raw logs and unpacked data for TokenUnbound events raised by the Modularcompliance contract.
type ModularcomplianceTokenUnboundIterator struct {
	Event *ModularcomplianceTokenUnbound // Event containing the contract specifics and raw log

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
func (it *ModularcomplianceTokenUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModularcomplianceTokenUnbound)
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
		it.Event = new(ModularcomplianceTokenUnbound)
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
func (it *ModularcomplianceTokenUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModularcomplianceTokenUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModularcomplianceTokenUnbound represents a TokenUnbound event raised by the Modularcompliance contract.
type ModularcomplianceTokenUnbound struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUnbound is a free log retrieval operation binding the contract event 0x28a4ca7134a3b3f9aff286e79ad3daadb4a06d1b43d037a3a98bdc074edd9b7a.
//
// Solidity: event TokenUnbound(address _token)
func (_Modularcompliance *ModularcomplianceFilterer) FilterTokenUnbound(opts *bind.FilterOpts) (*ModularcomplianceTokenUnboundIterator, error) {

	logs, sub, err := _Modularcompliance.contract.FilterLogs(opts, "TokenUnbound")
	if err != nil {
		return nil, err
	}
	return &ModularcomplianceTokenUnboundIterator{contract: _Modularcompliance.contract, event: "TokenUnbound", logs: logs, sub: sub}, nil
}

// WatchTokenUnbound is a free log subscription operation binding the contract event 0x28a4ca7134a3b3f9aff286e79ad3daadb4a06d1b43d037a3a98bdc074edd9b7a.
//
// Solidity: event TokenUnbound(address _token)
func (_Modularcompliance *ModularcomplianceFilterer) WatchTokenUnbound(opts *bind.WatchOpts, sink chan<- *ModularcomplianceTokenUnbound) (event.Subscription, error) {

	logs, sub, err := _Modularcompliance.contract.WatchLogs(opts, "TokenUnbound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModularcomplianceTokenUnbound)
				if err := _Modularcompliance.contract.UnpackLog(event, "TokenUnbound", log); err != nil {
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

// ParseTokenUnbound is a log parse operation binding the contract event 0x28a4ca7134a3b3f9aff286e79ad3daadb4a06d1b43d037a3a98bdc074edd9b7a.
//
// Solidity: event TokenUnbound(address _token)
func (_Modularcompliance *ModularcomplianceFilterer) ParseTokenUnbound(log types.Log) (*ModularcomplianceTokenUnbound, error) {
	event := new(ModularcomplianceTokenUnbound)
	if err := _Modularcompliance.contract.UnpackLog(event, "TokenUnbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
