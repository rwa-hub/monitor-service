// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identityregistry

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

// IdentityregistryMetaData contains all meta data concerning the Identityregistry contract.
var IdentityregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchRegisterIdentity\",\"inputs\":[{\"name\":\"_userAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_identities\",\"type\":\"address[]\",\"internalType\":\"contractIIdentity[]\"},{\"name\":\"_countries\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contains\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"identity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentityRegistryStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_trustedIssuersRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_claimTopicsRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identityStorage\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"investorCountry\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isVerified\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"issuersRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITrustedIssuersRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identity\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"},{\"name\":\"_country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimTopicsRegistry\",\"inputs\":[{\"name\":\"_claimTopicsRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIdentityRegistryStorage\",\"inputs\":[{\"name\":\"_identityRegistryStorage\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrustedIssuersRegistry\",\"inputs\":[{\"name\":\"_trustedIssuersRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topicsRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIClaimTopicsRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCountry\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identity\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AgentAdded\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentRemoved\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimTopicsRegistrySet\",\"inputs\":[{\"name\":\"claimTopicsRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CountryUpdated\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"country\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRegistered\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRemoved\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityStorageSet\",\"inputs\":[{\"name\":\"identityStorage\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityUpdated\",\"inputs\":[{\"name\":\"oldIdentity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"},{\"name\":\"newIdentity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TrustedIssuersRegistrySet\",\"inputs\":[{\"name\":\"trustedIssuersRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// IdentityregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityregistryMetaData.ABI instead.
var IdentityregistryABI = IdentityregistryMetaData.ABI

// Identityregistry is an auto generated Go binding around an Ethereum contract.
type Identityregistry struct {
	IdentityregistryCaller     // Read-only binding to the contract
	IdentityregistryTransactor // Write-only binding to the contract
	IdentityregistryFilterer   // Log filterer for contract events
}

// IdentityregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityregistrySession struct {
	Contract     *Identityregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityregistryCallerSession struct {
	Contract *IdentityregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityregistryTransactorSession struct {
	Contract     *IdentityregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityregistryRaw struct {
	Contract *Identityregistry // Generic contract binding to access the raw methods on
}

// IdentityregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityregistryCallerRaw struct {
	Contract *IdentityregistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityregistryTransactorRaw struct {
	Contract *IdentityregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityregistry creates a new instance of Identityregistry, bound to a specific deployed contract.
func NewIdentityregistry(address common.Address, backend bind.ContractBackend) (*Identityregistry, error) {
	contract, err := bindIdentityregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identityregistry{IdentityregistryCaller: IdentityregistryCaller{contract: contract}, IdentityregistryTransactor: IdentityregistryTransactor{contract: contract}, IdentityregistryFilterer: IdentityregistryFilterer{contract: contract}}, nil
}

// NewIdentityregistryCaller creates a new read-only instance of Identityregistry, bound to a specific deployed contract.
func NewIdentityregistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityregistryCaller, error) {
	contract, err := bindIdentityregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryCaller{contract: contract}, nil
}

// NewIdentityregistryTransactor creates a new write-only instance of Identityregistry, bound to a specific deployed contract.
func NewIdentityregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityregistryTransactor, error) {
	contract, err := bindIdentityregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryTransactor{contract: contract}, nil
}

// NewIdentityregistryFilterer creates a new log filterer instance of Identityregistry, bound to a specific deployed contract.
func NewIdentityregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityregistryFilterer, error) {
	contract, err := bindIdentityregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryFilterer{contract: contract}, nil
}

// bindIdentityregistry binds a generic wrapper to an already deployed contract.
func bindIdentityregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identityregistry *IdentityregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identityregistry.Contract.IdentityregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identityregistry *IdentityregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityregistry.Contract.IdentityregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identityregistry *IdentityregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identityregistry.Contract.IdentityregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identityregistry *IdentityregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identityregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identityregistry *IdentityregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identityregistry *IdentityregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identityregistry.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _userAddress) view returns(bool)
func (_Identityregistry *IdentityregistryCaller) Contains(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "contains", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _userAddress) view returns(bool)
func (_Identityregistry *IdentityregistrySession) Contains(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.Contains(&_Identityregistry.CallOpts, _userAddress)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _userAddress) view returns(bool)
func (_Identityregistry *IdentityregistryCallerSession) Contains(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.Contains(&_Identityregistry.CallOpts, _userAddress)
}

