package auth

import (
	"apisvr/lib/firebase"
	"apisvr/lib/log"
	"context"
)

func NewClientRaw(ctx context.Context, app *firebase.App) (*OrigClient, error) {
	return app.Auth(ctx)
}

func NewClientWithLogger(ctx context.Context, app *firebase.App, logger *log.Logger) (*ClientLogger, error) {
	origClient, err := NewClientRaw(ctx, app)
	if err != nil {
		return nil, err
	}
	return NewClientLogger(NewClientErrorWrapper(origClient), logger), nil
}

type ClientFactoryFunc = func(ctx context.Context, app *firebase.App, logger *log.Logger) (Client, error)

var clientFactory ClientFactoryFunc = func(ctx context.Context, app *firebase.App, logger *log.Logger) (Client, error) {
	cli, err := NewClientWithLogger(ctx, app, logger)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
