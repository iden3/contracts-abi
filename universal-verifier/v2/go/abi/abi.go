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

// IRequestValidatorResponseField is an auto generated low-level Go binding around an user-defined struct.
type IRequestValidatorResponseField struct {
	Name     string
	Value    *big.Int
	RawValue []byte
}

// IVerifierAuthMethod is an auto generated low-level Go binding around an user-defined struct.
type IVerifierAuthMethod struct {
	AuthMethod string
	Validator  common.Address
	Params     []byte
}

// IVerifierAuthResponse is an auto generated low-level Go binding around an user-defined struct.
type IVerifierAuthResponse struct {
	AuthMethod string
	Proof      []byte
}

// IVerifierMultiRequest is an auto generated low-level Go binding around an user-defined struct.
type IVerifierMultiRequest struct {
	MultiRequestId *big.Int
	RequestIds     []*big.Int
	GroupIds       []*big.Int
	Metadata       []byte
}

// IVerifierRequest is an auto generated low-level Go binding around an user-defined struct.
type IVerifierRequest struct {
	RequestId *big.Int
	Metadata  string
	Validator common.Address
	Params    []byte
	Owner     common.Address
}

// IVerifierRequestInfo is an auto generated low-level Go binding around an user-defined struct.
type IVerifierRequestInfo struct {
	RequestId *big.Int
	Metadata  string
	Validator common.Address
	Params    []byte
	Creator   common.Address
}

// IVerifierRequestProofStatus is an auto generated low-level Go binding around an user-defined struct.
type IVerifierRequestProofStatus struct {
	RequestId        *big.Int
	IsVerified       bool
	ValidatorVersion string
	Timestamp        *big.Int
}

// IVerifierResponse is an auto generated low-level Go binding around an user-defined struct.
type IVerifierResponse struct {
	RequestId *big.Int
	Proof     []byte
	Metadata  []byte
}

// VerifierAuthMethodData is an auto generated low-level Go binding around an user-defined struct.
type VerifierAuthMethodData struct {
	Validator common.Address
	Params    []byte
	IsActive  bool
}

// UniversalVerifierMetaData contains all meta data concerning the UniversalVerifier contract.
var UniversalVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"AuthMethodAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"AuthMethodIsNotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"AuthMethodNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengeIsInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ChecksumLengthRequired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"GroupIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"GroupIdNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GroupIdNotValid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"GroupMustHaveAtLeastTwoRequests\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"IdBytesLengthRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"InvalidRequestOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkIDNotTheSameForGroupedRequests\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MetadataNotSupportedYet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"MissingUserIDInGroupOfRequests\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"MissingUserIDInRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"}],\"name\":\"MultiRequestIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"}],\"name\":\"MultiRequestIdNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedMultiRequestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"}],\"name\":\"MultiRequestIdNotValid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotAnOwnerOrRequestOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nullifierSessionID\",\"type\":\"uint256\"}],\"name\":\"NullifierSessionIDAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestIdNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRequestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestIdNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestIdTypeNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestIdUsesReservedBytes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestIsDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestShouldNotHaveAGroup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"responseFieldName\",\"type\":\"string\"}],\"name\":\"ResponseFieldAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userIDFromAuth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userIDFromResponse\",\"type\":\"uint256\"}],\"name\":\"UserIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotAuthenticated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorIsNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorNotSupportInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestVerifierID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedVerifierID\",\"type\":\"uint256\"}],\"name\":\"VerifierIDIsNotValid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"AuthMethodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthResponseSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"groupIds\",\"type\":\"uint256[]\"}],\"name\":\"MultiRequestSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"RequestSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"RequestUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ResponseSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidatorToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"areMultiRequestProofsVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"authMethodExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"disableAuthMethod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"disableRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"enableAuthMethod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"enableRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"}],\"name\":\"getAuthMethod\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIAuthValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structVerifier.AuthMethodData\",\"name\":\"authMethodData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupID\",\"type\":\"uint256\"}],\"name\":\"getGroupedRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"internalType\":\"structIVerifier.RequestInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGroupsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"}],\"name\":\"getMultiRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"groupIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structIVerifier.MultiRequest\",\"name\":\"multiRequest\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getMultiRequestProofsStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"validatorVersion\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIVerifier.RequestProofStatus[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"internalType\":\"structIVerifier.RequestInfo\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestProofStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"validatorVersion\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIVerifier.RequestProofStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"responseFieldName\",\"type\":\"string\"}],\"name\":\"getResponseFieldValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getResponseFields\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawValue\",\"type\":\"bytes\"}],\"internalType\":\"structIRequestValidator.ResponseField[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifierID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"groupIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIState\",\"name\":\"state\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"isRequestEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"isRequestProofVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isWhitelistedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"}],\"name\":\"multiRequestIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidatorFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"requestIdExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"},{\"internalType\":\"contractIAuthValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structIVerifier.AuthMethod\",\"name\":\"authMethod\",\"type\":\"tuple\"}],\"name\":\"setAuthMethod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"multiRequestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"groupIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structIVerifier.MultiRequest\",\"name\":\"multiRequest\",\"type\":\"tuple\"}],\"name\":\"setMultiRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requestOwner\",\"type\":\"address\"}],\"name\":\"setRequestOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIVerifier.Request[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"name\":\"setRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIState\",\"name\":\"state\",\"type\":\"address\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"verifierID\",\"type\":\"uint256\"}],\"name\":\"setVerifierID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"authMethod\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structIVerifier.AuthResponse\",\"name\":\"authResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structIVerifier.Response[]\",\"name\":\"responses\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"crossChainProofs\",\"type\":\"bytes\"}],\"name\":\"submitResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"contractIRequestValidator\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIVerifier.Request\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"updateRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// AreMultiRequestProofsVerified is a free data retrieval call binding the contract method 0xd4d9f766.
//
// Solidity: function areMultiRequestProofsVerified(uint256 multiRequestId, address userAddress) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) AreMultiRequestProofsVerified(opts *bind.CallOpts, multiRequestId *big.Int, userAddress common.Address) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "areMultiRequestProofsVerified", multiRequestId, userAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreMultiRequestProofsVerified is a free data retrieval call binding the contract method 0xd4d9f766.
//
// Solidity: function areMultiRequestProofsVerified(uint256 multiRequestId, address userAddress) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) AreMultiRequestProofsVerified(multiRequestId *big.Int, userAddress common.Address) (bool, error) {
	return _UniversalVerifier.Contract.AreMultiRequestProofsVerified(&_UniversalVerifier.CallOpts, multiRequestId, userAddress)
}

