package time

import "time"

func Now() time.Time {
	return nowFunc()
}
