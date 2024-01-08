// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels gRPC server types
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	channels "apisvr/services/gen/channels"
	channelsviews "apisvr/services/gen/channels/views"
	channelspb "apisvr/services/gen/grpc/channels/pb"
)

// NewListPayload builds the payload of the "list" endpoint of the "channels"
// service from the gRPC request type.
func NewListPayload(message *channelspb.ListRequest) *channels.ListPayload {
	v := &channels.ListPayload{
		SessionID: message.SessionId,
	}
	return v
}

// NewProtoListResponse builds the gRPC response type from the result of the
// "list" endpoint of the "channels" service.
func NewProtoListResponse(result *channelsviews.ChannelListView) *channelspb.ListResponse {
	message := &channelspb.ListResponse{
		Total:  *result.Total,
		Offset: *result.Offset,
	}
	if result.Items != nil {
		message.Items = &channelspb.ChannelListItemCollection{}
		message.Items.Field = make([]*channelspb.ChannelListItem, len(result.Items))
		for i, val := range result.Items {
			message.Items.Field[i] = &channelspb.ChannelListItem{
				Id:        *val.ID,
				CreatedAt: *val.CreatedAt,
				UpdatedAt: *val.UpdatedAt,
				Name:      *val.Name,
			}
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the "channels"
// service from the gRPC request type.
func NewShowPayload(message *channelspb.ShowRequest) *channels.ShowPayload {
	v := &channels.ShowPayload{
		SessionID: message.SessionId,
		ID:        message.Id,
	}
	return v
}

// NewProtoShowResponse builds the gRPC response type from the result of the
// "show" endpoint of the "channels" service.
func NewProtoShowResponse(result *channelsviews.ChannelView) *channelspb.ShowResponse {
	message := &channelspb.ShowResponse{
		Id:        *result.ID,
		CreatedAt: *result.CreatedAt,
		UpdatedAt: *result.UpdatedAt,
		Name:      *result.Name,
	}
	return message
}

// NewCreatePayload builds the payload of the "create" endpoint of the
// "channels" service from the gRPC request type.
func NewCreatePayload(message *channelspb.CreateRequest) *channels.ChannelCreatePayload {
	v := &channels.ChannelCreatePayload{
		SessionID: message.SessionId,
		Name:      message.Name,
	}
	return v
}

// NewProtoCreateResponse builds the gRPC response type from the result of the
// "create" endpoint of the "channels" service.
func NewProtoCreateResponse(result *channelsviews.ChannelView) *channelspb.CreateResponse {
	message := &channelspb.CreateResponse{
		Id:        *result.ID,
		CreatedAt: *result.CreatedAt,
		UpdatedAt: *result.UpdatedAt,
		Name:      *result.Name,
	}
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the
// "channels" service from the gRPC request type.
func NewUpdatePayload(message *channelspb.UpdateRequest) *channels.ChannelUpdatePayload {
	v := &channels.ChannelUpdatePayload{
		SessionID: message.SessionId,
		ID:        message.Id,
		Name:      message.Name,
	}
	return v
}

// NewProtoUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "channels" service.
func NewProtoUpdateResponse(result *channelsviews.ChannelView) *channelspb.UpdateResponse {
	message := &channelspb.UpdateResponse{
		Id:        *result.ID,
		CreatedAt: *result.CreatedAt,
		UpdatedAt: *result.UpdatedAt,
		Name:      *result.Name,
	}
	return message
}

// NewDeletePayload builds the payload of the "delete" endpoint of the
// "channels" service from the gRPC request type.
func NewDeletePayload(message *channelspb.DeleteRequest) *channels.DeletePayload {
	v := &channels.DeletePayload{
		SessionID: message.SessionId,
		ID:        message.Id,
	}
	return v
}

// NewProtoDeleteResponse builds the gRPC response type from the result of the
// "delete" endpoint of the "channels" service.
func NewProtoDeleteResponse(result *channelsviews.ChannelView) *channelspb.DeleteResponse {
	message := &channelspb.DeleteResponse{
		Id:        *result.ID,
		CreatedAt: *result.CreatedAt,
		UpdatedAt: *result.UpdatedAt,
		Name:      *result.Name,
	}
	return message
}
