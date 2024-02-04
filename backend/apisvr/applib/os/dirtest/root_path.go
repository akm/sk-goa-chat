package dirtest

import (
	"apisvr/applib/errors"
	"os"
	"path/filepath"
	"testing"
)

func RootPath(t *testing.T) string {
	dir, err := GetRootPath()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	return dir
}
func GetRootPath() (string, error) {
	// テスト実行時のカレントディレクトリはテストのファイルがあるディレクトリ
	initialDir, err := os.Getwd()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get current directory")
	}

	dir := initialDir
	for {
		path := filepath.Join(dir, "go.mod")
		// fmt.Printf("path: %s\n", path)
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				dir = filepath.Dir(dir)
				if dir == "." || dir == "/" {
					return "", errors.Errorf("go.mod not found for %s", initialDir)
				}
			} else {
				return "", errors.Wrapf(err, "failed to stat %s", path)
			}
		} else {
			break
		}
	}
	return dir, nil
}
