package testsql

import (
	"apisvr/lib/sql"
	"apisvr/testlib/testdir"
	"path/filepath"
	"testing"

	"github.com/pressly/goose/v3"
)

// 呼び出し元で必ず defer Close() すること
func Setup(t *testing.T) *sql.DB {
	conn := Open(t)

	// 全テーブルをDROP
	DropAll(t, conn)

	// マイグレーションを実行
	rootPath := testdir.RootPath(t)
	if err := goose.Up(conn, filepath.Join(rootPath, "../dbmigrations")); err != nil {
		t.Fatalf("failed to migrate: %s", err)
	}

	return conn
}
