// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrymdcompliance

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

// RegistryMDCompliancePropertyInfo is an auto generated low-level Go binding around an user-defined struct.
type RegistryMDCompliancePropertyInfo struct {
	MatriculaId     *big.Int
	Folha           *big.Int
	Oficio          uint16
	NomeOficio      string
	Comarca         string
	Endereco        string
	Metragem        *big.Int
	Proprietario    common.Address
	MatriculaOrigem *big.Int
	Iptu            *big.Int
	Itr             *big.Int
	Tipo            uint8
	IsRegular       bool
}

// RegistrymdcomplianceMetaData contains all meta data concerning the Registrymdcompliance contract.
var RegistrymdcomplianceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addAverbacao\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"averbacao\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bindCompliance\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canComplianceBind\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getAverbacoes\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProperty\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRegistryMDCompliance.PropertyInfo\",\"components\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"folha\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oficio\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"nomeOficio\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"comarca\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"endereco\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metragem\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proprietario\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"matriculaOrigem\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"iptu\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itr\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tipo\",\"type\":\"uint8\",\"internalType\":\"enumRegistryMDCompliance.PropertyType\"},{\"name\":\"isRegular\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isComplianceBound\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPlugAndPlay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"moduleBurnAction\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"moduleCheck\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"moduleMintAction\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"moduleTransferAction\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerProperty\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"folha\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oficio\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"nomeOficio\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"comarca\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"endereco\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metragem\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proprietario\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"matriculaOrigem\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"iptu\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itr\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tipo\",\"type\":\"uint8\",\"internalType\":\"enumRegistryMDCompliance.PropertyType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferProperty\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"novoProprietario\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbindCompliance\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProperty\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endereco\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"metragem\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"iptu\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"itr\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isRegular\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AverbacaoAdded\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"averbacao\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceBound\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceUnbound\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PropertyRegistered\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proprietario\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PropertyTransferred\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PropertyUpdated\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ComplianceCheckFailed\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"DuplicateRegistry\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PropertyNotRegistered\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PropertyTransferNotAllowed\",\"inputs\":[{\"name\":\"matriculaId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// RegistrymdcomplianceABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrymdcomplianceMetaData.ABI instead.
var RegistrymdcomplianceABI = RegistrymdcomplianceMetaData.ABI

// Registrymdcompliance is an auto generated Go binding around an Ethereum contract.
type Registrymdcompliance struct {
	RegistrymdcomplianceCaller     // Read-only binding to the contract
	RegistrymdcomplianceTransactor // Write-only binding to the contract
	RegistrymdcomplianceFilterer   // Log filterer for contract events
}

// RegistrymdcomplianceCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrymdcomplianceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrymdcomplianceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrymdcomplianceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrymdcomplianceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrymdcomplianceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrymdcomplianceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrymdcomplianceSession struct {
	Contract     *Registrymdcompliance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RegistrymdcomplianceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrymdcomplianceCallerSession struct {
	Contract *RegistrymdcomplianceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RegistrymdcomplianceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrymdcomplianceTransactorSession struct {
	Contract     *RegistrymdcomplianceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RegistrymdcomplianceRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrymdcomplianceRaw struct {
	Contract *Registrymdcompliance // Generic contract binding to access the raw methods on
}

// RegistrymdcomplianceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrymdcomplianceCallerRaw struct {
	Contract *RegistrymdcomplianceCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrymdcomplianceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrymdcomplianceTransactorRaw struct {
	Contract *RegistrymdcomplianceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrymdcompliance creates a new instance of Registrymdcompliance, bound to a specific deployed contract.
func NewRegistrymdcompliance(address common.Address, backend bind.ContractBackend) (*Registrymdcompliance, error) {
	contract, err := bindRegistrymdcompliance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrymdcompliance{RegistrymdcomplianceCaller: RegistrymdcomplianceCaller{contract: contract}, RegistrymdcomplianceTransactor: RegistrymdcomplianceTransactor{contract: contract}, RegistrymdcomplianceFilterer: RegistrymdcomplianceFilterer{contract: contract}}, nil
}

// NewRegistrymdcomplianceCaller creates a new read-only instance of Registrymdcompliance, bound to a specific deployed contract.
func NewRegistrymdcomplianceCaller(address common.Address, caller bind.ContractCaller) (*RegistrymdcomplianceCaller, error) {
	contract, err := bindRegistrymdcompliance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceCaller{contract: contract}, nil
}

// NewRegistrymdcomplianceTransactor creates a new write-only instance of Registrymdcompliance, bound to a specific deployed contract.
func NewRegistrymdcomplianceTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrymdcomplianceTransactor, error) {
	contract, err := bindRegistrymdcompliance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceTransactor{contract: contract}, nil
}

// NewRegistrymdcomplianceFilterer creates a new log filterer instance of Registrymdcompliance, bound to a specific deployed contract.
func NewRegistrymdcomplianceFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrymdcomplianceFilterer, error) {
	contract, err := bindRegistrymdcompliance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceFilterer{contract: contract}, nil
}

// bindRegistrymdcompliance binds a generic wrapper to an already deployed contract.
func bindRegistrymdcompliance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrymdcomplianceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrymdcompliance *RegistrymdcomplianceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrymdcompliance.Contract.RegistrymdcomplianceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrymdcompliance *RegistrymdcomplianceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.RegistrymdcomplianceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrymdcompliance *RegistrymdcomplianceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.RegistrymdcomplianceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrymdcompliance *RegistrymdcomplianceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrymdcompliance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrymdcompliance *RegistrymdcomplianceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrymdcompliance *RegistrymdcomplianceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.contract.Transact(opts, method, params...)
}

