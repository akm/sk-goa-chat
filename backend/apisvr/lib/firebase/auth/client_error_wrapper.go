package auth

import (
	"apisvr/lib/errors"
	"apisvr/lib/time"
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

func (c *ClientErrorWrapper) SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error) {
	res, err := c.impl.SessionCookie(ctx, idToken, expiresIn)
	if err != nil {
		return "", errors.Wrapf(err, "failed to SessionCookie %s", idToken)
	}
	return res, nil
}

// func (c *ClientErrorWrapper) VerifyIDToken(ctx context.Context, idToken string) (*Token, error) {
// 	res, err := c.impl.VerifyIDToken(ctx, idToken)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "failed to VerifyIDToken %s", idToken)
// 	}
// 	return res, nil
// }

// func (c *ClientErrorWrapper) VerifyIDTokenAndCheckRevoked(ctx context.Context, idToken string) (*Token, error) {
// 	res, err := c.impl.VerifyIDTokenAndCheckRevoked(ctx, idToken)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "failed to VerifyIDTokenAndCheckRevoked %s", idToken)
// 	}
// 	return res, nil
// }

func (c *ClientErrorWrapper) VerifySessionCookie(ctx context.Context, sessionCookie string) (*Token, error) {
	res, err := c.impl.VerifySessionCookie(ctx, sessionCookie)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to VerifySessionCookie %s", sessionCookie)
	}
	return res, nil
}

func (c *ClientErrorWrapper) VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*Token, error) {
	res, err := c.impl.VerifySessionCookieAndCheckRevoked(ctx, sessionCookie)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to VerifySessionCookieAndCheckRevoked %s", sessionCookie)
	}
	return res, nil
}

func (c *ClientErrorWrapper) RevokeRefreshTokens(ctx context.Context, uid string) error {
	err := c.impl.RevokeRefreshTokens(ctx, uid)
	if err != nil {
		return errors.Wrapf(err, "failed to RevokeRefreshTokens %s", uid)
	}
	return nil
}
