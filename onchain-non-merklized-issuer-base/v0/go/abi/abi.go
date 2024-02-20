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

// INonMerklizedIssuerCredentialData is an auto generated low-level Go binding around an user-defined struct.
type INonMerklizedIssuerCredentialData struct {
	Id               *big.Int
	Context          []string
	Type             string
	IssuanceDate     uint64
	CredentialSchema string
	DisplayMethod    INonMerklizedIssuerDisplayMethod
}

// INonMerklizedIssuerDisplayMethod is an auto generated low-level Go binding around an user-defined struct.
type INonMerklizedIssuerDisplayMethod struct {
	Id   string
	Type string
}

// INonMerklizedIssuerSubjectField is an auto generated low-level Go binding around an user-defined struct.
type INonMerklizedIssuerSubjectField struct {
	Key      string
	Value    *big.Int
	RawValue []byte
}

// NonMerklizedIssuerBaseMetaData contains all meta data concerning the NonMerklizedIssuerBase contract.
var NonMerklizedIssuerBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CREDENTIAL_PROTOCOL_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_credentialId\",\"type\":\"uint256\"}],\"name\":\"getCredential\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"context\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issuanceDate\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"credentialSchema\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"}],\"internalType\":\"structINonMerklizedIssuer.DisplayMethod\",\"name\":\"displayMethod\",\"type\":\"tuple\"}],\"internalType\":\"structINonMerklizedIssuer.CredentialData\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawValue\",\"type\":\"bytes\"}],\"internalType\":\"structINonMerklizedIssuer.SubjectField[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCredentialProtocolVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"}],\"name\":\"listUserCredentialIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NonMerklizedIssuerBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use NonMerklizedIssuerBaseMetaData.ABI instead.
var NonMerklizedIssuerBaseABI = NonMerklizedIssuerBaseMetaData.ABI

// NonMerklizedIssuerBase is an auto generated Go binding around an Ethereum contract.
type NonMerklizedIssuerBase struct {
	NonMerklizedIssuerBaseCaller     // Read-only binding to the contract
	NonMerklizedIssuerBaseTransactor // Write-only binding to the contract
	NonMerklizedIssuerBaseFilterer   // Log filterer for contract events
}

// NonMerklizedIssuerBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonMerklizedIssuerBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonMerklizedIssuerBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonMerklizedIssuerBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonMerklizedIssuerBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonMerklizedIssuerBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonMerklizedIssuerBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonMerklizedIssuerBaseSession struct {
	Contract     *NonMerklizedIssuerBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NonMerklizedIssuerBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonMerklizedIssuerBaseCallerSession struct {
	Contract *NonMerklizedIssuerBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NonMerklizedIssuerBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonMerklizedIssuerBaseTransactorSession struct {
	Contract     *NonMerklizedIssuerBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NonMerklizedIssuerBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonMerklizedIssuerBaseRaw struct {
	Contract *NonMerklizedIssuerBase // Generic contract binding to access the raw methods on
}

// NonMerklizedIssuerBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonMerklizedIssuerBaseCallerRaw struct {
	Contract *NonMerklizedIssuerBaseCaller // Generic read-only contract binding to access the raw methods on
}

// NonMerklizedIssuerBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonMerklizedIssuerBaseTransactorRaw struct {
	Contract *NonMerklizedIssuerBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonMerklizedIssuerBase creates a new instance of NonMerklizedIssuerBase, bound to a specific deployed contract.
func NewNonMerklizedIssuerBase(address common.Address, backend bind.ContractBackend) (*NonMerklizedIssuerBase, error) {
	contract, err := bindNonMerklizedIssuerBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonMerklizedIssuerBase{NonMerklizedIssuerBaseCaller: NonMerklizedIssuerBaseCaller{contract: contract}, NonMerklizedIssuerBaseTransactor: NonMerklizedIssuerBaseTransactor{contract: contract}, NonMerklizedIssuerBaseFilterer: NonMerklizedIssuerBaseFilterer{contract: contract}}, nil
}

// NewNonMerklizedIssuerBaseCaller creates a new read-only instance of NonMerklizedIssuerBase, bound to a specific deployed contract.
func NewNonMerklizedIssuerBaseCaller(address common.Address, caller bind.ContractCaller) (*NonMerklizedIssuerBaseCaller, error) {
	contract, err := bindNonMerklizedIssuerBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonMerklizedIssuerBaseCaller{contract: contract}, nil
}

// NewNonMerklizedIssuerBaseTransactor creates a new write-only instance of NonMerklizedIssuerBase, bound to a specific deployed contract.
func NewNonMerklizedIssuerBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*NonMerklizedIssuerBaseTransactor, error) {
	contract, err := bindNonMerklizedIssuerBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonMerklizedIssuerBaseTransactor{contract: contract}, nil
}

// NewNonMerklizedIssuerBaseFilterer creates a new log filterer instance of NonMerklizedIssuerBase, bound to a specific deployed contract.
func NewNonMerklizedIssuerBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*NonMerklizedIssuerBaseFilterer, error) {
	contract, err := bindNonMerklizedIssuerBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonMerklizedIssuerBaseFilterer{contract: contract}, nil
}

