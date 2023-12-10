package dbmigrations

import (
	"context"
	"database/sql"

	"dbmigrations/local"
)

func init() {
	local.AddNamedMigrationContext(upGo, downGo)
}

func upGo(ctx context.Context, tx *sql.Tx) error {
	statement :=
		"INSERT INTO channels (id, name, visibility) VALUES " +
			"(1, 'general', 'public')," +
			"(2, 'random', 'private');"
	if _, err := tx.Exec(statement); err != nil {
		return err
	}
	return nil
}

func downGo(ctx context.Context, tx *sql.Tx) error {
	statement := "DELETE FROM channels WHERE id IN (1,2);"
	if _, err := tx.Exec(statement); err != nil {
		return err
	}
	return nil
}
