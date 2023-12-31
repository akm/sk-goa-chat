// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels gRPC client types
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	channels "apisvr/services/gen/channels"
	channelsviews "apisvr/services/gen/channels/views"
	channelspb "apisvr/services/gen/grpc/channels/pb"

	goa "goa.design/goa/v3/pkg"
)

// NewProtoListRequest builds the gRPC request type from the payload of the
// "list" endpoint of the "channels" service.
func NewProtoListRequest() *channelspb.ListRequest {
	message := &channelspb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the
// "channels" service from the gRPC response type.
func NewListResult(message *channelspb.ListResponse) *channelsviews.ChannelListView {
	result := &channelsviews.ChannelListView{
		Total:  &message.Total,
		Offset: &message.Offset,
	}
	if message.Items != nil {
		result.Items = make([]*channelsviews.ChannelListItemView, len(message.Items.Field))
		for i, val := range message.Items.Field {
			result.Items[i] = &channelsviews.ChannelListItemView{
				ID:        &val.Id,
				CreatedAt: &val.CreatedAt,
				UpdatedAt: &val.UpdatedAt,
				Name:      &val.Name,
			}
		}
	}
	return result
}

// NewProtoShowRequest builds the gRPC request type from the payload of the
// "show" endpoint of the "channels" service.
func NewProtoShowRequest(payload *channels.ShowPayload) *channelspb.ShowRequest {
	message := &channelspb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the
// "channels" service from the gRPC response type.
func NewShowResult(message *channelspb.ShowResponse) *channelsviews.ChannelView {
	result := &channelsviews.ChannelView{
		ID:        &message.Id,
		CreatedAt: &message.CreatedAt,
		UpdatedAt: &message.UpdatedAt,
		Name:      &message.Name,
	}
	return result
}

// NewProtoCreateRequest builds the gRPC request type from the payload of the
// "create" endpoint of the "channels" service.
func NewProtoCreateRequest(payload *channels.ChannelCreatePayload) *channelspb.CreateRequest {
	message := &channelspb.CreateRequest{
		Name: payload.Name,
	}
	return message
}

// NewCreateResult builds the result type of the "create" endpoint of the
// "channels" service from the gRPC response type.
func NewCreateResult(message *channelspb.CreateResponse) *channelsviews.ChannelView {
	result := &channelsviews.ChannelView{
		ID:        &message.Id,
		CreatedAt: &message.CreatedAt,
		UpdatedAt: &message.UpdatedAt,
		Name:      &message.Name,
	}
	return result
}

// NewProtoUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "channels" service.
func NewProtoUpdateRequest(payload *channels.ChannelUpdatePayload) *channelspb.UpdateRequest {
	message := &channelspb.UpdateRequest{
		Id:   payload.ID,
		Name: payload.Name,
	}
	return message
}

// NewUpdateResult builds the result type of the "update" endpoint of the
// "channels" service from the gRPC response type.
func NewUpdateResult(message *channelspb.UpdateResponse) *channelsviews.ChannelView {
	result := &channelsviews.ChannelView{
		ID:        &message.Id,
		CreatedAt: &message.CreatedAt,
		UpdatedAt: &message.UpdatedAt,
		Name:      &message.Name,
	}
	return result
}

// NewProtoDeleteRequest builds the gRPC request type from the payload of the
// "delete" endpoint of the "channels" service.
func NewProtoDeleteRequest(payload *channels.DeletePayload) *channelspb.DeleteRequest {
	message := &channelspb.DeleteRequest{
		Id: payload.ID,
	}
	return message
}

// NewDeleteResult builds the result type of the "delete" endpoint of the
// "channels" service from the gRPC response type.
func NewDeleteResult(message *channelspb.DeleteResponse) *channelsviews.ChannelView {
	result := &channelsviews.ChannelView{
		ID:        &message.Id,
		CreatedAt: &message.CreatedAt,
		UpdatedAt: &message.UpdatedAt,
		Name:      &message.Name,
	}
	return result
}

// ValidateListResponse runs the validations defined on ListResponse.
func ValidateListResponse(message *channelspb.ListResponse) (err error) {
	if message.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "message"))
	}
	if message.Items != nil {
		if err2 := ValidateChannelListItemCollection(message.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChannelListItemCollection runs the validations defined on
// ChannelListItemCollection.
func ValidateChannelListItemCollection(items *channelspb.ChannelListItemCollection) (err error) {
	for _, e := range items.Field {
		if e != nil {
			if err2 := ValidateChannelListItem(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChannelListItem runs the validations defined on ChannelListItem.
func ValidateChannelListItem(elem *channelspb.ChannelListItem) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("elem.created_at", elem.CreatedAt, goa.FormatDateTime))
	err = goa.MergeErrors(err, goa.ValidateFormat("elem.updated_at", elem.UpdatedAt, goa.FormatDateTime))
	return
}

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *channelspb.ShowResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created_at", message.CreatedAt, goa.FormatDateTime))
	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated_at", message.UpdatedAt, goa.FormatDateTime))
	return
}

// ValidateCreateResponse runs the validations defined on CreateResponse.
func ValidateCreateResponse(message *channelspb.CreateResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created_at", message.CreatedAt, goa.FormatDateTime))
	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated_at", message.UpdatedAt, goa.FormatDateTime))
	return
}

// ValidateUpdateResponse runs the validations defined on UpdateResponse.
func ValidateUpdateResponse(message *channelspb.UpdateResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created_at", message.CreatedAt, goa.FormatDateTime))
	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated_at", message.UpdatedAt, goa.FormatDateTime))
	return
}

// ValidateDeleteResponse runs the validations defined on DeleteResponse.
func ValidateDeleteResponse(message *channelspb.DeleteResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.created_at", message.CreatedAt, goa.FormatDateTime))
	err = goa.MergeErrors(err, goa.ValidateFormat("message.updated_at", message.UpdatedAt, goa.FormatDateTime))
	return
}
