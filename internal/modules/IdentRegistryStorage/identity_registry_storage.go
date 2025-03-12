// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identregistrystorage

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

// IdentregistrystorageMetaData contains all meta data concerning the Identregistrystorage contract.
var IdentregistrystorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addIdentityToStorage\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identity\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"},{\"name\":\"_country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bindIdentityRegistry\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkedIdentityRegistries\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyStoredIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identity\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyStoredInvestorCountry\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeIdentityFromStorage\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storedIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"storedInvestorCountry\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbindIdentityRegistry\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AgentAdded\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentRemoved\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CountryModified\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"country\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityModified\",\"inputs\":[{\"name\":\"oldIdentity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"},{\"name\":\"newIdentity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRegistryBound\",\"inputs\":[{\"name\":\"identityRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRegistryUnbound\",\"inputs\":[{\"name\":\"identityRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityStored\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityUnstored\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// IdentregistrystorageABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentregistrystorageMetaData.ABI instead.
var IdentregistrystorageABI = IdentregistrystorageMetaData.ABI

// Identregistrystorage is an auto generated Go binding around an Ethereum contract.
type Identregistrystorage struct {
	IdentregistrystorageCaller     // Read-only binding to the contract
	IdentregistrystorageTransactor // Write-only binding to the contract
	IdentregistrystorageFilterer   // Log filterer for contract events
}

// IdentregistrystorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentregistrystorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentregistrystorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentregistrystorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentregistrystorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentregistrystorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentregistrystorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentregistrystorageSession struct {
	Contract     *Identregistrystorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IdentregistrystorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentregistrystorageCallerSession struct {
	Contract *IdentregistrystorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IdentregistrystorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentregistrystorageTransactorSession struct {
	Contract     *IdentregistrystorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IdentregistrystorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentregistrystorageRaw struct {
	Contract *Identregistrystorage // Generic contract binding to access the raw methods on
}

// IdentregistrystorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentregistrystorageCallerRaw struct {
	Contract *IdentregistrystorageCaller // Generic read-only contract binding to access the raw methods on
}

// IdentregistrystorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentregistrystorageTransactorRaw struct {
	Contract *IdentregistrystorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentregistrystorage creates a new instance of Identregistrystorage, bound to a specific deployed contract.
func NewIdentregistrystorage(address common.Address, backend bind.ContractBackend) (*Identregistrystorage, error) {
	contract, err := bindIdentregistrystorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identregistrystorage{IdentregistrystorageCaller: IdentregistrystorageCaller{contract: contract}, IdentregistrystorageTransactor: IdentregistrystorageTransactor{contract: contract}, IdentregistrystorageFilterer: IdentregistrystorageFilterer{contract: contract}}, nil
}

// NewIdentregistrystorageCaller creates a new read-only instance of Identregistrystorage, bound to a specific deployed contract.
func NewIdentregistrystorageCaller(address common.Address, caller bind.ContractCaller) (*IdentregistrystorageCaller, error) {
	contract, err := bindIdentregistrystorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageCaller{contract: contract}, nil
}

// NewIdentregistrystorageTransactor creates a new write-only instance of Identregistrystorage, bound to a specific deployed contract.
func NewIdentregistrystorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentregistrystorageTransactor, error) {
	contract, err := bindIdentregistrystorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageTransactor{contract: contract}, nil
}

// NewIdentregistrystorageFilterer creates a new log filterer instance of Identregistrystorage, bound to a specific deployed contract.
func NewIdentregistrystorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentregistrystorageFilterer, error) {
	contract, err := bindIdentregistrystorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageFilterer{contract: contract}, nil
}