// CanComplianceBind is a free data retrieval call binding the contract method 0xbcc21053.
//
// Solidity: function canComplianceBind(address ) pure returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCaller) CanComplianceBind(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "canComplianceBind", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanComplianceBind is a free data retrieval call binding the contract method 0xbcc21053.
//
// Solidity: function canComplianceBind(address ) pure returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceSession) CanComplianceBind(arg0 common.Address) (bool, error) {
	return _Registrymdcompliance.Contract.CanComplianceBind(&_Registrymdcompliance.CallOpts, arg0)
}

// CanComplianceBind is a free data retrieval call binding the contract method 0xbcc21053.
//
// Solidity: function canComplianceBind(address ) pure returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) CanComplianceBind(arg0 common.Address) (bool, error) {
	return _Registrymdcompliance.Contract.CanComplianceBind(&_Registrymdcompliance.CallOpts, arg0)
}

// GetAverbacoes is a free data retrieval call binding the contract method 0xd132eaa7.
//
// Solidity: function getAverbacoes(uint256 matriculaId) view returns(string[])
func (_Registrymdcompliance *RegistrymdcomplianceCaller) GetAverbacoes(opts *bind.CallOpts, matriculaId *big.Int) ([]string, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "getAverbacoes", matriculaId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAverbacoes is a free data retrieval call binding the contract method 0xd132eaa7.
//
// Solidity: function getAverbacoes(uint256 matriculaId) view returns(string[])
func (_Registrymdcompliance *RegistrymdcomplianceSession) GetAverbacoes(matriculaId *big.Int) ([]string, error) {
	return _Registrymdcompliance.Contract.GetAverbacoes(&_Registrymdcompliance.CallOpts, matriculaId)
}

// GetAverbacoes is a free data retrieval call binding the contract method 0xd132eaa7.
//
// Solidity: function getAverbacoes(uint256 matriculaId) view returns(string[])
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) GetAverbacoes(matriculaId *big.Int) ([]string, error) {
	return _Registrymdcompliance.Contract.GetAverbacoes(&_Registrymdcompliance.CallOpts, matriculaId)
}

// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
//
// Solidity: function getProperty(uint256 matriculaId) view returns((uint256,uint256,uint16,string,string,string,uint256,address,uint256,uint256,uint256,uint8,bool))
func (_Registrymdcompliance *RegistrymdcomplianceCaller) GetProperty(opts *bind.CallOpts, matriculaId *big.Int) (RegistryMDCompliancePropertyInfo, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "getProperty", matriculaId)

	if err != nil {
		return *new(RegistryMDCompliancePropertyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryMDCompliancePropertyInfo)).(*RegistryMDCompliancePropertyInfo)

	return out0, err

}

// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
//
// Solidity: function getProperty(uint256 matriculaId) view returns((uint256,uint256,uint16,string,string,string,uint256,address,uint256,uint256,uint256,uint8,bool))
func (_Registrymdcompliance *RegistrymdcomplianceSession) GetProperty(matriculaId *big.Int) (RegistryMDCompliancePropertyInfo, error) {
	return _Registrymdcompliance.Contract.GetProperty(&_Registrymdcompliance.CallOpts, matriculaId)
}

// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
//
// Solidity: function getProperty(uint256 matriculaId) view returns((uint256,uint256,uint16,string,string,string,uint256,address,uint256,uint256,uint256,uint8,bool))
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) GetProperty(matriculaId *big.Int) (RegistryMDCompliancePropertyInfo, error) {
	return _Registrymdcompliance.Contract.GetProperty(&_Registrymdcompliance.CallOpts, matriculaId)
}

// IsComplianceBound is a free data retrieval call binding the contract method 0x4cf4d295.
//
// Solidity: function isComplianceBound(address _compliance) view returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCaller) IsComplianceBound(opts *bind.CallOpts, _compliance common.Address) (bool, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "isComplianceBound", _compliance)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComplianceBound is a free data retrieval call binding the contract method 0x4cf4d295.
//
// Solidity: function isComplianceBound(address _compliance) view returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceSession) IsComplianceBound(_compliance common.Address) (bool, error) {
	return _Registrymdcompliance.Contract.IsComplianceBound(&_Registrymdcompliance.CallOpts, _compliance)
}

// IsComplianceBound is a free data retrieval call binding the contract method 0x4cf4d295.
//
// Solidity: function isComplianceBound(address _compliance) view returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) IsComplianceBound(_compliance common.Address) (bool, error) {
	return _Registrymdcompliance.Contract.IsComplianceBound(&_Registrymdcompliance.CallOpts, _compliance)
}

