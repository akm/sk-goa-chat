package chatapi

import (
	log "apisvr/services/gen/log"
	notifications "apisvr/services/gen/notifications"
	"context"
)

// notifications service example implementation.
// The example methods log the requests and return zero values.
type notificationssrvc struct {
	logger *log.Logger
}

// NewNotifications returns the notifications service implementation.
func NewNotifications(logger *log.Logger) notifications.Service {
	return &notificationssrvc{logger}
}

// Subscribe to events sent such new chat messages.
func (s *notificationssrvc) Subscribe(ctx context.Context, p *notifications.SubscribePayload, stream notifications.SubscribeServerStream) (err error) {
	s.logger.Info().Msg("notifications.subscribe")
	return
}
