package auth

import (
	"context"
)

type Client interface {
	CreateUser(ctx context.Context, user *UserToCreate) (*UserRecord, error)
	GetUserByEmail(ctx context.Context, email string) (*UserRecord, error)

	// SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error)
	VerifyIDToken(ctx context.Context, idToken string) (*Token, error)
	VerifyIDTokenAndCheckRevoked(ctx context.Context, idToken string) (*Token, error)
	// VerifySessionCookie(ctx context.Context, sessionCookie string) (*Token, error)
	// VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*Token, error)
	RevokeRefreshTokens(ctx context.Context, uid string) error
}

var _ (Client) = (*OrigClient)(nil)