// IsPlugAndPlay is a free data retrieval call binding the contract method 0xe6f5e807.
//
// Solidity: function isPlugAndPlay() pure returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCaller) IsPlugAndPlay(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "isPlugAndPlay")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPlugAndPlay is a free data retrieval call binding the contract method 0xe6f5e807.
//
// Solidity: function isPlugAndPlay() pure returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceSession) IsPlugAndPlay() (bool, error) {
	return _Registrymdcompliance.Contract.IsPlugAndPlay(&_Registrymdcompliance.CallOpts)
}

// IsPlugAndPlay is a free data retrieval call binding the contract method 0xe6f5e807.
//
// Solidity: function isPlugAndPlay() pure returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) IsPlugAndPlay() (bool, error) {
	return _Registrymdcompliance.Contract.IsPlugAndPlay(&_Registrymdcompliance.CallOpts)
}

// ModuleCheck is a free data retrieval call binding the contract method 0x013b7ce4.
//
// Solidity: function moduleCheck(address from, address to, uint256 matriculaId, address compliance) view returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCaller) ModuleCheck(opts *bind.CallOpts, from common.Address, to common.Address, matriculaId *big.Int, compliance common.Address) (bool, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "moduleCheck", from, to, matriculaId, compliance)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ModuleCheck is a free data retrieval call binding the contract method 0x013b7ce4.
//
// Solidity: function moduleCheck(address from, address to, uint256 matriculaId, address compliance) view returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceSession) ModuleCheck(from common.Address, to common.Address, matriculaId *big.Int, compliance common.Address) (bool, error) {
	return _Registrymdcompliance.Contract.ModuleCheck(&_Registrymdcompliance.CallOpts, from, to, matriculaId, compliance)
}

// ModuleCheck is a free data retrieval call binding the contract method 0x013b7ce4.
//
// Solidity: function moduleCheck(address from, address to, uint256 matriculaId, address compliance) view returns(bool)
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) ModuleCheck(from common.Address, to common.Address, matriculaId *big.Int, compliance common.Address) (bool, error) {
	return _Registrymdcompliance.Contract.ModuleCheck(&_Registrymdcompliance.CallOpts, from, to, matriculaId, compliance)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Registrymdcompliance *RegistrymdcomplianceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Registrymdcompliance *RegistrymdcomplianceSession) Name() (string, error) {
	return _Registrymdcompliance.Contract.Name(&_Registrymdcompliance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) Name() (string, error) {
	return _Registrymdcompliance.Contract.Name(&_Registrymdcompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrymdcompliance *RegistrymdcomplianceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registrymdcompliance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrymdcompliance *RegistrymdcomplianceSession) Owner() (common.Address, error) {
	return _Registrymdcompliance.Contract.Owner(&_Registrymdcompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrymdcompliance *RegistrymdcomplianceCallerSession) Owner() (common.Address, error) {
	return _Registrymdcompliance.Contract.Owner(&_Registrymdcompliance.CallOpts)
}

// AddAverbacao is a paid mutator transaction binding the contract method 0x6a2cb48a.
//
// Solidity: function addAverbacao(uint256 matriculaId, string averbacao) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) AddAverbacao(opts *bind.TransactOpts, matriculaId *big.Int, averbacao string) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "addAverbacao", matriculaId, averbacao)
}

// AddAverbacao is a paid mutator transaction binding the contract method 0x6a2cb48a.
//
// Solidity: function addAverbacao(uint256 matriculaId, string averbacao) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) AddAverbacao(matriculaId *big.Int, averbacao string) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.AddAverbacao(&_Registrymdcompliance.TransactOpts, matriculaId, averbacao)
}

// AddAverbacao is a paid mutator transaction binding the contract method 0x6a2cb48a.
//
// Solidity: function addAverbacao(uint256 matriculaId, string averbacao) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) AddAverbacao(matriculaId *big.Int, averbacao string) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.AddAverbacao(&_Registrymdcompliance.TransactOpts, matriculaId, averbacao)
}

// BindCompliance is a paid mutator transaction binding the contract method 0x4a932544.
//
// Solidity: function bindCompliance(address _compliance) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) BindCompliance(opts *bind.TransactOpts, _compliance common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "bindCompliance", _compliance)
}

// BindCompliance is a paid mutator transaction binding the contract method 0x4a932544.
//
// Solidity: function bindCompliance(address _compliance) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) BindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.BindCompliance(&_Registrymdcompliance.TransactOpts, _compliance)
}

// BindCompliance is a paid mutator transaction binding the contract method 0x4a932544.
//
// Solidity: function bindCompliance(address _compliance) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) BindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.BindCompliance(&_Registrymdcompliance.TransactOpts, _compliance)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) Init() (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.Init(&_Registrymdcompliance.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) Init() (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.Init(&_Registrymdcompliance.TransactOpts)
}

// ModuleBurnAction is a paid mutator transaction binding the contract method 0x372491a2.
//
// Solidity: function moduleBurnAction(address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) ModuleBurnAction(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "moduleBurnAction", arg0, arg1)
}

