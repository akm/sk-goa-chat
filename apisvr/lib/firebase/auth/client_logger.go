package auth

import (
	"apisvr/lib/log"
	"context"
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
