// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages HTTP client encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	chatmessages "apisvr/services/gen/chat_messages"
	chatmessagesviews "apisvr/services/gen/chat_messages/views"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "chat_messages" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListChatMessagesPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat_messages", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the chat_messages
// list server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chatmessages.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("chat_messages", "list", "*chatmessages.ListPayload", v)
		}
		{
			v := p.SessionID
			req.AddCookie(&http.Cookie{
				Name:  "session_id",
				Value: v,
			})
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.ChannelID != nil {
			values.Add("channel_id", fmt.Sprintf("%v", *p.ChannelID))
		}
		if p.After != nil {
			values.Add("after", fmt.Sprintf("%v", *p.After))
		}
		if p.Before != nil {
			values.Add("before", fmt.Sprintf("%v", *p.Before))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the
// chat_messages list endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListResponse may return the following errors:
//   - "unauthenticated" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "list", err)
			}
			p := NewListChatMessageListOK(&body)
			view := "default"
			vres := &chatmessagesviews.ChatMessageList{Projected: p, View: view}
			if err = chatmessagesviews.ValidateChatMessageList(vres); err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "list", err)
			}
			res := chatmessages.NewChatMessageList(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListUnauthenticatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "list", err)
			}
			err = ValidateListUnauthenticatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "list", err)
			}
			return nil, NewListUnauthenticated(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat_messages", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "chat_messages" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*chatmessages.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chat_messages", "show", "*chatmessages.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowChatMessagesPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat_messages", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the chat_messages
// show server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chatmessages.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("chat_messages", "show", "*chatmessages.ShowPayload", v)
		}
		{
			v := p.SessionID
			req.AddCookie(&http.Cookie{
				Name:  "session_id",
				Value: v,
			})
		}
		return nil
	}
}

// DecodeShowResponse returns a decoder for responses returned by the
// chat_messages show endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeShowResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "unauthenticated" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "show", err)
			}
			p := NewShowChatMessageOK(&body)
			view := "default"
			vres := &chatmessagesviews.ChatMessage{Projected: p, View: view}
			if err = chatmessagesviews.ValidateChatMessage(vres); err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "show", err)
			}
			res := chatmessages.NewChatMessage(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "show", err)
			}
			err = ValidateShowNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "show", err)
			}
			return nil, NewShowNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ShowUnauthenticatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "show", err)
			}
			err = ValidateShowUnauthenticatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "show", err)
			}
			return nil, NewShowUnauthenticated(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat_messages", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "chat_messages" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateChatMessagesPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat_messages", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the
// chat_messages create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chatmessages.ChatMessageCreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("chat_messages", "create", "*chatmessages.ChatMessageCreatePayload", v)
		}
		{
			v := p.SessionID
			req.AddCookie(&http.Cookie{
				Name:  "session_id",
				Value: v,
			})
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chat_messages", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// chat_messages create endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "invalid_payload" (type *goa.ServiceError): http.StatusBadRequest
//   - "unauthenticated" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "create", err)
			}
			p := NewCreateChatMessageCreated(&body)
			view := "default"
			vres := &chatmessagesviews.ChatMessage{Projected: p, View: view}
			if err = chatmessagesviews.ValidateChatMessage(vres); err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "create", err)
			}
			res := chatmessages.NewChatMessage(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateInvalidPayloadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "create", err)
			}
			err = ValidateCreateInvalidPayloadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "create", err)
			}
			return nil, NewCreateInvalidPayload(&body)
		case http.StatusUnauthorized:
			var (
				body CreateUnauthenticatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "create", err)
			}
			err = ValidateCreateUnauthenticatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "create", err)
			}
			return nil, NewCreateUnauthenticated(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat_messages", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "chat_messages" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*chatmessages.ChatMessageUpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chat_messages", "update", "*chatmessages.ChatMessageUpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateChatMessagesPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat_messages", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the
// chat_messages update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chatmessages.ChatMessageUpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("chat_messages", "update", "*chatmessages.ChatMessageUpdatePayload", v)
		}
		{
			v := p.SessionID
			req.AddCookie(&http.Cookie{
				Name:  "session_id",
				Value: v,
			})
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chat_messages", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// chat_messages update endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "invalid_payload" (type *goa.ServiceError): http.StatusBadRequest
//   - "unauthenticated" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "update", err)
			}
			p := NewUpdateChatMessageOK(&body)
			view := "default"
			vres := &chatmessagesviews.ChatMessage{Projected: p, View: view}
			if err = chatmessagesviews.ValidateChatMessage(vres); err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "update", err)
			}
			res := chatmessages.NewChatMessage(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		case http.StatusBadRequest:
			var (
				body UpdateInvalidPayloadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "update", err)
			}
			err = ValidateUpdateInvalidPayloadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "update", err)
			}
			return nil, NewUpdateInvalidPayload(&body)
		case http.StatusUnauthorized:
			var (
				body UpdateUnauthenticatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "update", err)
			}
			err = ValidateUpdateUnauthenticatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "update", err)
			}
			return nil, NewUpdateUnauthenticated(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat_messages", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "chat_messages" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*chatmessages.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chat_messages", "delete", "*chatmessages.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteChatMessagesPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat_messages", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the
// chat_messages delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chatmessages.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("chat_messages", "delete", "*chatmessages.DeletePayload", v)
		}
		{
			v := p.SessionID
			req.AddCookie(&http.Cookie{
				Name:  "session_id",
				Value: v,
			})
		}
		return nil
	}
}

// DecodeDeleteResponse returns a decoder for responses returned by the
// chat_messages delete endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "unauthenticated" (type *goa.ServiceError): http.StatusUnauthorized
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "delete", err)
			}
			p := NewDeleteChatMessageOK(&body)
			view := "default"
			vres := &chatmessagesviews.ChatMessage{Projected: p, View: view}
			if err = chatmessagesviews.ValidateChatMessage(vres); err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "delete", err)
			}
			res := chatmessages.NewChatMessage(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body DeleteNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "delete", err)
			}
			err = ValidateDeleteNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "delete", err)
			}
			return nil, NewDeleteNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body DeleteUnauthenticatedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat_messages", "delete", err)
			}
			err = ValidateDeleteUnauthenticatedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat_messages", "delete", err)
			}
			return nil, NewDeleteUnauthenticated(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat_messages", "delete", resp.StatusCode, string(body))
		}
	}
}

// unmarshalChatMessageListItemResponseBodyToChatmessagesviewsChatMessageListItemView
// builds a value of type *chatmessagesviews.ChatMessageListItemView from a
// value of type *ChatMessageListItemResponseBody.
func unmarshalChatMessageListItemResponseBodyToChatmessagesviewsChatMessageListItemView(v *ChatMessageListItemResponseBody) *chatmessagesviews.ChatMessageListItemView {
	res := &chatmessagesviews.ChatMessageListItemView{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		ChannelID: v.ChannelID,
		UserID:    v.UserID,
		UserName:  v.UserName,
		Content:   v.Content,
	}

	return res
}
