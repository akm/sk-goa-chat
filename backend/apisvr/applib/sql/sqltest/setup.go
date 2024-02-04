package sqltest

import (
	"apisvr/applib/log"
	"apisvr/applib/os/dirtest"
	"apisvr/applib/sql"
	"path/filepath"
	"testing"

	"github.com/pressly/goose/v3"

	_ "dbmigrations"
)

// 呼び出し元で必ず defer Close() すること
func Setup(t *testing.T, logger *log.Logger) *sql.DB {
	conn := Open(t, logger)

	// 全テーブルをDROP
	DropAll(t, conn)

	// マイグレーションを実行
	goose.SetBaseFS(nil)
	if err := goose.SetDialect("mysql"); err != nil {
		t.Fatalf("failed to set dialect: %s", err)
	}

	rootPath := dirtest.RootPath(t)
	if err := goose.Up(conn, filepath.Join(rootPath, "../dbmigrations")); err != nil {
		t.Fatalf("failed to migrate: %s", err)
	}

	return conn
}
