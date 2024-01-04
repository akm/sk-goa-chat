package chatapi

import (
	"apisvr/lib/time"
	"apisvr/models"
	log "apisvr/services/gen/log"
	"apisvr/services/gen/users"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"apisvr/testlib/testsqlboiler"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestUsers(t *testing.T) {
	logger := testlog.New(t)
	conn := testsql.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := NewUsers(&log.Logger{Logger: logger})
	conv := NewUsersConvertor()

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			res, err := srvc.List(ctx)
			assert.NoError(t, err)
			assert.Zero(t, res.Total)
			assert.Zero(t, res.Offset)
			assert.Len(t, res.Items, 0)
		})
	})

	userFoo := &models.User{Email: "foo@example.com", Name: "Foo", FbauthUID: "foo-uid"}
	userBar := &models.User{Email: "bar@example.com", Name: "Bar", FbauthUID: "bar-uid"}
	testsqlboiler.Insert(t, ctx, conn, boil.Infer(), userFoo, userBar)

	t.Run("list", func(t *testing.T) {
		res, err := srvc.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), res.Total)
		assert.Zero(t, res.Offset)
		assert.Len(t, res.Items, 2)
		assert.Equal(t, conv.ModelsToList([]*models.User{userFoo, userBar}), res)
	})

	checkFooOnDB := func(t *testing.T, id uint64) *models.User {
		user1, err := models.FindUser(ctx, conn, id)
		assert.NoError(t, err)
		assert.Equal(t, "foo@example.com", user1.Email)
		assert.Equal(t, "Foo", user1.Name)
		assert.NotEmpty(t, user1.FbauthUID)
		return user1
	}

	var fooID uint64
	t.Run("create foo first time", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "foo@example.com", Name: "Foo"})
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
		fooID = res.ID
		user := checkFooOnDB(t, fooID)
		assert.Equal(t, conv.ModelToResult(user), res)
	})

	t.Run("create foo again", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "foo@example.com", Name: "Foo2"})
		assert.NoError(t, err)
		assert.Equal(t, fooID, res.ID)
		assert.Equal(t, "Foo", res.Name) // Foo2 にはならない
		user := checkFooOnDB(t, fooID)
		assert.Equal(t, conv.ModelToResult(user), res)
	})
}
