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

// IdentityLibStateInfo is an auto generated low-level Go binding around an user-defined struct.
type IdentityLibStateInfo struct {
	State           *big.Int
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

// NonMerklizedIssuerBaseMetaData contains all meta data concerning the NonMerklizedIssuerBase contract.
var NonMerklizedIssuerBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREDENTIAL_PROTOCOL_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"}],\"name\":\"getClaimProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getClaimProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"}],\"name\":\"getClaimProofWithStateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_credentialId\",\"type\":\"uint256\"}],\"name\":\"getCredential\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"context\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issuanceDate\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"credentialSchema\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"}],\"internalType\":\"structINonMerklizedIssuer.DisplayMethod\",\"name\":\"displayMethod\",\"type\":\"tuple\"}],\"internalType\":\"structINonMerklizedIssuer.CredentialData\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawValue\",\"type\":\"bytes\"}],\"internalType\":\"structINonMerklizedIssuer.SubjectField[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCredentialProtocolVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsOldStateGenesis\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedClaimsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRevocationsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRootsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRevocationProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationProofWithStateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatus\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatusByIdAndState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevocationsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rootsTreeRoot\",\"type\":\"uint256\"}],\"name\":\"getRootProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRootProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rootsTreeRoot\",\"type\":\"uint256\"}],\"name\":\"getRootProofWithStateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getRootsByState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.Roots\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmtDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateContractAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"}],\"name\":\"issueCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"}],\"name\":\"listUserCredentialIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_revocationNonce\",\"type\":\"uint64\"}],\"name\":\"revokeClaimAndTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) VERSION() (string, error) {
	return _NonMerklizedIssuerBase.Contract.VERSION(&_NonMerklizedIssuerBase.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) VERSION() (string, error) {
	return _NonMerklizedIssuerBase.Contract.VERSION(&_NonMerklizedIssuerBase.CallOpts)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetClaimProof(opts *bind.CallOpts, claimIndexHash *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getClaimProof", claimIndexHash)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimProof(&_NonMerklizedIssuerBase.CallOpts, claimIndexHash)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimProof(&_NonMerklizedIssuerBase.CallOpts, claimIndexHash)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetClaimProofByRoot(opts *bind.CallOpts, claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getClaimProofByRoot", claimIndexHash, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimProofByRoot(&_NonMerklizedIssuerBase.CallOpts, claimIndexHash, root)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimProofByRoot(&_NonMerklizedIssuerBase.CallOpts, claimIndexHash, root)
}

// GetClaimProofWithStateInfo is a free data retrieval call binding the contract method 0xb37feda4.
//
// Solidity: function getClaimProofWithStateInfo(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetClaimProofWithStateInfo(opts *bind.CallOpts, claimIndexHash *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getClaimProofWithStateInfo", claimIndexHash)

	if err != nil {
		return *new(SmtLibProof), *new(IdentityLibStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)
	out1 := *abi.ConvertType(out[1], new(IdentityLibStateInfo)).(*IdentityLibStateInfo)

	return out0, out1, err

}

// GetClaimProofWithStateInfo is a free data retrieval call binding the contract method 0xb37feda4.
//
// Solidity: function getClaimProofWithStateInfo(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetClaimProofWithStateInfo(claimIndexHash *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimProofWithStateInfo(&_NonMerklizedIssuerBase.CallOpts, claimIndexHash)
}

// GetClaimProofWithStateInfo is a free data retrieval call binding the contract method 0xb37feda4.
//
// Solidity: function getClaimProofWithStateInfo(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetClaimProofWithStateInfo(claimIndexHash *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimProofWithStateInfo(&_NonMerklizedIssuerBase.CallOpts, claimIndexHash)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetClaimsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getClaimsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetClaimsTreeRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimsTreeRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetClaimsTreeRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetClaimsTreeRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetCredential is a free data retrieval call binding the contract method 0x37c1d9ff.
//
// Solidity: function getCredential(uint256 _userId, uint256 _credentialId) view returns((uint256,string[],string,uint64,string,(string,string)), uint256[8], (string,uint256,bytes)[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetCredential(opts *bind.CallOpts, _userId *big.Int, _credentialId *big.Int) (INonMerklizedIssuerCredentialData, [8]*big.Int, []INonMerklizedIssuerSubjectField, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getCredential", _userId, _credentialId)

	if err != nil {
		return *new(INonMerklizedIssuerCredentialData), *new([8]*big.Int), *new([]INonMerklizedIssuerSubjectField), err
	}

	out0 := *abi.ConvertType(out[0], new(INonMerklizedIssuerCredentialData)).(*INonMerklizedIssuerCredentialData)
	out1 := *abi.ConvertType(out[1], new([8]*big.Int)).(*[8]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]INonMerklizedIssuerSubjectField)).(*[]INonMerklizedIssuerSubjectField)

	return out0, out1, out2, err

}

