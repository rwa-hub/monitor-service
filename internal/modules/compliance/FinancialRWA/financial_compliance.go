// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package financialcompliance

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

// FinancialcomplianceMetaData contains all meta data concerning the Financialcompliance contract.
var FinancialcomplianceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"approveBuyer\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditInsuranceApproved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"serasaClearance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"documentsVerified\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"income\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressVerified\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"saleRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signedAgreement\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bindCompliance\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canComplianceBind\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_minIncomeRequired\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isComplianceBound\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPlugAndPlay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"minIncomeRequired\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minTransactionValue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"moduleBurnAction\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"moduleCheck\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"moduleMintAction\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"moduleTransferAction\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbindCompliance\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BuyerApproved\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceBound\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceCheckFailedEvent\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceCheckPassed\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceUnbound\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ComplianceCheckFailed\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"IncomeTooLow\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotApprovedBuyer\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// FinancialcomplianceABI is the input ABI used to generate the binding from.
// Deprecated: Use FinancialcomplianceMetaData.ABI instead.
var FinancialcomplianceABI = FinancialcomplianceMetaData.ABI

// Financialcompliance is an auto generated Go binding around an Ethereum contract.
type Financialcompliance struct {
	FinancialcomplianceCaller     // Read-only binding to the contract
	FinancialcomplianceTransactor // Write-only binding to the contract
	FinancialcomplianceFilterer   // Log filterer for contract events
}

// FinancialcomplianceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinancialcomplianceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinancialcomplianceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinancialcomplianceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinancialcomplianceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinancialcomplianceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinancialcomplianceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinancialcomplianceSession struct {
	Contract     *Financialcompliance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FinancialcomplianceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinancialcomplianceCallerSession struct {
	Contract *FinancialcomplianceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FinancialcomplianceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinancialcomplianceTransactorSession struct {
	Contract     *FinancialcomplianceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FinancialcomplianceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinancialcomplianceRaw struct {
	Contract *Financialcompliance // Generic contract binding to access the raw methods on
}

// FinancialcomplianceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinancialcomplianceCallerRaw struct {
	Contract *FinancialcomplianceCaller // Generic read-only contract binding to access the raw methods on
}

// FinancialcomplianceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinancialcomplianceTransactorRaw struct {
	Contract *FinancialcomplianceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinancialcompliance creates a new instance of Financialcompliance, bound to a specific deployed contract.
func NewFinancialcompliance(address common.Address, backend bind.ContractBackend) (*Financialcompliance, error) {
	contract, err := bindFinancialcompliance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Financialcompliance{FinancialcomplianceCaller: FinancialcomplianceCaller{contract: contract}, FinancialcomplianceTransactor: FinancialcomplianceTransactor{contract: contract}, FinancialcomplianceFilterer: FinancialcomplianceFilterer{contract: contract}}, nil
}

// NewFinancialcomplianceCaller creates a new read-only instance of Financialcompliance, bound to a specific deployed contract.
func NewFinancialcomplianceCaller(address common.Address, caller bind.ContractCaller) (*FinancialcomplianceCaller, error) {
	contract, err := bindFinancialcompliance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceCaller{contract: contract}, nil
}

// NewFinancialcomplianceTransactor creates a new write-only instance of Financialcompliance, bound to a specific deployed contract.
func NewFinancialcomplianceTransactor(address common.Address, transactor bind.ContractTransactor) (*FinancialcomplianceTransactor, error) {
	contract, err := bindFinancialcompliance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceTransactor{contract: contract}, nil
}

// NewFinancialcomplianceFilterer creates a new log filterer instance of Financialcompliance, bound to a specific deployed contract.
func NewFinancialcomplianceFilterer(address common.Address, filterer bind.ContractFilterer) (*FinancialcomplianceFilterer, error) {
	contract, err := bindFinancialcompliance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceFilterer{contract: contract}, nil
}

// bindFinancialcompliance binds a generic wrapper to an already deployed contract.
func bindFinancialcompliance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinancialcomplianceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Financialcompliance *FinancialcomplianceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Financialcompliance.Contract.FinancialcomplianceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Financialcompliance *FinancialcomplianceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Financialcompliance.Contract.FinancialcomplianceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Financialcompliance *FinancialcomplianceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Financialcompliance.Contract.FinancialcomplianceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Financialcompliance *FinancialcomplianceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Financialcompliance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Financialcompliance *FinancialcomplianceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Financialcompliance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Financialcompliance *FinancialcomplianceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Financialcompliance.Contract.contract.Transact(opts, method, params...)
}

