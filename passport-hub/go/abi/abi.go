// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// IDscCircuitVerifierDscCircuitProof is an auto generated low-level Go binding around an user-defined struct.
type IDscCircuitVerifierDscCircuitProof struct {
	A          [2]*big.Int
	B          [2][2]*big.Int
	C          [2]*big.Int
	PubSignals [2]*big.Int
}

// IdentityVerificationHubImplV1MetaData contains all meta data concerning the IdentityVerificationHubImplV1 contract.
var IdentityVerificationHubImplV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CURRENT_DATE_NOT_IN_VALID_RANGE\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_COMMITMENT_ROOT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_CSCA_ROOT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_DSC_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_OFAC_ROOT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_OLDER_THAN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_VERIFIER_SET\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"DscCircuitVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"dscCircuitVerifierIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"dscCircuitVerifiers\",\"type\":\"address[]\"}],\"name\":\"HubInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"typeIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"verifierAddresses\",\"type\":\"address[]\"}],\"name\":\"batchUpdateDscCircuitVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"dscCircuitVerifierIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"dscCircuitVerifierAddresses\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dscCircuitVerifierId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubSignals\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIDscCircuitVerifier.DscCircuitProof\",\"name\":\"dscCircuitProof\",\"type\":\"tuple\"}],\"name\":\"registerDscKeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"}],\"name\":\"sigTypeToDscCircuitVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"name\":\"updateDscVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"updateRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IdentityVerificationHubImplV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityVerificationHubImplV1MetaData.ABI instead.
var IdentityVerificationHubImplV1ABI = IdentityVerificationHubImplV1MetaData.ABI

// IdentityVerificationHubImplV1 is an auto generated Go binding around an Ethereum contract.
type IdentityVerificationHubImplV1 struct {
	IdentityVerificationHubImplV1Caller     // Read-only binding to the contract
	IdentityVerificationHubImplV1Transactor // Write-only binding to the contract
	IdentityVerificationHubImplV1Filterer   // Log filterer for contract events
}

// IdentityVerificationHubImplV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityVerificationHubImplV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityVerificationHubImplV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityVerificationHubImplV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityVerificationHubImplV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityVerificationHubImplV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityVerificationHubImplV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityVerificationHubImplV1Session struct {
	Contract     *IdentityVerificationHubImplV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IdentityVerificationHubImplV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityVerificationHubImplV1CallerSession struct {
	Contract *IdentityVerificationHubImplV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// IdentityVerificationHubImplV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityVerificationHubImplV1TransactorSession struct {
	Contract     *IdentityVerificationHubImplV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// IdentityVerificationHubImplV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityVerificationHubImplV1Raw struct {
	Contract *IdentityVerificationHubImplV1 // Generic contract binding to access the raw methods on
}

// IdentityVerificationHubImplV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityVerificationHubImplV1CallerRaw struct {
	Contract *IdentityVerificationHubImplV1Caller // Generic read-only contract binding to access the raw methods on
}

// IdentityVerificationHubImplV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityVerificationHubImplV1TransactorRaw struct {
	Contract *IdentityVerificationHubImplV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityVerificationHubImplV1 creates a new instance of IdentityVerificationHubImplV1, bound to a specific deployed contract.
func NewIdentityVerificationHubImplV1(address common.Address, backend bind.ContractBackend) (*IdentityVerificationHubImplV1, error) {
	contract, err := bindIdentityVerificationHubImplV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1{IdentityVerificationHubImplV1Caller: IdentityVerificationHubImplV1Caller{contract: contract}, IdentityVerificationHubImplV1Transactor: IdentityVerificationHubImplV1Transactor{contract: contract}, IdentityVerificationHubImplV1Filterer: IdentityVerificationHubImplV1Filterer{contract: contract}}, nil
}

