package chatapi

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func SetupContext(ctx context.Context) context.Context {
	return boil.SkipTimestamps(ctx)
}
