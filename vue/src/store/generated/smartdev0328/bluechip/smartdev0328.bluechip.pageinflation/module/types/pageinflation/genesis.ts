/* eslint-disable */
import { Params } from "../pageinflation/params";
import { MintDenom } from "../pageinflation/mint_denom";
import { StartBlock } from "../pageinflation/start_block";
import { MintedPool } from "../pageinflation/minted_pool";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "smartdev0328.bluechip.pageinflation";

/** GenesisState defines the pageinflation module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  mintDenom: MintDenom | undefined;
  startBlock: StartBlock | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  mintedPool: MintedPool | undefined;
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.mintDenom !== undefined) {
      MintDenom.encode(message.mintDenom, writer.uint32(18).fork()).ldelim();
    }
    if (message.startBlock !== undefined) {
      StartBlock.encode(message.startBlock, writer.uint32(26).fork()).ldelim();
    }
    if (message.mintedPool !== undefined) {
      MintedPool.encode(message.mintedPool, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.mintDenom = MintDenom.decode(reader, reader.uint32());
          break;
        case 3:
          message.startBlock = StartBlock.decode(reader, reader.uint32());
          break;
        case 4:
          message.mintedPool = MintedPool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.mintDenom !== undefined && object.mintDenom !== null) {
      message.mintDenom = MintDenom.fromJSON(object.mintDenom);
    } else {
      message.mintDenom = undefined;
    }
    if (object.startBlock !== undefined && object.startBlock !== null) {
      message.startBlock = StartBlock.fromJSON(object.startBlock);
    } else {
      message.startBlock = undefined;
    }
    if (object.mintedPool !== undefined && object.mintedPool !== null) {
      message.mintedPool = MintedPool.fromJSON(object.mintedPool);
    } else {
      message.mintedPool = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.mintDenom !== undefined &&
      (obj.mintDenom = message.mintDenom
        ? MintDenom.toJSON(message.mintDenom)
        : undefined);
    message.startBlock !== undefined &&
      (obj.startBlock = message.startBlock
        ? StartBlock.toJSON(message.startBlock)
        : undefined);
    message.mintedPool !== undefined &&
      (obj.mintedPool = message.mintedPool
        ? MintedPool.toJSON(message.mintedPool)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.mintDenom !== undefined && object.mintDenom !== null) {
      message.mintDenom = MintDenom.fromPartial(object.mintDenom);
    } else {
      message.mintDenom = undefined;
    }
    if (object.startBlock !== undefined && object.startBlock !== null) {
      message.startBlock = StartBlock.fromPartial(object.startBlock);
    } else {
      message.startBlock = undefined;
    }
    if (object.mintedPool !== undefined && object.mintedPool !== null) {
      message.mintedPool = MintedPool.fromPartial(object.mintedPool);
    } else {
      message.mintedPool = undefined;
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
