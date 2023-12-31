// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels HTTP client types
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	channels "apisvr/services/gen/channels"
	channelsviews "apisvr/services/gen/channels/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "channels" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateRequestBody is the type of the "channels" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// ListResponseBody is the type of the "channels" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Items
	Items ChannelListItemCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Total number of items
	Total *uint64 `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// Offset
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty" xml:"offset,omitempty"`
}

// ShowResponseBody is the type of the "channels" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CreateResponseBody is the type of the "channels" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateResponseBody is the type of the "channels" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// DeleteResponseBody is the type of the "channels" service "delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "channels" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreateInvalidPayloadResponseBody is the type of the "channels" service
// "create" endpoint HTTP response body for the "invalid_payload" error.
type CreateInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateNotFoundResponseBody is the type of the "channels" service "update"
// endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateInvalidPayloadResponseBody is the type of the "channels" service
// "update" endpoint HTTP response body for the "invalid_payload" error.
type UpdateInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteNotFoundResponseBody is the type of the "channels" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ChannelListItemCollectionResponseBody is used to define fields on response
// body types.
type ChannelListItemCollectionResponseBody []*ChannelListItemResponseBody

// ChannelListItemResponseBody is used to define fields on response body types.
type ChannelListItemResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "channels" service.
func NewCreateRequestBody(p *channels.ChannelCreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name: p.Name,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "channels" service.
func NewUpdateRequestBody(p *channels.ChannelUpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name: p.Name,
	}
	return body
}

// NewListChannelListOK builds a "channels" service "list" endpoint result from
// a HTTP "OK" response.
func NewListChannelListOK(body *ListResponseBody) *channelsviews.ChannelListView {
	v := &channelsviews.ChannelListView{
		Total:  body.Total,
		Offset: body.Offset,
	}
	v.Items = make([]*channelsviews.ChannelListItemView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalChannelListItemResponseBodyToChannelsviewsChannelListItemView(val)
	}

	return v
}

// NewShowChannelOK builds a "channels" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowChannelOK(body *ShowResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewShowNotFound builds a channels service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateChannelCreated builds a "channels" service "create" endpoint result
// from a HTTP "Created" response.
func NewCreateChannelCreated(body *CreateResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewCreateInvalidPayload builds a channels service create endpoint
// invalid_payload error.
func NewCreateInvalidPayload(body *CreateInvalidPayloadResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateChannelOK builds a "channels" service "update" endpoint result from
// a HTTP "OK" response.
func NewUpdateChannelOK(body *UpdateResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewUpdateNotFound builds a channels service update endpoint not_found error.
func NewUpdateNotFound(body *UpdateNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateInvalidPayload builds a channels service update endpoint
// invalid_payload error.
func NewUpdateInvalidPayload(body *UpdateInvalidPayloadResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteChannelOK builds a "channels" service "delete" endpoint result from
// a HTTP "OK" response.
func NewDeleteChannelOK(body *DeleteResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewDeleteNotFound builds a channels service delete endpoint not_found error.
func NewDeleteNotFound(body *DeleteNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreateInvalidPayloadResponseBody runs the validations defined on
// create_invalid_payload_response_body
func ValidateCreateInvalidPayloadResponseBody(body *CreateInvalidPayloadResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateNotFoundResponseBody runs the validations defined on
// update_not_found_response_body
func ValidateUpdateNotFoundResponseBody(body *UpdateNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateInvalidPayloadResponseBody runs the validations defined on
// update_invalid_payload_response_body
func ValidateUpdateInvalidPayloadResponseBody(body *UpdateInvalidPayloadResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteNotFoundResponseBody runs the validations defined on
// delete_not_found_response_body
func ValidateDeleteNotFoundResponseBody(body *DeleteNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateChannelListItemCollectionResponseBody runs the validations defined
// on Channel-List-ItemCollectionResponseBody
func ValidateChannelListItemCollectionResponseBody(body ChannelListItemCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateChannelListItemResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChannelListItemResponseBody runs the validations defined on
// Channel-List-ItemResponseBody
func ValidateChannelListItemResponseBody(body *ChannelListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created_at", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updated_at", *body.UpdatedAt, goa.FormatDateTime))
	}
	return
}
