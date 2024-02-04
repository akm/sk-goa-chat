// Code generated by goa v3.14.1, DO NOT EDIT.
//
// sessions gRPC client encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	sessionspb "apisvr/services/gen/grpc/sessions/pb"
	sessions "apisvr/services/gen/sessions"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateFunc builds the remote method to invoke for "sessions" service
// "create" endpoint.
func BuildCreateFunc(grpccli sessionspb.SessionsClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Create(ctx, reqpb.(*sessionspb.CreateRequest), opts...)
		}
		return grpccli.Create(ctx, &sessionspb.CreateRequest{}, opts...)
	}
}

// EncodeCreateRequest encodes requests sent to sessions create endpoint.
func EncodeCreateRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*sessions.CreatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sessions", "create", "*sessions.CreatePayload", v)
	}
	return NewProtoCreateRequest(payload), nil
}

// DecodeCreateResponse decodes responses from the sessions create endpoint.
func DecodeCreateResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*sessionspb.CreateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sessions", "create", "*sessionspb.CreateResponse", v)
	}
	res := NewCreateResult(message)
	return res, nil
}

// BuildDeleteFunc builds the remote method to invoke for "sessions" service
// "delete" endpoint.
func BuildDeleteFunc(grpccli sessionspb.SessionsClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Delete(ctx, reqpb.(*sessionspb.DeleteRequest), opts...)
		}
		return grpccli.Delete(ctx, &sessionspb.DeleteRequest{}, opts...)
	}
}

// EncodeDeleteRequest encodes requests sent to sessions delete endpoint.
func EncodeDeleteRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*sessions.DeletePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sessions", "delete", "*sessions.DeletePayload", v)
	}
	return NewProtoDeleteRequest(payload), nil
}

// DecodeDeleteResponse decodes responses from the sessions delete endpoint.
func DecodeDeleteResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*sessionspb.DeleteResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sessions", "delete", "*sessionspb.DeleteResponse", v)
	}
	res := NewDeleteResult(message)
	return res, nil
}