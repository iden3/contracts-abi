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

// IZKPVerifierZKPResponse is an auto generated low-level Go binding around an user-defined struct.
type IZKPVerifierZKPResponse struct {
	RequestId uint64
	ZkProof   []byte
	Data      []byte
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

// PassportCredentialIssuerImplV1MetaData contains all meta data concerning the PassportCredentialIssuerImplV1 contract.
var PassportCredentialIssuerImplV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredentialProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDscCommitmentRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hashIndex\",\"type\":\"uint256\"}],\"name\":\"InvalidHashIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hashValue\",\"type\":\"uint256\"}],\"name\":\"InvalidHashValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkId1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkId2\",\"type\":\"uint256\"}],\"name\":\"InvalidLinkId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"InvalidResponsesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"templateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedTemplateRoot\",\"type\":\"uint256\"}],\"name\":\"InvalidTemplateRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"issuanceDate\",\"type\":\"uint256\"}],\"name\":\"IssuanceDateExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length2\",\"type\":\"uint256\"}],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"NoCredentialCircuitForRequestId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"NoSignatureCircuitForRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoVerifierSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"}],\"name\":\"NullifierAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"circuitId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"CredentialCircuitVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"ExpirationTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"circuitId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"SignatureCircuitVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tampletRoot\",\"type\":\"uint256\"}],\"name\":\"TemplateRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"}],\"name\":\"cleanNullifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"circuitId\",\"type\":\"string\"}],\"name\":\"credentialCircuitIdToRequestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"credentialRequestIdToCircuitIds\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"circuitId\",\"type\":\"string\"}],\"name\":\"credentialVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"}],\"name\":\"getClaimProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getClaimProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"}],\"name\":\"getClaimProofWithStateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsOldStateGenesis\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedClaimsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRevocationsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRootsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRevocationProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationProofWithStateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatus\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatusByIdAndState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevocationsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rootsTreeRoot\",\"type\":\"uint256\"}],\"name\":\"getRootProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRootProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rootsTreeRoot\",\"type\":\"uint256\"}],\"name\":\"getRootProofWithStateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getRootsByState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.Roots\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmtDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTemplateRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateContractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"idType\",\"type\":\"bytes2\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"templateRoot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"credentialCircuitIds\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"credentialVerifierAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"signatureCircuitIds\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"signatureVerifierAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"stateAddress\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"idType\",\"type\":\"bytes2\"}],\"name\":\"initializeIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"setExpirationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"templateRoot\",\"type\":\"uint256\"}],\"name\":\"setTemplateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"circuitId\",\"type\":\"string\"}],\"name\":\"signatureCircuitIdToRequestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"signatureRequestIdToCircuitIds\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"circuitId\",\"type\":\"string\"}],\"name\":\"signatureVerifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"requestId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"zkProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIZKPVerifier.ZKPResponse[]\",\"name\":\"responses\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"crossChainProofs\",\"type\":\"bytes\"}],\"name\":\"submitZKPResponseV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"circuitIds\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"verifierAddresses\",\"type\":\"address[]\"}],\"name\":\"updateCredentialVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"circuitIds\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"verifierAddresses\",\"type\":\"address[]\"}],\"name\":\"updateSignatureVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// PassportCredentialIssuerImplV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use PassportCredentialIssuerImplV1MetaData.ABI instead.
var PassportCredentialIssuerImplV1ABI = PassportCredentialIssuerImplV1MetaData.ABI

// PassportCredentialIssuerImplV1 is an auto generated Go binding around an Ethereum contract.
type PassportCredentialIssuerImplV1 struct {
	PassportCredentialIssuerImplV1Caller     // Read-only binding to the contract
	PassportCredentialIssuerImplV1Transactor // Write-only binding to the contract
	PassportCredentialIssuerImplV1Filterer   // Log filterer for contract events
}

// PassportCredentialIssuerImplV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type PassportCredentialIssuerImplV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportCredentialIssuerImplV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PassportCredentialIssuerImplV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportCredentialIssuerImplV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PassportCredentialIssuerImplV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PassportCredentialIssuerImplV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PassportCredentialIssuerImplV1Session struct {
	Contract     *PassportCredentialIssuerImplV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PassportCredentialIssuerImplV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PassportCredentialIssuerImplV1CallerSession struct {
	Contract *PassportCredentialIssuerImplV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// PassportCredentialIssuerImplV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PassportCredentialIssuerImplV1TransactorSession struct {
	Contract     *PassportCredentialIssuerImplV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// PassportCredentialIssuerImplV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type PassportCredentialIssuerImplV1Raw struct {
	Contract *PassportCredentialIssuerImplV1 // Generic contract binding to access the raw methods on
}

// PassportCredentialIssuerImplV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PassportCredentialIssuerImplV1CallerRaw struct {
	Contract *PassportCredentialIssuerImplV1Caller // Generic read-only contract binding to access the raw methods on
}

// PassportCredentialIssuerImplV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PassportCredentialIssuerImplV1TransactorRaw struct {
	Contract *PassportCredentialIssuerImplV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPassportCredentialIssuerImplV1 creates a new instance of PassportCredentialIssuerImplV1, bound to a specific deployed contract.
func NewPassportCredentialIssuerImplV1(address common.Address, backend bind.ContractBackend) (*PassportCredentialIssuerImplV1, error) {
	contract, err := bindPassportCredentialIssuerImplV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1{PassportCredentialIssuerImplV1Caller: PassportCredentialIssuerImplV1Caller{contract: contract}, PassportCredentialIssuerImplV1Transactor: PassportCredentialIssuerImplV1Transactor{contract: contract}, PassportCredentialIssuerImplV1Filterer: PassportCredentialIssuerImplV1Filterer{contract: contract}}, nil
}

