package chatapi

import (
	"applib/errors"
	"biz/models"
	"context"
)

type userContextKey string

const UserContextKey userContextKey = "user"

func NewContextWithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

func UserFromContext(ctx context.Context) (*models.User, error) {
	v := ctx.Value(UserContextKey)
	if v == nil {
		return nil, errors.Errorf("user not found")
	}
	return v.(*models.User), nil
}