// bindIdentregistrystorage binds a generic wrapper to an already deployed contract.
func bindIdentregistrystorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentregistrystorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identregistrystorage *IdentregistrystorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identregistrystorage.Contract.IdentregistrystorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identregistrystorage *IdentregistrystorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.IdentregistrystorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identregistrystorage *IdentregistrystorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.IdentregistrystorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identregistrystorage *IdentregistrystorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identregistrystorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identregistrystorage *IdentregistrystorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identregistrystorage *IdentregistrystorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.contract.Transact(opts, method, params...)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Identregistrystorage *IdentregistrystorageCaller) IsAgent(opts *bind.CallOpts, _agent common.Address) (bool, error) {
	var out []interface{}
	err := _Identregistrystorage.contract.Call(opts, &out, "isAgent", _agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Identregistrystorage *IdentregistrystorageSession) IsAgent(_agent common.Address) (bool, error) {
	return _Identregistrystorage.Contract.IsAgent(&_Identregistrystorage.CallOpts, _agent)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Identregistrystorage *IdentregistrystorageCallerSession) IsAgent(_agent common.Address) (bool, error) {
	return _Identregistrystorage.Contract.IsAgent(&_Identregistrystorage.CallOpts, _agent)
}

// LinkedIdentityRegistries is a free data retrieval call binding the contract method 0xbf9eb959.
//
// Solidity: function linkedIdentityRegistries() view returns(address[])
func (_Identregistrystorage *IdentregistrystorageCaller) LinkedIdentityRegistries(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Identregistrystorage.contract.Call(opts, &out, "linkedIdentityRegistries")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LinkedIdentityRegistries is a free data retrieval call binding the contract method 0xbf9eb959.
//
// Solidity: function linkedIdentityRegistries() view returns(address[])
func (_Identregistrystorage *IdentregistrystorageSession) LinkedIdentityRegistries() ([]common.Address, error) {
	return _Identregistrystorage.Contract.LinkedIdentityRegistries(&_Identregistrystorage.CallOpts)
}

// LinkedIdentityRegistries is a free data retrieval call binding the contract method 0xbf9eb959.
//
// Solidity: function linkedIdentityRegistries() view returns(address[])
func (_Identregistrystorage *IdentregistrystorageCallerSession) LinkedIdentityRegistries() ([]common.Address, error) {
	return _Identregistrystorage.Contract.LinkedIdentityRegistries(&_Identregistrystorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identregistrystorage *IdentregistrystorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identregistrystorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identregistrystorage *IdentregistrystorageSession) Owner() (common.Address, error) {
	return _Identregistrystorage.Contract.Owner(&_Identregistrystorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identregistrystorage *IdentregistrystorageCallerSession) Owner() (common.Address, error) {
	return _Identregistrystorage.Contract.Owner(&_Identregistrystorage.CallOpts)
}

// StoredIdentity is a free data retrieval call binding the contract method 0x7988d3a5.
//
// Solidity: function storedIdentity(address _userAddress) view returns(address)
func (_Identregistrystorage *IdentregistrystorageCaller) StoredIdentity(opts *bind.CallOpts, _userAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Identregistrystorage.contract.Call(opts, &out, "storedIdentity", _userAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StoredIdentity is a free data retrieval call binding the contract method 0x7988d3a5.
//
// Solidity: function storedIdentity(address _userAddress) view returns(address)
func (_Identregistrystorage *IdentregistrystorageSession) StoredIdentity(_userAddress common.Address) (common.Address, error) {
	return _Identregistrystorage.Contract.StoredIdentity(&_Identregistrystorage.CallOpts, _userAddress)
}

// StoredIdentity is a free data retrieval call binding the contract method 0x7988d3a5.
//
// Solidity: function storedIdentity(address _userAddress) view returns(address)
func (_Identregistrystorage *IdentregistrystorageCallerSession) StoredIdentity(_userAddress common.Address) (common.Address, error) {
	return _Identregistrystorage.Contract.StoredIdentity(&_Identregistrystorage.CallOpts, _userAddress)
}

// StoredInvestorCountry is a free data retrieval call binding the contract method 0x727e13bc.
//
// Solidity: function storedInvestorCountry(address _userAddress) view returns(uint16)
func (_Identregistrystorage *IdentregistrystorageCaller) StoredInvestorCountry(opts *bind.CallOpts, _userAddress common.Address) (uint16, error) {
	var out []interface{}
	err := _Identregistrystorage.contract.Call(opts, &out, "storedInvestorCountry", _userAddress)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// StoredInvestorCountry is a free data retrieval call binding the contract method 0x727e13bc.
//
// Solidity: function storedInvestorCountry(address _userAddress) view returns(uint16)
func (_Identregistrystorage *IdentregistrystorageSession) StoredInvestorCountry(_userAddress common.Address) (uint16, error) {
	return _Identregistrystorage.Contract.StoredInvestorCountry(&_Identregistrystorage.CallOpts, _userAddress)
}

// StoredInvestorCountry is a free data retrieval call binding the contract method 0x727e13bc.
//
// Solidity: function storedInvestorCountry(address _userAddress) view returns(uint16)
func (_Identregistrystorage *IdentregistrystorageCallerSession) StoredInvestorCountry(_userAddress common.Address) (uint16, error) {
	return _Identregistrystorage.Contract.StoredInvestorCountry(&_Identregistrystorage.CallOpts, _userAddress)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) AddAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "addAgent", _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Identregistrystorage *IdentregistrystorageSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.AddAgent(&_Identregistrystorage.TransactOpts, _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.AddAgent(&_Identregistrystorage.TransactOpts, _agent)
}

// AddIdentityToStorage is a paid mutator transaction binding the contract method 0xa53410dd.
//
// Solidity: function addIdentityToStorage(address _userAddress, address _identity, uint16 _country) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) AddIdentityToStorage(opts *bind.TransactOpts, _userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "addIdentityToStorage", _userAddress, _identity, _country)
}

// AddIdentityToStorage is a paid mutator transaction binding the contract method 0xa53410dd.
//
// Solidity: function addIdentityToStorage(address _userAddress, address _identity, uint16 _country) returns()
func (_Identregistrystorage *IdentregistrystorageSession) AddIdentityToStorage(_userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.AddIdentityToStorage(&_Identregistrystorage.TransactOpts, _userAddress, _identity, _country)
}

// AddIdentityToStorage is a paid mutator transaction binding the contract method 0xa53410dd.
//
// Solidity: function addIdentityToStorage(address _userAddress, address _identity, uint16 _country) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) AddIdentityToStorage(_userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.AddIdentityToStorage(&_Identregistrystorage.TransactOpts, _userAddress, _identity, _country)
}

// BindIdentityRegistry is a paid mutator transaction binding the contract method 0x690a49f9.
//
// Solidity: function bindIdentityRegistry(address _identityRegistry) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) BindIdentityRegistry(opts *bind.TransactOpts, _identityRegistry common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "bindIdentityRegistry", _identityRegistry)
}

// BindIdentityRegistry is a paid mutator transaction binding the contract method 0x690a49f9.
//
// Solidity: function bindIdentityRegistry(address _identityRegistry) returns()
func (_Identregistrystorage *IdentregistrystorageSession) BindIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.BindIdentityRegistry(&_Identregistrystorage.TransactOpts, _identityRegistry)
}

// BindIdentityRegistry is a paid mutator transaction binding the contract method 0x690a49f9.
//
// Solidity: function bindIdentityRegistry(address _identityRegistry) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) BindIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.BindIdentityRegistry(&_Identregistrystorage.TransactOpts, _identityRegistry)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Identregistrystorage *IdentregistrystorageSession) Init() (*types.Transaction, error) {
	return _Identregistrystorage.Contract.Init(&_Identregistrystorage.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) Init() (*types.Transaction, error) {
	return _Identregistrystorage.Contract.Init(&_Identregistrystorage.TransactOpts)
}

// ModifyStoredIdentity is a paid mutator transaction binding the contract method 0xe805cf86.
//
// Solidity: function modifyStoredIdentity(address _userAddress, address _identity) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) ModifyStoredIdentity(opts *bind.TransactOpts, _userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "modifyStoredIdentity", _userAddress, _identity)
}

// ModifyStoredIdentity is a paid mutator transaction binding the contract method 0xe805cf86.
//
// Solidity: function modifyStoredIdentity(address _userAddress, address _identity) returns()
func (_Identregistrystorage *IdentregistrystorageSession) ModifyStoredIdentity(_userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.ModifyStoredIdentity(&_Identregistrystorage.TransactOpts, _userAddress, _identity)
}

// ModifyStoredIdentity is a paid mutator transaction binding the contract method 0xe805cf86.
//
// Solidity: function modifyStoredIdentity(address _userAddress, address _identity) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) ModifyStoredIdentity(_userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.ModifyStoredIdentity(&_Identregistrystorage.TransactOpts, _userAddress, _identity)
}

// ModifyStoredInvestorCountry is a paid mutator transaction binding the contract method 0x9f3418d5.
//
// Solidity: function modifyStoredInvestorCountry(address _userAddress, uint16 _country) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) ModifyStoredInvestorCountry(opts *bind.TransactOpts, _userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "modifyStoredInvestorCountry", _userAddress, _country)
}

