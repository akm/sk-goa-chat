package testsql

import (
	"apisvr/applib/log"
	"apisvr/applib/sql"
	"testing"

	_ "apisvr/applib/sqlboiler"
)

func Open(t *testing.T, logger *log.Logger) *sql.DB {
	t.Helper()
	db, err := sql.Open(logger)
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	return db
}