// GetCredential is a free data retrieval call binding the contract method 0x37c1d9ff.
//
// Solidity: function getCredential(uint256 _userId, uint256 _credentialId) view returns((uint256,string[],string,uint64,string,(string,string)), uint256[8], (string,uint256,bytes)[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetCredential(_userId *big.Int, _credentialId *big.Int) (INonMerklizedIssuerCredentialData, [8]*big.Int, []INonMerklizedIssuerSubjectField, error) {
	return _NonMerklizedIssuerBase.Contract.GetCredential(&_NonMerklizedIssuerBase.CallOpts, _userId, _credentialId)
}

// GetCredential is a free data retrieval call binding the contract method 0x37c1d9ff.
//
// Solidity: function getCredential(uint256 _userId, uint256 _credentialId) view returns((uint256,string[],string,uint64,string,(string,string)), uint256[8], (string,uint256,bytes)[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetCredential(_userId *big.Int, _credentialId *big.Int) (INonMerklizedIssuerCredentialData, [8]*big.Int, []INonMerklizedIssuerSubjectField, error) {
	return _NonMerklizedIssuerBase.Contract.GetCredential(&_NonMerklizedIssuerBase.CallOpts, _userId, _credentialId)
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

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetId() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetId(&_NonMerklizedIssuerBase.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetId() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetId(&_NonMerklizedIssuerBase.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetIsOldStateGenesis(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getIsOldStateGenesis")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetIsOldStateGenesis() (bool, error) {
	return _NonMerklizedIssuerBase.Contract.GetIsOldStateGenesis(&_NonMerklizedIssuerBase.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetIsOldStateGenesis() (bool, error) {
	return _NonMerklizedIssuerBase.Contract.GetIsOldStateGenesis(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetLatestPublishedClaimsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getLatestPublishedClaimsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedClaimsRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedClaimsRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetLatestPublishedRevocationsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getLatestPublishedRevocationsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedRevocationsRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedRevocationsRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetLatestPublishedRootsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getLatestPublishedRootsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedRootsRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedRootsRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetLatestPublishedState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getLatestPublishedState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetLatestPublishedState() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedState(&_NonMerklizedIssuerBase.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetLatestPublishedState() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetLatestPublishedState(&_NonMerklizedIssuerBase.CallOpts)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRevocationProof(opts *bind.CallOpts, revocationNonce uint64) (SmtLibProof, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRevocationProof", revocationNonce)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationProof(&_NonMerklizedIssuerBase.CallOpts, revocationNonce)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationProof(&_NonMerklizedIssuerBase.CallOpts, revocationNonce)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRevocationProofByRoot(opts *bind.CallOpts, revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRevocationProofByRoot", revocationNonce, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationProofByRoot(&_NonMerklizedIssuerBase.CallOpts, revocationNonce, root)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationProofByRoot(&_NonMerklizedIssuerBase.CallOpts, revocationNonce, root)
}

// GetRevocationProofWithStateInfo is a free data retrieval call binding the contract method 0x0033058d.
//
// Solidity: function getRevocationProofWithStateInfo(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRevocationProofWithStateInfo(opts *bind.CallOpts, revocationNonce uint64) (SmtLibProof, IdentityLibStateInfo, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRevocationProofWithStateInfo", revocationNonce)

	if err != nil {
		return *new(SmtLibProof), *new(IdentityLibStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)
	out1 := *abi.ConvertType(out[1], new(IdentityLibStateInfo)).(*IdentityLibStateInfo)

	return out0, out1, err

}

// GetRevocationProofWithStateInfo is a free data retrieval call binding the contract method 0x0033058d.
//
// Solidity: function getRevocationProofWithStateInfo(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRevocationProofWithStateInfo(revocationNonce uint64) (SmtLibProof, IdentityLibStateInfo, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationProofWithStateInfo(&_NonMerklizedIssuerBase.CallOpts, revocationNonce)
}

// GetRevocationProofWithStateInfo is a free data retrieval call binding the contract method 0x0033058d.
//
// Solidity: function getRevocationProofWithStateInfo(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRevocationProofWithStateInfo(revocationNonce uint64) (SmtLibProof, IdentityLibStateInfo, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationProofWithStateInfo(&_NonMerklizedIssuerBase.CallOpts, revocationNonce)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRevocationStatus(opts *bind.CallOpts, id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRevocationStatus", id, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationStatus(&_NonMerklizedIssuerBase.CallOpts, id, nonce)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationStatus(&_NonMerklizedIssuerBase.CallOpts, id, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRevocationStatusByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRevocationStatusByIdAndState", id, state, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationStatusByIdAndState(&_NonMerklizedIssuerBase.CallOpts, id, state, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationStatusByIdAndState(&_NonMerklizedIssuerBase.CallOpts, id, state, nonce)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRevocationsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRevocationsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationsTreeRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetRevocationsTreeRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRootProof(opts *bind.CallOpts, rootsTreeRoot *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRootProof", rootsTreeRoot)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRootProof(rootsTreeRoot *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootProof(&_NonMerklizedIssuerBase.CallOpts, rootsTreeRoot)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRootProof(rootsTreeRoot *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootProof(&_NonMerklizedIssuerBase.CallOpts, rootsTreeRoot)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRootProofByRoot(opts *bind.CallOpts, claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRootProofByRoot", claimsTreeRoot, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootProofByRoot(&_NonMerklizedIssuerBase.CallOpts, claimsTreeRoot, root)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootProofByRoot(&_NonMerklizedIssuerBase.CallOpts, claimsTreeRoot, root)
}

// GetRootProofWithStateInfo is a free data retrieval call binding the contract method 0x443d7534.
//
// Solidity: function getRootProofWithStateInfo(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRootProofWithStateInfo(opts *bind.CallOpts, rootsTreeRoot *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRootProofWithStateInfo", rootsTreeRoot)

	if err != nil {
		return *new(SmtLibProof), *new(IdentityLibStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)
	out1 := *abi.ConvertType(out[1], new(IdentityLibStateInfo)).(*IdentityLibStateInfo)

	return out0, out1, err

}

// GetRootProofWithStateInfo is a free data retrieval call binding the contract method 0x443d7534.
//
// Solidity: function getRootProofWithStateInfo(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRootProofWithStateInfo(rootsTreeRoot *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootProofWithStateInfo(&_NonMerklizedIssuerBase.CallOpts, rootsTreeRoot)
}

// GetRootProofWithStateInfo is a free data retrieval call binding the contract method 0x443d7534.
//
// Solidity: function getRootProofWithStateInfo(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRootProofWithStateInfo(rootsTreeRoot *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootProofWithStateInfo(&_NonMerklizedIssuerBase.CallOpts, rootsTreeRoot)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRootsByState(opts *bind.CallOpts, state *big.Int) (IdentityLibRoots, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRootsByState", state)

	if err != nil {
		return *new(IdentityLibRoots), err
	}

	out0 := *abi.ConvertType(out[0], new(IdentityLibRoots)).(*IdentityLibRoots)

	return out0, err

}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootsByState(&_NonMerklizedIssuerBase.CallOpts, state)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootsByState(&_NonMerklizedIssuerBase.CallOpts, state)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetRootsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getRootsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetRootsTreeRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootsTreeRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetRootsTreeRoot() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetRootsTreeRoot(&_NonMerklizedIssuerBase.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) GetSmtDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "getSmtDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) GetSmtDepth() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetSmtDepth(&_NonMerklizedIssuerBase.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) GetSmtDepth() (*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.GetSmtDepth(&_NonMerklizedIssuerBase.CallOpts)
}

// ListUserCredentialIds is a free data retrieval call binding the contract method 0x58381619.
//
// Solidity: function listUserCredentialIds(uint256 _userId) view returns(uint256[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) ListUserCredentialIds(opts *bind.CallOpts, _userId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "listUserCredentialIds", _userId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListUserCredentialIds is a free data retrieval call binding the contract method 0x58381619.
//
// Solidity: function listUserCredentialIds(uint256 _userId) view returns(uint256[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) ListUserCredentialIds(_userId *big.Int) ([]*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.ListUserCredentialIds(&_NonMerklizedIssuerBase.CallOpts, _userId)
}

// ListUserCredentialIds is a free data retrieval call binding the contract method 0x58381619.
//
// Solidity: function listUserCredentialIds(uint256 _userId) view returns(uint256[])
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) ListUserCredentialIds(_userId *big.Int) ([]*big.Int, error) {
	return _NonMerklizedIssuerBase.Contract.ListUserCredentialIds(&_NonMerklizedIssuerBase.CallOpts, _userId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NonMerklizedIssuerBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) Owner() (common.Address, error) {
	return _NonMerklizedIssuerBase.Contract.Owner(&_NonMerklizedIssuerBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseCallerSession) Owner() (common.Address, error) {
	return _NonMerklizedIssuerBase.Contract.Owner(&_NonMerklizedIssuerBase.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) Initialize(opts *bind.TransactOpts, _stateContractAddr common.Address) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "initialize", _stateContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) Initialize(_stateContractAddr common.Address) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.Initialize(&_NonMerklizedIssuerBase.TransactOpts, _stateContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) Initialize(_stateContractAddr common.Address) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.Initialize(&_NonMerklizedIssuerBase.TransactOpts, _stateContractAddr)
}

// IssueCredential is a paid mutator transaction binding the contract method 0x7d4ee3df.
//
// Solidity: function issueCredential(uint256 _userId) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) IssueCredential(opts *bind.TransactOpts, _userId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "issueCredential", _userId)
}

