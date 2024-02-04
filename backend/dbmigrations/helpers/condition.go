package helpers

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

var EmptyMigrationContext = func(context.Context, *sql.Tx) error { return nil }

type Condition struct {
	predicate func() bool
}

func NewCondition(predicate func() bool) *Condition {
	return &Condition{predicate: predicate}
}

func (c *Condition) AddNamedMigrationContext(up, down goose.GoMigrationContext) {
	_, filename, _, _ := runtime.Caller(1)
	if c.predicate() {
		goose.AddNamedMigrationContext(filename, up, down)
	} else {
		goose.AddNamedMigrationContext(filename, EmptyMigrationContext, EmptyMigrationContext)
	}
}
