// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications WebSocket client streaming
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	notifications "apisvr/services/gen/notifications"
	"io"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "notifications" service.
type ConnConfigurer struct {
	SubscribeFn goahttp.ConnConfigureFunc
}

// SubscribeClientStream implements the notifications.SubscribeClientStream
// interface.
type SubscribeClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "notifications" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		SubscribeFn: fn,
	}
}

// Recv reads instances of "notifications.NotificationEvent" from the
// "subscribe" endpoint websocket connection.
func (s *SubscribeClientStream) Recv() (*notifications.NotificationEvent, error) {
	var (
		rv   *notifications.NotificationEvent
		body SubscribeResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	err = ValidateSubscribeResponseBody(&body)
	if err != nil {
		return rv, err
	}
	res := NewSubscribeNotificationEventOK(&body)
	return res, nil
}