// IssueCredential is a paid mutator transaction binding the contract method 0x7d4ee3df.
//
// Solidity: function issueCredential(uint256 _userId) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) IssueCredential(_userId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.IssueCredential(&_NonMerklizedIssuerBase.TransactOpts, _userId)
}

// IssueCredential is a paid mutator transaction binding the contract method 0x7d4ee3df.
//
// Solidity: function issueCredential(uint256 _userId) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) IssueCredential(_userId *big.Int) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.IssueCredential(&_NonMerklizedIssuerBase.TransactOpts, _userId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.RenounceOwnership(&_NonMerklizedIssuerBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.RenounceOwnership(&_NonMerklizedIssuerBase.TransactOpts)
}

// RevokeClaimAndTransit is a paid mutator transaction binding the contract method 0xf2a8ed5a.
//
// Solidity: function revokeClaimAndTransit(uint64 _revocationNonce) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) RevokeClaimAndTransit(opts *bind.TransactOpts, _revocationNonce uint64) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "revokeClaimAndTransit", _revocationNonce)
}

// RevokeClaimAndTransit is a paid mutator transaction binding the contract method 0xf2a8ed5a.
//
// Solidity: function revokeClaimAndTransit(uint64 _revocationNonce) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) RevokeClaimAndTransit(_revocationNonce uint64) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.RevokeClaimAndTransit(&_NonMerklizedIssuerBase.TransactOpts, _revocationNonce)
}

