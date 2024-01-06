package chatapi

import (
	"apisvr/lib/errors"
	"apisvr/lib/firebase"
	"apisvr/lib/firebase/auth"
	log "apisvr/services/gen/log"
	sessions "apisvr/services/gen/sessions"
	"context"
	"time"
)

// sessions service example implementation.
// The example methods log the requests and return zero values.
type sessionssrvc struct {
	logger *log.Logger
}

// NewSessions returns the sessions service implementation.
func NewSessions(logger *log.Logger) sessions.Service {
	return &sessionssrvc{logger}
}

// Create implements create.
func (s *sessionssrvc) Create(ctx context.Context, p *sessions.CreatePayload) (res *sessions.CreateResult, err error) {
	s.logger.Info().Msg("sessions.create")

	fbapp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "firebase.NewApp")
	}
	fbauth, err := auth.NewClientWithLogger(ctx, fbapp, s.logger.Logger)
	if err != nil {
		return nil, errors.Wrapf(err, "auth.NewClientWithLogger")
	}

	sessionID, err := fbauth.SessionCookie(ctx, p.IDToken, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	res = &sessions.CreateResult{SessionID: sessionID}
	return
}

// Delete implements delete.
func (s *sessionssrvc) Delete(ctx context.Context, p *sessions.DeletePayload) (err error) {
	s.logger.Info().Msg("sessions.delete")

	fbapp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return errors.Wrapf(err, "firebase.NewApp")
	}
	fbauth, err := auth.NewClientWithLogger(ctx, fbapp, s.logger.Logger)
	if err != nil {
		return errors.Wrapf(err, "auth.NewClientWithLogger")
	}

	token, err := fbauth.VerifySessionCookie(ctx, p.SessionID)
	if err != nil {
		return err
	}

	if err = fbauth.RevokeRefreshTokens(ctx, token.Subject); err != nil {
		return err
	}

	return
}
