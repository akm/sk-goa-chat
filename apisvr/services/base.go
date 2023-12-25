package chatapi

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "apisvr/models_ext"
)

func SetupContext(ctx context.Context) context.Context {
	return boil.SkipTimestamps(ctx)
}
