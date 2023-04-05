// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package state_contract

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

// IStateGistProof is an auto generated low-level Go binding around an user-defined struct.
type IStateGistProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     [64]*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IStateGistRootInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateGistRootInfo struct {
	Root                *big.Int
	ReplacedByRoot      *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// IStateStateInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateStateInfo struct {
	Id                  *big.Int
	State               *big.Int
	ReplacedByState     *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// StateContractMetaData contains all meta data concerning the StateContract contract.
var StateContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockN\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"StateUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGISTProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getGISTRootHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRootHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateInfoById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getStateInfoByIdAndState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryLengthById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"idExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStateTransitionVerifier\",\"name\":\"verifierContractAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifierAddr\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"stateExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOldStateGenesis\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"transitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StateContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StateContractMetaData.ABI instead.
var StateContractABI = StateContractMetaData.ABI

// StateContract is an auto generated Go binding around an Ethereum contract.
type StateContract struct {
	StateContractCaller     // Read-only binding to the contract
	StateContractTransactor // Write-only binding to the contract
	StateContractFilterer   // Log filterer for contract events
}

// StateContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateContractSession struct {
	Contract     *StateContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateContractCallerSession struct {
	Contract *StateContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StateContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateContractTransactorSession struct {
	Contract     *StateContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StateContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateContractRaw struct {
	Contract *StateContract // Generic contract binding to access the raw methods on
}

// StateContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateContractCallerRaw struct {
	Contract *StateContractCaller // Generic read-only contract binding to access the raw methods on
}

// StateContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateContractTransactorRaw struct {
	Contract *StateContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateContract creates a new instance of StateContract, bound to a specific deployed contract.
func NewStateContract(address common.Address, backend bind.ContractBackend) (*StateContract, error) {
	contract, err := bindStateContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateContract{StateContractCaller: StateContractCaller{contract: contract}, StateContractTransactor: StateContractTransactor{contract: contract}, StateContractFilterer: StateContractFilterer{contract: contract}}, nil
}

// NewStateContractCaller creates a new read-only instance of StateContract, bound to a specific deployed contract.
func NewStateContractCaller(address common.Address, caller bind.ContractCaller) (*StateContractCaller, error) {
	contract, err := bindStateContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateContractCaller{contract: contract}, nil
}

// NewStateContractTransactor creates a new write-only instance of StateContract, bound to a specific deployed contract.
func NewStateContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StateContractTransactor, error) {
	contract, err := bindStateContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateContractTransactor{contract: contract}, nil
}

// NewStateContractFilterer creates a new log filterer instance of StateContract, bound to a specific deployed contract.
func NewStateContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StateContractFilterer, error) {
	contract, err := bindStateContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateContractFilterer{contract: contract}, nil
}

