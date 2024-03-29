package time

import orig "time"

var nowFunc = orig.Now

func Now() orig.Time {
	return LocalNow()
}

func LocalNow() orig.Time {
	return nowFunc().In(defaultLocation)
}

func DefaultNow() orig.Time {
	return nowFunc()
}
