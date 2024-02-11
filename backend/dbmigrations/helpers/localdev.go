package helpers

import (
	"os"
	"strings"
)

var DemoData = NewCondition(func() bool {
	if os.Getenv("DEMO_DATA") != "" {
		return os.Getenv("DEMO_DATA") == "true"
	}
	return os.Getenv("STAGE_TYPE") == "local" &&
		!strings.Contains(os.Getenv("APP_ENV"), "test")
})