// ModifyStoredInvestorCountry is a paid mutator transaction binding the contract method 0x9f3418d5.
//
// Solidity: function modifyStoredInvestorCountry(address _userAddress, uint16 _country) returns()
func (_Identregistrystorage *IdentregistrystorageSession) ModifyStoredInvestorCountry(_userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.ModifyStoredInvestorCountry(&_Identregistrystorage.TransactOpts, _userAddress, _country)
}

// ModifyStoredInvestorCountry is a paid mutator transaction binding the contract method 0x9f3418d5.
//
// Solidity: function modifyStoredInvestorCountry(address _userAddress, uint16 _country) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) ModifyStoredInvestorCountry(_userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.ModifyStoredInvestorCountry(&_Identregistrystorage.TransactOpts, _userAddress, _country)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) RemoveAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "removeAgent", _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Identregistrystorage *IdentregistrystorageSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.RemoveAgent(&_Identregistrystorage.TransactOpts, _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.RemoveAgent(&_Identregistrystorage.TransactOpts, _agent)
}

// RemoveIdentityFromStorage is a paid mutator transaction binding the contract method 0xcf191bcd.
//
// Solidity: function removeIdentityFromStorage(address _userAddress) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) RemoveIdentityFromStorage(opts *bind.TransactOpts, _userAddress common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "removeIdentityFromStorage", _userAddress)
}

