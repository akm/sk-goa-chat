package local

import "os"

func IsLocal() bool {
	return os.Getenv("STAGE") == "local"
}
