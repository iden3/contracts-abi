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

// IdentityRegistryImplV1MetaData contains all meta data concerning the IdentityRegistryImplV1 contract.
var IdentityRegistryImplV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HUB_NOT_SET\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeafAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeafCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeafDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeafGreaterThanSnarkScalarField\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_HUB_CAN_ACCESS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"REGISTERED_COMMITMENT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongSiblingNodes\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtIndex\",\"type\":\"uint256\"}],\"name\":\"CommitmentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cscaRoot\",\"type\":\"uint256\"}],\"name\":\"CscaRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtIndex\",\"type\":\"uint256\"}],\"name\":\"DevCommitmentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DevCommitmentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLeaf\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DevCommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtIndex\",\"type\":\"uint256\"}],\"name\":\"DevDscKeyCommitmentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"}],\"name\":\"DevDscKeyCommitmentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"DevDscKeyCommitmentStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLeaf\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"}],\"name\":\"DevDscKeyCommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"DevNullifierStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"imtIndex\",\"type\":\"uint256\"}],\"name\":\"DscKeyCommitmentRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"}],\"name\":\"HubUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"}],\"name\":\"RegistryInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"checkCscaRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"checkDscKeyCommitmentMerkleRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"checkIdentityCommitmentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dscCommitment\",\"type\":\"uint256\"}],\"name\":\"devAddDscKeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"}],\"name\":\"devAddIdentityCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dscCommitment\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"devChangeDscKeyCommitmentState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"devChangeNullifierState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"siblingNodes\",\"type\":\"uint256[]\"}],\"name\":\"devRemoveCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"siblingNodes\",\"type\":\"uint256[]\"}],\"name\":\"devRemoveDscKeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLeaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"siblingNodes\",\"type\":\"uint256[]\"}],\"name\":\"devUpdateCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldLeaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLeaf\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"siblingNodes\",\"type\":\"uint256[]\"}],\"name\":\"devUpdateDscKeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCscaRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"}],\"name\":\"getDscKeyCommitmentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDscKeyCommitmentMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDscKeyCommitmentTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"}],\"name\":\"getIdentityCommitmentIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentityCommitmentMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentityCommitmentMerkleTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hub\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"}],\"name\":\"isRegisteredDscKeyCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"}],\"name\":\"nullifiers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"attestationId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"}],\"name\":\"registerCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dscCommitment\",\"type\":\"uint256\"}],\"name\":\"registerDscKeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"rootTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCscaRoot\",\"type\":\"uint256\"}],\"name\":\"updateCscaRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHubAddress\",\"type\":\"address\"}],\"name\":\"updateHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IdentityRegistryImplV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryImplV1MetaData.ABI instead.
var IdentityRegistryImplV1ABI = IdentityRegistryImplV1MetaData.ABI

// IdentityRegistryImplV1 is an auto generated Go binding around an Ethereum contract.
type IdentityRegistryImplV1 struct {
	IdentityRegistryImplV1Caller     // Read-only binding to the contract
	IdentityRegistryImplV1Transactor // Write-only binding to the contract
	IdentityRegistryImplV1Filterer   // Log filterer for contract events
}

// IdentityRegistryImplV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryImplV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryImplV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryImplV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryImplV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryImplV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryImplV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistryImplV1Session struct {
	Contract     *IdentityRegistryImplV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IdentityRegistryImplV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryImplV1CallerSession struct {
	Contract *IdentityRegistryImplV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IdentityRegistryImplV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryImplV1TransactorSession struct {
	Contract     *IdentityRegistryImplV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IdentityRegistryImplV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryImplV1Raw struct {
	Contract *IdentityRegistryImplV1 // Generic contract binding to access the raw methods on
}

// IdentityRegistryImplV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryImplV1CallerRaw struct {
	Contract *IdentityRegistryImplV1Caller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryImplV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryImplV1TransactorRaw struct {
	Contract *IdentityRegistryImplV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistryImplV1 creates a new instance of IdentityRegistryImplV1, bound to a specific deployed contract.
func NewIdentityRegistryImplV1(address common.Address, backend bind.ContractBackend) (*IdentityRegistryImplV1, error) {
	contract, err := bindIdentityRegistryImplV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1{IdentityRegistryImplV1Caller: IdentityRegistryImplV1Caller{contract: contract}, IdentityRegistryImplV1Transactor: IdentityRegistryImplV1Transactor{contract: contract}, IdentityRegistryImplV1Filterer: IdentityRegistryImplV1Filterer{contract: contract}}, nil
}

// NewIdentityRegistryImplV1Caller creates a new read-only instance of IdentityRegistryImplV1, bound to a specific deployed contract.
func NewIdentityRegistryImplV1Caller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryImplV1Caller, error) {
	contract, err := bindIdentityRegistryImplV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1Caller{contract: contract}, nil
}

// NewIdentityRegistryImplV1Transactor creates a new write-only instance of IdentityRegistryImplV1, bound to a specific deployed contract.
func NewIdentityRegistryImplV1Transactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryImplV1Transactor, error) {
	contract, err := bindIdentityRegistryImplV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1Transactor{contract: contract}, nil
}

// NewIdentityRegistryImplV1Filterer creates a new log filterer instance of IdentityRegistryImplV1, bound to a specific deployed contract.
func NewIdentityRegistryImplV1Filterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryImplV1Filterer, error) {
	contract, err := bindIdentityRegistryImplV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1Filterer{contract: contract}, nil
}

// bindIdentityRegistryImplV1 binds a generic wrapper to an already deployed contract.
func bindIdentityRegistryImplV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityRegistryImplV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistryImplV1.Contract.IdentityRegistryImplV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.IdentityRegistryImplV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.IdentityRegistryImplV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistryImplV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _IdentityRegistryImplV1.Contract.UPGRADEINTERFACEVERSION(&_IdentityRegistryImplV1.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _IdentityRegistryImplV1.Contract.UPGRADEINTERFACEVERSION(&_IdentityRegistryImplV1.CallOpts)
}

// CheckCscaRoot is a free data retrieval call binding the contract method 0x1465ff7f.
//
// Solidity: function checkCscaRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) CheckCscaRoot(opts *bind.CallOpts, root *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "checkCscaRoot", root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCscaRoot is a free data retrieval call binding the contract method 0x1465ff7f.
//
// Solidity: function checkCscaRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) CheckCscaRoot(root *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.CheckCscaRoot(&_IdentityRegistryImplV1.CallOpts, root)
}

// CheckCscaRoot is a free data retrieval call binding the contract method 0x1465ff7f.
//
// Solidity: function checkCscaRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) CheckCscaRoot(root *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.CheckCscaRoot(&_IdentityRegistryImplV1.CallOpts, root)
}

// CheckDscKeyCommitmentMerkleRoot is a free data retrieval call binding the contract method 0x9fc1e423.
//
// Solidity: function checkDscKeyCommitmentMerkleRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) CheckDscKeyCommitmentMerkleRoot(opts *bind.CallOpts, root *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "checkDscKeyCommitmentMerkleRoot", root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDscKeyCommitmentMerkleRoot is a free data retrieval call binding the contract method 0x9fc1e423.
//
// Solidity: function checkDscKeyCommitmentMerkleRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) CheckDscKeyCommitmentMerkleRoot(root *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.CheckDscKeyCommitmentMerkleRoot(&_IdentityRegistryImplV1.CallOpts, root)
}

// CheckDscKeyCommitmentMerkleRoot is a free data retrieval call binding the contract method 0x9fc1e423.
//
// Solidity: function checkDscKeyCommitmentMerkleRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) CheckDscKeyCommitmentMerkleRoot(root *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.CheckDscKeyCommitmentMerkleRoot(&_IdentityRegistryImplV1.CallOpts, root)
}

// CheckIdentityCommitmentRoot is a free data retrieval call binding the contract method 0xfd05ecd8.
//
// Solidity: function checkIdentityCommitmentRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) CheckIdentityCommitmentRoot(opts *bind.CallOpts, root *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "checkIdentityCommitmentRoot", root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIdentityCommitmentRoot is a free data retrieval call binding the contract method 0xfd05ecd8.
//
// Solidity: function checkIdentityCommitmentRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) CheckIdentityCommitmentRoot(root *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.CheckIdentityCommitmentRoot(&_IdentityRegistryImplV1.CallOpts, root)
}

