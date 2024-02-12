//go:build timetravel

package time

func SetNowFunc(f func() Time) func() {
	var backup func() Time
	nowFunc, backup = f, nowFunc
	return func() {
		nowFunc = backup
	}
}
