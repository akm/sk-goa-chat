package sql

import (
	"applib/errors"
	"applib/log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func Open(logger *log.Logger) (*DB, error) {
	dsn := os.Getenv("APP_MYSQL_DSN")
	db, err := OrigOpen("mysql", dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open database with %s", dsn)
	}
	// See https://github.com/simukti/sqldb-logger
	loggerAdapter := zerologadapter.New(*logger)
	db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter) // db is STILL *sql.DB
	return db, nil
}