// CheckIdentityCommitmentRoot is a free data retrieval call binding the contract method 0xfd05ecd8.
//
// Solidity: function checkIdentityCommitmentRoot(uint256 root) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) CheckIdentityCommitmentRoot(root *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.CheckIdentityCommitmentRoot(&_IdentityRegistryImplV1.CallOpts, root)
}

// GetCscaRoot is a free data retrieval call binding the contract method 0xaedba93c.
//
// Solidity: function getCscaRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetCscaRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getCscaRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCscaRoot is a free data retrieval call binding the contract method 0xaedba93c.
//
// Solidity: function getCscaRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetCscaRoot() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetCscaRoot(&_IdentityRegistryImplV1.CallOpts)
}

// GetCscaRoot is a free data retrieval call binding the contract method 0xaedba93c.
//
// Solidity: function getCscaRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetCscaRoot() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetCscaRoot(&_IdentityRegistryImplV1.CallOpts)
}

// GetDscKeyCommitmentIndex is a free data retrieval call binding the contract method 0x1bcc023d.
//
// Solidity: function getDscKeyCommitmentIndex(uint256 commitment) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetDscKeyCommitmentIndex(opts *bind.CallOpts, commitment *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getDscKeyCommitmentIndex", commitment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDscKeyCommitmentIndex is a free data retrieval call binding the contract method 0x1bcc023d.
//
// Solidity: function getDscKeyCommitmentIndex(uint256 commitment) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetDscKeyCommitmentIndex(commitment *big.Int) (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetDscKeyCommitmentIndex(&_IdentityRegistryImplV1.CallOpts, commitment)
}

// GetDscKeyCommitmentIndex is a free data retrieval call binding the contract method 0x1bcc023d.
//
// Solidity: function getDscKeyCommitmentIndex(uint256 commitment) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetDscKeyCommitmentIndex(commitment *big.Int) (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetDscKeyCommitmentIndex(&_IdentityRegistryImplV1.CallOpts, commitment)
}

// GetDscKeyCommitmentMerkleRoot is a free data retrieval call binding the contract method 0xd78ff6f5.
//
// Solidity: function getDscKeyCommitmentMerkleRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetDscKeyCommitmentMerkleRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getDscKeyCommitmentMerkleRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDscKeyCommitmentMerkleRoot is a free data retrieval call binding the contract method 0xd78ff6f5.
//
// Solidity: function getDscKeyCommitmentMerkleRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetDscKeyCommitmentMerkleRoot() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetDscKeyCommitmentMerkleRoot(&_IdentityRegistryImplV1.CallOpts)
}

// GetDscKeyCommitmentMerkleRoot is a free data retrieval call binding the contract method 0xd78ff6f5.
//
// Solidity: function getDscKeyCommitmentMerkleRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetDscKeyCommitmentMerkleRoot() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetDscKeyCommitmentMerkleRoot(&_IdentityRegistryImplV1.CallOpts)
}

// GetDscKeyCommitmentTreeSize is a free data retrieval call binding the contract method 0xfcceca2e.
//
// Solidity: function getDscKeyCommitmentTreeSize() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetDscKeyCommitmentTreeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getDscKeyCommitmentTreeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDscKeyCommitmentTreeSize is a free data retrieval call binding the contract method 0xfcceca2e.
//
// Solidity: function getDscKeyCommitmentTreeSize() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetDscKeyCommitmentTreeSize() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetDscKeyCommitmentTreeSize(&_IdentityRegistryImplV1.CallOpts)
}

// GetDscKeyCommitmentTreeSize is a free data retrieval call binding the contract method 0xfcceca2e.
//
// Solidity: function getDscKeyCommitmentTreeSize() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetDscKeyCommitmentTreeSize() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetDscKeyCommitmentTreeSize(&_IdentityRegistryImplV1.CallOpts)
}

// GetIdentityCommitmentIndex is a free data retrieval call binding the contract method 0x98966e3c.
//
// Solidity: function getIdentityCommitmentIndex(uint256 commitment) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetIdentityCommitmentIndex(opts *bind.CallOpts, commitment *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getIdentityCommitmentIndex", commitment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdentityCommitmentIndex is a free data retrieval call binding the contract method 0x98966e3c.
//
// Solidity: function getIdentityCommitmentIndex(uint256 commitment) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetIdentityCommitmentIndex(commitment *big.Int) (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetIdentityCommitmentIndex(&_IdentityRegistryImplV1.CallOpts, commitment)
}

// GetIdentityCommitmentIndex is a free data retrieval call binding the contract method 0x98966e3c.
//
// Solidity: function getIdentityCommitmentIndex(uint256 commitment) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetIdentityCommitmentIndex(commitment *big.Int) (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetIdentityCommitmentIndex(&_IdentityRegistryImplV1.CallOpts, commitment)
}

// GetIdentityCommitmentMerkleRoot is a free data retrieval call binding the contract method 0x4bf4dc7a.
//
// Solidity: function getIdentityCommitmentMerkleRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetIdentityCommitmentMerkleRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getIdentityCommitmentMerkleRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdentityCommitmentMerkleRoot is a free data retrieval call binding the contract method 0x4bf4dc7a.
//
// Solidity: function getIdentityCommitmentMerkleRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetIdentityCommitmentMerkleRoot() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetIdentityCommitmentMerkleRoot(&_IdentityRegistryImplV1.CallOpts)
}

// GetIdentityCommitmentMerkleRoot is a free data retrieval call binding the contract method 0x4bf4dc7a.
//
// Solidity: function getIdentityCommitmentMerkleRoot() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetIdentityCommitmentMerkleRoot() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetIdentityCommitmentMerkleRoot(&_IdentityRegistryImplV1.CallOpts)
}

// GetIdentityCommitmentMerkleTreeSize is a free data retrieval call binding the contract method 0x64567e8a.
//
// Solidity: function getIdentityCommitmentMerkleTreeSize() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) GetIdentityCommitmentMerkleTreeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "getIdentityCommitmentMerkleTreeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdentityCommitmentMerkleTreeSize is a free data retrieval call binding the contract method 0x64567e8a.
//
// Solidity: function getIdentityCommitmentMerkleTreeSize() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) GetIdentityCommitmentMerkleTreeSize() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetIdentityCommitmentMerkleTreeSize(&_IdentityRegistryImplV1.CallOpts)
}

// GetIdentityCommitmentMerkleTreeSize is a free data retrieval call binding the contract method 0x64567e8a.
//
// Solidity: function getIdentityCommitmentMerkleTreeSize() view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) GetIdentityCommitmentMerkleTreeSize() (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.GetIdentityCommitmentMerkleTreeSize(&_IdentityRegistryImplV1.CallOpts)
}

// Hub is a free data retrieval call binding the contract method 0x365a86fc.
//
// Solidity: function hub() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) Hub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "hub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hub is a free data retrieval call binding the contract method 0x365a86fc.
//
// Solidity: function hub() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) Hub() (common.Address, error) {
	return _IdentityRegistryImplV1.Contract.Hub(&_IdentityRegistryImplV1.CallOpts)
}

// Hub is a free data retrieval call binding the contract method 0x365a86fc.
//
// Solidity: function hub() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) Hub() (common.Address, error) {
	return _IdentityRegistryImplV1.Contract.Hub(&_IdentityRegistryImplV1.CallOpts)
}

// IsRegisteredDscKeyCommitment is a free data retrieval call binding the contract method 0xfc20a024.
//
// Solidity: function isRegisteredDscKeyCommitment(uint256 commitment) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) IsRegisteredDscKeyCommitment(opts *bind.CallOpts, commitment *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "isRegisteredDscKeyCommitment", commitment)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredDscKeyCommitment is a free data retrieval call binding the contract method 0xfc20a024.
//
// Solidity: function isRegisteredDscKeyCommitment(uint256 commitment) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) IsRegisteredDscKeyCommitment(commitment *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.IsRegisteredDscKeyCommitment(&_IdentityRegistryImplV1.CallOpts, commitment)
}

// IsRegisteredDscKeyCommitment is a free data retrieval call binding the contract method 0xfc20a024.
//
// Solidity: function isRegisteredDscKeyCommitment(uint256 commitment) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) IsRegisteredDscKeyCommitment(commitment *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.IsRegisteredDscKeyCommitment(&_IdentityRegistryImplV1.CallOpts, commitment)
}

