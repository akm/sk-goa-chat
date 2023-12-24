//go:build timetravel

package time

import (
	orig "time"
)

func SetNowFunc(f func() orig.Time) func() {
	var backup func() orig.Time
	nowFunc, backup = f, nowFunc
	return func() {
		nowFunc = backup
	}
}

func SetTime(t orig.Time) func() {
	actual := ClearNanoSec(t)
	return SetNowFunc(func() orig.Time { return actual })
}

func ClearNanoSec(t orig.Time) orig.Time {
	return orig.Unix(t.Unix(), 0).In(t.Location())
}
