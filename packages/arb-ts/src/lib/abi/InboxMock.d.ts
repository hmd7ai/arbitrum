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
  BaseContract,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  CallOverrides,
} from 'ethers'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'
import type { TypedEventFilter, TypedEvent, TypedListener } from './common'

interface InboxMockInterface extends ethers.utils.Interface {
  functions: {
    'activeOutbox()': FunctionFragment
    'bridge()': FunctionFragment
    'createRetryableTicket(address,uint256,uint256,address,address,uint256,uint256,bytes)': FunctionFragment
    'l2ToL1Sender()': FunctionFragment
    'setL2ToL1Sender(address)': FunctionFragment
  }

  encodeFunctionData(
    functionFragment: 'activeOutbox',
    values?: undefined
  ): string
  encodeFunctionData(functionFragment: 'bridge', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'createRetryableTicket',
    values: [
      string,
      BigNumberish,
      BigNumberish,
      string,
      string,
      BigNumberish,
      BigNumberish,
      BytesLike
    ]
  ): string
  encodeFunctionData(
    functionFragment: 'l2ToL1Sender',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'setL2ToL1Sender',
    values: [string]
  ): string

  decodeFunctionResult(
    functionFragment: 'activeOutbox',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'bridge', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'createRetryableTicket',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'l2ToL1Sender',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'setL2ToL1Sender',
    data: BytesLike
  ): Result

  events: {
    'TicketData(uint256)': EventFragment
  }

  getEvent(nameOrSignatureOrTopic: 'TicketData'): EventFragment
}

export type TicketDataEvent = TypedEvent<
  [BigNumber] & { maxSubmissionCost: BigNumber }
>

export class InboxMock extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  listeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter?: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): Array<TypedListener<EventArgsArray, EventArgsObject>>
  off<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  on<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  once<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  removeListener<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  removeAllListeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): this

  listeners(eventName?: string): Array<Listener>
  off(eventName: string, listener: Listener): this
  on(eventName: string, listener: Listener): this
  once(eventName: string, listener: Listener): this
  removeListener(eventName: string, listener: Listener): this
  removeAllListeners(eventName?: string): this

  queryFilter<EventArgsArray extends Array<any>, EventArgsObject>(
    event: TypedEventFilter<EventArgsArray, EventArgsObject>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEvent<EventArgsArray & EventArgsObject>>>

  interface: InboxMockInterface

  functions: {
    activeOutbox(overrides?: CallOverrides): Promise<[string]>

    bridge(overrides?: CallOverrides): Promise<[string]>

    createRetryableTicket(
      arg0: string,
      arg1: BigNumberish,
      maxSubmissionCost: BigNumberish,
      arg3: string,
      arg4: string,
      arg5: BigNumberish,
      arg6: BigNumberish,
      arg7: BytesLike,
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>

    l2ToL1Sender(overrides?: CallOverrides): Promise<[string]>

    setL2ToL1Sender(
      sender: string,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>
  }

  activeOutbox(overrides?: CallOverrides): Promise<string>

  bridge(overrides?: CallOverrides): Promise<string>

  createRetryableTicket(
    arg0: string,
    arg1: BigNumberish,
    maxSubmissionCost: BigNumberish,
    arg3: string,
    arg4: string,
    arg5: BigNumberish,
    arg6: BigNumberish,
    arg7: BytesLike,
    overrides?: PayableOverrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  l2ToL1Sender(overrides?: CallOverrides): Promise<string>

  setL2ToL1Sender(
    sender: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  callStatic: {
    activeOutbox(overrides?: CallOverrides): Promise<string>

    bridge(overrides?: CallOverrides): Promise<string>

    createRetryableTicket(
      arg0: string,
      arg1: BigNumberish,
      maxSubmissionCost: BigNumberish,
      arg3: string,
      arg4: string,
      arg5: BigNumberish,
      arg6: BigNumberish,
      arg7: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    l2ToL1Sender(overrides?: CallOverrides): Promise<string>

    setL2ToL1Sender(sender: string, overrides?: CallOverrides): Promise<void>
  }

  filters: {
    'TicketData(uint256)'(
      maxSubmissionCost?: null
    ): TypedEventFilter<[BigNumber], { maxSubmissionCost: BigNumber }>

    TicketData(
      maxSubmissionCost?: null
    ): TypedEventFilter<[BigNumber], { maxSubmissionCost: BigNumber }>
  }

  estimateGas: {
    activeOutbox(overrides?: CallOverrides): Promise<BigNumber>

    bridge(overrides?: CallOverrides): Promise<BigNumber>

    createRetryableTicket(
      arg0: string,
      arg1: BigNumberish,
      maxSubmissionCost: BigNumberish,
      arg3: string,
      arg4: string,
      arg5: BigNumberish,
      arg6: BigNumberish,
      arg7: BytesLike,
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>

    l2ToL1Sender(overrides?: CallOverrides): Promise<BigNumber>

    setL2ToL1Sender(
      sender: string,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>
  }

  populateTransaction: {
    activeOutbox(overrides?: CallOverrides): Promise<PopulatedTransaction>

    bridge(overrides?: CallOverrides): Promise<PopulatedTransaction>

    createRetryableTicket(
      arg0: string,
      arg1: BigNumberish,
      maxSubmissionCost: BigNumberish,
      arg3: string,
      arg4: string,
      arg5: BigNumberish,
      arg6: BigNumberish,
      arg7: BytesLike,
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>

    l2ToL1Sender(overrides?: CallOverrides): Promise<PopulatedTransaction>

    setL2ToL1Sender(
      sender: string,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>
  }
}
