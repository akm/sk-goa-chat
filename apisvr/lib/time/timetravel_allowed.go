//go:build timetravel

package time

import (
	orig "time"
)

var nowFunc = func() Time {
	return ClearNanoSec(orig.Now())
}

func SetNowFunc(f func() Time) func() {
	var backup func() Time
	nowFunc, backup = f, nowFunc
	return func() {
		nowFunc = backup
	}
}

func SetTime(t Time) func() {
	actual := ClearNanoSec(t)
	return SetNowFunc(func() Time { return actual })
}

func ClearNanoSec(t Time) Time {
	return orig.Unix(t.Unix(), 0).In(t.Location())
}