// Identity is a free data retrieval call binding the contract method 0xf0eb5e54.
//
// Solidity: function identity(address _userAddress) view returns(address)
func (_Identityregistry *IdentityregistryCaller) Identity(opts *bind.CallOpts, _userAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "identity", _userAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Identity is a free data retrieval call binding the contract method 0xf0eb5e54.
//
// Solidity: function identity(address _userAddress) view returns(address)
func (_Identityregistry *IdentityregistrySession) Identity(_userAddress common.Address) (common.Address, error) {
	return _Identityregistry.Contract.Identity(&_Identityregistry.CallOpts, _userAddress)
}

// Identity is a free data retrieval call binding the contract method 0xf0eb5e54.
//
// Solidity: function identity(address _userAddress) view returns(address)
func (_Identityregistry *IdentityregistryCallerSession) Identity(_userAddress common.Address) (common.Address, error) {
	return _Identityregistry.Contract.Identity(&_Identityregistry.CallOpts, _userAddress)
}

// IdentityStorage is a free data retrieval call binding the contract method 0xf11abfd8.
//
// Solidity: function identityStorage() view returns(address)
func (_Identityregistry *IdentityregistryCaller) IdentityStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "identityStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityStorage is a free data retrieval call binding the contract method 0xf11abfd8.
//
// Solidity: function identityStorage() view returns(address)
func (_Identityregistry *IdentityregistrySession) IdentityStorage() (common.Address, error) {
	return _Identityregistry.Contract.IdentityStorage(&_Identityregistry.CallOpts)
}

// IdentityStorage is a free data retrieval call binding the contract method 0xf11abfd8.
//
// Solidity: function identityStorage() view returns(address)
func (_Identityregistry *IdentityregistryCallerSession) IdentityStorage() (common.Address, error) {
	return _Identityregistry.Contract.IdentityStorage(&_Identityregistry.CallOpts)
}

// InvestorCountry is a free data retrieval call binding the contract method 0x7e42683b.
//
// Solidity: function investorCountry(address _userAddress) view returns(uint16)
func (_Identityregistry *IdentityregistryCaller) InvestorCountry(opts *bind.CallOpts, _userAddress common.Address) (uint16, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "investorCountry", _userAddress)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InvestorCountry is a free data retrieval call binding the contract method 0x7e42683b.
//
// Solidity: function investorCountry(address _userAddress) view returns(uint16)
func (_Identityregistry *IdentityregistrySession) InvestorCountry(_userAddress common.Address) (uint16, error) {
	return _Identityregistry.Contract.InvestorCountry(&_Identityregistry.CallOpts, _userAddress)
}

// InvestorCountry is a free data retrieval call binding the contract method 0x7e42683b.
//
// Solidity: function investorCountry(address _userAddress) view returns(uint16)
func (_Identityregistry *IdentityregistryCallerSession) InvestorCountry(_userAddress common.Address) (uint16, error) {
	return _Identityregistry.Contract.InvestorCountry(&_Identityregistry.CallOpts, _userAddress)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Identityregistry *IdentityregistryCaller) IsAgent(opts *bind.CallOpts, _agent common.Address) (bool, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "isAgent", _agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Identityregistry *IdentityregistrySession) IsAgent(_agent common.Address) (bool, error) {
	return _Identityregistry.Contract.IsAgent(&_Identityregistry.CallOpts, _agent)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Identityregistry *IdentityregistryCallerSession) IsAgent(_agent common.Address) (bool, error) {
	return _Identityregistry.Contract.IsAgent(&_Identityregistry.CallOpts, _agent)
}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address _userAddress) view returns(bool)
func (_Identityregistry *IdentityregistryCaller) IsVerified(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "isVerified", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address _userAddress) view returns(bool)
func (_Identityregistry *IdentityregistrySession) IsVerified(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.IsVerified(&_Identityregistry.CallOpts, _userAddress)
}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address _userAddress) view returns(bool)
func (_Identityregistry *IdentityregistryCallerSession) IsVerified(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.IsVerified(&_Identityregistry.CallOpts, _userAddress)
}

// IssuersRegistry is a free data retrieval call binding the contract method 0xb4f3fcb7.
//
// Solidity: function issuersRegistry() view returns(address)
func (_Identityregistry *IdentityregistryCaller) IssuersRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "issuersRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IssuersRegistry is a free data retrieval call binding the contract method 0xb4f3fcb7.
//
// Solidity: function issuersRegistry() view returns(address)
func (_Identityregistry *IdentityregistrySession) IssuersRegistry() (common.Address, error) {
	return _Identityregistry.Contract.IssuersRegistry(&_Identityregistry.CallOpts)
}

// IssuersRegistry is a free data retrieval call binding the contract method 0xb4f3fcb7.
//
// Solidity: function issuersRegistry() view returns(address)
func (_Identityregistry *IdentityregistryCallerSession) IssuersRegistry() (common.Address, error) {
	return _Identityregistry.Contract.IssuersRegistry(&_Identityregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identityregistry *IdentityregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identityregistry *IdentityregistrySession) Owner() (common.Address, error) {
	return _Identityregistry.Contract.Owner(&_Identityregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identityregistry *IdentityregistryCallerSession) Owner() (common.Address, error) {
	return _Identityregistry.Contract.Owner(&_Identityregistry.CallOpts)
}

// TopicsRegistry is a free data retrieval call binding the contract method 0x3b3e12f4.
//
// Solidity: function topicsRegistry() view returns(address)
func (_Identityregistry *IdentityregistryCaller) TopicsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "topicsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TopicsRegistry is a free data retrieval call binding the contract method 0x3b3e12f4.
//
// Solidity: function topicsRegistry() view returns(address)
func (_Identityregistry *IdentityregistrySession) TopicsRegistry() (common.Address, error) {
	return _Identityregistry.Contract.TopicsRegistry(&_Identityregistry.CallOpts)
}

// TopicsRegistry is a free data retrieval call binding the contract method 0x3b3e12f4.
//
// Solidity: function topicsRegistry() view returns(address)
func (_Identityregistry *IdentityregistryCallerSession) TopicsRegistry() (common.Address, error) {
	return _Identityregistry.Contract.TopicsRegistry(&_Identityregistry.CallOpts)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Identityregistry *IdentityregistryTransactor) AddAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "addAgent", _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Identityregistry *IdentityregistrySession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.AddAgent(&_Identityregistry.TransactOpts, _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Identityregistry *IdentityregistryTransactorSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.AddAgent(&_Identityregistry.TransactOpts, _agent)
}

// BatchRegisterIdentity is a paid mutator transaction binding the contract method 0x653dc9f1.
//
// Solidity: function batchRegisterIdentity(address[] _userAddresses, address[] _identities, uint16[] _countries) returns()
func (_Identityregistry *IdentityregistryTransactor) BatchRegisterIdentity(opts *bind.TransactOpts, _userAddresses []common.Address, _identities []common.Address, _countries []uint16) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "batchRegisterIdentity", _userAddresses, _identities, _countries)
}

