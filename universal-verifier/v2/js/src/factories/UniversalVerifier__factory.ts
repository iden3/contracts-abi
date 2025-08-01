/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  UniversalVerifier,
  UniversalVerifierInterface,
} from "../UniversalVerifier";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "AuthMethodAlreadyExists",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "AuthMethodIsNotActive",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "AuthMethodNotFound",
    type: "error",
  },
  {
    inputs: [],
    name: "ChallengeIsInvalid",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "length",
        type: "uint256",
      },
    ],
    name: "ChecksumLengthRequired",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "groupId",
        type: "uint256",
      },
    ],
    name: "GroupIdAlreadyExists",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "groupId",
        type: "uint256",
      },
    ],
    name: "GroupIdNotFound",
    type: "error",
  },
  {
    inputs: [],
    name: "GroupIdNotValid",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "groupId",
        type: "uint256",
      },
    ],
    name: "GroupMustHaveAtLeastTwoRequests",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "length",
        type: "uint256",
      },
    ],
    name: "IdBytesLengthRequired",
    type: "error",
  },
  {
    inputs: [],
    name: "InvalidInitialization",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "requestOwner",
        type: "address",
      },
      {
        internalType: "address",
        name: "sender",
        type: "address",
      },
    ],
    name: "InvalidRequestOwner",
    type: "error",
  },
  {
    inputs: [],
    name: "LinkIDNotTheSameForGroupedRequests",
    type: "error",
  },
  {
    inputs: [],
    name: "MetadataNotSupportedYet",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "groupId",
        type: "uint256",
      },
    ],
    name: "MissingUserIDInGroupOfRequests",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "MissingUserIDInRequest",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
    ],
    name: "MultiRequestIdAlreadyExists",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
    ],
    name: "MultiRequestIdNotFound",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "expectedMultiRequestId",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
    ],
    name: "MultiRequestIdNotValid",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "NotAnOwnerOrRequestOwner",
    type: "error",
  },
  {
    inputs: [],
    name: "NotInitializing",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "nullifierSessionID",
        type: "uint256",
      },
    ],
    name: "NullifierSessionIDAlreadyExists",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
    ],
    name: "OwnableInvalidOwner",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "account",
        type: "address",
      },
    ],
    name: "OwnableUnauthorizedAccount",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "RequestIdAlreadyExists",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "RequestIdNotFound",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "expectedRequestId",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "RequestIdNotValid",
    type: "error",
  },
  {
    inputs: [],
    name: "RequestIdTypeNotValid",
    type: "error",
  },
  {
    inputs: [],
    name: "RequestIdUsesReservedBytes",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "RequestIsDisabled",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "RequestShouldNotHaveAGroup",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "responseFieldName",
        type: "string",
      },
    ],
    name: "ResponseFieldAlreadyExists",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "userIDFromAuth",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "userIDFromResponse",
        type: "uint256",
      },
    ],
    name: "UserIDMismatch",
    type: "error",
  },
  {
    inputs: [],
    name: "UserNotAuthenticated",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "validator",
        type: "address",
      },
    ],
    name: "ValidatorIsNotWhitelisted",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "validator",
        type: "address",
      },
    ],
    name: "ValidatorNotSupportInterface",
    type: "error",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestVerifierID",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "expectedVerifierID",
        type: "uint256",
      },
    ],
    name: "VerifierIDIsNotValid",
    type: "error",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
      {
        indexed: false,
        internalType: "address",
        name: "validator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "params",
        type: "bytes",
      },
    ],
    name: "AuthMethodSet",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
      {
        indexed: true,
        internalType: "address",
        name: "caller",
        type: "address",
      },
    ],
    name: "AuthResponseSubmitted",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "uint64",
        name: "version",
        type: "uint64",
      },
    ],
    name: "Initialized",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
      {
        indexed: false,
        internalType: "uint256[]",
        name: "requestIds",
        type: "uint256[]",
      },
      {
        indexed: false,
        internalType: "uint256[]",
        name: "groupIds",
        type: "uint256[]",
      },
    ],
    name: "MultiRequestSet",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferStarted",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
      {
        indexed: true,
        internalType: "address",
        name: "creator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "string",
        name: "metadata",
        type: "string",
      },
      {
        indexed: false,
        internalType: "address",
        name: "validator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "params",
        type: "bytes",
      },
    ],
    name: "RequestSet",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
      {
        indexed: true,
        internalType: "address",
        name: "creator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "string",
        name: "metadata",
        type: "string",
      },
      {
        indexed: false,
        internalType: "address",
        name: "validator",
        type: "address",
      },
      {
        indexed: false,
        internalType: "bytes",
        name: "params",
        type: "bytes",
      },
    ],
    name: "RequestUpdate",
    type: "event",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
      {
        indexed: true,
        internalType: "address",
        name: "caller",
        type: "address",
      },
    ],
    name: "ResponseSubmitted",
    type: "event",
  },
  {
    inputs: [],
    name: "VERSION",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "acceptOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "contract IRequestValidator",
        name: "validator",
        type: "address",
      },
    ],
    name: "addValidatorToWhitelist",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "userAddress",
        type: "address",
      },
    ],
    name: "areMultiRequestProofsVerified",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "authMethodExists",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "disableAuthMethod",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "disableRequest",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "enableAuthMethod",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "enableRequest",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "authMethod",
        type: "string",
      },
    ],
    name: "getAuthMethod",
    outputs: [
      {
        components: [
          {
            internalType: "contract IAuthValidator",
            name: "validator",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "params",
            type: "bytes",
          },
          {
            internalType: "bool",
            name: "isActive",
            type: "bool",
          },
        ],
        internalType: "struct Verifier.AuthMethodData",
        name: "authMethodData",
        type: "tuple",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "groupID",
        type: "uint256",
      },
    ],
    name: "getGroupedRequests",
    outputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "string",
            name: "metadata",
            type: "string",
          },
          {
            internalType: "contract IRequestValidator",
            name: "validator",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "params",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "creator",
            type: "address",
          },
        ],
        internalType: "struct IVerifier.RequestInfo[]",
        name: "",
        type: "tuple[]",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "getGroupsCount",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
    ],
    name: "getMultiRequest",
    outputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "multiRequestId",
            type: "uint256",
          },
          {
            internalType: "uint256[]",
            name: "requestIds",
            type: "uint256[]",
          },
          {
            internalType: "uint256[]",
            name: "groupIds",
            type: "uint256[]",
          },
          {
            internalType: "bytes",
            name: "metadata",
            type: "bytes",
          },
        ],
        internalType: "struct IVerifier.MultiRequest",
        name: "multiRequest",
        type: "tuple",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "userAddress",
        type: "address",
      },
    ],
    name: "getMultiRequestProofsStatus",
    outputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "bool",
            name: "isVerified",
            type: "bool",
          },
          {
            internalType: "string",
            name: "validatorVersion",
            type: "string",
          },
          {
            internalType: "uint256",
            name: "timestamp",
            type: "uint256",
          },
        ],
        internalType: "struct IVerifier.RequestProofStatus[]",
        name: "",
        type: "tuple[]",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "getRequest",
    outputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "string",
            name: "metadata",
            type: "string",
          },
          {
            internalType: "contract IRequestValidator",
            name: "validator",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "params",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "creator",
            type: "address",
          },
        ],
        internalType: "struct IVerifier.RequestInfo",
        name: "request",
        type: "tuple",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "getRequestOwner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "sender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "getRequestProofStatus",
    outputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "bool",
            name: "isVerified",
            type: "bool",
          },
          {
            internalType: "string",
            name: "validatorVersion",
            type: "string",
          },
          {
            internalType: "uint256",
            name: "timestamp",
            type: "uint256",
          },
        ],
        internalType: "struct IVerifier.RequestProofStatus",
        name: "",
        type: "tuple",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "getRequestsCount",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "sender",
        type: "address",
      },
      {
        internalType: "string",
        name: "responseFieldName",
        type: "string",
      },
    ],
    name: "getResponseFieldValue",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "sender",
        type: "address",
      },
    ],
    name: "getResponseFields",
    outputs: [
      {
        components: [
          {
            internalType: "string",
            name: "name",
            type: "string",
          },
          {
            internalType: "uint256",
            name: "value",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "rawValue",
            type: "bytes",
          },
        ],
        internalType: "struct IRequestValidator.ResponseField[]",
        name: "",
        type: "tuple[]",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "getStateAddress",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "getVerifierID",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "groupId",
        type: "uint256",
      },
    ],
    name: "groupIdExists",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "contract IState",
        name: "state",
        type: "address",
      },
      {
        internalType: "address",
        name: "owner",
        type: "address",
      },
    ],
    name: "initialize",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "isRequestEnabled",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "sender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "isRequestProofVerified",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "contract IRequestValidator",
        name: "validator",
        type: "address",
      },
    ],
    name: "isWhitelistedValidator",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "multiRequestId",
        type: "uint256",
      },
    ],
    name: "multiRequestIdExists",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "pendingOwner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "contract IRequestValidator",
        name: "validator",
        type: "address",
      },
    ],
    name: "removeValidatorFromWhitelist",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
    ],
    name: "requestIdExists",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "string",
            name: "authMethod",
            type: "string",
          },
          {
            internalType: "contract IAuthValidator",
            name: "validator",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "params",
            type: "bytes",
          },
        ],
        internalType: "struct IVerifier.AuthMethod",
        name: "authMethod",
        type: "tuple",
      },
    ],
    name: "setAuthMethod",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "multiRequestId",
            type: "uint256",
          },
          {
            internalType: "uint256[]",
            name: "requestIds",
            type: "uint256[]",
          },
          {
            internalType: "uint256[]",
            name: "groupIds",
            type: "uint256[]",
          },
          {
            internalType: "bytes",
            name: "metadata",
            type: "bytes",
          },
        ],
        internalType: "struct IVerifier.MultiRequest",
        name: "multiRequest",
        type: "tuple",
      },
    ],
    name: "setMultiRequest",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "requestId",
        type: "uint256",
      },
      {
        internalType: "address",
        name: "requestOwner",
        type: "address",
      },
    ],
    name: "setRequestOwner",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "string",
            name: "metadata",
            type: "string",
          },
          {
            internalType: "contract IRequestValidator",
            name: "validator",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "params",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "creator",
            type: "address",
          },
        ],
        internalType: "struct IVerifier.Request[]",
        name: "requests",
        type: "tuple[]",
      },
    ],
    name: "setRequests",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "contract IState",
        name: "state",
        type: "address",
      },
    ],
    name: "setState",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "verifierID",
        type: "uint256",
      },
    ],
    name: "setVerifierID",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "string",
            name: "authMethod",
            type: "string",
          },
          {
            internalType: "bytes",
            name: "proof",
            type: "bytes",
          },
        ],
        internalType: "struct IVerifier.AuthResponse",
        name: "authResponse",
        type: "tuple",
      },
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "bytes",
            name: "proof",
            type: "bytes",
          },
          {
            internalType: "bytes",
            name: "metadata",
            type: "bytes",
          },
        ],
        internalType: "struct IVerifier.Response[]",
        name: "responses",
        type: "tuple[]",
      },
      {
        internalType: "bytes",
        name: "crossChainProofs",
        type: "bytes",
      },
    ],
    name: "submitResponse",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "uint256",
            name: "requestId",
            type: "uint256",
          },
          {
            internalType: "string",
            name: "metadata",
            type: "string",
          },
          {
            internalType: "contract IRequestValidator",
            name: "validator",
            type: "address",
          },
          {
            internalType: "bytes",
            name: "params",
            type: "bytes",
          },
          {
            internalType: "address",
            name: "creator",
            type: "address",
          },
        ],
        internalType: "struct IVerifier.Request",
        name: "request",
        type: "tuple",
      },
    ],
    name: "updateRequest",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "version",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "pure",
    type: "function",
  },
] as const;

export class UniversalVerifier__factory {
  static readonly abi = _abi;
  static createInterface(): UniversalVerifierInterface {
    return new Interface(_abi) as UniversalVerifierInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): UniversalVerifier {
    return new Contract(address, _abi, runner) as unknown as UniversalVerifier;
  }
}
