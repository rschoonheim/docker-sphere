import * as jspb from 'google-protobuf'



export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class GetEnvironmentHealthResponse extends jspb.Message {
  getDockerStatus(): string;
  setDockerStatus(value: string): GetEnvironmentHealthResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetEnvironmentHealthResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetEnvironmentHealthResponse): GetEnvironmentHealthResponse.AsObject;
  static serializeBinaryToWriter(message: GetEnvironmentHealthResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetEnvironmentHealthResponse;
  static deserializeBinaryFromReader(message: GetEnvironmentHealthResponse, reader: jspb.BinaryReader): GetEnvironmentHealthResponse;
}

export namespace GetEnvironmentHealthResponse {
  export type AsObject = {
    dockerStatus: string,
  }
}

