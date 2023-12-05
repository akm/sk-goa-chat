package local

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

var emptyMigrationContext = func(context.Context, *sql.Tx) error { return nil }

func AddNamedMigrationContext(up, down goose.GoMigrationContext) {
	_, filename, _, _ := runtime.Caller(1)
	if IsLocal() {
		goose.AddNamedMigrationContext(filename, up, down)
	} else {
		goose.AddNamedMigrationContext(filename, emptyMigrationContext, emptyMigrationContext)
	}
}
