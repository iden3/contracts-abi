/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "./common";

export declare namespace Verifier {
  export type AuthMethodDataStruct = {
    validator: AddressLike;
    params: BytesLike;
    isActive: boolean;
  };

  export type AuthMethodDataStructOutput = [
    validator: string,
    params: string,
    isActive: boolean
  ] & { validator: string; params: string; isActive: boolean };
}

export declare namespace IVerifier {
  export type RequestInfoStruct = {
    requestId: BigNumberish;
    metadata: string;
    validator: AddressLike;
    params: BytesLike;
    creator: AddressLike;
  };

  export type RequestInfoStructOutput = [
    requestId: bigint,
    metadata: string,
    validator: string,
    params: string,
    creator: string
  ] & {
    requestId: bigint;
    metadata: string;
    validator: string;
    params: string;
    creator: string;
  };

  export type MultiRequestStruct = {
    multiRequestId: BigNumberish;
    requestIds: BigNumberish[];
    groupIds: BigNumberish[];
    metadata: BytesLike;
  };

  export type MultiRequestStructOutput = [
    multiRequestId: bigint,
    requestIds: bigint[],
    groupIds: bigint[],
    metadata: string
  ] & {
    multiRequestId: bigint;
    requestIds: bigint[];
    groupIds: bigint[];
    metadata: string;
  };

  export type RequestProofStatusStruct = {
    requestId: BigNumberish;
    isVerified: boolean;
    validatorVersion: string;
    timestamp: BigNumberish;
  };

  export type RequestProofStatusStructOutput = [
    requestId: bigint,
    isVerified: boolean,
    validatorVersion: string,
    timestamp: bigint
  ] & {
    requestId: bigint;
    isVerified: boolean;
    validatorVersion: string;
    timestamp: bigint;
  };

  export type AuthMethodStruct = {
    authMethod: string;
    validator: AddressLike;
    params: BytesLike;
  };

  export type AuthMethodStructOutput = [
    authMethod: string,
    validator: string,
    params: string
  ] & { authMethod: string; validator: string; params: string };

  export type RequestStruct = {
    requestId: BigNumberish;
    metadata: string;
    validator: AddressLike;
    params: BytesLike;
    creator: AddressLike;
  };

  export type RequestStructOutput = [
    requestId: bigint,
    metadata: string,
    validator: string,
    params: string,
    creator: string
  ] & {
    requestId: bigint;
    metadata: string;
    validator: string;
    params: string;
    creator: string;
  };

  export type AuthResponseStruct = { authMethod: string; proof: BytesLike };

  export type AuthResponseStructOutput = [authMethod: string, proof: string] & {
    authMethod: string;
    proof: string;
  };

  export type ResponseStruct = {
    requestId: BigNumberish;
    proof: BytesLike;
    metadata: BytesLike;
  };

  export type ResponseStructOutput = [
    requestId: bigint,
    proof: string,
    metadata: string
  ] & { requestId: bigint; proof: string; metadata: string };
}

export declare namespace IRequestValidator {
  export type ResponseFieldStruct = {
    name: string;
    value: BigNumberish;
    rawValue: BytesLike;
  };

  export type ResponseFieldStructOutput = [
    name: string,
    value: bigint,
    rawValue: string
  ] & { name: string; value: bigint; rawValue: string };
}