// RevokeClaimAndTransit is a paid mutator transaction binding the contract method 0xf2a8ed5a.
//
// Solidity: function revokeClaimAndTransit(uint64 _revocationNonce) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) RevokeClaimAndTransit(_revocationNonce uint64) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.RevokeClaimAndTransit(&_NonMerklizedIssuerBase.TransactOpts, _revocationNonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.TransferOwnership(&_NonMerklizedIssuerBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NonMerklizedIssuerBase.Contract.TransferOwnership(&_NonMerklizedIssuerBase.TransactOpts, newOwner)
}

// NonMerklizedIssuerBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NonMerklizedIssuerBase contract.
type NonMerklizedIssuerBaseInitializedIterator struct {
	Event *NonMerklizedIssuerBaseInitialized // Event containing the contract specifics and raw log

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
func (it *NonMerklizedIssuerBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonMerklizedIssuerBaseInitialized)
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
		it.Event = new(NonMerklizedIssuerBaseInitialized)
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
func (it *NonMerklizedIssuerBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonMerklizedIssuerBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonMerklizedIssuerBaseInitialized represents a Initialized event raised by the NonMerklizedIssuerBase contract.
type NonMerklizedIssuerBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*NonMerklizedIssuerBaseInitializedIterator, error) {

	logs, sub, err := _NonMerklizedIssuerBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NonMerklizedIssuerBaseInitializedIterator{contract: _NonMerklizedIssuerBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NonMerklizedIssuerBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _NonMerklizedIssuerBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonMerklizedIssuerBaseInitialized)
				if err := _NonMerklizedIssuerBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseFilterer) ParseInitialized(log types.Log) (*NonMerklizedIssuerBaseInitialized, error) {
	event := new(NonMerklizedIssuerBaseInitialized)
	if err := _NonMerklizedIssuerBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NonMerklizedIssuerBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NonMerklizedIssuerBase contract.
type NonMerklizedIssuerBaseOwnershipTransferredIterator struct {
	Event *NonMerklizedIssuerBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NonMerklizedIssuerBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonMerklizedIssuerBaseOwnershipTransferred)
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
		it.Event = new(NonMerklizedIssuerBaseOwnershipTransferred)
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
func (it *NonMerklizedIssuerBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonMerklizedIssuerBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonMerklizedIssuerBaseOwnershipTransferred represents a OwnershipTransferred event raised by the NonMerklizedIssuerBase contract.
type NonMerklizedIssuerBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NonMerklizedIssuerBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NonMerklizedIssuerBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NonMerklizedIssuerBaseOwnershipTransferredIterator{contract: _NonMerklizedIssuerBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NonMerklizedIssuerBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NonMerklizedIssuerBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonMerklizedIssuerBaseOwnershipTransferred)
				if err := _NonMerklizedIssuerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NonMerklizedIssuerBase *NonMerklizedIssuerBaseFilterer) ParseOwnershipTransferred(log types.Log) (*NonMerklizedIssuerBaseOwnershipTransferred, error) {
	event := new(NonMerklizedIssuerBaseOwnershipTransferred)
	if err := _NonMerklizedIssuerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
