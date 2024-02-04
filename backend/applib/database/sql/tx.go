package sql

import (
	"applib/errors"
	"context"
)

func BeginTx(ctx context.Context, db *DB, cb func(context.Context, *Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := cb(ctx, tx); err != nil {
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return errors.Join(err, errors.Wrapf(rollbackError, "rollback error"))
		}
		return err
	}

	return tx.Commit()
}