// NewPassportCredentialIssuerImplV1Caller creates a new read-only instance of PassportCredentialIssuerImplV1, bound to a specific deployed contract.
func NewPassportCredentialIssuerImplV1Caller(address common.Address, caller bind.ContractCaller) (*PassportCredentialIssuerImplV1Caller, error) {
	contract, err := bindPassportCredentialIssuerImplV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1Caller{contract: contract}, nil
}

// NewPassportCredentialIssuerImplV1Transactor creates a new write-only instance of PassportCredentialIssuerImplV1, bound to a specific deployed contract.
func NewPassportCredentialIssuerImplV1Transactor(address common.Address, transactor bind.ContractTransactor) (*PassportCredentialIssuerImplV1Transactor, error) {
	contract, err := bindPassportCredentialIssuerImplV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1Transactor{contract: contract}, nil
}

// NewPassportCredentialIssuerImplV1Filterer creates a new log filterer instance of PassportCredentialIssuerImplV1, bound to a specific deployed contract.
func NewPassportCredentialIssuerImplV1Filterer(address common.Address, filterer bind.ContractFilterer) (*PassportCredentialIssuerImplV1Filterer, error) {
	contract, err := bindPassportCredentialIssuerImplV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1Filterer{contract: contract}, nil
}

// bindPassportCredentialIssuerImplV1 binds a generic wrapper to an already deployed contract.
func bindPassportCredentialIssuerImplV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PassportCredentialIssuerImplV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PassportCredentialIssuerImplV1.Contract.PassportCredentialIssuerImplV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.PassportCredentialIssuerImplV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.PassportCredentialIssuerImplV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PassportCredentialIssuerImplV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.UPGRADEINTERFACEVERSION(&_PassportCredentialIssuerImplV1.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.UPGRADEINTERFACEVERSION(&_PassportCredentialIssuerImplV1.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) VERSION() (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.VERSION(&_PassportCredentialIssuerImplV1.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) VERSION() (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.VERSION(&_PassportCredentialIssuerImplV1.CallOpts)
}

// CredentialCircuitIdToRequestIds is a free data retrieval call binding the contract method 0xcfa350c8.
//
// Solidity: function credentialCircuitIdToRequestIds(string circuitId) view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) CredentialCircuitIdToRequestIds(opts *bind.CallOpts, circuitId string) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "credentialCircuitIdToRequestIds", circuitId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CredentialCircuitIdToRequestIds is a free data retrieval call binding the contract method 0xcfa350c8.
//
// Solidity: function credentialCircuitIdToRequestIds(string circuitId) view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) CredentialCircuitIdToRequestIds(circuitId string) (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.CredentialCircuitIdToRequestIds(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// CredentialCircuitIdToRequestIds is a free data retrieval call binding the contract method 0xcfa350c8.
//
// Solidity: function credentialCircuitIdToRequestIds(string circuitId) view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) CredentialCircuitIdToRequestIds(circuitId string) (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.CredentialCircuitIdToRequestIds(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// CredentialRequestIdToCircuitIds is a free data retrieval call binding the contract method 0xd622ea6e.
//
// Solidity: function credentialRequestIdToCircuitIds(uint256 requestId) view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) CredentialRequestIdToCircuitIds(opts *bind.CallOpts, requestId *big.Int) (string, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "credentialRequestIdToCircuitIds", requestId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CredentialRequestIdToCircuitIds is a free data retrieval call binding the contract method 0xd622ea6e.
//
// Solidity: function credentialRequestIdToCircuitIds(uint256 requestId) view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) CredentialRequestIdToCircuitIds(requestId *big.Int) (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.CredentialRequestIdToCircuitIds(&_PassportCredentialIssuerImplV1.CallOpts, requestId)
}

// CredentialRequestIdToCircuitIds is a free data retrieval call binding the contract method 0xd622ea6e.
//
// Solidity: function credentialRequestIdToCircuitIds(uint256 requestId) view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) CredentialRequestIdToCircuitIds(requestId *big.Int) (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.CredentialRequestIdToCircuitIds(&_PassportCredentialIssuerImplV1.CallOpts, requestId)
}

// CredentialVerifiers is a free data retrieval call binding the contract method 0xa635c785.
//
// Solidity: function credentialVerifiers(string circuitId) view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) CredentialVerifiers(opts *bind.CallOpts, circuitId string) (common.Address, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "credentialVerifiers", circuitId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredentialVerifiers is a free data retrieval call binding the contract method 0xa635c785.
//
// Solidity: function credentialVerifiers(string circuitId) view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) CredentialVerifiers(circuitId string) (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.CredentialVerifiers(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// CredentialVerifiers is a free data retrieval call binding the contract method 0xa635c785.
//
// Solidity: function credentialVerifiers(string circuitId) view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) CredentialVerifiers(circuitId string) (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.CredentialVerifiers(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetClaimProof(opts *bind.CallOpts, claimIndexHash *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getClaimProof", claimIndexHash)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimProof(&_PassportCredentialIssuerImplV1.CallOpts, claimIndexHash)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimProof(&_PassportCredentialIssuerImplV1.CallOpts, claimIndexHash)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetClaimProofByRoot(opts *bind.CallOpts, claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getClaimProofByRoot", claimIndexHash, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimProofByRoot(&_PassportCredentialIssuerImplV1.CallOpts, claimIndexHash, root)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimProofByRoot(&_PassportCredentialIssuerImplV1.CallOpts, claimIndexHash, root)
}

