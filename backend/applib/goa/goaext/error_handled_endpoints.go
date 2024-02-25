package goaext

import (
	goaendpoints "applib/goa/endpoints"
)

func ErrorHandledEndpoints[T any](ts T, eh func(error) error) T {
	return goaendpoints.Wrap[T](ts, goaendpoints.ErrorHandler(eh))
}
