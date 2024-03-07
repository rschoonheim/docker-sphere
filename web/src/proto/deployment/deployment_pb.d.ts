import * as jspb from 'google-protobuf'



export class Deployment extends jspb.Message {
  getId(): string;
  setId(value: string): Deployment;

  getName(): string;
  setName(value: string): Deployment;

  getStatus(): number;
  setStatus(value: number): Deployment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Deployment.AsObject;
  static toObject(includeInstance: boolean, msg: Deployment): Deployment.AsObject;
  static serializeBinaryToWriter(message: Deployment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Deployment;
  static deserializeBinaryFromReader(message: Deployment, reader: jspb.BinaryReader): Deployment;
}

export namespace Deployment {
  export type AsObject = {
    id: string,
    name: string,
    status: number,
  }
}

export class CreateDeploymentRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateDeploymentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDeploymentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDeploymentRequest): CreateDeploymentRequest.AsObject;
  static serializeBinaryToWriter(message: CreateDeploymentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDeploymentRequest;
  static deserializeBinaryFromReader(message: CreateDeploymentRequest, reader: jspb.BinaryReader): CreateDeploymentRequest;
}

export namespace CreateDeploymentRequest {
  export type AsObject = {
    name: string,
  }
}

export class CreateDeploymentResponse extends jspb.Message {
  getId(): string;
  setId(value: string): CreateDeploymentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateDeploymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateDeploymentResponse): CreateDeploymentResponse.AsObject;
  static serializeBinaryToWriter(message: CreateDeploymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateDeploymentResponse;
  static deserializeBinaryFromReader(message: CreateDeploymentResponse, reader: jspb.BinaryReader): CreateDeploymentResponse;
}

export namespace CreateDeploymentResponse {
  export type AsObject = {
    id: string,
  }
}

export class GetDeploymentRequest extends jspb.Message {
  getId(): string;
  setId(value: string): GetDeploymentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeploymentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeploymentRequest): GetDeploymentRequest.AsObject;
  static serializeBinaryToWriter(message: GetDeploymentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeploymentRequest;
  static deserializeBinaryFromReader(message: GetDeploymentRequest, reader: jspb.BinaryReader): GetDeploymentRequest;
}

export namespace GetDeploymentRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetDeploymentResponse extends jspb.Message {
  getId(): string;
  setId(value: string): GetDeploymentResponse;

  getName(): string;
  setName(value: string): GetDeploymentResponse;

  getStatus(): number;
  setStatus(value: number): GetDeploymentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeploymentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeploymentResponse): GetDeploymentResponse.AsObject;
  static serializeBinaryToWriter(message: GetDeploymentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeploymentResponse;
  static deserializeBinaryFromReader(message: GetDeploymentResponse, reader: jspb.BinaryReader): GetDeploymentResponse;
}

export namespace GetDeploymentResponse {
  export type AsObject = {
    id: string,
    name: string,
    status: number,
  }
}

export class ListDeploymentsRequest extends jspb.Message {
  getLimit(): number;
  setLimit(value: number): ListDeploymentsRequest;

  getOffset(): number;
  setOffset(value: number): ListDeploymentsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDeploymentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDeploymentsRequest): ListDeploymentsRequest.AsObject;
  static serializeBinaryToWriter(message: ListDeploymentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDeploymentsRequest;
  static deserializeBinaryFromReader(message: ListDeploymentsRequest, reader: jspb.BinaryReader): ListDeploymentsRequest;
}

export namespace ListDeploymentsRequest {
  export type AsObject = {
    limit: number,
    offset: number,
  }
}

export class ListDeploymentsResponse extends jspb.Message {
  getDeploymentsList(): Array<Deployment>;
  setDeploymentsList(value: Array<Deployment>): ListDeploymentsResponse;
  clearDeploymentsList(): ListDeploymentsResponse;
  addDeployments(value?: Deployment, index?: number): Deployment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDeploymentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDeploymentsResponse): ListDeploymentsResponse.AsObject;
  static serializeBinaryToWriter(message: ListDeploymentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDeploymentsResponse;
  static deserializeBinaryFromReader(message: ListDeploymentsResponse, reader: jspb.BinaryReader): ListDeploymentsResponse;
}

export namespace ListDeploymentsResponse {
  export type AsObject = {
    deploymentsList: Array<Deployment.AsObject>,
  }
}

