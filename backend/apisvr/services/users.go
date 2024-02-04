package chatapi

import (
	"applib/errors"
	"applib/firebase/auth"
	"applib/database/sql"
	"applib/time"
	"biz/models"
	log "apisvr/services/gen/log"
	users "apisvr/services/gen/users"
	"context"
	"fmt"
	"strings"

	"firebase.google.com/go/v4/errorutils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	baseService
	*UsersConvertor
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &userssrvc{baseService: newBaseService(logger), UsersConvertor: NewUsersConvertor()}
}

// List implements list.
func (s *userssrvc) List(ctx context.Context) (res *users.UserList, err error) {
	err = s.actionWithDB(ctx, "users.list", func(ctx context.Context, db *sql.DB) error {
		results, err := models.Users().All(ctx, db)
		if err != nil {
			return err
		}
		res = s.ModelsToList(results)
		return nil
	})
	return
}

// Create implements create.
func (s *userssrvc) Create(ctx context.Context, p *users.UserCreatePayload) (res *users.User, err error) {
	err = s.actionWithDB(ctx, "users.create", func(ctx context.Context, db *sql.DB) error {
		if p.Name == "" {
			return users.MakeInvalidPayload(fmt.Errorf("name is required"))
		} else {
			runes := []rune(p.Name)
			if len(runes) > 255 {
				return users.MakeInvalidPayload(fmt.Errorf("name is too long"))
			}
		}
		if p.Email == "" {
			return users.MakeInvalidPayload(fmt.Errorf("email is required"))
		} else {
			runes := []rune(p.Email)
			if len(runes) > 255 {
				return users.MakeInvalidPayload(fmt.Errorf("email is too long"))
			}
		}

		fbauth, err := s.firebaseAuthClient(ctx)
		if err != nil {
			return err
		}

		return sql.BeginTx(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
			if m, err := models.Users(qm.Where("email = ?", p.Email)).One(ctx, db); err != nil {
				if err == sql.ErrNoRows {
					// OK
				} else {
					return err
				}
			} else {
				_, err := fbauth.GetUserByEmail(ctx, p.Email)
				s.logger.Debug().Msgf("fbauth.GetUserByEmail: %+v", err)
				if err != nil {
					if !strings.Contains(err.Error(), "no user exists with the email") {
						return err
					}
					fbInput := &auth.UserToCreate{}
					fbInput.Email(p.Email)
					fbInput.EmailVerified(false)
					fbInput.DisplayName(p.Name)
					fbOutput, err := fbauth.CreateUser(ctx, fbInput)
					if err != nil {
						return errors.Wrapf(err, "something wrong. user was not find by email and cannot create user")
					}
					m.Name = fbOutput.DisplayName
					m.FbauthUID = fbOutput.UID
					if _, err := m.Update(ctx, db, boil.Infer()); err != nil {
						return err
					}
				}
				res = s.ModelToResult(m)
				return nil
			}

			m := &models.User{
				Name:  p.Name,
				Email: p.Email,
			}

			fbUser, err := fbauth.GetUserByEmail(ctx, p.Email)
			if err != nil {
				err := errors.Cause(err)
				if !errorutils.IsNotFound(err) {
					return err
				}
				fbInput := &auth.UserToCreate{}
				fbInput.Email(p.Email)
				fbInput.EmailVerified(false)
				fbInput.DisplayName(p.Name)
				fbOutput, err := fbauth.CreateUser(ctx, fbInput)
				if err != nil {
					return errors.Wrapf(err, "something wrong. user was not find by email and cannot create user")
				}
				m.FbauthUID = fbOutput.UID

			} else {
				m.FbauthUID = fbUser.UID
			}

			if err := m.Insert(ctx, db, boil.Infer()); err != nil {
				return err
			}

			res = s.ModelToResult(m)
			return nil
		})
	})
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
