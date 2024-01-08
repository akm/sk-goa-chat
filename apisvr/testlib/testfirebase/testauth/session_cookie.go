package testauth

import (
	"apisvr/lib/firebase/auth"
	"apisvr/testlib/testgoogle/testidentitytoolkit"
	"context"
	"testing"
	"time"

	"google.golang.org/api/option"
)

func GetSessionCookie(t testing.TB, ctx context.Context, authClient *auth.OrigClient, email, password string, opts ...option.ClientOption) string {
	idToken := testidentitytoolkit.GetIdToken(t, ctx, email, password, opts...)
	cookie, err := authClient.SessionCookie(ctx, idToken, 1*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	return cookie
}
