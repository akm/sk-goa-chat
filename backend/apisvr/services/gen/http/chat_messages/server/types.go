// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages HTTP server types
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	chatmessages "apisvr/services/gen/chat_messages"
	chatmessagesviews "apisvr/services/gen/chat_messages/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "chat_messages" service "create"
// endpoint HTTP request body.
type CreateRequestBody struct {
	// Channel ID
	ChannelID *uint64 `form:"channel_id,omitempty" json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// Content
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
}

// UpdateRequestBody is the type of the "chat_messages" service "update"
// endpoint HTTP request body.
type UpdateRequestBody struct {
	// Content
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
}

// ListResponseBody is the type of the "chat_messages" service "list" endpoint
// HTTP response body.
type ListResponseBody struct {
	// Items
	Items ChatMessageListItemResponseBodyCollection `form:"items" json:"items" xml:"items"`
}

// ShowResponseBody is the type of the "chat_messages" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Channel ID
	ChannelID uint64 `form:"channel_id" json:"channel_id" xml:"channel_id"`
	// User ID
	UserID *uint64 `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// Content
	Content string `form:"content" json:"content" xml:"content"`
}

// CreateResponseBody is the type of the "chat_messages" service "create"
// endpoint HTTP response body.
type CreateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Channel ID
	ChannelID uint64 `form:"channel_id" json:"channel_id" xml:"channel_id"`
	// User ID
	UserID *uint64 `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// Content
	Content string `form:"content" json:"content" xml:"content"`
}

// UpdateResponseBody is the type of the "chat_messages" service "update"
// endpoint HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Channel ID
	ChannelID uint64 `form:"channel_id" json:"channel_id" xml:"channel_id"`
	// User ID
	UserID *uint64 `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// Content
	Content string `form:"content" json:"content" xml:"content"`
}