// NewIdentityVerificationHubImplV1Caller creates a new read-only instance of IdentityVerificationHubImplV1, bound to a specific deployed contract.
func NewIdentityVerificationHubImplV1Caller(address common.Address, caller bind.ContractCaller) (*IdentityVerificationHubImplV1Caller, error) {
	contract, err := bindIdentityVerificationHubImplV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1Caller{contract: contract}, nil
}

// NewIdentityVerificationHubImplV1Transactor creates a new write-only instance of IdentityVerificationHubImplV1, bound to a specific deployed contract.
func NewIdentityVerificationHubImplV1Transactor(address common.Address, transactor bind.ContractTransactor) (*IdentityVerificationHubImplV1Transactor, error) {
	contract, err := bindIdentityVerificationHubImplV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1Transactor{contract: contract}, nil
}

// NewIdentityVerificationHubImplV1Filterer creates a new log filterer instance of IdentityVerificationHubImplV1, bound to a specific deployed contract.
func NewIdentityVerificationHubImplV1Filterer(address common.Address, filterer bind.ContractFilterer) (*IdentityVerificationHubImplV1Filterer, error) {
	contract, err := bindIdentityVerificationHubImplV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1Filterer{contract: contract}, nil
}

// bindIdentityVerificationHubImplV1 binds a generic wrapper to an already deployed contract.
func bindIdentityVerificationHubImplV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityVerificationHubImplV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityVerificationHubImplV1.Contract.IdentityVerificationHubImplV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.IdentityVerificationHubImplV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.IdentityVerificationHubImplV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityVerificationHubImplV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdentityVerificationHubImplV1.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _IdentityVerificationHubImplV1.Contract.UPGRADEINTERFACEVERSION(&_IdentityVerificationHubImplV1.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _IdentityVerificationHubImplV1.Contract.UPGRADEINTERFACEVERSION(&_IdentityVerificationHubImplV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityVerificationHubImplV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) Owner() (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.Owner(&_IdentityVerificationHubImplV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerSession) Owner() (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.Owner(&_IdentityVerificationHubImplV1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityVerificationHubImplV1.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) PendingOwner() (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.PendingOwner(&_IdentityVerificationHubImplV1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerSession) PendingOwner() (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.PendingOwner(&_IdentityVerificationHubImplV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdentityVerificationHubImplV1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) ProxiableUUID() ([32]byte, error) {
	return _IdentityVerificationHubImplV1.Contract.ProxiableUUID(&_IdentityVerificationHubImplV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _IdentityVerificationHubImplV1.Contract.ProxiableUUID(&_IdentityVerificationHubImplV1.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Caller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityVerificationHubImplV1.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) Registry() (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.Registry(&_IdentityVerificationHubImplV1.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerSession) Registry() (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.Registry(&_IdentityVerificationHubImplV1.CallOpts)
}

// SigTypeToDscCircuitVerifiers is a free data retrieval call binding the contract method 0x2ee85c3e.
//
// Solidity: function sigTypeToDscCircuitVerifiers(uint256 typeId) view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Caller) SigTypeToDscCircuitVerifiers(opts *bind.CallOpts, typeId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IdentityVerificationHubImplV1.contract.Call(opts, &out, "sigTypeToDscCircuitVerifiers", typeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigTypeToDscCircuitVerifiers is a free data retrieval call binding the contract method 0x2ee85c3e.
//
// Solidity: function sigTypeToDscCircuitVerifiers(uint256 typeId) view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) SigTypeToDscCircuitVerifiers(typeId *big.Int) (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.SigTypeToDscCircuitVerifiers(&_IdentityVerificationHubImplV1.CallOpts, typeId)
}

// SigTypeToDscCircuitVerifiers is a free data retrieval call binding the contract method 0x2ee85c3e.
//
// Solidity: function sigTypeToDscCircuitVerifiers(uint256 typeId) view returns(address)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1CallerSession) SigTypeToDscCircuitVerifiers(typeId *big.Int) (common.Address, error) {
	return _IdentityVerificationHubImplV1.Contract.SigTypeToDscCircuitVerifiers(&_IdentityVerificationHubImplV1.CallOpts, typeId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) AcceptOwnership() (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.AcceptOwnership(&_IdentityVerificationHubImplV1.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.AcceptOwnership(&_IdentityVerificationHubImplV1.TransactOpts)
}

// BatchUpdateDscCircuitVerifiers is a paid mutator transaction binding the contract method 0xf015f579.
//
// Solidity: function batchUpdateDscCircuitVerifiers(uint256[] typeIds, address[] verifierAddresses) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) BatchUpdateDscCircuitVerifiers(opts *bind.TransactOpts, typeIds []*big.Int, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "batchUpdateDscCircuitVerifiers", typeIds, verifierAddresses)
}

// BatchUpdateDscCircuitVerifiers is a paid mutator transaction binding the contract method 0xf015f579.
//
// Solidity: function batchUpdateDscCircuitVerifiers(uint256[] typeIds, address[] verifierAddresses) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) BatchUpdateDscCircuitVerifiers(typeIds []*big.Int, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.BatchUpdateDscCircuitVerifiers(&_IdentityVerificationHubImplV1.TransactOpts, typeIds, verifierAddresses)
}

// BatchUpdateDscCircuitVerifiers is a paid mutator transaction binding the contract method 0xf015f579.
//
// Solidity: function batchUpdateDscCircuitVerifiers(uint256[] typeIds, address[] verifierAddresses) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) BatchUpdateDscCircuitVerifiers(typeIds []*big.Int, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.BatchUpdateDscCircuitVerifiers(&_IdentityVerificationHubImplV1.TransactOpts, typeIds, verifierAddresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x27cff759.
//
// Solidity: function initialize(address registryAddress, uint256[] dscCircuitVerifierIds, address[] dscCircuitVerifierAddresses) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, dscCircuitVerifierIds []*big.Int, dscCircuitVerifierAddresses []common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "initialize", registryAddress, dscCircuitVerifierIds, dscCircuitVerifierAddresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x27cff759.
//
// Solidity: function initialize(address registryAddress, uint256[] dscCircuitVerifierIds, address[] dscCircuitVerifierAddresses) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) Initialize(registryAddress common.Address, dscCircuitVerifierIds []*big.Int, dscCircuitVerifierAddresses []common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.Initialize(&_IdentityVerificationHubImplV1.TransactOpts, registryAddress, dscCircuitVerifierIds, dscCircuitVerifierAddresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x27cff759.
//
// Solidity: function initialize(address registryAddress, uint256[] dscCircuitVerifierIds, address[] dscCircuitVerifierAddresses) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) Initialize(registryAddress common.Address, dscCircuitVerifierIds []*big.Int, dscCircuitVerifierAddresses []common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.Initialize(&_IdentityVerificationHubImplV1.TransactOpts, registryAddress, dscCircuitVerifierIds, dscCircuitVerifierAddresses)
}

// RegisterDscKeyCommitment is a paid mutator transaction binding the contract method 0x3d098d8b.
//
// Solidity: function registerDscKeyCommitment(uint256 dscCircuitVerifierId, (uint256[2],uint256[2][2],uint256[2],uint256[2]) dscCircuitProof) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) RegisterDscKeyCommitment(opts *bind.TransactOpts, dscCircuitVerifierId *big.Int, dscCircuitProof IDscCircuitVerifierDscCircuitProof) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "registerDscKeyCommitment", dscCircuitVerifierId, dscCircuitProof)
}

// RegisterDscKeyCommitment is a paid mutator transaction binding the contract method 0x3d098d8b.
//
// Solidity: function registerDscKeyCommitment(uint256 dscCircuitVerifierId, (uint256[2],uint256[2][2],uint256[2],uint256[2]) dscCircuitProof) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) RegisterDscKeyCommitment(dscCircuitVerifierId *big.Int, dscCircuitProof IDscCircuitVerifierDscCircuitProof) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.RegisterDscKeyCommitment(&_IdentityVerificationHubImplV1.TransactOpts, dscCircuitVerifierId, dscCircuitProof)
}

// RegisterDscKeyCommitment is a paid mutator transaction binding the contract method 0x3d098d8b.
//
// Solidity: function registerDscKeyCommitment(uint256 dscCircuitVerifierId, (uint256[2],uint256[2][2],uint256[2],uint256[2]) dscCircuitProof) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) RegisterDscKeyCommitment(dscCircuitVerifierId *big.Int, dscCircuitProof IDscCircuitVerifierDscCircuitProof) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.RegisterDscKeyCommitment(&_IdentityVerificationHubImplV1.TransactOpts, dscCircuitVerifierId, dscCircuitProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.RenounceOwnership(&_IdentityVerificationHubImplV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.RenounceOwnership(&_IdentityVerificationHubImplV1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.TransferOwnership(&_IdentityVerificationHubImplV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.TransferOwnership(&_IdentityVerificationHubImplV1.TransactOpts, newOwner)
}

// UpdateDscVerifier is a paid mutator transaction binding the contract method 0x30aec59e.
//
// Solidity: function updateDscVerifier(uint256 typeId, address verifierAddress) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) UpdateDscVerifier(opts *bind.TransactOpts, typeId *big.Int, verifierAddress common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "updateDscVerifier", typeId, verifierAddress)
}

// UpdateDscVerifier is a paid mutator transaction binding the contract method 0x30aec59e.
//
// Solidity: function updateDscVerifier(uint256 typeId, address verifierAddress) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) UpdateDscVerifier(typeId *big.Int, verifierAddress common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.UpdateDscVerifier(&_IdentityVerificationHubImplV1.TransactOpts, typeId, verifierAddress)
}

