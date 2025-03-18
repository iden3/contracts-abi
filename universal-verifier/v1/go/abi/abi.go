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

// ICircuitValidatorKeyToInputIndex is an auto generated low-level Go binding around an user-defined struct.
type ICircuitValidatorKeyToInputIndex struct {
	Key        string
	InputIndex *big.Int
}

// IZKPVerifierProofStatus is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierProofStatus struct {
	IsVerified       bool
	ValidatorVersion string
	BlockNumber      *big.Int
	BlockTimestamp   *big.Int
}

// IZKPVerifierZKPRequest is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierZKPRequest struct {
	Metadata  string
	Validator common.Address
	Data      []byte
}

// IZKPVerifierZKPResponse is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierZKPResponse struct {
	RequestId uint64
	ZkProof   []byte
	Data      []byte
}

// UniversalVerifierMetaData contains all meta data concerning the UniversalVerifier contract.
var UniversalVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"requestIdToCompare\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkIdToCompare\",\"type\":\"uint256\"}],\"name\":\"LinkedProofError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ZKPRequestSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ZKPRequestUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ZKPResponseSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REQUESTS_RETURN_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidatorToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"disableZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"enableZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getProofStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"validatorVersion\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIZKPVerifier.ProofStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProofStorageField\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getRequestOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"getZKPRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"zkpRequest\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getZKPRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZKPRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIState\",\"name\":\"state\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"isProofVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isWhitelistedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"isZKPRequestEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidatorFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"}],\"name\":\"requestIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"}],\"name\":\"setRequestOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIState\",\"name\":\"state\",\"type\":\"address\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"setZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"requestIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"name\":\"setZKPRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"submitZKPResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"zkProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPResponse[]\",\"name\":\"responses\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"crossChainProof\",\"type\":\"bytes\"}],\"name\":\"submitZKPResponseV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractICircuitValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"updateZKPRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"requestIds\",\"type\":\"uint64[]\"}],\"name\":\"verifyLinkedProofs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"verifyZKPResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"inputIndex\",\"type\":\"uint256\"}],\"internalType\":\"structICircuitValidator.KeyToInputIndex[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// UniversalVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalVerifierMetaData.ABI instead.
var UniversalVerifierABI = UniversalVerifierMetaData.ABI

// UniversalVerifier is an auto generated Go binding around an Ethereum contract.
type UniversalVerifier struct {
	UniversalVerifierCaller     // Read-only binding to the contract
	UniversalVerifierTransactor // Write-only binding to the contract
	UniversalVerifierFilterer   // Log filterer for contract events
}

// UniversalVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalVerifierSession struct {
	Contract     *UniversalVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniversalVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalVerifierCallerSession struct {
	Contract *UniversalVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniversalVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalVerifierTransactorSession struct {
	Contract     *UniversalVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniversalVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalVerifierRaw struct {
	Contract *UniversalVerifier // Generic contract binding to access the raw methods on
}

// UniversalVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalVerifierCallerRaw struct {
	Contract *UniversalVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalVerifierTransactorRaw struct {
	Contract *UniversalVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalVerifier creates a new instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifier(address common.Address, backend bind.ContractBackend) (*UniversalVerifier, error) {
	contract, err := bindUniversalVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifier{UniversalVerifierCaller: UniversalVerifierCaller{contract: contract}, UniversalVerifierTransactor: UniversalVerifierTransactor{contract: contract}, UniversalVerifierFilterer: UniversalVerifierFilterer{contract: contract}}, nil
}

// NewUniversalVerifierCaller creates a new read-only instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifierCaller(address common.Address, caller bind.ContractCaller) (*UniversalVerifierCaller, error) {
	contract, err := bindUniversalVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierCaller{contract: contract}, nil
}

// NewUniversalVerifierTransactor creates a new write-only instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalVerifierTransactor, error) {
	contract, err := bindUniversalVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierTransactor{contract: contract}, nil
}

// NewUniversalVerifierFilterer creates a new log filterer instance of UniversalVerifier, bound to a specific deployed contract.
func NewUniversalVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalVerifierFilterer, error) {
	contract, err := bindUniversalVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierFilterer{contract: contract}, nil
}