// CanComplianceBind is a free data retrieval call binding the contract method 0xbcc21053.
//
// Solidity: function canComplianceBind(address ) pure returns(bool)
func (_Financialcompliance *FinancialcomplianceCaller) CanComplianceBind(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "canComplianceBind", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanComplianceBind is a free data retrieval call binding the contract method 0xbcc21053.
//
// Solidity: function canComplianceBind(address ) pure returns(bool)
func (_Financialcompliance *FinancialcomplianceSession) CanComplianceBind(arg0 common.Address) (bool, error) {
	return _Financialcompliance.Contract.CanComplianceBind(&_Financialcompliance.CallOpts, arg0)
}

// CanComplianceBind is a free data retrieval call binding the contract method 0xbcc21053.
//
// Solidity: function canComplianceBind(address ) pure returns(bool)
func (_Financialcompliance *FinancialcomplianceCallerSession) CanComplianceBind(arg0 common.Address) (bool, error) {
	return _Financialcompliance.Contract.CanComplianceBind(&_Financialcompliance.CallOpts, arg0)
}

// IsComplianceBound is a free data retrieval call binding the contract method 0x4cf4d295.
//
// Solidity: function isComplianceBound(address _compliance) view returns(bool)
func (_Financialcompliance *FinancialcomplianceCaller) IsComplianceBound(opts *bind.CallOpts, _compliance common.Address) (bool, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "isComplianceBound", _compliance)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComplianceBound is a free data retrieval call binding the contract method 0x4cf4d295.
//
// Solidity: function isComplianceBound(address _compliance) view returns(bool)
func (_Financialcompliance *FinancialcomplianceSession) IsComplianceBound(_compliance common.Address) (bool, error) {
	return _Financialcompliance.Contract.IsComplianceBound(&_Financialcompliance.CallOpts, _compliance)
}

// IsComplianceBound is a free data retrieval call binding the contract method 0x4cf4d295.
//
// Solidity: function isComplianceBound(address _compliance) view returns(bool)
func (_Financialcompliance *FinancialcomplianceCallerSession) IsComplianceBound(_compliance common.Address) (bool, error) {
	return _Financialcompliance.Contract.IsComplianceBound(&_Financialcompliance.CallOpts, _compliance)
}

// IsPlugAndPlay is a free data retrieval call binding the contract method 0xe6f5e807.
//
// Solidity: function isPlugAndPlay() pure returns(bool)
func (_Financialcompliance *FinancialcomplianceCaller) IsPlugAndPlay(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "isPlugAndPlay")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPlugAndPlay is a free data retrieval call binding the contract method 0xe6f5e807.
//
// Solidity: function isPlugAndPlay() pure returns(bool)
func (_Financialcompliance *FinancialcomplianceSession) IsPlugAndPlay() (bool, error) {
	return _Financialcompliance.Contract.IsPlugAndPlay(&_Financialcompliance.CallOpts)
}

// IsPlugAndPlay is a free data retrieval call binding the contract method 0xe6f5e807.
//
// Solidity: function isPlugAndPlay() pure returns(bool)
func (_Financialcompliance *FinancialcomplianceCallerSession) IsPlugAndPlay() (bool, error) {
	return _Financialcompliance.Contract.IsPlugAndPlay(&_Financialcompliance.CallOpts)
}

// MinIncomeRequired is a free data retrieval call binding the contract method 0x8515efe5.
//
// Solidity: function minIncomeRequired() view returns(uint256)
func (_Financialcompliance *FinancialcomplianceCaller) MinIncomeRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "minIncomeRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinIncomeRequired is a free data retrieval call binding the contract method 0x8515efe5.
//
// Solidity: function minIncomeRequired() view returns(uint256)
func (_Financialcompliance *FinancialcomplianceSession) MinIncomeRequired() (*big.Int, error) {
	return _Financialcompliance.Contract.MinIncomeRequired(&_Financialcompliance.CallOpts)
}

// MinIncomeRequired is a free data retrieval call binding the contract method 0x8515efe5.
//
// Solidity: function minIncomeRequired() view returns(uint256)
func (_Financialcompliance *FinancialcomplianceCallerSession) MinIncomeRequired() (*big.Int, error) {
	return _Financialcompliance.Contract.MinIncomeRequired(&_Financialcompliance.CallOpts)
}

// MinTransactionValue is a free data retrieval call binding the contract method 0x97ca05e4.
//
// Solidity: function minTransactionValue() view returns(uint256)
func (_Financialcompliance *FinancialcomplianceCaller) MinTransactionValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "minTransactionValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransactionValue is a free data retrieval call binding the contract method 0x97ca05e4.
//
// Solidity: function minTransactionValue() view returns(uint256)
func (_Financialcompliance *FinancialcomplianceSession) MinTransactionValue() (*big.Int, error) {
	return _Financialcompliance.Contract.MinTransactionValue(&_Financialcompliance.CallOpts)
}

// MinTransactionValue is a free data retrieval call binding the contract method 0x97ca05e4.
//
// Solidity: function minTransactionValue() view returns(uint256)
func (_Financialcompliance *FinancialcomplianceCallerSession) MinTransactionValue() (*big.Int, error) {
	return _Financialcompliance.Contract.MinTransactionValue(&_Financialcompliance.CallOpts)
}

// ModuleCheck is a free data retrieval call binding the contract method 0x013b7ce4.
//
// Solidity: function moduleCheck(address from, address to, uint256 value, address compliance) view returns(bool)
func (_Financialcompliance *FinancialcomplianceCaller) ModuleCheck(opts *bind.CallOpts, from common.Address, to common.Address, value *big.Int, compliance common.Address) (bool, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "moduleCheck", from, to, value, compliance)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ModuleCheck is a free data retrieval call binding the contract method 0x013b7ce4.
//
// Solidity: function moduleCheck(address from, address to, uint256 value, address compliance) view returns(bool)
func (_Financialcompliance *FinancialcomplianceSession) ModuleCheck(from common.Address, to common.Address, value *big.Int, compliance common.Address) (bool, error) {
	return _Financialcompliance.Contract.ModuleCheck(&_Financialcompliance.CallOpts, from, to, value, compliance)
}

// ModuleCheck is a free data retrieval call binding the contract method 0x013b7ce4.
//
// Solidity: function moduleCheck(address from, address to, uint256 value, address compliance) view returns(bool)
func (_Financialcompliance *FinancialcomplianceCallerSession) ModuleCheck(from common.Address, to common.Address, value *big.Int, compliance common.Address) (bool, error) {
	return _Financialcompliance.Contract.ModuleCheck(&_Financialcompliance.CallOpts, from, to, value, compliance)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Financialcompliance *FinancialcomplianceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Financialcompliance *FinancialcomplianceSession) Name() (string, error) {
	return _Financialcompliance.Contract.Name(&_Financialcompliance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Financialcompliance *FinancialcomplianceCallerSession) Name() (string, error) {
	return _Financialcompliance.Contract.Name(&_Financialcompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Financialcompliance *FinancialcomplianceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Financialcompliance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Financialcompliance *FinancialcomplianceSession) Owner() (common.Address, error) {
	return _Financialcompliance.Contract.Owner(&_Financialcompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Financialcompliance *FinancialcomplianceCallerSession) Owner() (common.Address, error) {
	return _Financialcompliance.Contract.Owner(&_Financialcompliance.CallOpts)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0x5cdeccf8.
//
// Solidity: function approveBuyer(address buyer, bool creditInsuranceApproved, bool serasaClearance, bool documentsVerified, uint256 income, string addressVerified, bool saleRegistered, bool signedAgreement) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) ApproveBuyer(opts *bind.TransactOpts, buyer common.Address, creditInsuranceApproved bool, serasaClearance bool, documentsVerified bool, income *big.Int, addressVerified string, saleRegistered bool, signedAgreement bool) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "approveBuyer", buyer, creditInsuranceApproved, serasaClearance, documentsVerified, income, addressVerified, saleRegistered, signedAgreement)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0x5cdeccf8.
//
// Solidity: function approveBuyer(address buyer, bool creditInsuranceApproved, bool serasaClearance, bool documentsVerified, uint256 income, string addressVerified, bool saleRegistered, bool signedAgreement) returns()
func (_Financialcompliance *FinancialcomplianceSession) ApproveBuyer(buyer common.Address, creditInsuranceApproved bool, serasaClearance bool, documentsVerified bool, income *big.Int, addressVerified string, saleRegistered bool, signedAgreement bool) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ApproveBuyer(&_Financialcompliance.TransactOpts, buyer, creditInsuranceApproved, serasaClearance, documentsVerified, income, addressVerified, saleRegistered, signedAgreement)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0x5cdeccf8.
//
// Solidity: function approveBuyer(address buyer, bool creditInsuranceApproved, bool serasaClearance, bool documentsVerified, uint256 income, string addressVerified, bool saleRegistered, bool signedAgreement) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) ApproveBuyer(buyer common.Address, creditInsuranceApproved bool, serasaClearance bool, documentsVerified bool, income *big.Int, addressVerified string, saleRegistered bool, signedAgreement bool) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ApproveBuyer(&_Financialcompliance.TransactOpts, buyer, creditInsuranceApproved, serasaClearance, documentsVerified, income, addressVerified, saleRegistered, signedAgreement)
}

// BindCompliance is a paid mutator transaction binding the contract method 0x4a932544.
//
// Solidity: function bindCompliance(address _compliance) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) BindCompliance(opts *bind.TransactOpts, _compliance common.Address) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "bindCompliance", _compliance)
}

// BindCompliance is a paid mutator transaction binding the contract method 0x4a932544.
//
// Solidity: function bindCompliance(address _compliance) returns()
func (_Financialcompliance *FinancialcomplianceSession) BindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Financialcompliance.Contract.BindCompliance(&_Financialcompliance.TransactOpts, _compliance)
}

// BindCompliance is a paid mutator transaction binding the contract method 0x4a932544.
//
// Solidity: function bindCompliance(address _compliance) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) BindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Financialcompliance.Contract.BindCompliance(&_Financialcompliance.TransactOpts, _compliance)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(uint256 _minIncomeRequired) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) Init(opts *bind.TransactOpts, _minIncomeRequired *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "init", _minIncomeRequired)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(uint256 _minIncomeRequired) returns()
func (_Financialcompliance *FinancialcomplianceSession) Init(_minIncomeRequired *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.Init(&_Financialcompliance.TransactOpts, _minIncomeRequired)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(uint256 _minIncomeRequired) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) Init(_minIncomeRequired *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.Init(&_Financialcompliance.TransactOpts, _minIncomeRequired)
}

// ModuleBurnAction is a paid mutator transaction binding the contract method 0x372491a2.
//
// Solidity: function moduleBurnAction(address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) ModuleBurnAction(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "moduleBurnAction", arg0, arg1)
}

// ModuleBurnAction is a paid mutator transaction binding the contract method 0x372491a2.
//
// Solidity: function moduleBurnAction(address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceSession) ModuleBurnAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ModuleBurnAction(&_Financialcompliance.TransactOpts, arg0, arg1)
}

