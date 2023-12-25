package time

import (
	orig "time"
)

type (
	Time     = orig.Time
	Duration = orig.Duration
	Location = orig.Location
)

const (
	Nanosecond  = orig.Nanosecond
	Microsecond = orig.Microsecond
	Millisecond = orig.Millisecond
	Second      = orig.Second
	Minute      = orig.Minute
	Hour        = orig.Hour

	RFC3339 = orig.RFC3339
)
