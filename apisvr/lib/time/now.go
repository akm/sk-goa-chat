package time

import orig "time"

func Now() orig.Time {
	return nowFunc()
}
