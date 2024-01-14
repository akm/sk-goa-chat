package chatapi

import (
	"apisvr/lib/errors"
	"apisvr/lib/sql"
	log "apisvr/services/gen/log"
	notifications "apisvr/services/gen/notifications"
	"context"
	"time"
)

// notifications service example implementation.
// The example methods log the requests and return zero values.
type notificationssrvc struct {
	baseAuthService
}

// NewNotifications returns the notifications service implementation.
func NewNotifications(logger *log.Logger) notifications.Service {
	return &notificationssrvc{
		baseAuthService: newBaseAuthService(
			logger,
			notifications.MakeUnauthenticated,
		),
	}
}

// Subscribe to events sent such new chat messages.
func (s *notificationssrvc) Subscribe(ctx context.Context, p *notifications.SubscribePayload, stream notifications.SubscribeServerStream) (err error) {
	err = s.actionWithDB(ctx, "notifications.subscribe", func(ctx context.Context, db *sql.DB) error {
		fbauth, err := s.firebaseAuthClient(ctx)
		if err != nil {
			return err
		}

		query := func(ctx context.Context) (map[uint64]uint64, error) {
			if _, err := s.authenticate(ctx, db, fbauth, p.SessionID); err != nil {
				return nil, err
			}

			rows, err := db.QueryContext(ctx, "SELECT channel_id, MAX(id) AS max_id FROM chat_messages GROUP BY channel_id")
			if err != nil {
				return nil, errors.Wrapf(err, "failed to query chat messages")
			}
			defer rows.Close()

			r := map[uint64]uint64{}
			for rows.Next() {
				var channelID, maxID uint64
				if err := rows.Scan(&channelID, &maxID); err != nil {
					return nil, errors.Wrapf(err, "failed to scan row")
				}
				r[channelID] = maxID
			}
			return r, nil
		}

		interval := 500 * time.Millisecond
		ticker := time.NewTicker(interval)

		lastMap, err := query(ctx)
		if err != nil {
			return err
		}

		done := false
		for {
			select {
			case <-ctx.Done():
				done = true
			case <-ticker.C:
				currMap, err := query(ctx)
				if err != nil {
					if err := stream.Close(); err != nil {
						s.logger.Printf("failed to close stream: [%T] %+v", err, err)
					} else {
						return err
					}
				}

				channelIDs := []uint64{}
				for channelID, maxID := range currMap {
					if lastMap[channelID] < maxID {
						channelIDs = append(channelIDs, channelID)
					}
				}
				if len(channelIDs) > 0 {
					if err := stream.Send(&notifications.NotificationEvent{ChannelIds: channelIDs}); err != nil {
						sendError := errors.Wrapf(err, "failed to send notification event")
						if err := stream.Close(); err != nil {
							s.logger.Printf("failed to close stream: [%T] %+v", err, err)
							return errors.Join(sendError, errors.Wrapf(err, "failed to close stream"))
						} else {
							return sendError
						}
					}
					lastMap = currMap
				}
			}
			if done {
				break
			}
		}
		s.logger.Printf("subscription ended")
		return stream.Close()
	})

	return
}
