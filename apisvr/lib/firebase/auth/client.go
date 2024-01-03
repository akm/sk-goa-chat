package auth

import "context"

type Client interface {
	CreateUser(ctx context.Context, user *UserToCreate) (*UserRecord, error)
	UpdateUser(ctx context.Context, uid string, user *UserToUpdate) (ur *UserRecord, err error)
	DeleteUser(ctx context.Context, uid string) error
}

var _ (Client) = (*OrigClient)(nil)
