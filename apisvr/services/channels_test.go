package chatapi

import (
	"apisvr/lib/sql"
	"apisvr/lib/time"
	"apisvr/models"
	"apisvr/services/gen/channels"
	"apisvr/services/gen/log"
	"apisvr/testlib/testgoa"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"apisvr/testlib/testsqlboiler"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestChannels(t *testing.T) {
	logger := testlog.New(t)
	conn := testsql.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := NewChannels(&log.Logger{Logger: logger})
	conv := NewChannelsConvertor()

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			res, err := srvc.List(ctx, &channels.ListPayload{})
			assert.NoError(t, err)
			assert.Zero(t, res.Total)
			assert.Zero(t, res.Offset)
			assert.Len(t, res.Items, 0)
		})
	})

	ch1 := &models.Channel{Name: "general", Visibility: models.ChannelsVisibilityPublic}
	ch2 := &models.Channel{Name: "random", Visibility: models.ChannelsVisibilityPublic}
	testsqlboiler.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
	assert.Equal(t, now, ch1.CreatedAt)

	t.Run("list", func(t *testing.T) {
		res, err := srvc.List(ctx, &channels.ListPayload{})
		assert.NoError(t, err)
		assert.Equal(t, uint64(2), res.Total)
		assert.Zero(t, res.Offset)
		assert.Equal(t, conv.ModelsToList([]*models.Channel{ch1, ch2}), res)
	})

	t.Run("show", func(t *testing.T) {
		for _, ch := range []*models.Channel{ch1, ch2} {
			t.Run(ch.Name, func(t *testing.T) {
				res, err := srvc.Show(ctx, &channels.ShowPayload{ID: ch.ID})
				assert.NoError(t, err)
				assert.Equal(t, conv.ModelToResult(ch), res)
			})
		}
		t.Run("not found", func(t *testing.T) {
			res, err := srvc.Show(ctx, &channels.ShowPayload{ID: 999})
			testgoa.AssertServiceError(t, "not_found", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})
	})

	t.Run("create", func(t *testing.T) {
		t.Run("valid name", func(t *testing.T) {
			name := "test1"
			res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: name})
			assert.NoError(t, err)
			ch := &models.Channel{ID: res.ID, Name: name, CreatedAt: now, UpdatedAt: now}
			assert.Equal(t, conv.ModelToResult(ch), res)
		})
		t.Run("empty name", func(t *testing.T) {
			res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: ""})
			testgoa.AssertServiceError(t, "invalid_payload", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})
		t.Run("too long name", func(t *testing.T) {
			res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: strings.Repeat("a", 256)})
			testgoa.AssertServiceError(t, "invalid_payload", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})

		t.Run("multi byte characters", func(t *testing.T) {
			maxMultiByteCharacters := strings.Repeat("薔", 255)
			t.Run("max", func(t *testing.T) {
				res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: maxMultiByteCharacters})
				assert.NoError(t, err)
				ch := &models.Channel{ID: res.ID, Name: maxMultiByteCharacters, CreatedAt: now, UpdatedAt: now}
				assert.Equal(t, conv.ModelToResult(ch), res)
			})
			t.Run("max plus 1", func(t *testing.T) {
				res, err := srvc.Create(ctx, &channels.ChannelCreatePayload{Name: maxMultiByteCharacters + "a"})
				testgoa.AssertServiceError(t, "invalid_payload", err) // ステータスコードは確認できない
				assert.Nil(t, res)
			})
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			res, err := srvc.Update(ctx, &channels.ChannelUpdatePayload{ID: 999, Name: "test"})
			testgoa.AssertServiceError(t, "not_found", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})
		t.Run("valid name", func(t *testing.T) {
			now, before := time.Now(), now
			defer time.SetTime(now)
			name := ch1.Name + "-dash"
			res, err := srvc.Update(ctx, &channels.ChannelUpdatePayload{ID: ch1.ID, Name: name})
			assert.NoError(t, err)
			ch := &models.Channel{ID: ch1.ID, Name: name, CreatedAt: before, UpdatedAt: now}
			assert.Equal(t, conv.ModelToResult(ch), res)
		})
		t.Run("empty name", func(t *testing.T) {
			res, err := srvc.Update(ctx, &channels.ChannelUpdatePayload{ID: ch1.ID, Name: ""})
			testgoa.AssertServiceError(t, "invalid_payload", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})
		t.Run("too long name", func(t *testing.T) {
			res, err := srvc.Update(ctx, &channels.ChannelUpdatePayload{ID: ch1.ID, Name: strings.Repeat("a", 256)})
			testgoa.AssertServiceError(t, "invalid_payload", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})

		t.Run("multi byte characters", func(t *testing.T) {
			maxMultiByteCharacters := strings.Repeat("薔", 255)
			t.Run("max", func(t *testing.T) {
				res, err := srvc.Update(ctx, &channels.ChannelUpdatePayload{ID: ch1.ID, Name: maxMultiByteCharacters})
				assert.NoError(t, err)
				ch := &models.Channel{ID: res.ID, Name: maxMultiByteCharacters, CreatedAt: now, UpdatedAt: now}
				assert.Equal(t, conv.ModelToResult(ch), res)
			})
			t.Run("max plus 1", func(t *testing.T) {
				res, err := srvc.Update(ctx, &channels.ChannelUpdatePayload{ID: ch1.ID, Name: maxMultiByteCharacters + "a"})
				testgoa.AssertServiceError(t, "invalid_payload", err) // ステータスコードは確認できない
				assert.Nil(t, res)
			})
		})
	})

	t.Run("delete", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			res, err := srvc.Delete(ctx, &channels.DeletePayload{ID: 999})
			testgoa.AssertServiceError(t, "not_found", err) // ステータスコードは確認できない
			assert.Nil(t, res)
		})
		t.Run("valid id", func(t *testing.T) {
			ch, err := models.FindChannel(ctx, conn, ch1.ID)
			assert.NoError(t, err)
			res, err := srvc.Delete(ctx, &channels.DeletePayload{ID: ch1.ID})
			assert.NoError(t, err)
			assert.Equal(t, conv.ModelToResult(ch), res)
			// DBから削除されていることを確認
			ch2, err := models.FindChannel(ctx, conn, ch1.ID)
			if assert.Error(t, err) {
				assert.Equal(t, sql.ErrNoRows, err)
			}
			assert.Nil(t, ch2)
		})
	})
}

