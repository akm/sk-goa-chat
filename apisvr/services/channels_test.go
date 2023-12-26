package chatapi

import (
	"apisvr/lib/time"
	"apisvr/models"
	"apisvr/services/gen/channels"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"apisvr/testlib/testsqlboiler"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestChannels(t *testing.T) {
	conn := testsql.Setup(t)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := channelssrvc{logger: testlog.New(t)}

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			res, err := srvc.List(ctx)
			assert.NoError(t, err)
			assert.Zero(t, res.Total)
			assert.Zero(t, res.Offset)
			assert.Len(t, res.Items, 0)
		})
	})

	t.Run("with data", func(t *testing.T) {
		ch1 := &models.Channel{Name: "general", Visibility: models.ChannelsVisibilityPublic}
		ch2 := &models.Channel{Name: "random", Visibility: models.ChannelsVisibilityPublic}
		testsqlboiler.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
		assert.Equal(t, now, ch1.CreatedAt)

		t.Run("list", func(t *testing.T) {
			res, err := srvc.List(ctx)
			assert.NoError(t, err)
			assert.Equal(t, uint64(2), res.Total)
			assert.Zero(t, res.Offset)
			assert.Equal(t, srvc.ConvertModelsToListResult([]*models.Channel{ch1, ch2}), res)
		})

		t.Run("show", func(t *testing.T) {
			for _, ch := range []*models.Channel{ch1, ch2} {
				t.Run(ch.Name, func(t *testing.T) {
					res, err := srvc.Show(ctx, &channels.ShowPayload{ID: ch.ID})
					assert.NoError(t, err)
					assert.Equal(t, srvc.ConvertModelToResult(ch), res)
				})
			}
			t.Run("not found", func(t *testing.T) {
				res, err := srvc.Show(ctx, &channels.ShowPayload{ID: 999})
				assert.Error(t, err)
				assert.Nil(t, res)
			})
		})

		t.Run("create", func(t *testing.T) {
			t.Run("valid name", func(t *testing.T) {
				name := "test1"
				res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: name})
				assert.NoError(t, err)
				ch := &models.Channel{ID: res.ID, Name: name, CreatedAt: now, UpdatedAt: now}
				assert.Equal(t, srvc.ConvertModelToResult(ch), res)
			})
			t.Run("empty name", func(t *testing.T) {
				res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: ""})
				assert.Error(t, err)
				assert.Nil(t, res)
			})
		})
	})
}
