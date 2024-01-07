package auth_test

import (
	"apisvr/lib/firebase"
	"apisvr/lib/firebase/auth"
	"apisvr/lib/firebase/errorutils"
	"apisvr/testlib/testfirebase/testauth"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/iterator"
)

func TestOrigClient(t *testing.T) {
	ctx := context.Background()
	fbapp, err := firebase.NewApp(ctx, nil)
	require.NoError(t, err)
	fbauth, err := auth.NewClientRaw(ctx, fbapp)
	require.NoError(t, err)

	t.Run("delete all of users before test", func(t *testing.T) {
		testauth.DeleteUsers(t, ctx, fbauth)
	})

	var fooUID string

	t.Run("create foo", func(t *testing.T) {
		args := &auth.UserToCreate{}
		args.Email("foo@example.com")
		args.DisplayName("Foo")
		args.Password("Passw0rd!")
		res, err := fbauth.CreateUser(ctx, args)
		require.NoError(t, err)
		// t.Logf("result: %+v", res)
		require.NotEmpty(t, res.UID)
		fooUID = res.UID
	})

	t.Run("get by uid", func(t *testing.T) {
		t.Run("registered uid", func(t *testing.T) {
			res, err := fbauth.GetUser(ctx, fooUID)
			require.NoError(t, err)
			// t.Logf("result: %+v", res)
			assert.NotEmpty(t, res.UID)
		})
		t.Run("not registered uid", func(t *testing.T) {
			res, err := fbauth.GetUser(ctx, "bar")
			require.Nil(t, res)
			require.Error(t, err)
			assert.True(t, errorutils.IsNotFound(err))
			assert.Equal(t, "no user exists with the uid: \"bar\"", err.Error())
			assert.Equal(t, "*internal.FirebaseError", fmt.Sprintf("%T", err))
		})
	})

	t.Run("get user by email", func(t *testing.T) {
		t.Run("registered email", func(t *testing.T) {
			res, err := fbauth.GetUserByEmail(ctx, "foo@example.com")
			require.NoError(t, err)
			// t.Logf("result: %+v", res)
			assert.NotEmpty(t, res.UID)
		})
		t.Run("not registered email", func(t *testing.T) {
			res, err := fbauth.GetUserByEmail(ctx, "bar@example.com")
			require.Nil(t, res)
			require.Error(t, err)
			// t.Logf("error: %+v", err)
			assert.True(t, errorutils.IsNotFound(err))
			assert.Equal(t, "no user exists with the email: \"bar@example.com\"", err.Error())
			assert.Equal(t, "*internal.FirebaseError", fmt.Sprintf("%T", err))
		})
	})

	t.Run("update user", func(t *testing.T) {
		t.Run("registered uid", func(t *testing.T) {
			args := &auth.UserToUpdate{}
			args.DisplayName("Foo Bar")
			res, err := fbauth.UpdateUser(ctx, fooUID, args)
			require.NoError(t, err)
			// t.Logf("result: %+v", res)
			assert.NotEmpty(t, res.UID)
		})
		t.Run("not registered uid", func(t *testing.T) {
			args := &auth.UserToUpdate{}
			args.DisplayName("Foo Bar")
			res, err := fbauth.UpdateUser(ctx, "baz", args)
			require.Nil(t, res)
			require.Error(t, err)
			assert.True(t, errorutils.IsNotFound(err))
			assert.Equal(t, "no user record found for the given identifier", err.Error())
			assert.Equal(t, "*internal.FirebaseError", fmt.Sprintf("%T", err))
		})
	})

	t.Run("create bar", func(t *testing.T) {
		args := &auth.UserToCreate{}
		args.Email("bar@example.com")
		args.DisplayName("Bar")
		args.Password("Passw0rd!")
		res, err := fbauth.CreateUser(ctx, args)
		require.NoError(t, err)
		// t.Logf("result: %+v", res)
		require.NotEmpty(t, res.UID)
		fooUID = res.UID
	})

	t.Run("list users", func(t *testing.T) {
		users := []*auth.ExportedUserRecord{}

		iter := fbauth.Users(ctx, "")
		for {
			user, err := iter.Next()
			if err != nil {
				if err == iterator.Done {
					break
				}
				require.NoError(t, err)
			}
			t.Logf("user: %+v", user)
			users = append(users, user)
		}

		require.Len(t, users, 2)
		for _, user := range users {
			assert.NotEmpty(t, user.UID)
			switch user.Email {
			case "foo@example.com":
				assert.Equal(t, "Foo Bar", user.DisplayName)
			case "bar@example.com":
				assert.Equal(t, "Bar", user.DisplayName)
			default:
				t.Errorf("unexpected user: %+v", user)
			}
		}
	})

	t.Run("delete user", func(t *testing.T) {
		t.Run("registered uid", func(t *testing.T) {
			err := fbauth.DeleteUser(ctx, fooUID)
			require.NoError(t, err)
			// t.Logf("result: %+v", res)
		})
		t.Run("not registered uid", func(t *testing.T) {
			err := fbauth.DeleteUser(ctx, "baz")
			require.Error(t, err)
			assert.True(t, errorutils.IsNotFound(err))
			assert.Equal(t, "no user record found for the given identifier", err.Error())
			assert.Equal(t, "*internal.FirebaseError", fmt.Sprintf("%T", err))
		})
	})
}
