package chatapi

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	goa "goa.design/goa/v3/pkg"

	"apisvr/lib/errors"
	"apisvr/lib/firebase"
	"apisvr/lib/firebase/auth"
	"apisvr/lib/sql"
	"apisvr/models"
	_ "apisvr/models_ext"
	log "apisvr/services/gen/log"
)

func SetupContext(ctx context.Context) context.Context {
	return boil.SkipTimestamps(ctx)
}

type baseService struct {
	logger *log.Logger
}

func newBaseService(logger *log.Logger) baseService {
	return baseService{logger: logger}
}

func (s *baseService) sqlOpen() (*sql.DB, error) {
	return sql.Open(s.logger.Logger)
}

func (s *baseService) action(ctx context.Context, name string, cb func(context.Context) error) error {
	s.logger.Info().Msg(name)
	return cb(SetupContext(ctx))
}

func (s *baseService) actionWithDB(ctx context.Context, name string, cb func(context.Context, *sql.DB) error) error {
	s.logger.Info().Msg(name)
	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return err
	}
	defer db.Close()
	return cb(ctx, db)
}

func (s *baseService) firebaseAuthClient(ctx context.Context) (auth.Client, error) {
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

type baseAuthService struct {
	baseService
	ConvToAuthenticationError func(error) *goa.ServiceError
}

func newBaseAuthService(logger *log.Logger, convToAuthenticationError func(error) *goa.ServiceError) baseAuthService {
	return baseAuthService{
		baseService:               newBaseService(logger),
		ConvToAuthenticationError: convToAuthenticationError,
	}
}

func (s *baseAuthService) authenticate(ctx context.Context, db *sql.DB, sessionID string) (*models.User, error) {
	fbauth, err := s.firebaseAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	token, err := fbauth.VerifySessionCookie(ctx, sessionID)
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

func (s *baseAuthService) actionWithAuth(ctx context.Context, name string, sessionID string, cb func(context.Context, *sql.DB, *models.User) error) error {
	s.logger.Info().Msg(name)
	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	user, err := s.authenticate(ctx, db, sessionID)
	if err != nil {
		return err
	}

	return cb(ctx, db, user)
}
