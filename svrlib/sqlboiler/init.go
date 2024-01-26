package sqlboiler

import (
	"svrlib/time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	boil.SetLocation(loc)
}