// AreMultiRequestProofsVerified is a free data retrieval call binding the contract method 0xd4d9f766.
//
// Solidity: function areMultiRequestProofsVerified(uint256 multiRequestId, address userAddress) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) AreMultiRequestProofsVerified(multiRequestId *big.Int, userAddress common.Address) (bool, error) {
	return _UniversalVerifier.Contract.AreMultiRequestProofsVerified(&_UniversalVerifier.CallOpts, multiRequestId, userAddress)
}

// AuthMethodExists is a free data retrieval call binding the contract method 0xbbd8b7a1.
//
// Solidity: function authMethodExists(string authMethod) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) AuthMethodExists(opts *bind.CallOpts, authMethod string) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "authMethodExists", authMethod)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthMethodExists is a free data retrieval call binding the contract method 0xbbd8b7a1.
//
// Solidity: function authMethodExists(string authMethod) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) AuthMethodExists(authMethod string) (bool, error) {
	return _UniversalVerifier.Contract.AuthMethodExists(&_UniversalVerifier.CallOpts, authMethod)
}

// AuthMethodExists is a free data retrieval call binding the contract method 0xbbd8b7a1.
//
// Solidity: function authMethodExists(string authMethod) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) AuthMethodExists(authMethod string) (bool, error) {
	return _UniversalVerifier.Contract.AuthMethodExists(&_UniversalVerifier.CallOpts, authMethod)
}

// GetAuthMethod is a free data retrieval call binding the contract method 0x751adcd2.
//
// Solidity: function getAuthMethod(string authMethod) view returns((address,bytes,bool) authMethodData)
func (_UniversalVerifier *UniversalVerifierCaller) GetAuthMethod(opts *bind.CallOpts, authMethod string) (VerifierAuthMethodData, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getAuthMethod", authMethod)

	if err != nil {
		return *new(VerifierAuthMethodData), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierAuthMethodData)).(*VerifierAuthMethodData)

	return out0, err

}

// GetAuthMethod is a free data retrieval call binding the contract method 0x751adcd2.
//
// Solidity: function getAuthMethod(string authMethod) view returns((address,bytes,bool) authMethodData)
func (_UniversalVerifier *UniversalVerifierSession) GetAuthMethod(authMethod string) (VerifierAuthMethodData, error) {
	return _UniversalVerifier.Contract.GetAuthMethod(&_UniversalVerifier.CallOpts, authMethod)
}

// GetAuthMethod is a free data retrieval call binding the contract method 0x751adcd2.
//
// Solidity: function getAuthMethod(string authMethod) view returns((address,bytes,bool) authMethodData)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetAuthMethod(authMethod string) (VerifierAuthMethodData, error) {
	return _UniversalVerifier.Contract.GetAuthMethod(&_UniversalVerifier.CallOpts, authMethod)
}

// GetGroupedRequests is a free data retrieval call binding the contract method 0x30837e06.
//
// Solidity: function getGroupedRequests(uint256 groupID) view returns((uint256,string,address,bytes,address)[])
func (_UniversalVerifier *UniversalVerifierCaller) GetGroupedRequests(opts *bind.CallOpts, groupID *big.Int) ([]IVerifierRequestInfo, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getGroupedRequests", groupID)

	if err != nil {
		return *new([]IVerifierRequestInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IVerifierRequestInfo)).(*[]IVerifierRequestInfo)

	return out0, err

}

// GetGroupedRequests is a free data retrieval call binding the contract method 0x30837e06.
//
// Solidity: function getGroupedRequests(uint256 groupID) view returns((uint256,string,address,bytes,address)[])
func (_UniversalVerifier *UniversalVerifierSession) GetGroupedRequests(groupID *big.Int) ([]IVerifierRequestInfo, error) {
	return _UniversalVerifier.Contract.GetGroupedRequests(&_UniversalVerifier.CallOpts, groupID)
}

// GetGroupedRequests is a free data retrieval call binding the contract method 0x30837e06.
//
// Solidity: function getGroupedRequests(uint256 groupID) view returns((uint256,string,address,bytes,address)[])
func (_UniversalVerifier *UniversalVerifierCallerSession) GetGroupedRequests(groupID *big.Int) ([]IVerifierRequestInfo, error) {
	return _UniversalVerifier.Contract.GetGroupedRequests(&_UniversalVerifier.CallOpts, groupID)
}

// GetGroupsCount is a free data retrieval call binding the contract method 0xb84afbb4.
//
// Solidity: function getGroupsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetGroupsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getGroupsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGroupsCount is a free data retrieval call binding the contract method 0xb84afbb4.
//
// Solidity: function getGroupsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetGroupsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetGroupsCount(&_UniversalVerifier.CallOpts)
}

// GetGroupsCount is a free data retrieval call binding the contract method 0xb84afbb4.
//
// Solidity: function getGroupsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetGroupsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetGroupsCount(&_UniversalVerifier.CallOpts)
}

// GetMultiRequest is a free data retrieval call binding the contract method 0xf2e27c68.
//
// Solidity: function getMultiRequest(uint256 multiRequestId) view returns((uint256,uint256[],uint256[],bytes) multiRequest)
func (_UniversalVerifier *UniversalVerifierCaller) GetMultiRequest(opts *bind.CallOpts, multiRequestId *big.Int) (IVerifierMultiRequest, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getMultiRequest", multiRequestId)

	if err != nil {
		return *new(IVerifierMultiRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IVerifierMultiRequest)).(*IVerifierMultiRequest)

	return out0, err

}

// GetMultiRequest is a free data retrieval call binding the contract method 0xf2e27c68.
//
// Solidity: function getMultiRequest(uint256 multiRequestId) view returns((uint256,uint256[],uint256[],bytes) multiRequest)
func (_UniversalVerifier *UniversalVerifierSession) GetMultiRequest(multiRequestId *big.Int) (IVerifierMultiRequest, error) {
	return _UniversalVerifier.Contract.GetMultiRequest(&_UniversalVerifier.CallOpts, multiRequestId)
}

// GetMultiRequest is a free data retrieval call binding the contract method 0xf2e27c68.
//
// Solidity: function getMultiRequest(uint256 multiRequestId) view returns((uint256,uint256[],uint256[],bytes) multiRequest)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetMultiRequest(multiRequestId *big.Int) (IVerifierMultiRequest, error) {
	return _UniversalVerifier.Contract.GetMultiRequest(&_UniversalVerifier.CallOpts, multiRequestId)
}