// GetClaimProofWithStateInfo is a free data retrieval call binding the contract method 0xb37feda4.
//
// Solidity: function getClaimProofWithStateInfo(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetClaimProofWithStateInfo(opts *bind.CallOpts, claimIndexHash *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getClaimProofWithStateInfo", claimIndexHash)

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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetClaimProofWithStateInfo(claimIndexHash *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimProofWithStateInfo(&_PassportCredentialIssuerImplV1.CallOpts, claimIndexHash)
}

// GetClaimProofWithStateInfo is a free data retrieval call binding the contract method 0xb37feda4.
//
// Solidity: function getClaimProofWithStateInfo(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetClaimProofWithStateInfo(claimIndexHash *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimProofWithStateInfo(&_PassportCredentialIssuerImplV1.CallOpts, claimIndexHash)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetClaimsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getClaimsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetClaimsTreeRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimsTreeRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetClaimsTreeRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetClaimsTreeRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetExpirationTime is a free data retrieval call binding the contract method 0xe23a845a.
//
// Solidity: function getExpirationTime() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getExpirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpirationTime is a free data retrieval call binding the contract method 0xe23a845a.
//
// Solidity: function getExpirationTime() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetExpirationTime() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetExpirationTime(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetExpirationTime is a free data retrieval call binding the contract method 0xe23a845a.
//
// Solidity: function getExpirationTime() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetExpirationTime() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetExpirationTime(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetId() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetId(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetId() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetId(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetIsOldStateGenesis(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getIsOldStateGenesis")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetIsOldStateGenesis() (bool, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetIsOldStateGenesis(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetIsOldStateGenesis() (bool, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetIsOldStateGenesis(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetLatestPublishedClaimsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getLatestPublishedClaimsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedClaimsRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedClaimsRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetLatestPublishedRevocationsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getLatestPublishedRevocationsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedRevocationsRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedRevocationsRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetLatestPublishedRootsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getLatestPublishedRootsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedRootsRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedRootsRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetLatestPublishedState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getLatestPublishedState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetLatestPublishedState() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedState(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetLatestPublishedState() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetLatestPublishedState(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRegistry() (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRegistry(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRegistry() (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRegistry(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRevocationProof(opts *bind.CallOpts, revocationNonce uint64) (SmtLibProof, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRevocationProof", revocationNonce)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationProof(&_PassportCredentialIssuerImplV1.CallOpts, revocationNonce)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationProof(&_PassportCredentialIssuerImplV1.CallOpts, revocationNonce)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRevocationProofByRoot(opts *bind.CallOpts, revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRevocationProofByRoot", revocationNonce, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationProofByRoot(&_PassportCredentialIssuerImplV1.CallOpts, revocationNonce, root)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationProofByRoot(&_PassportCredentialIssuerImplV1.CallOpts, revocationNonce, root)
}

// GetRevocationProofWithStateInfo is a free data retrieval call binding the contract method 0x0033058d.
//
// Solidity: function getRevocationProofWithStateInfo(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRevocationProofWithStateInfo(opts *bind.CallOpts, revocationNonce uint64) (SmtLibProof, IdentityLibStateInfo, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRevocationProofWithStateInfo", revocationNonce)

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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRevocationProofWithStateInfo(revocationNonce uint64) (SmtLibProof, IdentityLibStateInfo, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationProofWithStateInfo(&_PassportCredentialIssuerImplV1.CallOpts, revocationNonce)
}

// GetRevocationProofWithStateInfo is a free data retrieval call binding the contract method 0x0033058d.
//
// Solidity: function getRevocationProofWithStateInfo(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRevocationProofWithStateInfo(revocationNonce uint64) (SmtLibProof, IdentityLibStateInfo, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationProofWithStateInfo(&_PassportCredentialIssuerImplV1.CallOpts, revocationNonce)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRevocationStatus(opts *bind.CallOpts, id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRevocationStatus", id, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationStatus(&_PassportCredentialIssuerImplV1.CallOpts, id, nonce)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationStatus(&_PassportCredentialIssuerImplV1.CallOpts, id, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRevocationStatusByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRevocationStatusByIdAndState", id, state, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationStatusByIdAndState(&_PassportCredentialIssuerImplV1.CallOpts, id, state, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationStatusByIdAndState(&_PassportCredentialIssuerImplV1.CallOpts, id, state, nonce)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRevocationsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRevocationsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRevocationsTreeRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationsTreeRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRevocationsTreeRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRootProof(opts *bind.CallOpts, rootsTreeRoot *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRootProof", rootsTreeRoot)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRootProof(rootsTreeRoot *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootProof(&_PassportCredentialIssuerImplV1.CallOpts, rootsTreeRoot)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRootProof(rootsTreeRoot *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootProof(&_PassportCredentialIssuerImplV1.CallOpts, rootsTreeRoot)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRootProofByRoot(opts *bind.CallOpts, claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRootProofByRoot", claimsTreeRoot, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootProofByRoot(&_PassportCredentialIssuerImplV1.CallOpts, claimsTreeRoot, root)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootProofByRoot(&_PassportCredentialIssuerImplV1.CallOpts, claimsTreeRoot, root)
}

// GetRootProofWithStateInfo is a free data retrieval call binding the contract method 0x443d7534.
//
// Solidity: function getRootProofWithStateInfo(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRootProofWithStateInfo(opts *bind.CallOpts, rootsTreeRoot *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRootProofWithStateInfo", rootsTreeRoot)

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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRootProofWithStateInfo(rootsTreeRoot *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootProofWithStateInfo(&_PassportCredentialIssuerImplV1.CallOpts, rootsTreeRoot)
}

