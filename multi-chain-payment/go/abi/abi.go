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

// MCPaymentIden3PaymentRailsERC20RequestV1 is an auto generated low-level Go binding around an user-defined struct.
type MCPaymentIden3PaymentRailsERC20RequestV1 struct {
	TokenAddress   common.Address
	Recipient      common.Address
	Amount         *big.Int
	ExpirationDate *big.Int
	Nonce          *big.Int
	Metadata       []byte
}

// MCPaymentIden3PaymentRailsRequestV1 is an auto generated low-level Go binding around an user-defined struct.
type MCPaymentIden3PaymentRailsRequestV1 struct {
	Recipient      common.Address
	Amount         *big.Int
	ExpirationDate *big.Int
	Nonce          *big.Int
	Metadata       []byte
}

// MCPaymentMetaData contains all meta data concerning the MCPayment contract.
var MCPaymentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidOwnerPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"PaymentError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"WithdrawError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC_20_PAYMENT_DATA_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAYMENT_DATA_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnerPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"ownerPercentage\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"isPaymentDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuerWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structMCPayment.Iden3PaymentRailsRequestV1\",\"name\":\"paymentData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structMCPayment.Iden3PaymentRailsERC20RequestV1\",\"name\":\"paymentData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"payERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"permitSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structMCPayment.Iden3PaymentRailsERC20RequestV1\",\"name\":\"paymentData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"payERC20Permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structMCPayment.Iden3PaymentRailsERC20RequestV1\",\"name\":\"paymentData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverIden3PaymentRailsERC20RequestV1Signature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structMCPayment.Iden3PaymentRailsRequestV1\",\"name\":\"paymentData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recoverIden3PaymentRailsRequestV1Signature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ownerPercentage\",\"type\":\"uint8\"}],\"name\":\"updateOwnerPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MCPaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use MCPaymentMetaData.ABI instead.
var MCPaymentABI = MCPaymentMetaData.ABI

// MCPayment is an auto generated Go binding around an Ethereum contract.
type MCPayment struct {
	MCPaymentCaller     // Read-only binding to the contract
	MCPaymentTransactor // Write-only binding to the contract
	MCPaymentFilterer   // Log filterer for contract events
}

// MCPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCPaymentSession struct {
	Contract     *MCPayment        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCPaymentCallerSession struct {
	Contract *MCPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MCPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCPaymentTransactorSession struct {
	Contract     *MCPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MCPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCPaymentRaw struct {
	Contract *MCPayment // Generic contract binding to access the raw methods on
}

// MCPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCPaymentCallerRaw struct {
	Contract *MCPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// MCPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCPaymentTransactorRaw struct {
	Contract *MCPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCPayment creates a new instance of MCPayment, bound to a specific deployed contract.
func NewMCPayment(address common.Address, backend bind.ContractBackend) (*MCPayment, error) {
	contract, err := bindMCPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCPayment{MCPaymentCaller: MCPaymentCaller{contract: contract}, MCPaymentTransactor: MCPaymentTransactor{contract: contract}, MCPaymentFilterer: MCPaymentFilterer{contract: contract}}, nil
}

// NewMCPaymentCaller creates a new read-only instance of MCPayment, bound to a specific deployed contract.
func NewMCPaymentCaller(address common.Address, caller bind.ContractCaller) (*MCPaymentCaller, error) {
	contract, err := bindMCPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCPaymentCaller{contract: contract}, nil
}

// NewMCPaymentTransactor creates a new write-only instance of MCPayment, bound to a specific deployed contract.
func NewMCPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*MCPaymentTransactor, error) {
	contract, err := bindMCPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCPaymentTransactor{contract: contract}, nil
}

// NewMCPaymentFilterer creates a new log filterer instance of MCPayment, bound to a specific deployed contract.
func NewMCPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*MCPaymentFilterer, error) {
	contract, err := bindMCPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCPaymentFilterer{contract: contract}, nil
}

// bindMCPayment binds a generic wrapper to an already deployed contract.
func bindMCPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MCPaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCPayment *MCPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCPayment.Contract.MCPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCPayment *MCPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCPayment.Contract.MCPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCPayment *MCPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCPayment.Contract.MCPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCPayment *MCPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCPayment *MCPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCPayment *MCPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCPayment.Contract.contract.Transact(opts, method, params...)
}

