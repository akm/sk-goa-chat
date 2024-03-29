// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels HTTP server types
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	channels "apisvr/services/gen/channels"
	channelsviews "apisvr/services/gen/channels/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "channels" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateRequestBody is the type of the "channels" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ListResponseBody is the type of the "channels" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Items
	Items ChannelListItemResponseBodyCollection `form:"items" json:"items" xml:"items"`
	// Total number of items
	Total uint64 `form:"total" json:"total" xml:"total"`
	// Offset
	Offset uint64 `form:"offset" json:"offset" xml:"offset"`
}

// ShowResponseBody is the type of the "channels" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateResponseBody is the type of the "channels" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateResponseBody is the type of the "channels" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// DeleteResponseBody is the type of the "channels" service "delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// ListUnauthenticatedResponseBody is the type of the "channels" service "list"
// endpoint HTTP response body for the "unauthenticated" error.
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

// ShowNotFoundResponseBody is the type of the "channels" service "show"
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

// ShowUnauthenticatedResponseBody is the type of the "channels" service "show"
// endpoint HTTP response body for the "unauthenticated" error.
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

// CreateInvalidPayloadResponseBody is the type of the "channels" service
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

// CreateUnauthenticatedResponseBody is the type of the "channels" service
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

// UpdateNotFoundResponseBody is the type of the "channels" service "update"
// endpoint HTTP response body for the "not_found" error.
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

// UpdateInvalidPayloadResponseBody is the type of the "channels" service
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

// UpdateUnauthenticatedResponseBody is the type of the "channels" service
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

// DeleteNotFoundResponseBody is the type of the "channels" service "delete"
// endpoint HTTP response body for the "not_found" error.
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

// DeleteUnauthenticatedResponseBody is the type of the "channels" service
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

// ChannelListItemResponseBodyCollection is used to define fields on response
// body types.
type ChannelListItemResponseBodyCollection []*ChannelListItemResponseBody

// ChannelListItemResponseBody is used to define fields on response body types.
type ChannelListItemResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "channels" service.
func NewListResponseBody(res *channelsviews.ChannelListView) *ListResponseBody {
	body := &ListResponseBody{
		Total:  *res.Total,
		Offset: *res.Offset,
	}
	if res.Items != nil {
		body.Items = make([]*ChannelListItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalChannelsviewsChannelListItemViewToChannelListItemResponseBody(val)
		}
	} else {
		body.Items = []*ChannelListItemResponseBody{}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "channels" service.
func NewShowResponseBody(res *channelsviews.ChannelView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		Name:      *res.Name,
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "channels" service.
func NewCreateResponseBody(res *channelsviews.ChannelView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		Name:      *res.Name,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "channels" service.
func NewUpdateResponseBody(res *channelsviews.ChannelView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		Name:      *res.Name,
	}
	return body
}

// NewDeleteResponseBody builds the HTTP response body from the result of the
// "delete" endpoint of the "channels" service.
func NewDeleteResponseBody(res *channelsviews.ChannelView) *DeleteResponseBody {
	body := &DeleteResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		Name:      *res.Name,
	}
	return body
}

// NewListUnauthenticatedResponseBody builds the HTTP response body from the
// result of the "list" endpoint of the "channels" service.
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
// the "show" endpoint of the "channels" service.
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
// result of the "show" endpoint of the "channels" service.
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
// result of the "create" endpoint of the "channels" service.
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
// result of the "create" endpoint of the "channels" service.
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
// of the "update" endpoint of the "channels" service.
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
// result of the "update" endpoint of the "channels" service.
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
// result of the "update" endpoint of the "channels" service.
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
// of the "delete" endpoint of the "channels" service.
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
// result of the "delete" endpoint of the "channels" service.
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

// NewListPayload builds a channels service list endpoint payload.
func NewListPayload(idToken string) *channels.ListPayload {
	v := &channels.ListPayload{}
	v.IDToken = idToken

	return v
}

// NewShowPayload builds a channels service show endpoint payload.
func NewShowPayload(id uint64, idToken string) *channels.ShowPayload {
	v := &channels.ShowPayload{}
	v.ID = id
	v.IDToken = idToken

	return v
}

// NewCreateChannelCreatePayload builds a channels service create endpoint
// payload.
func NewCreateChannelCreatePayload(body *CreateRequestBody, idToken string) *channels.ChannelCreatePayload {
	v := &channels.ChannelCreatePayload{
		Name: *body.Name,
	}
	v.IDToken = idToken

	return v
}

// NewUpdateChannelUpdatePayload builds a channels service update endpoint
// payload.
func NewUpdateChannelUpdatePayload(body *UpdateRequestBody, id uint64, idToken string) *channels.ChannelUpdatePayload {
	v := &channels.ChannelUpdatePayload{
		Name: *body.Name,
	}
	v.ID = id
	v.IDToken = idToken

	return v
}

// NewDeletePayload builds a channels service delete endpoint payload.
func NewDeletePayload(id uint64, idToken string) *channels.DeletePayload {
	v := &channels.DeletePayload{}
	v.ID = id
	v.IDToken = idToken

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
