// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file api/user/user.proto (package docker_sphere_gateway.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { User } from "./user_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service docker_sphere_gateway.v1.UserService
 */
export const UserService = {
  typeName: "docker_sphere_gateway.v1.UserService",
  methods: {
    /**
     * @generated from rpc docker_sphere_gateway.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: User,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_sphere_gateway.v1.UserService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: User,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_sphere_gateway.v1.UserService.UpdateUser
     */
    updateUser: {
      name: "UpdateUser",
      I: User,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc docker_sphere_gateway.v1.UserService.DeleteUser
     */
    deleteUser: {
      name: "DeleteUser",
      I: User,
      O: User,
      kind: MethodKind.Unary,
    },
  }
} as const;

