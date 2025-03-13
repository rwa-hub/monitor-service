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

var IdentityregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchRegisterIdentity\",\"inputs\":[{\"name\":\"_userAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_identities\",\"type\":\"address[]\",\"internalType\":\"contractIIdentity[]\"},{\"name\":\"_countries\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contains\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"identity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIdentityRegistryStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_trustedIssuersRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_claimTopicsRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identityStorage\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"investorCountry\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isVerified\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"issuersRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITrustedIssuersRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identity\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"},{\"name\":\"_country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAgent\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimTopicsRegistry\",\"inputs\":[{\"name\":\"_claimTopicsRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIdentityRegistryStorage\",\"inputs\":[{\"name\":\"_identityRegistryStorage\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrustedIssuersRegistry\",\"inputs\":[{\"name\":\"_trustedIssuersRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topicsRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIClaimTopicsRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCountry\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateIdentity\",\"inputs\":[{\"name\":\"_userAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_identity\",\"type\":\"address\",\"internalType\":\"contractIIdentity\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AgentAdded\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AgentRemoved\",\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimTopicsRegistrySet\",\"inputs\":[{\"name\":\"claimTopicsRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CountryUpdated\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"country\",\"type\":\"uint16\",\"indexed\":true,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRegistered\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityRemoved\",\"inputs\":[{\"name\":\"investorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityStorageSet\",\"inputs\":[{\"name\":\"identityStorage\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityUpdated\",\"inputs\":[{\"name\":\"oldIdentity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"},{\"name\":\"newIdentity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIIdentity\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TrustedIssuersRegistrySet\",\"inputs\":[{\"name\":\"trustedIssuersRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

var IdentityregistryABI = IdentityregistryMetaData.ABI

type Identityregistry struct {
	IdentityregistryCaller
	IdentityregistryTransactor
	IdentityregistryFilterer
}

type IdentityregistryCaller struct {
	contract *bind.BoundContract
}

type IdentityregistryTransactor struct {
	contract *bind.BoundContract
}

type IdentityregistryFilterer struct {
	contract *bind.BoundContract
}

type IdentityregistrySession struct {
	Contract     *Identityregistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IdentityregistryCallerSession struct {
	Contract *IdentityregistryCaller
	CallOpts bind.CallOpts
}

type IdentityregistryTransactorSession struct {
	Contract     *IdentityregistryTransactor
	TransactOpts bind.TransactOpts
}

type IdentityregistryRaw struct {
	Contract *Identityregistry
}

type IdentityregistryCallerRaw struct {
	Contract *IdentityregistryCaller
}

type IdentityregistryTransactorRaw struct {
	Contract *IdentityregistryTransactor
}

func NewIdentityregistry(address common.Address, backend bind.ContractBackend) (*Identityregistry, error) {
	contract, err := bindIdentityregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identityregistry{IdentityregistryCaller: IdentityregistryCaller{contract: contract}, IdentityregistryTransactor: IdentityregistryTransactor{contract: contract}, IdentityregistryFilterer: IdentityregistryFilterer{contract: contract}}, nil
}

func NewIdentityregistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityregistryCaller, error) {
	contract, err := bindIdentityregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryCaller{contract: contract}, nil
}

func NewIdentityregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityregistryTransactor, error) {
	contract, err := bindIdentityregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryTransactor{contract: contract}, nil
}

func NewIdentityregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityregistryFilterer, error) {
	contract, err := bindIdentityregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityregistryFilterer{contract: contract}, nil
}

func bindIdentityregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Identityregistry *IdentityregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identityregistry.Contract.IdentityregistryCaller.contract.Call(opts, result, method, params...)
}

func (_Identityregistry *IdentityregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityregistry.Contract.IdentityregistryTransactor.contract.Transfer(opts)
}

func (_Identityregistry *IdentityregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identityregistry.Contract.IdentityregistryTransactor.contract.Transact(opts, method, params...)
}

func (_Identityregistry *IdentityregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identityregistry.Contract.contract.Call(opts, result, method, params...)
}

func (_Identityregistry *IdentityregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityregistry.Contract.contract.Transfer(opts)
}

func (_Identityregistry *IdentityregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identityregistry.Contract.contract.Transact(opts, method, params...)
}

