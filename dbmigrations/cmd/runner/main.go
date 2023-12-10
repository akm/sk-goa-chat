package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"

	_ "dbmigrations"
)

var (
	flags        = flag.NewFlagSet("goose", flag.ExitOnError)
	dir          = flags.String("dir", ".", "directory with migration files")
	table        = flags.String("table", "", "database table to store the db version")
	allowMissing = flags.Bool("allow-missing", false, "allow missing migrations")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 3 {
		if len(args) > 0 && args[0] == "build-check" {
			fmt.Println("build-check: ok")
			return
		}
		flags.Usage()
		return
	}

	dialect, dsn, command := args[0], args[1], args[2]

	goose.SetBaseFS(nil)
	if err := goose.SetDialect(dialect); err != nil {
		log.Fatalf("goose: failed to set dialect: %v\n", err)
	}

	if table != nil {
		goose.SetTableName(*table)
	}

	optFuncs := []goose.OptionsFunc{}
	if allowMissing != nil && *allowMissing {
		optFuncs = append(optFuncs, goose.WithAllowMissing())
	}

	db, err := goose.OpenDBWithDriver("mysql", dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB with %q because of %v\n", dsn, err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.RunWithOptions(command, db, *dir, arguments, optFuncs...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
