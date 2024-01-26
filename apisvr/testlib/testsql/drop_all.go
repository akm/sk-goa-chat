package testsql

import (
	"slices"
	"svrlib/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

var skipTableNames = []string{
	"goose_db_version",
}

// FKの依存関係を考慮して全てのテーブルをDROPします
func DropAll(t *testing.T, conn *sql.DB) {
	tableNames := SortedTableNames(t, conn)
	errs := []error{}
	for _, tbl := range tableNames {
		if slices.Contains[[]string, string](skipTableNames, tbl) {
			t.Logf("DELETE FROM\t%s\ttrying\n", tbl)
			_, err := conn.Exec("DELETE FROM " + tbl + " WHERE id > 1")
			if !assert.NoError(t, err) {
				t.Logf("DELETE FROM\t%s\tERROR\n", tbl)
				errs = append(errs, err)
			} else {
				t.Logf("DELETE FROM\t%s\tSUCCESS\n", tbl)
			}
			continue
		}

		t.Logf("DROP TABLE\t%s\ttrying\n", tbl)
		_, err := conn.Exec("DROP TABLE IF EXISTS " + tbl)
		if !assert.NoError(t, err) {
			t.Logf("DROP TABLE\t%s\tERROR\n", tbl)
			errs = append(errs, err)
		} else {
			t.Logf("DROP TABLE\t%s\tSUCCESS\n", tbl)
		}
	}
	if len(errs) > 0 {
		t.Logf("DROP TABLE\tERROR\n%+v\n", errs)
		t.Fatalf("failed to drop tables: %+v", errs)
	} else {
		t.Logf("DROP TABLE\tCOMPLETE\n")
	}
}