// ModuleBurnAction is a paid mutator transaction binding the contract method 0x372491a2.
//
// Solidity: function moduleBurnAction(address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) ModuleBurnAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.ModuleBurnAction(&_Registrymdcompliance.TransactOpts, arg0, arg1)
}

// ModuleBurnAction is a paid mutator transaction binding the contract method 0x372491a2.
//
// Solidity: function moduleBurnAction(address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) ModuleBurnAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.ModuleBurnAction(&_Registrymdcompliance.TransactOpts, arg0, arg1)
}

// ModuleMintAction is a paid mutator transaction binding the contract method 0xf104a8c9.
//
// Solidity: function moduleMintAction(address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) ModuleMintAction(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "moduleMintAction", arg0, arg1)
}

// ModuleMintAction is a paid mutator transaction binding the contract method 0xf104a8c9.
//
// Solidity: function moduleMintAction(address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) ModuleMintAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.ModuleMintAction(&_Registrymdcompliance.TransactOpts, arg0, arg1)
}

// ModuleMintAction is a paid mutator transaction binding the contract method 0xf104a8c9.
//
// Solidity: function moduleMintAction(address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) ModuleMintAction(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.ModuleMintAction(&_Registrymdcompliance.TransactOpts, arg0, arg1)
}

// ModuleTransferAction is a paid mutator transaction binding the contract method 0x2cb7e1ec.
//
// Solidity: function moduleTransferAction(address , address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) ModuleTransferAction(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "moduleTransferAction", arg0, arg1, arg2)
}

// ModuleTransferAction is a paid mutator transaction binding the contract method 0x2cb7e1ec.
//
// Solidity: function moduleTransferAction(address , address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) ModuleTransferAction(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.ModuleTransferAction(&_Registrymdcompliance.TransactOpts, arg0, arg1, arg2)
}

// ModuleTransferAction is a paid mutator transaction binding the contract method 0x2cb7e1ec.
//
// Solidity: function moduleTransferAction(address , address , uint256 ) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) ModuleTransferAction(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.ModuleTransferAction(&_Registrymdcompliance.TransactOpts, arg0, arg1, arg2)
}

// RegisterProperty is a paid mutator transaction binding the contract method 0x8dcc0935.
//
// Solidity: function registerProperty(uint256 matriculaId, uint256 folha, uint16 oficio, string nomeOficio, string comarca, string endereco, uint256 metragem, address proprietario, uint256 matriculaOrigem, uint256 iptu, uint256 itr, uint8 tipo) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) RegisterProperty(opts *bind.TransactOpts, matriculaId *big.Int, folha *big.Int, oficio uint16, nomeOficio string, comarca string, endereco string, metragem *big.Int, proprietario common.Address, matriculaOrigem *big.Int, iptu *big.Int, itr *big.Int, tipo uint8) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "registerProperty", matriculaId, folha, oficio, nomeOficio, comarca, endereco, metragem, proprietario, matriculaOrigem, iptu, itr, tipo)
}

// RegisterProperty is a paid mutator transaction binding the contract method 0x8dcc0935.
//
// Solidity: function registerProperty(uint256 matriculaId, uint256 folha, uint16 oficio, string nomeOficio, string comarca, string endereco, uint256 metragem, address proprietario, uint256 matriculaOrigem, uint256 iptu, uint256 itr, uint8 tipo) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) RegisterProperty(matriculaId *big.Int, folha *big.Int, oficio uint16, nomeOficio string, comarca string, endereco string, metragem *big.Int, proprietario common.Address, matriculaOrigem *big.Int, iptu *big.Int, itr *big.Int, tipo uint8) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.RegisterProperty(&_Registrymdcompliance.TransactOpts, matriculaId, folha, oficio, nomeOficio, comarca, endereco, metragem, proprietario, matriculaOrigem, iptu, itr, tipo)
}

// RegisterProperty is a paid mutator transaction binding the contract method 0x8dcc0935.
//
// Solidity: function registerProperty(uint256 matriculaId, uint256 folha, uint16 oficio, string nomeOficio, string comarca, string endereco, uint256 metragem, address proprietario, uint256 matriculaOrigem, uint256 iptu, uint256 itr, uint8 tipo) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) RegisterProperty(matriculaId *big.Int, folha *big.Int, oficio uint16, nomeOficio string, comarca string, endereco string, metragem *big.Int, proprietario common.Address, matriculaOrigem *big.Int, iptu *big.Int, itr *big.Int, tipo uint8) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.RegisterProperty(&_Registrymdcompliance.TransactOpts, matriculaId, folha, oficio, nomeOficio, comarca, endereco, metragem, proprietario, matriculaOrigem, iptu, itr, tipo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.RenounceOwnership(&_Registrymdcompliance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.RenounceOwnership(&_Registrymdcompliance.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.TransferOwnership(&_Registrymdcompliance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.TransferOwnership(&_Registrymdcompliance.TransactOpts, newOwner)
}

