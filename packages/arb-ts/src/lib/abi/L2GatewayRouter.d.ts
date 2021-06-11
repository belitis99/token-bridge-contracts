/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import {
  ethers,
  EventFilter,
  Signer,
  BigNumber,
  BigNumberish,
  PopulatedTransaction,
} from 'ethers'
import {
  Contract,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  CallOverrides,
} from '@ethersproject/contracts'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'

interface L2GatewayRouterInterface extends ethers.utils.Interface {
  functions: {
    'calculateL2TokenAddress(address)': FunctionFragment
    'counterpartGateway()': FunctionFragment
    'defaultGateway()': FunctionFragment
    'finalizeInboundTransfer(address,address,address,uint256,bytes)': FunctionFragment
    'getGateway(address)': FunctionFragment
    'getOutboundCalldata(address,address,address,uint256,bytes)': FunctionFragment
    'initialize(address,address)': FunctionFragment
    'l1TokenToGateway(address)': FunctionFragment
    'outboundTransfer(address,address,uint256,bytes)': FunctionFragment
    'router()': FunctionFragment
    'setDefaultGateway(address)': FunctionFragment
    'setGateway(address[],address[])': FunctionFragment
  }

  encodeFunctionData(
    functionFragment: 'calculateL2TokenAddress',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'counterpartGateway',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'defaultGateway',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'finalizeInboundTransfer',
    values: [string, string, string, BigNumberish, BytesLike]
  ): string
  encodeFunctionData(functionFragment: 'getGateway', values: [string]): string
  encodeFunctionData(
    functionFragment: 'getOutboundCalldata',
    values: [string, string, string, BigNumberish, BytesLike]
  ): string
  encodeFunctionData(
    functionFragment: 'initialize',
    values: [string, string]
  ): string
  encodeFunctionData(
    functionFragment: 'l1TokenToGateway',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'outboundTransfer',
    values: [string, string, BigNumberish, BytesLike]
  ): string
  encodeFunctionData(functionFragment: 'router', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'setDefaultGateway',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'setGateway',
    values: [string[], string[]]
  ): string

  decodeFunctionResult(
    functionFragment: 'calculateL2TokenAddress',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'counterpartGateway',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'defaultGateway',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'finalizeInboundTransfer',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'getGateway', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'getOutboundCalldata',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'initialize', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'l1TokenToGateway',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'outboundTransfer',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'router', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'setDefaultGateway',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'setGateway', data: BytesLike): Result

  events: {
    'DefaultGatewayUpdated(address)': EventFragment
    'GatewaySet(address,address)': EventFragment
    'InboundTransferFinalized(address,address,address,uint256,uint256,bytes)': EventFragment
    'OutboundTransferInitiated(address,address,address,uint256,uint256,bytes)': EventFragment
    'TransferRouted(address,address,address,address)': EventFragment
  }

  getEvent(nameOrSignatureOrTopic: 'DefaultGatewayUpdated'): EventFragment
  getEvent(nameOrSignatureOrTopic: 'GatewaySet'): EventFragment
  getEvent(nameOrSignatureOrTopic: 'InboundTransferFinalized'): EventFragment
  getEvent(nameOrSignatureOrTopic: 'OutboundTransferInitiated'): EventFragment
  getEvent(nameOrSignatureOrTopic: 'TransferRouted'): EventFragment
}

