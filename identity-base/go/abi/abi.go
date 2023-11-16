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

// IOnchainCredentialStatusResolverCredentialStatus is an auto generated low-level Go binding around an user-defined struct.
type IOnchainCredentialStatusResolverCredentialStatus struct {
	Issuer IOnchainCredentialStatusResolverIdentityStateRoots
	Mtp    IOnchainCredentialStatusResolverProof
}

// IOnchainCredentialStatusResolverIdentityStateRoots is an auto generated low-level Go binding around an user-defined struct.
type IOnchainCredentialStatusResolverIdentityStateRoots struct {
	State              *big.Int
	ClaimsTreeRoot     *big.Int
	RevocationTreeRoot *big.Int
	RootOfRoots        *big.Int
}

// IOnchainCredentialStatusResolverProof is an auto generated low-level Go binding around an user-defined struct.
type IOnchainCredentialStatusResolverProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     []*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IdentityLibRoots is an auto generated low-level Go binding around an user-defined struct.
type IdentityLibRoots struct {
	ClaimsRoot      *big.Int
	RevocationsRoot *big.Int
	RootsRoot       *big.Int
}

// SmtLibProof is an auto generated low-level Go binding around an user-defined struct.
type SmtLibProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     []*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IdentityBaseMetaData contains all meta data concerning the IdentityBase contract.
var IdentityBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"}],\"name\":\"getClaimProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getClaimProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsOldStateGenesis\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedClaimsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRevocationsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRootsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRevocationProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatus\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatusByIdAndState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevocationsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"}],\"name\":\"getRootProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRootProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getRootsByState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.Roots\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmtDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateContractAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdentityBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityBaseMetaData.ABI instead.
var IdentityBaseABI = IdentityBaseMetaData.ABI

// IdentityBase is an auto generated Go binding around an Ethereum contract.
type IdentityBase struct {
	IdentityBaseCaller     // Read-only binding to the contract
	IdentityBaseTransactor // Write-only binding to the contract
	IdentityBaseFilterer   // Log filterer for contract events
}

// IdentityBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityBaseSession struct {
	Contract     *IdentityBase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityBaseCallerSession struct {
	Contract *IdentityBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IdentityBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityBaseTransactorSession struct {
	Contract     *IdentityBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IdentityBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityBaseRaw struct {
	Contract *IdentityBase // Generic contract binding to access the raw methods on
}

// IdentityBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityBaseCallerRaw struct {
	Contract *IdentityBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityBaseTransactorRaw struct {
	Contract *IdentityBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityBase creates a new instance of IdentityBase, bound to a specific deployed contract.
func NewIdentityBase(address common.Address, backend bind.ContractBackend) (*IdentityBase, error) {
	contract, err := bindIdentityBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityBase{IdentityBaseCaller: IdentityBaseCaller{contract: contract}, IdentityBaseTransactor: IdentityBaseTransactor{contract: contract}, IdentityBaseFilterer: IdentityBaseFilterer{contract: contract}}, nil
}

// NewIdentityBaseCaller creates a new read-only instance of IdentityBase, bound to a specific deployed contract.
func NewIdentityBaseCaller(address common.Address, caller bind.ContractCaller) (*IdentityBaseCaller, error) {
	contract, err := bindIdentityBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityBaseCaller{contract: contract}, nil
}

// NewIdentityBaseTransactor creates a new write-only instance of IdentityBase, bound to a specific deployed contract.
func NewIdentityBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityBaseTransactor, error) {
	contract, err := bindIdentityBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityBaseTransactor{contract: contract}, nil
}

// NewIdentityBaseFilterer creates a new log filterer instance of IdentityBase, bound to a specific deployed contract.
func NewIdentityBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityBaseFilterer, error) {
	contract, err := bindIdentityBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityBaseFilterer{contract: contract}, nil
}

// bindIdentityBase binds a generic wrapper to an already deployed contract.
func bindIdentityBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityBase *IdentityBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityBase.Contract.IdentityBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityBase *IdentityBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityBase.Contract.IdentityBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityBase *IdentityBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityBase.Contract.IdentityBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityBase *IdentityBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityBase *IdentityBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityBase *IdentityBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityBase.Contract.contract.Transact(opts, method, params...)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetClaimProof(opts *bind.CallOpts, claimIndexHash *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getClaimProof", claimIndexHash)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetClaimProof(&_IdentityBase.CallOpts, claimIndexHash)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetClaimProof(&_IdentityBase.CallOpts, claimIndexHash)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetClaimProofByRoot(opts *bind.CallOpts, claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getClaimProofByRoot", claimIndexHash, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetClaimProofByRoot(&_IdentityBase.CallOpts, claimIndexHash, root)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetClaimProofByRoot(&_IdentityBase.CallOpts, claimIndexHash, root)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetClaimsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getClaimsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetClaimsTreeRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetClaimsTreeRoot(&_IdentityBase.CallOpts)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetClaimsTreeRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetClaimsTreeRoot(&_IdentityBase.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetId() (*big.Int, error) {
	return _IdentityBase.Contract.GetId(&_IdentityBase.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetId() (*big.Int, error) {
	return _IdentityBase.Contract.GetId(&_IdentityBase.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_IdentityBase *IdentityBaseCaller) GetIsOldStateGenesis(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getIsOldStateGenesis")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_IdentityBase *IdentityBaseSession) GetIsOldStateGenesis() (bool, error) {
	return _IdentityBase.Contract.GetIsOldStateGenesis(&_IdentityBase.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_IdentityBase *IdentityBaseCallerSession) GetIsOldStateGenesis() (bool, error) {
	return _IdentityBase.Contract.GetIsOldStateGenesis(&_IdentityBase.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetLatestPublishedClaimsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getLatestPublishedClaimsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedClaimsRoot(&_IdentityBase.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedClaimsRoot(&_IdentityBase.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetLatestPublishedRevocationsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getLatestPublishedRevocationsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedRevocationsRoot(&_IdentityBase.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedRevocationsRoot(&_IdentityBase.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetLatestPublishedRootsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getLatestPublishedRootsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedRootsRoot(&_IdentityBase.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedRootsRoot(&_IdentityBase.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetLatestPublishedState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getLatestPublishedState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetLatestPublishedState() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedState(&_IdentityBase.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetLatestPublishedState() (*big.Int, error) {
	return _IdentityBase.Contract.GetLatestPublishedState(&_IdentityBase.CallOpts)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetRevocationProof(opts *bind.CallOpts, revocationNonce uint64) (SmtLibProof, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRevocationProof", revocationNonce)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRevocationProof(&_IdentityBase.CallOpts, revocationNonce)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRevocationProof(&_IdentityBase.CallOpts, revocationNonce)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetRevocationProofByRoot(opts *bind.CallOpts, revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRevocationProofByRoot", revocationNonce, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRevocationProofByRoot(&_IdentityBase.CallOpts, revocationNonce, root)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRevocationProofByRoot(&_IdentityBase.CallOpts, revocationNonce, root)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_IdentityBase *IdentityBaseCaller) GetRevocationStatus(opts *bind.CallOpts, id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRevocationStatus", id, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_IdentityBase *IdentityBaseSession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _IdentityBase.Contract.GetRevocationStatus(&_IdentityBase.CallOpts, id, nonce)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_IdentityBase *IdentityBaseCallerSession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _IdentityBase.Contract.GetRevocationStatus(&_IdentityBase.CallOpts, id, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_IdentityBase *IdentityBaseCaller) GetRevocationStatusByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRevocationStatusByIdAndState", id, state, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_IdentityBase *IdentityBaseSession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _IdentityBase.Contract.GetRevocationStatusByIdAndState(&_IdentityBase.CallOpts, id, state, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_IdentityBase *IdentityBaseCallerSession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _IdentityBase.Contract.GetRevocationStatusByIdAndState(&_IdentityBase.CallOpts, id, state, nonce)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetRevocationsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRevocationsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetRevocationsTreeRoot(&_IdentityBase.CallOpts)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetRevocationsTreeRoot(&_IdentityBase.CallOpts)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 claimsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetRootProof(opts *bind.CallOpts, claimsTreeRoot *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRootProof", claimsTreeRoot)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 claimsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetRootProof(claimsTreeRoot *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRootProof(&_IdentityBase.CallOpts, claimsTreeRoot)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 claimsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetRootProof(claimsTreeRoot *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRootProof(&_IdentityBase.CallOpts, claimsTreeRoot)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetRootProofByRoot(opts *bind.CallOpts, claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRootProofByRoot", claimsTreeRoot, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRootProofByRoot(&_IdentityBase.CallOpts, claimsTreeRoot, root)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _IdentityBase.Contract.GetRootProofByRoot(&_IdentityBase.CallOpts, claimsTreeRoot, root)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_IdentityBase *IdentityBaseCaller) GetRootsByState(opts *bind.CallOpts, state *big.Int) (IdentityLibRoots, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRootsByState", state)

	if err != nil {
		return *new(IdentityLibRoots), err
	}

	out0 := *abi.ConvertType(out[0], new(IdentityLibRoots)).(*IdentityLibRoots)

	return out0, err

}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_IdentityBase *IdentityBaseSession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _IdentityBase.Contract.GetRootsByState(&_IdentityBase.CallOpts, state)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_IdentityBase *IdentityBaseCallerSession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _IdentityBase.Contract.GetRootsByState(&_IdentityBase.CallOpts, state)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetRootsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getRootsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetRootsTreeRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetRootsTreeRoot(&_IdentityBase.CallOpts)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetRootsTreeRoot() (*big.Int, error) {
	return _IdentityBase.Contract.GetRootsTreeRoot(&_IdentityBase.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_IdentityBase *IdentityBaseCaller) GetSmtDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityBase.contract.Call(opts, &out, "getSmtDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_IdentityBase *IdentityBaseSession) GetSmtDepth() (*big.Int, error) {
	return _IdentityBase.Contract.GetSmtDepth(&_IdentityBase.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_IdentityBase *IdentityBaseCallerSession) GetSmtDepth() (*big.Int, error) {
	return _IdentityBase.Contract.GetSmtDepth(&_IdentityBase.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_IdentityBase *IdentityBaseTransactor) Initialize(opts *bind.TransactOpts, _stateContractAddr common.Address) (*types.Transaction, error) {
	return _IdentityBase.contract.Transact(opts, "initialize", _stateContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_IdentityBase *IdentityBaseSession) Initialize(_stateContractAddr common.Address) (*types.Transaction, error) {
	return _IdentityBase.Contract.Initialize(&_IdentityBase.TransactOpts, _stateContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_IdentityBase *IdentityBaseTransactorSession) Initialize(_stateContractAddr common.Address) (*types.Transaction, error) {
	return _IdentityBase.Contract.Initialize(&_IdentityBase.TransactOpts, _stateContractAddr)
}