// TransferProperty is a paid mutator transaction binding the contract method 0xff356237.
//
// Solidity: function transferProperty(uint256 matriculaId, address novoProprietario) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) TransferProperty(opts *bind.TransactOpts, matriculaId *big.Int, novoProprietario common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "transferProperty", matriculaId, novoProprietario)
}

// TransferProperty is a paid mutator transaction binding the contract method 0xff356237.
//
// Solidity: function transferProperty(uint256 matriculaId, address novoProprietario) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) TransferProperty(matriculaId *big.Int, novoProprietario common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.TransferProperty(&_Registrymdcompliance.TransactOpts, matriculaId, novoProprietario)
}

// TransferProperty is a paid mutator transaction binding the contract method 0xff356237.
//
// Solidity: function transferProperty(uint256 matriculaId, address novoProprietario) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) TransferProperty(matriculaId *big.Int, novoProprietario common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.TransferProperty(&_Registrymdcompliance.TransactOpts, matriculaId, novoProprietario)
}

// UnbindCompliance is a paid mutator transaction binding the contract method 0x0694a5fb.
//
// Solidity: function unbindCompliance(address _compliance) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) UnbindCompliance(opts *bind.TransactOpts, _compliance common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "unbindCompliance", _compliance)
}

// UnbindCompliance is a paid mutator transaction binding the contract method 0x0694a5fb.
//
// Solidity: function unbindCompliance(address _compliance) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) UnbindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.UnbindCompliance(&_Registrymdcompliance.TransactOpts, _compliance)
}

// UnbindCompliance is a paid mutator transaction binding the contract method 0x0694a5fb.
//
// Solidity: function unbindCompliance(address _compliance) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) UnbindCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.UnbindCompliance(&_Registrymdcompliance.TransactOpts, _compliance)
}

// UpdateProperty is a paid mutator transaction binding the contract method 0xb330e069.
//
// Solidity: function updateProperty(uint256 matriculaId, string endereco, uint256 metragem, uint256 iptu, uint256 itr, bool isRegular) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactor) UpdateProperty(opts *bind.TransactOpts, matriculaId *big.Int, endereco string, metragem *big.Int, iptu *big.Int, itr *big.Int, isRegular bool) (*types.Transaction, error) {
	return _Registrymdcompliance.contract.Transact(opts, "updateProperty", matriculaId, endereco, metragem, iptu, itr, isRegular)
}

// UpdateProperty is a paid mutator transaction binding the contract method 0xb330e069.
//
// Solidity: function updateProperty(uint256 matriculaId, string endereco, uint256 metragem, uint256 iptu, uint256 itr, bool isRegular) returns()
func (_Registrymdcompliance *RegistrymdcomplianceSession) UpdateProperty(matriculaId *big.Int, endereco string, metragem *big.Int, iptu *big.Int, itr *big.Int, isRegular bool) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.UpdateProperty(&_Registrymdcompliance.TransactOpts, matriculaId, endereco, metragem, iptu, itr, isRegular)
}

// UpdateProperty is a paid mutator transaction binding the contract method 0xb330e069.
//
// Solidity: function updateProperty(uint256 matriculaId, string endereco, uint256 metragem, uint256 iptu, uint256 itr, bool isRegular) returns()
func (_Registrymdcompliance *RegistrymdcomplianceTransactorSession) UpdateProperty(matriculaId *big.Int, endereco string, metragem *big.Int, iptu *big.Int, itr *big.Int, isRegular bool) (*types.Transaction, error) {
	return _Registrymdcompliance.Contract.UpdateProperty(&_Registrymdcompliance.TransactOpts, matriculaId, endereco, metragem, iptu, itr, isRegular)
}