// DeleteResponseBody is the type of the "chat_messages" service "delete"
// endpoint HTTP response body.
type DeleteResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Channel ID
	ChannelID uint64 `form:"channel_id" json:"channel_id" xml:"channel_id"`
	// User ID
	UserID *uint64 `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// Content
	Content string `form:"content" json:"content" xml:"content"`
}

// ListUnauthenticatedResponseBody is the type of the "chat_messages" service
// "list" endpoint HTTP response body for the "unauthenticated" error.
type ListUnauthenticatedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowNotFoundResponseBody is the type of the "chat_messages" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowUnauthenticatedResponseBody is the type of the "chat_messages" service
// "show" endpoint HTTP response body for the "unauthenticated" error.
type ShowUnauthenticatedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateInvalidPayloadResponseBody is the type of the "chat_messages" service
// "create" endpoint HTTP response body for the "invalid_payload" error.
type CreateInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateUnauthenticatedResponseBody is the type of the "chat_messages" service
// "create" endpoint HTTP response body for the "unauthenticated" error.
type CreateUnauthenticatedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotFoundResponseBody is the type of the "chat_messages" service
// "update" endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInvalidPayloadResponseBody is the type of the "chat_messages" service
// "update" endpoint HTTP response body for the "invalid_payload" error.
type UpdateInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateUnauthenticatedResponseBody is the type of the "chat_messages" service
// "update" endpoint HTTP response body for the "unauthenticated" error.
type UpdateUnauthenticatedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotFoundResponseBody is the type of the "chat_messages" service
// "delete" endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteUnauthenticatedResponseBody is the type of the "chat_messages" service
// "delete" endpoint HTTP response body for the "unauthenticated" error.
type DeleteUnauthenticatedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ChatMessageListItemResponseBodyCollection is used to define fields on
// response body types.
type ChatMessageListItemResponseBodyCollection []*ChatMessageListItemResponseBody

// ChatMessageListItemResponseBody is used to define fields on response body
// types.
type ChatMessageListItemResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Channel ID
	ChannelID uint64 `form:"channel_id" json:"channel_id" xml:"channel_id"`
	// User ID
	UserID *uint64 `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// User Name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// Content
	Content string `form:"content" json:"content" xml:"content"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "chat_messages" service.
func NewListResponseBody(res *chatmessagesviews.ChatMessageListView) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Items != nil {
		body.Items = make([]*ChatMessageListItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalChatmessagesviewsChatMessageListItemViewToChatMessageListItemResponseBody(val)
		}
	} else {
		body.Items = []*ChatMessageListItemResponseBody{}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "chat_messages" service.
func NewShowResponseBody(res *chatmessagesviews.ChatMessageView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		ChannelID: *res.ChannelID,
		UserID:    res.UserID,
		UserName:  *res.UserName,
		Content:   *res.Content,
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "chat_messages" service.
func NewCreateResponseBody(res *chatmessagesviews.ChatMessageView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		ChannelID: *res.ChannelID,
		UserID:    res.UserID,
		UserName:  *res.UserName,
		Content:   *res.Content,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "chat_messages" service.
func NewUpdateResponseBody(res *chatmessagesviews.ChatMessageView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		ChannelID: *res.ChannelID,
		UserID:    res.UserID,
		UserName:  *res.UserName,
		Content:   *res.Content,
	}
	return body
}

// NewDeleteResponseBody builds the HTTP response body from the result of the
// "delete" endpoint of the "chat_messages" service.
func NewDeleteResponseBody(res *chatmessagesviews.ChatMessageView) *DeleteResponseBody {
	body := &DeleteResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		ChannelID: *res.ChannelID,
		UserID:    res.UserID,
		UserName:  *res.UserName,
		Content:   *res.Content,
	}
	return body
}

// NewListUnauthenticatedResponseBody builds the HTTP response body from the
// result of the "list" endpoint of the "chat_messages" service.
func NewListUnauthenticatedResponseBody(res *goa.ServiceError) *ListUnauthenticatedResponseBody {
	body := &ListUnauthenticatedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "chat_messages" service.
func NewShowNotFoundResponseBody(res *goa.ServiceError) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowUnauthenticatedResponseBody builds the HTTP response body from the
// result of the "show" endpoint of the "chat_messages" service.
func NewShowUnauthenticatedResponseBody(res *goa.ServiceError) *ShowUnauthenticatedResponseBody {
	body := &ShowUnauthenticatedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateInvalidPayloadResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "chat_messages" service.
func NewCreateInvalidPayloadResponseBody(res *goa.ServiceError) *CreateInvalidPayloadResponseBody {
	body := &CreateInvalidPayloadResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateUnauthenticatedResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "chat_messages" service.
func NewCreateUnauthenticatedResponseBody(res *goa.ServiceError) *CreateUnauthenticatedResponseBody {
	body := &CreateUnauthenticatedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "chat_messages" service.
func NewUpdateNotFoundResponseBody(res *goa.ServiceError) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInvalidPayloadResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "chat_messages" service.
func NewUpdateInvalidPayloadResponseBody(res *goa.ServiceError) *UpdateInvalidPayloadResponseBody {
	body := &UpdateInvalidPayloadResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateUnauthenticatedResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "chat_messages" service.
func NewUpdateUnauthenticatedResponseBody(res *goa.ServiceError) *UpdateUnauthenticatedResponseBody {
	body := &UpdateUnauthenticatedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "chat_messages" service.
func NewDeleteNotFoundResponseBody(res *goa.ServiceError) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteUnauthenticatedResponseBody builds the HTTP response body from the
// result of the "delete" endpoint of the "chat_messages" service.
func NewDeleteUnauthenticatedResponseBody(res *goa.ServiceError) *DeleteUnauthenticatedResponseBody {
	body := &DeleteUnauthenticatedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListPayload builds a chat_messages service list endpoint payload.
func NewListPayload(limit int, channelID *uint64, after *uint64, before *uint64, idToken string) *chatmessages.ListPayload {
	v := &chatmessages.ListPayload{}
	v.Limit = limit
	v.ChannelID = channelID
	v.After = after
	v.Before = before
	v.IDToken = idToken

	return v
}

// NewShowPayload builds a chat_messages service show endpoint payload.
func NewShowPayload(id uint64, idToken string) *chatmessages.ShowPayload {
	v := &chatmessages.ShowPayload{}
	v.ID = id
	v.IDToken = idToken

	return v
}

// NewCreateChatMessageCreatePayload builds a chat_messages service create
// endpoint payload.
func NewCreateChatMessageCreatePayload(body *CreateRequestBody, idToken string) *chatmessages.ChatMessageCreatePayload {
	v := &chatmessages.ChatMessageCreatePayload{
		ChannelID: *body.ChannelID,
		Content:   *body.Content,
	}
	v.IDToken = idToken

	return v
}

// NewUpdateChatMessageUpdatePayload builds a chat_messages service update
// endpoint payload.
func NewUpdateChatMessageUpdatePayload(body *UpdateRequestBody, id uint64, idToken string) *chatmessages.ChatMessageUpdatePayload {
	v := &chatmessages.ChatMessageUpdatePayload{
		Content: *body.Content,
	}
	v.ID = id
	v.IDToken = idToken

	return v
}

// NewDeletePayload builds a chat_messages service delete endpoint payload.
func NewDeletePayload(id uint64, idToken string) *chatmessages.DeletePayload {
	v := &chatmessages.DeletePayload{}
	v.ID = id
	v.IDToken = idToken

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.ChannelID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("channel_id", "body"))
	}
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	return
}
