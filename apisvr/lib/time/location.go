package time

var defaultLocation = func() *Location {
	loc, err := LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return loc
}()
