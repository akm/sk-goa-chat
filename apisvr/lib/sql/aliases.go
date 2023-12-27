package sql

import (
	orig "database/sql"
)

type (
	DB = orig.DB
)

var (
	OrigOpen = orig.Open

	ErrNoRows = orig.ErrNoRows
)