// bindUniversalVerifier binds a generic wrapper to an already deployed contract.
func bindUniversalVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalVerifier *UniversalVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalVerifier.Contract.UniversalVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalVerifier *UniversalVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UniversalVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalVerifier *UniversalVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UniversalVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalVerifier *UniversalVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalVerifier *UniversalVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalVerifier *UniversalVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.contract.Transact(opts, method, params...)
}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) REQUESTSRETURNLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "REQUESTS_RETURN_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) REQUESTSRETURNLIMIT() (*big.Int, error) {
	return _UniversalVerifier.Contract.REQUESTSRETURNLIMIT(&_UniversalVerifier.CallOpts)
}

// REQUESTSRETURNLIMIT is a free data retrieval call binding the contract method 0x1905e7b1.
//
// Solidity: function REQUESTS_RETURN_LIMIT() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) REQUESTSRETURNLIMIT() (*big.Int, error) {
	return _UniversalVerifier.Contract.REQUESTSRETURNLIMIT(&_UniversalVerifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalVerifier *UniversalVerifierCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalVerifier *UniversalVerifierSession) VERSION() (string, error) {
	return _UniversalVerifier.Contract.VERSION(&_UniversalVerifier.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalVerifier *UniversalVerifierCallerSession) VERSION() (string, error) {
	return _UniversalVerifier.Contract.VERSION(&_UniversalVerifier.CallOpts)
}

// GetProofStatus is a free data retrieval call binding the contract method 0x8c1da2c9.
//
// Solidity: function getProofStatus(address sender, uint64 requestId) view returns((bool,string,uint256,uint256))
func (_UniversalVerifier *UniversalVerifierCaller) GetProofStatus(opts *bind.CallOpts, sender common.Address, requestId uint64) (IZKPVerifierProofStatus, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getProofStatus", sender, requestId)

	if err != nil {
		return *new(IZKPVerifierProofStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IZKPVerifierProofStatus)).(*IZKPVerifierProofStatus)

	return out0, err

}

// GetProofStatus is a free data retrieval call binding the contract method 0x8c1da2c9.
//
// Solidity: function getProofStatus(address sender, uint64 requestId) view returns((bool,string,uint256,uint256))
func (_UniversalVerifier *UniversalVerifierSession) GetProofStatus(sender common.Address, requestId uint64) (IZKPVerifierProofStatus, error) {
	return _UniversalVerifier.Contract.GetProofStatus(&_UniversalVerifier.CallOpts, sender, requestId)
}

// GetProofStatus is a free data retrieval call binding the contract method 0x8c1da2c9.
//
// Solidity: function getProofStatus(address sender, uint64 requestId) view returns((bool,string,uint256,uint256))
func (_UniversalVerifier *UniversalVerifierCallerSession) GetProofStatus(sender common.Address, requestId uint64) (IZKPVerifierProofStatus, error) {
	return _UniversalVerifier.Contract.GetProofStatus(&_UniversalVerifier.CallOpts, sender, requestId)
}

// GetProofStorageField is a free data retrieval call binding the contract method 0xd4858051.
//
// Solidity: function getProofStorageField(address user, uint64 requestId, string key) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetProofStorageField(opts *bind.CallOpts, user common.Address, requestId uint64, key string) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getProofStorageField", user, requestId, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProofStorageField is a free data retrieval call binding the contract method 0xd4858051.
//
// Solidity: function getProofStorageField(address user, uint64 requestId, string key) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetProofStorageField(user common.Address, requestId uint64, key string) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetProofStorageField(&_UniversalVerifier.CallOpts, user, requestId, key)
}

// GetProofStorageField is a free data retrieval call binding the contract method 0xd4858051.
//
// Solidity: function getProofStorageField(address user, uint64 requestId, string key) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetProofStorageField(user common.Address, requestId uint64, key string) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetProofStorageField(&_UniversalVerifier.CallOpts, user, requestId, key)
}

// GetRequestOwner is a free data retrieval call binding the contract method 0x6f2477ad.
//
// Solidity: function getRequestOwner(uint64 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) GetRequestOwner(opts *bind.CallOpts, requestId uint64) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRequestOwner", requestId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRequestOwner is a free data retrieval call binding the contract method 0x6f2477ad.
//
// Solidity: function getRequestOwner(uint64 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) GetRequestOwner(requestId uint64) (common.Address, error) {
	return _UniversalVerifier.Contract.GetRequestOwner(&_UniversalVerifier.CallOpts, requestId)
}

