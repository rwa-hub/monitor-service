// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchBurn\",\"inputs\":[{\"name\":\"_userAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchForcedTransfer\",\"inputs\":[{\"name\":\"_fromList\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_toList\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchFreezePartialTokens\",\"inputs\":[{\"name\":\"_userAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchMint\",\"inputs\":[{\"name\":\"_toList\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchSetAddressFrozen\",\"inputs\":[{\"name\":\"_userAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_freeze\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchTransfer\",\"inputs\":[{\"name\":\"_toList\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchUnfreezePartialTokens\",\"inputs\":[{\"name\":\"_userAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"compliance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIModularCompliance\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forcedTransfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezePartialTokens\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFrozenTokens\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentityRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_onchainID\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoveryAddress\",\"inputs\":[{\"name\":\"_lostWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_investorOnchainID\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddressFrozen\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_freeze\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCompliance\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIdentityRegistry\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOnchainID\",\"inputs\":[{\"name\":\"_onchainID\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unfreezePartialTokens\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AddressFrozen\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_isFrozen\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentAdded\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentRemoved\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceAdded\",\"inputs\":[{\"name\":\"_compliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRegistryAdded\",\"inputs\":[{\"name\":\"_identityRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecoverySuccess\",\"inputs\":[{\"name\":\"_lostWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_newWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_investorOnchainID\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensFrozen\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensUnfrozen\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedTokenInformation\",\"inputs\":[{\"name\":\"_newName\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"_newSymbol\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"_newDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"_newVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"_newOnchainID\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Token *TokenSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _userAddress)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _userAddress) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(_userAddress common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _userAddress)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_Token *TokenCaller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "compliance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_Token *TokenSession) Compliance() (common.Address, error) {
	return _Token.Contract.Compliance(&_Token.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() view returns(address)
func (_Token *TokenCallerSession) Compliance() (common.Address, error) {
	return _Token.Contract.Compliance(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// GetFrozenTokens is a free data retrieval call binding the contract method 0x158b1a57.
//
// Solidity: function getFrozenTokens(address _userAddress) view returns(uint256)
func (_Token *TokenCaller) GetFrozenTokens(opts *bind.CallOpts, _userAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getFrozenTokens", _userAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFrozenTokens is a free data retrieval call binding the contract method 0x158b1a57.
//
// Solidity: function getFrozenTokens(address _userAddress) view returns(uint256)
func (_Token *TokenSession) GetFrozenTokens(_userAddress common.Address) (*big.Int, error) {
	return _Token.Contract.GetFrozenTokens(&_Token.CallOpts, _userAddress)
}

// GetFrozenTokens is a free data retrieval call binding the contract method 0x158b1a57.
//
// Solidity: function getFrozenTokens(address _userAddress) view returns(uint256)
func (_Token *TokenCallerSession) GetFrozenTokens(_userAddress common.Address) (*big.Int, error) {
	return _Token.Contract.GetFrozenTokens(&_Token.CallOpts, _userAddress)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_Token *TokenCaller) IdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "identityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_Token *TokenSession) IdentityRegistry() (common.Address, error) {
	return _Token.Contract.IdentityRegistry(&_Token.CallOpts)
}

// IdentityRegistry is a free data retrieval call binding the contract method 0x134e18f4.
//
// Solidity: function identityRegistry() view returns(address)
func (_Token *TokenCallerSession) IdentityRegistry() (common.Address, error) {
	return _Token.Contract.IdentityRegistry(&_Token.CallOpts)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Token *TokenCaller) IsAgent(opts *bind.CallOpts, _agent common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isAgent", _agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Token *TokenSession) IsAgent(_agent common.Address) (bool, error) {
	return _Token.Contract.IsAgent(&_Token.CallOpts, _agent)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address _agent) view returns(bool)
func (_Token *TokenCallerSession) IsAgent(_agent common.Address) (bool, error) {
	return _Token.Contract.IsAgent(&_Token.CallOpts, _agent)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address _userAddress) view returns(bool)
func (_Token *TokenCaller) IsFrozen(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isFrozen", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address _userAddress) view returns(bool)
func (_Token *TokenSession) IsFrozen(_userAddress common.Address) (bool, error) {
	return _Token.Contract.IsFrozen(&_Token.CallOpts, _userAddress)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address _userAddress) view returns(bool)
func (_Token *TokenCallerSession) IsFrozen(_userAddress common.Address) (bool, error) {
	return _Token.Contract.IsFrozen(&_Token.CallOpts, _userAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// OnchainID is a free data retrieval call binding the contract method 0xaba63705.
//
// Solidity: function onchainID() view returns(address)
func (_Token *TokenCaller) OnchainID(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "onchainID")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnchainID is a free data retrieval call binding the contract method 0xaba63705.
//
// Solidity: function onchainID() view returns(address)
func (_Token *TokenSession) OnchainID() (common.Address, error) {
	return _Token.Contract.OnchainID(&_Token.CallOpts)
}

// OnchainID is a free data retrieval call binding the contract method 0xaba63705.
//
// Solidity: function onchainID() view returns(address)
func (_Token *TokenCallerSession) OnchainID() (common.Address, error) {
	return _Token.Contract.OnchainID(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenCallerSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Token *TokenCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Token *TokenSession) Version() (string, error) {
	return _Token.Contract.Version(&_Token.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Token *TokenCallerSession) Version() (string, error) {
	return _Token.Contract.Version(&_Token.CallOpts)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Token *TokenTransactor) AddAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addAgent", _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Token *TokenSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddAgent(&_Token.TransactOpts, _agent)
}

// AddAgent is a paid mutator transaction binding the contract method 0x84e79842.
//
// Solidity: function addAgent(address _agent) returns()
func (_Token *TokenTransactorSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddAgent(&_Token.TransactOpts, _agent)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Token *TokenSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _amount)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenTransactor) BatchBurn(opts *bind.TransactOpts, _userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchBurn", _userAddresses, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenSession) BatchBurn(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchBurn(&_Token.TransactOpts, _userAddresses, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0x4a6cc677.
//
// Solidity: function batchBurn(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenTransactorSession) BatchBurn(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchBurn(&_Token.TransactOpts, _userAddresses, _amounts)
}

// BatchForcedTransfer is a paid mutator transaction binding the contract method 0x42a47abc.
//
// Solidity: function batchForcedTransfer(address[] _fromList, address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenTransactor) BatchForcedTransfer(opts *bind.TransactOpts, _fromList []common.Address, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchForcedTransfer", _fromList, _toList, _amounts)
}

// BatchForcedTransfer is a paid mutator transaction binding the contract method 0x42a47abc.
//
// Solidity: function batchForcedTransfer(address[] _fromList, address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenSession) BatchForcedTransfer(_fromList []common.Address, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchForcedTransfer(&_Token.TransactOpts, _fromList, _toList, _amounts)
}

// BatchForcedTransfer is a paid mutator transaction binding the contract method 0x42a47abc.
//
// Solidity: function batchForcedTransfer(address[] _fromList, address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenTransactorSession) BatchForcedTransfer(_fromList []common.Address, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchForcedTransfer(&_Token.TransactOpts, _fromList, _toList, _amounts)
}

// BatchFreezePartialTokens is a paid mutator transaction binding the contract method 0xfc7e5fa8.
//
// Solidity: function batchFreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenTransactor) BatchFreezePartialTokens(opts *bind.TransactOpts, _userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchFreezePartialTokens", _userAddresses, _amounts)
}

// BatchFreezePartialTokens is a paid mutator transaction binding the contract method 0xfc7e5fa8.
//
// Solidity: function batchFreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenSession) BatchFreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchFreezePartialTokens(&_Token.TransactOpts, _userAddresses, _amounts)
}

// BatchFreezePartialTokens is a paid mutator transaction binding the contract method 0xfc7e5fa8.
//
// Solidity: function batchFreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenTransactorSession) BatchFreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchFreezePartialTokens(&_Token.TransactOpts, _userAddresses, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenTransactor) BatchMint(opts *bind.TransactOpts, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchMint", _toList, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenSession) BatchMint(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchMint(&_Token.TransactOpts, _toList, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenTransactorSession) BatchMint(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchMint(&_Token.TransactOpts, _toList, _amounts)
}