// Nullifiers is a free data retrieval call binding the contract method 0xaddee776.
//
// Solidity: function nullifiers(bytes32 attestationId, uint256 nullifier) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) Nullifiers(opts *bind.CallOpts, attestationId [32]byte, nullifier *big.Int) (bool, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "nullifiers", attestationId, nullifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Nullifiers is a free data retrieval call binding the contract method 0xaddee776.
//
// Solidity: function nullifiers(bytes32 attestationId, uint256 nullifier) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) Nullifiers(attestationId [32]byte, nullifier *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.Nullifiers(&_IdentityRegistryImplV1.CallOpts, attestationId, nullifier)
}

// Nullifiers is a free data retrieval call binding the contract method 0xaddee776.
//
// Solidity: function nullifiers(bytes32 attestationId, uint256 nullifier) view returns(bool)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) Nullifiers(attestationId [32]byte, nullifier *big.Int) (bool, error) {
	return _IdentityRegistryImplV1.Contract.Nullifiers(&_IdentityRegistryImplV1.CallOpts, attestationId, nullifier)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) Owner() (common.Address, error) {
	return _IdentityRegistryImplV1.Contract.Owner(&_IdentityRegistryImplV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) Owner() (common.Address, error) {
	return _IdentityRegistryImplV1.Contract.Owner(&_IdentityRegistryImplV1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) PendingOwner() (common.Address, error) {
	return _IdentityRegistryImplV1.Contract.PendingOwner(&_IdentityRegistryImplV1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) PendingOwner() (common.Address, error) {
	return _IdentityRegistryImplV1.Contract.PendingOwner(&_IdentityRegistryImplV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) ProxiableUUID() ([32]byte, error) {
	return _IdentityRegistryImplV1.Contract.ProxiableUUID(&_IdentityRegistryImplV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _IdentityRegistryImplV1.Contract.ProxiableUUID(&_IdentityRegistryImplV1.CallOpts)
}

// RootTimestamps is a free data retrieval call binding the contract method 0xede12d9f.
//
// Solidity: function rootTimestamps(uint256 root) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Caller) RootTimestamps(opts *bind.CallOpts, root *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IdentityRegistryImplV1.contract.Call(opts, &out, "rootTimestamps", root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RootTimestamps is a free data retrieval call binding the contract method 0xede12d9f.
//
// Solidity: function rootTimestamps(uint256 root) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) RootTimestamps(root *big.Int) (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.RootTimestamps(&_IdentityRegistryImplV1.CallOpts, root)
}

// RootTimestamps is a free data retrieval call binding the contract method 0xede12d9f.
//
// Solidity: function rootTimestamps(uint256 root) view returns(uint256)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1CallerSession) RootTimestamps(root *big.Int) (*big.Int, error) {
	return _IdentityRegistryImplV1.Contract.RootTimestamps(&_IdentityRegistryImplV1.CallOpts, root)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) AcceptOwnership() (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.AcceptOwnership(&_IdentityRegistryImplV1.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.AcceptOwnership(&_IdentityRegistryImplV1.TransactOpts)
}

// DevAddDscKeyCommitment is a paid mutator transaction binding the contract method 0xc9c17b3f.
//
// Solidity: function devAddDscKeyCommitment(uint256 dscCommitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevAddDscKeyCommitment(opts *bind.TransactOpts, dscCommitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devAddDscKeyCommitment", dscCommitment)
}

// DevAddDscKeyCommitment is a paid mutator transaction binding the contract method 0xc9c17b3f.
//
// Solidity: function devAddDscKeyCommitment(uint256 dscCommitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevAddDscKeyCommitment(dscCommitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevAddDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, dscCommitment)
}

// DevAddDscKeyCommitment is a paid mutator transaction binding the contract method 0xc9c17b3f.
//
// Solidity: function devAddDscKeyCommitment(uint256 dscCommitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevAddDscKeyCommitment(dscCommitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevAddDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, dscCommitment)
}

// DevAddIdentityCommitment is a paid mutator transaction binding the contract method 0xad26e675.
//
// Solidity: function devAddIdentityCommitment(bytes32 attestationId, uint256 nullifier, uint256 commitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevAddIdentityCommitment(opts *bind.TransactOpts, attestationId [32]byte, nullifier *big.Int, commitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devAddIdentityCommitment", attestationId, nullifier, commitment)
}

// DevAddIdentityCommitment is a paid mutator transaction binding the contract method 0xad26e675.
//
// Solidity: function devAddIdentityCommitment(bytes32 attestationId, uint256 nullifier, uint256 commitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevAddIdentityCommitment(attestationId [32]byte, nullifier *big.Int, commitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevAddIdentityCommitment(&_IdentityRegistryImplV1.TransactOpts, attestationId, nullifier, commitment)
}

// DevAddIdentityCommitment is a paid mutator transaction binding the contract method 0xad26e675.
//
// Solidity: function devAddIdentityCommitment(bytes32 attestationId, uint256 nullifier, uint256 commitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevAddIdentityCommitment(attestationId [32]byte, nullifier *big.Int, commitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevAddIdentityCommitment(&_IdentityRegistryImplV1.TransactOpts, attestationId, nullifier, commitment)
}

// DevChangeDscKeyCommitmentState is a paid mutator transaction binding the contract method 0x1f299f57.
//
// Solidity: function devChangeDscKeyCommitmentState(uint256 dscCommitment, bool state) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevChangeDscKeyCommitmentState(opts *bind.TransactOpts, dscCommitment *big.Int, state bool) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devChangeDscKeyCommitmentState", dscCommitment, state)
}

// DevChangeDscKeyCommitmentState is a paid mutator transaction binding the contract method 0x1f299f57.
//
// Solidity: function devChangeDscKeyCommitmentState(uint256 dscCommitment, bool state) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevChangeDscKeyCommitmentState(dscCommitment *big.Int, state bool) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevChangeDscKeyCommitmentState(&_IdentityRegistryImplV1.TransactOpts, dscCommitment, state)
}

// DevChangeDscKeyCommitmentState is a paid mutator transaction binding the contract method 0x1f299f57.
//
// Solidity: function devChangeDscKeyCommitmentState(uint256 dscCommitment, bool state) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevChangeDscKeyCommitmentState(dscCommitment *big.Int, state bool) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevChangeDscKeyCommitmentState(&_IdentityRegistryImplV1.TransactOpts, dscCommitment, state)
}

// DevChangeNullifierState is a paid mutator transaction binding the contract method 0x88581195.
//
// Solidity: function devChangeNullifierState(bytes32 attestationId, uint256 nullifier, bool state) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevChangeNullifierState(opts *bind.TransactOpts, attestationId [32]byte, nullifier *big.Int, state bool) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devChangeNullifierState", attestationId, nullifier, state)
}

// DevChangeNullifierState is a paid mutator transaction binding the contract method 0x88581195.
//
// Solidity: function devChangeNullifierState(bytes32 attestationId, uint256 nullifier, bool state) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevChangeNullifierState(attestationId [32]byte, nullifier *big.Int, state bool) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevChangeNullifierState(&_IdentityRegistryImplV1.TransactOpts, attestationId, nullifier, state)
}

// DevChangeNullifierState is a paid mutator transaction binding the contract method 0x88581195.
//
// Solidity: function devChangeNullifierState(bytes32 attestationId, uint256 nullifier, bool state) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevChangeNullifierState(attestationId [32]byte, nullifier *big.Int, state bool) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevChangeNullifierState(&_IdentityRegistryImplV1.TransactOpts, attestationId, nullifier, state)
}

// DevRemoveCommitment is a paid mutator transaction binding the contract method 0x0bc19639.
//
// Solidity: function devRemoveCommitment(uint256 oldLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevRemoveCommitment(opts *bind.TransactOpts, oldLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devRemoveCommitment", oldLeaf, siblingNodes)
}

// DevRemoveCommitment is a paid mutator transaction binding the contract method 0x0bc19639.
//
// Solidity: function devRemoveCommitment(uint256 oldLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevRemoveCommitment(oldLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevRemoveCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, siblingNodes)
}

// DevRemoveCommitment is a paid mutator transaction binding the contract method 0x0bc19639.
//
// Solidity: function devRemoveCommitment(uint256 oldLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevRemoveCommitment(oldLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevRemoveCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, siblingNodes)
}

// DevRemoveDscKeyCommitment is a paid mutator transaction binding the contract method 0xb3ccecb8.
//
// Solidity: function devRemoveDscKeyCommitment(uint256 oldLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevRemoveDscKeyCommitment(opts *bind.TransactOpts, oldLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devRemoveDscKeyCommitment", oldLeaf, siblingNodes)
}

