// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications gRPC server
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	notificationspb "apisvr/services/gen/grpc/notifications/pb"
	notifications "apisvr/services/gen/notifications"
	"context"
	"errors"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the notificationspb.NotificationsServer interface.
type Server struct {
	SubscribeH goagrpc.StreamHandler
	notificationspb.UnimplementedNotificationsServer
}

// SubscribeServerStream implements the notifications.SubscribeServerStream
// interface.
type SubscribeServerStream struct {
	stream notificationspb.Notifications_SubscribeServer
}

// New instantiates the server struct with the notifications service endpoints.
func New(e *notifications.Endpoints, sh goagrpc.StreamHandler) *Server {
	return &Server{
		SubscribeH: NewSubscribeHandler(e.Subscribe, sh),
	}
}

// NewSubscribeHandler creates a gRPC handler which serves the "notifications"
// service "subscribe" endpoint.
func NewSubscribeHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, DecodeSubscribeRequest)
	}
	return h
}

// Subscribe implements the "Subscribe" method in
// notificationspb.NotificationsServer interface.
func (s *Server) Subscribe(message *notificationspb.SubscribeRequest, stream notificationspb.Notifications_SubscribeServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
	ctx = context.WithValue(ctx, goa.ServiceKey, "notifications")
	p, err := s.SubscribeH.Decode(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthenticated":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &notifications.SubscribeEndpointInput{
		Stream:  &SubscribeServerStream{stream: stream},
		Payload: p.(*notifications.SubscribePayload),
	}
	err = s.SubscribeH.Handle(ctx, ep)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthenticated":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// Send streams instances of "notificationspb.SubscribeResponse" to the
// "subscribe" endpoint gRPC stream.
func (s *SubscribeServerStream) Send(res *notifications.NotificationEvent) error {
	v := NewProtoNotificationEventSubscribeResponse(res)
	return s.stream.Send(v)
}

func (s *SubscribeServerStream) Close() error {
	// nothing to do here
	return nil
}