// BatchSetAddressFrozen is a paid mutator transaction binding the contract method 0x1a7af379.
//
// Solidity: function batchSetAddressFrozen(address[] _userAddresses, bool[] _freeze) returns()
func (_Token *TokenTransactor) BatchSetAddressFrozen(opts *bind.TransactOpts, _userAddresses []common.Address, _freeze []bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchSetAddressFrozen", _userAddresses, _freeze)
}

// BatchSetAddressFrozen is a paid mutator transaction binding the contract method 0x1a7af379.
//
// Solidity: function batchSetAddressFrozen(address[] _userAddresses, bool[] _freeze) returns()
func (_Token *TokenSession) BatchSetAddressFrozen(_userAddresses []common.Address, _freeze []bool) (*types.Transaction, error) {
	return _Token.Contract.BatchSetAddressFrozen(&_Token.TransactOpts, _userAddresses, _freeze)
}

// BatchSetAddressFrozen is a paid mutator transaction binding the contract method 0x1a7af379.
//
// Solidity: function batchSetAddressFrozen(address[] _userAddresses, bool[] _freeze) returns()
func (_Token *TokenTransactorSession) BatchSetAddressFrozen(_userAddresses []common.Address, _freeze []bool) (*types.Transaction, error) {
	return _Token.Contract.BatchSetAddressFrozen(&_Token.TransactOpts, _userAddresses, _freeze)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenTransactor) BatchTransfer(opts *bind.TransactOpts, _toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchTransfer", _toList, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenSession) BatchTransfer(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchTransfer(&_Token.TransactOpts, _toList, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _toList, uint256[] _amounts) returns()
func (_Token *TokenTransactorSession) BatchTransfer(_toList []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchTransfer(&_Token.TransactOpts, _toList, _amounts)
}

// BatchUnfreezePartialTokens is a paid mutator transaction binding the contract method 0x4710362d.
//
// Solidity: function batchUnfreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenTransactor) BatchUnfreezePartialTokens(opts *bind.TransactOpts, _userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "batchUnfreezePartialTokens", _userAddresses, _amounts)
}

// BatchUnfreezePartialTokens is a paid mutator transaction binding the contract method 0x4710362d.
//
// Solidity: function batchUnfreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenSession) BatchUnfreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchUnfreezePartialTokens(&_Token.TransactOpts, _userAddresses, _amounts)
}