// bindStateContract binds a generic wrapper to an already deployed contract.
func bindStateContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateContract *StateContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateContract.Contract.StateContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateContract *StateContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateContract.Contract.StateContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateContract *StateContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateContract.Contract.StateContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateContract *StateContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateContract *StateContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateContract *StateContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateContract.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateContract *StateContractCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateContract *StateContractSession) VERSION() (string, error) {
	return _StateContract.Contract.VERSION(&_StateContract.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateContract *StateContractCallerSession) VERSION() (string, error) {
	return _StateContract.Contract.VERSION(&_StateContract.CallOpts)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTProof(opts *bind.CallOpts, id *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTProof", id)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTProof(id *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProof(&_StateContract.CallOpts, id)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTProof(id *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProof(&_StateContract.CallOpts, id)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTProofByBlock(opts *bind.CallOpts, id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTProofByBlock", id, blockNumber)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTProofByBlock(id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProofByBlock(&_StateContract.CallOpts, id, blockNumber)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTProofByBlock(id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProofByBlock(&_StateContract.CallOpts, id, blockNumber)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTProofByRoot(opts *bind.CallOpts, id *big.Int, root *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTProofByRoot", id, root)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTProofByRoot(id *big.Int, root *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProofByRoot(&_StateContract.CallOpts, id, root)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTProofByRoot(id *big.Int, root *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProofByRoot(&_StateContract.CallOpts, id, root)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTProofByTime(opts *bind.CallOpts, id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTProofByTime", id, timestamp)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTProofByTime(id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProofByTime(&_StateContract.CallOpts, id, timestamp)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTProofByTime(id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	return _StateContract.Contract.GetGISTProofByTime(&_StateContract.CallOpts, id, timestamp)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateContract *StateContractCaller) GetGISTRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateContract *StateContractSession) GetGISTRoot() (*big.Int, error) {
	return _StateContract.Contract.GetGISTRoot(&_StateContract.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateContract *StateContractCallerSession) GetGISTRoot() (*big.Int, error) {
	return _StateContract.Contract.GetGISTRoot(&_StateContract.CallOpts)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateContract *StateContractCaller) GetGISTRootHistory(opts *bind.CallOpts, start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTRootHistory", start, length)

	if err != nil {
		return *new([]IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateGistRootInfo)).(*[]IStateGistRootInfo)

	return out0, err

}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateContract *StateContractSession) GetGISTRootHistory(start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootHistory(&_StateContract.CallOpts, start, length)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateContract *StateContractCallerSession) GetGISTRootHistory(start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootHistory(&_StateContract.CallOpts, start, length)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateContract *StateContractCaller) GetGISTRootHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTRootHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateContract *StateContractSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _StateContract.Contract.GetGISTRootHistoryLength(&_StateContract.CallOpts)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateContract *StateContractCallerSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _StateContract.Contract.GetGISTRootHistoryLength(&_StateContract.CallOpts)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTRootInfo(opts *bind.CallOpts, root *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTRootInfo", root)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTRootInfo(root *big.Int) (IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootInfo(&_StateContract.CallOpts, root)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTRootInfo(root *big.Int) (IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootInfo(&_StateContract.CallOpts, root)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTRootInfoByBlock(opts *bind.CallOpts, blockNumber *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTRootInfoByBlock", blockNumber)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTRootInfoByBlock(blockNumber *big.Int) (IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootInfoByBlock(&_StateContract.CallOpts, blockNumber)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTRootInfoByBlock(blockNumber *big.Int) (IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootInfoByBlock(&_StateContract.CallOpts, blockNumber)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCaller) GetGISTRootInfoByTime(opts *bind.CallOpts, timestamp *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getGISTRootInfoByTime", timestamp)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractSession) GetGISTRootInfoByTime(timestamp *big.Int) (IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootInfoByTime(&_StateContract.CallOpts, timestamp)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetGISTRootInfoByTime(timestamp *big.Int) (IStateGistRootInfo, error) {
	return _StateContract.Contract.GetGISTRootInfoByTime(&_StateContract.CallOpts, timestamp)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCaller) GetStateInfoById(opts *bind.CallOpts, id *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getStateInfoById", id)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractSession) GetStateInfoById(id *big.Int) (IStateStateInfo, error) {
	return _StateContract.Contract.GetStateInfoById(&_StateContract.CallOpts, id)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetStateInfoById(id *big.Int) (IStateStateInfo, error) {
	return _StateContract.Contract.GetStateInfoById(&_StateContract.CallOpts, id)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCaller) GetStateInfoByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getStateInfoByIdAndState", id, state)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractSession) GetStateInfoByIdAndState(id *big.Int, state *big.Int) (IStateStateInfo, error) {
	return _StateContract.Contract.GetStateInfoByIdAndState(&_StateContract.CallOpts, id, state)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateContract *StateContractCallerSession) GetStateInfoByIdAndState(id *big.Int, state *big.Int) (IStateStateInfo, error) {
	return _StateContract.Contract.GetStateInfoByIdAndState(&_StateContract.CallOpts, id, state)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateContract *StateContractCaller) GetStateInfoHistoryById(opts *bind.CallOpts, id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getStateInfoHistoryById", id, startIndex, length)

	if err != nil {
		return *new([]IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateStateInfo)).(*[]IStateStateInfo)

	return out0, err

}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateContract *StateContractSession) GetStateInfoHistoryById(id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	return _StateContract.Contract.GetStateInfoHistoryById(&_StateContract.CallOpts, id, startIndex, length)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateContract *StateContractCallerSession) GetStateInfoHistoryById(id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	return _StateContract.Contract.GetStateInfoHistoryById(&_StateContract.CallOpts, id, startIndex, length)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateContract *StateContractCaller) GetStateInfoHistoryLengthById(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getStateInfoHistoryLengthById", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateContract *StateContractSession) GetStateInfoHistoryLengthById(id *big.Int) (*big.Int, error) {
	return _StateContract.Contract.GetStateInfoHistoryLengthById(&_StateContract.CallOpts, id)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateContract *StateContractCallerSession) GetStateInfoHistoryLengthById(id *big.Int) (*big.Int, error) {
	return _StateContract.Contract.GetStateInfoHistoryLengthById(&_StateContract.CallOpts, id)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateContract *StateContractCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateContract *StateContractSession) GetVerifier() (common.Address, error) {
	return _StateContract.Contract.GetVerifier(&_StateContract.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateContract *StateContractCallerSession) GetVerifier() (common.Address, error) {
	return _StateContract.Contract.GetVerifier(&_StateContract.CallOpts)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateContract *StateContractCaller) IdExists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "idExists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateContract *StateContractSession) IdExists(id *big.Int) (bool, error) {
	return _StateContract.Contract.IdExists(&_StateContract.CallOpts, id)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateContract *StateContractCallerSession) IdExists(id *big.Int) (bool, error) {
	return _StateContract.Contract.IdExists(&_StateContract.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateContract *StateContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateContract *StateContractSession) Owner() (common.Address, error) {
	return _StateContract.Contract.Owner(&_StateContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateContract *StateContractCallerSession) Owner() (common.Address, error) {
	return _StateContract.Contract.Owner(&_StateContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateContract *StateContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateContract *StateContractSession) PendingOwner() (common.Address, error) {
	return _StateContract.Contract.PendingOwner(&_StateContract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateContract *StateContractCallerSession) PendingOwner() (common.Address, error) {
	return _StateContract.Contract.PendingOwner(&_StateContract.CallOpts)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateContract *StateContractCaller) StateExists(opts *bind.CallOpts, id *big.Int, state *big.Int) (bool, error) {
	var out []interface{}
	err := _StateContract.contract.Call(opts, &out, "stateExists", id, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateContract *StateContractSession) StateExists(id *big.Int, state *big.Int) (bool, error) {
	return _StateContract.Contract.StateExists(&_StateContract.CallOpts, id, state)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateContract *StateContractCallerSession) StateExists(id *big.Int, state *big.Int) (bool, error) {
	return _StateContract.Contract.StateExists(&_StateContract.CallOpts, id, state)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateContract *StateContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateContract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateContract *StateContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _StateContract.Contract.AcceptOwnership(&_StateContract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateContract *StateContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _StateContract.Contract.AcceptOwnership(&_StateContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address verifierContractAddr) returns()
func (_StateContract *StateContractTransactor) Initialize(opts *bind.TransactOpts, verifierContractAddr common.Address) (*types.Transaction, error) {
	return _StateContract.contract.Transact(opts, "initialize", verifierContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address verifierContractAddr) returns()
func (_StateContract *StateContractSession) Initialize(verifierContractAddr common.Address) (*types.Transaction, error) {
	return _StateContract.Contract.Initialize(&_StateContract.TransactOpts, verifierContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address verifierContractAddr) returns()
func (_StateContract *StateContractTransactorSession) Initialize(verifierContractAddr common.Address) (*types.Transaction, error) {
	return _StateContract.Contract.Initialize(&_StateContract.TransactOpts, verifierContractAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateContract *StateContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateContract *StateContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateContract.Contract.RenounceOwnership(&_StateContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateContract *StateContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateContract.Contract.RenounceOwnership(&_StateContract.TransactOpts)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateContract *StateContractTransactor) SetVerifier(opts *bind.TransactOpts, newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateContract.contract.Transact(opts, "setVerifier", newVerifierAddr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateContract *StateContractSession) SetVerifier(newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateContract.Contract.SetVerifier(&_StateContract.TransactOpts, newVerifierAddr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateContract *StateContractTransactorSession) SetVerifier(newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateContract.Contract.SetVerifier(&_StateContract.TransactOpts, newVerifierAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateContract *StateContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StateContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateContract *StateContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateContract.Contract.TransferOwnership(&_StateContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateContract *StateContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateContract.Contract.TransferOwnership(&_StateContract.TransactOpts, newOwner)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateContract *StateContractTransactor) TransitState(opts *bind.TransactOpts, id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateContract.contract.Transact(opts, "transitState", id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateContract *StateContractSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateContract.Contract.TransitState(&_StateContract.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateContract *StateContractTransactorSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateContract.Contract.TransitState(&_StateContract.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// StateContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StateContract contract.
type StateContractInitializedIterator struct {
	Event *StateContractInitialized // Event containing the contract specifics and raw log

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
func (it *StateContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateContractInitialized)
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
		it.Event = new(StateContractInitialized)
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
func (it *StateContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateContractInitialized represents a Initialized event raised by the StateContract contract.
type StateContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateContract *StateContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*StateContractInitializedIterator, error) {

	logs, sub, err := _StateContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StateContractInitializedIterator{contract: _StateContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateContract *StateContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StateContractInitialized) (event.Subscription, error) {

	logs, sub, err := _StateContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateContractInitialized)
				if err := _StateContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StateContract *StateContractFilterer) ParseInitialized(log types.Log) (*StateContractInitialized, error) {
	event := new(StateContractInitialized)
	if err := _StateContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateContractOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the StateContract contract.
type StateContractOwnershipTransferStartedIterator struct {
	Event *StateContractOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *StateContractOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateContractOwnershipTransferStarted)
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
		it.Event = new(StateContractOwnershipTransferStarted)
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
func (it *StateContractOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateContractOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateContractOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the StateContract contract.
type StateContractOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateContract *StateContractFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateContractOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateContract.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateContractOwnershipTransferStartedIterator{contract: _StateContract.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateContract *StateContractFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *StateContractOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateContract.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateContractOwnershipTransferStarted)
				if err := _StateContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_StateContract *StateContractFilterer) ParseOwnershipTransferStarted(log types.Log) (*StateContractOwnershipTransferStarted, error) {
	event := new(StateContractOwnershipTransferStarted)
	if err := _StateContract.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StateContract contract.
type StateContractOwnershipTransferredIterator struct {
	Event *StateContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StateContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateContractOwnershipTransferred)
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
		it.Event = new(StateContractOwnershipTransferred)
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
func (it *StateContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateContractOwnershipTransferred represents a OwnershipTransferred event raised by the StateContract contract.
type StateContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateContract *StateContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateContractOwnershipTransferredIterator{contract: _StateContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateContract *StateContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StateContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateContractOwnershipTransferred)
				if err := _StateContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StateContract *StateContractFilterer) ParseOwnershipTransferred(log types.Log) (*StateContractOwnershipTransferred, error) {
	event := new(StateContractOwnershipTransferred)
	if err := _StateContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateContractStateUpdatedIterator is returned from FilterStateUpdated and is used to iterate over the raw logs and unpacked data for StateUpdated events raised by the StateContract contract.
type StateContractStateUpdatedIterator struct {
	Event *StateContractStateUpdated // Event containing the contract specifics and raw log

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
func (it *StateContractStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateContractStateUpdated)
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
		it.Event = new(StateContractStateUpdated)
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
func (it *StateContractStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateContractStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateContractStateUpdated represents a StateUpdated event raised by the StateContract contract.
type StateContractStateUpdated struct {
	Id        *big.Int
	BlockN    *big.Int
	Timestamp *big.Int
	State     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateUpdated is a free log retrieval operation binding the contract event 0x88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a8838.
//
// Solidity: event StateUpdated(uint256 id, uint256 blockN, uint256 timestamp, uint256 state)
func (_StateContract *StateContractFilterer) FilterStateUpdated(opts *bind.FilterOpts) (*StateContractStateUpdatedIterator, error) {

	logs, sub, err := _StateContract.contract.FilterLogs(opts, "StateUpdated")
	if err != nil {
		return nil, err
	}
	return &StateContractStateUpdatedIterator{contract: _StateContract.contract, event: "StateUpdated", logs: logs, sub: sub}, nil
}

// WatchStateUpdated is a free log subscription operation binding the contract event 0x88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a8838.
//
// Solidity: event StateUpdated(uint256 id, uint256 blockN, uint256 timestamp, uint256 state)
func (_StateContract *StateContractFilterer) WatchStateUpdated(opts *bind.WatchOpts, sink chan<- *StateContractStateUpdated) (event.Subscription, error) {

	logs, sub, err := _StateContract.contract.WatchLogs(opts, "StateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateContractStateUpdated)
				if err := _StateContract.contract.UnpackLog(event, "StateUpdated", log); err != nil {
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

// ParseStateUpdated is a log parse operation binding the contract event 0x88aef4d78ad30d12a12a98e96007f5b09c1610b5364b2b99960b7d07e00a8838.
//
// Solidity: event StateUpdated(uint256 id, uint256 blockN, uint256 timestamp, uint256 state)
func (_StateContract *StateContractFilterer) ParseStateUpdated(log types.Log) (*StateContractStateUpdated, error) {
	event := new(StateContractStateUpdated)
	if err := _StateContract.contract.UnpackLog(event, "StateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
