package chatapi

import (
	log "apisvr/services/gen/log"
	sessions "apisvr/services/gen/sessions"
	"context"
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
	res = &sessions.CreateResult{}
	s.logger.Info().Msg("sessions.create")
	return
}

// Delete implements delete.
func (s *sessionssrvc) Delete(ctx context.Context, p *sessions.DeletePayload) (err error) {
	s.logger.Info().Msg("sessions.delete")
	return
}
