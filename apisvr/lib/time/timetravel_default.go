//go:build !timetravel

package time

import (
	"fmt"
	orig "time"
)

var nowFunc = orig.Now

func SetNowFunc(f func() Time) func() {
	panic(fmt.Errorf("timetravel is not enabled"))
}

func SetTime(t Time) func() {
	panic(fmt.Errorf("timetravel is not enabled"))
}

func ClearNanoSec(t Time) Time {
	panic(fmt.Errorf("timetravel is not enabled"))
}
