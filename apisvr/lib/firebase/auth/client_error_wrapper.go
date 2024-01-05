package auth

import (
	"apisvr/lib/errors"
	"context"
)

type ClientErrorWrapper struct {
	impl Client
}

var _ (Client) = (*ClientErrorWrapper)(nil)

func NewClientErrorWrapper(impl Client) *ClientErrorWrapper {
	return &ClientErrorWrapper{impl: impl}
}

func (c *ClientErrorWrapper) CreateUser(ctx context.Context, user *UserToCreate) (*UserRecord, error) {
	res, err := c.impl.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to CreateUser: %+v", user)
	}
	return res, nil
}

func (c *ClientErrorWrapper) GetUserByEmail(ctx context.Context, email string) (*UserRecord, error) {
	res, err := c.impl.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to GetUserByEmail %s", email)
	}
	return res, nil
}
