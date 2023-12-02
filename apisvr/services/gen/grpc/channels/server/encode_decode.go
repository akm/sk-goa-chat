// Code generated by goa v3.14.0, DO NOT EDIT.
//
// channels gRPC server encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	channels "apisvr/services/gen/channels"
	channelsviews "apisvr/services/gen/channels/views"
	channelspb "apisvr/services/gen/grpc/channels/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeListResponse encodes responses from the "channels" service "list"
// endpoint.
func EncodeListResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*channelsviews.ChannelList)
	if !ok {
		return nil, goagrpc.ErrInvalidType("channels", "list", "*channelsviews.ChannelList", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoListResponse(result)
	return resp, nil
}

// EncodeShowResponse encodes responses from the "channels" service "show"
// endpoint.
func EncodeShowResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*channelsviews.Channel)
	if !ok {
		return nil, goagrpc.ErrInvalidType("channels", "show", "*channelsviews.Channel", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoShowResponse(result)
	return resp, nil
}

// DecodeShowRequest decodes requests sent to "channels" service "show"
// endpoint.
func DecodeShowRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *channelspb.ShowRequest
		ok      bool
	)
	{
		if message, ok = v.(*channelspb.ShowRequest); !ok {
			return nil, goagrpc.ErrInvalidType("channels", "show", "*channelspb.ShowRequest", v)
		}
	}
	var payload *channels.ShowPayload
	{
		payload = NewShowPayload(message)
	}
	return payload, nil
}

// EncodeCreateResponse encodes responses from the "channels" service "create"
// endpoint.
func EncodeCreateResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*channelsviews.Channel)
	if !ok {
		return nil, goagrpc.ErrInvalidType("channels", "create", "*channelsviews.Channel", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoCreateResponse(result)
	return resp, nil
}

// DecodeCreateRequest decodes requests sent to "channels" service "create"
// endpoint.
func DecodeCreateRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *channelspb.CreateRequest
		ok      bool
	)
	{
		if message, ok = v.(*channelspb.CreateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("channels", "create", "*channelspb.CreateRequest", v)
		}
	}
	var payload *channels.ChannelCreatePayload
	{
		payload = NewCreatePayload(message)
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "channels" service "update"
// endpoint.
func EncodeUpdateResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*channelsviews.Channel)
	if !ok {
		return nil, goagrpc.ErrInvalidType("channels", "update", "*channelsviews.Channel", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoUpdateResponse(result)
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "channels" service "update"
// endpoint.
func DecodeUpdateRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *channelspb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*channelspb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("channels", "update", "*channelspb.UpdateRequest", v)
		}
	}
	var payload *channels.ChannelUpdatePayload
	{
		payload = NewUpdatePayload(message)
	}
	return payload, nil
}

// EncodeDeleteResponse encodes responses from the "channels" service "delete"
// endpoint.
func EncodeDeleteResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*channelsviews.Channel)
	if !ok {
		return nil, goagrpc.ErrInvalidType("channels", "delete", "*channelsviews.Channel", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoDeleteResponse(result)
	return resp, nil
}

// DecodeDeleteRequest decodes requests sent to "channels" service "delete"
// endpoint.
func DecodeDeleteRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *channelspb.DeleteRequest
		ok      bool
	)
	{
		if message, ok = v.(*channelspb.DeleteRequest); !ok {
			return nil, goagrpc.ErrInvalidType("channels", "delete", "*channelspb.DeleteRequest", v)
		}
	}
	var payload *channels.DeletePayload
	{
		payload = NewDeletePayload(message)
	}
	return payload, nil
}
