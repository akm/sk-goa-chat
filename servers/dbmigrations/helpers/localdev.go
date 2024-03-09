package helpers

import (
	"os"
	"strings"
)

var DemoData = NewCondition(func() bool {
	if os.Getenv("DEMO_DATA") != "" {
		return os.Getenv("DEMO_DATA") == "true"
	}
	appEnv := os.Getenv("APP_ENV")
	return os.Getenv("APP_STAGE_TYPE") == "local" &&
		!strings.Contains(appEnv, "test") ||
		appEnv == "e2e_test"
})
