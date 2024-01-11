// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications gRPC client
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	notificationspb "apisvr/services/gen/grpc/notifications/pb"
	notifications "apisvr/services/gen/notifications"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli notificationspb.NotificationsClient
	opts    []grpc.CallOption
}

// SubscribeClientStream implements the notifications.SubscribeClientStream
// interface.
type SubscribeClientStream struct {
	stream notificationspb.Notifications_SubscribeClient
}

// NewClient instantiates gRPC client for all the notifications service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: notificationspb.NewNotificationsClient(cc),
		opts:    opts,
	}
}

// Subscribe calls the "Subscribe" function in
// notificationspb.NotificationsClient interface.
func (c *Client) Subscribe() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSubscribeFunc(c.grpccli, c.opts...),
			EncodeSubscribeRequest,
			DecodeSubscribeResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Recv reads instances of "notificationspb.SubscribeResponse" from the
// "subscribe" endpoint gRPC stream.
func (s *SubscribeClientStream) Recv() (*notifications.NotificationEvent, error) {
	var res *notifications.NotificationEvent
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	if err = ValidateSubscribeResponse(v); err != nil {
		return res, err
	}
	return NewSubscribeResponseNotificationEvent(v), nil
}