// ModuleBurnAction is a paid mutator transaction binding the contract method 0x372491a2.
//
// Solidity: function moduleBurnAction(address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) ModuleBurnAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ModuleBurnAction(&_Financialcompliance.TransactOpts, arg0, arg1)
}

// ModuleMintAction is a paid mutator transaction binding the contract method 0xf104a8c9.
//
// Solidity: function moduleMintAction(address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) ModuleMintAction(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "moduleMintAction", arg0, arg1)
}

// ModuleMintAction is a paid mutator transaction binding the contract method 0xf104a8c9.
//
// Solidity: function moduleMintAction(address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceSession) ModuleMintAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ModuleMintAction(&_Financialcompliance.TransactOpts, arg0, arg1)
}

// ModuleMintAction is a paid mutator transaction binding the contract method 0xf104a8c9.
//
// Solidity: function moduleMintAction(address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) ModuleMintAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ModuleMintAction(&_Financialcompliance.TransactOpts, arg0, arg1)
}

// ModuleTransferAction is a paid mutator transaction binding the contract method 0x2cb7e1ec.
//
// Solidity: function moduleTransferAction(address , address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) ModuleTransferAction(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "moduleTransferAction", arg0, arg1, arg2)
}

// ModuleTransferAction is a paid mutator transaction binding the contract method 0x2cb7e1ec.
//
// Solidity: function moduleTransferAction(address , address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceSession) ModuleTransferAction(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ModuleTransferAction(&_Financialcompliance.TransactOpts, arg0, arg1, arg2)
}

