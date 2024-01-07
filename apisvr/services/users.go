package chatapi

import (
	"apisvr/lib/errors"
	"apisvr/lib/firebase"
	"apisvr/lib/firebase/auth"
	"apisvr/lib/sql"
	"apisvr/models"
	log "apisvr/services/gen/log"
	users "apisvr/services/gen/users"
	"context"
	"fmt"
	"strings"
	"time"

	"firebase.google.com/go/v4/errorutils"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
	*UsersConvertor
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &userssrvc{logger: logger, UsersConvertor: NewUsersConvertor()}
}

func (s *userssrvc) sqlOpen() (*sql.DB, error) {
	return sql.Open(s.logger.Logger)
}

// List implements list.
func (s *userssrvc) List(ctx context.Context) (res *users.UserList, err error) {
	s.logger.Info().Msg("users.list")
	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := models.Users().All(ctx, db)
	if err != nil {
		return nil, err
	}

	res = s.ModelsToList(results)
	return
}

// Create implements create.
func (s *userssrvc) Create(ctx context.Context, p *users.UserCreatePayload) (res *users.User, err error) {
	s.logger.Info().Msg("users.create")

	if p.Name == "" {
		return nil, users.MakeInvalidPayload(fmt.Errorf("name is required"))
	} else {
		runes := []rune(p.Name)
		if len(runes) > 255 {
			return nil, users.MakeInvalidPayload(fmt.Errorf("name is too long"))
		}
	}
	if p.Email == "" {
		return nil, users.MakeInvalidPayload(fmt.Errorf("email is required"))
	} else {
		runes := []rune(p.Email)
		if len(runes) > 255 {
			return nil, users.MakeInvalidPayload(fmt.Errorf("email is too long"))
		}
	}

	fbapp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "firebase.NewApp")
	}
	fbauth, err := auth.NewClientWithLogger(ctx, fbapp, s.logger.Logger)
	if err != nil {
		return nil, errors.Wrapf(err, "auth.NewClientWithLogger")
	}

	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = sql.BeginTx(ctx, db, func(ctx context.Context, tx *sql.Tx) error {
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