// GetRootProofWithStateInfo is a free data retrieval call binding the contract method 0x443d7534.
//
// Solidity: function getRootProofWithStateInfo(uint256 rootsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256), (uint256,uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRootProofWithStateInfo(rootsTreeRoot *big.Int) (SmtLibProof, IdentityLibStateInfo, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootProofWithStateInfo(&_PassportCredentialIssuerImplV1.CallOpts, rootsTreeRoot)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRootsByState(opts *bind.CallOpts, state *big.Int) (IdentityLibRoots, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRootsByState", state)

	if err != nil {
		return *new(IdentityLibRoots), err
	}

	out0 := *abi.ConvertType(out[0], new(IdentityLibRoots)).(*IdentityLibRoots)

	return out0, err

}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootsByState(&_PassportCredentialIssuerImplV1.CallOpts, state)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootsByState(&_PassportCredentialIssuerImplV1.CallOpts, state)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetRootsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getRootsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetRootsTreeRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootsTreeRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetRootsTreeRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetRootsTreeRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetSmtDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getSmtDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetSmtDepth() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetSmtDepth(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetSmtDepth() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetSmtDepth(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetTemplateRoot is a free data retrieval call binding the contract method 0x2c1d3204.
//
// Solidity: function getTemplateRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) GetTemplateRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "getTemplateRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTemplateRoot is a free data retrieval call binding the contract method 0x2c1d3204.
//
// Solidity: function getTemplateRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) GetTemplateRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetTemplateRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// GetTemplateRoot is a free data retrieval call binding the contract method 0x2c1d3204.
//
// Solidity: function getTemplateRoot() view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) GetTemplateRoot() (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.GetTemplateRoot(&_PassportCredentialIssuerImplV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) Owner() (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.Owner(&_PassportCredentialIssuerImplV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) Owner() (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.Owner(&_PassportCredentialIssuerImplV1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) PendingOwner() (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.PendingOwner(&_PassportCredentialIssuerImplV1.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) PendingOwner() (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.PendingOwner(&_PassportCredentialIssuerImplV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) ProxiableUUID() ([32]byte, error) {
	return _PassportCredentialIssuerImplV1.Contract.ProxiableUUID(&_PassportCredentialIssuerImplV1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _PassportCredentialIssuerImplV1.Contract.ProxiableUUID(&_PassportCredentialIssuerImplV1.CallOpts)
}

// SignatureCircuitIdToRequestIds is a free data retrieval call binding the contract method 0x8affdc60.
//
// Solidity: function signatureCircuitIdToRequestIds(string circuitId) view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) SignatureCircuitIdToRequestIds(opts *bind.CallOpts, circuitId string) (*big.Int, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "signatureCircuitIdToRequestIds", circuitId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatureCircuitIdToRequestIds is a free data retrieval call binding the contract method 0x8affdc60.
//
// Solidity: function signatureCircuitIdToRequestIds(string circuitId) view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SignatureCircuitIdToRequestIds(circuitId string) (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.SignatureCircuitIdToRequestIds(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// SignatureCircuitIdToRequestIds is a free data retrieval call binding the contract method 0x8affdc60.
//
// Solidity: function signatureCircuitIdToRequestIds(string circuitId) view returns(uint256)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) SignatureCircuitIdToRequestIds(circuitId string) (*big.Int, error) {
	return _PassportCredentialIssuerImplV1.Contract.SignatureCircuitIdToRequestIds(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// SignatureRequestIdToCircuitIds is a free data retrieval call binding the contract method 0xdbdf1852.
//
// Solidity: function signatureRequestIdToCircuitIds(uint256 requestId) view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) SignatureRequestIdToCircuitIds(opts *bind.CallOpts, requestId *big.Int) (string, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "signatureRequestIdToCircuitIds", requestId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SignatureRequestIdToCircuitIds is a free data retrieval call binding the contract method 0xdbdf1852.
//
// Solidity: function signatureRequestIdToCircuitIds(uint256 requestId) view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SignatureRequestIdToCircuitIds(requestId *big.Int) (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.SignatureRequestIdToCircuitIds(&_PassportCredentialIssuerImplV1.CallOpts, requestId)
}

// SignatureRequestIdToCircuitIds is a free data retrieval call binding the contract method 0xdbdf1852.
//
// Solidity: function signatureRequestIdToCircuitIds(uint256 requestId) view returns(string)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) SignatureRequestIdToCircuitIds(requestId *big.Int) (string, error) {
	return _PassportCredentialIssuerImplV1.Contract.SignatureRequestIdToCircuitIds(&_PassportCredentialIssuerImplV1.CallOpts, requestId)
}

// SignatureVerifiers is a free data retrieval call binding the contract method 0xcde1cacb.
//
// Solidity: function signatureVerifiers(string circuitId) view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) SignatureVerifiers(opts *bind.CallOpts, circuitId string) (common.Address, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "signatureVerifiers", circuitId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignatureVerifiers is a free data retrieval call binding the contract method 0xcde1cacb.
//
// Solidity: function signatureVerifiers(string circuitId) view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SignatureVerifiers(circuitId string) (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.SignatureVerifiers(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// SignatureVerifiers is a free data retrieval call binding the contract method 0xcde1cacb.
//
// Solidity: function signatureVerifiers(string circuitId) view returns(address)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) SignatureVerifiers(circuitId string) (common.Address, error) {
	return _PassportCredentialIssuerImplV1.Contract.SignatureVerifiers(&_PassportCredentialIssuerImplV1.CallOpts, circuitId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PassportCredentialIssuerImplV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PassportCredentialIssuerImplV1.Contract.SupportsInterface(&_PassportCredentialIssuerImplV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PassportCredentialIssuerImplV1.Contract.SupportsInterface(&_PassportCredentialIssuerImplV1.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) AcceptOwnership() (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.AcceptOwnership(&_PassportCredentialIssuerImplV1.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.AcceptOwnership(&_PassportCredentialIssuerImplV1.TransactOpts)
}

// CleanNullifier is a paid mutator transaction binding the contract method 0xd94b0271.
//
// Solidity: function cleanNullifier(uint256 nullifier) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) CleanNullifier(opts *bind.TransactOpts, nullifier *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "cleanNullifier", nullifier)
}

// CleanNullifier is a paid mutator transaction binding the contract method 0xd94b0271.
//
// Solidity: function cleanNullifier(uint256 nullifier) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) CleanNullifier(nullifier *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.CleanNullifier(&_PassportCredentialIssuerImplV1.TransactOpts, nullifier)
}

// CleanNullifier is a paid mutator transaction binding the contract method 0xd94b0271.
//
// Solidity: function cleanNullifier(uint256 nullifier) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) CleanNullifier(nullifier *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.CleanNullifier(&_PassportCredentialIssuerImplV1.TransactOpts, nullifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x406e8b4c.
//
// Solidity: function initialize(address _stateContractAddr, bytes2 idType) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) Initialize(opts *bind.TransactOpts, _stateContractAddr common.Address, idType [2]byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "initialize", _stateContractAddr, idType)
}

// Initialize is a paid mutator transaction binding the contract method 0x406e8b4c.
//
// Solidity: function initialize(address _stateContractAddr, bytes2 idType) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) Initialize(_stateContractAddr common.Address, idType [2]byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.Initialize(&_PassportCredentialIssuerImplV1.TransactOpts, _stateContractAddr, idType)
}