// ModuleTransferAction is a paid mutator transaction binding the contract method 0x2cb7e1ec.
//
// Solidity: function moduleTransferAction(address , address , uint256 ) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) ModuleTransferAction(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Financialcompliance.Contract.ModuleTransferAction(&_Financialcompliance.TransactOpts, arg0, arg1, arg2)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Financialcompliance *FinancialcomplianceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Financialcompliance *FinancialcomplianceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Financialcompliance.Contract.RenounceOwnership(&_Financialcompliance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Financialcompliance.Contract.RenounceOwnership(&_Financialcompliance.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Financialcompliance *FinancialcomplianceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Financialcompliance.Contract.TransferOwnership(&_Financialcompliance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Financialcompliance.Contract.TransferOwnership(&_Financialcompliance.TransactOpts, newOwner)
}

// UnbindCompliance is a paid mutator transaction binding the contract method 0x0694a5fb.
//
// Solidity: function unbindCompliance(address _compliance) returns()
func (_Financialcompliance *FinancialcomplianceTransactor) UnbindCompliance(opts *bind.TransactOpts, _compliance common.Address) (*types.Transaction, error) {
	return _Financialcompliance.contract.Transact(opts, "unbindCompliance", _compliance)
}

// UnbindCompliance is a paid mutator transaction binding the contract method 0x0694a5fb.
//
// Solidity: function unbindCompliance(address _compliance) returns()
func (_Financialcompliance *FinancialcomplianceSession) UnbindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Financialcompliance.Contract.UnbindCompliance(&_Financialcompliance.TransactOpts, _compliance)
}