// bindNonMerklizedIssuerBase binds a generic wrapper to an already deployed contract.
func bindNonMerklizedIssuerBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonMerklizedIssuerBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonMerklizedIssuerBase.Contract.NonMerklizedIssuerBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.NonMerklizedIssuerBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.NonMerklizedIssuerBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonMerklizedIssuerBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.contract.Transact(opts, method, params...)
}

// CREDENTIALPROTOCOLVERSION is a free data retrieval call binding the contract method 0x8e44f31b.
//
// Solidity: function CREDENTIAL_PROTOCOL_VERSION() view returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) CREDENTIALPROTOCOLVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "CREDENTIAL_PROTOCOL_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CREDENTIALPROTOCOLVERSION is a free data retrieval call binding the contract method 0x8e44f31b.
//
// Solidity: function CREDENTIAL_PROTOCOL_VERSION() view returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) CREDENTIALPROTOCOLVERSION() (string, error) {
	return _NonMerklizedIssuerBase.Contract.CREDENTIALPROTOCOLVERSION(&_NonMerklizedIssuerBase.CallOpts)
}

// CREDENTIALPROTOCOLVERSION is a free data retrieval call binding the contract method 0x8e44f31b.
//
// Solidity: function CREDENTIAL_PROTOCOL_VERSION() view returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) CREDENTIALPROTOCOLVERSION() (string, error) {
	return _NonMerklizedIssuerBase.Contract.CREDENTIALPROTOCOLVERSION(&_NonMerklizedIssuerBase.CallOpts)
}

// GetCredentialProtocolVersion is a free data retrieval call binding the contract method 0xb9f5eb4d.
//
// Solidity: function getCredentialProtocolVersion() pure returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetCredentialProtocolVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getCredentialProtocolVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCredentialProtocolVersion is a free data retrieval call binding the contract method 0xb9f5eb4d.
//
// Solidity: function getCredentialProtocolVersion() pure returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetCredentialProtocolVersion() (string, error) {
	return _NonMerklizedIssuerBase.Contract.GetCredentialProtocolVersion(&_NonMerklizedIssuerBase.CallOpts)
}

// GetCredentialProtocolVersion is a free data retrieval call binding the contract method 0xb9f5eb4d.
//
// Solidity: function getCredentialProtocolVersion() pure returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetCredentialProtocolVersion() (string, error) {
	return _NonMerklizedIssuerBase.Contract.GetCredentialProtocolVersion(&_NonMerklizedIssuerBase.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NonMerklizedIssuerBase.Contract.SupportsInterface(&_NonMerklizedIssuerBase.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NonMerklizedIssuerBase.Contract.SupportsInterface(&_NonMerklizedIssuerBase.CallOpts, interfaceId)
}

// GetCredential is a paid mutator transaction binding the contract method 0x37c1d9ff.
//
// Solidity: function getCredential(uint256 _userId, uint256 _credentialId) returns((uint256,string[],string,uint64,string,(string,string)), uint256[8], (string,uint256,bytes)[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) GetCredential(opts *bind.TransactOpts, _userId *big.Int, _credentialId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "getCredential", _userId, _credentialId)
}

// GetCredential is a paid mutator transaction binding the contract method 0x37c1d9ff.
//
// Solidity: function getCredential(uint256 _userId, uint256 _credentialId) returns((uint256,string[],string,uint64,string,(string,string)), uint256[8], (string,uint256,bytes)[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetCredential(_userId *big.Int, _credentialId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.GetCredential(&_NonMerklizedIssuerBase.TransactOpts, _userId, _credentialId)
}

// GetCredential is a paid mutator transaction binding the contract method 0x37c1d9ff.
//
// Solidity: function getCredential(uint256 _userId, uint256 _credentialId) returns((uint256,string[],string,uint64,string,(string,string)), uint256[8], (string,uint256,bytes)[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) GetCredential(_userId *big.Int, _credentialId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.GetCredential(&_NonMerklizedIssuerBase.TransactOpts, _userId, _credentialId)
}

// ListUserCredentialIds is a paid mutator transaction binding the contract method 0x58381619.
//
// Solidity: function listUserCredentialIds(uint256 _userId) returns(uint256[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) ListUserCredentialIds(opts *bind.TransactOpts, _userId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "listUserCredentialIds", _userId)
}

// ListUserCredentialIds is a paid mutator transaction binding the contract method 0x58381619.
//
// Solidity: function listUserCredentialIds(uint256 _userId) returns(uint256[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) ListUserCredentialIds(_userId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.ListUserCredentialIds(&_NonMerklizedIssuerBase.TransactOpts, _userId)
}

// ListUserCredentialIds is a paid mutator transaction binding the contract method 0x58381619.
//
// Solidity: function listUserCredentialIds(uint256 _userId) returns(uint256[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) ListUserCredentialIds(_userId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.ListUserCredentialIds(&_NonMerklizedIssuerBase.TransactOpts, _userId)
}