// GetRequestOwner is a free data retrieval call binding the contract method 0x6f2477ad.
//
// Solidity: function getRequestOwner(uint64 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRequestOwner(requestId uint64) (common.Address, error) {
	return _UniversalVerifier.Contract.GetRequestOwner(&_UniversalVerifier.CallOpts, requestId)
}

// GetStateAddress is a free data retrieval call binding the contract method 0x31969e57.
//
// Solidity: function getStateAddress() view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) GetStateAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getStateAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStateAddress is a free data retrieval call binding the contract method 0x31969e57.
//
// Solidity: function getStateAddress() view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) GetStateAddress() (common.Address, error) {
	return _UniversalVerifier.Contract.GetStateAddress(&_UniversalVerifier.CallOpts)
}

// GetStateAddress is a free data retrieval call binding the contract method 0x31969e57.
//
// Solidity: function getStateAddress() view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetStateAddress() (common.Address, error) {
	return _UniversalVerifier.Contract.GetStateAddress(&_UniversalVerifier.CallOpts)
}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes) zkpRequest)
func (_UniversalVerifier *UniversalVerifierCaller) GetZKPRequest(opts *bind.CallOpts, requestId uint64) (IZKPVerifierZKPRequest, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getZKPRequest", requestId)

	if err != nil {
		return *new(IZKPVerifierZKPRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IZKPVerifierZKPRequest)).(*IZKPVerifierZKPRequest)

	return out0, err

}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes) zkpRequest)
func (_UniversalVerifier *UniversalVerifierSession) GetZKPRequest(requestId uint64) (IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequest(&_UniversalVerifier.CallOpts, requestId)
}

// GetZKPRequest is a free data retrieval call binding the contract method 0xc76d0845.
//
// Solidity: function getZKPRequest(uint64 requestId) view returns((string,address,bytes) zkpRequest)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetZKPRequest(requestId uint64) (IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequest(&_UniversalVerifier.CallOpts, requestId)
}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_UniversalVerifier *UniversalVerifierCaller) GetZKPRequests(opts *bind.CallOpts, startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getZKPRequests", startIndex, length)

	if err != nil {
		return *new([]IZKPVerifierZKPRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IZKPVerifierZKPRequest)).(*[]IZKPVerifierZKPRequest)

	return out0, err

}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_UniversalVerifier *UniversalVerifierSession) GetZKPRequests(startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequests(&_UniversalVerifier.CallOpts, startIndex, length)
}

// GetZKPRequests is a free data retrieval call binding the contract method 0x5f9e60d7.
//
// Solidity: function getZKPRequests(uint256 startIndex, uint256 length) view returns((string,address,bytes)[])
func (_UniversalVerifier *UniversalVerifierCallerSession) GetZKPRequests(startIndex *big.Int, length *big.Int) ([]IZKPVerifierZKPRequest, error) {
	return _UniversalVerifier.Contract.GetZKPRequests(&_UniversalVerifier.CallOpts, startIndex, length)
}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetZKPRequestsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getZKPRequestsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetZKPRequestsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetZKPRequestsCount(&_UniversalVerifier.CallOpts)
}

// GetZKPRequestsCount is a free data retrieval call binding the contract method 0x6508e1b4.
//
// Solidity: function getZKPRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetZKPRequestsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetZKPRequestsCount(&_UniversalVerifier.CallOpts)
}

// IsProofVerified is a free data retrieval call binding the contract method 0x49555fb1.
//
// Solidity: function isProofVerified(address sender, uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsProofVerified(opts *bind.CallOpts, sender common.Address, requestId uint64) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isProofVerified", sender, requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProofVerified is a free data retrieval call binding the contract method 0x49555fb1.
//
// Solidity: function isProofVerified(address sender, uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsProofVerified(sender common.Address, requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsProofVerified(&_UniversalVerifier.CallOpts, sender, requestId)
}