// GetMultiRequestProofsStatus is a free data retrieval call binding the contract method 0xf1fd9964.
//
// Solidity: function getMultiRequestProofsStatus(uint256 multiRequestId, address userAddress) view returns((uint256,bool,string,uint256)[])
func (_UniversalVerifier *UniversalVerifierCaller) GetMultiRequestProofsStatus(opts *bind.CallOpts, multiRequestId *big.Int, userAddress common.Address) ([]IVerifierRequestProofStatus, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getMultiRequestProofsStatus", multiRequestId, userAddress)

	if err != nil {
		return *new([]IVerifierRequestProofStatus), err
	}

	out0 := *abi.ConvertType(out[0], new([]IVerifierRequestProofStatus)).(*[]IVerifierRequestProofStatus)

	return out0, err

}

// GetMultiRequestProofsStatus is a free data retrieval call binding the contract method 0xf1fd9964.
//
// Solidity: function getMultiRequestProofsStatus(uint256 multiRequestId, address userAddress) view returns((uint256,bool,string,uint256)[])
func (_UniversalVerifier *UniversalVerifierSession) GetMultiRequestProofsStatus(multiRequestId *big.Int, userAddress common.Address) ([]IVerifierRequestProofStatus, error) {
	return _UniversalVerifier.Contract.GetMultiRequestProofsStatus(&_UniversalVerifier.CallOpts, multiRequestId, userAddress)
}

// GetMultiRequestProofsStatus is a free data retrieval call binding the contract method 0xf1fd9964.
//
// Solidity: function getMultiRequestProofsStatus(uint256 multiRequestId, address userAddress) view returns((uint256,bool,string,uint256)[])
func (_UniversalVerifier *UniversalVerifierCallerSession) GetMultiRequestProofsStatus(multiRequestId *big.Int, userAddress common.Address) ([]IVerifierRequestProofStatus, error) {
	return _UniversalVerifier.Contract.GetMultiRequestProofsStatus(&_UniversalVerifier.CallOpts, multiRequestId, userAddress)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((uint256,string,address,bytes,address) request)
func (_UniversalVerifier *UniversalVerifierCaller) GetRequest(opts *bind.CallOpts, requestId *big.Int) (IVerifierRequestInfo, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRequest", requestId)

	if err != nil {
		return *new(IVerifierRequestInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IVerifierRequestInfo)).(*IVerifierRequestInfo)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((uint256,string,address,bytes,address) request)
func (_UniversalVerifier *UniversalVerifierSession) GetRequest(requestId *big.Int) (IVerifierRequestInfo, error) {
	return _UniversalVerifier.Contract.GetRequest(&_UniversalVerifier.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((uint256,string,address,bytes,address) request)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRequest(requestId *big.Int) (IVerifierRequestInfo, error) {
	return _UniversalVerifier.Contract.GetRequest(&_UniversalVerifier.CallOpts, requestId)
}

// GetRequestOwner is a free data retrieval call binding the contract method 0x8611cb67.
//
// Solidity: function getRequestOwner(uint256 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierCaller) GetRequestOwner(opts *bind.CallOpts, requestId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRequestOwner", requestId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRequestOwner is a free data retrieval call binding the contract method 0x8611cb67.
//
// Solidity: function getRequestOwner(uint256 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierSession) GetRequestOwner(requestId *big.Int) (common.Address, error) {
	return _UniversalVerifier.Contract.GetRequestOwner(&_UniversalVerifier.CallOpts, requestId)
}

// GetRequestOwner is a free data retrieval call binding the contract method 0x8611cb67.
//
// Solidity: function getRequestOwner(uint256 requestId) view returns(address)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRequestOwner(requestId *big.Int) (common.Address, error) {
	return _UniversalVerifier.Contract.GetRequestOwner(&_UniversalVerifier.CallOpts, requestId)
}

// GetRequestProofStatus is a free data retrieval call binding the contract method 0xad9d21e7.
//
// Solidity: function getRequestProofStatus(address sender, uint256 requestId) view returns((uint256,bool,string,uint256))
func (_UniversalVerifier *UniversalVerifierCaller) GetRequestProofStatus(opts *bind.CallOpts, sender common.Address, requestId *big.Int) (IVerifierRequestProofStatus, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRequestProofStatus", sender, requestId)

	if err != nil {
		return *new(IVerifierRequestProofStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IVerifierRequestProofStatus)).(*IVerifierRequestProofStatus)

	return out0, err

}

// GetRequestProofStatus is a free data retrieval call binding the contract method 0xad9d21e7.
//
// Solidity: function getRequestProofStatus(address sender, uint256 requestId) view returns((uint256,bool,string,uint256))
func (_UniversalVerifier *UniversalVerifierSession) GetRequestProofStatus(sender common.Address, requestId *big.Int) (IVerifierRequestProofStatus, error) {
	return _UniversalVerifier.Contract.GetRequestProofStatus(&_UniversalVerifier.CallOpts, sender, requestId)
}

// GetRequestProofStatus is a free data retrieval call binding the contract method 0xad9d21e7.
//
// Solidity: function getRequestProofStatus(address sender, uint256 requestId) view returns((uint256,bool,string,uint256))
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRequestProofStatus(sender common.Address, requestId *big.Int) (IVerifierRequestProofStatus, error) {
	return _UniversalVerifier.Contract.GetRequestProofStatus(&_UniversalVerifier.CallOpts, sender, requestId)
}

// GetRequestsCount is a free data retrieval call binding the contract method 0x3410452a.
//
// Solidity: function getRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetRequestsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getRequestsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestsCount is a free data retrieval call binding the contract method 0x3410452a.
//
// Solidity: function getRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetRequestsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetRequestsCount(&_UniversalVerifier.CallOpts)
}

// GetRequestsCount is a free data retrieval call binding the contract method 0x3410452a.
//
// Solidity: function getRequestsCount() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetRequestsCount() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetRequestsCount(&_UniversalVerifier.CallOpts)
}

// GetResponseFieldValue is a free data retrieval call binding the contract method 0xbfcbcc31.
//
// Solidity: function getResponseFieldValue(uint256 requestId, address sender, string responseFieldName) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetResponseFieldValue(opts *bind.CallOpts, requestId *big.Int, sender common.Address, responseFieldName string) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getResponseFieldValue", requestId, sender, responseFieldName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResponseFieldValue is a free data retrieval call binding the contract method 0xbfcbcc31.
//
// Solidity: function getResponseFieldValue(uint256 requestId, address sender, string responseFieldName) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetResponseFieldValue(requestId *big.Int, sender common.Address, responseFieldName string) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetResponseFieldValue(&_UniversalVerifier.CallOpts, requestId, sender, responseFieldName)
}

