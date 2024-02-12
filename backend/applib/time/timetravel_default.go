//go:build !timetravel

package time

import (
	"fmt"
)

func SetNowFunc(f func() Time) func() {
	panic(fmt.Errorf("timetravel is not enabled"))
}
