// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file proto/gateway/authentication/v1/authentication.proto (package proto.gateway.authentication.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { LoginRequest, LoginResponse, LogoutRequest, LogoutResponse } from "./authentication_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Authentication service
 *
 * @generated from service proto.gateway.authentication.v1.AuthenticationService
 */
export const AuthenticationService = {
  typeName: "proto.gateway.authentication.v1.AuthenticationService",
  methods: {
    /**
     * @generated from rpc proto.gateway.authentication.v1.AuthenticationService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.gateway.authentication.v1.AuthenticationService.Logout
     */
    logout: {
      name: "Logout",
      I: LogoutRequest,
      O: LogoutResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