// GetResponseFieldValue is a free data retrieval call binding the contract method 0xbfcbcc31.
//
// Solidity: function getResponseFieldValue(uint256 requestId, address sender, string responseFieldName) view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetResponseFieldValue(requestId *big.Int, sender common.Address, responseFieldName string) (*big.Int, error) {
	return _UniversalVerifier.Contract.GetResponseFieldValue(&_UniversalVerifier.CallOpts, requestId, sender, responseFieldName)
}

// GetResponseFields is a free data retrieval call binding the contract method 0x829b3ae1.
//
// Solidity: function getResponseFields(uint256 requestId, address sender) view returns((string,uint256,bytes)[])
func (_UniversalVerifier *UniversalVerifierCaller) GetResponseFields(opts *bind.CallOpts, requestId *big.Int, sender common.Address) ([]IRequestValidatorResponseField, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getResponseFields", requestId, sender)

	if err != nil {
		return *new([]IRequestValidatorResponseField), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRequestValidatorResponseField)).(*[]IRequestValidatorResponseField)

	return out0, err

}

// GetResponseFields is a free data retrieval call binding the contract method 0x829b3ae1.
//
// Solidity: function getResponseFields(uint256 requestId, address sender) view returns((string,uint256,bytes)[])
func (_UniversalVerifier *UniversalVerifierSession) GetResponseFields(requestId *big.Int, sender common.Address) ([]IRequestValidatorResponseField, error) {
	return _UniversalVerifier.Contract.GetResponseFields(&_UniversalVerifier.CallOpts, requestId, sender)
}

// GetResponseFields is a free data retrieval call binding the contract method 0x829b3ae1.
//
// Solidity: function getResponseFields(uint256 requestId, address sender) view returns((string,uint256,bytes)[])
func (_UniversalVerifier *UniversalVerifierCallerSession) GetResponseFields(requestId *big.Int, sender common.Address) ([]IRequestValidatorResponseField, error) {
	return _UniversalVerifier.Contract.GetResponseFields(&_UniversalVerifier.CallOpts, requestId, sender)
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

// GetVerifierID is a free data retrieval call binding the contract method 0xb208eae8.
//
// Solidity: function getVerifierID() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCaller) GetVerifierID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "getVerifierID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVerifierID is a free data retrieval call binding the contract method 0xb208eae8.
//
// Solidity: function getVerifierID() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierSession) GetVerifierID() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetVerifierID(&_UniversalVerifier.CallOpts)
}

// GetVerifierID is a free data retrieval call binding the contract method 0xb208eae8.
//
// Solidity: function getVerifierID() view returns(uint256)
func (_UniversalVerifier *UniversalVerifierCallerSession) GetVerifierID() (*big.Int, error) {
	return _UniversalVerifier.Contract.GetVerifierID(&_UniversalVerifier.CallOpts)
}

// GroupIdExists is a free data retrieval call binding the contract method 0x4288c7ae.
//
// Solidity: function groupIdExists(uint256 groupId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) GroupIdExists(opts *bind.CallOpts, groupId *big.Int) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "groupIdExists", groupId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GroupIdExists is a free data retrieval call binding the contract method 0x4288c7ae.
//
// Solidity: function groupIdExists(uint256 groupId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) GroupIdExists(groupId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.GroupIdExists(&_UniversalVerifier.CallOpts, groupId)
}

// GroupIdExists is a free data retrieval call binding the contract method 0x4288c7ae.
//
// Solidity: function groupIdExists(uint256 groupId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) GroupIdExists(groupId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.GroupIdExists(&_UniversalVerifier.CallOpts, groupId)
}

// IsRequestEnabled is a free data retrieval call binding the contract method 0x0f12a241.
//
// Solidity: function isRequestEnabled(uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsRequestEnabled(opts *bind.CallOpts, requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isRequestEnabled", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRequestEnabled is a free data retrieval call binding the contract method 0x0f12a241.
//
// Solidity: function isRequestEnabled(uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsRequestEnabled(requestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.IsRequestEnabled(&_UniversalVerifier.CallOpts, requestId)
}

// IsRequestEnabled is a free data retrieval call binding the contract method 0x0f12a241.
//
// Solidity: function isRequestEnabled(uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsRequestEnabled(requestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.IsRequestEnabled(&_UniversalVerifier.CallOpts, requestId)
}

// IsRequestProofVerified is a free data retrieval call binding the contract method 0xe8555603.
//
// Solidity: function isRequestProofVerified(address sender, uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) IsRequestProofVerified(opts *bind.CallOpts, sender common.Address, requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "isRequestProofVerified", sender, requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRequestProofVerified is a free data retrieval call binding the contract method 0xe8555603.
//
// Solidity: function isRequestProofVerified(address sender, uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) IsRequestProofVerified(sender common.Address, requestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.IsRequestProofVerified(&_UniversalVerifier.CallOpts, sender, requestId)
}

// IsRequestProofVerified is a free data retrieval call binding the contract method 0xe8555603.
//
// Solidity: function isRequestProofVerified(address sender, uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) IsRequestProofVerified(sender common.Address, requestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.IsRequestProofVerified(&_UniversalVerifier.CallOpts, sender, requestId)
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

// MultiRequestIdExists is a free data retrieval call binding the contract method 0xcea0607f.
//
// Solidity: function multiRequestIdExists(uint256 multiRequestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) MultiRequestIdExists(opts *bind.CallOpts, multiRequestId *big.Int) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "multiRequestIdExists", multiRequestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiRequestIdExists is a free data retrieval call binding the contract method 0xcea0607f.
//
// Solidity: function multiRequestIdExists(uint256 multiRequestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) MultiRequestIdExists(multiRequestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.MultiRequestIdExists(&_UniversalVerifier.CallOpts, multiRequestId)
}

// MultiRequestIdExists is a free data retrieval call binding the contract method 0xcea0607f.
//
// Solidity: function multiRequestIdExists(uint256 multiRequestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) MultiRequestIdExists(multiRequestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.MultiRequestIdExists(&_UniversalVerifier.CallOpts, multiRequestId)
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

// RequestIdExists is a free data retrieval call binding the contract method 0x7b8f6d37.
//
// Solidity: function requestIdExists(uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCaller) RequestIdExists(opts *bind.CallOpts, requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _UniversalVerifier.contract.Call(opts, &out, "requestIdExists", requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestIdExists is a free data retrieval call binding the contract method 0x7b8f6d37.
//
// Solidity: function requestIdExists(uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierSession) RequestIdExists(requestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.RequestIdExists(&_UniversalVerifier.CallOpts, requestId)
}

