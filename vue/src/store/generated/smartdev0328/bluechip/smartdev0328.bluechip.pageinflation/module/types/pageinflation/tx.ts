/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "smartdev0328.bluechip.pageinflation";

export interface MsgCreatorPoolMint {
  creator: string;
}

export interface MsgCreatorPoolMintResponse {}

const baseMsgCreatorPoolMint: object = { creator: "" };

export const MsgCreatorPoolMint = {
  encode(
    message: MsgCreatorPoolMint,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatorPoolMint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatorPoolMint } as MsgCreatorPoolMint;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatorPoolMint {
    const message = { ...baseMsgCreatorPoolMint } as MsgCreatorPoolMint;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgCreatorPoolMint): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatorPoolMint>): MsgCreatorPoolMint {
    const message = { ...baseMsgCreatorPoolMint } as MsgCreatorPoolMint;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgCreatorPoolMintResponse: object = {};

export const MsgCreatorPoolMintResponse = {
  encode(
    _: MsgCreatorPoolMintResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreatorPoolMintResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreatorPoolMintResponse,
    } as MsgCreatorPoolMintResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreatorPoolMintResponse {
    const message = {
      ...baseMsgCreatorPoolMintResponse,
    } as MsgCreatorPoolMintResponse;
    return message;
  },

  toJSON(_: MsgCreatorPoolMintResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreatorPoolMintResponse>
  ): MsgCreatorPoolMintResponse {
    const message = {
      ...baseMsgCreatorPoolMintResponse,
    } as MsgCreatorPoolMintResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreatorPoolMint(
    request: MsgCreatorPoolMint
  ): Promise<MsgCreatorPoolMintResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatorPoolMint(
    request: MsgCreatorPoolMint
  ): Promise<MsgCreatorPoolMintResponse> {
    const data = MsgCreatorPoolMint.encode(request).finish();
    const promise = this.rpc.request(
      "smartdev0328.bluechip.pageinflation.Msg",
      "CreatorPoolMint",
      data
    );
    return promise.then((data) =>
      MsgCreatorPoolMintResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
