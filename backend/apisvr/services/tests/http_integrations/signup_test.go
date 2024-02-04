package httpintegrations

import (
	"applib/encoding/json/jsontest"
	"applib/firebase"
	"applib/firebase/auth"
	"applib/firebase/auth/authtest"
	"applib/google/identitytoolkit/identitytoolkittest"
	"applib/log/logtest"
	"applib/database/sql/sqltest"
	"applib/time"
	"biz/models"
	chatapi "apisvr/services"
	"apisvr/services/gen/http/channels/server"
	sessionsserver "apisvr/services/gen/http/sessions/server"
	usersserver "apisvr/services/gen/http/users/server"
	"apisvr/services/gen/log"
	"apisvr/services/gen/sessions"
	"apisvr/services/gen/users"
	"context"
	"net/http"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestSignup(t *testing.T) {
	logger := logtest.New(t)
	conn := sqltest.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	usersConv := chatapi.NewUsersConvertor()

	fbapp, err := firebase.NewApp(ctx, nil)
	require.NoError(t, err)
	fbauth, err := auth.NewClientRaw(ctx, fbapp)
	require.NoError(t, err)

	t.Run("delete all of users before test", func(t *testing.T) {
		authtest.DeleteUsers(t, ctx, fbauth)
	})

	checker := goahttpcheck.New()
	// checker.Mount(usersServer.NewListHandler, server.MountListHandler, channels.NewListEndpoint(srvc))
	usersSrvc := chatapi.NewUsers(&log.Logger{Logger: logger})
	checker.Mount(usersserver.NewCreateHandler, usersserver.MountCreateHandler, users.NewCreateEndpoint(usersSrvc))

	sessionSrvc := chatapi.NewSessions(&log.Logger{Logger: logger})
	checker.Mount(sessionsserver.NewCreateHandler, sessionsserver.MountCreateHandler, sessions.NewCreateEndpoint(sessionSrvc))
	checker.Mount(sessionsserver.NewDeleteHandler, sessionsserver.MountDeleteHandler, sessions.NewDeleteEndpoint(sessionSrvc))

	fooEmail := "foo@example.com"
	fooName := "Foo"
	fooPassword := "Passw0rd!"
	var fooUID string
	var fooIDToken string

	t.Run("create user on firebase", func(t *testing.T) {
		args := &auth.UserToCreate{}
		args.Email(fooEmail)
		args.DisplayName(fooName)
		args.Password(fooPassword)
		res, err := fbauth.CreateUser(ctx, args)
		require.NoError(t, err)
		// t.Logf("result: %+v", res)
		require.NotEmpty(t, res.UID)
		fooUID = res.UID

		// ブラウザ上のJSではfirebaseのライブラリによって CreateUser の結果から IDToken を取得できますが、
		// Goでは実装されていないので、代わりに identitytoolkit のAPIの呼び出してIDTokenを取得します。
		fooIDToken = identitytoolkittest.GetIdToken(t, ctx, fooEmail, fooPassword)
		assert.NotEmpty(t, fooIDToken)
	})

	t.Run("create user on apisvr", func(t *testing.T) {
		checker.Test(t, http.MethodPost, "/api/users").
			WithJSON(map[string]any{"email": fooEmail, "name": fooName}).
			Check().HasStatus(http.StatusCreated).Cb(
			func(r *http.Response) {
				defer r.Body.Close()
				res := jsontest.UnmarshalFrom[server.CreateResponseBody](t, r.Body)
				expected := jsontest.Unmarshal[server.CreateResponseBody](t,
					jsontest.MarshalAndSnakeizeJsonKeys(t,
						usersConv.ModelToResult(&models.User{ID: res.ID, Email: fooEmail, Name: fooName, CreatedAt: now, UpdatedAt: now}),
					),
				)
				assert.Equal(t, expected, res)
			})

		t.Run("assert user on db", func(t *testing.T) {
			user, err := models.Users(qm.Where("email = ?", fooEmail)).One(ctx, conn)
			require.NoError(t, err)
			assert.Equal(t, fooEmail, user.Email)
			assert.Equal(t, fooName, user.Name)
			assert.Equal(t, fooUID, user.FbauthUID)
		})
	})

	var fooSessionID string
	t.Run("create session", func(t *testing.T) {
		checker.Test(t, http.MethodPost, "/api/session").
			WithJSON(map[string]any{"id_token": fooIDToken}).
			Check().HasStatus(http.StatusCreated).Cb(
			func(r *http.Response) {
				defer r.Body.Close()

				cookieMap := map[string]string{}
				for _, c := range r.Cookies() {
					cookieMap[c.Name] = c.Value
				}
				assert.NotEmpty(t, cookieMap["session_id"])
				fooSessionID = cookieMap["session_id"]
				assert.NotEmpty(t, fooSessionID)
			})
	})

	t.Run("delete session", func(t *testing.T) {
		checker.Test(t, http.MethodDelete, "/api/session").
			WithCookie("session_id", fooSessionID).
			Check().HasStatus(http.StatusOK)
	})
}