// BatchRegisterIdentity is a paid mutator transaction binding the contract method 0x653dc9f1.
//
// Solidity: function batchRegisterIdentity(address[] _userAddresses, address[] _identities, uint16[] _countries) returns()
func (_Identityregistry *IdentityregistrySession) BatchRegisterIdentity(_userAddresses []common.Address, _identities []common.Address, _countries []uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.BatchRegisterIdentity(&_Identityregistry.TransactOpts, _userAddresses, _identities, _countries)
}

// BatchRegisterIdentity is a paid mutator transaction binding the contract method 0x653dc9f1.
//
// Solidity: function batchRegisterIdentity(address[] _userAddresses, address[] _identities, uint16[] _countries) returns()
func (_Identityregistry *IdentityregistryTransactorSession) BatchRegisterIdentity(_userAddresses []common.Address, _identities []common.Address, _countries []uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.BatchRegisterIdentity(&_Identityregistry.TransactOpts, _userAddresses, _identities, _countries)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address _userAddress) returns()
func (_Identityregistry *IdentityregistryTransactor) DeleteIdentity(opts *bind.TransactOpts, _userAddress common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "deleteIdentity", _userAddress)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address _userAddress) returns()
func (_Identityregistry *IdentityregistrySession) DeleteIdentity(_userAddress common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.DeleteIdentity(&_Identityregistry.TransactOpts, _userAddress)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address _userAddress) returns()
func (_Identityregistry *IdentityregistryTransactorSession) DeleteIdentity(_userAddress common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.DeleteIdentity(&_Identityregistry.TransactOpts, _userAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _trustedIssuersRegistry, address _claimTopicsRegistry, address _identityStorage) returns()
func (_Identityregistry *IdentityregistryTransactor) Init(opts *bind.TransactOpts, _trustedIssuersRegistry common.Address, _claimTopicsRegistry common.Address, _identityStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "init", _trustedIssuersRegistry, _claimTopicsRegistry, _identityStorage)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _trustedIssuersRegistry, address _claimTopicsRegistry, address _identityStorage) returns()
func (_Identityregistry *IdentityregistrySession) Init(_trustedIssuersRegistry common.Address, _claimTopicsRegistry common.Address, _identityStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.Init(&_Identityregistry.TransactOpts, _trustedIssuersRegistry, _claimTopicsRegistry, _identityStorage)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _trustedIssuersRegistry, address _claimTopicsRegistry, address _identityStorage) returns()
func (_Identityregistry *IdentityregistryTransactorSession) Init(_trustedIssuersRegistry common.Address, _claimTopicsRegistry common.Address, _identityStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.Init(&_Identityregistry.TransactOpts, _trustedIssuersRegistry, _claimTopicsRegistry, _identityStorage)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x454a03e0.
//
// Solidity: function registerIdentity(address _userAddress, address _identity, uint16 _country) returns()
func (_Identityregistry *IdentityregistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "registerIdentity", _userAddress, _identity, _country)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x454a03e0.
//
// Solidity: function registerIdentity(address _userAddress, address _identity, uint16 _country) returns()
func (_Identityregistry *IdentityregistrySession) RegisterIdentity(_userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.RegisterIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity, _country)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x454a03e0.
//
// Solidity: function registerIdentity(address _userAddress, address _identity, uint16 _country) returns()
func (_Identityregistry *IdentityregistryTransactorSession) RegisterIdentity(_userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.RegisterIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity, _country)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Identityregistry *IdentityregistryTransactor) RemoveAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "removeAgent", _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Identityregistry *IdentityregistrySession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.RemoveAgent(&_Identityregistry.TransactOpts, _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Identityregistry *IdentityregistryTransactorSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.RemoveAgent(&_Identityregistry.TransactOpts, _agent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identityregistry *IdentityregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identityregistry *IdentityregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Identityregistry.Contract.RenounceOwnership(&_Identityregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identityregistry *IdentityregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Identityregistry.Contract.RenounceOwnership(&_Identityregistry.TransactOpts)
}

// SetClaimTopicsRegistry is a paid mutator transaction binding the contract method 0x670af6a9.
//
// Solidity: function setClaimTopicsRegistry(address _claimTopicsRegistry) returns()
func (_Identityregistry *IdentityregistryTransactor) SetClaimTopicsRegistry(opts *bind.TransactOpts, _claimTopicsRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "setClaimTopicsRegistry", _claimTopicsRegistry)
}

// SetClaimTopicsRegistry is a paid mutator transaction binding the contract method 0x670af6a9.
//
// Solidity: function setClaimTopicsRegistry(address _claimTopicsRegistry) returns()
func (_Identityregistry *IdentityregistrySession) SetClaimTopicsRegistry(_claimTopicsRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetClaimTopicsRegistry(&_Identityregistry.TransactOpts, _claimTopicsRegistry)
}

// SetClaimTopicsRegistry is a paid mutator transaction binding the contract method 0x670af6a9.
//
// Solidity: function setClaimTopicsRegistry(address _claimTopicsRegistry) returns()
func (_Identityregistry *IdentityregistryTransactorSession) SetClaimTopicsRegistry(_claimTopicsRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetClaimTopicsRegistry(&_Identityregistry.TransactOpts, _claimTopicsRegistry)
}

// SetIdentityRegistryStorage is a paid mutator transaction binding the contract method 0x26d941ae.
//
// Solidity: function setIdentityRegistryStorage(address _identityRegistryStorage) returns()
func (_Identityregistry *IdentityregistryTransactor) SetIdentityRegistryStorage(opts *bind.TransactOpts, _identityRegistryStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "setIdentityRegistryStorage", _identityRegistryStorage)
}

