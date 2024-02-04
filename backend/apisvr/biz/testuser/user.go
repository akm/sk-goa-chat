package testuser

import (
	"apisvr/applib/firebase/auth"
	"apisvr/applib/google/identitytoolkit/identitytoolkittest"
	"apisvr/applib/time"
	"apisvr/models"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User struct {
	Email    string
	Password string
	Name     string

	UID     string
	IDToken string

	Model *models.User

	SessionID string
}

func New(email, password, name string) *User {
	return &User{Email: email, Password: password, Name: name}
}

func (u *User) CreateOnFirebaseAuth(t *testing.T, ctx context.Context, fbauth auth.Client) {
	args := &auth.UserToCreate{}
	args.Email(u.Email)
	args.DisplayName(u.Name)
	args.Password(u.Password)
	res, err := fbauth.CreateUser(ctx, args)
	require.NoError(t, err)
	// t.Logf("result: %+v", res)
	require.NotEmpty(t, res.UID)
	u.UID = res.UID
}

func (u *User) CreateOnDB(t *testing.T, ctx context.Context, conn *sql.DB) {
	user := &models.User{FbauthUID: u.UID, Email: u.Email, Name: u.Name}
	err := user.Insert(ctx, conn, boil.Infer())
	require.NoError(t, err)
	u.Model = user
}

func (u *User) GetIdToken(t *testing.T, ctx context.Context, fbauth auth.Client) string {
	if u.UID == "" {
		u.CreateOnFirebaseAuth(t, ctx, fbauth)
	}
	u.IDToken = identitytoolkittest.GetIdToken(t, ctx, u.Email, u.Password)
	assert.NotEmpty(t, u.IDToken)
	return u.IDToken
}

func (u *User) GetSessionID(t *testing.T, ctx context.Context, fbauth auth.Client) string {
	u.GetIdToken(t, ctx, fbauth)
	cookie, err := fbauth.SessionCookie(ctx, u.IDToken, 1*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	u.SessionID = cookie
	return u.SessionID
}

func (u *User) Setup(t *testing.T, ctx context.Context, fbauth auth.Client, conn *sql.DB) *User {
	u.CreateOnFirebaseAuth(t, ctx, fbauth)
	u.CreateOnDB(t, ctx, conn)
	u.GetSessionID(t, ctx, fbauth)
	return u
}
