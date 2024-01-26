package testauth

import (
	"context"
	"svrlib/firebase"
	"svrlib/firebase/auth"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewClient(t *testing.T, ctx context.Context) *auth.OrigClient {
	fbapp, err := firebase.NewApp(ctx, nil)
	require.NoError(t, err)
	fbauth, err := auth.NewClientRaw(ctx, fbapp)
	require.NoError(t, err)
	return fbauth
}

func Setup(t *testing.T, ctx context.Context) *auth.OrigClient {
	fbauth := NewClient(t, ctx)
	DeleteUsers(t, ctx, fbauth)
	return fbauth
}