// RequestIdExists is a free data retrieval call binding the contract method 0x7b8f6d37.
//
// Solidity: function requestIdExists(uint256 requestId) view returns(bool)
func (_UniversalVerifier *UniversalVerifierCallerSession) RequestIdExists(requestId *big.Int) (bool, error) {
	return _UniversalVerifier.Contract.RequestIdExists(&_UniversalVerifier.CallOpts, requestId)
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

// DisableAuthMethod is a paid mutator transaction binding the contract method 0x7c7a9a5c.
//
// Solidity: function disableAuthMethod(string authMethod) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) DisableAuthMethod(opts *bind.TransactOpts, authMethod string) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "disableAuthMethod", authMethod)
}

// DisableAuthMethod is a paid mutator transaction binding the contract method 0x7c7a9a5c.
//
// Solidity: function disableAuthMethod(string authMethod) returns()
func (_UniversalVerifier *UniversalVerifierSession) DisableAuthMethod(authMethod string) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableAuthMethod(&_UniversalVerifier.TransactOpts, authMethod)
}

// DisableAuthMethod is a paid mutator transaction binding the contract method 0x7c7a9a5c.
//
// Solidity: function disableAuthMethod(string authMethod) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) DisableAuthMethod(authMethod string) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableAuthMethod(&_UniversalVerifier.TransactOpts, authMethod)
}

// DisableRequest is a paid mutator transaction binding the contract method 0xb33a998f.
//
// Solidity: function disableRequest(uint256 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) DisableRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "disableRequest", requestId)
}

// DisableRequest is a paid mutator transaction binding the contract method 0xb33a998f.
//
// Solidity: function disableRequest(uint256 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) DisableRequest(requestId *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// DisableRequest is a paid mutator transaction binding the contract method 0xb33a998f.
//
// Solidity: function disableRequest(uint256 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) DisableRequest(requestId *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.DisableRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// EnableAuthMethod is a paid mutator transaction binding the contract method 0x1cf5cf24.
//
// Solidity: function enableAuthMethod(string authMethod) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) EnableAuthMethod(opts *bind.TransactOpts, authMethod string) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "enableAuthMethod", authMethod)
}

// EnableAuthMethod is a paid mutator transaction binding the contract method 0x1cf5cf24.
//
// Solidity: function enableAuthMethod(string authMethod) returns()
func (_UniversalVerifier *UniversalVerifierSession) EnableAuthMethod(authMethod string) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableAuthMethod(&_UniversalVerifier.TransactOpts, authMethod)
}

// EnableAuthMethod is a paid mutator transaction binding the contract method 0x1cf5cf24.
//
// Solidity: function enableAuthMethod(string authMethod) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) EnableAuthMethod(authMethod string) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableAuthMethod(&_UniversalVerifier.TransactOpts, authMethod)
}

// EnableRequest is a paid mutator transaction binding the contract method 0x47861fe2.
//
// Solidity: function enableRequest(uint256 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) EnableRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "enableRequest", requestId)
}

// EnableRequest is a paid mutator transaction binding the contract method 0x47861fe2.
//
// Solidity: function enableRequest(uint256 requestId) returns()
func (_UniversalVerifier *UniversalVerifierSession) EnableRequest(requestId *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableRequest(&_UniversalVerifier.TransactOpts, requestId)
}

// EnableRequest is a paid mutator transaction binding the contract method 0x47861fe2.
//
// Solidity: function enableRequest(uint256 requestId) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) EnableRequest(requestId *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.EnableRequest(&_UniversalVerifier.TransactOpts, requestId)
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

// SetAuthMethod is a paid mutator transaction binding the contract method 0xe1470f9a.
//
// Solidity: function setAuthMethod((string,address,bytes) authMethod) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetAuthMethod(opts *bind.TransactOpts, authMethod IVerifierAuthMethod) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setAuthMethod", authMethod)
}

// SetAuthMethod is a paid mutator transaction binding the contract method 0xe1470f9a.
//
// Solidity: function setAuthMethod((string,address,bytes) authMethod) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetAuthMethod(authMethod IVerifierAuthMethod) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetAuthMethod(&_UniversalVerifier.TransactOpts, authMethod)
}

// SetAuthMethod is a paid mutator transaction binding the contract method 0xe1470f9a.
//
// Solidity: function setAuthMethod((string,address,bytes) authMethod) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetAuthMethod(authMethod IVerifierAuthMethod) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetAuthMethod(&_UniversalVerifier.TransactOpts, authMethod)
}

// SetMultiRequest is a paid mutator transaction binding the contract method 0x7bc499e5.
//
// Solidity: function setMultiRequest((uint256,uint256[],uint256[],bytes) multiRequest) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetMultiRequest(opts *bind.TransactOpts, multiRequest IVerifierMultiRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setMultiRequest", multiRequest)
}

// SetMultiRequest is a paid mutator transaction binding the contract method 0x7bc499e5.
//
// Solidity: function setMultiRequest((uint256,uint256[],uint256[],bytes) multiRequest) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetMultiRequest(multiRequest IVerifierMultiRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetMultiRequest(&_UniversalVerifier.TransactOpts, multiRequest)
}

// SetMultiRequest is a paid mutator transaction binding the contract method 0x7bc499e5.
//
// Solidity: function setMultiRequest((uint256,uint256[],uint256[],bytes) multiRequest) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetMultiRequest(multiRequest IVerifierMultiRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetMultiRequest(&_UniversalVerifier.TransactOpts, multiRequest)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0x7ad8dc61.
//
// Solidity: function setRequestOwner(uint256 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetRequestOwner(opts *bind.TransactOpts, requestId *big.Int, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setRequestOwner", requestId, requestOwner)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0x7ad8dc61.
//
// Solidity: function setRequestOwner(uint256 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetRequestOwner(requestId *big.Int, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequestOwner(&_UniversalVerifier.TransactOpts, requestId, requestOwner)
}

// SetRequestOwner is a paid mutator transaction binding the contract method 0x7ad8dc61.
//
// Solidity: function setRequestOwner(uint256 requestId, address requestOwner) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetRequestOwner(requestId *big.Int, requestOwner common.Address) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequestOwner(&_UniversalVerifier.TransactOpts, requestId, requestOwner)
}

