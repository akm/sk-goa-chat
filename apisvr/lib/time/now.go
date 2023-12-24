package time

import "time"

var nowFunc = time.Now

func Now() time.Time {
	return nowFunc()
}
