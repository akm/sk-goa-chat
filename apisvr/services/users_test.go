package chatapi

import (
	"apisvr/lib/time"
	"apisvr/models"
	log "apisvr/services/gen/log"
	"apisvr/services/gen/users"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	logger := testlog.New(t)
	conn := testsql.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := NewUsers(&log.Logger{Logger: logger})

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			res, err := srvc.List(ctx)
			assert.NoError(t, err)
			assert.Zero(t, res.Total)
			assert.Zero(t, res.Offset)
			assert.Len(t, res.Items, 0)
		})
	})

	checkFooOnDB := func(id uint64) func(t *testing.T) {
		return func(t *testing.T) {
			user1, err := models.FindUser(ctx, conn, id)
			assert.NoError(t, err)
			assert.Equal(t, "foo@example.com", user1.Email)
			assert.Equal(t, "Foo", user1.Name)
			assert.NotEmpty(t, user1.FbauthUID)
		}
	}

	var fooID uint64
	t.Run("create foo first time", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "foo@example.com", Name: "Foo"})
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
		fooID = res.ID
	})
	t.Run("check data on DB", checkFooOnDB(fooID))

	t.Run("create foo again", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "foo@example.com", Name: "Foo2"})
		assert.NoError(t, err)
		assert.Equal(t, fooID, res.ID)
		assert.Equal(t, "Foo", res.Name) // Foo2 にはならない
	})
	t.Run("check data on DB again", checkFooOnDB(fooID))

	t.Run("create bar", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "bar@example.com", Name: "Bar"})
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
	})

	t.Run("list", func(t *testing.T) {
		res, err := srvc.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), res.Total)
		assert.Zero(t, res.Offset)
		assert.Len(t, res.Items, 2)
	})
}