// BatchUnfreezePartialTokens is a paid mutator transaction binding the contract method 0x4710362d.
//
// Solidity: function batchUnfreezePartialTokens(address[] _userAddresses, uint256[] _amounts) returns()
func (_Token *TokenTransactorSession) BatchUnfreezePartialTokens(_userAddresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.BatchUnfreezePartialTokens(&_Token.TransactOpts, _userAddresses, _amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _userAddress, uint256 _amount) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", _userAddress, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _userAddress, uint256 _amount) returns()
func (_Token *TokenSession) Burn(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _userAddress, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _userAddress, uint256 _amount) returns()
func (_Token *TokenTransactorSession) Burn(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, _userAddress, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Token *TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Token *TokenSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Token *TokenTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, _spender, _subtractedValue)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_Token *TokenTransactor) ForcedTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "forcedTransfer", _from, _to, _amount)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_Token *TokenSession) ForcedTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ForcedTransfer(&_Token.TransactOpts, _from, _to, _amount)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_Token *TokenTransactorSession) ForcedTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ForcedTransfer(&_Token.TransactOpts, _from, _to, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Token *TokenTransactor) FreezePartialTokens(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "freezePartialTokens", _userAddress, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Token *TokenSession) FreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FreezePartialTokens(&_Token.TransactOpts, _userAddress, _amount)
}

// FreezePartialTokens is a paid mutator transaction binding the contract method 0x125c4a33.
//
// Solidity: function freezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Token *TokenTransactorSession) FreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.FreezePartialTokens(&_Token.TransactOpts, _userAddress, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Token *TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Token *TokenSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Token *TokenTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, _spender, _addedValue)
}

// Init is a paid mutator transaction binding the contract method 0xf91b619c.
//
// Solidity: function init(address _identityRegistry, address _compliance, string _name, string _symbol, uint8 _decimals, address _onchainID) returns()
func (_Token *TokenTransactor) Init(opts *bind.TransactOpts, _identityRegistry common.Address, _compliance common.Address, _name string, _symbol string, _decimals uint8, _onchainID common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "init", _identityRegistry, _compliance, _name, _symbol, _decimals, _onchainID)
}

// Init is a paid mutator transaction binding the contract method 0xf91b619c.
//
// Solidity: function init(address _identityRegistry, address _compliance, string _name, string _symbol, uint8 _decimals, address _onchainID) returns()
func (_Token *TokenSession) Init(_identityRegistry common.Address, _compliance common.Address, _name string, _symbol string, _decimals uint8, _onchainID common.Address) (*types.Transaction, error) {
	return _Token.Contract.Init(&_Token.TransactOpts, _identityRegistry, _compliance, _name, _symbol, _decimals, _onchainID)
}

// Init is a paid mutator transaction binding the contract method 0xf91b619c.
//
// Solidity: function init(address _identityRegistry, address _compliance, string _name, string _symbol, uint8 _decimals, address _onchainID) returns()
func (_Token *TokenTransactorSession) Init(_identityRegistry common.Address, _compliance common.Address, _name string, _symbol string, _decimals uint8, _onchainID common.Address) (*types.Transaction, error) {
	return _Token.Contract.Init(&_Token.TransactOpts, _identityRegistry, _compliance, _name, _symbol, _decimals, _onchainID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Token *TokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Token *TokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// RecoveryAddress is a paid mutator transaction binding the contract method 0x9285948a.
//
// Solidity: function recoveryAddress(address _lostWallet, address _newWallet, address _investorOnchainID) returns(bool)
func (_Token *TokenTransactor) RecoveryAddress(opts *bind.TransactOpts, _lostWallet common.Address, _newWallet common.Address, _investorOnchainID common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "recoveryAddress", _lostWallet, _newWallet, _investorOnchainID)
}

// RecoveryAddress is a paid mutator transaction binding the contract method 0x9285948a.
//
// Solidity: function recoveryAddress(address _lostWallet, address _newWallet, address _investorOnchainID) returns(bool)
func (_Token *TokenSession) RecoveryAddress(_lostWallet common.Address, _newWallet common.Address, _investorOnchainID common.Address) (*types.Transaction, error) {
	return _Token.Contract.RecoveryAddress(&_Token.TransactOpts, _lostWallet, _newWallet, _investorOnchainID)
}

// RecoveryAddress is a paid mutator transaction binding the contract method 0x9285948a.
//
// Solidity: function recoveryAddress(address _lostWallet, address _newWallet, address _investorOnchainID) returns(bool)
func (_Token *TokenTransactorSession) RecoveryAddress(_lostWallet common.Address, _newWallet common.Address, _investorOnchainID common.Address) (*types.Transaction, error) {
	return _Token.Contract.RecoveryAddress(&_Token.TransactOpts, _lostWallet, _newWallet, _investorOnchainID)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Token *TokenTransactor) RemoveAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeAgent", _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Token *TokenSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveAgent(&_Token.TransactOpts, _agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address _agent) returns()
func (_Token *TokenTransactorSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveAgent(&_Token.TransactOpts, _agent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_Token *TokenTransactor) SetAddressFrozen(opts *bind.TransactOpts, _userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAddressFrozen", _userAddress, _freeze)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_Token *TokenSession) SetAddressFrozen(_userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _Token.Contract.SetAddressFrozen(&_Token.TransactOpts, _userAddress, _freeze)
}

// SetAddressFrozen is a paid mutator transaction binding the contract method 0xc69c09cf.
//
// Solidity: function setAddressFrozen(address _userAddress, bool _freeze) returns()
func (_Token *TokenTransactorSession) SetAddressFrozen(_userAddress common.Address, _freeze bool) (*types.Transaction, error) {
	return _Token.Contract.SetAddressFrozen(&_Token.TransactOpts, _userAddress, _freeze)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _compliance) returns()
func (_Token *TokenTransactor) SetCompliance(opts *bind.TransactOpts, _compliance common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setCompliance", _compliance)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _compliance) returns()
func (_Token *TokenSession) SetCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCompliance(&_Token.TransactOpts, _compliance)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _compliance) returns()
func (_Token *TokenTransactorSession) SetCompliance(_compliance common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetCompliance(&_Token.TransactOpts, _compliance)
}

// SetIdentityRegistry is a paid mutator transaction binding the contract method 0xcbf3f861.
//
// Solidity: function setIdentityRegistry(address _identityRegistry) returns()
func (_Token *TokenTransactor) SetIdentityRegistry(opts *bind.TransactOpts, _identityRegistry common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setIdentityRegistry", _identityRegistry)
}

// SetIdentityRegistry is a paid mutator transaction binding the contract method 0xcbf3f861.
//
// Solidity: function setIdentityRegistry(address _identityRegistry) returns()
func (_Token *TokenSession) SetIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetIdentityRegistry(&_Token.TransactOpts, _identityRegistry)
}

// SetIdentityRegistry is a paid mutator transaction binding the contract method 0xcbf3f861.
//
// Solidity: function setIdentityRegistry(address _identityRegistry) returns()
func (_Token *TokenTransactorSession) SetIdentityRegistry(_identityRegistry common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetIdentityRegistry(&_Token.TransactOpts, _identityRegistry)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Token *TokenTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Token *TokenSession) SetName(_name string) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Token *TokenTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, _name)
}