// UnbindCompliance is a paid mutator transaction binding the contract method 0x0694a5fb.
//
// Solidity: function unbindCompliance(address _compliance) returns()
func (_Financialcompliance *FinancialcomplianceTransactorSession) UnbindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Financialcompliance.Contract.UnbindCompliance(&_Financialcompliance.TransactOpts, _compliance)
}

// FinancialcomplianceBuyerApprovedIterator is returned from FilterBuyerApproved and is used to iterate over the raw logs and unpacked data for BuyerApproved events raised by the Financialcompliance contract.
type FinancialcomplianceBuyerApprovedIterator struct {
	Event *FinancialcomplianceBuyerApproved // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceBuyerApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceBuyerApproved)
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
		it.Event = new(FinancialcomplianceBuyerApproved)
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
func (it *FinancialcomplianceBuyerApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceBuyerApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceBuyerApproved represents a BuyerApproved event raised by the Financialcompliance contract.
type FinancialcomplianceBuyerApproved struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBuyerApproved is a free log retrieval operation binding the contract event 0x72bd472c4d4dc7af2e7c012d0325cf266dfe9edfe2f2685807bbb26f8c9fe652.
//
// Solidity: event BuyerApproved(address indexed buyer)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterBuyerApproved(opts *bind.FilterOpts, buyer []common.Address) (*FinancialcomplianceBuyerApprovedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "BuyerApproved", buyerRule)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceBuyerApprovedIterator{contract: _Financialcompliance.contract, event: "BuyerApproved", logs: logs, sub: sub}, nil
}

// WatchBuyerApproved is a free log subscription operation binding the contract event 0x72bd472c4d4dc7af2e7c012d0325cf266dfe9edfe2f2685807bbb26f8c9fe652.
//
// Solidity: event BuyerApproved(address indexed buyer)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchBuyerApproved(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceBuyerApproved, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "BuyerApproved", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceBuyerApproved)
				if err := _Financialcompliance.contract.UnpackLog(event, "BuyerApproved", log); err != nil {
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

// ParseBuyerApproved is a log parse operation binding the contract event 0x72bd472c4d4dc7af2e7c012d0325cf266dfe9edfe2f2685807bbb26f8c9fe652.
//
// Solidity: event BuyerApproved(address indexed buyer)
func (_Financialcompliance *FinancialcomplianceFilterer) ParseBuyerApproved(log types.Log) (*FinancialcomplianceBuyerApproved, error) {
	event := new(FinancialcomplianceBuyerApproved)
	if err := _Financialcompliance.contract.UnpackLog(event, "BuyerApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialcomplianceComplianceBoundIterator is returned from FilterComplianceBound and is used to iterate over the raw logs and unpacked data for ComplianceBound events raised by the Financialcompliance contract.
type FinancialcomplianceComplianceBoundIterator struct {
	Event *FinancialcomplianceComplianceBound // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceComplianceBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceComplianceBound)
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
		it.Event = new(FinancialcomplianceComplianceBound)
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
func (it *FinancialcomplianceComplianceBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceComplianceBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceComplianceBound represents a ComplianceBound event raised by the Financialcompliance contract.
type FinancialcomplianceComplianceBound struct {
	Compliance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComplianceBound is a free log retrieval operation binding the contract event 0x1f7b76c58fb697eb53c6c7c1becb96911516a136e24d7ced386b2355358b75a3.
//
// Solidity: event ComplianceBound(address indexed _compliance)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterComplianceBound(opts *bind.FilterOpts, _compliance []common.Address) (*FinancialcomplianceComplianceBoundIterator, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "ComplianceBound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceComplianceBoundIterator{contract: _Financialcompliance.contract, event: "ComplianceBound", logs: logs, sub: sub}, nil
}

// WatchComplianceBound is a free log subscription operation binding the contract event 0x1f7b76c58fb697eb53c6c7c1becb96911516a136e24d7ced386b2355358b75a3.
//
// Solidity: event ComplianceBound(address indexed _compliance)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchComplianceBound(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceComplianceBound, _compliance []common.Address) (event.Subscription, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "ComplianceBound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceComplianceBound)
				if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceBound", log); err != nil {
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

// ParseComplianceBound is a log parse operation binding the contract event 0x1f7b76c58fb697eb53c6c7c1becb96911516a136e24d7ced386b2355358b75a3.
//
// Solidity: event ComplianceBound(address indexed _compliance)
func (_Financialcompliance *FinancialcomplianceFilterer) ParseComplianceBound(log types.Log) (*FinancialcomplianceComplianceBound, error) {
	event := new(FinancialcomplianceComplianceBound)
	if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialcomplianceComplianceCheckFailedEventIterator is returned from FilterComplianceCheckFailedEvent and is used to iterate over the raw logs and unpacked data for ComplianceCheckFailedEvent events raised by the Financialcompliance contract.
type FinancialcomplianceComplianceCheckFailedEventIterator struct {
	Event *FinancialcomplianceComplianceCheckFailedEvent // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceComplianceCheckFailedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceComplianceCheckFailedEvent)
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
		it.Event = new(FinancialcomplianceComplianceCheckFailedEvent)
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
func (it *FinancialcomplianceComplianceCheckFailedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceComplianceCheckFailedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceComplianceCheckFailedEvent represents a ComplianceCheckFailedEvent event raised by the Financialcompliance contract.
type FinancialcomplianceComplianceCheckFailedEvent struct {
	From   common.Address
	To     common.Address
	Value  *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterComplianceCheckFailedEvent is a free log retrieval operation binding the contract event 0x8634ff10ead29d455f64f535da1121ff50f50b826344b5c834681295f5041195.
//
// Solidity: event ComplianceCheckFailedEvent(address indexed from, address indexed to, uint256 value, string reason)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterComplianceCheckFailedEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FinancialcomplianceComplianceCheckFailedEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "ComplianceCheckFailedEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceComplianceCheckFailedEventIterator{contract: _Financialcompliance.contract, event: "ComplianceCheckFailedEvent", logs: logs, sub: sub}, nil
}

// WatchComplianceCheckFailedEvent is a free log subscription operation binding the contract event 0x8634ff10ead29d455f64f535da1121ff50f50b826344b5c834681295f5041195.
//
// Solidity: event ComplianceCheckFailedEvent(address indexed from, address indexed to, uint256 value, string reason)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchComplianceCheckFailedEvent(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceComplianceCheckFailedEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "ComplianceCheckFailedEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceComplianceCheckFailedEvent)
				if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceCheckFailedEvent", log); err != nil {
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

// ParseComplianceCheckFailedEvent is a log parse operation binding the contract event 0x8634ff10ead29d455f64f535da1121ff50f50b826344b5c834681295f5041195.
//
// Solidity: event ComplianceCheckFailedEvent(address indexed from, address indexed to, uint256 value, string reason)
func (_Financialcompliance *FinancialcomplianceFilterer) ParseComplianceCheckFailedEvent(log types.Log) (*FinancialcomplianceComplianceCheckFailedEvent, error) {
	event := new(FinancialcomplianceComplianceCheckFailedEvent)
	if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceCheckFailedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialcomplianceComplianceCheckPassedIterator is returned from FilterComplianceCheckPassed and is used to iterate over the raw logs and unpacked data for ComplianceCheckPassed events raised by the Financialcompliance contract.
type FinancialcomplianceComplianceCheckPassedIterator struct {
	Event *FinancialcomplianceComplianceCheckPassed // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceComplianceCheckPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceComplianceCheckPassed)
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
		it.Event = new(FinancialcomplianceComplianceCheckPassed)
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
func (it *FinancialcomplianceComplianceCheckPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceComplianceCheckPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceComplianceCheckPassed represents a ComplianceCheckPassed event raised by the Financialcompliance contract.
type FinancialcomplianceComplianceCheckPassed struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterComplianceCheckPassed is a free log retrieval operation binding the contract event 0xcf998f09a5dd6c9f016b8bb7b360ad34c2a2383fedee312773430a0646b7e8a9.
//
// Solidity: event ComplianceCheckPassed(address indexed from, address indexed to, uint256 value)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterComplianceCheckPassed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FinancialcomplianceComplianceCheckPassedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "ComplianceCheckPassed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceComplianceCheckPassedIterator{contract: _Financialcompliance.contract, event: "ComplianceCheckPassed", logs: logs, sub: sub}, nil
}

// WatchComplianceCheckPassed is a free log subscription operation binding the contract event 0xcf998f09a5dd6c9f016b8bb7b360ad34c2a2383fedee312773430a0646b7e8a9.
//
// Solidity: event ComplianceCheckPassed(address indexed from, address indexed to, uint256 value)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchComplianceCheckPassed(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceComplianceCheckPassed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "ComplianceCheckPassed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceComplianceCheckPassed)
				if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceCheckPassed", log); err != nil {
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

// ParseComplianceCheckPassed is a log parse operation binding the contract event 0xcf998f09a5dd6c9f016b8bb7b360ad34c2a2383fedee312773430a0646b7e8a9.
//
// Solidity: event ComplianceCheckPassed(address indexed from, address indexed to, uint256 value)
func (_Financialcompliance *FinancialcomplianceFilterer) ParseComplianceCheckPassed(log types.Log) (*FinancialcomplianceComplianceCheckPassed, error) {
	event := new(FinancialcomplianceComplianceCheckPassed)
	if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceCheckPassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialcomplianceComplianceUnboundIterator is returned from FilterComplianceUnbound and is used to iterate over the raw logs and unpacked data for ComplianceUnbound events raised by the Financialcompliance contract.
type FinancialcomplianceComplianceUnboundIterator struct {
	Event *FinancialcomplianceComplianceUnbound // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceComplianceUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceComplianceUnbound)
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
		it.Event = new(FinancialcomplianceComplianceUnbound)
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
func (it *FinancialcomplianceComplianceUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceComplianceUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceComplianceUnbound represents a ComplianceUnbound event raised by the Financialcompliance contract.
type FinancialcomplianceComplianceUnbound struct {
	Compliance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComplianceUnbound is a free log retrieval operation binding the contract event 0x408b49d9be1c914c52a0227e18a077e5a892dddf32a26cfa94a5d9708fad7718.
//
// Solidity: event ComplianceUnbound(address indexed _compliance)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterComplianceUnbound(opts *bind.FilterOpts, _compliance []common.Address) (*FinancialcomplianceComplianceUnboundIterator, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "ComplianceUnbound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceComplianceUnboundIterator{contract: _Financialcompliance.contract, event: "ComplianceUnbound", logs: logs, sub: sub}, nil
}

// WatchComplianceUnbound is a free log subscription operation binding the contract event 0x408b49d9be1c914c52a0227e18a077e5a892dddf32a26cfa94a5d9708fad7718.
//
// Solidity: event ComplianceUnbound(address indexed _compliance)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchComplianceUnbound(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceComplianceUnbound, _compliance []common.Address) (event.Subscription, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "ComplianceUnbound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceComplianceUnbound)
				if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceUnbound", log); err != nil {
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

// ParseComplianceUnbound is a log parse operation binding the contract event 0x408b49d9be1c914c52a0227e18a077e5a892dddf32a26cfa94a5d9708fad7718.
//
// Solidity: event ComplianceUnbound(address indexed _compliance)
func (_Financialcompliance *FinancialcomplianceFilterer) ParseComplianceUnbound(log types.Log) (*FinancialcomplianceComplianceUnbound, error) {
	event := new(FinancialcomplianceComplianceUnbound)
	if err := _Financialcompliance.contract.UnpackLog(event, "ComplianceUnbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialcomplianceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Financialcompliance contract.
type FinancialcomplianceInitializedIterator struct {
	Event *FinancialcomplianceInitialized // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceInitialized)
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
		it.Event = new(FinancialcomplianceInitialized)
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
func (it *FinancialcomplianceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceInitialized represents a Initialized event raised by the Financialcompliance contract.
type FinancialcomplianceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterInitialized(opts *bind.FilterOpts) (*FinancialcomplianceInitializedIterator, error) {

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceInitializedIterator{contract: _Financialcompliance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceInitialized) (event.Subscription, error) {

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceInitialized)
				if err := _Financialcompliance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Financialcompliance *FinancialcomplianceFilterer) ParseInitialized(log types.Log) (*FinancialcomplianceInitialized, error) {
	event := new(FinancialcomplianceInitialized)
	if err := _Financialcompliance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinancialcomplianceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Financialcompliance contract.
type FinancialcomplianceOwnershipTransferredIterator struct {
	Event *FinancialcomplianceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FinancialcomplianceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinancialcomplianceOwnershipTransferred)
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
		it.Event = new(FinancialcomplianceOwnershipTransferred)
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
func (it *FinancialcomplianceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinancialcomplianceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinancialcomplianceOwnershipTransferred represents a OwnershipTransferred event raised by the Financialcompliance contract.
type FinancialcomplianceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Financialcompliance *FinancialcomplianceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinancialcomplianceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Financialcompliance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinancialcomplianceOwnershipTransferredIterator{contract: _Financialcompliance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Financialcompliance *FinancialcomplianceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinancialcomplianceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Financialcompliance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinancialcomplianceOwnershipTransferred)
				if err := _Financialcompliance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Financialcompliance *FinancialcomplianceFilterer) ParseOwnershipTransferred(log types.Log) (*FinancialcomplianceOwnershipTransferred, error) {
	event := new(FinancialcomplianceOwnershipTransferred)
	if err := _Financialcompliance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