// ERC20PAYMENTDATATYPEHASH is a free data retrieval call binding the contract method 0xc6bfaa3f.
//
// Solidity: function ERC_20_PAYMENT_DATA_TYPE_HASH() view returns(bytes32)
func (_MCPayment *MCPaymentCaller) ERC20PAYMENTDATATYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "ERC_20_PAYMENT_DATA_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20PAYMENTDATATYPEHASH is a free data retrieval call binding the contract method 0xc6bfaa3f.
//
// Solidity: function ERC_20_PAYMENT_DATA_TYPE_HASH() view returns(bytes32)
func (_MCPayment *MCPaymentSession) ERC20PAYMENTDATATYPEHASH() ([32]byte, error) {
	return _MCPayment.Contract.ERC20PAYMENTDATATYPEHASH(&_MCPayment.CallOpts)
}

// ERC20PAYMENTDATATYPEHASH is a free data retrieval call binding the contract method 0xc6bfaa3f.
//
// Solidity: function ERC_20_PAYMENT_DATA_TYPE_HASH() view returns(bytes32)
func (_MCPayment *MCPaymentCallerSession) ERC20PAYMENTDATATYPEHASH() ([32]byte, error) {
	return _MCPayment.Contract.ERC20PAYMENTDATATYPEHASH(&_MCPayment.CallOpts)
}

// PAYMENTDATATYPEHASH is a free data retrieval call binding the contract method 0xf0dd6899.
//
// Solidity: function PAYMENT_DATA_TYPE_HASH() view returns(bytes32)
func (_MCPayment *MCPaymentCaller) PAYMENTDATATYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "PAYMENT_DATA_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAYMENTDATATYPEHASH is a free data retrieval call binding the contract method 0xf0dd6899.
//
// Solidity: function PAYMENT_DATA_TYPE_HASH() view returns(bytes32)
func (_MCPayment *MCPaymentSession) PAYMENTDATATYPEHASH() ([32]byte, error) {
	return _MCPayment.Contract.PAYMENTDATATYPEHASH(&_MCPayment.CallOpts)
}

