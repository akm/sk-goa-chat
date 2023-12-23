package testsql

import (
	"apisvr/lib/sql"
	"testing"

	_ "apisvr/lib/sqlboiler"
)

func Open(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open()
	if err != nil {
		t.Fatalf("failed to open database: %s", err)
	}
	return db
}