// SetRequests is a paid mutator transaction binding the contract method 0xfd8a8d7d.
//
// Solidity: function setRequests((uint256,string,address,bytes,address)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetRequests(opts *bind.TransactOpts, requests []IVerifierRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setRequests", requests)
}

// SetRequests is a paid mutator transaction binding the contract method 0xfd8a8d7d.
//
// Solidity: function setRequests((uint256,string,address,bytes,address)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetRequests(requests []IVerifierRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequests(&_UniversalVerifier.TransactOpts, requests)
}

// SetRequests is a paid mutator transaction binding the contract method 0xfd8a8d7d.
//
// Solidity: function setRequests((uint256,string,address,bytes,address)[] requests) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetRequests(requests []IVerifierRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetRequests(&_UniversalVerifier.TransactOpts, requests)
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

// SetVerifierID is a paid mutator transaction binding the contract method 0xc087b0d4.
//
// Solidity: function setVerifierID(uint256 verifierID) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SetVerifierID(opts *bind.TransactOpts, verifierID *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "setVerifierID", verifierID)
}

// SetVerifierID is a paid mutator transaction binding the contract method 0xc087b0d4.
//
// Solidity: function setVerifierID(uint256 verifierID) returns()
func (_UniversalVerifier *UniversalVerifierSession) SetVerifierID(verifierID *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetVerifierID(&_UniversalVerifier.TransactOpts, verifierID)
}

// SetVerifierID is a paid mutator transaction binding the contract method 0xc087b0d4.
//
// Solidity: function setVerifierID(uint256 verifierID) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SetVerifierID(verifierID *big.Int) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SetVerifierID(&_UniversalVerifier.TransactOpts, verifierID)
}

// SubmitResponse is a paid mutator transaction binding the contract method 0x06c86a91.
//
// Solidity: function submitResponse((string,bytes) authResponse, (uint256,bytes,bytes)[] responses, bytes crossChainProofs) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) SubmitResponse(opts *bind.TransactOpts, authResponse IVerifierAuthResponse, responses []IVerifierResponse, crossChainProofs []byte) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "submitResponse", authResponse, responses, crossChainProofs)
}

// SubmitResponse is a paid mutator transaction binding the contract method 0x06c86a91.
//
// Solidity: function submitResponse((string,bytes) authResponse, (uint256,bytes,bytes)[] responses, bytes crossChainProofs) returns()
func (_UniversalVerifier *UniversalVerifierSession) SubmitResponse(authResponse IVerifierAuthResponse, responses []IVerifierResponse, crossChainProofs []byte) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitResponse(&_UniversalVerifier.TransactOpts, authResponse, responses, crossChainProofs)
}

// SubmitResponse is a paid mutator transaction binding the contract method 0x06c86a91.
//
// Solidity: function submitResponse((string,bytes) authResponse, (uint256,bytes,bytes)[] responses, bytes crossChainProofs) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) SubmitResponse(authResponse IVerifierAuthResponse, responses []IVerifierResponse, crossChainProofs []byte) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.SubmitResponse(&_UniversalVerifier.TransactOpts, authResponse, responses, crossChainProofs)
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

// UpdateRequest is a paid mutator transaction binding the contract method 0xa0366644.
//
// Solidity: function updateRequest((uint256,string,address,bytes,address) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactor) UpdateRequest(opts *bind.TransactOpts, request IVerifierRequest) (*types.Transaction, error) {
	return _UniversalVerifier.contract.Transact(opts, "updateRequest", request)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xa0366644.
