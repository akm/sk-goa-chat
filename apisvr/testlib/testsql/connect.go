package testsql

import (
	"apisvr/lib/sql"
	"apisvr/services/gen/log"
	"testing"

	_ "apisvr/lib/sqlboiler"
)

func Open(t *testing.T, logger *log.Logger) *sql.DB {
	t.Helper()
	db, err := sql.Open(logger)
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	return db
}