// IsProofVerified is a free data retrieval call binding the contract method 0x49555fb1.
//
// Solidity: function isProofVerified(address sender, uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsProofVerified(sender common.Address, requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsProofVerified(&_UniversalVerifier.CallOpts, sender, requestId)
}

// IsWhitelistedValidator is a free data retrieval call binding the contract method 0x9b36d036.
//
// Solidity: function isWhitelistedValidator(address validator) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsWhitelistedValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isWhitelistedValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedValidator is a free data retrieval call binding the contract method 0x9b36d036.
//
// Solidity: function isWhitelistedValidator(address validator) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsWhitelistedValidator(validator common.Address) (bool, error) {
	return _UniversalVerifier.Contract.IsWhitelistedValidator(&_UniversalVerifier.CallOpts, validator)
}

// IsWhitelistedValidator is a free data retrieval call binding the contract method 0x9b36d036.
//
// Solidity: function isWhitelistedValidator(address validator) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsWhitelistedValidator(validator common.Address) (bool, error) {
	return _UniversalVerifier.Contract.IsWhitelistedValidator(&_UniversalVerifier.CallOpts, validator)
}

// IsZKPRequestEnabled is a free data retrieval call binding the contract method 0x130a73ac.
//
// Solidity: function isZKPRequestEnabled(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsZKPRequestEnabled(opts *bind.CallOpts, requestId uint64) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isZKPRequestEnabled", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZKPRequestEnabled is a free data retrieval call binding the contract method 0x130a73ac.
//
// Solidity: function isZKPRequestEnabled(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsZKPRequestEnabled(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsZKPRequestEnabled(&_UniversalVerifier.CallOpts, requestId)
}

// IsZKPRequestEnabled is a free data retrieval call binding the contract method 0x130a73ac.
//
// Solidity: function isZKPRequestEnabled(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsZKPRequestEnabled(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.IsZKPRequestEnabled(&_UniversalVerifier.CallOpts, requestId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) Owner() (common.Address, error) {
	return _UniversalVerifier.Contract.Owner(&_UniversalVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) Owner() (common.Address, error) {
	return _UniversalVerifier.Contract.Owner(&_UniversalVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) PendingOwner() (common.Address, error) {
	return _UniversalVerifier.Contract.PendingOwner(&_UniversalVerifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) PendingOwner() (common.Address, error) {
	return _UniversalVerifier.Contract.PendingOwner(&_UniversalVerifier.CallOpts)
}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) RequestIdExists(opts *bind.CallOpts, requestId uint64) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "requestIdExists", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) RequestIdExists(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.RequestIdExists(&_UniversalVerifier.CallOpts, requestId)
}

// RequestIdExists is a free data retrieval call binding the contract method 0xab7bcfb7.
//
// Solidity: function requestIdExists(uint64 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) RequestIdExists(requestId uint64) (bool, error) {
	return _UniversalVerifier.Contract.RequestIdExists(&_UniversalVerifier.CallOpts, requestId)
}

// VerifyLinkedProofs is a free data retrieval call binding the contract method 0x0106a959.
//
// Solidity: function verifyLinkedProofs(address sender, uint64[] requestIds) view returns()
func (_UniversalVerifier *UniversalVerifierCaller) VerifyLinkedProofs(opts *bind.CallOpts, sender common.Address, requestIds []uint64) error {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "verifyLinkedProofs", sender, requestIds)

	if err != nil {
		return err
	}

	return err

}

// VerifyLinkedProofs is a free data retrieval call binding the contract method 0x0106a959.
//
// Solidity: function verifyLinkedProofs(address sender, uint64[] requestIds) view returns()
func (_UniversalVerifier *UniversalVerifierSession) VerifyLinkedProofs(sender common.Address, requestIds []uint64) error {
	return _UniversalVerifier.Contract.VerifyLinkedProofs(&_UniversalVerifier.CallOpts, sender, requestIds)
}

