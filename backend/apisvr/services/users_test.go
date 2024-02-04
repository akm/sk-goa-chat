package chatapi

import (
	"apisvr/models"
	log "apisvr/services/gen/log"
	"apisvr/services/gen/users"
	"applib/database/sql/sqltest"
	"applib/log/logtest"
	"applib/sqlboiler/sqlboilertest"
	"applib/time"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestUsers(t *testing.T) {
	logger := logtest.New(t)
	conn := sqltest.Setup(t, logger)
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
	sqlboilertest.Insert(t, ctx, conn, boil.Infer(), userFoo, userBar)

	t.Run("list", func(t *testing.T) {
		res, err := srvc.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), res.Total)
		assert.Zero(t, res.Offset)
		assert.Len(t, res.Items, 2)
		assert.Equal(t, conv.ModelsToList([]*models.User{userFoo, userBar}), res)
	})

	checkFooOnDB := func(t *testing.T, id uint64, expectedEmail, expectedName string) *models.User {
		user1, err := models.FindUser(ctx, conn, id)
		assert.NoError(t, err)
		assert.Equal(t, expectedEmail, user1.Email)
		assert.Equal(t, expectedName, user1.Name)
		assert.NotEmpty(t, user1.FbauthUID)
		return user1
	}

	var bazID uint64
	t.Run("create baz first time", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "baz@example.com", Name: "Baz"})
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
		bazID = res.ID
		user := checkFooOnDB(t, bazID, "baz@example.com", "Baz")
		assert.Equal(t, conv.ModelToResult(user), res)
	})

	t.Run("create baz again", func(t *testing.T) {
		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "baz@example.com", Name: "Baz2"})
		assert.NoError(t, err)
		assert.Equal(t, bazID, res.ID)
		assert.Equal(t, "Baz", res.Name) // Baz2 にはならない
		user := checkFooOnDB(t, bazID, "baz@example.com", "Baz")
		assert.Equal(t, conv.ModelToResult(user), res)
	})

	t.Run("user on DB got gone", func(t *testing.T) {
		user1, err := models.FindUser(ctx, conn, bazID)
		assert.NoError(t, err)
		affected, err := user1.Delete(ctx, conn)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), affected)

		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: "baz@example.com", Name: "Baz3"})
		assert.NoError(t, err)
		assert.NotEqual(t, bazID, res.ID)
		assert.Equal(t, "Baz3", res.Name) // Baz3 になる
		user := checkFooOnDB(t, res.ID, "baz@example.com", "Baz3")
		assert.Equal(t, conv.ModelToResult(user), res)
	})

	t.Run("user on Firebase got gone", func(t *testing.T) {
		email := "qux@example.com"
		name := "Qux"
		newName := "Qux2"
		oldUID := "qux-uid"
		userQux := &models.User{Email: email, Name: name, FbauthUID: oldUID} // "user on Firebase got lost" テストで使用されます
		sqlboilertest.Insert(t, ctx, conn, boil.Infer(), userQux)

		res, err := srvc.Create(ctx, &users.UserCreatePayload{Email: email, Name: newName})
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, newName, res.Name) // Foo2 になる
		user := checkFooOnDB(t, res.ID, email, newName)
		assert.NotEqual(t, oldUID, user.FbauthUID) // 更新される
		assert.Equal(t, conv.ModelToResult(user), res)
	})
}

func TestUsersConvertor(t *testing.T) {
	now := time.Now()
	defer time.SetTime(now)

	conv := &UsersConvertor{}

	// Convertor はテストで期待する値を作成するためにも使うものなので、メソッド単体のテストが必要
	t.Run("ModelsToList", func(t *testing.T) {
		user1 := &models.User{ID: 1, Name: "test1", CreatedAt: now, UpdatedAt: now}
		user2 := &models.User{ID: 2, Name: "test2", CreatedAt: now, UpdatedAt: now}
		res := conv.ModelsToList([]*models.User{user1, user2})
		assert.Equal(t, &users.UserList{
			Total:  2,
			Offset: 0,
			Items: []*users.UserListItem{
				{ID: 1, Name: "test1"},
				{ID: 2, Name: "test2"},
			},
		}, res)
	})

	t.Run("ModelToResult", func(t *testing.T) {
		user1 := &models.User{ID: 1, Name: "test1", CreatedAt: now, UpdatedAt: now}
		res := conv.ModelToResult(user1)
		assert.Equal(t, &users.User{
			ID:        1,
			Name:      "test1",
			CreatedAt: now.Format(time.RFC3339),
			UpdatedAt: now.Format(time.RFC3339),
		}, res)
	})
}
