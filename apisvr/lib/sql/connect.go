package sql

import (
	"apisvr/lib/errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func Open() (*DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := OrigOpen("mysql", dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open database with %s", dsn)
	}
	// See https://github.com/simukti/sqldb-logger
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stderr))
	db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter) // db is STILL *sql.DB
	return db, nil
}
