// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages gRPC server encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	chatmessages "apisvr/services/gen/chat_messages"
	chatmessagesviews "apisvr/services/gen/chat_messages/views"
	chat_messagespb "apisvr/services/gen/grpc/chat_messages/pb"
	"context"
	"strings"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeListResponse encodes responses from the "chat_messages" service "list"
// endpoint.
func EncodeListResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*chatmessagesviews.ChatMessageList)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chat_messages", "list", "*chatmessagesviews.ChatMessageList", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoListResponse(result)
	return resp, nil
}

// DecodeListRequest decodes requests sent to "chat_messages" service "list"
// endpoint.
func DecodeListRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		idToken string
		err     error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			idToken = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chat_messagespb.ListRequest
		ok      bool
	)
	{
		if message, ok = v.(*chat_messagespb.ListRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chat_messages", "list", "*chat_messagespb.ListRequest", v)
		}
	}
	var payload *chatmessages.ListPayload
	{
		payload = NewListPayload(message, idToken)
		if strings.Contains(payload.IDToken, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.IDToken, " ", 2)[1]
			payload.IDToken = cred
		}
	}
	return payload, nil
}

// EncodeShowResponse encodes responses from the "chat_messages" service "show"
// endpoint.
func EncodeShowResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*chatmessagesviews.ChatMessage)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chat_messages", "show", "*chatmessagesviews.ChatMessage", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoShowResponse(result)
	return resp, nil
}

// DecodeShowRequest decodes requests sent to "chat_messages" service "show"
// endpoint.
func DecodeShowRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		idToken string
		err     error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			idToken = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chat_messagespb.ShowRequest
		ok      bool
	)
	{
		if message, ok = v.(*chat_messagespb.ShowRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chat_messages", "show", "*chat_messagespb.ShowRequest", v)
		}
	}
	var payload *chatmessages.ShowPayload
	{
		payload = NewShowPayload(message, idToken)
		if strings.Contains(payload.IDToken, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.IDToken, " ", 2)[1]
			payload.IDToken = cred
		}
	}
	return payload, nil
}

// EncodeCreateResponse encodes responses from the "chat_messages" service
// "create" endpoint.
func EncodeCreateResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*chatmessagesviews.ChatMessage)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chat_messages", "create", "*chatmessagesviews.ChatMessage", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoCreateResponse(result)
	return resp, nil
}

// DecodeCreateRequest decodes requests sent to "chat_messages" service
// "create" endpoint.
func DecodeCreateRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		idToken string
		err     error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			idToken = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chat_messagespb.CreateRequest
		ok      bool
	)
	{
		if message, ok = v.(*chat_messagespb.CreateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chat_messages", "create", "*chat_messagespb.CreateRequest", v)
		}
	}
	var payload *chatmessages.ChatMessageCreatePayload
	{
		payload = NewCreatePayload(message, idToken)
		if strings.Contains(payload.IDToken, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.IDToken, " ", 2)[1]
			payload.IDToken = cred
		}
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "chat_messages" service
// "update" endpoint.
func EncodeUpdateResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*chatmessagesviews.ChatMessage)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chat_messages", "update", "*chatmessagesviews.ChatMessage", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoUpdateResponse(result)
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "chat_messages" service
// "update" endpoint.
func DecodeUpdateRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		idToken string
		err     error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			idToken = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chat_messagespb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*chat_messagespb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chat_messages", "update", "*chat_messagespb.UpdateRequest", v)
		}
	}
	var payload *chatmessages.ChatMessageUpdatePayload
	{
		payload = NewUpdatePayload(message, idToken)
		if strings.Contains(payload.IDToken, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.IDToken, " ", 2)[1]
			payload.IDToken = cred
		}
	}
	return payload, nil
}

// EncodeDeleteResponse encodes responses from the "chat_messages" service
// "delete" endpoint.
func EncodeDeleteResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*chatmessagesviews.ChatMessage)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chat_messages", "delete", "*chatmessagesviews.ChatMessage", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoDeleteResponse(result)
	return resp, nil
}

// DecodeDeleteRequest decodes requests sent to "chat_messages" service
// "delete" endpoint.
func DecodeDeleteRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		idToken string
		err     error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			idToken = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *chat_messagespb.DeleteRequest
		ok      bool
	)
	{
		if message, ok = v.(*chat_messagespb.DeleteRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chat_messages", "delete", "*chat_messagespb.DeleteRequest", v)
		}
	}
	var payload *chatmessages.DeletePayload
	{
		payload = NewDeletePayload(message, idToken)
		if strings.Contains(payload.IDToken, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.IDToken, " ", 2)[1]
			payload.IDToken = cred
		}
	}
	return payload, nil
}
