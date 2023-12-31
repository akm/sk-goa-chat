// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels gRPC client
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	channelspb "apisvr/services/gen/grpc/channels/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli channelspb.ChannelsClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the channels service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: channelspb.NewChannelsClient(cc),
		opts:    opts,
	}
}

// List calls the "List" function in channelspb.ChannelsClient interface.
func (c *Client) List() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildListFunc(c.grpccli, c.opts...),
			nil,
			DecodeListResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Show calls the "Show" function in channelspb.ChannelsClient interface.
func (c *Client) Show() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildShowFunc(c.grpccli, c.opts...),
			EncodeShowRequest,
			DecodeShowResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Create calls the "Create" function in channelspb.ChannelsClient interface.
func (c *Client) Create() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildCreateFunc(c.grpccli, c.opts...),
			EncodeCreateRequest,
			DecodeCreateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Update calls the "Update" function in channelspb.ChannelsClient interface.
func (c *Client) Update() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildUpdateFunc(c.grpccli, c.opts...),
			EncodeUpdateRequest,
			DecodeUpdateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Delete calls the "Delete" function in channelspb.ChannelsClient interface.
func (c *Client) Delete() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildDeleteFunc(c.grpccli, c.opts...),
			EncodeDeleteRequest,
			DecodeDeleteResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