// VerifyLinkedProofs is a free data retrieval call binding the contract method 0x0106a959.
//
// Solidity: function verifyLinkedProofs(address sender, uint64[] requestIds) view returns()
func (_UniversalVerifier *UniversalVerifierCallerSession) VerifyLinkedProofs(sender common.Address, requestIds []uint64) error {
	return _UniversalVerifier.Contract.VerifyLinkedProofs(&_UniversalVerifier.CallOpts, sender, requestIds)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UniversalVerifier *UniversalVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UniversalVerifier *UniversalVerifierSession) Version() (string, error) {
	return _UniversalVerifier.Contract.Version(&_UniversalVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UniversalVerifier *UniversalVerifierCallerSession) Version() (string, error) {
	return _UniversalVerifier.Contract.Version(&_UniversalVerifier.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalVerifier *UniversalVerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AcceptOwnership(&_UniversalVerifier.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AcceptOwnership(&_UniversalVerifier.TransactOpts)
}

// AddValidatorToWhitelist is a paid mutator transaction binding the contract method 0x59f871a1.
//
// Solidity: function addValidatorToWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) AddValidatorToWhitelist(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "addValidatorToWhitelist", validator)
}

// AddValidatorToWhitelist is a paid mutator transaction binding the contract method 0x59f871a1.
//
// Solidity: function addValidatorToWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierSession) AddValidatorToWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AddValidatorToWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// AddValidatorToWhitelist is a paid mutator transaction binding the contract method 0x59f871a1.
//
// Solidity: function addValidatorToWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) AddValidatorToWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.AddValidatorToWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// DisableZKPRequest is a paid mutator transaction binding the contract method 0x82aff29f.
//
// Solidity: function disableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) DisableZKPRequest(opts *bind.TransactOpts, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "disableZKPRequest", requestId)
}

// DisableZKPRequest is a paid mutator transaction binding the contract method 0x82aff29f.
//
// Solidity: function disableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) DisableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// DisableZKPRequest is a paid mutator transaction binding the contract method 0x82aff29f.
//
// Solidity: function disableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) DisableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// EnableZKPRequest is a paid mutator transaction binding the contract method 0x30def4ac.
//
// Solidity: function enableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) EnableZKPRequest(opts *bind.TransactOpts, requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "enableZKPRequest", requestId)
}

// EnableZKPRequest is a paid mutator transaction binding the contract method 0x30def4ac.
//
// Solidity: function enableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) EnableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// EnableZKPRequest is a paid mutator transaction binding the contract method 0x30def4ac.
//
// Solidity: function enableZKPRequest(uint64 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) EnableZKPRequest(requestId uint64) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableZKPRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address state, address owner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) Initialize(opts *bind.TransactOpts, state common.Address, owner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "initialize", state, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address state, address owner) returns()
func (_UniversalVerifier *UniversalVerifierSession) Initialize(state common.Address, owner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.Initialize(&_UniversalVerifier.TransactOpts, state, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address state, address owner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) Initialize(state common.Address, owner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.Initialize(&_UniversalVerifier.TransactOpts, state, owner)
}

// RemoveValidatorFromWhitelist is a paid mutator transaction binding the contract method 0x617879fe.
//
// Solidity: function removeValidatorFromWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) RemoveValidatorFromWhitelist(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "removeValidatorFromWhitelist", validator)
}

// RemoveValidatorFromWhitelist is a paid mutator transaction binding the contract method 0x617879fe.
//
// Solidity: function removeValidatorFromWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierSession) RemoveValidatorFromWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RemoveValidatorFromWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// RemoveValidatorFromWhitelist is a paid mutator transaction binding the contract method 0x617879fe.
//
// Solidity: function removeValidatorFromWhitelist(address validator) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) RemoveValidatorFromWhitelist(validator common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RemoveValidatorFromWhitelist(&_UniversalVerifier.TransactOpts, validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalVerifier *UniversalVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RenounceOwnership(&_UniversalVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalVerifier.Contract.RenounceOwnership(&_UniversalVerifier.TransactOpts)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0xc9c23ea5.
//
// Solidity: function setRequestOwner(uint64 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetRequestOwner(opts *bind.TransactOpts, requestId uint64, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setRequestOwner", requestId, requestOwner)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0xc9c23ea5.
//
// Solidity: function setRequestOwner(uint64 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetRequestOwner(requestId uint64, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequestOwner(&_UniversalVerifier.TransactOpts, requestId, requestOwner)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0xc9c23ea5.
//
// Solidity: function setRequestOwner(uint64 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetRequestOwner(requestId uint64, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequestOwner(&_UniversalVerifier.TransactOpts, requestId, requestOwner)
}

