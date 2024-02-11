package helpers

import "os"

func IsLocalDev() bool {
	return os.Getenv("STAGE_ENV") == "localdev"
}

var DemoData = NewCondition(IsLocalDev)
