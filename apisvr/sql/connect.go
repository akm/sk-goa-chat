package sql

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"apisvr/errors"
)

func Open() (*DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := OrigOpen("mysql", dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open database with %s", dsn)
	}
	return db, nil
}