// SetState is a paid mutator transaction binding the contract method 0x34c901af.
//
// Solidity: function setState(address state) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetState(opts *bind.TransactOpts, state common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setState", state)
}

// SetState is a paid mutator transaction binding the contract method 0x34c901af.
//
// Solidity: function setState(address state) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetState(state common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetState(&_UniversalVerifier.TransactOpts, state)
}

// SetState is a paid mutator transaction binding the contract method 0x34c901af.
//
// Solidity: function setState(address state) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetState(state common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetState(&_UniversalVerifier.TransactOpts, state)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetZKPRequest(opts *bind.TransactOpts, requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setZKPRequest", requestId, request)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// SetZKPRequest is a paid mutator transaction binding the contract method 0x9f5223e0.
//
// Solidity: function setZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// SetZKPRequests is a paid mutator transaction binding the contract method 0x6882ee95.
//
// Solidity: function setZKPRequests(uint64[] requestIds, (string,address,bytes)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetZKPRequests(opts *bind.TransactOpts, requestIds []uint64, requests []IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setZKPRequests", requestIds, requests)
}

// SetZKPRequests is a paid mutator transaction binding the contract method 0x6882ee95.
//
// Solidity: function setZKPRequests(uint64[] requestIds, (string,address,bytes)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetZKPRequests(requestIds []uint64, requests []IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequests(&_UniversalVerifier.TransactOpts, requestIds, requests)
}

// SetZKPRequests is a paid mutator transaction binding the contract method 0x6882ee95.
//
// Solidity: function setZKPRequests(uint64[] requestIds, (string,address,bytes)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetZKPRequests(requestIds []uint64, requests []IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetZKPRequests(&_UniversalVerifier.TransactOpts, requestIds, requests)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponse(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponse", requestId, inputs, a, b, c)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c)
}

// SubmitZKPResponse is a paid mutator transaction binding the contract method 0xb68967e2.
//
// Solidity: function submitZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProof) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitZKPResponseV2(opts *bind.TransactOpts, responses []IZKPVerifierZKPResponse, crossChainProof []byte) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitZKPResponseV2", responses, crossChainProof)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProof) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitZKPResponseV2(responses []IZKPVerifierZKPResponse, crossChainProof []byte) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseV2(&_UniversalVerifier.TransactOpts, responses, crossChainProof)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProof) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitZKPResponseV2(responses []IZKPVerifierZKPResponse, crossChainProof []byte) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitZKPResponseV2(&_UniversalVerifier.TransactOpts, responses, crossChainProof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalVerifier *UniversalVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.TransferOwnership(&_UniversalVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.TransferOwnership(&_UniversalVerifier.TransactOpts, newOwner)
}

// UpdateZKPRequest is a paid mutator transaction binding the contract method 0x011f5bd6.
//
// Solidity: function updateZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) UpdateZKPRequest(opts *bind.TransactOpts, requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "updateZKPRequest", requestId, request)
}

// UpdateZKPRequest is a paid mutator transaction binding the contract method 0x011f5bd6.
//
// Solidity: function updateZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierSession) UpdateZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UpdateZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// UpdateZKPRequest is a paid mutator transaction binding the contract method 0x011f5bd6.
//
// Solidity: function updateZKPRequest(uint64 requestId, (string,address,bytes) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) UpdateZKPRequest(requestId uint64, request IZKPVerifierZKPRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UpdateZKPRequest(&_UniversalVerifier.TransactOpts, requestId, request)
}

// VerifyZKPResponse is a paid mutator transaction binding the contract method 0x5176983b.
//
// Solidity: function verifyZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address sender) returns((string,uint256)[])
func (_UniversalVerifier *UniversalVerifierTransactor) VerifyZKPResponse(opts *bind.TransactOpts, requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, sender common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "verifyZKPResponse", requestId, inputs, a, b, c, sender)
}

// VerifyZKPResponse is a paid mutator transaction binding the contract method 0x5176983b.
//
// Solidity: function verifyZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address sender) returns((string,uint256)[])
func (_UniversalVerifier *UniversalVerifierSession) VerifyZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, sender common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.VerifyZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, sender)
}

