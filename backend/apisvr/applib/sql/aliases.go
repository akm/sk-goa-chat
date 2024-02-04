package sql

import (
	orig "database/sql"
)

type (
	DB = orig.DB
	Tx = orig.Tx
)

var (
	OrigOpen = orig.Open

	ErrNoRows = orig.ErrNoRows
)