// PAYMENTDATATYPEHASH is a free data retrieval call binding the contract method 0xf0dd6899.
//
// Solidity: function PAYMENT_DATA_TYPE_HASH() view returns(bytes32)
func (_MCPayment *MCPaymentCallerSession) PAYMENTDATATYPEHASH() ([32]byte, error) {
	return _MCPayment.Contract.PAYMENTDATATYPEHASH(&_MCPayment.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_MCPayment *MCPaymentCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_MCPayment *MCPaymentSession) VERSION() (string, error) {
	return _MCPayment.Contract.VERSION(&_MCPayment.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_MCPayment *MCPaymentCallerSession) VERSION() (string, error) {
	return _MCPayment.Contract.VERSION(&_MCPayment.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MCPayment *MCPaymentCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MCPayment *MCPaymentSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MCPayment.Contract.Eip712Domain(&_MCPayment.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MCPayment *MCPaymentCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MCPayment.Contract.Eip712Domain(&_MCPayment.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address issuer) view returns(uint256)
func (_MCPayment *MCPaymentCaller) GetBalance(opts *bind.CallOpts, issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "getBalance", issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address issuer) view returns(uint256)
func (_MCPayment *MCPaymentSession) GetBalance(issuer common.Address) (*big.Int, error) {
	return _MCPayment.Contract.GetBalance(&_MCPayment.CallOpts, issuer)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address issuer) view returns(uint256)
func (_MCPayment *MCPaymentCallerSession) GetBalance(issuer common.Address) (*big.Int, error) {
	return _MCPayment.Contract.GetBalance(&_MCPayment.CallOpts, issuer)
}

// GetOwnerBalance is a free data retrieval call binding the contract method 0x590791f2.
//
// Solidity: function getOwnerBalance() view returns(uint256)
func (_MCPayment *MCPaymentCaller) GetOwnerBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "getOwnerBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnerBalance is a free data retrieval call binding the contract method 0x590791f2.
//
// Solidity: function getOwnerBalance() view returns(uint256)
func (_MCPayment *MCPaymentSession) GetOwnerBalance() (*big.Int, error) {
	return _MCPayment.Contract.GetOwnerBalance(&_MCPayment.CallOpts)
}

// GetOwnerBalance is a free data retrieval call binding the contract method 0x590791f2.
//
// Solidity: function getOwnerBalance() view returns(uint256)
func (_MCPayment *MCPaymentCallerSession) GetOwnerBalance() (*big.Int, error) {
	return _MCPayment.Contract.GetOwnerBalance(&_MCPayment.CallOpts)
}

// GetOwnerPercentage is a free data retrieval call binding the contract method 0x309a042c.
//
// Solidity: function getOwnerPercentage() view returns(uint8)
func (_MCPayment *MCPaymentCaller) GetOwnerPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "getOwnerPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOwnerPercentage is a free data retrieval call binding the contract method 0x309a042c.
//
// Solidity: function getOwnerPercentage() view returns(uint8)
func (_MCPayment *MCPaymentSession) GetOwnerPercentage() (uint8, error) {
	return _MCPayment.Contract.GetOwnerPercentage(&_MCPayment.CallOpts)
}

// GetOwnerPercentage is a free data retrieval call binding the contract method 0x309a042c.
//
// Solidity: function getOwnerPercentage() view returns(uint8)
func (_MCPayment *MCPaymentCallerSession) GetOwnerPercentage() (uint8, error) {
	return _MCPayment.Contract.GetOwnerPercentage(&_MCPayment.CallOpts)
}

// IsPaymentDone is a free data retrieval call binding the contract method 0x9d9c12b7.
//
// Solidity: function isPaymentDone(address issuer, uint256 nonce) view returns(bool)
func (_MCPayment *MCPaymentCaller) IsPaymentDone(opts *bind.CallOpts, issuer common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "isPaymentDone", issuer, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaymentDone is a free data retrieval call binding the contract method 0x9d9c12b7.
//
// Solidity: function isPaymentDone(address issuer, uint256 nonce) view returns(bool)
func (_MCPayment *MCPaymentSession) IsPaymentDone(issuer common.Address, nonce *big.Int) (bool, error) {
	return _MCPayment.Contract.IsPaymentDone(&_MCPayment.CallOpts, issuer, nonce)
}

// IsPaymentDone is a free data retrieval call binding the contract method 0x9d9c12b7.
//
// Solidity: function isPaymentDone(address issuer, uint256 nonce) view returns(bool)
func (_MCPayment *MCPaymentCallerSession) IsPaymentDone(issuer common.Address, nonce *big.Int) (bool, error) {
	return _MCPayment.Contract.IsPaymentDone(&_MCPayment.CallOpts, issuer, nonce)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCPayment *MCPaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCPayment *MCPaymentSession) Owner() (common.Address, error) {
	return _MCPayment.Contract.Owner(&_MCPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCPayment *MCPaymentCallerSession) Owner() (common.Address, error) {
	return _MCPayment.Contract.Owner(&_MCPayment.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MCPayment *MCPaymentCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MCPayment *MCPaymentSession) PendingOwner() (common.Address, error) {
	return _MCPayment.Contract.PendingOwner(&_MCPayment.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MCPayment *MCPaymentCallerSession) PendingOwner() (common.Address, error) {
	return _MCPayment.Contract.PendingOwner(&_MCPayment.CallOpts)
}

// RecoverIden3PaymentRailsERC20RequestV1Signature is a free data retrieval call binding the contract method 0xc52f5c83.
//
// Solidity: function recoverIden3PaymentRailsERC20RequestV1Signature((address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) view returns(address)
func (_MCPayment *MCPaymentCaller) RecoverIden3PaymentRailsERC20RequestV1Signature(opts *bind.CallOpts, paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "recoverIden3PaymentRailsERC20RequestV1Signature", paymentData, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverIden3PaymentRailsERC20RequestV1Signature is a free data retrieval call binding the contract method 0xc52f5c83.
//
// Solidity: function recoverIden3PaymentRailsERC20RequestV1Signature((address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) view returns(address)
func (_MCPayment *MCPaymentSession) RecoverIden3PaymentRailsERC20RequestV1Signature(paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (common.Address, error) {
	return _MCPayment.Contract.RecoverIden3PaymentRailsERC20RequestV1Signature(&_MCPayment.CallOpts, paymentData, signature)
}

// RecoverIden3PaymentRailsERC20RequestV1Signature is a free data retrieval call binding the contract method 0xc52f5c83.
//
// Solidity: function recoverIden3PaymentRailsERC20RequestV1Signature((address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) view returns(address)
func (_MCPayment *MCPaymentCallerSession) RecoverIden3PaymentRailsERC20RequestV1Signature(paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (common.Address, error) {
	return _MCPayment.Contract.RecoverIden3PaymentRailsERC20RequestV1Signature(&_MCPayment.CallOpts, paymentData, signature)
}

// RecoverIden3PaymentRailsRequestV1Signature is a free data retrieval call binding the contract method 0xc5fd08d1.
//
// Solidity: function recoverIden3PaymentRailsRequestV1Signature((address,uint256,uint256,uint256,bytes) paymentData, bytes signature) view returns(address)
func (_MCPayment *MCPaymentCaller) RecoverIden3PaymentRailsRequestV1Signature(opts *bind.CallOpts, paymentData MCPaymentIden3PaymentRailsRequestV1, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _MCPayment.contract.Call(opts, &out, "recoverIden3PaymentRailsRequestV1Signature", paymentData, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverIden3PaymentRailsRequestV1Signature is a free data retrieval call binding the contract method 0xc5fd08d1.
//
// Solidity: function recoverIden3PaymentRailsRequestV1Signature((address,uint256,uint256,uint256,bytes) paymentData, bytes signature) view returns(address)
func (_MCPayment *MCPaymentSession) RecoverIden3PaymentRailsRequestV1Signature(paymentData MCPaymentIden3PaymentRailsRequestV1, signature []byte) (common.Address, error) {
	return _MCPayment.Contract.RecoverIden3PaymentRailsRequestV1Signature(&_MCPayment.CallOpts, paymentData, signature)
}

// RecoverIden3PaymentRailsRequestV1Signature is a free data retrieval call binding the contract method 0xc5fd08d1.
//
// Solidity: function recoverIden3PaymentRailsRequestV1Signature((address,uint256,uint256,uint256,bytes) paymentData, bytes signature) view returns(address)
func (_MCPayment *MCPaymentCallerSession) RecoverIden3PaymentRailsRequestV1Signature(paymentData MCPaymentIden3PaymentRailsRequestV1, signature []byte) (common.Address, error) {
	return _MCPayment.Contract.RecoverIden3PaymentRailsRequestV1Signature(&_MCPayment.CallOpts, paymentData, signature)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MCPayment *MCPaymentTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MCPayment *MCPaymentSession) AcceptOwnership() (*types.Transaction, error) {
	return _MCPayment.Contract.AcceptOwnership(&_MCPayment.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MCPayment *MCPaymentTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MCPayment.Contract.AcceptOwnership(&_MCPayment.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 ownerPercentage) returns()
func (_MCPayment *MCPaymentTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, ownerPercentage uint8) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "initialize", owner, ownerPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 ownerPercentage) returns()
func (_MCPayment *MCPaymentSession) Initialize(owner common.Address, ownerPercentage uint8) (*types.Transaction, error) {
	return _MCPayment.Contract.Initialize(&_MCPayment.TransactOpts, owner, ownerPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 ownerPercentage) returns()
func (_MCPayment *MCPaymentTransactorSession) Initialize(owner common.Address, ownerPercentage uint8) (*types.Transaction, error) {
	return _MCPayment.Contract.Initialize(&_MCPayment.TransactOpts, owner, ownerPercentage)
}

// IssuerWithdraw is a paid mutator transaction binding the contract method 0xcc9cd961.
//
// Solidity: function issuerWithdraw() returns()
func (_MCPayment *MCPaymentTransactor) IssuerWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "issuerWithdraw")
}

// IssuerWithdraw is a paid mutator transaction binding the contract method 0xcc9cd961.
//
// Solidity: function issuerWithdraw() returns()
func (_MCPayment *MCPaymentSession) IssuerWithdraw() (*types.Transaction, error) {
	return _MCPayment.Contract.IssuerWithdraw(&_MCPayment.TransactOpts)
}

// IssuerWithdraw is a paid mutator transaction binding the contract method 0xcc9cd961.
//
// Solidity: function issuerWithdraw() returns()
func (_MCPayment *MCPaymentTransactorSession) IssuerWithdraw() (*types.Transaction, error) {
	return _MCPayment.Contract.IssuerWithdraw(&_MCPayment.TransactOpts)
}

// OwnerWithdraw is a paid mutator transaction binding the contract method 0x4311de8f.
//
// Solidity: function ownerWithdraw() returns()
func (_MCPayment *MCPaymentTransactor) OwnerWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "ownerWithdraw")
}

// OwnerWithdraw is a paid mutator transaction binding the contract method 0x4311de8f.
//
// Solidity: function ownerWithdraw() returns()
func (_MCPayment *MCPaymentSession) OwnerWithdraw() (*types.Transaction, error) {
	return _MCPayment.Contract.OwnerWithdraw(&_MCPayment.TransactOpts)
}

// OwnerWithdraw is a paid mutator transaction binding the contract method 0x4311de8f.
//
// Solidity: function ownerWithdraw() returns()
func (_MCPayment *MCPaymentTransactorSession) OwnerWithdraw() (*types.Transaction, error) {
	return _MCPayment.Contract.OwnerWithdraw(&_MCPayment.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xaa021669.
//
// Solidity: function pay((address,uint256,uint256,uint256,bytes) paymentData, bytes signature) payable returns()
func (_MCPayment *MCPaymentTransactor) Pay(opts *bind.TransactOpts, paymentData MCPaymentIden3PaymentRailsRequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "pay", paymentData, signature)
}

// Pay is a paid mutator transaction binding the contract method 0xaa021669.
//
// Solidity: function pay((address,uint256,uint256,uint256,bytes) paymentData, bytes signature) payable returns()
func (_MCPayment *MCPaymentSession) Pay(paymentData MCPaymentIden3PaymentRailsRequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.Contract.Pay(&_MCPayment.TransactOpts, paymentData, signature)
}

// Pay is a paid mutator transaction binding the contract method 0xaa021669.
//
// Solidity: function pay((address,uint256,uint256,uint256,bytes) paymentData, bytes signature) payable returns()
func (_MCPayment *MCPaymentTransactorSession) Pay(paymentData MCPaymentIden3PaymentRailsRequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.Contract.Pay(&_MCPayment.TransactOpts, paymentData, signature)
}

// PayERC20 is a paid mutator transaction binding the contract method 0x57615a3a.
//
// Solidity: function payERC20((address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) returns()
func (_MCPayment *MCPaymentTransactor) PayERC20(opts *bind.TransactOpts, paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "payERC20", paymentData, signature)
}

// PayERC20 is a paid mutator transaction binding the contract method 0x57615a3a.
//
// Solidity: function payERC20((address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) returns()
func (_MCPayment *MCPaymentSession) PayERC20(paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.Contract.PayERC20(&_MCPayment.TransactOpts, paymentData, signature)
}

// PayERC20 is a paid mutator transaction binding the contract method 0x57615a3a.
//
// Solidity: function payERC20((address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) returns()
func (_MCPayment *MCPaymentTransactorSession) PayERC20(paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.Contract.PayERC20(&_MCPayment.TransactOpts, paymentData, signature)
}

// PayERC20Permit is a paid mutator transaction binding the contract method 0x3dbbdbe5.
//
// Solidity: function payERC20Permit(bytes permitSignature, (address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) returns()
func (_MCPayment *MCPaymentTransactor) PayERC20Permit(opts *bind.TransactOpts, permitSignature []byte, paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "payERC20Permit", permitSignature, paymentData, signature)
}

// PayERC20Permit is a paid mutator transaction binding the contract method 0x3dbbdbe5.
//
// Solidity: function payERC20Permit(bytes permitSignature, (address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) returns()
func (_MCPayment *MCPaymentSession) PayERC20Permit(permitSignature []byte, paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.Contract.PayERC20Permit(&_MCPayment.TransactOpts, permitSignature, paymentData, signature)
}

// PayERC20Permit is a paid mutator transaction binding the contract method 0x3dbbdbe5.
//
// Solidity: function payERC20Permit(bytes permitSignature, (address,address,uint256,uint256,uint256,bytes) paymentData, bytes signature) returns()
func (_MCPayment *MCPaymentTransactorSession) PayERC20Permit(permitSignature []byte, paymentData MCPaymentIden3PaymentRailsERC20RequestV1, signature []byte) (*types.Transaction, error) {
	return _MCPayment.Contract.PayERC20Permit(&_MCPayment.TransactOpts, permitSignature, paymentData, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MCPayment *MCPaymentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MCPayment *MCPaymentSession) RenounceOwnership() (*types.Transaction, error) {
	return _MCPayment.Contract.RenounceOwnership(&_MCPayment.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MCPayment *MCPaymentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MCPayment.Contract.RenounceOwnership(&_MCPayment.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MCPayment *MCPaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MCPayment *MCPaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MCPayment.Contract.TransferOwnership(&_MCPayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MCPayment *MCPaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MCPayment.Contract.TransferOwnership(&_MCPayment.TransactOpts, newOwner)
}

// UpdateOwnerPercentage is a paid mutator transaction binding the contract method 0x0cea58aa.
//
// Solidity: function updateOwnerPercentage(uint8 ownerPercentage) returns()
func (_MCPayment *MCPaymentTransactor) UpdateOwnerPercentage(opts *bind.TransactOpts, ownerPercentage uint8) (*types.Transaction, error) {
	return _MCPayment.contract.Transact(opts, "updateOwnerPercentage", ownerPercentage)
}

// UpdateOwnerPercentage is a paid mutator transaction binding the contract method 0x0cea58aa.
//
// Solidity: function updateOwnerPercentage(uint8 ownerPercentage) returns()
func (_MCPayment *MCPaymentSession) UpdateOwnerPercentage(ownerPercentage uint8) (*types.Transaction, error) {
	return _MCPayment.Contract.UpdateOwnerPercentage(&_MCPayment.TransactOpts, ownerPercentage)
}

// UpdateOwnerPercentage is a paid mutator transaction binding the contract method 0x0cea58aa.
//
// Solidity: function updateOwnerPercentage(uint8 ownerPercentage) returns()
func (_MCPayment *MCPaymentTransactorSession) UpdateOwnerPercentage(ownerPercentage uint8) (*types.Transaction, error) {
	return _MCPayment.Contract.UpdateOwnerPercentage(&_MCPayment.TransactOpts, ownerPercentage)
}

// MCPaymentEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the MCPayment contract.
type MCPaymentEIP712DomainChangedIterator struct {
	Event *MCPaymentEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *MCPaymentEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCPaymentEIP712DomainChanged)
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
		it.Event = new(MCPaymentEIP712DomainChanged)
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
func (it *MCPaymentEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCPaymentEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCPaymentEIP712DomainChanged represents a EIP712DomainChanged event raised by the MCPayment contract.
type MCPaymentEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MCPayment *MCPaymentFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MCPaymentEIP712DomainChangedIterator, error) {

	logs, sub, err := _MCPayment.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MCPaymentEIP712DomainChangedIterator{contract: _MCPayment.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MCPayment *MCPaymentFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MCPaymentEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _MCPayment.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCPaymentEIP712DomainChanged)
				if err := _MCPayment.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MCPayment *MCPaymentFilterer) ParseEIP712DomainChanged(log types.Log) (*MCPaymentEIP712DomainChanged, error) {
	event := new(MCPaymentEIP712DomainChanged)
	if err := _MCPayment.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCPaymentInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MCPayment contract.
type MCPaymentInitializedIterator struct {
	Event *MCPaymentInitialized // Event containing the contract specifics and raw log

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
func (it *MCPaymentInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCPaymentInitialized)
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
		it.Event = new(MCPaymentInitialized)
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
func (it *MCPaymentInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCPaymentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCPaymentInitialized represents a Initialized event raised by the MCPayment contract.
type MCPaymentInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MCPayment *MCPaymentFilterer) FilterInitialized(opts *bind.FilterOpts) (*MCPaymentInitializedIterator, error) {

	logs, sub, err := _MCPayment.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MCPaymentInitializedIterator{contract: _MCPayment.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MCPayment *MCPaymentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MCPaymentInitialized) (event.Subscription, error) {

	logs, sub, err := _MCPayment.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCPaymentInitialized)
				if err := _MCPayment.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MCPayment *MCPaymentFilterer) ParseInitialized(log types.Log) (*MCPaymentInitialized, error) {
	event := new(MCPaymentInitialized)
	if err := _MCPayment.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCPaymentOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MCPayment contract.
type MCPaymentOwnershipTransferStartedIterator struct {
	Event *MCPaymentOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *MCPaymentOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCPaymentOwnershipTransferStarted)
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
		it.Event = new(MCPaymentOwnershipTransferStarted)
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
func (it *MCPaymentOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCPaymentOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCPaymentOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MCPayment contract.
type MCPaymentOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MCPayment *MCPaymentFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MCPaymentOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MCPayment.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MCPaymentOwnershipTransferStartedIterator{contract: _MCPayment.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MCPayment *MCPaymentFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MCPaymentOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MCPayment.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCPaymentOwnershipTransferStarted)
				if err := _MCPayment.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_MCPayment *MCPaymentFilterer) ParseOwnershipTransferStarted(log types.Log) (*MCPaymentOwnershipTransferStarted, error) {
	event := new(MCPaymentOwnershipTransferStarted)
	if err := _MCPayment.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCPaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MCPayment contract.
type MCPaymentOwnershipTransferredIterator struct {
	Event *MCPaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MCPaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCPaymentOwnershipTransferred)
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
		it.Event = new(MCPaymentOwnershipTransferred)
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
func (it *MCPaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCPaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCPaymentOwnershipTransferred represents a OwnershipTransferred event raised by the MCPayment contract.
type MCPaymentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MCPayment *MCPaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MCPaymentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MCPayment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MCPaymentOwnershipTransferredIterator{contract: _MCPayment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MCPayment *MCPaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MCPaymentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MCPayment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCPaymentOwnershipTransferred)
				if err := _MCPayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MCPayment *MCPaymentFilterer) ParseOwnershipTransferred(log types.Log) (*MCPaymentOwnershipTransferred, error) {
	event := new(MCPaymentOwnershipTransferred)
	if err := _MCPayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCPaymentPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the MCPayment contract.
type MCPaymentPaymentIterator struct {
	Event *MCPaymentPayment // Event containing the contract specifics and raw log

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
func (it *MCPaymentPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCPaymentPayment)
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
		it.Event = new(MCPaymentPayment)
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
func (it *MCPaymentPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCPaymentPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCPaymentPayment represents a Payment event raised by the MCPayment contract.
type MCPaymentPayment struct {
	Recipient common.Address
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0xd4f43975feb89f48dd30cabbb32011045be187d1e11c8ea9faa43efc35282519.
//
// Solidity: event Payment(address indexed recipient, uint256 indexed nonce)
func (_MCPayment *MCPaymentFilterer) FilterPayment(opts *bind.FilterOpts, recipient []common.Address, nonce []*big.Int) (*MCPaymentPaymentIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MCPayment.contract.FilterLogs(opts, "Payment", recipientRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MCPaymentPaymentIterator{contract: _MCPayment.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0xd4f43975feb89f48dd30cabbb32011045be187d1e11c8ea9faa43efc35282519.
//
// Solidity: event Payment(address indexed recipient, uint256 indexed nonce)
func (_MCPayment *MCPaymentFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *MCPaymentPayment, recipient []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MCPayment.contract.WatchLogs(opts, "Payment", recipientRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCPaymentPayment)
				if err := _MCPayment.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0xd4f43975feb89f48dd30cabbb32011045be187d1e11c8ea9faa43efc35282519.
//
// Solidity: event Payment(address indexed recipient, uint256 indexed nonce)
func (_MCPayment *MCPaymentFilterer) ParsePayment(log types.Log) (*MCPaymentPayment, error) {
	event := new(MCPaymentPayment)
	if err := _MCPayment.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
