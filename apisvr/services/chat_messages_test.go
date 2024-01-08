package chatapi

import (
	"apisvr/lib/time"
	"apisvr/models"
	chatmessages "apisvr/services/gen/chat_messages"
	log "apisvr/services/gen/log"
	"apisvr/testlib/testfirebase/testauth"
	"apisvr/testlib/testgoa"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"apisvr/testlib/testsqlboiler"
	"apisvr/testlib/testuser"
	"context"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestChaeMessages(t *testing.T) {
	logger := testlog.New(t)
	conn := testsql.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := NewChatMessages(&log.Logger{Logger: logger})
	conv := NewChatMessageConvertor()

	fbauth := testauth.Setup(t, ctx)

	userFoo := testuser.Foo().Setup(t, ctx, fbauth, conn)
	userBar := testuser.Bar().Setup(t, ctx, fbauth, conn)
	sessionID := userFoo.SessionID

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			res, err := srvc.List(ctx, &chatmessages.ListPayload{SessionID: sessionID})
			assert.NoError(t, err)
			assert.Len(t, res.Items, 0)
		})
	})

	ch1 := &models.Channel{Name: "general", Visibility: models.ChannelsVisibilityPublic}
	ch2 := &models.Channel{Name: "random", Visibility: models.ChannelsVisibilityPublic}
	testsqlboiler.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
	assert.Equal(t, now, ch1.CreatedAt)

	newChatMessage := func(ch *models.Channel, user *testuser.User, content string) *models.ChatMessage {
		return &models.ChatMessage{ChannelID: ch.ID, UserID: null.Uint64From(user.Model.ID), UserName: user.Name, Content: content}
	}

	ch1Msg1 := newChatMessage(ch1, userFoo, "Hello")
	ch1Msg2 := newChatMessage(ch1, userBar, "Hi")
	ch1Msg3 := newChatMessage(ch1, userFoo, "Yo")
	testsqlboiler.Insert(t, ctx, conn, boil.Infer(), ch1Msg1, ch1Msg2, ch1Msg3)

	t.Run("list", func(t *testing.T) {
		res, err := srvc.List(ctx, &chatmessages.ListPayload{SessionID: sessionID, ChannelID: &ch1.ID})
		assert.NoError(t, err)
		assert.Len(t, res.Items, 3)
		assert.Equal(t, conv.ModelsToList([]*models.ChatMessage{ch1Msg1, ch1Msg2, ch1Msg3}), res)
	})

	t.Run("show", func(t *testing.T) {
		for _, ch := range []*models.ChatMessage{ch1Msg1, ch1Msg2, ch1Msg3} {
			t.Run(ch.Content, func(t *testing.T) {
				res, err := srvc.Show(ctx, &chatmessages.ShowPayload{SessionID: sessionID, ID: ch.ID})
				assert.NoError(t, err)
				assert.Equal(t, conv.ModelToResult(ch), res)
			})
		}
	})

	t.Run("create", func(t *testing.T) {
		t.Run("valid content", func(t *testing.T) {
			res, err := srvc.Create(ctx, &chatmessages.ChatMessageCreatePayload{SessionID: sessionID, ChannelID: ch1.ID, Content: "Hola"})
			assert.NoError(t, err)
			assert.Equal(t, conv.ModelToResult(&models.ChatMessage{ID: res.ID, ChannelID: ch1.ID, UserID: null.Uint64From(userFoo.Model.ID), UserName: userFoo.Name, Content: "Hola", CreatedAt: now, UpdatedAt: now}), res)
		})
		t.Run("empty content", func(t *testing.T) {
			res, err := srvc.Create(ctx, &chatmessages.ChatMessageCreatePayload{SessionID: sessionID, ChannelID: ch1.ID, Content: ""})
			testgoa.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
		t.Run("too long content", func(t *testing.T) {
			// mediumtext は 最大長が 16,777,215 (2 ^ 24 − 1) 文字の TEXT カラム。
			// https://dev.mysql.com/doc/refman/8.0/ja/string-type-syntax.html
			res, err := srvc.Create(ctx, &chatmessages.ChatMessageCreatePayload{SessionID: sessionID, ChannelID: ch1.ID, Content: strings.Repeat("a", int(math.Pow(2, 24)))})
			testgoa.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: 999, Content: "bonjour"})
			testgoa.AssertServiceError(t, "not_found", err)
			assert.Nil(t, res)
		})
		t.Run("valid content", func(t *testing.T) {
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: ch1Msg1.ID, Content: "Ni hao"})
			assert.NoError(t, err)
			assert.Equal(t, conv.ModelToResult(&models.ChatMessage{ID: res.ID, ChannelID: ch1.ID, UserID: null.Uint64From(userFoo.Model.ID), UserName: userFoo.Name, Content: "Ni hao", CreatedAt: now, UpdatedAt: now}), res)
		})
		t.Run("empty content", func(t *testing.T) {
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: ch1Msg1.ID, Content: ""})
			testgoa.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
		t.Run("too long content", func(t *testing.T) {
			// mediumtext は 最大長が 16,777,215 (2 ^ 24 − 1) 文字の TEXT カラム。
			// https://dev.mysql.com/doc/refman/8.0/ja/string-type-syntax.html
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: ch1Msg1.ID, Content: strings.Repeat("a", int(math.Pow(2, 24)))})
			testgoa.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
	})

	t.Run("delete", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			res, err := srvc.Delete(ctx, &chatmessages.DeletePayload{SessionID: sessionID, ID: 999})
			testgoa.AssertServiceError(t, "not_found", err)
			assert.Nil(t, res)
		})
		t.Run("valid id", func(t *testing.T) {
			res, err := srvc.Delete(ctx, &chatmessages.DeletePayload{SessionID: sessionID, ID: ch1Msg2.ID})
			assert.NoError(t, err)
			assert.Equal(t, conv.ModelToResult(ch1Msg2), res)
		})
	})
}

