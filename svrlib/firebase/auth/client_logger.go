package auth

import (
	"context"
	"svrlib/log"
	"svrlib/time"
)

type ClientLogger struct {
	impl   Client
	logger *log.Logger
}

var _ (Client) = (*ClientLogger)(nil)

func NewClientLogger(impl Client, logger *log.Logger) *ClientLogger {
	return &ClientLogger{impl: impl, logger: logger}
}

func (c *ClientLogger) CreateUser(ctx context.Context, user *UserToCreate) (*UserRecord, error) {
	c.logger.Info().Msgf("CreateUser calling: %+v", user)
	res, err := c.impl.CreateUser(ctx, user)
	if err != nil {
		// c.logger.Error().Err(err).Msgf("CreateUser: %+v", user)
		return nil, err
	} else {
		c.logger.Info().Msgf("CreateUser result: %+v", res)
		return res, nil
	}
}

func (c *ClientLogger) GetUserByEmail(ctx context.Context, email string) (*UserRecord, error) {
	c.logger.Info().Msgf("GetUserByEmail calling %s", email)
	res, err := c.impl.GetUserByEmail(ctx, email)
	if err != nil {
		// c.logger.Error().Err(err).Msgf("GetUserByEmail %s", email)
		return nil, err
	} else {
		c.logger.Info().Msgf("GetUserByEmail result %s: %+v", email, res)
		return res, nil
	}
}

// func (c *ClientLogger) UpdateUser(ctx context.Context, uid string, user *UserToUpdate) (ur *UserRecord, err error) {
// 	c.logger.Info().Msgf("UpdateUser calling %s: %+v", uid, user)
// 	res, err := c.impl.UpdateUser(ctx, uid, user)
// 	if err != nil {
// 		c.logger.Error().Err(err).Msgf("UpdateUser %s : %+v", uid, user)
// 		return nil, err
// 	} else {
// 		c.logger.Info().Msgf("UpdateUser result %s: %+v", uid, res)
// 		return res, nil
// 	}
// }

// func (c *ClientLogger) DeleteUser(ctx context.Context, uid string) error {
// 	c.logger.Info().Msgf("DeleteUser calling %s", uid)
// 	err := c.impl.DeleteUser(ctx, uid)
// 	if err != nil {
// 		c.logger.Error().Err(err).Msgf("DeleteUser %s", uid)
// 		return err
// 	} else {
// 		c.logger.Info().Msgf("DeleteUser result %s", uid)
// 		return nil
// 	}
// }

func (c *ClientLogger) SessionCookie(ctx context.Context, idToken string, expiresIn time.Duration) (string, error) {
	c.logger.Info().Msgf("SessionCookie calling %s", idToken)
	res, err := c.impl.SessionCookie(ctx, idToken, expiresIn)
	if err != nil {
		c.logger.Error().Err(err).Msgf("SessionCookie %s", idToken)
		return "", err
	} else {
		c.logger.Info().Msgf("SessionCookie result %s", idToken)
		return res, nil
	}
}

// func (c *ClientLogger) VerifyIDToken(ctx context.Context, idToken string) (*Token, error) {
// 	c.logger.Info().Msgf("VerifyIDToken calling %s", idToken)
// 	res, err := c.impl.VerifyIDToken(ctx, idToken)
// 	if err != nil {
// 		c.logger.Error().Err(err).Msgf("VerifyIDToken %s", idToken)
// 		return nil, err
// 	} else {
// 		c.logger.Info().Msgf("VerifyIDToken result %s", idToken)
// 		return res, nil
// 	}
// }

// func (c *ClientLogger) VerifyIDTokenAndCheckRevoked(ctx context.Context, idToken string) (*Token, error) {
// 	c.logger.Info().Msgf("VerifyIDTokenAndCheckRevoked calling %s", idToken)
// 	res, err := c.impl.VerifyIDTokenAndCheckRevoked(ctx, idToken)
// 	if err != nil {
// 		c.logger.Error().Err(err).Msgf("VerifyIDTokenAndCheckRevoked %s", idToken)
// 		return nil, err
// 	} else {
// 		c.logger.Info().Msgf("VerifyIDTokenAndCheckRevoked result %s", idToken)
// 		return res, nil
// 	}
// }

func (c *ClientLogger) VerifySessionCookie(ctx context.Context, sessionCookie string) (*Token, error) {
	c.logger.Info().Msgf("VerifySessionCookie calling %s", sessionCookie)
	res, err := c.impl.VerifySessionCookie(ctx, sessionCookie)
	if err != nil {
		c.logger.Error().Err(err).Msgf("VerifySessionCookie %s", sessionCookie)
		return nil, err
	} else {
		c.logger.Info().Msgf("VerifySessionCookie result %s", sessionCookie)
		return res, nil
	}
}

func (c *ClientLogger) VerifySessionCookieAndCheckRevoked(ctx context.Context, sessionCookie string) (*Token, error) {
	c.logger.Info().Msgf("VerifySessionCookieAndCheckRevoked calling %s", sessionCookie)
	res, err := c.impl.VerifySessionCookieAndCheckRevoked(ctx, sessionCookie)
	if err != nil {
		c.logger.Error().Err(err).Msgf("VerifySessionCookieAndCheckRevoked %s", sessionCookie)
		return nil, err
	} else {
		c.logger.Info().Msgf("VerifySessionCookieAndCheckRevoked result %s", sessionCookie)
		return res, nil
	}
}

func (c *ClientLogger) RevokeRefreshTokens(ctx context.Context, uid string) error {
	c.logger.Info().Msgf("RevokeRefreshTokens calling %s", uid)
	err := c.impl.RevokeRefreshTokens(ctx, uid)
	if err != nil {
		c.logger.Error().Err(err).Msgf("RevokeRefreshTokens %s", uid)
		return err
	} else {
		c.logger.Info().Msgf("RevokeRefreshTokens result %s", uid)
		return nil
	}
}
