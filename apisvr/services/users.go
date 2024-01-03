package chatapi

import (
	log "apisvr/services/gen/log"
	users "apisvr/services/gen/users"
	"context"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &userssrvc{logger}
}

// List implements list.
func (s *userssrvc) List(ctx context.Context) (res *users.UserList, err error) {
	res = &users.UserList{}
	s.logger.Info().Msg("users.list")
	return
}

// Create implements create.
func (s *userssrvc) Create(ctx context.Context, p *users.UserCreatePayload) (res *users.User, err error) {
	res = &users.User{}
	s.logger.Info().Msg("users.create")
	return
}