// SetIdentityRegistryStorage is a paid mutator transaction binding the contract method 0x26d941ae.
//
// Solidity: function setIdentityRegistryStorage(address _identityRegistryStorage) returns()
func (_Identityregistry *IdentityregistrySession) SetIdentityRegistryStorage(_identityRegistryStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetIdentityRegistryStorage(&_Identityregistry.TransactOpts, _identityRegistryStorage)
}

// SetIdentityRegistryStorage is a paid mutator transaction binding the contract method 0x26d941ae.
//
// Solidity: function setIdentityRegistryStorage(address _identityRegistryStorage) returns()
func (_Identityregistry *IdentityregistryTransactorSession) SetIdentityRegistryStorage(_identityRegistryStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetIdentityRegistryStorage(&_Identityregistry.TransactOpts, _identityRegistryStorage)
}

// SetTrustedIssuersRegistry is a paid mutator transaction binding the contract method 0xe744d789.
//
// Solidity: function setTrustedIssuersRegistry(address _trustedIssuersRegistry) returns()
func (_Identityregistry *IdentityregistryTransactor) SetTrustedIssuersRegistry(opts *bind.TransactOpts, _trustedIssuersRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "setTrustedIssuersRegistry", _trustedIssuersRegistry)
}

// SetTrustedIssuersRegistry is a paid mutator transaction binding the contract method 0xe744d789.
//
// Solidity: function setTrustedIssuersRegistry(address _trustedIssuersRegistry) returns()
func (_Identityregistry *IdentityregistrySession) SetTrustedIssuersRegistry(_trustedIssuersRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetTrustedIssuersRegistry(&_Identityregistry.TransactOpts, _trustedIssuersRegistry)
}

