package auth

import (
	"apisvr/lib/errors"
	"apisvr/lib/firebase"
	"apisvr/lib/log"
	"context"
)

func NewClientRaw(ctx context.Context, app *firebase.App) (*OrigClient, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to app.Auth")
	}
	return client, nil
}

func NewClientWithLogger(ctx context.Context, app *firebase.App, logger *log.Logger) (*ClientLogger, error) {
	origClient, err := NewClientRaw(ctx, app)
	if err != nil {
		return nil, err
	}
	return NewClientLogger(NewClientErrorWrapper(origClient), logger), nil
}

type ClientFactoryFunc = func(ctx context.Context, app *firebase.App, logger *log.Logger) (Client, error)

//nolint:unused
var clientFactory ClientFactoryFunc = func(ctx context.Context, app *firebase.App, logger *log.Logger) (Client, error) {
	cli, err := NewClientWithLogger(ctx, app, logger)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
