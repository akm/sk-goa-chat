package testidentitytoolkit

import (
	"apisvr/lib/errors"
	"context"
	"fmt"
	"os"
	"testing"

	"google.golang.org/api/identitytoolkit/v1"
	"google.golang.org/api/option"
)

// See
// https://zenn.dev/nbstsh/scraps/421f07c3e01c29
// https://pkg.go.dev/google.golang.org/api@v0.155.0/identitytoolkit/v1#AccountsService.SignInWithPassword
func GetIdToken(t testing.TB, ctx context.Context, email, password string, opts ...option.ClientOption) string {
	if len(opts) == 0 {
		opts = []option.ClientOption{}
		if host := os.Getenv("FIREBASE_AUTH_EMULATOR_HOST"); host != "" {
			opts = append(opts, option.WithEndpoint(fmt.Sprintf("http://%s/identitytoolkit.googleapis.com/v1", host)))
		} else {
			t.Fatalf("%+v\n", errors.Errorf("FIREBASE_AUTH_EMULATOR_HOST is not set"))
		}
		if apiKey := os.Getenv("FIREBASE_API_KEY"); apiKey != "" {
			opts = append(opts, option.WithAPIKey(apiKey))
		} else {
			t.Fatalf("%+v\n", errors.Errorf("FIREBASE_API_KEY is not set"))
		}
	}
	// option.WithEndpoint("http://localhost:9099/identitytoolkit.googleapis.com/v1"),
	// option.WithAPIKey("AIzaSyBRcOfsczLRC7VX0iirMU2dlAFKYPRo-9U"),
	identitytoolkitService, err := identitytoolkit.NewService(ctx, opts...)
	if err != nil {
		t.Fatal(err)
	}
	call := identitytoolkitService.Accounts.SignInWithPassword(&identitytoolkit.GoogleCloudIdentitytoolkitV1SignInWithPasswordRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	})
	res, err := call.Do()
	if err != nil {
		t.Fatal(err)
	}

	return res.IdToken
}
