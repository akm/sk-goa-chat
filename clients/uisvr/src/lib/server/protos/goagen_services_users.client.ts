// @generated by protobuf-ts 2.9.3
// @generated from protobuf file "goagen_services_users.proto" (package "users", syntax proto3)
// tslint:disable
//
// Code generated with goa v3.14.1, DO NOT EDIT.
//
// users protocol buffer definition
//
// Command:
// $ goa gen apisvr/design -o ./services
//
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { Users } from "./goagen_services_users";
import type { CreateResponse } from "./goagen_services_users";
import type { CreateRequest } from "./goagen_services_users";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListResponse } from "./goagen_services_users";
import type { ListRequest } from "./goagen_services_users";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * Service is the users service interface.
 *
 * @generated from protobuf service users.Users
 */
export interface IUsersClient {
    /**
     * List implements list.
     *
     * @generated from protobuf rpc: List(users.ListRequest) returns (users.ListResponse);
     */
    list(input: ListRequest, options?: RpcOptions): UnaryCall<ListRequest, ListResponse>;
    /**
     * Create implements create.
     *
     * @generated from protobuf rpc: Create(users.CreateRequest) returns (users.CreateResponse);
     */
    create(input: CreateRequest, options?: RpcOptions): UnaryCall<CreateRequest, CreateResponse>;
}
/**
 * Service is the users service interface.
 *
 * @generated from protobuf service users.Users
 */
export class UsersClient implements IUsersClient, ServiceInfo {
    typeName = Users.typeName;
    methods = Users.methods;
    options = Users.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * List implements list.
     *
     * @generated from protobuf rpc: List(users.ListRequest) returns (users.ListResponse);
     */
    list(input: ListRequest, options?: RpcOptions): UnaryCall<ListRequest, ListResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListRequest, ListResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Create implements create.
     *
     * @generated from protobuf rpc: Create(users.CreateRequest) returns (users.CreateResponse);
     */
    create(input: CreateRequest, options?: RpcOptions): UnaryCall<CreateRequest, CreateResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateRequest, CreateResponse>("unary", this._transport, method, opt, input);
    }
}
