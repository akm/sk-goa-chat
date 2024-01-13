package chatapi

import (
	"apisvr/lib/time"
	log "apisvr/services/gen/log"
	sessions "apisvr/services/gen/sessions"
	"context"
)

// sessions service example implementation.
// The example methods log the requests and return zero values.
type sessionssrvc struct {
	baseService
}

// NewSessions returns the sessions service implementation.
func NewSessions(logger *log.Logger) sessions.Service {
	return &sessionssrvc{baseService: newBaseService(logger)}
}

// Create implements create.
func (s *sessionssrvc) Create(ctx context.Context, p *sessions.CreatePayload) (res *sessions.CreateResult, err error) {
	err = s.action(ctx, "sessions.create", func(ctx context.Context) error {
		fbauth, err := s.firebaseAuthClient(ctx)
		if err != nil {
			return err
		}
		sessionID, err := fbauth.SessionCookie(ctx, p.IDToken, 1*time.Hour)
		if err != nil {
			return err
		}
		res = &sessions.CreateResult{SessionID: sessionID}
		return nil
	})
	return
}

// Delete implements delete.
func (s *sessionssrvc) Delete(ctx context.Context, p *sessions.DeletePayload) (res *sessions.DeleteResult, err error) {
	err = s.action(ctx, "sessions.delete", func(ctx context.Context) error {
		fbauth, err := s.firebaseAuthClient(ctx)
		if err != nil {
			return err
		}
		token, err := fbauth.VerifySessionCookie(ctx, p.SessionID)
		if err != nil {
			return err
		}
		if err = fbauth.RevokeRefreshTokens(ctx, token.Subject); err != nil {
			return err
		}
		res = &sessions.DeleteResult{SessionID: ""} // sessionID を空文字列に設定
		return nil
	})
	return
}