func (_Identityregistry *IdentityregistryCaller) Contains(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "contains", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) Contains(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.Contains(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCallerSession) Contains(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.Contains(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCaller) Identity(opts *bind.CallOpts, _userAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "identity", _userAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) Identity(_userAddress common.Address) (common.Address, error) {
	return _Identityregistry.Contract.Identity(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCallerSession) Identity(_userAddress common.Address) (common.Address, error) {
	return _Identityregistry.Contract.Identity(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCaller) IdentityStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "identityStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) IdentityStorage() (common.Address, error) {
	return _Identityregistry.Contract.IdentityStorage(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCallerSession) IdentityStorage() (common.Address, error) {
	return _Identityregistry.Contract.IdentityStorage(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCaller) InvestorCountry(opts *bind.CallOpts, _userAddress common.Address) (uint16, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "investorCountry", _userAddress)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) InvestorCountry(_userAddress common.Address) (uint16, error) {
	return _Identityregistry.Contract.InvestorCountry(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCallerSession) InvestorCountry(_userAddress common.Address) (uint16, error) {
	return _Identityregistry.Contract.InvestorCountry(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCaller) IsAgent(opts *bind.CallOpts, _agent common.Address) (bool, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "isAgent", _agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) IsAgent(_agent common.Address) (bool, error) {
	return _Identityregistry.Contract.IsAgent(&_Identityregistry.CallOpts, _agent)
}

func (_Identityregistry *IdentityregistryCallerSession) IsAgent(_agent common.Address) (bool, error) {
	return _Identityregistry.Contract.IsAgent(&_Identityregistry.CallOpts, _agent)
}

func (_Identityregistry *IdentityregistryCaller) IsVerified(opts *bind.CallOpts, _userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "isVerified", _userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) IsVerified(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.IsVerified(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCallerSession) IsVerified(_userAddress common.Address) (bool, error) {
	return _Identityregistry.Contract.IsVerified(&_Identityregistry.CallOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryCaller) IssuersRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "issuersRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) IssuersRegistry() (common.Address, error) {
	return _Identityregistry.Contract.IssuersRegistry(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCallerSession) IssuersRegistry() (common.Address, error) {
	return _Identityregistry.Contract.IssuersRegistry(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) Owner() (common.Address, error) {
	return _Identityregistry.Contract.Owner(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCallerSession) Owner() (common.Address, error) {
	return _Identityregistry.Contract.Owner(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCaller) TopicsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identityregistry.contract.Call(opts, &out, "topicsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Identityregistry *IdentityregistrySession) TopicsRegistry() (common.Address, error) {
	return _Identityregistry.Contract.TopicsRegistry(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryCallerSession) TopicsRegistry() (common.Address, error) {
	return _Identityregistry.Contract.TopicsRegistry(&_Identityregistry.CallOpts)
}

func (_Identityregistry *IdentityregistryTransactor) AddAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "addAgent", _agent)
}

func (_Identityregistry *IdentityregistrySession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.AddAgent(&_Identityregistry.TransactOpts, _agent)
}

func (_Identityregistry *IdentityregistryTransactorSession) AddAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.AddAgent(&_Identityregistry.TransactOpts, _agent)
}

func (_Identityregistry *IdentityregistryTransactor) BatchRegisterIdentity(opts *bind.TransactOpts, _userAddresses []common.Address, _identities []common.Address, _countries []uint16) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "batchRegisterIdentity", _userAddresses, _identities, _countries)
}

func (_Identityregistry *IdentityregistrySession) BatchRegisterIdentity(_userAddresses []common.Address, _identities []common.Address, _countries []uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.BatchRegisterIdentity(&_Identityregistry.TransactOpts, _userAddresses, _identities, _countries)
}

func (_Identityregistry *IdentityregistryTransactorSession) BatchRegisterIdentity(_userAddresses []common.Address, _identities []common.Address, _countries []uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.BatchRegisterIdentity(&_Identityregistry.TransactOpts, _userAddresses, _identities, _countries)
}

func (_Identityregistry *IdentityregistryTransactor) DeleteIdentity(opts *bind.TransactOpts, _userAddress common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "deleteIdentity", _userAddress)
}

func (_Identityregistry *IdentityregistrySession) DeleteIdentity(_userAddress common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.DeleteIdentity(&_Identityregistry.TransactOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryTransactorSession) DeleteIdentity(_userAddress common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.DeleteIdentity(&_Identityregistry.TransactOpts, _userAddress)
}

func (_Identityregistry *IdentityregistryTransactor) Init(opts *bind.TransactOpts, _trustedIssuersRegistry common.Address, _claimTopicsRegistry common.Address, _identityStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "init", _trustedIssuersRegistry, _claimTopicsRegistry, _identityStorage)
}

func (_Identityregistry *IdentityregistrySession) Init(_trustedIssuersRegistry common.Address, _claimTopicsRegistry common.Address, _identityStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.Init(&_Identityregistry.TransactOpts, _trustedIssuersRegistry, _claimTopicsRegistry, _identityStorage)
}

func (_Identityregistry *IdentityregistryTransactorSession) Init(_trustedIssuersRegistry common.Address, _claimTopicsRegistry common.Address, _identityStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.Init(&_Identityregistry.TransactOpts, _trustedIssuersRegistry, _claimTopicsRegistry, _identityStorage)
}

func (_Identityregistry *IdentityregistryTransactor) RegisterIdentity(opts *bind.TransactOpts, _userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "registerIdentity", _userAddress, _identity, _country)
}

func (_Identityregistry *IdentityregistrySession) RegisterIdentity(_userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.RegisterIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity, _country)
}

func (_Identityregistry *IdentityregistryTransactorSession) RegisterIdentity(_userAddress common.Address, _identity common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.RegisterIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity, _country)
}

func (_Identityregistry *IdentityregistryTransactor) RemoveAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "removeAgent", _agent)
}

func (_Identityregistry *IdentityregistrySession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.RemoveAgent(&_Identityregistry.TransactOpts, _agent)
}

func (_Identityregistry *IdentityregistryTransactorSession) RemoveAgent(_agent common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.RemoveAgent(&_Identityregistry.TransactOpts, _agent)
}

func (_Identityregistry *IdentityregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "renounceOwnership")
}

