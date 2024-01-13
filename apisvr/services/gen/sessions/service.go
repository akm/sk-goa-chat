// Code generated by goa v3.14.1, DO NOT EDIT.
//
// sessions service
//
// Command:
// $ goa gen apisvr/design -o ./services

package sessions

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the sessions service interface.
type Service interface {
	// Create implements create.
	Create(context.Context, *CreatePayload) (res *CreateResult, err error)
	// Delete implements delete.
	Delete(context.Context, *DeletePayload) (res *DeleteResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "sessions"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"create", "delete"}

// CreatePayload is the payload type of the sessions service create method.
type CreatePayload struct {
	// ID Token
	IDToken string
}

// CreateResult is the result type of the sessions service create method.
type CreateResult struct {
	// Session ID
	SessionID string
}

// DeletePayload is the payload type of the sessions service delete method.
type DeletePayload struct {
	// Session ID
	SessionID string
}

// DeleteResult is the result type of the sessions service delete method.
type DeleteResult struct {
	// Session ID
	SessionID string
}

// MakeInvalidPayload builds a goa.ServiceError from an error.
func MakeInvalidPayload(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_payload", false, false, false)
}
