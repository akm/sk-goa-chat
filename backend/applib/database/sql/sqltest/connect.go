package sqltest

import (
	"applib/database/sql"
	"applib/log"
	"testing"

	_ "applib/sqlboiler"
)

func Open(t *testing.T, logger *log.Logger) *sql.DB {
	t.Helper()
	db, err := sql.Open(logger)
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	return db
}