// VerifyZKPResponse is a paid mutator transaction binding the contract method 0x5176983b.
//
// Solidity: function verifyZKPResponse(uint64 requestId, uint256[] inputs, uint256[2] a, uint256[2][2] b, uint256[2] c, address sender) returns((string,uint256)[])
func (_UniversalVerifier *UniversalVerifierTransactorSession) VerifyZKPResponse(requestId uint64, inputs []*big.Int, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, sender common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.VerifyZKPResponse(&_UniversalVerifier.TransactOpts, requestId, inputs, a, b, c, sender)
}

// UniversalVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UniversalVerifier contract.
type UniversalVerifierInitializedIterator struct {
	Event *UniversalVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierInitialized)
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
		it.Event = new(UniversalVerifierInitialized)
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
func (it *UniversalVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierInitialized represents a Initialized event raised by the UniversalVerifier contract.
type UniversalVerifierInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*UniversalVerifierInitializedIterator, error) {

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierInitializedIterator{contract: _UniversalVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UniversalVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierInitialized)
				if err := _UniversalVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseInitialized(log types.Log) (*UniversalVerifierInitialized, error) {
	event := new(UniversalVerifierInitialized)
	if err := _UniversalVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferStartedIterator struct {
	Event *UniversalVerifierOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierOwnershipTransferStarted)
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
		it.Event = new(UniversalVerifierOwnershipTransferStarted)
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
func (it *UniversalVerifierOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalVerifierOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierOwnershipTransferStartedIterator{contract: _UniversalVerifier.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierOwnershipTransferStarted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseOwnershipTransferStarted(log types.Log) (*UniversalVerifierOwnershipTransferStarted, error) {
	event := new(UniversalVerifierOwnershipTransferStarted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferredIterator struct {
	Event *UniversalVerifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierOwnershipTransferred)
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
		it.Event = new(UniversalVerifierOwnershipTransferred)
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
func (it *UniversalVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the UniversalVerifier contract.
type UniversalVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierOwnershipTransferredIterator{contract: _UniversalVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniversalVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierOwnershipTransferred)
				if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UniversalVerifier *UniversalVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*UniversalVerifierOwnershipTransferred, error) {
	event := new(UniversalVerifierOwnershipTransferred)
	if err := _UniversalVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierZKPRequestSetIterator is returned from FilterZKPRequestSet and is used to iterate over the raw logs and unpacked data for ZKPRequestSet events raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestSetIterator struct {
	Event *UniversalVerifierZKPRequestSet // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierZKPRequestSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierZKPRequestSet)
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
		it.Event = new(UniversalVerifierZKPRequestSet)
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
func (it *UniversalVerifierZKPRequestSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierZKPRequestSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierZKPRequestSet represents a ZKPRequestSet event raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestSet struct {
	RequestId    uint64
	RequestOwner common.Address
	Metadata     string
	Validator    common.Address
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterZKPRequestSet is a free log retrieval operation binding the contract event 0x26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8.
//
// Solidity: event ZKPRequestSet(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterZKPRequestSet(opts *bind.FilterOpts, requestId []uint64, requestOwner []common.Address) (*UniversalVerifierZKPRequestSetIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ZKPRequestSet", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierZKPRequestSetIterator{contract: _UniversalVerifier.contract, event: "ZKPRequestSet", logs: logs, sub: sub}, nil
}

// WatchZKPRequestSet is a free log subscription operation binding the contract event 0x26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8.
//
// Solidity: event ZKPRequestSet(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchZKPRequestSet(opts *bind.WatchOpts, sink chan<- *UniversalVerifierZKPRequestSet, requestId []uint64, requestOwner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ZKPRequestSet", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierZKPRequestSet)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestSet", log); err != nil {
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

// ParseZKPRequestSet is a log parse operation binding the contract event 0x26db2c5d21d517fa8c11af3ae555af7f4a560b9f70a5a32e73e6cd94d0cae3e8.
//
// Solidity: event ZKPRequestSet(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseZKPRequestSet(log types.Log) (*UniversalVerifierZKPRequestSet, error) {
	event := new(UniversalVerifierZKPRequestSet)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierZKPRequestUpdateIterator is returned from FilterZKPRequestUpdate and is used to iterate over the raw logs and unpacked data for ZKPRequestUpdate events raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestUpdateIterator struct {
	Event *UniversalVerifierZKPRequestUpdate // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierZKPRequestUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierZKPRequestUpdate)
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
		it.Event = new(UniversalVerifierZKPRequestUpdate)
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
func (it *UniversalVerifierZKPRequestUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierZKPRequestUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierZKPRequestUpdate represents a ZKPRequestUpdate event raised by the UniversalVerifier contract.
type UniversalVerifierZKPRequestUpdate struct {
	RequestId    uint64
	RequestOwner common.Address
	Metadata     string
	Validator    common.Address
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterZKPRequestUpdate is a free log retrieval operation binding the contract event 0xa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0.
//
// Solidity: event ZKPRequestUpdate(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterZKPRequestUpdate(opts *bind.FilterOpts, requestId []uint64, requestOwner []common.Address) (*UniversalVerifierZKPRequestUpdateIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ZKPRequestUpdate", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierZKPRequestUpdateIterator{contract: _UniversalVerifier.contract, event: "ZKPRequestUpdate", logs: logs, sub: sub}, nil
}

// WatchZKPRequestUpdate is a free log subscription operation binding the contract event 0xa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0.
//
// Solidity: event ZKPRequestUpdate(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchZKPRequestUpdate(opts *bind.WatchOpts, sink chan<- *UniversalVerifierZKPRequestUpdate, requestId []uint64, requestOwner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ZKPRequestUpdate", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierZKPRequestUpdate)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestUpdate", log); err != nil {
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

// ParseZKPRequestUpdate is a log parse operation binding the contract event 0xa2c61fcbd9637e91178d0dea7f9b5cce13f60c453603b3dc056b8f01bb3d4cb0.
//
// Solidity: event ZKPRequestUpdate(uint64 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes data)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseZKPRequestUpdate(log types.Log) (*UniversalVerifierZKPRequestUpdate, error) {
	event := new(UniversalVerifierZKPRequestUpdate)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPRequestUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierZKPResponseSubmittedIterator is returned from FilterZKPResponseSubmitted and is used to iterate over the raw logs and unpacked data for ZKPResponseSubmitted events raised by the UniversalVerifier contract.
type UniversalVerifierZKPResponseSubmittedIterator struct {
	Event *UniversalVerifierZKPResponseSubmitted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierZKPResponseSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierZKPResponseSubmitted)
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
		it.Event = new(UniversalVerifierZKPResponseSubmitted)
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
func (it *UniversalVerifierZKPResponseSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierZKPResponseSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierZKPResponseSubmitted represents a ZKPResponseSubmitted event raised by the UniversalVerifier contract.
type UniversalVerifierZKPResponseSubmitted struct {
	RequestId uint64
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZKPResponseSubmitted is a free log retrieval operation binding the contract event 0x6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c.
//
// Solidity: event ZKPResponseSubmitted(uint64 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterZKPResponseSubmitted(opts *bind.FilterOpts, requestId []uint64, caller []common.Address) (*UniversalVerifierZKPResponseSubmittedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ZKPResponseSubmitted", requestIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierZKPResponseSubmittedIterator{contract: _UniversalVerifier.contract, event: "ZKPResponseSubmitted", logs: logs, sub: sub}, nil
}

// WatchZKPResponseSubmitted is a free log subscription operation binding the contract event 0x6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c.
//
// Solidity: event ZKPResponseSubmitted(uint64 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchZKPResponseSubmitted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierZKPResponseSubmitted, requestId []uint64, caller []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ZKPResponseSubmitted", requestIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierZKPResponseSubmitted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPResponseSubmitted", log); err != nil {
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

// ParseZKPResponseSubmitted is a log parse operation binding the contract event 0x6979bc9c3e552c05dd9859285f1ed7a172e52ef39e1dce9c720e5bf8d82c9f4c.
//
// Solidity: event ZKPResponseSubmitted(uint64 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseZKPResponseSubmitted(log types.Log) (*UniversalVerifierZKPResponseSubmitted, error) {
	event := new(UniversalVerifierZKPResponseSubmitted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ZKPResponseSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