//
// Solidity: function updateRequest((uint256,string,address,bytes,address) request) returns()
func (_UniversalVerifier *UniversalVerifierSession) UpdateRequest(request IVerifierRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UpdateRequest(&_UniversalVerifier.TransactOpts, request)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xa0366644.
//
// Solidity: function updateRequest((uint256,string,address,bytes,address) request) returns()
func (_UniversalVerifier *UniversalVerifierTransactorSession) UpdateRequest(request IVerifierRequest) (*types.Transaction, error) {
	return _UniversalVerifier.Contract.UpdateRequest(&_UniversalVerifier.TransactOpts, request)
}

// UniversalVerifierAuthMethodSetIterator is returned from FilterAuthMethodSet and is used to iterate over the raw logs and unpacked data for AuthMethodSet events raised by the UniversalVerifier contract.
type UniversalVerifierAuthMethodSetIterator struct {
	Event *UniversalVerifierAuthMethodSet // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierAuthMethodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierAuthMethodSet)
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
		it.Event = new(UniversalVerifierAuthMethodSet)
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
func (it *UniversalVerifierAuthMethodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierAuthMethodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierAuthMethodSet represents a AuthMethodSet event raised by the UniversalVerifier contract.
type UniversalVerifierAuthMethodSet struct {
	AuthMethod common.Hash
	Validator  common.Address
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthMethodSet is a free log retrieval operation binding the contract event 0xf9782ba766756e9c6ca90d64bfd1978844530aae1eab911d4630ffc1720e15e2.
//
// Solidity: event AuthMethodSet(string indexed authMethod, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterAuthMethodSet(opts *bind.FilterOpts, authMethod []string) (*UniversalVerifierAuthMethodSetIterator, error) {

	var authMethodRule []interface{}
	for _, authMethodItem := range authMethod {
		authMethodRule = append(authMethodRule, authMethodItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "AuthMethodSet", authMethodRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierAuthMethodSetIterator{contract: _UniversalVerifier.contract, event: "AuthMethodSet", logs: logs, sub: sub}, nil
}

// WatchAuthMethodSet is a free log subscription operation binding the contract event 0xf9782ba766756e9c6ca90d64bfd1978844530aae1eab911d4630ffc1720e15e2.
//
// Solidity: event AuthMethodSet(string indexed authMethod, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchAuthMethodSet(opts *bind.WatchOpts, sink chan<- *UniversalVerifierAuthMethodSet, authMethod []string) (event.Subscription, error) {

	var authMethodRule []interface{}
	for _, authMethodItem := range authMethod {
		authMethodRule = append(authMethodRule, authMethodItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "AuthMethodSet", authMethodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierAuthMethodSet)
				if err := _UniversalVerifier.contract.UnpackLog(event, "AuthMethodSet", log); err != nil {
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

// ParseAuthMethodSet is a log parse operation binding the contract event 0xf9782ba766756e9c6ca90d64bfd1978844530aae1eab911d4630ffc1720e15e2.
//
// Solidity: event AuthMethodSet(string indexed authMethod, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseAuthMethodSet(log types.Log) (*UniversalVerifierAuthMethodSet, error) {
	event := new(UniversalVerifierAuthMethodSet)
	if err := _UniversalVerifier.contract.UnpackLog(event, "AuthMethodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierAuthResponseSubmittedIterator is returned from FilterAuthResponseSubmitted and is used to iterate over the raw logs and unpacked data for AuthResponseSubmitted events raised by the UniversalVerifier contract.
type UniversalVerifierAuthResponseSubmittedIterator struct {
	Event *UniversalVerifierAuthResponseSubmitted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierAuthResponseSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierAuthResponseSubmitted)
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
		it.Event = new(UniversalVerifierAuthResponseSubmitted)
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
func (it *UniversalVerifierAuthResponseSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierAuthResponseSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierAuthResponseSubmitted represents a AuthResponseSubmitted event raised by the UniversalVerifier contract.
type UniversalVerifierAuthResponseSubmitted struct {
	AuthMethod common.Hash
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuthResponseSubmitted is a free log retrieval operation binding the contract event 0x16eed29a4411d6f696189e32f94cb87b6e2817e1aa0b9821f6d53191cfb5de77.
//
// Solidity: event AuthResponseSubmitted(string indexed authMethod, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterAuthResponseSubmitted(opts *bind.FilterOpts, authMethod []string, caller []common.Address) (*UniversalVerifierAuthResponseSubmittedIterator, error) {

	var authMethodRule []interface{}
	for _, authMethodItem := range authMethod {
		authMethodRule = append(authMethodRule, authMethodItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "AuthResponseSubmitted", authMethodRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierAuthResponseSubmittedIterator{contract: _UniversalVerifier.contract, event: "AuthResponseSubmitted", logs: logs, sub: sub}, nil
}

// WatchAuthResponseSubmitted is a free log subscription operation binding the contract event 0x16eed29a4411d6f696189e32f94cb87b6e2817e1aa0b9821f6d53191cfb5de77.
//
// Solidity: event AuthResponseSubmitted(string indexed authMethod, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchAuthResponseSubmitted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierAuthResponseSubmitted, authMethod []string, caller []common.Address) (event.Subscription, error) {

	var authMethodRule []interface{}
	for _, authMethodItem := range authMethod {
		authMethodRule = append(authMethodRule, authMethodItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "AuthResponseSubmitted", authMethodRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierAuthResponseSubmitted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "AuthResponseSubmitted", log); err != nil {
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

// ParseAuthResponseSubmitted is a log parse operation binding the contract event 0x16eed29a4411d6f696189e32f94cb87b6e2817e1aa0b9821f6d53191cfb5de77.
//
// Solidity: event AuthResponseSubmitted(string indexed authMethod, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseAuthResponseSubmitted(log types.Log) (*UniversalVerifierAuthResponseSubmitted, error) {
	event := new(UniversalVerifierAuthResponseSubmitted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "AuthResponseSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// UniversalVerifierMultiRequestSetIterator is returned from FilterMultiRequestSet and is used to iterate over the raw logs and unpacked data for MultiRequestSet events raised by the UniversalVerifier contract.
type UniversalVerifierMultiRequestSetIterator struct {
	Event *UniversalVerifierMultiRequestSet // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierMultiRequestSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierMultiRequestSet)
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
		it.Event = new(UniversalVerifierMultiRequestSet)
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
func (it *UniversalVerifierMultiRequestSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierMultiRequestSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierMultiRequestSet represents a MultiRequestSet event raised by the UniversalVerifier contract.
type UniversalVerifierMultiRequestSet struct {
	MultiRequestId *big.Int
	RequestIds     []*big.Int
	GroupIds       []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMultiRequestSet is a free log retrieval operation binding the contract event 0x3f79cffed6655f27556642061782f30ad1196bf040a8de80eee5d526da9d0bc5.
//
// Solidity: event MultiRequestSet(uint256 indexed multiRequestId, uint256[] requestIds, uint256[] groupIds)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterMultiRequestSet(opts *bind.FilterOpts, multiRequestId []*big.Int) (*UniversalVerifierMultiRequestSetIterator, error) {

	var multiRequestIdRule []interface{}
	for _, multiRequestIdItem := range multiRequestId {
		multiRequestIdRule = append(multiRequestIdRule, multiRequestIdItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "MultiRequestSet", multiRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierMultiRequestSetIterator{contract: _UniversalVerifier.contract, event: "MultiRequestSet", logs: logs, sub: sub}, nil
}

// WatchMultiRequestSet is a free log subscription operation binding the contract event 0x3f79cffed6655f27556642061782f30ad1196bf040a8de80eee5d526da9d0bc5.
//
// Solidity: event MultiRequestSet(uint256 indexed multiRequestId, uint256[] requestIds, uint256[] groupIds)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchMultiRequestSet(opts *bind.WatchOpts, sink chan<- *UniversalVerifierMultiRequestSet, multiRequestId []*big.Int) (event.Subscription, error) {

	var multiRequestIdRule []interface{}
	for _, multiRequestIdItem := range multiRequestId {
		multiRequestIdRule = append(multiRequestIdRule, multiRequestIdItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "MultiRequestSet", multiRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierMultiRequestSet)
				if err := _UniversalVerifier.contract.UnpackLog(event, "MultiRequestSet", log); err != nil {
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

// ParseMultiRequestSet is a log parse operation binding the contract event 0x3f79cffed6655f27556642061782f30ad1196bf040a8de80eee5d526da9d0bc5.
//
// Solidity: event MultiRequestSet(uint256 indexed multiRequestId, uint256[] requestIds, uint256[] groupIds)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseMultiRequestSet(log types.Log) (*UniversalVerifierMultiRequestSet, error) {
	event := new(UniversalVerifierMultiRequestSet)
	if err := _UniversalVerifier.contract.UnpackLog(event, "MultiRequestSet", log); err != nil {
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

// UniversalVerifierRequestSetIterator is returned from FilterRequestSet and is used to iterate over the raw logs and unpacked data for RequestSet events raised by the UniversalVerifier contract.
type UniversalVerifierRequestSetIterator struct {
	Event *UniversalVerifierRequestSet // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierRequestSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierRequestSet)
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
		it.Event = new(UniversalVerifierRequestSet)
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
func (it *UniversalVerifierRequestSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierRequestSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierRequestSet represents a RequestSet event raised by the UniversalVerifier contract.
type UniversalVerifierRequestSet struct {
	RequestId    *big.Int
	RequestOwner common.Address
	Metadata     string
	Validator    common.Address
	Params       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestSet is a free log retrieval operation binding the contract event 0x6c42013d79e93b13536220e027462dfb104def77ecb0a75ffbe92e853fc7cf09.
//
// Solidity: event RequestSet(uint256 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterRequestSet(opts *bind.FilterOpts, requestId []*big.Int, requestOwner []common.Address) (*UniversalVerifierRequestSetIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "RequestSet", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierRequestSetIterator{contract: _UniversalVerifier.contract, event: "RequestSet", logs: logs, sub: sub}, nil
}

// WatchRequestSet is a free log subscription operation binding the contract event 0x6c42013d79e93b13536220e027462dfb104def77ecb0a75ffbe92e853fc7cf09.
//
// Solidity: event RequestSet(uint256 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchRequestSet(opts *bind.WatchOpts, sink chan<- *UniversalVerifierRequestSet, requestId []*big.Int, requestOwner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "RequestSet", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierRequestSet)
				if err := _UniversalVerifier.contract.UnpackLog(event, "RequestSet", log); err != nil {
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

// ParseRequestSet is a log parse operation binding the contract event 0x6c42013d79e93b13536220e027462dfb104def77ecb0a75ffbe92e853fc7cf09.
//
// Solidity: event RequestSet(uint256 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseRequestSet(log types.Log) (*UniversalVerifierRequestSet, error) {
	event := new(UniversalVerifierRequestSet)
	if err := _UniversalVerifier.contract.UnpackLog(event, "RequestSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierRequestUpdateIterator is returned from FilterRequestUpdate and is used to iterate over the raw logs and unpacked data for RequestUpdate events raised by the UniversalVerifier contract.
type UniversalVerifierRequestUpdateIterator struct {
	Event *UniversalVerifierRequestUpdate // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierRequestUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierRequestUpdate)
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
		it.Event = new(UniversalVerifierRequestUpdate)
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
func (it *UniversalVerifierRequestUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierRequestUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierRequestUpdate represents a RequestUpdate event raised by the UniversalVerifier contract.
type UniversalVerifierRequestUpdate struct {
	RequestId    *big.Int
	RequestOwner common.Address
	Metadata     string
	Validator    common.Address
	Params       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestUpdate is a free log retrieval operation binding the contract event 0x2f517356721b15606e3bdde8939a6e814c9becc2fcef57d1c631893360c85736.
//
// Solidity: event RequestUpdate(uint256 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterRequestUpdate(opts *bind.FilterOpts, requestId []*big.Int, requestOwner []common.Address) (*UniversalVerifierRequestUpdateIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "RequestUpdate", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierRequestUpdateIterator{contract: _UniversalVerifier.contract, event: "RequestUpdate", logs: logs, sub: sub}, nil
}

// WatchRequestUpdate is a free log subscription operation binding the contract event 0x2f517356721b15606e3bdde8939a6e814c9becc2fcef57d1c631893360c85736.
//
// Solidity: event RequestUpdate(uint256 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchRequestUpdate(opts *bind.WatchOpts, sink chan<- *UniversalVerifierRequestUpdate, requestId []*big.Int, requestOwner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestOwnerRule []interface{}
	for _, requestOwnerItem := range requestOwner {
		requestOwnerRule = append(requestOwnerRule, requestOwnerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "RequestUpdate", requestIdRule, requestOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierRequestUpdate)
				if err := _UniversalVerifier.contract.UnpackLog(event, "RequestUpdate", log); err != nil {
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

// ParseRequestUpdate is a log parse operation binding the contract event 0x2f517356721b15606e3bdde8939a6e814c9becc2fcef57d1c631893360c85736.
//
// Solidity: event RequestUpdate(uint256 indexed requestId, address indexed requestOwner, string metadata, address validator, bytes params)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseRequestUpdate(log types.Log) (*UniversalVerifierRequestUpdate, error) {
	event := new(UniversalVerifierRequestUpdate)
	if err := _UniversalVerifier.contract.UnpackLog(event, "RequestUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalVerifierResponseSubmittedIterator is returned from FilterResponseSubmitted and is used to iterate over the raw logs and unpacked data for ResponseSubmitted events raised by the UniversalVerifier contract.
type UniversalVerifierResponseSubmittedIterator struct {
	Event *UniversalVerifierResponseSubmitted // Event containing the contract specifics and raw log

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
func (it *UniversalVerifierResponseSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalVerifierResponseSubmitted)
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
		it.Event = new(UniversalVerifierResponseSubmitted)
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
func (it *UniversalVerifierResponseSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalVerifierResponseSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalVerifierResponseSubmitted represents a ResponseSubmitted event raised by the UniversalVerifier contract.
type UniversalVerifierResponseSubmitted struct {
	RequestId *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResponseSubmitted is a free log retrieval operation binding the contract event 0xd34a91eddebc3a46f0193e4f54124f90bd00de34c64b5407562def103949d006.
//
// Solidity: event ResponseSubmitted(uint256 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) FilterResponseSubmitted(opts *bind.FilterOpts, requestId []*big.Int, caller []common.Address) (*UniversalVerifierResponseSubmittedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.FilterLogs(opts, "ResponseSubmitted", requestIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalVerifierResponseSubmittedIterator{contract: _UniversalVerifier.contract, event: "ResponseSubmitted", logs: logs, sub: sub}, nil
}

// WatchResponseSubmitted is a free log subscription operation binding the contract event 0xd34a91eddebc3a46f0193e4f54124f90bd00de34c64b5407562def103949d006.
//
// Solidity: event ResponseSubmitted(uint256 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) WatchResponseSubmitted(opts *bind.WatchOpts, sink chan<- *UniversalVerifierResponseSubmitted, requestId []*big.Int, caller []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _UniversalVerifier.contract.WatchLogs(opts, "ResponseSubmitted", requestIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalVerifierResponseSubmitted)
				if err := _UniversalVerifier.contract.UnpackLog(event, "ResponseSubmitted", log); err != nil {
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

// ParseResponseSubmitted is a log parse operation binding the contract event 0xd34a91eddebc3a46f0193e4f54124f90bd00de34c64b5407562def103949d006.
//
// Solidity: event ResponseSubmitted(uint256 indexed requestId, address indexed caller)
func (_UniversalVerifier *UniversalVerifierFilterer) ParseResponseSubmitted(log types.Log) (*UniversalVerifierResponseSubmitted, error) {
	event := new(UniversalVerifierResponseSubmitted)
	if err := _UniversalVerifier.contract.UnpackLog(event, "ResponseSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