// DevRemoveDscKeyCommitment is a paid mutator transaction binding the contract method 0xb3ccecb8.
//
// Solidity: function devRemoveDscKeyCommitment(uint256 oldLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevRemoveDscKeyCommitment(oldLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevRemoveDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, siblingNodes)
}

// DevRemoveDscKeyCommitment is a paid mutator transaction binding the contract method 0xb3ccecb8.
//
// Solidity: function devRemoveDscKeyCommitment(uint256 oldLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevRemoveDscKeyCommitment(oldLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevRemoveDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, siblingNodes)
}

// DevUpdateCommitment is a paid mutator transaction binding the contract method 0x4378cb53.
//
// Solidity: function devUpdateCommitment(uint256 oldLeaf, uint256 newLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevUpdateCommitment(opts *bind.TransactOpts, oldLeaf *big.Int, newLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devUpdateCommitment", oldLeaf, newLeaf, siblingNodes)
}

// DevUpdateCommitment is a paid mutator transaction binding the contract method 0x4378cb53.
//
// Solidity: function devUpdateCommitment(uint256 oldLeaf, uint256 newLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevUpdateCommitment(oldLeaf *big.Int, newLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevUpdateCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, newLeaf, siblingNodes)
}

// DevUpdateCommitment is a paid mutator transaction binding the contract method 0x4378cb53.
//
// Solidity: function devUpdateCommitment(uint256 oldLeaf, uint256 newLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevUpdateCommitment(oldLeaf *big.Int, newLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevUpdateCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, newLeaf, siblingNodes)
}

// DevUpdateDscKeyCommitment is a paid mutator transaction binding the contract method 0x0f48a159.
//
// Solidity: function devUpdateDscKeyCommitment(uint256 oldLeaf, uint256 newLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) DevUpdateDscKeyCommitment(opts *bind.TransactOpts, oldLeaf *big.Int, newLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "devUpdateDscKeyCommitment", oldLeaf, newLeaf, siblingNodes)
}

// DevUpdateDscKeyCommitment is a paid mutator transaction binding the contract method 0x0f48a159.
//
// Solidity: function devUpdateDscKeyCommitment(uint256 oldLeaf, uint256 newLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) DevUpdateDscKeyCommitment(oldLeaf *big.Int, newLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevUpdateDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, newLeaf, siblingNodes)
}

// DevUpdateDscKeyCommitment is a paid mutator transaction binding the contract method 0x0f48a159.
//
// Solidity: function devUpdateDscKeyCommitment(uint256 oldLeaf, uint256 newLeaf, uint256[] siblingNodes) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) DevUpdateDscKeyCommitment(oldLeaf *big.Int, newLeaf *big.Int, siblingNodes []*big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.DevUpdateDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, oldLeaf, newLeaf, siblingNodes)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _hub) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) Initialize(opts *bind.TransactOpts, _hub common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "initialize", _hub)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _hub) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) Initialize(_hub common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.Initialize(&_IdentityRegistryImplV1.TransactOpts, _hub)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _hub) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) Initialize(_hub common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.Initialize(&_IdentityRegistryImplV1.TransactOpts, _hub)
}

// RegisterCommitment is a paid mutator transaction binding the contract method 0x9a85d40a.
//
// Solidity: function registerCommitment(bytes32 attestationId, uint256 nullifier, uint256 commitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) RegisterCommitment(opts *bind.TransactOpts, attestationId [32]byte, nullifier *big.Int, commitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "registerCommitment", attestationId, nullifier, commitment)
}

// RegisterCommitment is a paid mutator transaction binding the contract method 0x9a85d40a.
//
// Solidity: function registerCommitment(bytes32 attestationId, uint256 nullifier, uint256 commitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) RegisterCommitment(attestationId [32]byte, nullifier *big.Int, commitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.RegisterCommitment(&_IdentityRegistryImplV1.TransactOpts, attestationId, nullifier, commitment)
}

// RegisterCommitment is a paid mutator transaction binding the contract method 0x9a85d40a.
//
// Solidity: function registerCommitment(bytes32 attestationId, uint256 nullifier, uint256 commitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) RegisterCommitment(attestationId [32]byte, nullifier *big.Int, commitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.RegisterCommitment(&_IdentityRegistryImplV1.TransactOpts, attestationId, nullifier, commitment)
}

// RegisterDscKeyCommitment is a paid mutator transaction binding the contract method 0x445257da.
//
// Solidity: function registerDscKeyCommitment(uint256 dscCommitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) RegisterDscKeyCommitment(opts *bind.TransactOpts, dscCommitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "registerDscKeyCommitment", dscCommitment)
}

// RegisterDscKeyCommitment is a paid mutator transaction binding the contract method 0x445257da.
//
// Solidity: function registerDscKeyCommitment(uint256 dscCommitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) RegisterDscKeyCommitment(dscCommitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.RegisterDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, dscCommitment)
}

// RegisterDscKeyCommitment is a paid mutator transaction binding the contract method 0x445257da.
//
// Solidity: function registerDscKeyCommitment(uint256 dscCommitment) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) RegisterDscKeyCommitment(dscCommitment *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.RegisterDscKeyCommitment(&_IdentityRegistryImplV1.TransactOpts, dscCommitment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.RenounceOwnership(&_IdentityRegistryImplV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.RenounceOwnership(&_IdentityRegistryImplV1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.TransferOwnership(&_IdentityRegistryImplV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.TransferOwnership(&_IdentityRegistryImplV1.TransactOpts, newOwner)
}

// UpdateCscaRoot is a paid mutator transaction binding the contract method 0x88b7b189.
//
// Solidity: function updateCscaRoot(uint256 newCscaRoot) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) UpdateCscaRoot(opts *bind.TransactOpts, newCscaRoot *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "updateCscaRoot", newCscaRoot)
}

// UpdateCscaRoot is a paid mutator transaction binding the contract method 0x88b7b189.
//
// Solidity: function updateCscaRoot(uint256 newCscaRoot) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) UpdateCscaRoot(newCscaRoot *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.UpdateCscaRoot(&_IdentityRegistryImplV1.TransactOpts, newCscaRoot)
}

// UpdateCscaRoot is a paid mutator transaction binding the contract method 0x88b7b189.
//
// Solidity: function updateCscaRoot(uint256 newCscaRoot) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) UpdateCscaRoot(newCscaRoot *big.Int) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.UpdateCscaRoot(&_IdentityRegistryImplV1.TransactOpts, newCscaRoot)
}

// UpdateHub is a paid mutator transaction binding the contract method 0xfbf08da9.
//
// Solidity: function updateHub(address newHubAddress) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) UpdateHub(opts *bind.TransactOpts, newHubAddress common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "updateHub", newHubAddress)
}

// UpdateHub is a paid mutator transaction binding the contract method 0xfbf08da9.
//
// Solidity: function updateHub(address newHubAddress) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) UpdateHub(newHubAddress common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.UpdateHub(&_IdentityRegistryImplV1.TransactOpts, newHubAddress)
}

// UpdateHub is a paid mutator transaction binding the contract method 0xfbf08da9.
//
// Solidity: function updateHub(address newHubAddress) returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) UpdateHub(newHubAddress common.Address) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.UpdateHub(&_IdentityRegistryImplV1.TransactOpts, newHubAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.UpgradeToAndCall(&_IdentityRegistryImplV1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityRegistryImplV1.Contract.UpgradeToAndCall(&_IdentityRegistryImplV1.TransactOpts, newImplementation, data)
}