// Initialize is a paid mutator transaction binding the contract method 0x406e8b4c.
//
// Solidity: function initialize(address _stateContractAddr, bytes2 idType) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) Initialize(_stateContractAddr common.Address, idType [2]byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.Initialize(&_PassportCredentialIssuerImplV1.TransactOpts, _stateContractAddr, idType)
}

// InitializeIssuer is a paid mutator transaction binding the contract method 0x09bdc4b1.
//
// Solidity: function initializeIssuer(uint256 expirationTime, uint256 templateRoot, address registry, string[] credentialCircuitIds, address[] credentialVerifierAddresses, string[] signatureCircuitIds, address[] signatureVerifierAddresses, address stateAddress, bytes2 idType) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) InitializeIssuer(opts *bind.TransactOpts, expirationTime *big.Int, templateRoot *big.Int, registry common.Address, credentialCircuitIds []string, credentialVerifierAddresses []common.Address, signatureCircuitIds []string, signatureVerifierAddresses []common.Address, stateAddress common.Address, idType [2]byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "initializeIssuer", expirationTime, templateRoot, registry, credentialCircuitIds, credentialVerifierAddresses, signatureCircuitIds, signatureVerifierAddresses, stateAddress, idType)
}

// InitializeIssuer is a paid mutator transaction binding the contract method 0x09bdc4b1.
//
// Solidity: function initializeIssuer(uint256 expirationTime, uint256 templateRoot, address registry, string[] credentialCircuitIds, address[] credentialVerifierAddresses, string[] signatureCircuitIds, address[] signatureVerifierAddresses, address stateAddress, bytes2 idType) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) InitializeIssuer(expirationTime *big.Int, templateRoot *big.Int, registry common.Address, credentialCircuitIds []string, credentialVerifierAddresses []common.Address, signatureCircuitIds []string, signatureVerifierAddresses []common.Address, stateAddress common.Address, idType [2]byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.InitializeIssuer(&_PassportCredentialIssuerImplV1.TransactOpts, expirationTime, templateRoot, registry, credentialCircuitIds, credentialVerifierAddresses, signatureCircuitIds, signatureVerifierAddresses, stateAddress, idType)
}

// InitializeIssuer is a paid mutator transaction binding the contract method 0x09bdc4b1.
//
// Solidity: function initializeIssuer(uint256 expirationTime, uint256 templateRoot, address registry, string[] credentialCircuitIds, address[] credentialVerifierAddresses, string[] signatureCircuitIds, address[] signatureVerifierAddresses, address stateAddress, bytes2 idType) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) InitializeIssuer(expirationTime *big.Int, templateRoot *big.Int, registry common.Address, credentialCircuitIds []string, credentialVerifierAddresses []common.Address, signatureCircuitIds []string, signatureVerifierAddresses []common.Address, stateAddress common.Address, idType [2]byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.InitializeIssuer(&_PassportCredentialIssuerImplV1.TransactOpts, expirationTime, templateRoot, registry, credentialCircuitIds, credentialVerifierAddresses, signatureCircuitIds, signatureVerifierAddresses, stateAddress, idType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.RenounceOwnership(&_PassportCredentialIssuerImplV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.RenounceOwnership(&_PassportCredentialIssuerImplV1.TransactOpts)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 expirationTime) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) SetExpirationTime(opts *bind.TransactOpts, expirationTime *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "setExpirationTime", expirationTime)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 expirationTime) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SetExpirationTime(expirationTime *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SetExpirationTime(&_PassportCredentialIssuerImplV1.TransactOpts, expirationTime)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 expirationTime) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) SetExpirationTime(expirationTime *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SetExpirationTime(&_PassportCredentialIssuerImplV1.TransactOpts, expirationTime)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registry) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) SetRegistry(opts *bind.TransactOpts, registry common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "setRegistry", registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registry) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SetRegistry(registry common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SetRegistry(&_PassportCredentialIssuerImplV1.TransactOpts, registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registry) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) SetRegistry(registry common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SetRegistry(&_PassportCredentialIssuerImplV1.TransactOpts, registry)
}