export interface UniversalVerifierInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "VERSION"
      | "acceptOwnership"
      | "addValidatorToWhitelist"
      | "areMultiRequestProofsVerified"
      | "authMethodExists"
      | "disableAuthMethod"
      | "disableRequest"
      | "enableAuthMethod"
      | "enableRequest"
      | "getAuthMethod"
      | "getGroupedRequests"
      | "getGroupsCount"
      | "getMultiRequest"
      | "getMultiRequestProofsStatus"
      | "getRequest"
      | "getRequestOwner"
      | "getRequestProofStatus"
      | "getRequestsCount"
      | "getResponseFieldValue"
      | "getResponseFields"
      | "getStateAddress"
      | "getVerifierID"
      | "groupIdExists"
      | "initialize"
      | "isRequestEnabled"
      | "isRequestProofVerified"
      | "isWhitelistedValidator"
      | "multiRequestIdExists"
      | "owner"
      | "pendingOwner"
      | "removeValidatorFromWhitelist"
      | "renounceOwnership"
      | "requestIdExists"
      | "setAuthMethod"
      | "setMultiRequest"
      | "setRequestOwner"
      | "setRequests"
      | "setState"
      | "setVerifierID"
      | "submitResponse"
      | "transferOwnership"
      | "updateRequest"
      | "version"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "AuthMethodSet"
      | "AuthResponseSubmitted"
      | "Initialized"
      | "MultiRequestSet"
      | "OwnershipTransferStarted"
      | "OwnershipTransferred"
      | "RequestSet"
      | "RequestUpdate"
      | "ResponseSubmitted"
  ): EventFragment;

  encodeFunctionData(functionFragment: "VERSION", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "acceptOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "addValidatorToWhitelist",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "areMultiRequestProofsVerified",
    values: [BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "authMethodExists",
    values: [string]
  ): string;
  encodeFunctionData(
    functionFragment: "disableAuthMethod",
    values: [string]
  ): string;
  encodeFunctionData(
    functionFragment: "disableRequest",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "enableAuthMethod",
    values: [string]
  ): string;
  encodeFunctionData(
    functionFragment: "enableRequest",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getAuthMethod",
    values: [string]
  ): string;
  encodeFunctionData(
    functionFragment: "getGroupedRequests",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getGroupsCount",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getMultiRequest",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getMultiRequestProofsStatus",
    values: [BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getRequest",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getRequestOwner",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getRequestProofStatus",
    values: [AddressLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getRequestsCount",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getResponseFieldValue",
    values: [BigNumberish, AddressLike, string]
  ): string;
  encodeFunctionData(
    functionFragment: "getResponseFields",
    values: [BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getStateAddress",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getVerifierID",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "groupIdExists",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [AddressLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "isRequestEnabled",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "isRequestProofVerified",
    values: [AddressLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "isWhitelistedValidator",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "multiRequestIdExists",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "pendingOwner",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "removeValidatorFromWhitelist",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "requestIdExists",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "setAuthMethod",
    values: [IVerifier.AuthMethodStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "setMultiRequest",
    values: [IVerifier.MultiRequestStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "setRequestOwner",
    values: [BigNumberish, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setRequests",
    values: [IVerifier.RequestStruct[]]
  ): string;
  encodeFunctionData(
    functionFragment: "setState",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setVerifierID",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "submitResponse",
    values: [
      IVerifier.AuthResponseStruct,
      IVerifier.ResponseStruct[],
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "updateRequest",
    values: [IVerifier.RequestStruct]
  ): string;
  encodeFunctionData(functionFragment: "version", values?: undefined): string;

  decodeFunctionResult(functionFragment: "VERSION", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "acceptOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "addValidatorToWhitelist",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "areMultiRequestProofsVerified",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "authMethodExists",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "disableAuthMethod",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "disableRequest",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "enableAuthMethod",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "enableRequest",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getAuthMethod",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getGroupedRequests",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getGroupsCount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getMultiRequest",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getMultiRequestProofsStatus",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "getRequest", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getRequestOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getRequestProofStatus",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getRequestsCount",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getResponseFieldValue",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getResponseFields",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getStateAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getVerifierID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "groupIdExists",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "isRequestEnabled",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "isRequestProofVerified",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "isWhitelistedValidator",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "multiRequestIdExists",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "pendingOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "removeValidatorFromWhitelist",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "requestIdExists",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setAuthMethod",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setMultiRequest",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setRequestOwner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setRequests",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "setState", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setVerifierID",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "submitResponse",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateRequest",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "version", data: BytesLike): Result;
}

export namespace AuthMethodSetEvent {
  export type InputTuple = [
    authMethod: string,
    validator: AddressLike,
    params: BytesLike
  ];
  export type OutputTuple = [
    authMethod: string,
    validator: string,
    params: string
  ];
  export interface OutputObject {
    authMethod: string;
    validator: string;
    params: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace AuthResponseSubmittedEvent {
  export type InputTuple = [authMethod: string, caller: AddressLike];
  export type OutputTuple = [authMethod: string, caller: string];
  export interface OutputObject {
    authMethod: string;
    caller: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace InitializedEvent {
  export type InputTuple = [version: BigNumberish];
  export type OutputTuple = [version: bigint];
  export interface OutputObject {
    version: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace MultiRequestSetEvent {
  export type InputTuple = [
    multiRequestId: BigNumberish,
    requestIds: BigNumberish[],
    groupIds: BigNumberish[]
  ];
  export type OutputTuple = [
    multiRequestId: bigint,
    requestIds: bigint[],
    groupIds: bigint[]
  ];
  export interface OutputObject {
    multiRequestId: bigint;
    requestIds: bigint[];
    groupIds: bigint[];
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace OwnershipTransferStartedEvent {
  export type InputTuple = [previousOwner: AddressLike, newOwner: AddressLike];
  export type OutputTuple = [previousOwner: string, newOwner: string];
  export interface OutputObject {
    previousOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace OwnershipTransferredEvent {
  export type InputTuple = [previousOwner: AddressLike, newOwner: AddressLike];
  export type OutputTuple = [previousOwner: string, newOwner: string];
  export interface OutputObject {
    previousOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RequestSetEvent {
  export type InputTuple = [
    requestId: BigNumberish,
    creator: AddressLike,
    metadata: string,
    validator: AddressLike,
    params: BytesLike
  ];
  export type OutputTuple = [
    requestId: bigint,
    creator: string,
    metadata: string,
    validator: string,
    params: string
  ];
  export interface OutputObject {
    requestId: bigint;
    creator: string;
    metadata: string;
    validator: string;
    params: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace RequestUpdateEvent {
  export type InputTuple = [
    requestId: BigNumberish,
    creator: AddressLike,
    metadata: string,
    validator: AddressLike,
    params: BytesLike
  ];
  export type OutputTuple = [
    requestId: bigint,
    creator: string,
    metadata: string,
    validator: string,
    params: string
  ];
  export interface OutputObject {
    requestId: bigint;
    creator: string;
    metadata: string;
    validator: string;
    params: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace ResponseSubmittedEvent {
  export type InputTuple = [requestId: BigNumberish, caller: AddressLike];
  export type OutputTuple = [requestId: bigint, caller: string];
  export interface OutputObject {
    requestId: bigint;
    caller: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface UniversalVerifier extends BaseContract {
  connect(runner?: ContractRunner | null): UniversalVerifier;
  waitForDeployment(): Promise<this>;

  interface: UniversalVerifierInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  VERSION: TypedContractMethod<[], [string], "view">;

  acceptOwnership: TypedContractMethod<[], [void], "nonpayable">;

  addValidatorToWhitelist: TypedContractMethod<
    [validator: AddressLike],
    [void],
    "nonpayable"
  >;

  areMultiRequestProofsVerified: TypedContractMethod<
    [multiRequestId: BigNumberish, userAddress: AddressLike],
    [boolean],
    "view"
  >;

  authMethodExists: TypedContractMethod<
    [authMethod: string],
    [boolean],
    "view"
  >;

  disableAuthMethod: TypedContractMethod<
    [authMethod: string],
    [void],
    "nonpayable"
  >;

  disableRequest: TypedContractMethod<
    [requestId: BigNumberish],
    [void],
    "nonpayable"
  >;

  enableAuthMethod: TypedContractMethod<
    [authMethod: string],
    [void],
    "nonpayable"
  >;

  enableRequest: TypedContractMethod<
    [requestId: BigNumberish],
    [void],
    "nonpayable"
  >;

  getAuthMethod: TypedContractMethod<
    [authMethod: string],
    [Verifier.AuthMethodDataStructOutput],
    "view"
  >;

  getGroupedRequests: TypedContractMethod<
    [groupID: BigNumberish],
    [IVerifier.RequestInfoStructOutput[]],
    "view"
  >;

  getGroupsCount: TypedContractMethod<[], [bigint], "view">;

  getMultiRequest: TypedContractMethod<
    [multiRequestId: BigNumberish],
    [IVerifier.MultiRequestStructOutput],
    "view"
  >;

  getMultiRequestProofsStatus: TypedContractMethod<
    [multiRequestId: BigNumberish, userAddress: AddressLike],
    [IVerifier.RequestProofStatusStructOutput[]],
    "view"
  >;

  getRequest: TypedContractMethod<
    [requestId: BigNumberish],
    [IVerifier.RequestInfoStructOutput],
    "view"
  >;

  getRequestOwner: TypedContractMethod<
    [requestId: BigNumberish],
    [string],
    "view"
  >;

  getRequestProofStatus: TypedContractMethod<
    [sender: AddressLike, requestId: BigNumberish],
    [IVerifier.RequestProofStatusStructOutput],
    "view"
  >;

  getRequestsCount: TypedContractMethod<[], [bigint], "view">;

  getResponseFieldValue: TypedContractMethod<
    [requestId: BigNumberish, sender: AddressLike, responseFieldName: string],
    [bigint],
    "view"
  >;

  getResponseFields: TypedContractMethod<
    [requestId: BigNumberish, sender: AddressLike],
    [IRequestValidator.ResponseFieldStructOutput[]],
    "view"
  >;

  getStateAddress: TypedContractMethod<[], [string], "view">;

  getVerifierID: TypedContractMethod<[], [bigint], "view">;

  groupIdExists: TypedContractMethod<
    [groupId: BigNumberish],
    [boolean],
    "view"
  >;

  initialize: TypedContractMethod<
    [state: AddressLike, owner: AddressLike],
    [void],
    "nonpayable"
  >;

  isRequestEnabled: TypedContractMethod<
    [requestId: BigNumberish],
    [boolean],
    "view"
  >;

  isRequestProofVerified: TypedContractMethod<
    [sender: AddressLike, requestId: BigNumberish],
    [boolean],
    "view"
  >;

  isWhitelistedValidator: TypedContractMethod<
    [validator: AddressLike],
    [boolean],
    "view"
  >;

  multiRequestIdExists: TypedContractMethod<
    [multiRequestId: BigNumberish],
    [boolean],
    "view"
  >;

  owner: TypedContractMethod<[], [string], "view">;

  pendingOwner: TypedContractMethod<[], [string], "view">;

  removeValidatorFromWhitelist: TypedContractMethod<
    [validator: AddressLike],
    [void],
    "nonpayable"
  >;

  renounceOwnership: TypedContractMethod<[], [void], "nonpayable">;

  requestIdExists: TypedContractMethod<
    [requestId: BigNumberish],
    [boolean],
    "view"
  >;

  setAuthMethod: TypedContractMethod<
    [authMethod: IVerifier.AuthMethodStruct],
    [void],
    "nonpayable"
  >;

  setMultiRequest: TypedContractMethod<
    [multiRequest: IVerifier.MultiRequestStruct],
    [void],
    "nonpayable"
  >;

  setRequestOwner: TypedContractMethod<
    [requestId: BigNumberish, requestOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  setRequests: TypedContractMethod<
    [requests: IVerifier.RequestStruct[]],
    [void],
    "nonpayable"
  >;

  setState: TypedContractMethod<[state: AddressLike], [void], "nonpayable">;

  setVerifierID: TypedContractMethod<
    [verifierID: BigNumberish],
    [void],
    "nonpayable"
  >;

  submitResponse: TypedContractMethod<
    [
      authResponse: IVerifier.AuthResponseStruct,
      responses: IVerifier.ResponseStruct[],
      crossChainProofs: BytesLike
    ],
    [void],
    "nonpayable"
  >;

  transferOwnership: TypedContractMethod<
    [newOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  updateRequest: TypedContractMethod<
    [request: IVerifier.RequestStruct],
    [void],
    "nonpayable"
  >;

  version: TypedContractMethod<[], [string], "view">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "VERSION"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "acceptOwnership"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "addValidatorToWhitelist"
  ): TypedContractMethod<[validator: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "areMultiRequestProofsVerified"
  ): TypedContractMethod<
    [multiRequestId: BigNumberish, userAddress: AddressLike],
    [boolean],
    "view"
  >;
  getFunction(
    nameOrSignature: "authMethodExists"
  ): TypedContractMethod<[authMethod: string], [boolean], "view">;
  getFunction(
    nameOrSignature: "disableAuthMethod"
  ): TypedContractMethod<[authMethod: string], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "disableRequest"
  ): TypedContractMethod<[requestId: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "enableAuthMethod"
  ): TypedContractMethod<[authMethod: string], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "enableRequest"
  ): TypedContractMethod<[requestId: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "getAuthMethod"
  ): TypedContractMethod<
    [authMethod: string],
    [Verifier.AuthMethodDataStructOutput],
    "view"
  >;
  getFunction(
    nameOrSignature: "getGroupedRequests"
  ): TypedContractMethod<
    [groupID: BigNumberish],
    [IVerifier.RequestInfoStructOutput[]],
    "view"
  >;
  getFunction(
    nameOrSignature: "getGroupsCount"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "getMultiRequest"
  ): TypedContractMethod<
    [multiRequestId: BigNumberish],
    [IVerifier.MultiRequestStructOutput],
    "view"
  >;
  getFunction(
    nameOrSignature: "getMultiRequestProofsStatus"
  ): TypedContractMethod<
    [multiRequestId: BigNumberish, userAddress: AddressLike],
    [IVerifier.RequestProofStatusStructOutput[]],
    "view"
  >;
  getFunction(
    nameOrSignature: "getRequest"
  ): TypedContractMethod<
    [requestId: BigNumberish],
    [IVerifier.RequestInfoStructOutput],
    "view"
  >;
  getFunction(
    nameOrSignature: "getRequestOwner"
  ): TypedContractMethod<[requestId: BigNumberish], [string], "view">;
  getFunction(
    nameOrSignature: "getRequestProofStatus"
  ): TypedContractMethod<
    [sender: AddressLike, requestId: BigNumberish],
    [IVerifier.RequestProofStatusStructOutput],
    "view"
  >;
  getFunction(
    nameOrSignature: "getRequestsCount"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "getResponseFieldValue"
  ): TypedContractMethod<
    [requestId: BigNumberish, sender: AddressLike, responseFieldName: string],
    [bigint],
    "view"
  >;
  getFunction(
    nameOrSignature: "getResponseFields"
  ): TypedContractMethod<
    [requestId: BigNumberish, sender: AddressLike],
    [IRequestValidator.ResponseFieldStructOutput[]],
    "view"
  >;
  getFunction(
    nameOrSignature: "getStateAddress"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "getVerifierID"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "groupIdExists"
  ): TypedContractMethod<[groupId: BigNumberish], [boolean], "view">;
  getFunction(
    nameOrSignature: "initialize"
  ): TypedContractMethod<
    [state: AddressLike, owner: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "isRequestEnabled"
  ): TypedContractMethod<[requestId: BigNumberish], [boolean], "view">;
  getFunction(
    nameOrSignature: "isRequestProofVerified"
  ): TypedContractMethod<
    [sender: AddressLike, requestId: BigNumberish],
    [boolean],
    "view"
  >;
  getFunction(
    nameOrSignature: "isWhitelistedValidator"
  ): TypedContractMethod<[validator: AddressLike], [boolean], "view">;
  getFunction(
    nameOrSignature: "multiRequestIdExists"
  ): TypedContractMethod<[multiRequestId: BigNumberish], [boolean], "view">;
  getFunction(
    nameOrSignature: "owner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "pendingOwner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "removeValidatorFromWhitelist"
  ): TypedContractMethod<[validator: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "renounceOwnership"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "requestIdExists"
  ): TypedContractMethod<[requestId: BigNumberish], [boolean], "view">;
  getFunction(
    nameOrSignature: "setAuthMethod"
  ): TypedContractMethod<
    [authMethod: IVerifier.AuthMethodStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setMultiRequest"
  ): TypedContractMethod<
    [multiRequest: IVerifier.MultiRequestStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setRequestOwner"
  ): TypedContractMethod<
    [requestId: BigNumberish, requestOwner: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setRequests"
  ): TypedContractMethod<
    [requests: IVerifier.RequestStruct[]],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setState"
  ): TypedContractMethod<[state: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "setVerifierID"
  ): TypedContractMethod<[verifierID: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "submitResponse"
  ): TypedContractMethod<
    [
      authResponse: IVerifier.AuthResponseStruct,
      responses: IVerifier.ResponseStruct[],
      crossChainProofs: BytesLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "transferOwnership"
  ): TypedContractMethod<[newOwner: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "updateRequest"
  ): TypedContractMethod<
    [request: IVerifier.RequestStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "version"
  ): TypedContractMethod<[], [string], "view">;

  getEvent(
    key: "AuthMethodSet"
  ): TypedContractEvent<
    AuthMethodSetEvent.InputTuple,
    AuthMethodSetEvent.OutputTuple,
    AuthMethodSetEvent.OutputObject
  >;
  getEvent(
    key: "AuthResponseSubmitted"
  ): TypedContractEvent<
    AuthResponseSubmittedEvent.InputTuple,
    AuthResponseSubmittedEvent.OutputTuple,
    AuthResponseSubmittedEvent.OutputObject
  >;
  getEvent(
    key: "Initialized"
  ): TypedContractEvent<
    InitializedEvent.InputTuple,
    InitializedEvent.OutputTuple,
    InitializedEvent.OutputObject
  >;
  getEvent(
    key: "MultiRequestSet"
  ): TypedContractEvent<
    MultiRequestSetEvent.InputTuple,
    MultiRequestSetEvent.OutputTuple,
    MultiRequestSetEvent.OutputObject
  >;
  getEvent(
    key: "OwnershipTransferStarted"
  ): TypedContractEvent<
    OwnershipTransferStartedEvent.InputTuple,
    OwnershipTransferStartedEvent.OutputTuple,
    OwnershipTransferStartedEvent.OutputObject
  >;
  getEvent(
    key: "OwnershipTransferred"
  ): TypedContractEvent<
    OwnershipTransferredEvent.InputTuple,
    OwnershipTransferredEvent.OutputTuple,
    OwnershipTransferredEvent.OutputObject
  >;
  getEvent(
    key: "RequestSet"
  ): TypedContractEvent<
    RequestSetEvent.InputTuple,
    RequestSetEvent.OutputTuple,
    RequestSetEvent.OutputObject
  >;
  getEvent(
    key: "RequestUpdate"
  ): TypedContractEvent<
    RequestUpdateEvent.InputTuple,
    RequestUpdateEvent.OutputTuple,
    RequestUpdateEvent.OutputObject
  >;
  getEvent(
    key: "ResponseSubmitted"
  ): TypedContractEvent<
    ResponseSubmittedEvent.InputTuple,
    ResponseSubmittedEvent.OutputTuple,
    ResponseSubmittedEvent.OutputObject
  >;

  filters: {
    "AuthMethodSet(string,address,bytes)": TypedContractEvent<
      AuthMethodSetEvent.InputTuple,
      AuthMethodSetEvent.OutputTuple,
      AuthMethodSetEvent.OutputObject
    >;
    AuthMethodSet: TypedContractEvent<
      AuthMethodSetEvent.InputTuple,
      AuthMethodSetEvent.OutputTuple,
      AuthMethodSetEvent.OutputObject
    >;

    "AuthResponseSubmitted(string,address)": TypedContractEvent<
      AuthResponseSubmittedEvent.InputTuple,
      AuthResponseSubmittedEvent.OutputTuple,
      AuthResponseSubmittedEvent.OutputObject
    >;
    AuthResponseSubmitted: TypedContractEvent<
      AuthResponseSubmittedEvent.InputTuple,
      AuthResponseSubmittedEvent.OutputTuple,
      AuthResponseSubmittedEvent.OutputObject
    >;

    "Initialized(uint64)": TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;
    Initialized: TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;

    "MultiRequestSet(uint256,uint256[],uint256[])": TypedContractEvent<
      MultiRequestSetEvent.InputTuple,
      MultiRequestSetEvent.OutputTuple,
      MultiRequestSetEvent.OutputObject
    >;
    MultiRequestSet: TypedContractEvent<
      MultiRequestSetEvent.InputTuple,
      MultiRequestSetEvent.OutputTuple,
      MultiRequestSetEvent.OutputObject
    >;

    "OwnershipTransferStarted(address,address)": TypedContractEvent<
      OwnershipTransferStartedEvent.InputTuple,
      OwnershipTransferStartedEvent.OutputTuple,
      OwnershipTransferStartedEvent.OutputObject
    >;
    OwnershipTransferStarted: TypedContractEvent<
      OwnershipTransferStartedEvent.InputTuple,
      OwnershipTransferStartedEvent.OutputTuple,
      OwnershipTransferStartedEvent.OutputObject
    >;

    "OwnershipTransferred(address,address)": TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;
    OwnershipTransferred: TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;

    "RequestSet(uint256,address,string,address,bytes)": TypedContractEvent<
      RequestSetEvent.InputTuple,
      RequestSetEvent.OutputTuple,
      RequestSetEvent.OutputObject
    >;
    RequestSet: TypedContractEvent<
      RequestSetEvent.InputTuple,
      RequestSetEvent.OutputTuple,
      RequestSetEvent.OutputObject
    >;

    "RequestUpdate(uint256,address,string,address,bytes)": TypedContractEvent<
      RequestUpdateEvent.InputTuple,
      RequestUpdateEvent.OutputTuple,
      RequestUpdateEvent.OutputObject
    >;
    RequestUpdate: TypedContractEvent<
      RequestUpdateEvent.InputTuple,
      RequestUpdateEvent.OutputTuple,
      RequestUpdateEvent.OutputObject
    >;

    "ResponseSubmitted(uint256,address)": TypedContractEvent<
      ResponseSubmittedEvent.InputTuple,
      ResponseSubmittedEvent.OutputTuple,
      ResponseSubmittedEvent.OutputObject
    >;
    ResponseSubmitted: TypedContractEvent<
      ResponseSubmittedEvent.InputTuple,
      ResponseSubmittedEvent.OutputTuple,
      ResponseSubmittedEvent.OutputObject
    >;
  };
}