// RemoveIdentityFromStorage is a paid mutator transaction binding the contract method 0xcf191bcd.
//
// Solidity: function removeIdentityFromStorage(address _userAddress) returns()
func (_Identregistrystorage *IdentregistrystorageSession) RemoveIdentityFromStorage(_userAddress common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.RemoveIdentityFromStorage(&_Identregistrystorage.TransactOpts, _userAddress)
}

// RemoveIdentityFromStorage is a paid mutator transaction binding the contract method 0xcf191bcd.
//
// Solidity: function removeIdentityFromStorage(address _userAddress) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) RemoveIdentityFromStorage(_userAddress common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.RemoveIdentityFromStorage(&_Identregistrystorage.TransactOpts, _userAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identregistrystorage *IdentregistrystorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Identregistrystorage.Contract.RenounceOwnership(&_Identregistrystorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Identregistrystorage.Contract.RenounceOwnership(&_Identregistrystorage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identregistrystorage *IdentregistrystorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.TransferOwnership(&_Identregistrystorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.TransferOwnership(&_Identregistrystorage.TransactOpts, newOwner)
}

// UnbindIdentityRegistry is a paid mutator transaction binding the contract method 0x97a012f7.
//
// Solidity: function unbindIdentityRegistry(address _identityRegistry) returns()
func (_Identregistrystorage *IdentregistrystorageTransactor) UnbindIdentityRegistry(opts *bind.TransactOpts, _identityRegistry common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.contract.Transact(opts, "unbindIdentityRegistry", _identityRegistry)
}

// UnbindIdentityRegistry is a paid mutator transaction binding the contract method 0x97a012f7.
//
// Solidity: function unbindIdentityRegistry(address _identityRegistry) returns()
func (_Identregistrystorage *IdentregistrystorageSession) UnbindIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.UnbindIdentityRegistry(&_Identregistrystorage.TransactOpts, _identityRegistry)
}

// UnbindIdentityRegistry is a paid mutator transaction binding the contract method 0x97a012f7.
//
// Solidity: function unbindIdentityRegistry(address _identityRegistry) returns()
func (_Identregistrystorage *IdentregistrystorageTransactorSession) UnbindIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _Identregistrystorage.Contract.UnbindIdentityRegistry(&_Identregistrystorage.TransactOpts, _identityRegistry)
}

// IdentregistrystorageAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the Identregistrystorage contract.
type IdentregistrystorageAgentAddedIterator struct {
	Event *IdentregistrystorageAgentAdded // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageAgentAdded)
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
		it.Event = new(IdentregistrystorageAgentAdded)
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
func (it *IdentregistrystorageAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageAgentAdded represents a AgentAdded event raised by the Identregistrystorage contract.
type IdentregistrystorageAgentAdded struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterAgentAdded(opts *bind.FilterOpts, _agent []common.Address) (*IdentregistrystorageAgentAddedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageAgentAddedIterator{contract: _Identregistrystorage.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageAgentAdded, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageAgentAdded)
				if err := _Identregistrystorage.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseAgentAdded(log types.Log) (*IdentregistrystorageAgentAdded, error) {
	event := new(IdentregistrystorageAgentAdded)
	if err := _Identregistrystorage.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the Identregistrystorage contract.
type IdentregistrystorageAgentRemovedIterator struct {
	Event *IdentregistrystorageAgentRemoved // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageAgentRemoved)
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
		it.Event = new(IdentregistrystorageAgentRemoved)
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
func (it *IdentregistrystorageAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageAgentRemoved represents a AgentRemoved event raised by the Identregistrystorage contract.
type IdentregistrystorageAgentRemoved struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterAgentRemoved(opts *bind.FilterOpts, _agent []common.Address) (*IdentregistrystorageAgentRemovedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageAgentRemovedIterator{contract: _Identregistrystorage.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageAgentRemoved, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageAgentRemoved)
				if err := _Identregistrystorage.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseAgentRemoved(log types.Log) (*IdentregistrystorageAgentRemoved, error) {
	event := new(IdentregistrystorageAgentRemoved)
	if err := _Identregistrystorage.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageCountryModifiedIterator is returned from FilterCountryModified and is used to iterate over the raw logs and unpacked data for CountryModified events raised by the Identregistrystorage contract.
type IdentregistrystorageCountryModifiedIterator struct {
	Event *IdentregistrystorageCountryModified // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageCountryModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageCountryModified)
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
		it.Event = new(IdentregistrystorageCountryModified)
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
func (it *IdentregistrystorageCountryModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageCountryModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageCountryModified represents a CountryModified event raised by the Identregistrystorage contract.
type IdentregistrystorageCountryModified struct {
	InvestorAddress common.Address
	Country         uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCountryModified is a free log retrieval operation binding the contract event 0x20965fcdc6eed7ae398065b40ece4e732ba667992ca819fc54e80e9f2047c4cf.
//
// Solidity: event CountryModified(address indexed investorAddress, uint16 indexed country)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterCountryModified(opts *bind.FilterOpts, investorAddress []common.Address, country []uint16) (*IdentregistrystorageCountryModifiedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var countryRule []interface{}
	for _, countryItem := range country {
		countryRule = append(countryRule, countryItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "CountryModified", investorAddressRule, countryRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageCountryModifiedIterator{contract: _Identregistrystorage.contract, event: "CountryModified", logs: logs, sub: sub}, nil
}

// WatchCountryModified is a free log subscription operation binding the contract event 0x20965fcdc6eed7ae398065b40ece4e732ba667992ca819fc54e80e9f2047c4cf.
//
// Solidity: event CountryModified(address indexed investorAddress, uint16 indexed country)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchCountryModified(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageCountryModified, investorAddress []common.Address, country []uint16) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var countryRule []interface{}
	for _, countryItem := range country {
		countryRule = append(countryRule, countryItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "CountryModified", investorAddressRule, countryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageCountryModified)
				if err := _Identregistrystorage.contract.UnpackLog(event, "CountryModified", log); err != nil {
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

// ParseCountryModified is a log parse operation binding the contract event 0x20965fcdc6eed7ae398065b40ece4e732ba667992ca819fc54e80e9f2047c4cf.
//
// Solidity: event CountryModified(address indexed investorAddress, uint16 indexed country)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseCountryModified(log types.Log) (*IdentregistrystorageCountryModified, error) {
	event := new(IdentregistrystorageCountryModified)
	if err := _Identregistrystorage.contract.UnpackLog(event, "CountryModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageIdentityModifiedIterator is returned from FilterIdentityModified and is used to iterate over the raw logs and unpacked data for IdentityModified events raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityModifiedIterator struct {
	Event *IdentregistrystorageIdentityModified // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageIdentityModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageIdentityModified)
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
		it.Event = new(IdentregistrystorageIdentityModified)
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
func (it *IdentregistrystorageIdentityModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageIdentityModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageIdentityModified represents a IdentityModified event raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityModified struct {
	OldIdentity common.Address
	NewIdentity common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIdentityModified is a free log retrieval operation binding the contract event 0x556ce885dfcea52155c773f1ed2e58781c51945c13030ab8f793c61f51d1b808.
//
// Solidity: event IdentityModified(address indexed oldIdentity, address indexed newIdentity)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterIdentityModified(opts *bind.FilterOpts, oldIdentity []common.Address, newIdentity []common.Address) (*IdentregistrystorageIdentityModifiedIterator, error) {

	var oldIdentityRule []interface{}
	for _, oldIdentityItem := range oldIdentity {
		oldIdentityRule = append(oldIdentityRule, oldIdentityItem)
	}
	var newIdentityRule []interface{}
	for _, newIdentityItem := range newIdentity {
		newIdentityRule = append(newIdentityRule, newIdentityItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "IdentityModified", oldIdentityRule, newIdentityRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageIdentityModifiedIterator{contract: _Identregistrystorage.contract, event: "IdentityModified", logs: logs, sub: sub}, nil
}

// WatchIdentityModified is a free log subscription operation binding the contract event 0x556ce885dfcea52155c773f1ed2e58781c51945c13030ab8f793c61f51d1b808.
//
// Solidity: event IdentityModified(address indexed oldIdentity, address indexed newIdentity)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchIdentityModified(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageIdentityModified, oldIdentity []common.Address, newIdentity []common.Address) (event.Subscription, error) {

	var oldIdentityRule []interface{}
	for _, oldIdentityItem := range oldIdentity {
		oldIdentityRule = append(oldIdentityRule, oldIdentityItem)
	}
	var newIdentityRule []interface{}
	for _, newIdentityItem := range newIdentity {
		newIdentityRule = append(newIdentityRule, newIdentityItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "IdentityModified", oldIdentityRule, newIdentityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageIdentityModified)
				if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityModified", log); err != nil {
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

// ParseIdentityModified is a log parse operation binding the contract event 0x556ce885dfcea52155c773f1ed2e58781c51945c13030ab8f793c61f51d1b808.
//
// Solidity: event IdentityModified(address indexed oldIdentity, address indexed newIdentity)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseIdentityModified(log types.Log) (*IdentregistrystorageIdentityModified, error) {
	event := new(IdentregistrystorageIdentityModified)
	if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageIdentityRegistryBoundIterator is returned from FilterIdentityRegistryBound and is used to iterate over the raw logs and unpacked data for IdentityRegistryBound events raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityRegistryBoundIterator struct {
	Event *IdentregistrystorageIdentityRegistryBound // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageIdentityRegistryBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageIdentityRegistryBound)
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
		it.Event = new(IdentregistrystorageIdentityRegistryBound)
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
func (it *IdentregistrystorageIdentityRegistryBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageIdentityRegistryBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageIdentityRegistryBound represents a IdentityRegistryBound event raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityRegistryBound struct {
	IdentityRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIdentityRegistryBound is a free log retrieval operation binding the contract event 0x500c250171aa20e861b680f93502547b9d436eda7d4c537fc360db6e0c6eedfb.
//
// Solidity: event IdentityRegistryBound(address indexed identityRegistry)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterIdentityRegistryBound(opts *bind.FilterOpts, identityRegistry []common.Address) (*IdentregistrystorageIdentityRegistryBoundIterator, error) {

	var identityRegistryRule []interface{}
	for _, identityRegistryItem := range identityRegistry {
		identityRegistryRule = append(identityRegistryRule, identityRegistryItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "IdentityRegistryBound", identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageIdentityRegistryBoundIterator{contract: _Identregistrystorage.contract, event: "IdentityRegistryBound", logs: logs, sub: sub}, nil
}

// WatchIdentityRegistryBound is a free log subscription operation binding the contract event 0x500c250171aa20e861b680f93502547b9d436eda7d4c537fc360db6e0c6eedfb.
//
// Solidity: event IdentityRegistryBound(address indexed identityRegistry)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchIdentityRegistryBound(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageIdentityRegistryBound, identityRegistry []common.Address) (event.Subscription, error) {

	var identityRegistryRule []interface{}
	for _, identityRegistryItem := range identityRegistry {
		identityRegistryRule = append(identityRegistryRule, identityRegistryItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "IdentityRegistryBound", identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageIdentityRegistryBound)
				if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityRegistryBound", log); err != nil {
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

// ParseIdentityRegistryBound is a log parse operation binding the contract event 0x500c250171aa20e861b680f93502547b9d436eda7d4c537fc360db6e0c6eedfb.
//
// Solidity: event IdentityRegistryBound(address indexed identityRegistry)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseIdentityRegistryBound(log types.Log) (*IdentregistrystorageIdentityRegistryBound, error) {
	event := new(IdentregistrystorageIdentityRegistryBound)
	if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityRegistryBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageIdentityRegistryUnboundIterator is returned from FilterIdentityRegistryUnbound and is used to iterate over the raw logs and unpacked data for IdentityRegistryUnbound events raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityRegistryUnboundIterator struct {
	Event *IdentregistrystorageIdentityRegistryUnbound // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageIdentityRegistryUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageIdentityRegistryUnbound)
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
		it.Event = new(IdentregistrystorageIdentityRegistryUnbound)
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
func (it *IdentregistrystorageIdentityRegistryUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageIdentityRegistryUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageIdentityRegistryUnbound represents a IdentityRegistryUnbound event raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityRegistryUnbound struct {
	IdentityRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIdentityRegistryUnbound is a free log retrieval operation binding the contract event 0x51f353eb5801583fdf2706e43c045b62fdf6b1566820b349390616363ecf72c9.
//
// Solidity: event IdentityRegistryUnbound(address indexed identityRegistry)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterIdentityRegistryUnbound(opts *bind.FilterOpts, identityRegistry []common.Address) (*IdentregistrystorageIdentityRegistryUnboundIterator, error) {

	var identityRegistryRule []interface{}
	for _, identityRegistryItem := range identityRegistry {
		identityRegistryRule = append(identityRegistryRule, identityRegistryItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "IdentityRegistryUnbound", identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageIdentityRegistryUnboundIterator{contract: _Identregistrystorage.contract, event: "IdentityRegistryUnbound", logs: logs, sub: sub}, nil
}

// WatchIdentityRegistryUnbound is a free log subscription operation binding the contract event 0x51f353eb5801583fdf2706e43c045b62fdf6b1566820b349390616363ecf72c9.
//
// Solidity: event IdentityRegistryUnbound(address indexed identityRegistry)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchIdentityRegistryUnbound(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageIdentityRegistryUnbound, identityRegistry []common.Address) (event.Subscription, error) {

	var identityRegistryRule []interface{}
	for _, identityRegistryItem := range identityRegistry {
		identityRegistryRule = append(identityRegistryRule, identityRegistryItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "IdentityRegistryUnbound", identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageIdentityRegistryUnbound)
				if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityRegistryUnbound", log); err != nil {
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

// ParseIdentityRegistryUnbound is a log parse operation binding the contract event 0x51f353eb5801583fdf2706e43c045b62fdf6b1566820b349390616363ecf72c9.
//
// Solidity: event IdentityRegistryUnbound(address indexed identityRegistry)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseIdentityRegistryUnbound(log types.Log) (*IdentregistrystorageIdentityRegistryUnbound, error) {
	event := new(IdentregistrystorageIdentityRegistryUnbound)
	if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityRegistryUnbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageIdentityStoredIterator is returned from FilterIdentityStored and is used to iterate over the raw logs and unpacked data for IdentityStored events raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityStoredIterator struct {
	Event *IdentregistrystorageIdentityStored // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageIdentityStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageIdentityStored)
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
		it.Event = new(IdentregistrystorageIdentityStored)
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
func (it *IdentregistrystorageIdentityStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageIdentityStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageIdentityStored represents a IdentityStored event raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityStored struct {
	InvestorAddress common.Address
	Identity        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityStored is a free log retrieval operation binding the contract event 0x0030dea7e9c9afaa2e3c9810f2fc9b5181f1bad74ca5a8db85f746a33585e747.
//
// Solidity: event IdentityStored(address indexed investorAddress, address indexed identity)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterIdentityStored(opts *bind.FilterOpts, investorAddress []common.Address, identity []common.Address) (*IdentregistrystorageIdentityStoredIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "IdentityStored", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageIdentityStoredIterator{contract: _Identregistrystorage.contract, event: "IdentityStored", logs: logs, sub: sub}, nil
}

// WatchIdentityStored is a free log subscription operation binding the contract event 0x0030dea7e9c9afaa2e3c9810f2fc9b5181f1bad74ca5a8db85f746a33585e747.
//
// Solidity: event IdentityStored(address indexed investorAddress, address indexed identity)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchIdentityStored(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageIdentityStored, investorAddress []common.Address, identity []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "IdentityStored", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageIdentityStored)
				if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityStored", log); err != nil {
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

// ParseIdentityStored is a log parse operation binding the contract event 0x0030dea7e9c9afaa2e3c9810f2fc9b5181f1bad74ca5a8db85f746a33585e747.
//
// Solidity: event IdentityStored(address indexed investorAddress, address indexed identity)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseIdentityStored(log types.Log) (*IdentregistrystorageIdentityStored, error) {
	event := new(IdentregistrystorageIdentityStored)
	if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageIdentityUnstoredIterator is returned from FilterIdentityUnstored and is used to iterate over the raw logs and unpacked data for IdentityUnstored events raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityUnstoredIterator struct {
	Event *IdentregistrystorageIdentityUnstored // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageIdentityUnstoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageIdentityUnstored)
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
		it.Event = new(IdentregistrystorageIdentityUnstored)
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
func (it *IdentregistrystorageIdentityUnstoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageIdentityUnstoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageIdentityUnstored represents a IdentityUnstored event raised by the Identregistrystorage contract.
type IdentregistrystorageIdentityUnstored struct {
	InvestorAddress common.Address
	Identity        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityUnstored is a free log retrieval operation binding the contract event 0xca6a4c3370b859312246e7f086284076e557997e10d856b716c23ab67067790b.
//
// Solidity: event IdentityUnstored(address indexed investorAddress, address indexed identity)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterIdentityUnstored(opts *bind.FilterOpts, investorAddress []common.Address, identity []common.Address) (*IdentregistrystorageIdentityUnstoredIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "IdentityUnstored", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageIdentityUnstoredIterator{contract: _Identregistrystorage.contract, event: "IdentityUnstored", logs: logs, sub: sub}, nil
}

// WatchIdentityUnstored is a free log subscription operation binding the contract event 0xca6a4c3370b859312246e7f086284076e557997e10d856b716c23ab67067790b.
//
// Solidity: event IdentityUnstored(address indexed investorAddress, address indexed identity)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchIdentityUnstored(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageIdentityUnstored, investorAddress []common.Address, identity []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "IdentityUnstored", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageIdentityUnstored)
				if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityUnstored", log); err != nil {
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

// ParseIdentityUnstored is a log parse operation binding the contract event 0xca6a4c3370b859312246e7f086284076e557997e10d856b716c23ab67067790b.
//
// Solidity: event IdentityUnstored(address indexed investorAddress, address indexed identity)
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseIdentityUnstored(log types.Log) (*IdentregistrystorageIdentityUnstored, error) {
	event := new(IdentregistrystorageIdentityUnstored)
	if err := _Identregistrystorage.contract.UnpackLog(event, "IdentityUnstored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Identregistrystorage contract.
type IdentregistrystorageInitializedIterator struct {
	Event *IdentregistrystorageInitialized // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageInitialized)
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
		it.Event = new(IdentregistrystorageInitialized)
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
func (it *IdentregistrystorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageInitialized represents a Initialized event raised by the Identregistrystorage contract.
type IdentregistrystorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*IdentregistrystorageInitializedIterator, error) {

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageInitializedIterator{contract: _Identregistrystorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageInitialized)
				if err := _Identregistrystorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseInitialized(log types.Log) (*IdentregistrystorageInitialized, error) {
	event := new(IdentregistrystorageInitialized)
	if err := _Identregistrystorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentregistrystorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Identregistrystorage contract.
type IdentregistrystorageOwnershipTransferredIterator struct {
	Event *IdentregistrystorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentregistrystorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentregistrystorageOwnershipTransferred)
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
		it.Event = new(IdentregistrystorageOwnershipTransferred)
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
func (it *IdentregistrystorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentregistrystorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentregistrystorageOwnershipTransferred represents a OwnershipTransferred event raised by the Identregistrystorage contract.
type IdentregistrystorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identregistrystorage *IdentregistrystorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentregistrystorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identregistrystorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentregistrystorageOwnershipTransferredIterator{contract: _Identregistrystorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identregistrystorage *IdentregistrystorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentregistrystorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identregistrystorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentregistrystorageOwnershipTransferred)
				if err := _Identregistrystorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Identregistrystorage *IdentregistrystorageFilterer) ParseOwnershipTransferred(log types.Log) (*IdentregistrystorageOwnershipTransferred, error) {
	event := new(IdentregistrystorageOwnershipTransferred)
	if err := _Identregistrystorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