// RegistrymdcomplianceAverbacaoAddedIterator is returned from FilterAverbacaoAdded and is used to iterate over the raw logs and unpacked data for AverbacaoAdded events raised by the Registrymdcompliance contract.
type RegistrymdcomplianceAverbacaoAddedIterator struct {
	Event *RegistrymdcomplianceAverbacaoAdded // Event containing the contract specifics and raw log

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
func (it *RegistrymdcomplianceAverbacaoAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcomplianceAverbacaoAdded)
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
		it.Event = new(RegistrymdcomplianceAverbacaoAdded)
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
func (it *RegistrymdcomplianceAverbacaoAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcomplianceAverbacaoAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcomplianceAverbacaoAdded represents a AverbacaoAdded event raised by the Registrymdcompliance contract.
type RegistrymdcomplianceAverbacaoAdded struct {
	MatriculaId *big.Int
	Averbacao   string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAverbacaoAdded is a free log retrieval operation binding the contract event 0xbf8ce0448bec8fa412f9b0cedb39e02ce79dee5343e3287490165f72876a68cd.
//
// Solidity: event AverbacaoAdded(uint256 indexed matriculaId, string averbacao)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterAverbacaoAdded(opts *bind.FilterOpts, matriculaId []*big.Int) (*RegistrymdcomplianceAverbacaoAddedIterator, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "AverbacaoAdded", matriculaIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceAverbacaoAddedIterator{contract: _Registrymdcompliance.contract, event: "AverbacaoAdded", logs: logs, sub: sub}, nil
}

// WatchAverbacaoAdded is a free log subscription operation binding the contract event 0xbf8ce0448bec8fa412f9b0cedb39e02ce79dee5343e3287490165f72876a68cd.
//
// Solidity: event AverbacaoAdded(uint256 indexed matriculaId, string averbacao)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchAverbacaoAdded(opts *bind.WatchOpts, sink chan<- *RegistrymdcomplianceAverbacaoAdded, matriculaId []*big.Int) (event.Subscription, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "AverbacaoAdded", matriculaIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcomplianceAverbacaoAdded)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "AverbacaoAdded", log); err != nil {
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

// ParseAverbacaoAdded is a log parse operation binding the contract event 0xbf8ce0448bec8fa412f9b0cedb39e02ce79dee5343e3287490165f72876a68cd.
//
// Solidity: event AverbacaoAdded(uint256 indexed matriculaId, string averbacao)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParseAverbacaoAdded(log types.Log) (*RegistrymdcomplianceAverbacaoAdded, error) {
	event := new(RegistrymdcomplianceAverbacaoAdded)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "AverbacaoAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcomplianceComplianceBoundIterator is returned from FilterComplianceBound and is used to iterate over the raw logs and unpacked data for ComplianceBound events raised by the Registrymdcompliance contract.
type RegistrymdcomplianceComplianceBoundIterator struct {
	Event *RegistrymdcomplianceComplianceBound // Event containing the contract specifics and raw log

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
func (it *RegistrymdcomplianceComplianceBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcomplianceComplianceBound)
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
		it.Event = new(RegistrymdcomplianceComplianceBound)
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
func (it *RegistrymdcomplianceComplianceBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcomplianceComplianceBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcomplianceComplianceBound represents a ComplianceBound event raised by the Registrymdcompliance contract.
type RegistrymdcomplianceComplianceBound struct {
	Compliance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComplianceBound is a free log retrieval operation binding the contract event 0x1f7b76c58fb697eb53c6c7c1becb96911516a136e24d7ced386b2355358b75a3.
//
// Solidity: event ComplianceBound(address indexed _compliance)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterComplianceBound(opts *bind.FilterOpts, _compliance []common.Address) (*RegistrymdcomplianceComplianceBoundIterator, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "ComplianceBound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceComplianceBoundIterator{contract: _Registrymdcompliance.contract, event: "ComplianceBound", logs: logs, sub: sub}, nil
}

// WatchComplianceBound is a free log subscription operation binding the contract event 0x1f7b76c58fb697eb53c6c7c1becb96911516a136e24d7ced386b2355358b75a3.
//
// Solidity: event ComplianceBound(address indexed _compliance)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchComplianceBound(opts *bind.WatchOpts, sink chan<- *RegistrymdcomplianceComplianceBound, _compliance []common.Address) (event.Subscription, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "ComplianceBound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcomplianceComplianceBound)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "ComplianceBound", log); err != nil {
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
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParseComplianceBound(log types.Log) (*RegistrymdcomplianceComplianceBound, error) {
	event := new(RegistrymdcomplianceComplianceBound)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "ComplianceBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcomplianceComplianceUnboundIterator is returned from FilterComplianceUnbound and is used to iterate over the raw logs and unpacked data for ComplianceUnbound events raised by the Registrymdcompliance contract.
type RegistrymdcomplianceComplianceUnboundIterator struct {
	Event *RegistrymdcomplianceComplianceUnbound // Event containing the contract specifics and raw log

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
func (it *RegistrymdcomplianceComplianceUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcomplianceComplianceUnbound)
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
		it.Event = new(RegistrymdcomplianceComplianceUnbound)
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
func (it *RegistrymdcomplianceComplianceUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcomplianceComplianceUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcomplianceComplianceUnbound represents a ComplianceUnbound event raised by the Registrymdcompliance contract.
type RegistrymdcomplianceComplianceUnbound struct {
	Compliance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComplianceUnbound is a free log retrieval operation binding the contract event 0x408b49d9be1c914c52a0227e18a077e5a892dddf32a26cfa94a5d9708fad7718.
//
// Solidity: event ComplianceUnbound(address indexed _compliance)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterComplianceUnbound(opts *bind.FilterOpts, _compliance []common.Address) (*RegistrymdcomplianceComplianceUnboundIterator, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "ComplianceUnbound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceComplianceUnboundIterator{contract: _Registrymdcompliance.contract, event: "ComplianceUnbound", logs: logs, sub: sub}, nil
}

// WatchComplianceUnbound is a free log subscription operation binding the contract event 0x408b49d9be1c914c52a0227e18a077e5a892dddf32a26cfa94a5d9708fad7718.
//
// Solidity: event ComplianceUnbound(address indexed _compliance)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchComplianceUnbound(opts *bind.WatchOpts, sink chan<- *RegistrymdcomplianceComplianceUnbound, _compliance []common.Address) (event.Subscription, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "ComplianceUnbound", _complianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcomplianceComplianceUnbound)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "ComplianceUnbound", log); err != nil {
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
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParseComplianceUnbound(log types.Log) (*RegistrymdcomplianceComplianceUnbound, error) {
	event := new(RegistrymdcomplianceComplianceUnbound)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "ComplianceUnbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcomplianceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registrymdcompliance contract.
type RegistrymdcomplianceInitializedIterator struct {
	Event *RegistrymdcomplianceInitialized // Event containing the contract specifics and raw log

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
func (it *RegistrymdcomplianceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcomplianceInitialized)
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
		it.Event = new(RegistrymdcomplianceInitialized)
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
func (it *RegistrymdcomplianceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcomplianceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcomplianceInitialized represents a Initialized event raised by the Registrymdcompliance contract.
type RegistrymdcomplianceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistrymdcomplianceInitializedIterator, error) {

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceInitializedIterator{contract: _Registrymdcompliance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistrymdcomplianceInitialized) (event.Subscription, error) {

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcomplianceInitialized)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParseInitialized(log types.Log) (*RegistrymdcomplianceInitialized, error) {
	event := new(RegistrymdcomplianceInitialized)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcomplianceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registrymdcompliance contract.
type RegistrymdcomplianceOwnershipTransferredIterator struct {
	Event *RegistrymdcomplianceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrymdcomplianceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcomplianceOwnershipTransferred)
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
		it.Event = new(RegistrymdcomplianceOwnershipTransferred)
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
func (it *RegistrymdcomplianceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcomplianceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcomplianceOwnershipTransferred represents a OwnershipTransferred event raised by the Registrymdcompliance contract.
type RegistrymdcomplianceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrymdcomplianceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcomplianceOwnershipTransferredIterator{contract: _Registrymdcompliance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrymdcomplianceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcomplianceOwnershipTransferred)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrymdcomplianceOwnershipTransferred, error) {
	event := new(RegistrymdcomplianceOwnershipTransferred)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcompliancePropertyRegisteredIterator is returned from FilterPropertyRegistered and is used to iterate over the raw logs and unpacked data for PropertyRegistered events raised by the Registrymdcompliance contract.
type RegistrymdcompliancePropertyRegisteredIterator struct {
	Event *RegistrymdcompliancePropertyRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrymdcompliancePropertyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcompliancePropertyRegistered)
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
		it.Event = new(RegistrymdcompliancePropertyRegistered)
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
func (it *RegistrymdcompliancePropertyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcompliancePropertyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcompliancePropertyRegistered represents a PropertyRegistered event raised by the Registrymdcompliance contract.
type RegistrymdcompliancePropertyRegistered struct {
	MatriculaId  *big.Int
	Proprietario common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPropertyRegistered is a free log retrieval operation binding the contract event 0xf16e2a0ca5373d3b3c5e53640bf2953ea25184f229ba20d421b5fc6a635e3945.
//
// Solidity: event PropertyRegistered(uint256 indexed matriculaId, address indexed proprietario)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterPropertyRegistered(opts *bind.FilterOpts, matriculaId []*big.Int, proprietario []common.Address) (*RegistrymdcompliancePropertyRegisteredIterator, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}
	var proprietarioRule []interface{}
	for _, proprietarioItem := range proprietario {
		proprietarioRule = append(proprietarioRule, proprietarioItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "PropertyRegistered", matriculaIdRule, proprietarioRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcompliancePropertyRegisteredIterator{contract: _Registrymdcompliance.contract, event: "PropertyRegistered", logs: logs, sub: sub}, nil
}

// WatchPropertyRegistered is a free log subscription operation binding the contract event 0xf16e2a0ca5373d3b3c5e53640bf2953ea25184f229ba20d421b5fc6a635e3945.
//
// Solidity: event PropertyRegistered(uint256 indexed matriculaId, address indexed proprietario)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchPropertyRegistered(opts *bind.WatchOpts, sink chan<- *RegistrymdcompliancePropertyRegistered, matriculaId []*big.Int, proprietario []common.Address) (event.Subscription, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}
	var proprietarioRule []interface{}
	for _, proprietarioItem := range proprietario {
		proprietarioRule = append(proprietarioRule, proprietarioItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "PropertyRegistered", matriculaIdRule, proprietarioRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcompliancePropertyRegistered)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "PropertyRegistered", log); err != nil {
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

// ParsePropertyRegistered is a log parse operation binding the contract event 0xf16e2a0ca5373d3b3c5e53640bf2953ea25184f229ba20d421b5fc6a635e3945.
//
// Solidity: event PropertyRegistered(uint256 indexed matriculaId, address indexed proprietario)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParsePropertyRegistered(log types.Log) (*RegistrymdcompliancePropertyRegistered, error) {
	event := new(RegistrymdcompliancePropertyRegistered)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "PropertyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcompliancePropertyTransferredIterator is returned from FilterPropertyTransferred and is used to iterate over the raw logs and unpacked data for PropertyTransferred events raised by the Registrymdcompliance contract.
type RegistrymdcompliancePropertyTransferredIterator struct {
	Event *RegistrymdcompliancePropertyTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrymdcompliancePropertyTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcompliancePropertyTransferred)
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
		it.Event = new(RegistrymdcompliancePropertyTransferred)
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
func (it *RegistrymdcompliancePropertyTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcompliancePropertyTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcompliancePropertyTransferred represents a PropertyTransferred event raised by the Registrymdcompliance contract.
type RegistrymdcompliancePropertyTransferred struct {
	MatriculaId *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPropertyTransferred is a free log retrieval operation binding the contract event 0x2231fc8bc2c20e26d85d66dc7f13a6b6d535c86d2a4d678bb2de61f5b7136083.
//
// Solidity: event PropertyTransferred(uint256 indexed matriculaId, address indexed from, address indexed to)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterPropertyTransferred(opts *bind.FilterOpts, matriculaId []*big.Int, from []common.Address, to []common.Address) (*RegistrymdcompliancePropertyTransferredIterator, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "PropertyTransferred", matriculaIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcompliancePropertyTransferredIterator{contract: _Registrymdcompliance.contract, event: "PropertyTransferred", logs: logs, sub: sub}, nil
}

// WatchPropertyTransferred is a free log subscription operation binding the contract event 0x2231fc8bc2c20e26d85d66dc7f13a6b6d535c86d2a4d678bb2de61f5b7136083.
//
// Solidity: event PropertyTransferred(uint256 indexed matriculaId, address indexed from, address indexed to)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchPropertyTransferred(opts *bind.WatchOpts, sink chan<- *RegistrymdcompliancePropertyTransferred, matriculaId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "PropertyTransferred", matriculaIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcompliancePropertyTransferred)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "PropertyTransferred", log); err != nil {
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

// ParsePropertyTransferred is a log parse operation binding the contract event 0x2231fc8bc2c20e26d85d66dc7f13a6b6d535c86d2a4d678bb2de61f5b7136083.
//
// Solidity: event PropertyTransferred(uint256 indexed matriculaId, address indexed from, address indexed to)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParsePropertyTransferred(log types.Log) (*RegistrymdcompliancePropertyTransferred, error) {
	event := new(RegistrymdcompliancePropertyTransferred)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "PropertyTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrymdcompliancePropertyUpdatedIterator is returned from FilterPropertyUpdated and is used to iterate over the raw logs and unpacked data for PropertyUpdated events raised by the Registrymdcompliance contract.
type RegistrymdcompliancePropertyUpdatedIterator struct {
	Event *RegistrymdcompliancePropertyUpdated // Event containing the contract specifics and raw log

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
func (it *RegistrymdcompliancePropertyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrymdcompliancePropertyUpdated)
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
		it.Event = new(RegistrymdcompliancePropertyUpdated)
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
func (it *RegistrymdcompliancePropertyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrymdcompliancePropertyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrymdcompliancePropertyUpdated represents a PropertyUpdated event raised by the Registrymdcompliance contract.
type RegistrymdcompliancePropertyUpdated struct {
	MatriculaId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPropertyUpdated is a free log retrieval operation binding the contract event 0x7ad43e2b215acc30fe68db10cbf57c6b2b699ef7572e54808c2a106fb904f3fb.
//
// Solidity: event PropertyUpdated(uint256 indexed matriculaId)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) FilterPropertyUpdated(opts *bind.FilterOpts, matriculaId []*big.Int) (*RegistrymdcompliancePropertyUpdatedIterator, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.FilterLogs(opts, "PropertyUpdated", matriculaIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistrymdcompliancePropertyUpdatedIterator{contract: _Registrymdcompliance.contract, event: "PropertyUpdated", logs: logs, sub: sub}, nil
}

// WatchPropertyUpdated is a free log subscription operation binding the contract event 0x7ad43e2b215acc30fe68db10cbf57c6b2b699ef7572e54808c2a106fb904f3fb.
//
// Solidity: event PropertyUpdated(uint256 indexed matriculaId)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) WatchPropertyUpdated(opts *bind.WatchOpts, sink chan<- *RegistrymdcompliancePropertyUpdated, matriculaId []*big.Int) (event.Subscription, error) {

	var matriculaIdRule []interface{}
	for _, matriculaIdItem := range matriculaId {
		matriculaIdRule = append(matriculaIdRule, matriculaIdItem)
	}

	logs, sub, err := _Registrymdcompliance.contract.WatchLogs(opts, "PropertyUpdated", matriculaIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrymdcompliancePropertyUpdated)
				if err := _Registrymdcompliance.contract.UnpackLog(event, "PropertyUpdated", log); err != nil {
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

// ParsePropertyUpdated is a log parse operation binding the contract event 0x7ad43e2b215acc30fe68db10cbf57c6b2b699ef7572e54808c2a106fb904f3fb.
//
// Solidity: event PropertyUpdated(uint256 indexed matriculaId)
func (_Registrymdcompliance *RegistrymdcomplianceFilterer) ParsePropertyUpdated(log types.Log) (*RegistrymdcompliancePropertyUpdated, error) {
	event := new(RegistrymdcompliancePropertyUpdated)
	if err := _Registrymdcompliance.contract.UnpackLog(event, "PropertyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
