package sql

import "context"

func BeginTx(ctx context.Context, db *DB, cb func(context.Context, *Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := cb(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
