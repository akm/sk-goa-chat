package modelsext

import (
	"apisvr/lib/time"
	"apisvr/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	models.AddChannelHook(boil.BeforeInsertHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.Channel) error {
		now := time.Now()
		p.CreatedAt = now
		p.UpdatedAt = now
		return nil
	})
}