// SetOnchainID is a paid mutator transaction binding the contract method 0x3d1ddc5b.
//
// Solidity: function setOnchainID(address _onchainID) returns()
func (_Token *TokenTransactor) SetOnchainID(opts *bind.TransactOpts, _onchainID common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setOnchainID", _onchainID)
}

// SetOnchainID is a paid mutator transaction binding the contract method 0x3d1ddc5b.
//
// Solidity: function setOnchainID(address _onchainID) returns()
func (_Token *TokenSession) SetOnchainID(_onchainID common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetOnchainID(&_Token.TransactOpts, _onchainID)
}

// SetOnchainID is a paid mutator transaction binding the contract method 0x3d1ddc5b.
//
// Solidity: function setOnchainID(address _onchainID) returns()
func (_Token *TokenTransactorSession) SetOnchainID(_onchainID common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetOnchainID(&_Token.TransactOpts, _onchainID)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_Token *TokenTransactor) SetSymbol(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSymbol", _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_Token *TokenSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _Token.Contract.SetSymbol(&_Token.TransactOpts, _symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string _symbol) returns()
func (_Token *TokenTransactorSession) SetSymbol(_symbol string) (*types.Transaction, error) {
	return _Token.Contract.SetSymbol(&_Token.TransactOpts, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Token *TokenSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Token *TokenTransactor) UnfreezePartialTokens(opts *bind.TransactOpts, _userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unfreezePartialTokens", _userAddress, _amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Token *TokenSession) UnfreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UnfreezePartialTokens(&_Token.TransactOpts, _userAddress, _amount)
}

// UnfreezePartialTokens is a paid mutator transaction binding the contract method 0x1fe56f7d.
//
// Solidity: function unfreezePartialTokens(address _userAddress, uint256 _amount) returns()
func (_Token *TokenTransactorSession) UnfreezePartialTokens(_userAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UnfreezePartialTokens(&_Token.TransactOpts, _userAddress, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// TokenAddressFrozenIterator is returned from FilterAddressFrozen and is used to iterate over the raw logs and unpacked data for AddressFrozen events raised by the Token contract.
type TokenAddressFrozenIterator struct {
	Event *TokenAddressFrozen // Event containing the contract specifics and raw log

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
func (it *TokenAddressFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddressFrozen)
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
		it.Event = new(TokenAddressFrozen)
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
func (it *TokenAddressFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddressFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddressFrozen represents a AddressFrozen event raised by the Token contract.
type TokenAddressFrozen struct {
	UserAddress common.Address
	IsFrozen    bool
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddressFrozen is a free log retrieval operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address indexed _userAddress, bool indexed _isFrozen, address indexed _owner)
func (_Token *TokenFilterer) FilterAddressFrozen(opts *bind.FilterOpts, _userAddress []common.Address, _isFrozen []bool, _owner []common.Address) (*TokenAddressFrozenIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}
	var _isFrozenRule []interface{}
	for _, _isFrozenItem := range _isFrozen {
		_isFrozenRule = append(_isFrozenRule, _isFrozenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AddressFrozen", _userAddressRule, _isFrozenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &TokenAddressFrozenIterator{contract: _Token.contract, event: "AddressFrozen", logs: logs, sub: sub}, nil
}

// WatchAddressFrozen is a free log subscription operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address indexed _userAddress, bool indexed _isFrozen, address indexed _owner)
func (_Token *TokenFilterer) WatchAddressFrozen(opts *bind.WatchOpts, sink chan<- *TokenAddressFrozen, _userAddress []common.Address, _isFrozen []bool, _owner []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}
	var _isFrozenRule []interface{}
	for _, _isFrozenItem := range _isFrozen {
		_isFrozenRule = append(_isFrozenRule, _isFrozenItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AddressFrozen", _userAddressRule, _isFrozenRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddressFrozen)
				if err := _Token.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
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

// ParseAddressFrozen is a log parse operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: event AddressFrozen(address indexed _userAddress, bool indexed _isFrozen, address indexed _owner)
func (_Token *TokenFilterer) ParseAddressFrozen(log types.Log) (*TokenAddressFrozen, error) {
	event := new(TokenAddressFrozen)
	if err := _Token.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the Token contract.
type TokenAgentAddedIterator struct {
	Event *TokenAgentAdded // Event containing the contract specifics and raw log

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
func (it *TokenAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAgentAdded)
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
		it.Event = new(TokenAgentAdded)
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
func (it *TokenAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAgentAdded represents a AgentAdded event raised by the Token contract.
type TokenAgentAdded struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Token *TokenFilterer) FilterAgentAdded(opts *bind.FilterOpts, _agent []common.Address) (*TokenAgentAddedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return &TokenAgentAddedIterator{contract: _Token.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf68e73cec97f2d70aa641fb26e87a4383686e2efacb648f2165aeb02ac562ec5.
//
// Solidity: event AgentAdded(address indexed _agent)
func (_Token *TokenFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *TokenAgentAdded, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AgentAdded", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAgentAdded)
				if err := _Token.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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
func (_Token *TokenFilterer) ParseAgentAdded(log types.Log) (*TokenAgentAdded, error) {
	event := new(TokenAgentAdded)
	if err := _Token.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the Token contract.
type TokenAgentRemovedIterator struct {
	Event *TokenAgentRemoved // Event containing the contract specifics and raw log

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
func (it *TokenAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAgentRemoved)
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
		it.Event = new(TokenAgentRemoved)
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
func (it *TokenAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAgentRemoved represents a AgentRemoved event raised by the Token contract.
type TokenAgentRemoved struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Token *TokenFilterer) FilterAgentRemoved(opts *bind.FilterOpts, _agent []common.Address) (*TokenAgentRemovedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return &TokenAgentRemovedIterator{contract: _Token.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0xed9c8ad8d5a0a66898ea49d2956929c93ae2e8bd50281b2ed897c5d1a6737e0b.
//
// Solidity: event AgentRemoved(address indexed _agent)
func (_Token *TokenFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *TokenAgentRemoved, _agent []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AgentRemoved", _agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAgentRemoved)
				if err := _Token.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_Token *TokenFilterer) ParseAgentRemoved(log types.Log) (*TokenAgentRemoved, error) {
	event := new(TokenAgentRemoved)
	if err := _Token.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenComplianceAddedIterator is returned from FilterComplianceAdded and is used to iterate over the raw logs and unpacked data for ComplianceAdded events raised by the Token contract.
type TokenComplianceAddedIterator struct {
	Event *TokenComplianceAdded // Event containing the contract specifics and raw log

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
func (it *TokenComplianceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenComplianceAdded)
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
		it.Event = new(TokenComplianceAdded)
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
func (it *TokenComplianceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenComplianceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenComplianceAdded represents a ComplianceAdded event raised by the Token contract.
type TokenComplianceAdded struct {
	Compliance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterComplianceAdded is a free log retrieval operation binding the contract event 0x7f3a888862559648ec01d97deb7b5012bff86dc91e654a1de397170db40e35b6.
//
// Solidity: event ComplianceAdded(address indexed _compliance)
func (_Token *TokenFilterer) FilterComplianceAdded(opts *bind.FilterOpts, _compliance []common.Address) (*TokenComplianceAddedIterator, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ComplianceAdded", _complianceRule)
	if err != nil {
		return nil, err
	}
	return &TokenComplianceAddedIterator{contract: _Token.contract, event: "ComplianceAdded", logs: logs, sub: sub}, nil
}

// WatchComplianceAdded is a free log subscription operation binding the contract event 0x7f3a888862559648ec01d97deb7b5012bff86dc91e654a1de397170db40e35b6.
//
// Solidity: event ComplianceAdded(address indexed _compliance)
func (_Token *TokenFilterer) WatchComplianceAdded(opts *bind.WatchOpts, sink chan<- *TokenComplianceAdded, _compliance []common.Address) (event.Subscription, error) {

	var _complianceRule []interface{}
	for _, _complianceItem := range _compliance {
		_complianceRule = append(_complianceRule, _complianceItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ComplianceAdded", _complianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenComplianceAdded)
				if err := _Token.contract.UnpackLog(event, "ComplianceAdded", log); err != nil {
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

// ParseComplianceAdded is a log parse operation binding the contract event 0x7f3a888862559648ec01d97deb7b5012bff86dc91e654a1de397170db40e35b6.
//
// Solidity: event ComplianceAdded(address indexed _compliance)
func (_Token *TokenFilterer) ParseComplianceAdded(log types.Log) (*TokenComplianceAdded, error) {
	event := new(TokenComplianceAdded)
	if err := _Token.contract.UnpackLog(event, "ComplianceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenIdentityRegistryAddedIterator is returned from FilterIdentityRegistryAdded and is used to iterate over the raw logs and unpacked data for IdentityRegistryAdded events raised by the Token contract.
type TokenIdentityRegistryAddedIterator struct {
	Event *TokenIdentityRegistryAdded // Event containing the contract specifics and raw log

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
func (it *TokenIdentityRegistryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenIdentityRegistryAdded)
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
		it.Event = new(TokenIdentityRegistryAdded)
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
func (it *TokenIdentityRegistryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenIdentityRegistryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenIdentityRegistryAdded represents a IdentityRegistryAdded event raised by the Token contract.
type TokenIdentityRegistryAdded struct {
	IdentityRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIdentityRegistryAdded is a free log retrieval operation binding the contract event 0xd2be862d755bca7e0d39772b2cab3a5578da9c285f69199f4c063c2294a7f36c.
//
// Solidity: event IdentityRegistryAdded(address indexed _identityRegistry)
func (_Token *TokenFilterer) FilterIdentityRegistryAdded(opts *bind.FilterOpts, _identityRegistry []common.Address) (*TokenIdentityRegistryAddedIterator, error) {

	var _identityRegistryRule []interface{}
	for _, _identityRegistryItem := range _identityRegistry {
		_identityRegistryRule = append(_identityRegistryRule, _identityRegistryItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "IdentityRegistryAdded", _identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return &TokenIdentityRegistryAddedIterator{contract: _Token.contract, event: "IdentityRegistryAdded", logs: logs, sub: sub}, nil
}

// WatchIdentityRegistryAdded is a free log subscription operation binding the contract event 0xd2be862d755bca7e0d39772b2cab3a5578da9c285f69199f4c063c2294a7f36c.
//
// Solidity: event IdentityRegistryAdded(address indexed _identityRegistry)
func (_Token *TokenFilterer) WatchIdentityRegistryAdded(opts *bind.WatchOpts, sink chan<- *TokenIdentityRegistryAdded, _identityRegistry []common.Address) (event.Subscription, error) {

	var _identityRegistryRule []interface{}
	for _, _identityRegistryItem := range _identityRegistry {
		_identityRegistryRule = append(_identityRegistryRule, _identityRegistryItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "IdentityRegistryAdded", _identityRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenIdentityRegistryAdded)
				if err := _Token.contract.UnpackLog(event, "IdentityRegistryAdded", log); err != nil {
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

// ParseIdentityRegistryAdded is a log parse operation binding the contract event 0xd2be862d755bca7e0d39772b2cab3a5578da9c285f69199f4c063c2294a7f36c.
//
// Solidity: event IdentityRegistryAdded(address indexed _identityRegistry)
func (_Token *TokenFilterer) ParseIdentityRegistryAdded(log types.Log) (*TokenIdentityRegistryAdded, error) {
	event := new(TokenIdentityRegistryAdded)
	if err := _Token.contract.UnpackLog(event, "IdentityRegistryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Token contract.
type TokenInitializedIterator struct {
	Event *TokenInitialized // Event containing the contract specifics and raw log

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
func (it *TokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenInitialized)
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
		it.Event = new(TokenInitialized)
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
func (it *TokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenInitialized represents a Initialized event raised by the Token contract.
type TokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Token *TokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenInitializedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenInitializedIterator{contract: _Token.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Token *TokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenInitialized) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenInitialized)
				if err := _Token.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Token *TokenFilterer) ParseInitialized(log types.Log) (*TokenInitialized, error) {
	event := new(TokenInitialized)
	if err := _Token.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Token contract.
type TokenPausedIterator struct {
	Event *TokenPaused // Event containing the contract specifics and raw log

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
func (it *TokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaused)
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
		it.Event = new(TokenPaused)
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
func (it *TokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaused represents a Paused event raised by the Token contract.
type TokenPaused struct {
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address _userAddress)
func (_Token *TokenFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenPausedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenPausedIterator{contract: _Token.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address _userAddress)
func (_Token *TokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenPaused) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaused)
				if err := _Token.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address _userAddress)
func (_Token *TokenFilterer) ParsePaused(log types.Log) (*TokenPaused, error) {
	event := new(TokenPaused)
	if err := _Token.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRecoverySuccessIterator is returned from FilterRecoverySuccess and is used to iterate over the raw logs and unpacked data for RecoverySuccess events raised by the Token contract.
type TokenRecoverySuccessIterator struct {
	Event *TokenRecoverySuccess // Event containing the contract specifics and raw log

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
func (it *TokenRecoverySuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRecoverySuccess)
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
		it.Event = new(TokenRecoverySuccess)
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
func (it *TokenRecoverySuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRecoverySuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRecoverySuccess represents a RecoverySuccess event raised by the Token contract.
type TokenRecoverySuccess struct {
	LostWallet        common.Address
	NewWallet         common.Address
	InvestorOnchainID common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRecoverySuccess is a free log retrieval operation binding the contract event 0xf0c9129a94f30f1caaceb63e44b9811d0a3edf1d6c23757f346093af5553fed0.
//
// Solidity: event RecoverySuccess(address indexed _lostWallet, address indexed _newWallet, address indexed _investorOnchainID)
func (_Token *TokenFilterer) FilterRecoverySuccess(opts *bind.FilterOpts, _lostWallet []common.Address, _newWallet []common.Address, _investorOnchainID []common.Address) (*TokenRecoverySuccessIterator, error) {

	var _lostWalletRule []interface{}
	for _, _lostWalletItem := range _lostWallet {
		_lostWalletRule = append(_lostWalletRule, _lostWalletItem)
	}
	var _newWalletRule []interface{}
	for _, _newWalletItem := range _newWallet {
		_newWalletRule = append(_newWalletRule, _newWalletItem)
	}
	var _investorOnchainIDRule []interface{}
	for _, _investorOnchainIDItem := range _investorOnchainID {
		_investorOnchainIDRule = append(_investorOnchainIDRule, _investorOnchainIDItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RecoverySuccess", _lostWalletRule, _newWalletRule, _investorOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TokenRecoverySuccessIterator{contract: _Token.contract, event: "RecoverySuccess", logs: logs, sub: sub}, nil
}

// WatchRecoverySuccess is a free log subscription operation binding the contract event 0xf0c9129a94f30f1caaceb63e44b9811d0a3edf1d6c23757f346093af5553fed0.
//
// Solidity: event RecoverySuccess(address indexed _lostWallet, address indexed _newWallet, address indexed _investorOnchainID)
func (_Token *TokenFilterer) WatchRecoverySuccess(opts *bind.WatchOpts, sink chan<- *TokenRecoverySuccess, _lostWallet []common.Address, _newWallet []common.Address, _investorOnchainID []common.Address) (event.Subscription, error) {

	var _lostWalletRule []interface{}
	for _, _lostWalletItem := range _lostWallet {
		_lostWalletRule = append(_lostWalletRule, _lostWalletItem)
	}
	var _newWalletRule []interface{}
	for _, _newWalletItem := range _newWallet {
		_newWalletRule = append(_newWalletRule, _newWalletItem)
	}
	var _investorOnchainIDRule []interface{}
	for _, _investorOnchainIDItem := range _investorOnchainID {
		_investorOnchainIDRule = append(_investorOnchainIDRule, _investorOnchainIDItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RecoverySuccess", _lostWalletRule, _newWalletRule, _investorOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRecoverySuccess)
				if err := _Token.contract.UnpackLog(event, "RecoverySuccess", log); err != nil {
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

// ParseRecoverySuccess is a log parse operation binding the contract event 0xf0c9129a94f30f1caaceb63e44b9811d0a3edf1d6c23757f346093af5553fed0.
//
// Solidity: event RecoverySuccess(address indexed _lostWallet, address indexed _newWallet, address indexed _investorOnchainID)
func (_Token *TokenFilterer) ParseRecoverySuccess(log types.Log) (*TokenRecoverySuccess, error) {
	event := new(TokenRecoverySuccess)
	if err := _Token.contract.UnpackLog(event, "RecoverySuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTokensFrozenIterator is returned from FilterTokensFrozen and is used to iterate over the raw logs and unpacked data for TokensFrozen events raised by the Token contract.
type TokenTokensFrozenIterator struct {
	Event *TokenTokensFrozen // Event containing the contract specifics and raw log

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
func (it *TokenTokensFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTokensFrozen)
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
		it.Event = new(TokenTokensFrozen)
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
func (it *TokenTokensFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTokensFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTokensFrozen represents a TokensFrozen event raised by the Token contract.
type TokenTokensFrozen struct {
	UserAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensFrozen is a free log retrieval operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address indexed _userAddress, uint256 _amount)
func (_Token *TokenFilterer) FilterTokensFrozen(opts *bind.FilterOpts, _userAddress []common.Address) (*TokenTokensFrozenIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "TokensFrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenTokensFrozenIterator{contract: _Token.contract, event: "TokensFrozen", logs: logs, sub: sub}, nil
}

// WatchTokensFrozen is a free log subscription operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address indexed _userAddress, uint256 _amount)
func (_Token *TokenFilterer) WatchTokensFrozen(opts *bind.WatchOpts, sink chan<- *TokenTokensFrozen, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "TokensFrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTokensFrozen)
				if err := _Token.contract.UnpackLog(event, "TokensFrozen", log); err != nil {
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

// ParseTokensFrozen is a log parse operation binding the contract event 0xa065e63c631c86f1b9f66a4a2f63f2093bf1c2168d23290259dbd969e0222a45.
//
// Solidity: event TokensFrozen(address indexed _userAddress, uint256 _amount)
func (_Token *TokenFilterer) ParseTokensFrozen(log types.Log) (*TokenTokensFrozen, error) {
	event := new(TokenTokensFrozen)
	if err := _Token.contract.UnpackLog(event, "TokensFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTokensUnfrozenIterator is returned from FilterTokensUnfrozen and is used to iterate over the raw logs and unpacked data for TokensUnfrozen events raised by the Token contract.
type TokenTokensUnfrozenIterator struct {
	Event *TokenTokensUnfrozen // Event containing the contract specifics and raw log

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
func (it *TokenTokensUnfrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTokensUnfrozen)
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
		it.Event = new(TokenTokensUnfrozen)
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
func (it *TokenTokensUnfrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTokensUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTokensUnfrozen represents a TokensUnfrozen event raised by the Token contract.
type TokenTokensUnfrozen struct {
	UserAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensUnfrozen is a free log retrieval operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address indexed _userAddress, uint256 _amount)
func (_Token *TokenFilterer) FilterTokensUnfrozen(opts *bind.FilterOpts, _userAddress []common.Address) (*TokenTokensUnfrozenIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "TokensUnfrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenTokensUnfrozenIterator{contract: _Token.contract, event: "TokensUnfrozen", logs: logs, sub: sub}, nil
}

// WatchTokensUnfrozen is a free log subscription operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address indexed _userAddress, uint256 _amount)
func (_Token *TokenFilterer) WatchTokensUnfrozen(opts *bind.WatchOpts, sink chan<- *TokenTokensUnfrozen, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "TokensUnfrozen", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTokensUnfrozen)
				if err := _Token.contract.UnpackLog(event, "TokensUnfrozen", log); err != nil {
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

// ParseTokensUnfrozen is a log parse operation binding the contract event 0x9bed35cb62ad0dba04f9d5bfee4b5bc91443e77da8a65c4c84834c51bb08b0d6.
//
// Solidity: event TokensUnfrozen(address indexed _userAddress, uint256 _amount)
func (_Token *TokenFilterer) ParseTokensUnfrozen(log types.Log) (*TokenTokensUnfrozen, error) {
	event := new(TokenTokensUnfrozen)
	if err := _Token.contract.UnpackLog(event, "TokensUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Token contract.
type TokenUnpausedIterator struct {
	Event *TokenUnpaused // Event containing the contract specifics and raw log

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
func (it *TokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUnpaused)
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
		it.Event = new(TokenUnpaused)
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
func (it *TokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUnpaused represents a Unpaused event raised by the Token contract.
type TokenUnpaused struct {
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address _userAddress)
func (_Token *TokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokenUnpausedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokenUnpausedIterator{contract: _Token.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address _userAddress)
func (_Token *TokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUnpaused)
				if err := _Token.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address _userAddress)
func (_Token *TokenFilterer) ParseUnpaused(log types.Log) (*TokenUnpaused, error) {
	event := new(TokenUnpaused)
	if err := _Token.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUpdatedTokenInformationIterator is returned from FilterUpdatedTokenInformation and is used to iterate over the raw logs and unpacked data for UpdatedTokenInformation events raised by the Token contract.
type TokenUpdatedTokenInformationIterator struct {
	Event *TokenUpdatedTokenInformation // Event containing the contract specifics and raw log

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
func (it *TokenUpdatedTokenInformationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUpdatedTokenInformation)
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
		it.Event = new(TokenUpdatedTokenInformation)
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
func (it *TokenUpdatedTokenInformationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUpdatedTokenInformationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUpdatedTokenInformation represents a UpdatedTokenInformation event raised by the Token contract.
type TokenUpdatedTokenInformation struct {
	NewName      common.Hash
	NewSymbol    common.Hash
	NewDecimals  uint8
	NewVersion   string
	NewOnchainID common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenInformation is a free log retrieval operation binding the contract event 0x6a1105ac8148a3c319adbc369f9072573e8a11d3a3d195e067e7c40767ec54d1.
//
// Solidity: event UpdatedTokenInformation(string indexed _newName, string indexed _newSymbol, uint8 _newDecimals, string _newVersion, address indexed _newOnchainID)
func (_Token *TokenFilterer) FilterUpdatedTokenInformation(opts *bind.FilterOpts, _newName []string, _newSymbol []string, _newOnchainID []common.Address) (*TokenUpdatedTokenInformationIterator, error) {

	var _newNameRule []interface{}
	for _, _newNameItem := range _newName {
		_newNameRule = append(_newNameRule, _newNameItem)
	}
	var _newSymbolRule []interface{}
	for _, _newSymbolItem := range _newSymbol {
		_newSymbolRule = append(_newSymbolRule, _newSymbolItem)
	}

	var _newOnchainIDRule []interface{}
	for _, _newOnchainIDItem := range _newOnchainID {
		_newOnchainIDRule = append(_newOnchainIDRule, _newOnchainIDItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "UpdatedTokenInformation", _newNameRule, _newSymbolRule, _newOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TokenUpdatedTokenInformationIterator{contract: _Token.contract, event: "UpdatedTokenInformation", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenInformation is a free log subscription operation binding the contract event 0x6a1105ac8148a3c319adbc369f9072573e8a11d3a3d195e067e7c40767ec54d1.
//
// Solidity: event UpdatedTokenInformation(string indexed _newName, string indexed _newSymbol, uint8 _newDecimals, string _newVersion, address indexed _newOnchainID)
func (_Token *TokenFilterer) WatchUpdatedTokenInformation(opts *bind.WatchOpts, sink chan<- *TokenUpdatedTokenInformation, _newName []string, _newSymbol []string, _newOnchainID []common.Address) (event.Subscription, error) {

	var _newNameRule []interface{}
	for _, _newNameItem := range _newName {
		_newNameRule = append(_newNameRule, _newNameItem)
	}
	var _newSymbolRule []interface{}
	for _, _newSymbolItem := range _newSymbol {
		_newSymbolRule = append(_newSymbolRule, _newSymbolItem)
	}

	var _newOnchainIDRule []interface{}
	for _, _newOnchainIDItem := range _newOnchainID {
		_newOnchainIDRule = append(_newOnchainIDRule, _newOnchainIDItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "UpdatedTokenInformation", _newNameRule, _newSymbolRule, _newOnchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUpdatedTokenInformation)
				if err := _Token.contract.UnpackLog(event, "UpdatedTokenInformation", log); err != nil {
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

// ParseUpdatedTokenInformation is a log parse operation binding the contract event 0x6a1105ac8148a3c319adbc369f9072573e8a11d3a3d195e067e7c40767ec54d1.
//
// Solidity: event UpdatedTokenInformation(string indexed _newName, string indexed _newSymbol, uint8 _newDecimals, string _newVersion, address indexed _newOnchainID)
func (_Token *TokenFilterer) ParseUpdatedTokenInformation(log types.Log) (*TokenUpdatedTokenInformation, error) {
	event := new(TokenUpdatedTokenInformation)
	if err := _Token.contract.UnpackLog(event, "UpdatedTokenInformation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
