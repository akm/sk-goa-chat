package testauth

import (
	"context"
	"testing"

	origauth "firebase.google.com/go/v4/auth"

	"github.com/stretchr/testify/require"
	"google.golang.org/api/iterator"
)

func DeleteUsers(t *testing.T, ctx context.Context, fbauth *origauth.Client) {
	t.Helper()

	uids := []string{}
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
		uids = append(uids, user.UID)
	}

	if len(uids) > 0 {
		result, err := fbauth.DeleteUsers(ctx, uids)
		require.NoError(t, err)
		t.Logf("result: %+v", result)
	}
}
