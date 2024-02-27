package goasql

import (
	"applib/database/sql"
	"applib/log"
	"context"

	goa "goa.design/goa/v3/pkg"
)

func ConnectionEndpointWrapper(logger *log.Logger) func(goa.Endpoint) goa.Endpoint {
	return func(ep goa.Endpoint) goa.Endpoint {
		return func(ctx context.Context, req any) (any, error) {
			db, err := sql.Open(logger)
			if err != nil {
				return nil, err
			}
			defer db.Close()
			return ep(sql.NewContextWithConnection(ctx, db), req)
		}
	}
}
