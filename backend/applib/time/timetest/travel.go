package timetest

import (
	"applib/time"
	orig "time"
)

var nowFunc = func() time.Time {
	return ClearNanoSec(orig.Now())
}

func init() {
	time.SetNowFunc(nowFunc)
}

func SetNow(t time.Time) func() {
	actual := ClearNanoSec(t)
	return time.SetNowFunc(func() time.Time { return actual })
}

func ClearNanoSec(t time.Time) time.Time {
	return orig.Unix(t.Unix(), 0).In(t.Location())
}
