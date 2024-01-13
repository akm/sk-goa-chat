package chatapi

import (
	"apisvr/lib/errors"
	"apisvr/lib/sql"
	"apisvr/models"
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
	err = s.actionWithAuth(ctx, "notifications.subscribe", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {

		query := func(ctx context.Context) (map[uint64]uint64, error) {
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
					return err
				}

				channelIDs := []uint64{}
				for channelID, maxID := range currMap {
					if lastMap[channelID] < maxID {
						channelIDs = append(channelIDs, channelID)
					}
				}
				if len(channelIDs) > 0 {
					stream.Send(&notifications.NotificationEvent{ChannelIds: channelIDs})
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