export class L2GatewayRouter extends Contract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  on(event: EventFilter | string, listener: Listener): this
  once(event: EventFilter | string, listener: Listener): this
  addListener(eventName: EventFilter | string, listener: Listener): this
  removeAllListeners(eventName: EventFilter | string): this
  removeListener(eventName: any, listener: Listener): this

  interface: L2GatewayRouterInterface

  functions: {
    calculateL2TokenAddress(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<[string]>

    'calculateL2TokenAddress(address)'(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<[string]>

    counterpartGateway(overrides?: CallOverrides): Promise<[string]>

    'counterpartGateway()'(overrides?: CallOverrides): Promise<[string]>

    defaultGateway(overrides?: CallOverrides): Promise<[string]>

    'defaultGateway()'(overrides?: CallOverrides): Promise<[string]>

    finalizeInboundTransfer(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'finalizeInboundTransfer(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    getGateway(
      _token: string,
      overrides?: CallOverrides
    ): Promise<[string] & { gateway: string }>

    'getGateway(address)'(
      _token: string,
      overrides?: CallOverrides
    ): Promise<[string] & { gateway: string }>

    getOutboundCalldata(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<[string]>

    'getOutboundCalldata(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<[string]>

    initialize(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'initialize(address,address)'(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    l1TokenToGateway(arg0: string, overrides?: CallOverrides): Promise<[string]>

    'l1TokenToGateway(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<[string]>

    'outboundTransfer(address,address,uint256,bytes)'(
      _l1Token: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'outboundTransfer(address,address,uint256,uint256,uint256,bytes)'(
      _token: string,
      _to: string,
      _amount: BigNumberish,
      _maxGas: BigNumberish,
      _gasPriceBid: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    router(overrides?: CallOverrides): Promise<[string]>

    'router()'(overrides?: CallOverrides): Promise<[string]>

    setDefaultGateway(
      newL2DefaultGateway: string,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'setDefaultGateway(address)'(
      newL2DefaultGateway: string,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    setGateway(
      _l1Token: string[],
      _gateway: string[],
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'setGateway(address[],address[])'(
      _l1Token: string[],
      _gateway: string[],
      overrides?: Overrides
    ): Promise<ContractTransaction>
  }

  calculateL2TokenAddress(
    l1ERC20: string,
    overrides?: CallOverrides
  ): Promise<string>

  'calculateL2TokenAddress(address)'(
    l1ERC20: string,
    overrides?: CallOverrides
  ): Promise<string>

  counterpartGateway(overrides?: CallOverrides): Promise<string>

  'counterpartGateway()'(overrides?: CallOverrides): Promise<string>

  defaultGateway(overrides?: CallOverrides): Promise<string>

  'defaultGateway()'(overrides?: CallOverrides): Promise<string>

  finalizeInboundTransfer(
    _token: string,
    _from: string,
    _to: string,
    _amount: BigNumberish,
    _data: BytesLike,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'finalizeInboundTransfer(address,address,address,uint256,bytes)'(
    _token: string,
    _from: string,
    _to: string,
    _amount: BigNumberish,
    _data: BytesLike,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  getGateway(_token: string, overrides?: CallOverrides): Promise<string>

  'getGateway(address)'(
    _token: string,
    overrides?: CallOverrides
  ): Promise<string>

  getOutboundCalldata(
    _token: string,
    _from: string,
    _to: string,
    _amount: BigNumberish,
    _data: BytesLike,
    overrides?: CallOverrides
  ): Promise<string>

  'getOutboundCalldata(address,address,address,uint256,bytes)'(
    _token: string,
    _from: string,
    _to: string,
    _amount: BigNumberish,
    _data: BytesLike,
    overrides?: CallOverrides
  ): Promise<string>

  initialize(
    _counterpartGateway: string,
    _defaultGateway: string,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'initialize(address,address)'(
    _counterpartGateway: string,
    _defaultGateway: string,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  l1TokenToGateway(arg0: string, overrides?: CallOverrides): Promise<string>

  'l1TokenToGateway(address)'(
    arg0: string,
    overrides?: CallOverrides
  ): Promise<string>

  'outboundTransfer(address,address,uint256,bytes)'(
    _l1Token: string,
    _to: string,
    _amount: BigNumberish,
    _data: BytesLike,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'outboundTransfer(address,address,uint256,uint256,uint256,bytes)'(
    _token: string,
    _to: string,
    _amount: BigNumberish,
    _maxGas: BigNumberish,
    _gasPriceBid: BigNumberish,
    _data: BytesLike,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  router(overrides?: CallOverrides): Promise<string>

  'router()'(overrides?: CallOverrides): Promise<string>

  setDefaultGateway(
    newL2DefaultGateway: string,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'setDefaultGateway(address)'(
    newL2DefaultGateway: string,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  setGateway(
    _l1Token: string[],
    _gateway: string[],
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'setGateway(address[],address[])'(
    _l1Token: string[],
    _gateway: string[],
    overrides?: Overrides
  ): Promise<ContractTransaction>

  callStatic: {
    calculateL2TokenAddress(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<string>

    'calculateL2TokenAddress(address)'(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<string>

    counterpartGateway(overrides?: CallOverrides): Promise<string>

    'counterpartGateway()'(overrides?: CallOverrides): Promise<string>

    defaultGateway(overrides?: CallOverrides): Promise<string>

    'defaultGateway()'(overrides?: CallOverrides): Promise<string>

    finalizeInboundTransfer(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<string>

    'finalizeInboundTransfer(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<string>

    getGateway(_token: string, overrides?: CallOverrides): Promise<string>

    'getGateway(address)'(
      _token: string,
      overrides?: CallOverrides
    ): Promise<string>

    getOutboundCalldata(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<string>

    'getOutboundCalldata(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<string>

    initialize(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: CallOverrides
    ): Promise<void>

    'initialize(address,address)'(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: CallOverrides
    ): Promise<void>

    l1TokenToGateway(arg0: string, overrides?: CallOverrides): Promise<string>

    'l1TokenToGateway(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<string>

    'outboundTransfer(address,address,uint256,bytes)'(
      _l1Token: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<string>

    'outboundTransfer(address,address,uint256,uint256,uint256,bytes)'(
      _token: string,
      _to: string,
      _amount: BigNumberish,
      _maxGas: BigNumberish,
      _gasPriceBid: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<string>

    router(overrides?: CallOverrides): Promise<string>

    'router()'(overrides?: CallOverrides): Promise<string>

    setDefaultGateway(
      newL2DefaultGateway: string,
      overrides?: CallOverrides
    ): Promise<void>

    'setDefaultGateway(address)'(
      newL2DefaultGateway: string,
      overrides?: CallOverrides
    ): Promise<void>

    setGateway(
      _l1Token: string[],
      _gateway: string[],
      overrides?: CallOverrides
    ): Promise<void>

    'setGateway(address[],address[])'(
      _l1Token: string[],
      _gateway: string[],
      overrides?: CallOverrides
    ): Promise<void>
  }

  filters: {
    DefaultGatewayUpdated(newDefaultGateway: null): EventFilter

    GatewaySet(l1Token: string | null, gateway: string | null): EventFilter

    InboundTransferFinalized(
      token: null,
      _from: string | null,
      _to: string | null,
      _transferId: BigNumberish | null,
      _amount: null,
      _data: null
    ): EventFilter

    OutboundTransferInitiated(
      token: null,
      _from: string | null,
      _to: string | null,
      _transferId: BigNumberish | null,
      _amount: null,
      _data: null
    ): EventFilter

    TransferRouted(
      token: string | null,
      _userFrom: string | null,
      _userTo: string | null,
      gateway: null
    ): EventFilter
  }

  estimateGas: {
    calculateL2TokenAddress(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    'calculateL2TokenAddress(address)'(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    counterpartGateway(overrides?: CallOverrides): Promise<BigNumber>

    'counterpartGateway()'(overrides?: CallOverrides): Promise<BigNumber>

    defaultGateway(overrides?: CallOverrides): Promise<BigNumber>

    'defaultGateway()'(overrides?: CallOverrides): Promise<BigNumber>

    finalizeInboundTransfer(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'finalizeInboundTransfer(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    getGateway(_token: string, overrides?: CallOverrides): Promise<BigNumber>

    'getGateway(address)'(
      _token: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getOutboundCalldata(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    'getOutboundCalldata(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    initialize(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: Overrides
    ): Promise<BigNumber>

    'initialize(address,address)'(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: Overrides
    ): Promise<BigNumber>

    l1TokenToGateway(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    'l1TokenToGateway(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    'outboundTransfer(address,address,uint256,bytes)'(
      _l1Token: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'outboundTransfer(address,address,uint256,uint256,uint256,bytes)'(
      _token: string,
      _to: string,
      _amount: BigNumberish,
      _maxGas: BigNumberish,
      _gasPriceBid: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    router(overrides?: CallOverrides): Promise<BigNumber>

    'router()'(overrides?: CallOverrides): Promise<BigNumber>

    setDefaultGateway(
      newL2DefaultGateway: string,
      overrides?: Overrides
    ): Promise<BigNumber>

    'setDefaultGateway(address)'(
      newL2DefaultGateway: string,
      overrides?: Overrides
    ): Promise<BigNumber>

    setGateway(
      _l1Token: string[],
      _gateway: string[],
      overrides?: Overrides
    ): Promise<BigNumber>

    'setGateway(address[],address[])'(
      _l1Token: string[],
      _gateway: string[],
      overrides?: Overrides
    ): Promise<BigNumber>
  }

  populateTransaction: {
    calculateL2TokenAddress(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    'calculateL2TokenAddress(address)'(
      l1ERC20: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    counterpartGateway(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'counterpartGateway()'(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    defaultGateway(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'defaultGateway()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    finalizeInboundTransfer(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'finalizeInboundTransfer(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    getGateway(
      _token: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    'getGateway(address)'(
      _token: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getOutboundCalldata(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    'getOutboundCalldata(address,address,address,uint256,bytes)'(
      _token: string,
      _from: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    initialize(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'initialize(address,address)'(
      _counterpartGateway: string,
      _defaultGateway: string,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    l1TokenToGateway(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    'l1TokenToGateway(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    'outboundTransfer(address,address,uint256,bytes)'(
      _l1Token: string,
      _to: string,
      _amount: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'outboundTransfer(address,address,uint256,uint256,uint256,bytes)'(
      _token: string,
      _to: string,
      _amount: BigNumberish,
      _maxGas: BigNumberish,
      _gasPriceBid: BigNumberish,
      _data: BytesLike,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    router(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'router()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    setDefaultGateway(
      newL2DefaultGateway: string,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'setDefaultGateway(address)'(
      newL2DefaultGateway: string,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    setGateway(
      _l1Token: string[],
      _gateway: string[],
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'setGateway(address[],address[])'(
      _l1Token: string[],
      _gateway: string[],
      overrides?: Overrides
    ): Promise<PopulatedTransaction>
  }
}