// SetTemplateRoot is a paid mutator transaction binding the contract method 0xb354c2d8.
//
// Solidity: function setTemplateRoot(uint256 templateRoot) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) SetTemplateRoot(opts *bind.TransactOpts, templateRoot *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "setTemplateRoot", templateRoot)
}

// SetTemplateRoot is a paid mutator transaction binding the contract method 0xb354c2d8.
//
// Solidity: function setTemplateRoot(uint256 templateRoot) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SetTemplateRoot(templateRoot *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SetTemplateRoot(&_PassportCredentialIssuerImplV1.TransactOpts, templateRoot)
}

// SetTemplateRoot is a paid mutator transaction binding the contract method 0xb354c2d8.
//
// Solidity: function setTemplateRoot(uint256 templateRoot) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) SetTemplateRoot(templateRoot *big.Int) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SetTemplateRoot(&_PassportCredentialIssuerImplV1.TransactOpts, templateRoot)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProofs) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) SubmitZKPResponseV2(opts *bind.TransactOpts, responses []IZKPVerifierZKPResponse, crossChainProofs []byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "submitZKPResponseV2", responses, crossChainProofs)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProofs) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) SubmitZKPResponseV2(responses []IZKPVerifierZKPResponse, crossChainProofs []byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SubmitZKPResponseV2(&_PassportCredentialIssuerImplV1.TransactOpts, responses, crossChainProofs)
}

// SubmitZKPResponseV2 is a paid mutator transaction binding the contract method 0xade09fcd.
//
// Solidity: function submitZKPResponseV2((uint64,bytes,bytes)[] responses, bytes crossChainProofs) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) SubmitZKPResponseV2(responses []IZKPVerifierZKPResponse, crossChainProofs []byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.SubmitZKPResponseV2(&_PassportCredentialIssuerImplV1.TransactOpts, responses, crossChainProofs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.TransferOwnership(&_PassportCredentialIssuerImplV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.TransferOwnership(&_PassportCredentialIssuerImplV1.TransactOpts, newOwner)
}

// UpdateCredentialVerifiers is a paid mutator transaction binding the contract method 0xbadf055d.
//
// Solidity: function updateCredentialVerifiers(string[] circuitIds, address[] verifierAddresses) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) UpdateCredentialVerifiers(opts *bind.TransactOpts, circuitIds []string, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "updateCredentialVerifiers", circuitIds, verifierAddresses)
}

// UpdateCredentialVerifiers is a paid mutator transaction binding the contract method 0xbadf055d.
//
// Solidity: function updateCredentialVerifiers(string[] circuitIds, address[] verifierAddresses) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) UpdateCredentialVerifiers(circuitIds []string, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.UpdateCredentialVerifiers(&_PassportCredentialIssuerImplV1.TransactOpts, circuitIds, verifierAddresses)
}

// UpdateCredentialVerifiers is a paid mutator transaction binding the contract method 0xbadf055d.
//
// Solidity: function updateCredentialVerifiers(string[] circuitIds, address[] verifierAddresses) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) UpdateCredentialVerifiers(circuitIds []string, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.UpdateCredentialVerifiers(&_PassportCredentialIssuerImplV1.TransactOpts, circuitIds, verifierAddresses)
}

// UpdateSignatureVerifiers is a paid mutator transaction binding the contract method 0x9dc3c7ba.
//
// Solidity: function updateSignatureVerifiers(string[] circuitIds, address[] verifierAddresses) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) UpdateSignatureVerifiers(opts *bind.TransactOpts, circuitIds []string, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "updateSignatureVerifiers", circuitIds, verifierAddresses)
}

// UpdateSignatureVerifiers is a paid mutator transaction binding the contract method 0x9dc3c7ba.
//
// Solidity: function updateSignatureVerifiers(string[] circuitIds, address[] verifierAddresses) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) UpdateSignatureVerifiers(circuitIds []string, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.UpdateSignatureVerifiers(&_PassportCredentialIssuerImplV1.TransactOpts, circuitIds, verifierAddresses)
}

// UpdateSignatureVerifiers is a paid mutator transaction binding the contract method 0x9dc3c7ba.
//
// Solidity: function updateSignatureVerifiers(string[] circuitIds, address[] verifierAddresses) returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) UpdateSignatureVerifiers(circuitIds []string, verifierAddresses []common.Address) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.UpdateSignatureVerifiers(&_PassportCredentialIssuerImplV1.TransactOpts, circuitIds, verifierAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.UpgradeToAndCall(&_PassportCredentialIssuerImplV1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PassportCredentialIssuerImplV1.Contract.UpgradeToAndCall(&_PassportCredentialIssuerImplV1.TransactOpts, newImplementation, data)
}

// PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator is returned from FilterCredentialCircuitVerifierUpdated and is used to iterate over the raw logs and unpacked data for CredentialCircuitVerifierUpdated events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator struct {
	Event *PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated)
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
		it.Event = new(PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated)
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
func (it *PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated represents a CredentialCircuitVerifierUpdated event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated struct {
	CircuitId string
	Verifier  common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCredentialCircuitVerifierUpdated is a free log retrieval operation binding the contract event 0x17b20bbf8df0ff1422db2abe4b28c298e32929a2d112f027c3f68d488cf3cf74.
//
// Solidity: event CredentialCircuitVerifierUpdated(string circuitId, address verifier, uint256 requestId)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterCredentialCircuitVerifierUpdated(opts *bind.FilterOpts) (*PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "CredentialCircuitVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdatedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "CredentialCircuitVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchCredentialCircuitVerifierUpdated is a free log subscription operation binding the contract event 0x17b20bbf8df0ff1422db2abe4b28c298e32929a2d112f027c3f68d488cf3cf74.
//
// Solidity: event CredentialCircuitVerifierUpdated(string circuitId, address verifier, uint256 requestId)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchCredentialCircuitVerifierUpdated(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "CredentialCircuitVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "CredentialCircuitVerifierUpdated", log); err != nil {
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

// ParseCredentialCircuitVerifierUpdated is a log parse operation binding the contract event 0x17b20bbf8df0ff1422db2abe4b28c298e32929a2d112f027c3f68d488cf3cf74.
//
// Solidity: event CredentialCircuitVerifierUpdated(string circuitId, address verifier, uint256 requestId)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseCredentialCircuitVerifierUpdated(log types.Log) (*PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated, error) {
	event := new(PassportCredentialIssuerImplV1CredentialCircuitVerifierUpdated)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "CredentialCircuitVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator is returned from FilterExpirationTimeUpdated and is used to iterate over the raw logs and unpacked data for ExpirationTimeUpdated events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator struct {
	Event *PassportCredentialIssuerImplV1ExpirationTimeUpdated // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1ExpirationTimeUpdated)
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
		it.Event = new(PassportCredentialIssuerImplV1ExpirationTimeUpdated)
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
func (it *PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1ExpirationTimeUpdated represents a ExpirationTimeUpdated event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1ExpirationTimeUpdated struct {
	ExpirationTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExpirationTimeUpdated is a free log retrieval operation binding the contract event 0x7f7f3071c3c286700c199acae2fe071beb4a63dae1afff85186fa75d18090ae8.
//
// Solidity: event ExpirationTimeUpdated(uint256 expirationTime)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterExpirationTimeUpdated(opts *bind.FilterOpts) (*PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "ExpirationTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1ExpirationTimeUpdatedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "ExpirationTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchExpirationTimeUpdated is a free log subscription operation binding the contract event 0x7f7f3071c3c286700c199acae2fe071beb4a63dae1afff85186fa75d18090ae8.
//
// Solidity: event ExpirationTimeUpdated(uint256 expirationTime)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchExpirationTimeUpdated(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1ExpirationTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "ExpirationTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1ExpirationTimeUpdated)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "ExpirationTimeUpdated", log); err != nil {
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

// ParseExpirationTimeUpdated is a log parse operation binding the contract event 0x7f7f3071c3c286700c199acae2fe071beb4a63dae1afff85186fa75d18090ae8.
//
// Solidity: event ExpirationTimeUpdated(uint256 expirationTime)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseExpirationTimeUpdated(log types.Log) (*PassportCredentialIssuerImplV1ExpirationTimeUpdated, error) {
	event := new(PassportCredentialIssuerImplV1ExpirationTimeUpdated)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "ExpirationTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1InitializedIterator struct {
	Event *PassportCredentialIssuerImplV1Initialized // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1Initialized)
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
		it.Event = new(PassportCredentialIssuerImplV1Initialized)
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
func (it *PassportCredentialIssuerImplV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1Initialized represents a Initialized event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*PassportCredentialIssuerImplV1InitializedIterator, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1InitializedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1Initialized) (event.Subscription, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1Initialized)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseInitialized(log types.Log) (*PassportCredentialIssuerImplV1Initialized, error) {
	event := new(PassportCredentialIssuerImplV1Initialized)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1OwnershipTransferStartedIterator struct {
	Event *PassportCredentialIssuerImplV1OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1OwnershipTransferStarted)
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
		it.Event = new(PassportCredentialIssuerImplV1OwnershipTransferStarted)
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
func (it *PassportCredentialIssuerImplV1OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportCredentialIssuerImplV1OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1OwnershipTransferStartedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1OwnershipTransferStarted)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseOwnershipTransferStarted(log types.Log) (*PassportCredentialIssuerImplV1OwnershipTransferStarted, error) {
	event := new(PassportCredentialIssuerImplV1OwnershipTransferStarted)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1OwnershipTransferredIterator struct {
	Event *PassportCredentialIssuerImplV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1OwnershipTransferred)
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
		it.Event = new(PassportCredentialIssuerImplV1OwnershipTransferred)
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
func (it *PassportCredentialIssuerImplV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1OwnershipTransferred represents a OwnershipTransferred event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PassportCredentialIssuerImplV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1OwnershipTransferredIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1OwnershipTransferred)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseOwnershipTransferred(log types.Log) (*PassportCredentialIssuerImplV1OwnershipTransferred, error) {
	event := new(PassportCredentialIssuerImplV1OwnershipTransferred)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1RegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1RegistryUpdatedIterator struct {
	Event *PassportCredentialIssuerImplV1RegistryUpdated // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1RegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1RegistryUpdated)
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
		it.Event = new(PassportCredentialIssuerImplV1RegistryUpdated)
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
func (it *PassportCredentialIssuerImplV1RegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1RegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1RegistryUpdated represents a RegistryUpdated event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1RegistryUpdated struct {
	Registry common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0xd6ceddf6d2a22f21c7c81675c518004eff43bc5c8a6fc32a0b748e69d58671cd.
//
// Solidity: event RegistryUpdated(address registry)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*PassportCredentialIssuerImplV1RegistryUpdatedIterator, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1RegistryUpdatedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0xd6ceddf6d2a22f21c7c81675c518004eff43bc5c8a6fc32a0b748e69d58671cd.
//
// Solidity: event RegistryUpdated(address registry)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1RegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1RegistryUpdated)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseRegistryUpdated(log types.Log) (*PassportCredentialIssuerImplV1RegistryUpdated, error) {
	event := new(PassportCredentialIssuerImplV1RegistryUpdated)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator is returned from FilterSignatureCircuitVerifierUpdated and is used to iterate over the raw logs and unpacked data for SignatureCircuitVerifierUpdated events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator struct {
	Event *PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated)
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
		it.Event = new(PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated)
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
func (it *PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated represents a SignatureCircuitVerifierUpdated event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated struct {
	CircuitId string
	Verifier  common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignatureCircuitVerifierUpdated is a free log retrieval operation binding the contract event 0x73293df9496ecb6cf8b93101e8333c4e3add8400c974d542a34600d304c92244.
//
// Solidity: event SignatureCircuitVerifierUpdated(string circuitId, address verifier, uint256 requestId)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterSignatureCircuitVerifierUpdated(opts *bind.FilterOpts) (*PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "SignatureCircuitVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdatedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "SignatureCircuitVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchSignatureCircuitVerifierUpdated is a free log subscription operation binding the contract event 0x73293df9496ecb6cf8b93101e8333c4e3add8400c974d542a34600d304c92244.
//
// Solidity: event SignatureCircuitVerifierUpdated(string circuitId, address verifier, uint256 requestId)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchSignatureCircuitVerifierUpdated(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "SignatureCircuitVerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "SignatureCircuitVerifierUpdated", log); err != nil {
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

// ParseSignatureCircuitVerifierUpdated is a log parse operation binding the contract event 0x73293df9496ecb6cf8b93101e8333c4e3add8400c974d542a34600d304c92244.
//
// Solidity: event SignatureCircuitVerifierUpdated(string circuitId, address verifier, uint256 requestId)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseSignatureCircuitVerifierUpdated(log types.Log) (*PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated, error) {
	event := new(PassportCredentialIssuerImplV1SignatureCircuitVerifierUpdated)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "SignatureCircuitVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1TemplateRootUpdatedIterator is returned from FilterTemplateRootUpdated and is used to iterate over the raw logs and unpacked data for TemplateRootUpdated events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1TemplateRootUpdatedIterator struct {
	Event *PassportCredentialIssuerImplV1TemplateRootUpdated // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1TemplateRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1TemplateRootUpdated)
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
		it.Event = new(PassportCredentialIssuerImplV1TemplateRootUpdated)
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
func (it *PassportCredentialIssuerImplV1TemplateRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1TemplateRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1TemplateRootUpdated represents a TemplateRootUpdated event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1TemplateRootUpdated struct {
	TampletRoot *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTemplateRootUpdated is a free log retrieval operation binding the contract event 0xf007a02ba428e63fa98b65269e70f9a9db847e0a839049a2f31f069282267734.
//
// Solidity: event TemplateRootUpdated(uint256 tampletRoot)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterTemplateRootUpdated(opts *bind.FilterOpts) (*PassportCredentialIssuerImplV1TemplateRootUpdatedIterator, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "TemplateRootUpdated")
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1TemplateRootUpdatedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "TemplateRootUpdated", logs: logs, sub: sub}, nil
}

// WatchTemplateRootUpdated is a free log subscription operation binding the contract event 0xf007a02ba428e63fa98b65269e70f9a9db847e0a839049a2f31f069282267734.
//
// Solidity: event TemplateRootUpdated(uint256 tampletRoot)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchTemplateRootUpdated(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1TemplateRootUpdated) (event.Subscription, error) {

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "TemplateRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1TemplateRootUpdated)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "TemplateRootUpdated", log); err != nil {
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

// ParseTemplateRootUpdated is a log parse operation binding the contract event 0xf007a02ba428e63fa98b65269e70f9a9db847e0a839049a2f31f069282267734.
//
// Solidity: event TemplateRootUpdated(uint256 tampletRoot)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseTemplateRootUpdated(log types.Log) (*PassportCredentialIssuerImplV1TemplateRootUpdated, error) {
	event := new(PassportCredentialIssuerImplV1TemplateRootUpdated)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "TemplateRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PassportCredentialIssuerImplV1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1UpgradedIterator struct {
	Event *PassportCredentialIssuerImplV1Upgraded // Event containing the contract specifics and raw log

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
func (it *PassportCredentialIssuerImplV1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PassportCredentialIssuerImplV1Upgraded)
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
		it.Event = new(PassportCredentialIssuerImplV1Upgraded)
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
func (it *PassportCredentialIssuerImplV1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PassportCredentialIssuerImplV1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PassportCredentialIssuerImplV1Upgraded represents a Upgraded event raised by the PassportCredentialIssuerImplV1 contract.
type PassportCredentialIssuerImplV1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PassportCredentialIssuerImplV1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PassportCredentialIssuerImplV1UpgradedIterator{contract: _PassportCredentialIssuerImplV1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PassportCredentialIssuerImplV1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PassportCredentialIssuerImplV1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PassportCredentialIssuerImplV1Upgraded)
				if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PassportCredentialIssuerImplV1 *PassportCredentialIssuerImplV1Filterer) ParseUpgraded(log types.Log) (*PassportCredentialIssuerImplV1Upgraded, error) {
	event := new(PassportCredentialIssuerImplV1Upgraded)
	if err := _PassportCredentialIssuerImplV1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
