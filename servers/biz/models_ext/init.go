package modelsext

import (
	"applib/time"
	"biz/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	// Channel
	models.AddChannelHook(boil.BeforeInsertHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.Channel) error {
		now := time.Now()
		p.CreatedAt = now
		p.UpdatedAt = now
		return nil
	})
	models.AddChannelHook(boil.BeforeUpdateHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.Channel) error {
		p.UpdatedAt = time.Now()
		return nil
	})

	// ChatMessage
	models.AddChatMessageHook(boil.BeforeInsertHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.ChatMessage) error {
		now := time.Now()
		p.CreatedAt = now
		p.UpdatedAt = now
		return nil
	})
	models.AddChatMessageHook(boil.BeforeUpdateHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.ChatMessage) error {
		p.UpdatedAt = time.Now()
		return nil
	})

	// User
	models.AddUserHook(boil.BeforeInsertHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.User) error {
		now := time.Now()
		p.CreatedAt = now
		p.UpdatedAt = now
		return nil
	})
	models.AddUserHook(boil.BeforeUpdateHook, func(ctx context.Context, exec boil.ContextExecutor, p *models.User) error {
		p.UpdatedAt = time.Now()
		return nil
	})
}