func (_Identityregistry *IdentityregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Identityregistry.Contract.RenounceOwnership(&_Identityregistry.TransactOpts)
}

func (_Identityregistry *IdentityregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Identityregistry.Contract.RenounceOwnership(&_Identityregistry.TransactOpts)
}

func (_Identityregistry *IdentityregistryTransactor) SetClaimTopicsRegistry(opts *bind.TransactOpts, _claimTopicsRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "setClaimTopicsRegistry", _claimTopicsRegistry)
}

func (_Identityregistry *IdentityregistrySession) SetClaimTopicsRegistry(_claimTopicsRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetClaimTopicsRegistry(&_Identityregistry.TransactOpts, _claimTopicsRegistry)
}

func (_Identityregistry *IdentityregistryTransactorSession) SetClaimTopicsRegistry(_claimTopicsRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetClaimTopicsRegistry(&_Identityregistry.TransactOpts, _claimTopicsRegistry)
}

func (_Identityregistry *IdentityregistryTransactor) SetIdentityRegistryStorage(opts *bind.TransactOpts, _identityRegistryStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "setIdentityRegistryStorage", _identityRegistryStorage)
}

func (_Identityregistry *IdentityregistrySession) SetIdentityRegistryStorage(_identityRegistryStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetIdentityRegistryStorage(&_Identityregistry.TransactOpts, _identityRegistryStorage)
}

func (_Identityregistry *IdentityregistryTransactorSession) SetIdentityRegistryStorage(_identityRegistryStorage common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetIdentityRegistryStorage(&_Identityregistry.TransactOpts, _identityRegistryStorage)
}

func (_Identityregistry *IdentityregistryTransactor) SetTrustedIssuersRegistry(opts *bind.TransactOpts, _trustedIssuersRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "setTrustedIssuersRegistry", _trustedIssuersRegistry)
}

func (_Identityregistry *IdentityregistrySession) SetTrustedIssuersRegistry(_trustedIssuersRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetTrustedIssuersRegistry(&_Identityregistry.TransactOpts, _trustedIssuersRegistry)
}

func (_Identityregistry *IdentityregistryTransactorSession) SetTrustedIssuersRegistry(_trustedIssuersRegistry common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.SetTrustedIssuersRegistry(&_Identityregistry.TransactOpts, _trustedIssuersRegistry)
}

func (_Identityregistry *IdentityregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_Identityregistry *IdentityregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.TransferOwnership(&_Identityregistry.TransactOpts, newOwner)
}

func (_Identityregistry *IdentityregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.TransferOwnership(&_Identityregistry.TransactOpts, newOwner)
}

func (_Identityregistry *IdentityregistryTransactor) UpdateCountry(opts *bind.TransactOpts, _userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "updateCountry", _userAddress, _country)
}

func (_Identityregistry *IdentityregistrySession) UpdateCountry(_userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateCountry(&_Identityregistry.TransactOpts, _userAddress, _country)
}

func (_Identityregistry *IdentityregistryTransactorSession) UpdateCountry(_userAddress common.Address, _country uint16) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateCountry(&_Identityregistry.TransactOpts, _userAddress, _country)
}

func (_Identityregistry *IdentityregistryTransactor) UpdateIdentity(opts *bind.TransactOpts, _userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identityregistry.contract.Transact(opts, "updateIdentity", _userAddress, _identity)
}

func (_Identityregistry *IdentityregistrySession) UpdateIdentity(_userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity)
}

func (_Identityregistry *IdentityregistryTransactorSession) UpdateIdentity(_userAddress common.Address, _identity common.Address) (*types.Transaction, error) {
	return _Identityregistry.Contract.UpdateIdentity(&_Identityregistry.TransactOpts, _userAddress, _identity)
}