// SetTrustedIssuersRegistry is a paid mutator transaction binding the contract method 0xe744d789.
//
// Solidity: function setTrustedIssuersRegistry(address _trustedIssuersRegistry) returns()
func (_Identityregistry *IdentityregistryTransactorSession) SetTrustedIssuersRegistry(_trustedIssuersRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetTrustedIssuersRegistry(&_Identityregistry.TransactOpts, _trustedIssuersRegistry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identityregistry *IdentityregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identityregistry *IdentityregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.TransferOwnership(&_Identityregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identityregistry *IdentityregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.TransferOwnership(&_Identityregistry.TransactOpts, newOwner)
}

// UpdateCountry is a paid mutator transaction binding the contract method 0x3b239a7f.
//
// Solidity: function updateCountry(address _userAddress, uint16 _country) returns()
func (_Identityregistry *IdentityregistryTransactor) UpdateCountry(opts *bind.TransactOpts, _userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "updateCountry", _userAddress, _country)
}

// UpdateCountry is a paid mutator transaction binding the contract method 0x3b239a7f.
//
// Solidity: function updateCountry(address _userAddress, uint16 _country) returns()
func (_Identityregistry *IdentityregistrySession) UpdateCountry(_userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateCountry(&_Identityregistry.TransactOpts, _userAddress, _country)
}

// UpdateCountry is a paid mutator transaction binding the contract method 0x3b239a7f.
//
// Solidity: function updateCountry(address _userAddress, uint16 _country) returns()
func (_Identityregistry *IdentityregistryTransactorSession) UpdateCountry(_userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateCountry(&_Identityregistry.TransactOpts, _userAddress, _country)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x8e098ca1.
//
// Solidity: function updateIdentity(address _userAddress, address _identity) returns()
func (_Identityregistry *IdentityregistryTransactor) UpdateIdentity(opts *bind.TransactOpts, _userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "updateIdentity", _userAddress, _identity)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x8e098ca1.
//
// Solidity: function updateIdentity(address _userAddress, address _identity) returns()
func (_Identityregistry *IdentityregistrySession) UpdateIdentity(_userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x8e098ca1.
//
// Solidity: function updateIdentity(address _userAddress, address _identity) returns()
func (_Identityregistry *IdentityregistryTransactorSession) UpdateIdentity(_userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity)
}

// IdentityregistryAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the Identityregistry contract.
type IdentityregistryAgentAddedIterator struct {
	Event *IdentityregistryAgentAdded // Event containing the contract specifics and raw log

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
func (it *IdentityregistryAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryAgentAdded)
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
		it.Event = new(IdentityregistryAgentAdded)
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
func (it *IdentityregistryAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryAgentAdded represents a AgentAdded event raised by the Identityregistry contract.
type IdentityregistryAgentAdded struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Identityregistry *IdentityregistryFilterer) FilterAgentAdded(opts *bind.FilterOpts, _agent []common.Address) (*IdentityregistryAgentAddedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryAgentAddedIterator{contract: _Identityregistry.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Identityregistry *IdentityregistryFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *IdentityregistryAgentAdded, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryAgentAdded)
				if err := _Identityregistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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
func (_Identityregistry *IdentityregistryFilterer) ParseAgentAdded(log types.Log) (*IdentityregistryAgentAdded, error) {
	event := new(IdentityregistryAgentAdded)
	if err := _Identityregistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the Identityregistry contract.
type IdentityregistryAgentRemovedIterator struct {
	Event *IdentityregistryAgentRemoved // Event containing the contract specifics and raw log

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
func (it *IdentityregistryAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryAgentRemoved)
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
		it.Event = new(IdentityregistryAgentRemoved)
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
func (it *IdentityregistryAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryAgentRemoved represents a AgentRemoved event raised by the Identityregistry contract.
type IdentityregistryAgentRemoved struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Identityregistry *IdentityregistryFilterer) FilterAgentRemoved(opts *bind.FilterOpts, _agent []common.Address) (*IdentityregistryAgentRemovedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryAgentRemovedIterator{contract: _Identityregistry.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Identityregistry *IdentityregistryFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *IdentityregistryAgentRemoved, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryAgentRemoved)
				if err := _Identityregistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_Identityregistry *IdentityregistryFilterer) ParseAgentRemoved(log types.Log) (*IdentityregistryAgentRemoved, error) {
	event := new(IdentityregistryAgentRemoved)
	if err := _Identityregistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryClaimTopicsRegistrySetIterator is returned from FilterClaimTopicsRegistrySet and is used to iterate over the raw logs and unpacked data for ClaimTopicsRegistrySet events raised by the Identityregistry contract.
type IdentityregistryClaimTopicsRegistrySetIterator struct {
	Event *IdentityregistryClaimTopicsRegistrySet // Event containing the contract specifics and raw log

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
func (it *IdentityregistryClaimTopicsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryClaimTopicsRegistrySet)
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
		it.Event = new(IdentityregistryClaimTopicsRegistrySet)
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
func (it *IdentityregistryClaimTopicsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryClaimTopicsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryClaimTopicsRegistrySet represents a ClaimTopicsRegistrySet event raised by the Identityregistry contract.
type IdentityregistryClaimTopicsRegistrySet struct {
	ClaimTopicsRegistry common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterClaimTopicsRegistrySet is a free log retrieval operation binding the contract event 0x7170bf15b246e880b2369cd7c67d057760d8a35149e8c64dde91efa22bcc76d0.
//
// Solidity: event ClaimTopicsRegistrySet(address indexed claimTopicsRegistry)
func (_Identityregistry *IdentityregistryFilterer) FilterClaimTopicsRegistrySet(opts *bind.FilterOpts, claimTopicsRegistry []common.Address) (*IdentityregistryClaimTopicsRegistrySetIterator, error) {

	var claimTopicsRegistryRule []interface{}
	for _, claimTopicsRegistryItem := range claimTopicsRegistry {
		claimTopicsRegistryRule = append(claimTopicsRegistryRule, claimTopicsRegistryItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "ClaimTopicsRegistrySet", claimTopicsRegistryRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryClaimTopicsRegistrySetIterator{contract: _Identityregistry.contract, event: "ClaimTopicsRegistrySet", logs: logs, sub: sub}, nil
}

// WatchClaimTopicsRegistrySet is a free log subscription operation binding the contract event 0x7170bf15b246e880b2369cd7c67d057760d8a35149e8c64dde91efa22bcc76d0.
//
// Solidity: event ClaimTopicsRegistrySet(address indexed claimTopicsRegistry)
func (_Identityregistry *IdentityregistryFilterer) WatchClaimTopicsRegistrySet(opts *bind.WatchOpts, sink chan<- *IdentityregistryClaimTopicsRegistrySet, claimTopicsRegistry []common.Address) (event.Subscription, error) {

	var claimTopicsRegistryRule []interface{}
	for _, claimTopicsRegistryItem := range claimTopicsRegistry {
		claimTopicsRegistryRule = append(claimTopicsRegistryRule, claimTopicsRegistryItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "ClaimTopicsRegistrySet", claimTopicsRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryClaimTopicsRegistrySet)
				if err := _Identityregistry.contract.UnpackLog(event, "ClaimTopicsRegistrySet", log); err != nil {
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

// ParseClaimTopicsRegistrySet is a log parse operation binding the contract event 0x7170bf15b246e880b2369cd7c67d057760d8a35149e8c64dde91efa22bcc76d0.
//
// Solidity: event ClaimTopicsRegistrySet(address indexed claimTopicsRegistry)
func (_Identityregistry *IdentityregistryFilterer) ParseClaimTopicsRegistrySet(log types.Log) (*IdentityregistryClaimTopicsRegistrySet, error) {
	event := new(IdentityregistryClaimTopicsRegistrySet)
	if err := _Identityregistry.contract.UnpackLog(event, "ClaimTopicsRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryCountryUpdatedIterator is returned from FilterCountryUpdated and is used to iterate over the raw logs and unpacked data for CountryUpdated events raised by the Identityregistry contract.
type IdentityregistryCountryUpdatedIterator struct {
	Event *IdentityregistryCountryUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityregistryCountryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryCountryUpdated)
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
		it.Event = new(IdentityregistryCountryUpdated)
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
func (it *IdentityregistryCountryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryCountryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryCountryUpdated represents a CountryUpdated event raised by the Identityregistry contract.
type IdentityregistryCountryUpdated struct {
	InvestorAddress common.Address
	Country         uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCountryUpdated is a free log retrieval operation binding the contract event 0x04ed3b726495c2dca1ff1215d9ca54e1a4030abb5e82b0f6ce55702416cee853.
//
// Solidity: event CountryUpdated(address indexed investorAddress, uint16 indexed country)
func (_Identityregistry *IdentityregistryFilterer) FilterCountryUpdated(opts *bind.FilterOpts, investorAddress []common.Address, country []uint16) (*IdentityregistryCountryUpdatedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var countryRule []interface{}
	for _, countryItem := range country {
		countryRule = append(countryRule, countryItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "CountryUpdated", investorAddressRule, countryRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryCountryUpdatedIterator{contract: _Identityregistry.contract, event: "CountryUpdated", logs: logs, sub: sub}, nil
}

// WatchCountryUpdated is a free log subscription operation binding the contract event 0x04ed3b726495c2dca1ff1215d9ca54e1a4030abb5e82b0f6ce55702416cee853.
//
// Solidity: event CountryUpdated(address indexed investorAddress, uint16 indexed country)
func (_Identityregistry *IdentityregistryFilterer) WatchCountryUpdated(opts *bind.WatchOpts, sink chan<- *IdentityregistryCountryUpdated, investorAddress []common.Address, country []uint16) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var countryRule []interface{}
	for _, countryItem := range country {
		countryRule = append(countryRule, countryItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "CountryUpdated", investorAddressRule, countryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryCountryUpdated)
				if err := _Identityregistry.contract.UnpackLog(event, "CountryUpdated", log); err != nil {
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

// ParseCountryUpdated is a log parse operation binding the contract event 0x04ed3b726495c2dca1ff1215d9ca54e1a4030abb5e82b0f6ce55702416cee853.
//
// Solidity: event CountryUpdated(address indexed investorAddress, uint16 indexed country)
func (_Identityregistry *IdentityregistryFilterer) ParseCountryUpdated(log types.Log) (*IdentityregistryCountryUpdated, error) {
	event := new(IdentityregistryCountryUpdated)
	if err := _Identityregistry.contract.UnpackLog(event, "CountryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryIdentityRegisteredIterator is returned from FilterIdentityRegistered and is used to iterate over the raw logs and unpacked data for IdentityRegistered events raised by the Identityregistry contract.
type IdentityregistryIdentityRegisteredIterator struct {
	Event *IdentityregistryIdentityRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityregistryIdentityRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryIdentityRegistered)
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
		it.Event = new(IdentityregistryIdentityRegistered)
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
func (it *IdentityregistryIdentityRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryIdentityRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryIdentityRegistered represents a IdentityRegistered event raised by the Identityregistry contract.
type IdentityregistryIdentityRegistered struct {
	InvestorAddress common.Address
	Identity        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityRegistered is a free log retrieval operation binding the contract event 0x6ae73635c50d24a45af6fbd5e016ac4bed179addbc8bf24e04ff0fcc6d33af19.
//
// Solidity: event IdentityRegistered(address indexed investorAddress, address indexed identity)
func (_Identityregistry *IdentityregistryFilterer) FilterIdentityRegistered(opts *bind.FilterOpts, investorAddress []common.Address, identity []common.Address) (*IdentityregistryIdentityRegisteredIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "IdentityRegistered", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryIdentityRegisteredIterator{contract: _Identityregistry.contract, event: "IdentityRegistered", logs: logs, sub: sub}, nil
}

// WatchIdentityRegistered is a free log subscription operation binding the contract event 0x6ae73635c50d24a45af6fbd5e016ac4bed179addbc8bf24e04ff0fcc6d33af19.
//
// Solidity: event IdentityRegistered(address indexed investorAddress, address indexed identity)
func (_Identityregistry *IdentityregistryFilterer) WatchIdentityRegistered(opts *bind.WatchOpts, sink chan<- *IdentityregistryIdentityRegistered, investorAddress []common.Address, identity []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "IdentityRegistered", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryIdentityRegistered)
				if err := _Identityregistry.contract.UnpackLog(event, "IdentityRegistered", log); err != nil {
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

// ParseIdentityRegistered is a log parse operation binding the contract event 0x6ae73635c50d24a45af6fbd5e016ac4bed179addbc8bf24e04ff0fcc6d33af19.
//
// Solidity: event IdentityRegistered(address indexed investorAddress, address indexed identity)
func (_Identityregistry *IdentityregistryFilterer) ParseIdentityRegistered(log types.Log) (*IdentityregistryIdentityRegistered, error) {
	event := new(IdentityregistryIdentityRegistered)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryIdentityRemovedIterator is returned from FilterIdentityRemoved and is used to iterate over the raw logs and unpacked data for IdentityRemoved events raised by the Identityregistry contract.
type IdentityregistryIdentityRemovedIterator struct {
	Event *IdentityregistryIdentityRemoved // Event containing the contract specifics and raw log

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
func (it *IdentityregistryIdentityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryIdentityRemoved)
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
		it.Event = new(IdentityregistryIdentityRemoved)
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
func (it *IdentityregistryIdentityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryIdentityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryIdentityRemoved represents a IdentityRemoved event raised by the Identityregistry contract.
type IdentityregistryIdentityRemoved struct {
	InvestorAddress common.Address
	Identity        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityRemoved is a free log retrieval operation binding the contract event 0x59d6590e225b81befe259af056324092801080acbb7feab310eb34678871f327.
//
// Solidity: event IdentityRemoved(address indexed investorAddress, address indexed identity)
func (_Identityregistry *IdentityregistryFilterer) FilterIdentityRemoved(opts *bind.FilterOpts, investorAddress []common.Address, identity []common.Address) (*IdentityregistryIdentityRemovedIterator, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "IdentityRemoved", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryIdentityRemovedIterator{contract: _Identityregistry.contract, event: "IdentityRemoved", logs: logs, sub: sub}, nil
}

// WatchIdentityRemoved is a free log subscription operation binding the contract event 0x59d6590e225b81befe259af056324092801080acbb7feab310eb34678871f327.
//
// Solidity: event IdentityRemoved(address indexed investorAddress, address indexed identity)
func (_Identityregistry *IdentityregistryFilterer) WatchIdentityRemoved(opts *bind.WatchOpts, sink chan<- *IdentityregistryIdentityRemoved, investorAddress []common.Address, identity []common.Address) (event.Subscription, error) {

	var investorAddressRule []interface{}
	for _, investorAddressItem := range investorAddress {
		investorAddressRule = append(investorAddressRule, investorAddressItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "IdentityRemoved", investorAddressRule, identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryIdentityRemoved)
				if err := _Identityregistry.contract.UnpackLog(event, "IdentityRemoved", log); err != nil {
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

// ParseIdentityRemoved is a log parse operation binding the contract event 0x59d6590e225b81befe259af056324092801080acbb7feab310eb34678871f327.
//
// Solidity: event IdentityRemoved(address indexed investorAddress, address indexed identity)
func (_Identityregistry *IdentityregistryFilterer) ParseIdentityRemoved(log types.Log) (*IdentityregistryIdentityRemoved, error) {
	event := new(IdentityregistryIdentityRemoved)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryIdentityStorageSetIterator is returned from FilterIdentityStorageSet and is used to iterate over the raw logs and unpacked data for IdentityStorageSet events raised by the Identityregistry contract.
type IdentityregistryIdentityStorageSetIterator struct {
	Event *IdentityregistryIdentityStorageSet // Event containing the contract specifics and raw log

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
func (it *IdentityregistryIdentityStorageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryIdentityStorageSet)
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
		it.Event = new(IdentityregistryIdentityStorageSet)
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
func (it *IdentityregistryIdentityStorageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryIdentityStorageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryIdentityStorageSet represents a IdentityStorageSet event raised by the Identityregistry contract.
type IdentityregistryIdentityStorageSet struct {
	IdentityStorage common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityStorageSet is a free log retrieval operation binding the contract event 0x2fa8b95c1db7afe99e3398f3792f008135cedc1fa26b0bb2ecd2352cd166d53c.
//
// Solidity: event IdentityStorageSet(address indexed identityStorage)
func (_Identityregistry *IdentityregistryFilterer) FilterIdentityStorageSet(opts *bind.FilterOpts, identityStorage []common.Address) (*IdentityregistryIdentityStorageSetIterator, error) {

	var identityStorageRule []interface{}
	for _, identityStorageItem := range identityStorage {
		identityStorageRule = append(identityStorageRule, identityStorageItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "IdentityStorageSet", identityStorageRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryIdentityStorageSetIterator{contract: _Identityregistry.contract, event: "IdentityStorageSet", logs: logs, sub: sub}, nil
}

// WatchIdentityStorageSet is a free log subscription operation binding the contract event 0x2fa8b95c1db7afe99e3398f3792f008135cedc1fa26b0bb2ecd2352cd166d53c.
//
// Solidity: event IdentityStorageSet(address indexed identityStorage)
func (_Identityregistry *IdentityregistryFilterer) WatchIdentityStorageSet(opts *bind.WatchOpts, sink chan<- *IdentityregistryIdentityStorageSet, identityStorage []common.Address) (event.Subscription, error) {

	var identityStorageRule []interface{}
	for _, identityStorageItem := range identityStorage {
		identityStorageRule = append(identityStorageRule, identityStorageItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "IdentityStorageSet", identityStorageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryIdentityStorageSet)
				if err := _Identityregistry.contract.UnpackLog(event, "IdentityStorageSet", log); err != nil {
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

// ParseIdentityStorageSet is a log parse operation binding the contract event 0x2fa8b95c1db7afe99e3398f3792f008135cedc1fa26b0bb2ecd2352cd166d53c.
//
// Solidity: event IdentityStorageSet(address indexed identityStorage)
func (_Identityregistry *IdentityregistryFilterer) ParseIdentityStorageSet(log types.Log) (*IdentityregistryIdentityStorageSet, error) {
	event := new(IdentityregistryIdentityStorageSet)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityStorageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryIdentityUpdatedIterator is returned from FilterIdentityUpdated and is used to iterate over the raw logs and unpacked data for IdentityUpdated events raised by the Identityregistry contract.
type IdentityregistryIdentityUpdatedIterator struct {
	Event *IdentityregistryIdentityUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityregistryIdentityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryIdentityUpdated)
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
		it.Event = new(IdentityregistryIdentityUpdated)
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
func (it *IdentityregistryIdentityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryIdentityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryIdentityUpdated represents a IdentityUpdated event raised by the Identityregistry contract.
type IdentityregistryIdentityUpdated struct {
	OldIdentity common.Address
	NewIdentity common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIdentityUpdated is a free log retrieval operation binding the contract event 0xe98082932c8056a0f514da9104e4a66bc2cbaef102ad59d90c4b24220ebf6010.
//
// Solidity: event IdentityUpdated(address indexed oldIdentity, address indexed newIdentity)
func (_Identityregistry *IdentityregistryFilterer) FilterIdentityUpdated(opts *bind.FilterOpts, oldIdentity []common.Address, newIdentity []common.Address) (*IdentityregistryIdentityUpdatedIterator, error) {

	var oldIdentityRule []interface{}
	for _, oldIdentityItem := range oldIdentity {
		oldIdentityRule = append(oldIdentityRule, oldIdentityItem)
	}
	var newIdentityRule []interface{}
	for _, newIdentityItem := range newIdentity {
		newIdentityRule = append(newIdentityRule, newIdentityItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "IdentityUpdated", oldIdentityRule, newIdentityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryIdentityUpdatedIterator{contract: _Identityregistry.contract, event: "IdentityUpdated", logs: logs, sub: sub}, nil
}

// WatchIdentityUpdated is a free log subscription operation binding the contract event 0xe98082932c8056a0f514da9104e4a66bc2cbaef102ad59d90c4b24220ebf6010.
//
// Solidity: event IdentityUpdated(address indexed oldIdentity, address indexed newIdentity)
func (_Identityregistry *IdentityregistryFilterer) WatchIdentityUpdated(opts *bind.WatchOpts, sink chan<- *IdentityregistryIdentityUpdated, oldIdentity []common.Address, newIdentity []common.Address) (event.Subscription, error) {

	var oldIdentityRule []interface{}
	for _, oldIdentityItem := range oldIdentity {
		oldIdentityRule = append(oldIdentityRule, oldIdentityItem)
	}
	var newIdentityRule []interface{}
	for _, newIdentityItem := range newIdentity {
		newIdentityRule = append(newIdentityRule, newIdentityItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "IdentityUpdated", oldIdentityRule, newIdentityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryIdentityUpdated)
				if err := _Identityregistry.contract.UnpackLog(event, "IdentityUpdated", log); err != nil {
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

// ParseIdentityUpdated is a log parse operation binding the contract event 0xe98082932c8056a0f514da9104e4a66bc2cbaef102ad59d90c4b24220ebf6010.
//
// Solidity: event IdentityUpdated(address indexed oldIdentity, address indexed newIdentity)
func (_Identityregistry *IdentityregistryFilterer) ParseIdentityUpdated(log types.Log) (*IdentityregistryIdentityUpdated, error) {
	event := new(IdentityregistryIdentityUpdated)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Identityregistry contract.
type IdentityregistryInitializedIterator struct {
	Event *IdentityregistryInitialized // Event containing the contract specifics and raw log

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
func (it *IdentityregistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryInitialized)
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
		it.Event = new(IdentityregistryInitialized)
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
func (it *IdentityregistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryInitialized represents a Initialized event raised by the Identityregistry contract.
type IdentityregistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identityregistry *IdentityregistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityregistryInitializedIterator, error) {

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityregistryInitializedIterator{contract: _Identityregistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identityregistry *IdentityregistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentityregistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryInitialized)
				if err := _Identityregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Identityregistry *IdentityregistryFilterer) ParseInitialized(log types.Log) (*IdentityregistryInitialized, error) {
	event := new(IdentityregistryInitialized)
	if err := _Identityregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Identityregistry contract.
type IdentityregistryOwnershipTransferredIterator struct {
	Event *IdentityregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryOwnershipTransferred)
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
		it.Event = new(IdentityregistryOwnershipTransferred)
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
func (it *IdentityregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Identityregistry contract.
type IdentityregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identityregistry *IdentityregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryOwnershipTransferredIterator{contract: _Identityregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identityregistry *IdentityregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryOwnershipTransferred)
				if err := _Identityregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Identityregistry *IdentityregistryFilterer) ParseOwnershipTransferred(log types.Log) (*IdentityregistryOwnershipTransferred, error) {
	event := new(IdentityregistryOwnershipTransferred)
	if err := _Identityregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityregistryTrustedIssuersRegistrySetIterator is returned from FilterTrustedIssuersRegistrySet and is used to iterate over the raw logs and unpacked data for TrustedIssuersRegistrySet events raised by the Identityregistry contract.
type IdentityregistryTrustedIssuersRegistrySetIterator struct {
	Event *IdentityregistryTrustedIssuersRegistrySet // Event containing the contract specifics and raw log

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
func (it *IdentityregistryTrustedIssuersRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityregistryTrustedIssuersRegistrySet)
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
		it.Event = new(IdentityregistryTrustedIssuersRegistrySet)
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
func (it *IdentityregistryTrustedIssuersRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityregistryTrustedIssuersRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityregistryTrustedIssuersRegistrySet represents a TrustedIssuersRegistrySet event raised by the Identityregistry contract.
type IdentityregistryTrustedIssuersRegistrySet struct {
	TrustedIssuersRegistry common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterTrustedIssuersRegistrySet is a free log retrieval operation binding the contract event 0x1b98cb79e6f73020175fe87333f1b91ad6a881519c0afe30340c2599b2b4bde0.
//
// Solidity: event TrustedIssuersRegistrySet(address indexed trustedIssuersRegistry)
func (_Identityregistry *IdentityregistryFilterer) FilterTrustedIssuersRegistrySet(opts *bind.FilterOpts, trustedIssuersRegistry []common.Address) (*IdentityregistryTrustedIssuersRegistrySetIterator, error) {

	var trustedIssuersRegistryRule []interface{}
	for _, trustedIssuersRegistryItem := range trustedIssuersRegistry {
		trustedIssuersRegistryRule = append(trustedIssuersRegistryRule, trustedIssuersRegistryItem)
	}

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "TrustedIssuersRegistrySet", trustedIssuersRegistryRule)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryTrustedIssuersRegistrySetIterator{contract: _Identityregistry.contract, event: "TrustedIssuersRegistrySet", logs: logs, sub: sub}, nil
}

// WatchTrustedIssuersRegistrySet is a free log subscription operation binding the contract event 0x1b98cb79e6f73020175fe87333f1b91ad6a881519c0afe30340c2599b2b4bde0.
//
// Solidity: event TrustedIssuersRegistrySet(address indexed trustedIssuersRegistry)
func (_Identityregistry *IdentityregistryFilterer) WatchTrustedIssuersRegistrySet(opts *bind.WatchOpts, sink chan<- *IdentityregistryTrustedIssuersRegistrySet, trustedIssuersRegistry []common.Address) (event.Subscription, error) {

	var trustedIssuersRegistryRule []interface{}
	for _, trustedIssuersRegistryItem := range trustedIssuersRegistry {
		trustedIssuersRegistryRule = append(trustedIssuersRegistryRule, trustedIssuersRegistryItem)
	}

	logs, sub, err := _Identityregistry.contract.WatchLogs(opts, "TrustedIssuersRegistrySet", trustedIssuersRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityregistryTrustedIssuersRegistrySet)
				if err := _Identityregistry.contract.UnpackLog(event, "TrustedIssuersRegistrySet", log); err != nil {
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

// ParseTrustedIssuersRegistrySet is a log parse operation binding the contract event 0x1b98cb79e6f73020175fe87333f1b91ad6a881519c0afe30340c2599b2b4bde0.
//
// Solidity: event TrustedIssuersRegistrySet(address indexed trustedIssuersRegistry)
func (_Identityregistry *IdentityregistryFilterer) ParseTrustedIssuersRegistrySet(log types.Log) (*IdentityregistryTrustedIssuersRegistrySet, error) {
	event := new(IdentityregistryTrustedIssuersRegistrySet)
	if err := _Identityregistry.contract.UnpackLog(event, "TrustedIssuersRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
