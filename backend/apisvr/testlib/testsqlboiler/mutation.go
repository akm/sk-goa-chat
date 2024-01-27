package testsqlboiler

import (
	"context"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type InsertModel interface {
	Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

func Insert(t *testing.T, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns, models ...InsertModel) {
	t.Helper()
	for _, m := range models {
		if err := m.Insert(ctx, exec, columns); err != nil {
			t.Fatal(err)
		}
	}
}

type UpdateModel interface {
	Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
}

func Update(t *testing.T, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns, models ...UpdateModel) {
	t.Helper()
	for _, m := range models {
		if _, err := m.Update(ctx, exec, columns); err != nil {
			t.Fatal(err)
		}
	}
}

type DeleteModel interface {
	Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

func Delete(t *testing.T, ctx context.Context, exec boil.ContextExecutor, models ...DeleteModel) {
	t.Helper()
	for _, m := range models {
		if _, err := m.Delete(ctx, exec); err != nil {
			t.Fatal(err)
		}
	}
}