type IdentityregistryAgentAddedIterator struct {
	Event *IdentityregistryAgentAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryAgentAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryAgentAddedIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryAgentAdded struct {
	Agent common.Address
	Raw   types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseAgentAdded(log types.Log) (*IdentityregistryAgentAdded, error) {
	event := new(IdentityregistryAgentAdded)
	if err := _Identityregistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryAgentRemovedIterator struct {
	Event *IdentityregistryAgentRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryAgentRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryAgentRemovedIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryAgentRemoved struct {
	Agent common.Address
	Raw   types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseAgentRemoved(log types.Log) (*IdentityregistryAgentRemoved, error) {
	event := new(IdentityregistryAgentRemoved)
	if err := _Identityregistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryClaimTopicsRegistrySetIterator struct {
	Event *IdentityregistryClaimTopicsRegistrySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryClaimTopicsRegistrySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryClaimTopicsRegistrySetIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryClaimTopicsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryClaimTopicsRegistrySet struct {
	ClaimTopicsRegistry common.Address
	Raw                 types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseClaimTopicsRegistrySet(log types.Log) (*IdentityregistryClaimTopicsRegistrySet, error) {
	event := new(IdentityregistryClaimTopicsRegistrySet)
	if err := _Identityregistry.contract.UnpackLog(event, "ClaimTopicsRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryCountryUpdatedIterator struct {
	Event *IdentityregistryCountryUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryCountryUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryCountryUpdatedIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryCountryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryCountryUpdated struct {
	InvestorAddress common.Address
	Country         uint16
	Raw             types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseCountryUpdated(log types.Log) (*IdentityregistryCountryUpdated, error) {
	event := new(IdentityregistryCountryUpdated)
	if err := _Identityregistry.contract.UnpackLog(event, "CountryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryIdentityRegisteredIterator struct {
	Event *IdentityregistryIdentityRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryIdentityRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryIdentityRegisteredIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryIdentityRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryIdentityRegistered struct {
	InvestorAddress common.Address
	Identity        common.Address
	Raw             types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseIdentityRegistered(log types.Log) (*IdentityregistryIdentityRegistered, error) {
	event := new(IdentityregistryIdentityRegistered)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryIdentityRemovedIterator struct {
	Event *IdentityregistryIdentityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryIdentityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryIdentityRemovedIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryIdentityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryIdentityRemoved struct {
	InvestorAddress common.Address
	Identity        common.Address
	Raw             types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseIdentityRemoved(log types.Log) (*IdentityregistryIdentityRemoved, error) {
	event := new(IdentityregistryIdentityRemoved)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryIdentityStorageSetIterator struct {
	Event *IdentityregistryIdentityStorageSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryIdentityStorageSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryIdentityStorageSetIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryIdentityStorageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryIdentityStorageSet struct {
	IdentityStorage common.Address
	Raw             types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseIdentityStorageSet(log types.Log) (*IdentityregistryIdentityStorageSet, error) {
	event := new(IdentityregistryIdentityStorageSet)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityStorageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryIdentityUpdatedIterator struct {
	Event *IdentityregistryIdentityUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryIdentityUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryIdentityUpdatedIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryIdentityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryIdentityUpdated struct {
	OldIdentity common.Address
	NewIdentity common.Address
	Raw         types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseIdentityUpdated(log types.Log) (*IdentityregistryIdentityUpdated, error) {
	event := new(IdentityregistryIdentityUpdated)
	if err := _Identityregistry.contract.UnpackLog(event, "IdentityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryInitializedIterator struct {
	Event *IdentityregistryInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryInitializedIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryInitialized struct {
	Version uint8
	Raw     types.Log
}

func (_Identityregistry *IdentityregistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityregistryInitializedIterator, error) {

	logs, sub, err := _Identityregistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityregistryInitializedIterator{contract: _Identityregistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseInitialized(log types.Log) (*IdentityregistryInitialized, error) {
	event := new(IdentityregistryInitialized)
	if err := _Identityregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryOwnershipTransferredIterator struct {
	Event *IdentityregistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseOwnershipTransferred(log types.Log) (*IdentityregistryOwnershipTransferred, error) {
	event := new(IdentityregistryOwnershipTransferred)
	if err := _Identityregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IdentityregistryTrustedIssuersRegistrySetIterator struct {
	Event *IdentityregistryTrustedIssuersRegistrySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IdentityregistryTrustedIssuersRegistrySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

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

func (it *IdentityregistryTrustedIssuersRegistrySetIterator) Error() error {
	return it.fail
}

func (it *IdentityregistryTrustedIssuersRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IdentityregistryTrustedIssuersRegistrySet struct {
	TrustedIssuersRegistry common.Address
	Raw                    types.Log
}

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

func (_Identityregistry *IdentityregistryFilterer) ParseTrustedIssuersRegistrySet(log types.Log) (*IdentityregistryTrustedIssuersRegistrySet, error) {
	event := new(IdentityregistryTrustedIssuersRegistrySet)
	if err := _Identityregistry.contract.UnpackLog(event, "TrustedIssuersRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
