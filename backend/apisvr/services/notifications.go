package chatapi

import (
	log "apisvr/services/gen/log"
	notifications "apisvr/services/gen/notifications"
	"applib/database/sql"
	"applib/errors"
	"applib/time"
	"context"
)

// notifications service example implementation.
// The example methods log the requests and return zero values.
type notificationssrvc struct {
	BaseAuthService
}

// NewNotifications returns the notifications service implementation.
func NewNotifications(logger *log.Logger) notifications.Service {
	return &notificationssrvc{
		BaseAuthService: NewBaseAuthService(
			logger,
			notifications.MakeUnauthenticated,
		),
	}
}

// Subscribe to events sent such new chat messages.
func (s *notificationssrvc) Subscribe(ctx context.Context, p *notifications.SubscribePayload, stream notifications.SubscribeServerStream) (err error) {
	defer func() {
		if err := stream.Close(); err != nil {
			s.logger.Printf("failed to close stream: [%T] %+v", err, err)
		}
	}()

	err = s.actionWithDB(ctx, "notifications.subscribe", func(ctx context.Context, db *sql.DB) error {
		fbauth, err := s.firebaseAuthClient(ctx)
		if err != nil {
			return err
		}
		if _, err := s.authenticate(ctx, db, fbauth, p.IDToken); err != nil {
			return err
		}

		ch := make(chan chatMessageEvent)
		newMessageChannels = append(newMessageChannels, ch)

		interval := 500 * time.Millisecond
		ticker := time.NewTicker(interval)

		done := false
		for {
			if _, err := s.authenticate(ctx, db, fbauth, p.IDToken); err != nil {
				return err
			}

			select {
			case <-ctx.Done():
				done = true
			case <-ticker.C:
				event := <-ch

				channelIDs := []uint64{event.ChannelID}
				if len(channelIDs) > 0 {
					if err := stream.Send(&notifications.NotificationEvent{ChannelIds: channelIDs}); err != nil {
						return errors.Wrapf(err, "failed to send notification event")
					}
				}
			}
			if done {
				break
			}
		}
		s.logger.Printf("subscription ended")
		return nil
	})

	return
}