// UpdateDscVerifier is a paid mutator transaction binding the contract method 0x30aec59e.
//
// Solidity: function updateDscVerifier(uint256 typeId, address verifierAddress) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) UpdateDscVerifier(typeId *big.Int, verifierAddress common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.UpdateDscVerifier(&_IdentityVerificationHubImplV1.TransactOpts, typeId, verifierAddress)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address registryAddress) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) UpdateRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "updateRegistry", registryAddress)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address registryAddress) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) UpdateRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.UpdateRegistry(&_IdentityVerificationHubImplV1.TransactOpts, registryAddress)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address registryAddress) returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) UpdateRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.UpdateRegistry(&_IdentityVerificationHubImplV1.TransactOpts, registryAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.UpgradeToAndCall(&_IdentityVerificationHubImplV1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityVerificationHubImplV1.Contract.UpgradeToAndCall(&_IdentityVerificationHubImplV1.TransactOpts, newImplementation, data)
}

// IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator is returned from FilterDscCircuitVerifierUpdated and is used to iterate over the raw logs and unpacked data for DscCircuitVerifierUpdated events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator struct {
	Event *IdentityVerificationHubImplV1DscCircuitVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1DscCircuitVerifierUpdated)
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
		it.Event = new(IdentityVerificationHubImplV1DscCircuitVerifierUpdated)
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
func (it *IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1DscCircuitVerifierUpdated represents a DscCircuitVerifierUpdated event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1DscCircuitVerifierUpdated struct {
	TypeId   *big.Int
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDscCircuitVerifierUpdated is a free log retrieval operation binding the contract event 0x02c7aeb52f414fb9b4dfa6100f38185b5e12b184349218c18700bc39423bedc9.
//
// Solidity: event DscCircuitVerifierUpdated(uint256 typeId, address verifier)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterDscCircuitVerifierUpdated(opts *bind.FilterOpts) (*IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "DscCircuitVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1DscCircuitVerifierUpdatedIterator{contract: _IdentityVerificationHubImplV1.contract, event: "DscCircuitVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchDscCircuitVerifierUpdated is a free log subscription operation binding the contract event 0x02c7aeb52f414fb9b4dfa6100f38185b5e12b184349218c18700bc39423bedc9.
//
// Solidity: event DscCircuitVerifierUpdated(uint256 typeId, address verifier)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchDscCircuitVerifierUpdated(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1DscCircuitVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "DscCircuitVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1DscCircuitVerifierUpdated)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "DscCircuitVerifierUpdated", log); err != nil {
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

// ParseDscCircuitVerifierUpdated is a log parse operation binding the contract event 0x02c7aeb52f414fb9b4dfa6100f38185b5e12b184349218c18700bc39423bedc9.
//
// Solidity: event DscCircuitVerifierUpdated(uint256 typeId, address verifier)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseDscCircuitVerifierUpdated(log types.Log) (*IdentityVerificationHubImplV1DscCircuitVerifierUpdated, error) {
	event := new(IdentityVerificationHubImplV1DscCircuitVerifierUpdated)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "DscCircuitVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityVerificationHubImplV1HubInitializedIterator is returned from FilterHubInitialized and is used to iterate over the raw logs and unpacked data for HubInitialized events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1HubInitializedIterator struct {
	Event *IdentityVerificationHubImplV1HubInitialized // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1HubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1HubInitialized)
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
		it.Event = new(IdentityVerificationHubImplV1HubInitialized)
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
func (it *IdentityVerificationHubImplV1HubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1HubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1HubInitialized represents a HubInitialized event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1HubInitialized struct {
	Registry              common.Address
	DscCircuitVerifierIds []*big.Int
	DscCircuitVerifiers   []common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterHubInitialized is a free log retrieval operation binding the contract event 0x520d7e312d5fd314a2b728ddd29ff0eba340e71955f127e6eb38bbdebf600dee.
//
// Solidity: event HubInitialized(address registry, uint256[] dscCircuitVerifierIds, address[] dscCircuitVerifiers)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterHubInitialized(opts *bind.FilterOpts) (*IdentityVerificationHubImplV1HubInitializedIterator, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "HubInitialized")
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1HubInitializedIterator{contract: _IdentityVerificationHubImplV1.contract, event: "HubInitialized", logs: logs, sub: sub}, nil
}

// WatchHubInitialized is a free log subscription operation binding the contract event 0x520d7e312d5fd314a2b728ddd29ff0eba340e71955f127e6eb38bbdebf600dee.
//
// Solidity: event HubInitialized(address registry, uint256[] dscCircuitVerifierIds, address[] dscCircuitVerifiers)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchHubInitialized(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1HubInitialized) (event.Subscription, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "HubInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1HubInitialized)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "HubInitialized", log); err != nil {
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

// ParseHubInitialized is a log parse operation binding the contract event 0x520d7e312d5fd314a2b728ddd29ff0eba340e71955f127e6eb38bbdebf600dee.
//
// Solidity: event HubInitialized(address registry, uint256[] dscCircuitVerifierIds, address[] dscCircuitVerifiers)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseHubInitialized(log types.Log) (*IdentityVerificationHubImplV1HubInitialized, error) {
	event := new(IdentityVerificationHubImplV1HubInitialized)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "HubInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityVerificationHubImplV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1InitializedIterator struct {
	Event *IdentityVerificationHubImplV1Initialized // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1Initialized)
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
		it.Event = new(IdentityVerificationHubImplV1Initialized)
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
func (it *IdentityVerificationHubImplV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1Initialized represents a Initialized event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityVerificationHubImplV1InitializedIterator, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1InitializedIterator{contract: _IdentityVerificationHubImplV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1Initialized) (event.Subscription, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1Initialized)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseInitialized(log types.Log) (*IdentityVerificationHubImplV1Initialized, error) {
	event := new(IdentityVerificationHubImplV1Initialized)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityVerificationHubImplV1OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1OwnershipTransferStartedIterator struct {
	Event *IdentityVerificationHubImplV1OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1OwnershipTransferStarted)
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
		it.Event = new(IdentityVerificationHubImplV1OwnershipTransferStarted)
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
func (it *IdentityVerificationHubImplV1OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityVerificationHubImplV1OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1OwnershipTransferStartedIterator{contract: _IdentityVerificationHubImplV1.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1OwnershipTransferStarted)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseOwnershipTransferStarted(log types.Log) (*IdentityVerificationHubImplV1OwnershipTransferStarted, error) {
	event := new(IdentityVerificationHubImplV1OwnershipTransferStarted)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityVerificationHubImplV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1OwnershipTransferredIterator struct {
	Event *IdentityVerificationHubImplV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1OwnershipTransferred)
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
		it.Event = new(IdentityVerificationHubImplV1OwnershipTransferred)
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
func (it *IdentityVerificationHubImplV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1OwnershipTransferred represents a OwnershipTransferred event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityVerificationHubImplV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1OwnershipTransferredIterator{contract: _IdentityVerificationHubImplV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1OwnershipTransferred)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseOwnershipTransferred(log types.Log) (*IdentityVerificationHubImplV1OwnershipTransferred, error) {
	event := new(IdentityVerificationHubImplV1OwnershipTransferred)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityVerificationHubImplV1RegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1RegistryUpdatedIterator struct {
	Event *IdentityVerificationHubImplV1RegistryUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1RegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1RegistryUpdated)
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
		it.Event = new(IdentityVerificationHubImplV1RegistryUpdated)
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
func (it *IdentityVerificationHubImplV1RegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1RegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1RegistryUpdated represents a RegistryUpdated event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1RegistryUpdated struct {
	Registry common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0xd6ceddf6d2a22f21c7c81675c518004eff43bc5c8a6fc32a0b748e69d58671cd.
//
// Solidity: event RegistryUpdated(address registry)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*IdentityVerificationHubImplV1RegistryUpdatedIterator, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1RegistryUpdatedIterator{contract: _IdentityVerificationHubImplV1.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0xd6ceddf6d2a22f21c7c81675c518004eff43bc5c8a6fc32a0b748e69d58671cd.
//
// Solidity: event RegistryUpdated(address registry)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1RegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1RegistryUpdated)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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

// ParseRegistryUpdated is a log parse operation binding the contract event 0xd6ceddf6d2a22f21c7c81675c518004eff43bc5c8a6fc32a0b748e69d58671cd.
//
// Solidity: event RegistryUpdated(address registry)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseRegistryUpdated(log types.Log) (*IdentityVerificationHubImplV1RegistryUpdated, error) {
	event := new(IdentityVerificationHubImplV1RegistryUpdated)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityVerificationHubImplV1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1UpgradedIterator struct {
	Event *IdentityVerificationHubImplV1Upgraded // Event containing the contract specifics and raw log

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
func (it *IdentityVerificationHubImplV1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityVerificationHubImplV1Upgraded)
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
		it.Event = new(IdentityVerificationHubImplV1Upgraded)
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
func (it *IdentityVerificationHubImplV1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityVerificationHubImplV1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityVerificationHubImplV1Upgraded represents a Upgraded event raised by the IdentityVerificationHubImplV1 contract.
type IdentityVerificationHubImplV1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IdentityVerificationHubImplV1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IdentityVerificationHubImplV1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IdentityVerificationHubImplV1UpgradedIterator{contract: _IdentityVerificationHubImplV1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IdentityVerificationHubImplV1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IdentityVerificationHubImplV1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityVerificationHubImplV1Upgraded)
				if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityVerificationHubImplV1 *IdentityVerificationHubImplV1Filterer) ParseUpgraded(log types.Log) (*IdentityVerificationHubImplV1Upgraded, error) {
	event := new(IdentityVerificationHubImplV1Upgraded)
	if err := _IdentityVerificationHubImplV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
