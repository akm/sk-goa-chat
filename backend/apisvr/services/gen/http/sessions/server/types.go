// Code generated by goa v3.14.1, DO NOT EDIT.
//
// sessions HTTP server types
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	sessions "apisvr/services/gen/sessions"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "sessions" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// ID Token
	IDToken *string `form:"id_token,omitempty" json:"id_token,omitempty" xml:"id_token,omitempty"`
}

// CreateInvalidPayloadResponseBody is the type of the "sessions" service
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

// DeleteInvalidPayloadResponseBody is the type of the "sessions" service
// "delete" endpoint HTTP response body for the "invalid_payload" error.
type DeleteInvalidPayloadResponseBody struct {
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

// NewCreateInvalidPayloadResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "sessions" service.
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

// NewDeleteInvalidPayloadResponseBody builds the HTTP response body from the
// result of the "delete" endpoint of the "sessions" service.
func NewDeleteInvalidPayloadResponseBody(res *goa.ServiceError) *DeleteInvalidPayloadResponseBody {
	body := &DeleteInvalidPayloadResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreatePayload builds a sessions service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody) *sessions.CreatePayload {
	v := &sessions.CreatePayload{
		IDToken: *body.IDToken,
	}

	return v
}

// NewDeletePayload builds a sessions service delete endpoint payload.
func NewDeletePayload(sessionID string) *sessions.DeletePayload {
	v := &sessions.DeletePayload{}
	v.SessionID = sessionID

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.IDToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id_token", "body"))
	}
	return
}