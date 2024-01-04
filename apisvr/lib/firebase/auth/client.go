package auth

import "context"

type Client interface {
	CreateUser(ctx context.Context, user *UserToCreate) (*UserRecord, error)
	GetUserByEmail(ctx context.Context, email string) (*UserRecord, error)
}

var _ (Client) = (*OrigClient)(nil)
