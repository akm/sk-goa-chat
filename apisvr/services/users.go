package chatapi

import (
	"apisvr/models"
	log "apisvr/services/gen/log"
	users "apisvr/services/gen/users"
	"context"
	"time"
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

type UsersConvertor struct{}

func NewUsersConvertor() *UsersConvertor {
	return &UsersConvertor{}
}

func (s *UsersConvertor) ModelsToList(models []*models.User) *users.UserList {
	items := s.ModelsToListItems(models)
	return &users.UserList{
		Items:  items,
		Total:  uint64(len(items)),
		Offset: 0,
	}
}

func (*UsersConvertor) ModelsToListItems(models []*models.User) users.UserListItemCollection {
	items := make(users.UserListItemCollection, len(models))
	for i, result := range models {
		items[i] = &users.UserListItem{
			ID:   result.ID,
			Name: result.Name,
		}
	}
	return items
}

func (*UsersConvertor) ModelToResult(model *models.User) *users.User {
	return &users.User{
		ID:        model.ID,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
		Name:      model.Name,
	}
}
