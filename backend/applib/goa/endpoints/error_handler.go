package goaendpoints

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

func ErrorHandler(eh func(error) error) func(goa.Endpoint) goa.Endpoint {
	return func(ep goa.Endpoint) goa.Endpoint {
		return func(ctx context.Context, req any) (any, error) {
			r, err := ep(ctx, req)
			if err != nil {
				err = eh(err)
			}
			return r, err
		}
	}
}
