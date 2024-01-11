// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications service
//
// Command:
// $ goa gen apisvr/design -o ./services

package notifications

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the notifications service interface.
type Service interface {
	// Subscribe to events sent such new chat messages.
	Subscribe(context.Context, *SubscribePayload, SubscribeServerStream) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "notifications"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"subscribe"}

// SubscribeServerStream is the interface a "subscribe" endpoint server stream
// must satisfy.
type SubscribeServerStream interface {
	// Send streams instances of "NotificationEvent".
	Send(*NotificationEvent) error
	// Close closes the stream.
	Close() error
}

// SubscribeClientStream is the interface a "subscribe" endpoint client stream
// must satisfy.
type SubscribeClientStream interface {
	// Recv reads instances of "NotificationEvent" from the stream.
	Recv() (*NotificationEvent, error)
}

// NotificationEvent is the result type of the notifications service subscribe
// method.
type NotificationEvent struct {
	// IDs of channels which got new messages
	ChannelIds []uint64
}

// SubscribePayload is the payload type of the notifications service subscribe
// method.
type SubscribePayload struct {
	// Session ID
	SessionID string
}

// MakeUnauthenticated builds a goa.ServiceError from an error.
func MakeUnauthenticated(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "unauthenticated", false, false, false)
}
