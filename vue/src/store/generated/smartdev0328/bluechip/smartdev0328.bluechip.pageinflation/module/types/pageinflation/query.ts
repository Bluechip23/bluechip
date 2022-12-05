/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../pageinflation/params";
import { MintDenom } from "../pageinflation/mint_denom";
import { StartBlock } from "../pageinflation/start_block";
import { MintedPool } from "../pageinflation/minted_pool";

export const protobufPackage = "smartdev0328.bluechip.pageinflation";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetMintDenomRequest {}

export interface QueryGetMintDenomResponse {
  MintDenom: MintDenom | undefined;
}

export interface QueryGetStartBlockRequest {}

export interface QueryGetStartBlockResponse {
  StartBlock: StartBlock | undefined;
}

export interface QueryGetMintedPoolRequest {}

export interface QueryGetMintedPoolResponse {
  MintedPool: MintedPool | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetMintDenomRequest: object = {};

export const QueryGetMintDenomRequest = {
  encode(
    _: QueryGetMintDenomRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMintDenomRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMintDenomRequest,
    } as QueryGetMintDenomRequest;
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

  fromJSON(_: any): QueryGetMintDenomRequest {
    const message = {
      ...baseQueryGetMintDenomRequest,
    } as QueryGetMintDenomRequest;
    return message;
  },

  toJSON(_: QueryGetMintDenomRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetMintDenomRequest>
  ): QueryGetMintDenomRequest {
    const message = {
      ...baseQueryGetMintDenomRequest,
    } as QueryGetMintDenomRequest;
    return message;
  },
};

const baseQueryGetMintDenomResponse: object = {};

export const QueryGetMintDenomResponse = {
  encode(
    message: QueryGetMintDenomResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.MintDenom !== undefined) {
      MintDenom.encode(message.MintDenom, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMintDenomResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMintDenomResponse,
    } as QueryGetMintDenomResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MintDenom = MintDenom.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMintDenomResponse {
    const message = {
      ...baseQueryGetMintDenomResponse,
    } as QueryGetMintDenomResponse;
    if (object.MintDenom !== undefined && object.MintDenom !== null) {
      message.MintDenom = MintDenom.fromJSON(object.MintDenom);
    } else {
      message.MintDenom = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMintDenomResponse): unknown {
    const obj: any = {};
    message.MintDenom !== undefined &&
      (obj.MintDenom = message.MintDenom
        ? MintDenom.toJSON(message.MintDenom)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMintDenomResponse>
  ): QueryGetMintDenomResponse {
    const message = {
      ...baseQueryGetMintDenomResponse,
    } as QueryGetMintDenomResponse;
    if (object.MintDenom !== undefined && object.MintDenom !== null) {
      message.MintDenom = MintDenom.fromPartial(object.MintDenom);
    } else {
      message.MintDenom = undefined;
    }
    return message;
  },
};

const baseQueryGetStartBlockRequest: object = {};

export const QueryGetStartBlockRequest = {
  encode(
    _: QueryGetStartBlockRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStartBlockRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStartBlockRequest,
    } as QueryGetStartBlockRequest;
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

  fromJSON(_: any): QueryGetStartBlockRequest {
    const message = {
      ...baseQueryGetStartBlockRequest,
    } as QueryGetStartBlockRequest;
    return message;
  },

  toJSON(_: QueryGetStartBlockRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetStartBlockRequest>
  ): QueryGetStartBlockRequest {
    const message = {
      ...baseQueryGetStartBlockRequest,
    } as QueryGetStartBlockRequest;
    return message;
  },
};

const baseQueryGetStartBlockResponse: object = {};

export const QueryGetStartBlockResponse = {
  encode(
    message: QueryGetStartBlockResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.StartBlock !== undefined) {
      StartBlock.encode(message.StartBlock, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStartBlockResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStartBlockResponse,
    } as QueryGetStartBlockResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.StartBlock = StartBlock.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStartBlockResponse {
    const message = {
      ...baseQueryGetStartBlockResponse,
    } as QueryGetStartBlockResponse;
    if (object.StartBlock !== undefined && object.StartBlock !== null) {
      message.StartBlock = StartBlock.fromJSON(object.StartBlock);
    } else {
      message.StartBlock = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStartBlockResponse): unknown {
    const obj: any = {};
    message.StartBlock !== undefined &&
      (obj.StartBlock = message.StartBlock
        ? StartBlock.toJSON(message.StartBlock)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStartBlockResponse>
  ): QueryGetStartBlockResponse {
    const message = {
      ...baseQueryGetStartBlockResponse,
    } as QueryGetStartBlockResponse;
    if (object.StartBlock !== undefined && object.StartBlock !== null) {
      message.StartBlock = StartBlock.fromPartial(object.StartBlock);
    } else {
      message.StartBlock = undefined;
    }
    return message;
  },
};

const baseQueryGetMintedPoolRequest: object = {};

export const QueryGetMintedPoolRequest = {
  encode(
    _: QueryGetMintedPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMintedPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMintedPoolRequest,
    } as QueryGetMintedPoolRequest;
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

  fromJSON(_: any): QueryGetMintedPoolRequest {
    const message = {
      ...baseQueryGetMintedPoolRequest,
    } as QueryGetMintedPoolRequest;
    return message;
  },

  toJSON(_: QueryGetMintedPoolRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetMintedPoolRequest>
  ): QueryGetMintedPoolRequest {
    const message = {
      ...baseQueryGetMintedPoolRequest,
    } as QueryGetMintedPoolRequest;
    return message;
  },
};

const baseQueryGetMintedPoolResponse: object = {};

export const QueryGetMintedPoolResponse = {
  encode(
    message: QueryGetMintedPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.MintedPool !== undefined) {
      MintedPool.encode(message.MintedPool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetMintedPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMintedPoolResponse,
    } as QueryGetMintedPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.MintedPool = MintedPool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMintedPoolResponse {
    const message = {
      ...baseQueryGetMintedPoolResponse,
    } as QueryGetMintedPoolResponse;
    if (object.MintedPool !== undefined && object.MintedPool !== null) {
      message.MintedPool = MintedPool.fromJSON(object.MintedPool);
    } else {
      message.MintedPool = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMintedPoolResponse): unknown {
    const obj: any = {};
    message.MintedPool !== undefined &&
      (obj.MintedPool = message.MintedPool
        ? MintedPool.toJSON(message.MintedPool)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMintedPoolResponse>
  ): QueryGetMintedPoolResponse {
    const message = {
      ...baseQueryGetMintedPoolResponse,
    } as QueryGetMintedPoolResponse;
    if (object.MintedPool !== undefined && object.MintedPool !== null) {
      message.MintedPool = MintedPool.fromPartial(object.MintedPool);
    } else {
      message.MintedPool = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a MintDenom by index. */
  MintDenom(
    request: QueryGetMintDenomRequest
  ): Promise<QueryGetMintDenomResponse>;
  /** Queries a StartBlock by index. */
  StartBlock(
    request: QueryGetStartBlockRequest
  ): Promise<QueryGetStartBlockResponse>;
  /** Queries a MintedPool by index. */
  MintedPool(
    request: QueryGetMintedPoolRequest
  ): Promise<QueryGetMintedPoolResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smartdev0328.bluechip.pageinflation.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  MintDenom(
    request: QueryGetMintDenomRequest
  ): Promise<QueryGetMintDenomResponse> {
    const data = QueryGetMintDenomRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smartdev0328.bluechip.pageinflation.Query",
      "MintDenom",
      data
    );
    return promise.then((data) =>
      QueryGetMintDenomResponse.decode(new Reader(data))
    );
  }

  StartBlock(
    request: QueryGetStartBlockRequest
  ): Promise<QueryGetStartBlockResponse> {
    const data = QueryGetStartBlockRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smartdev0328.bluechip.pageinflation.Query",
      "StartBlock",
      data
    );
    return promise.then((data) =>
      QueryGetStartBlockResponse.decode(new Reader(data))
    );
  }

  MintedPool(
    request: QueryGetMintedPoolRequest
  ): Promise<QueryGetMintedPoolResponse> {
    const data = QueryGetMintedPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "smartdev0328.bluechip.pageinflation.Query",
      "MintedPool",
      data
    );
    return promise.then((data) =>
      QueryGetMintedPoolResponse.decode(new Reader(data))
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
