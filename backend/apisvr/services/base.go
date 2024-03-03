package chatapi

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"

	log "apisvr/services/gen/log"
	"applib/database/sql"
	"applib/errors"
	"applib/firebase"
	"applib/firebase/auth"
	"biz/models"
	_ "biz/models_ext"
)

func SetupContext(ctx context.Context) context.Context {
	return boil.SkipTimestamps(ctx)
}

type BaseService struct {
	logger *log.Logger
}

func NewBaseService(logger *log.Logger) BaseService {
	return BaseService{logger: logger}
}

func (s *BaseService) action(ctx context.Context, name string, cb func(context.Context) error) error {
	s.logger.Info().Msg(name)
	return cb(SetupContext(ctx))
}

func (s *BaseService) sqlOpen() (*sql.DB, error) {
	return sql.Open(s.logger.Logger)
}

func (s *BaseService) actionWithDB(ctx context.Context, name string, cb func(context.Context, *sql.DB) error) error {
	return s.action(ctx, name, func(ctx context.Context) error {
		db, err := sql.ConnectionFromContext(ctx)
		if err != nil {
			return err
		}
		return cb(ctx, db)
	})
}

func (s *BaseService) firebaseAuthClient(ctx context.Context) (auth.Client, error) {
	fbapp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "firebase.NewApp")
	}
	fbauth, err := auth.NewClientWithLogger(ctx, fbapp, s.logger.Logger)
	if err != nil {
		return nil, errors.Wrapf(err, "auth.NewClientWithLogger")
	}
	return fbauth, nil
}

type BaseAuthService struct {
	BaseService
	ConvToAuthenticationError func(error) *goa.ServiceError
}

func NewBaseAuthService(logger *log.Logger, convToAuthenticationError func(error) *goa.ServiceError) BaseAuthService {
	return BaseAuthService{
		BaseService:               NewBaseService(logger),
		ConvToAuthenticationError: convToAuthenticationError,
	}
}

func (s *BaseAuthService) authenticate(ctx context.Context, db *sql.DB, fbauth auth.Client, idToken string) (*models.User, error) {
	token, err := fbauth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, s.ConvToAuthenticationError(err)
	}

	user, err := models.Users(qm.Where("fbauth_uid = ?", token.UID)).One(ctx, db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, s.ConvToAuthenticationError(fmt.Errorf("user not found"))
		} else {
			return nil, errors.Wrapf(err, "failed to query user")
		}
	}

	return user, nil
}

func (s *BaseAuthService) APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (newCtx context.Context, err error) {
	newCtx = ctx
	err = func() error {
		// TODO Get db from context for message processing
		db, err := s.sqlOpen()
		if err != nil {
			return err
		}
		defer db.Close()

		fbauth, err := s.firebaseAuthClient(ctx)
		if err != nil {
			return err
		}

		u, err := s.authenticate(ctx, db, fbauth, key)
		if err != nil {
			return err
		}
		newCtx = NewContextWithUser(ctx, u)
		return nil
	}()
	return
}

// func (s *baseAuthService) actionWithAuth(ctx context.Context, name string, idToken string, cb func(context.Context, *sql.DB, *models.User) error) error {
// 	return s.actionWithDB(ctx, name, func(ctx context.Context, db *sql.DB) error {
// 		fbauth, err := s.firebaseAuthClient(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		user, err := s.authenticate(ctx, db, fbauth, idToken)
// 		if err != nil {
// 			return err
// 		}
// 		return cb(ctx, db, user)
// 	})
// }

func (s *BaseAuthService) actionWithUser(ctx context.Context, name string, idToken string, cb func(context.Context, *sql.DB, *models.User) error) error {
	return s.actionWithDB(ctx, name, func(ctx context.Context, db *sql.DB) error {
		u, err := UserFromContext(ctx)
		if err != nil {
			return err
		}
		return cb(ctx, db, u)
	})
}
