package sql

import (
	"context"
	"errors"
)

type contextConnectionKey string

const ContextConnectionKey contextConnectionKey = "sql-connection"

func NewContextWithConnection(ctx context.Context, db *DB) context.Context {
	return context.WithValue(ctx, ContextConnectionKey, db)
}

var ErrNoConnectionInContext = errors.New("no connection in context")

func ConnectionFromContext(ctx context.Context) (*DB, error) {
	conn, ok := ctx.Value(ContextConnectionKey).(*DB)
	if !ok {
		return nil, ErrNoConnectionInContext
	}
	return conn, nil
}