func TestChatMessageConvertor(t *testing.T) {
	now := time.Now()
	defer time.SetTime(now)

	conv := NewChatMessageConvertor()

	t.Run("ModelsToList", func(t *testing.T) {
		userFooID := uint64(1)
		userBarID := uint64(2)
		msg1 := &models.ChatMessage{ID: 1, ChannelID: 1, UserID: null.Uint64From(userFooID), UserName: "Foo", Content: "hello", CreatedAt: now, UpdatedAt: now}
		msg2 := &models.ChatMessage{ID: 2, ChannelID: 1, UserID: null.Uint64From(userBarID), UserName: "Bar", Content: "world", CreatedAt: now, UpdatedAt: now}
		msgs := []*models.ChatMessage{msg1, msg2}
		res := conv.ModelsToList(msgs)
		assert.Equal(t, &chatmessages.ChatMessageList{
			Items: chatmessages.ChatMessageListItemCollection{
				&chatmessages.ChatMessageListItem{ID: 1, ChannelID: 1, UserID: &userFooID, UserName: "Foo", Content: "hello", CreatedAt: now.Format(time.RFC3339), UpdatedAt: now.Format(time.RFC3339)},
				&chatmessages.ChatMessageListItem{ID: 2, ChannelID: 1, UserID: &userBarID, UserName: "Bar", Content: "world", CreatedAt: now.Format(time.RFC3339), UpdatedAt: now.Format(time.RFC3339)},
			},
		}, res)
	})

	t.Run("ModelToResult", func(t *testing.T) {
		userFooID := uint64(1)
		msg := &models.ChatMessage{ID: 1, ChannelID: 1, UserID: null.Uint64From(userFooID), UserName: "Foo", Content: "hello", CreatedAt: now, UpdatedAt: now}
		res := conv.ModelToResult(msg)
		assert.Equal(t, &chatmessages.ChatMessage{ID: 1, ChannelID: 1, UserID: &userFooID, UserName: "Foo", Content: "hello", CreatedAt: now.Format(time.RFC3339), UpdatedAt: now.Format(time.RFC3339)}, res)
	})
}