func TestChannelsConvertor(t *testing.T) {
	now := time.Now()
	defer time.SetTime(now)

	conv := &ChannelsConvertor{}

	// Convertor はテストで期待する値を作成するためにも使うものなので、メソッド単体のテストが必要
	t.Run("ModelsToList", func(t *testing.T) {
		ch1 := &models.Channel{ID: 1, Name: "test1", CreatedAt: now, UpdatedAt: now}
		ch2 := &models.Channel{ID: 2, Name: "test2", CreatedAt: now, UpdatedAt: now}
		res := conv.ModelsToList([]*models.Channel{ch1, ch2})
		assert.Equal(t, &channels.ChannelList{
			Total:  2,
			Offset: 0,
			Items: []*channels.ChannelListItem{
				{ID: 1, Name: "test1", CreatedAt: now.Format(time.RFC3339), UpdatedAt: now.Format(time.RFC3339)},
				{ID: 2, Name: "test2", CreatedAt: now.Format(time.RFC3339), UpdatedAt: now.Format(time.RFC3339)},
			},
		}, res)
	})

	t.Run("ModelToResult", func(t *testing.T) {
		ch := &models.Channel{ID: 1, Name: "test1", CreatedAt: now, UpdatedAt: now}
		res := conv.ModelToResult(ch)
		assert.Equal(t, &channels.Channel{
			ID:        1,
			Name:      "test1",
			CreatedAt: now.Format(time.RFC3339),
			UpdatedAt: now.Format(time.RFC3339),
		}, res)
	})
}