// IdentityRegistryImplV1CommitmentRegisteredIterator is returned from FilterCommitmentRegistered and is used to iterate over the raw logs and unpacked data for CommitmentRegistered events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1CommitmentRegisteredIterator struct {
	Event *IdentityRegistryImplV1CommitmentRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1CommitmentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1CommitmentRegistered)
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
		it.Event = new(IdentityRegistryImplV1CommitmentRegistered)
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
func (it *IdentityRegistryImplV1CommitmentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1CommitmentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1CommitmentRegistered represents a CommitmentRegistered event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1CommitmentRegistered struct {
	AttestationId [32]byte
	Nullifier     *big.Int
	Commitment    *big.Int
	Timestamp     *big.Int
	ImtRoot       *big.Int
	ImtIndex      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitmentRegistered is a free log retrieval operation binding the contract event 0xeb119b7270209cdf8b25ab82dcae1e1f269d3244ed481853debf439b49415dff.
//
// Solidity: event CommitmentRegistered(bytes32 indexed attestationId, uint256 indexed nullifier, uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterCommitmentRegistered(opts *bind.FilterOpts, attestationId [][32]byte, nullifier []*big.Int, commitment []*big.Int) (*IdentityRegistryImplV1CommitmentRegisteredIterator, error) {

	var attestationIdRule []interface{}
	for _, attestationIdItem := range attestationId {
		attestationIdRule = append(attestationIdRule, attestationIdItem)
	}
	var nullifierRule []interface{}
	for _, nullifierItem := range nullifier {
		nullifierRule = append(nullifierRule, nullifierItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "CommitmentRegistered", attestationIdRule, nullifierRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1CommitmentRegisteredIterator{contract: _IdentityRegistryImplV1.contract, event: "CommitmentRegistered", logs: logs, sub: sub}, nil
}

// WatchCommitmentRegistered is a free log subscription operation binding the contract event 0xeb119b7270209cdf8b25ab82dcae1e1f269d3244ed481853debf439b49415dff.
//
// Solidity: event CommitmentRegistered(bytes32 indexed attestationId, uint256 indexed nullifier, uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchCommitmentRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1CommitmentRegistered, attestationId [][32]byte, nullifier []*big.Int, commitment []*big.Int) (event.Subscription, error) {

	var attestationIdRule []interface{}
	for _, attestationIdItem := range attestationId {
		attestationIdRule = append(attestationIdRule, attestationIdItem)
	}
	var nullifierRule []interface{}
	for _, nullifierItem := range nullifier {
		nullifierRule = append(nullifierRule, nullifierItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "CommitmentRegistered", attestationIdRule, nullifierRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1CommitmentRegistered)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "CommitmentRegistered", log); err != nil {
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

// ParseCommitmentRegistered is a log parse operation binding the contract event 0xeb119b7270209cdf8b25ab82dcae1e1f269d3244ed481853debf439b49415dff.
//
// Solidity: event CommitmentRegistered(bytes32 indexed attestationId, uint256 indexed nullifier, uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseCommitmentRegistered(log types.Log) (*IdentityRegistryImplV1CommitmentRegistered, error) {
	event := new(IdentityRegistryImplV1CommitmentRegistered)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "CommitmentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1CscaRootUpdatedIterator is returned from FilterCscaRootUpdated and is used to iterate over the raw logs and unpacked data for CscaRootUpdated events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1CscaRootUpdatedIterator struct {
	Event *IdentityRegistryImplV1CscaRootUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1CscaRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1CscaRootUpdated)
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
		it.Event = new(IdentityRegistryImplV1CscaRootUpdated)
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
func (it *IdentityRegistryImplV1CscaRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1CscaRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1CscaRootUpdated represents a CscaRootUpdated event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1CscaRootUpdated struct {
	CscaRoot *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCscaRootUpdated is a free log retrieval operation binding the contract event 0x02dd6b96a8b84ff5245724a410a499208af7b8fb2150fa85f297ae005f48ecc7.
//
// Solidity: event CscaRootUpdated(uint256 cscaRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterCscaRootUpdated(opts *bind.FilterOpts) (*IdentityRegistryImplV1CscaRootUpdatedIterator, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "CscaRootUpdated")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1CscaRootUpdatedIterator{contract: _IdentityRegistryImplV1.contract, event: "CscaRootUpdated", logs: logs, sub: sub}, nil
}

// WatchCscaRootUpdated is a free log subscription operation binding the contract event 0x02dd6b96a8b84ff5245724a410a499208af7b8fb2150fa85f297ae005f48ecc7.
//
// Solidity: event CscaRootUpdated(uint256 cscaRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchCscaRootUpdated(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1CscaRootUpdated) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "CscaRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1CscaRootUpdated)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "CscaRootUpdated", log); err != nil {
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

// ParseCscaRootUpdated is a log parse operation binding the contract event 0x02dd6b96a8b84ff5245724a410a499208af7b8fb2150fa85f297ae005f48ecc7.
//
// Solidity: event CscaRootUpdated(uint256 cscaRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseCscaRootUpdated(log types.Log) (*IdentityRegistryImplV1CscaRootUpdated, error) {
	event := new(IdentityRegistryImplV1CscaRootUpdated)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "CscaRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevCommitmentRegisteredIterator is returned from FilterDevCommitmentRegistered and is used to iterate over the raw logs and unpacked data for DevCommitmentRegistered events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevCommitmentRegisteredIterator struct {
	Event *IdentityRegistryImplV1DevCommitmentRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevCommitmentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevCommitmentRegistered)
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
		it.Event = new(IdentityRegistryImplV1DevCommitmentRegistered)
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
func (it *IdentityRegistryImplV1DevCommitmentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevCommitmentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevCommitmentRegistered represents a DevCommitmentRegistered event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevCommitmentRegistered struct {
	AttestationId [32]byte
	Nullifier     *big.Int
	Commitment    *big.Int
	Timestamp     *big.Int
	ImtRoot       *big.Int
	ImtIndex      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDevCommitmentRegistered is a free log retrieval operation binding the contract event 0xfc5b28a97d17b16ee3dd749eecd0a602288e98dc57b251e77a8095147f8366b5.
//
// Solidity: event DevCommitmentRegistered(bytes32 indexed attestationId, uint256 indexed nullifier, uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevCommitmentRegistered(opts *bind.FilterOpts, attestationId [][32]byte, nullifier []*big.Int, commitment []*big.Int) (*IdentityRegistryImplV1DevCommitmentRegisteredIterator, error) {

	var attestationIdRule []interface{}
	for _, attestationIdItem := range attestationId {
		attestationIdRule = append(attestationIdRule, attestationIdItem)
	}
	var nullifierRule []interface{}
	for _, nullifierItem := range nullifier {
		nullifierRule = append(nullifierRule, nullifierItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevCommitmentRegistered", attestationIdRule, nullifierRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevCommitmentRegisteredIterator{contract: _IdentityRegistryImplV1.contract, event: "DevCommitmentRegistered", logs: logs, sub: sub}, nil
}

// WatchDevCommitmentRegistered is a free log subscription operation binding the contract event 0xfc5b28a97d17b16ee3dd749eecd0a602288e98dc57b251e77a8095147f8366b5.
//
// Solidity: event DevCommitmentRegistered(bytes32 indexed attestationId, uint256 indexed nullifier, uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevCommitmentRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevCommitmentRegistered, attestationId [][32]byte, nullifier []*big.Int, commitment []*big.Int) (event.Subscription, error) {

	var attestationIdRule []interface{}
	for _, attestationIdItem := range attestationId {
		attestationIdRule = append(attestationIdRule, attestationIdItem)
	}
	var nullifierRule []interface{}
	for _, nullifierItem := range nullifier {
		nullifierRule = append(nullifierRule, nullifierItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevCommitmentRegistered", attestationIdRule, nullifierRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevCommitmentRegistered)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevCommitmentRegistered", log); err != nil {
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

// ParseDevCommitmentRegistered is a log parse operation binding the contract event 0xfc5b28a97d17b16ee3dd749eecd0a602288e98dc57b251e77a8095147f8366b5.
//
// Solidity: event DevCommitmentRegistered(bytes32 indexed attestationId, uint256 indexed nullifier, uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevCommitmentRegistered(log types.Log) (*IdentityRegistryImplV1DevCommitmentRegistered, error) {
	event := new(IdentityRegistryImplV1DevCommitmentRegistered)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevCommitmentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevCommitmentRemovedIterator is returned from FilterDevCommitmentRemoved and is used to iterate over the raw logs and unpacked data for DevCommitmentRemoved events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevCommitmentRemovedIterator struct {
	Event *IdentityRegistryImplV1DevCommitmentRemoved // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevCommitmentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevCommitmentRemoved)
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
		it.Event = new(IdentityRegistryImplV1DevCommitmentRemoved)
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
func (it *IdentityRegistryImplV1DevCommitmentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevCommitmentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevCommitmentRemoved represents a DevCommitmentRemoved event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevCommitmentRemoved struct {
	OldLeaf   *big.Int
	ImtRoot   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDevCommitmentRemoved is a free log retrieval operation binding the contract event 0x8814493b3599188c61d7bf20585a0355af60987c540db9988392ae830a54c64a.
//
// Solidity: event DevCommitmentRemoved(uint256 indexed oldLeaf, uint256 imtRoot, uint256 timestamp)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevCommitmentRemoved(opts *bind.FilterOpts, oldLeaf []*big.Int) (*IdentityRegistryImplV1DevCommitmentRemovedIterator, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevCommitmentRemoved", oldLeafRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevCommitmentRemovedIterator{contract: _IdentityRegistryImplV1.contract, event: "DevCommitmentRemoved", logs: logs, sub: sub}, nil
}

// WatchDevCommitmentRemoved is a free log subscription operation binding the contract event 0x8814493b3599188c61d7bf20585a0355af60987c540db9988392ae830a54c64a.
//
// Solidity: event DevCommitmentRemoved(uint256 indexed oldLeaf, uint256 imtRoot, uint256 timestamp)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevCommitmentRemoved(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevCommitmentRemoved, oldLeaf []*big.Int) (event.Subscription, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevCommitmentRemoved", oldLeafRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevCommitmentRemoved)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevCommitmentRemoved", log); err != nil {
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

// ParseDevCommitmentRemoved is a log parse operation binding the contract event 0x8814493b3599188c61d7bf20585a0355af60987c540db9988392ae830a54c64a.
//
// Solidity: event DevCommitmentRemoved(uint256 indexed oldLeaf, uint256 imtRoot, uint256 timestamp)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevCommitmentRemoved(log types.Log) (*IdentityRegistryImplV1DevCommitmentRemoved, error) {
	event := new(IdentityRegistryImplV1DevCommitmentRemoved)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevCommitmentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevCommitmentUpdatedIterator is returned from FilterDevCommitmentUpdated and is used to iterate over the raw logs and unpacked data for DevCommitmentUpdated events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevCommitmentUpdatedIterator struct {
	Event *IdentityRegistryImplV1DevCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevCommitmentUpdated)
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
		it.Event = new(IdentityRegistryImplV1DevCommitmentUpdated)
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
func (it *IdentityRegistryImplV1DevCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevCommitmentUpdated represents a DevCommitmentUpdated event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevCommitmentUpdated struct {
	OldLeaf   *big.Int
	NewLeaf   *big.Int
	ImtRoot   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDevCommitmentUpdated is a free log retrieval operation binding the contract event 0x9d5851adacc1956f388355223351dad2704088fae99b36711607250dd9afe3fc.
//
// Solidity: event DevCommitmentUpdated(uint256 indexed oldLeaf, uint256 indexed newLeaf, uint256 imtRoot, uint256 timestamp)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevCommitmentUpdated(opts *bind.FilterOpts, oldLeaf []*big.Int, newLeaf []*big.Int) (*IdentityRegistryImplV1DevCommitmentUpdatedIterator, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}
	var newLeafRule []interface{}
	for _, newLeafItem := range newLeaf {
		newLeafRule = append(newLeafRule, newLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevCommitmentUpdated", oldLeafRule, newLeafRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevCommitmentUpdatedIterator{contract: _IdentityRegistryImplV1.contract, event: "DevCommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchDevCommitmentUpdated is a free log subscription operation binding the contract event 0x9d5851adacc1956f388355223351dad2704088fae99b36711607250dd9afe3fc.
//
// Solidity: event DevCommitmentUpdated(uint256 indexed oldLeaf, uint256 indexed newLeaf, uint256 imtRoot, uint256 timestamp)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevCommitmentUpdated, oldLeaf []*big.Int, newLeaf []*big.Int) (event.Subscription, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}
	var newLeafRule []interface{}
	for _, newLeafItem := range newLeaf {
		newLeafRule = append(newLeafRule, newLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevCommitmentUpdated", oldLeafRule, newLeafRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevCommitmentUpdated)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevCommitmentUpdated", log); err != nil {
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

// ParseDevCommitmentUpdated is a log parse operation binding the contract event 0x9d5851adacc1956f388355223351dad2704088fae99b36711607250dd9afe3fc.
//
// Solidity: event DevCommitmentUpdated(uint256 indexed oldLeaf, uint256 indexed newLeaf, uint256 imtRoot, uint256 timestamp)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevCommitmentUpdated(log types.Log) (*IdentityRegistryImplV1DevCommitmentUpdated, error) {
	event := new(IdentityRegistryImplV1DevCommitmentUpdated)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevCommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator is returned from FilterDevDscKeyCommitmentRegistered and is used to iterate over the raw logs and unpacked data for DevDscKeyCommitmentRegistered events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator struct {
	Event *IdentityRegistryImplV1DevDscKeyCommitmentRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentRegistered)
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
		it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentRegistered)
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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentRegistered represents a DevDscKeyCommitmentRegistered event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentRegistered struct {
	Commitment *big.Int
	ImtRoot    *big.Int
	ImtIndex   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDevDscKeyCommitmentRegistered is a free log retrieval operation binding the contract event 0x324818d8d34a2d439a57172d0cccad7d8476d81b3cb18a43969b503ebe973da8.
//
// Solidity: event DevDscKeyCommitmentRegistered(uint256 indexed commitment, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevDscKeyCommitmentRegistered(opts *bind.FilterOpts, commitment []*big.Int) (*IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevDscKeyCommitmentRegistered", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevDscKeyCommitmentRegisteredIterator{contract: _IdentityRegistryImplV1.contract, event: "DevDscKeyCommitmentRegistered", logs: logs, sub: sub}, nil
}

// WatchDevDscKeyCommitmentRegistered is a free log subscription operation binding the contract event 0x324818d8d34a2d439a57172d0cccad7d8476d81b3cb18a43969b503ebe973da8.
//
// Solidity: event DevDscKeyCommitmentRegistered(uint256 indexed commitment, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevDscKeyCommitmentRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevDscKeyCommitmentRegistered, commitment []*big.Int) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevDscKeyCommitmentRegistered", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevDscKeyCommitmentRegistered)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentRegistered", log); err != nil {
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

// ParseDevDscKeyCommitmentRegistered is a log parse operation binding the contract event 0x324818d8d34a2d439a57172d0cccad7d8476d81b3cb18a43969b503ebe973da8.
//
// Solidity: event DevDscKeyCommitmentRegistered(uint256 indexed commitment, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevDscKeyCommitmentRegistered(log types.Log) (*IdentityRegistryImplV1DevDscKeyCommitmentRegistered, error) {
	event := new(IdentityRegistryImplV1DevDscKeyCommitmentRegistered)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator is returned from FilterDevDscKeyCommitmentRemoved and is used to iterate over the raw logs and unpacked data for DevDscKeyCommitmentRemoved events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator struct {
	Event *IdentityRegistryImplV1DevDscKeyCommitmentRemoved // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentRemoved)
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
		it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentRemoved)
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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentRemoved represents a DevDscKeyCommitmentRemoved event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentRemoved struct {
	OldLeaf *big.Int
	ImtRoot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDevDscKeyCommitmentRemoved is a free log retrieval operation binding the contract event 0xdb059cfc380ead4900d342dc84f5f44bfa83e73439c4740ba3a13b1b12e0ba26.
//
// Solidity: event DevDscKeyCommitmentRemoved(uint256 indexed oldLeaf, uint256 imtRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevDscKeyCommitmentRemoved(opts *bind.FilterOpts, oldLeaf []*big.Int) (*IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevDscKeyCommitmentRemoved", oldLeafRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevDscKeyCommitmentRemovedIterator{contract: _IdentityRegistryImplV1.contract, event: "DevDscKeyCommitmentRemoved", logs: logs, sub: sub}, nil
}

// WatchDevDscKeyCommitmentRemoved is a free log subscription operation binding the contract event 0xdb059cfc380ead4900d342dc84f5f44bfa83e73439c4740ba3a13b1b12e0ba26.
//
// Solidity: event DevDscKeyCommitmentRemoved(uint256 indexed oldLeaf, uint256 imtRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevDscKeyCommitmentRemoved(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevDscKeyCommitmentRemoved, oldLeaf []*big.Int) (event.Subscription, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevDscKeyCommitmentRemoved", oldLeafRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevDscKeyCommitmentRemoved)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentRemoved", log); err != nil {
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

// ParseDevDscKeyCommitmentRemoved is a log parse operation binding the contract event 0xdb059cfc380ead4900d342dc84f5f44bfa83e73439c4740ba3a13b1b12e0ba26.
//
// Solidity: event DevDscKeyCommitmentRemoved(uint256 indexed oldLeaf, uint256 imtRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevDscKeyCommitmentRemoved(log types.Log) (*IdentityRegistryImplV1DevDscKeyCommitmentRemoved, error) {
	event := new(IdentityRegistryImplV1DevDscKeyCommitmentRemoved)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator is returned from FilterDevDscKeyCommitmentStateChanged and is used to iterate over the raw logs and unpacked data for DevDscKeyCommitmentStateChanged events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator struct {
	Event *IdentityRegistryImplV1DevDscKeyCommitmentStateChanged // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentStateChanged)
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
		it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentStateChanged)
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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentStateChanged represents a DevDscKeyCommitmentStateChanged event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentStateChanged struct {
	Commitment *big.Int
	State      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDevDscKeyCommitmentStateChanged is a free log retrieval operation binding the contract event 0x11e1b30cd0b6eab7020e0fbd635ef39b211bda7696dfd324caaef4f1ae9cd315.
//
// Solidity: event DevDscKeyCommitmentStateChanged(uint256 indexed commitment, bool state)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevDscKeyCommitmentStateChanged(opts *bind.FilterOpts, commitment []*big.Int) (*IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevDscKeyCommitmentStateChanged", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevDscKeyCommitmentStateChangedIterator{contract: _IdentityRegistryImplV1.contract, event: "DevDscKeyCommitmentStateChanged", logs: logs, sub: sub}, nil
}

// WatchDevDscKeyCommitmentStateChanged is a free log subscription operation binding the contract event 0x11e1b30cd0b6eab7020e0fbd635ef39b211bda7696dfd324caaef4f1ae9cd315.
//
// Solidity: event DevDscKeyCommitmentStateChanged(uint256 indexed commitment, bool state)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevDscKeyCommitmentStateChanged(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevDscKeyCommitmentStateChanged, commitment []*big.Int) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevDscKeyCommitmentStateChanged", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevDscKeyCommitmentStateChanged)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentStateChanged", log); err != nil {
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

// ParseDevDscKeyCommitmentStateChanged is a log parse operation binding the contract event 0x11e1b30cd0b6eab7020e0fbd635ef39b211bda7696dfd324caaef4f1ae9cd315.
//
// Solidity: event DevDscKeyCommitmentStateChanged(uint256 indexed commitment, bool state)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevDscKeyCommitmentStateChanged(log types.Log) (*IdentityRegistryImplV1DevDscKeyCommitmentStateChanged, error) {
	event := new(IdentityRegistryImplV1DevDscKeyCommitmentStateChanged)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator is returned from FilterDevDscKeyCommitmentUpdated and is used to iterate over the raw logs and unpacked data for DevDscKeyCommitmentUpdated events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator struct {
	Event *IdentityRegistryImplV1DevDscKeyCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentUpdated)
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
		it.Event = new(IdentityRegistryImplV1DevDscKeyCommitmentUpdated)
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
func (it *IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevDscKeyCommitmentUpdated represents a DevDscKeyCommitmentUpdated event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevDscKeyCommitmentUpdated struct {
	OldLeaf *big.Int
	NewLeaf *big.Int
	ImtRoot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDevDscKeyCommitmentUpdated is a free log retrieval operation binding the contract event 0x0bb5433e87cfc429682aadc2e217fe3b6915922eb25540401aaa64d0b8eaa1c0.
//
// Solidity: event DevDscKeyCommitmentUpdated(uint256 indexed oldLeaf, uint256 indexed newLeaf, uint256 imtRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevDscKeyCommitmentUpdated(opts *bind.FilterOpts, oldLeaf []*big.Int, newLeaf []*big.Int) (*IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}
	var newLeafRule []interface{}
	for _, newLeafItem := range newLeaf {
		newLeafRule = append(newLeafRule, newLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevDscKeyCommitmentUpdated", oldLeafRule, newLeafRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevDscKeyCommitmentUpdatedIterator{contract: _IdentityRegistryImplV1.contract, event: "DevDscKeyCommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchDevDscKeyCommitmentUpdated is a free log subscription operation binding the contract event 0x0bb5433e87cfc429682aadc2e217fe3b6915922eb25540401aaa64d0b8eaa1c0.
//
// Solidity: event DevDscKeyCommitmentUpdated(uint256 indexed oldLeaf, uint256 indexed newLeaf, uint256 imtRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevDscKeyCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevDscKeyCommitmentUpdated, oldLeaf []*big.Int, newLeaf []*big.Int) (event.Subscription, error) {

	var oldLeafRule []interface{}
	for _, oldLeafItem := range oldLeaf {
		oldLeafRule = append(oldLeafRule, oldLeafItem)
	}
	var newLeafRule []interface{}
	for _, newLeafItem := range newLeaf {
		newLeafRule = append(newLeafRule, newLeafItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevDscKeyCommitmentUpdated", oldLeafRule, newLeafRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevDscKeyCommitmentUpdated)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentUpdated", log); err != nil {
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

// ParseDevDscKeyCommitmentUpdated is a log parse operation binding the contract event 0x0bb5433e87cfc429682aadc2e217fe3b6915922eb25540401aaa64d0b8eaa1c0.
//
// Solidity: event DevDscKeyCommitmentUpdated(uint256 indexed oldLeaf, uint256 indexed newLeaf, uint256 imtRoot)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevDscKeyCommitmentUpdated(log types.Log) (*IdentityRegistryImplV1DevDscKeyCommitmentUpdated, error) {
	event := new(IdentityRegistryImplV1DevDscKeyCommitmentUpdated)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevDscKeyCommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DevNullifierStateChangedIterator is returned from FilterDevNullifierStateChanged and is used to iterate over the raw logs and unpacked data for DevNullifierStateChanged events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevNullifierStateChangedIterator struct {
	Event *IdentityRegistryImplV1DevNullifierStateChanged // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DevNullifierStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DevNullifierStateChanged)
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
		it.Event = new(IdentityRegistryImplV1DevNullifierStateChanged)
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
func (it *IdentityRegistryImplV1DevNullifierStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DevNullifierStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DevNullifierStateChanged represents a DevNullifierStateChanged event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DevNullifierStateChanged struct {
	AttestationId [32]byte
	Nullifier     *big.Int
	State         bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDevNullifierStateChanged is a free log retrieval operation binding the contract event 0xfdea0e3abaa4a1f72ff03f918d65910a9e2ce74814cfc21561b188db582bbe15.
//
// Solidity: event DevNullifierStateChanged(bytes32 indexed attestationId, uint256 indexed nullifier, bool state)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDevNullifierStateChanged(opts *bind.FilterOpts, attestationId [][32]byte, nullifier []*big.Int) (*IdentityRegistryImplV1DevNullifierStateChangedIterator, error) {

	var attestationIdRule []interface{}
	for _, attestationIdItem := range attestationId {
		attestationIdRule = append(attestationIdRule, attestationIdItem)
	}
	var nullifierRule []interface{}
	for _, nullifierItem := range nullifier {
		nullifierRule = append(nullifierRule, nullifierItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DevNullifierStateChanged", attestationIdRule, nullifierRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DevNullifierStateChangedIterator{contract: _IdentityRegistryImplV1.contract, event: "DevNullifierStateChanged", logs: logs, sub: sub}, nil
}

// WatchDevNullifierStateChanged is a free log subscription operation binding the contract event 0xfdea0e3abaa4a1f72ff03f918d65910a9e2ce74814cfc21561b188db582bbe15.
//
// Solidity: event DevNullifierStateChanged(bytes32 indexed attestationId, uint256 indexed nullifier, bool state)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDevNullifierStateChanged(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DevNullifierStateChanged, attestationId [][32]byte, nullifier []*big.Int) (event.Subscription, error) {

	var attestationIdRule []interface{}
	for _, attestationIdItem := range attestationId {
		attestationIdRule = append(attestationIdRule, attestationIdItem)
	}
	var nullifierRule []interface{}
	for _, nullifierItem := range nullifier {
		nullifierRule = append(nullifierRule, nullifierItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DevNullifierStateChanged", attestationIdRule, nullifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DevNullifierStateChanged)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevNullifierStateChanged", log); err != nil {
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

// ParseDevNullifierStateChanged is a log parse operation binding the contract event 0xfdea0e3abaa4a1f72ff03f918d65910a9e2ce74814cfc21561b188db582bbe15.
//
// Solidity: event DevNullifierStateChanged(bytes32 indexed attestationId, uint256 indexed nullifier, bool state)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDevNullifierStateChanged(log types.Log) (*IdentityRegistryImplV1DevNullifierStateChanged, error) {
	event := new(IdentityRegistryImplV1DevNullifierStateChanged)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DevNullifierStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator is returned from FilterDscKeyCommitmentRegistered and is used to iterate over the raw logs and unpacked data for DscKeyCommitmentRegistered events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator struct {
	Event *IdentityRegistryImplV1DscKeyCommitmentRegistered // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1DscKeyCommitmentRegistered)
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
		it.Event = new(IdentityRegistryImplV1DscKeyCommitmentRegistered)
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
func (it *IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1DscKeyCommitmentRegistered represents a DscKeyCommitmentRegistered event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1DscKeyCommitmentRegistered struct {
	Commitment *big.Int
	Timestamp  *big.Int
	ImtRoot    *big.Int
	ImtIndex   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDscKeyCommitmentRegistered is a free log retrieval operation binding the contract event 0x8fd1e12daf7bbc111604e9312366b73b92a65058cc13768ee02edf176ab7ba8e.
//
// Solidity: event DscKeyCommitmentRegistered(uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterDscKeyCommitmentRegistered(opts *bind.FilterOpts, commitment []*big.Int) (*IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "DscKeyCommitmentRegistered", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1DscKeyCommitmentRegisteredIterator{contract: _IdentityRegistryImplV1.contract, event: "DscKeyCommitmentRegistered", logs: logs, sub: sub}, nil
}

// WatchDscKeyCommitmentRegistered is a free log subscription operation binding the contract event 0x8fd1e12daf7bbc111604e9312366b73b92a65058cc13768ee02edf176ab7ba8e.
//
// Solidity: event DscKeyCommitmentRegistered(uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchDscKeyCommitmentRegistered(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1DscKeyCommitmentRegistered, commitment []*big.Int) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "DscKeyCommitmentRegistered", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1DscKeyCommitmentRegistered)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DscKeyCommitmentRegistered", log); err != nil {
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

// ParseDscKeyCommitmentRegistered is a log parse operation binding the contract event 0x8fd1e12daf7bbc111604e9312366b73b92a65058cc13768ee02edf176ab7ba8e.
//
// Solidity: event DscKeyCommitmentRegistered(uint256 indexed commitment, uint256 timestamp, uint256 imtRoot, uint256 imtIndex)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseDscKeyCommitmentRegistered(log types.Log) (*IdentityRegistryImplV1DscKeyCommitmentRegistered, error) {
	event := new(IdentityRegistryImplV1DscKeyCommitmentRegistered)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "DscKeyCommitmentRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1HubUpdatedIterator is returned from FilterHubUpdated and is used to iterate over the raw logs and unpacked data for HubUpdated events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1HubUpdatedIterator struct {
	Event *IdentityRegistryImplV1HubUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1HubUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1HubUpdated)
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
		it.Event = new(IdentityRegistryImplV1HubUpdated)
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
func (it *IdentityRegistryImplV1HubUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1HubUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1HubUpdated represents a HubUpdated event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1HubUpdated struct {
	Hub common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHubUpdated is a free log retrieval operation binding the contract event 0x5b4b08c1a8c9cb40687a2bb0a22c161920757446a2a35c3c2856ec693946d4f0.
//
// Solidity: event HubUpdated(address hub)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterHubUpdated(opts *bind.FilterOpts) (*IdentityRegistryImplV1HubUpdatedIterator, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "HubUpdated")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1HubUpdatedIterator{contract: _IdentityRegistryImplV1.contract, event: "HubUpdated", logs: logs, sub: sub}, nil
}

// WatchHubUpdated is a free log subscription operation binding the contract event 0x5b4b08c1a8c9cb40687a2bb0a22c161920757446a2a35c3c2856ec693946d4f0.
//
// Solidity: event HubUpdated(address hub)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchHubUpdated(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1HubUpdated) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "HubUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1HubUpdated)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "HubUpdated", log); err != nil {
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

// ParseHubUpdated is a log parse operation binding the contract event 0x5b4b08c1a8c9cb40687a2bb0a22c161920757446a2a35c3c2856ec693946d4f0.
//
// Solidity: event HubUpdated(address hub)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseHubUpdated(log types.Log) (*IdentityRegistryImplV1HubUpdated, error) {
	event := new(IdentityRegistryImplV1HubUpdated)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "HubUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1InitializedIterator struct {
	Event *IdentityRegistryImplV1Initialized // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1Initialized)
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
		it.Event = new(IdentityRegistryImplV1Initialized)
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
func (it *IdentityRegistryImplV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1Initialized represents a Initialized event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityRegistryImplV1InitializedIterator, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1InitializedIterator{contract: _IdentityRegistryImplV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1Initialized) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1Initialized)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseInitialized(log types.Log) (*IdentityRegistryImplV1Initialized, error) {
	event := new(IdentityRegistryImplV1Initialized)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1OwnershipTransferStartedIterator struct {
	Event *IdentityRegistryImplV1OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1OwnershipTransferStarted)
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
		it.Event = new(IdentityRegistryImplV1OwnershipTransferStarted)
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
func (it *IdentityRegistryImplV1OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityRegistryImplV1OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1OwnershipTransferStartedIterator{contract: _IdentityRegistryImplV1.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1OwnershipTransferStarted)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseOwnershipTransferStarted(log types.Log) (*IdentityRegistryImplV1OwnershipTransferStarted, error) {
	event := new(IdentityRegistryImplV1OwnershipTransferStarted)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1OwnershipTransferredIterator struct {
	Event *IdentityRegistryImplV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1OwnershipTransferred)
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
		it.Event = new(IdentityRegistryImplV1OwnershipTransferred)
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
func (it *IdentityRegistryImplV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1OwnershipTransferred represents a OwnershipTransferred event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityRegistryImplV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1OwnershipTransferredIterator{contract: _IdentityRegistryImplV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1OwnershipTransferred)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseOwnershipTransferred(log types.Log) (*IdentityRegistryImplV1OwnershipTransferred, error) {
	event := new(IdentityRegistryImplV1OwnershipTransferred)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1RegistryInitializedIterator is returned from FilterRegistryInitialized and is used to iterate over the raw logs and unpacked data for RegistryInitialized events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1RegistryInitializedIterator struct {
	Event *IdentityRegistryImplV1RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1RegistryInitialized)
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
		it.Event = new(IdentityRegistryImplV1RegistryInitialized)
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
func (it *IdentityRegistryImplV1RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1RegistryInitialized represents a RegistryInitialized event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1RegistryInitialized struct {
	Hub common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRegistryInitialized is a free log retrieval operation binding the contract event 0x43dcf2166be267d8f6dc31580b9b1a605fddb5259a48dddd50b0348ac38fb24e.
//
// Solidity: event RegistryInitialized(address hub)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterRegistryInitialized(opts *bind.FilterOpts) (*IdentityRegistryImplV1RegistryInitializedIterator, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "RegistryInitialized")
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1RegistryInitializedIterator{contract: _IdentityRegistryImplV1.contract, event: "RegistryInitialized", logs: logs, sub: sub}, nil
}

// WatchRegistryInitialized is a free log subscription operation binding the contract event 0x43dcf2166be267d8f6dc31580b9b1a605fddb5259a48dddd50b0348ac38fb24e.
//
// Solidity: event RegistryInitialized(address hub)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchRegistryInitialized(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "RegistryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1RegistryInitialized)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "RegistryInitialized", log); err != nil {
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

// ParseRegistryInitialized is a log parse operation binding the contract event 0x43dcf2166be267d8f6dc31580b9b1a605fddb5259a48dddd50b0348ac38fb24e.
//
// Solidity: event RegistryInitialized(address hub)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseRegistryInitialized(log types.Log) (*IdentityRegistryImplV1RegistryInitialized, error) {
	event := new(IdentityRegistryImplV1RegistryInitialized)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "RegistryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryImplV1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1UpgradedIterator struct {
	Event *IdentityRegistryImplV1Upgraded // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryImplV1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryImplV1Upgraded)
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
		it.Event = new(IdentityRegistryImplV1Upgraded)
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
func (it *IdentityRegistryImplV1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryImplV1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryImplV1Upgraded represents a Upgraded event raised by the IdentityRegistryImplV1 contract.
type IdentityRegistryImplV1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IdentityRegistryImplV1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryImplV1UpgradedIterator{contract: _IdentityRegistryImplV1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IdentityRegistryImplV1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IdentityRegistryImplV1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryImplV1Upgraded)
				if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_IdentityRegistryImplV1 *IdentityRegistryImplV1Filterer) ParseUpgraded(log types.Log) (*IdentityRegistryImplV1Upgraded, error) {
	event := new(IdentityRegistryImplV1Upgraded)
	if err := _IdentityRegistryImplV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
