package authtest

import (
	"applib/firebase/auth"
	"applib/google/identitytoolkit/identitytoolkittest"
	"applib/time"
	"context"
	"testing"

	"google.golang.org/api/option"
)

func GetSessionCookie(t testing.TB, ctx context.Context, authClient *auth.OrigClient, email, password string, opts ...option.ClientOption) string {
	idToken := identitytoolkittest.GetIdToken(t, ctx, email, password, opts...)
	cookie, err := authClient.SessionCookie(ctx, idToken, 1*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	return cookie
}
