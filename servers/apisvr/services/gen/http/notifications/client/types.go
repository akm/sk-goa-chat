// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications HTTP client types
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	notifications "apisvr/services/gen/notifications"

	goa "goa.design/goa/v3/pkg"
)

// SubscribeResponseBody is the type of the "notifications" service "subscribe"
// endpoint HTTP response body.
type SubscribeResponseBody struct {
	// IDs of channels which got new messages
	ChannelIds []uint64 `form:"channel_ids,omitempty" json:"channel_ids,omitempty" xml:"channel_ids,omitempty"`
}

// SubscribeUnauthenticatedResponseBody is the type of the "notifications"
// service "subscribe" endpoint HTTP response body for the "unauthenticated"
// error.
type SubscribeUnauthenticatedResponseBody struct {
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

// NewSubscribeNotificationEventOK builds a "notifications" service "subscribe"
// endpoint result from a HTTP "OK" response.
func NewSubscribeNotificationEventOK(body *SubscribeResponseBody) *notifications.NotificationEvent {
	v := &notifications.NotificationEvent{}
	v.ChannelIds = make([]uint64, len(body.ChannelIds))
	for i, val := range body.ChannelIds {
		v.ChannelIds[i] = val
	}

	return v
}

// NewSubscribeUnauthenticated builds a notifications service subscribe
// endpoint unauthenticated error.
func NewSubscribeUnauthenticated(body *SubscribeUnauthenticatedResponseBody) *goa.ServiceError {
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

// ValidateSubscribeResponseBody runs the validations defined on
// SubscribeResponseBody
func ValidateSubscribeResponseBody(body *SubscribeResponseBody) (err error) {
	if body.ChannelIds == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("channel_ids", "body"))
	}
	return
}

// ValidateSubscribeUnauthenticatedResponseBody runs the validations defined on
// subscribe_unauthenticated_response_body
func ValidateSubscribeUnauthenticatedResponseBody(body *SubscribeUnauthenticatedResponseBody) (err error) {
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
