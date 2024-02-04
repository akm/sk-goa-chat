package chatapi

import (
	"apisvr/biz/testuser"
	"apisvr/models"
	chatmessages "apisvr/services/gen/chat_messages"
	log "apisvr/services/gen/log"
	"applib/collection"
	"applib/database/sql/sqltest"
	"applib/firebase/auth/authtest"
	"applib/goa/goatest"
	"applib/log/logtest"
	"applib/sqlboiler/sqlboilertest"
	"applib/time"
	"context"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestChaeMessages(t *testing.T) {
	logger := logtest.New(t)
	conn := sqltest.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := NewChatMessages(&log.Logger{Logger: logger})
	conv := NewChatMessageConvertor()

	fbauth := authtest.Setup(t, ctx)

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
	sqlboilertest.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
	assert.Equal(t, now, ch1.CreatedAt)

	newChatMessage := func(ch *models.Channel, user *testuser.User, content string) *models.ChatMessage {
		return &models.ChatMessage{ChannelID: ch.ID, UserID: null.Uint64From(user.Model.ID), UserName: user.Name, Content: content}
	}

	ch1Msg1 := newChatMessage(ch1, userFoo, "Hello")
	ch1Msg2 := newChatMessage(ch1, userBar, "Hi")
	ch1Msg3 := newChatMessage(ch1, userFoo, "Yo")
	sqlboilertest.Insert(t, ctx, conn, boil.Infer(), ch1Msg1, ch1Msg2, ch1Msg3)

	t.Run("list", func(t *testing.T) {
		res, err := srvc.List(ctx, &chatmessages.ListPayload{SessionID: sessionID, ChannelID: &ch1.ID, Limit: 50})
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
			goatest.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
		t.Run("too long content", func(t *testing.T) {
			// mediumtext は 最大長が 16,777,215 (2 ^ 24 − 1) 文字の TEXT カラム。
			// https://dev.mysql.com/doc/refman/8.0/ja/string-type-syntax.html
			res, err := srvc.Create(ctx, &chatmessages.ChatMessageCreatePayload{SessionID: sessionID, ChannelID: ch1.ID, Content: strings.Repeat("a", int(math.Pow(2, 24)))})
			goatest.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: 999, Content: "bonjour"})
			goatest.AssertServiceError(t, "not_found", err)
			assert.Nil(t, res)
		})
		t.Run("valid content", func(t *testing.T) {
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: ch1Msg1.ID, Content: "Ni hao"})
			assert.NoError(t, err)
			assert.Equal(t, conv.ModelToResult(&models.ChatMessage{ID: res.ID, ChannelID: ch1.ID, UserID: null.Uint64From(userFoo.Model.ID), UserName: userFoo.Name, Content: "Ni hao", CreatedAt: now, UpdatedAt: now}), res)
		})
		t.Run("empty content", func(t *testing.T) {
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: ch1Msg1.ID, Content: ""})
			goatest.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
		t.Run("too long content", func(t *testing.T) {
			// mediumtext は 最大長が 16,777,215 (2 ^ 24 − 1) 文字の TEXT カラム。
			// https://dev.mysql.com/doc/refman/8.0/ja/string-type-syntax.html
			res, err := srvc.Update(ctx, &chatmessages.ChatMessageUpdatePayload{SessionID: sessionID, ID: ch1Msg1.ID, Content: strings.Repeat("a", int(math.Pow(2, 24)))})
			goatest.AssertServiceError(t, "invalid_payload", err)
			assert.Nil(t, res)
		})
	})

	t.Run("delete", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			res, err := srvc.Delete(ctx, &chatmessages.DeletePayload{SessionID: sessionID, ID: 999})
			goatest.AssertServiceError(t, "not_found", err)
			assert.Nil(t, res)
		})
		t.Run("valid id", func(t *testing.T) {
			res, err := srvc.Delete(ctx, &chatmessages.DeletePayload{SessionID: sessionID, ID: ch1Msg2.ID})
			assert.NoError(t, err)
			assert.Equal(t, conv.ModelToResult(ch1Msg2), res)
		})
	})
}

func TestChaeMessagesList(t *testing.T) {
	logger := logtest.New(t)
	conn := sqltest.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	ctx := context.Background()
	srvc := NewChatMessages(&log.Logger{Logger: logger})
	conv := NewChatMessageConvertor()

	fbauth := authtest.Setup(t, ctx)

	userFoo := testuser.Foo().Setup(t, ctx, fbauth, conn)
	userBar := testuser.Bar().Setup(t, ctx, fbauth, conn)
	sessionID := userFoo.SessionID

	ch1 := &models.Channel{Name: "general", Visibility: models.ChannelsVisibilityPublic}
	ch2 := &models.Channel{Name: "random", Visibility: models.ChannelsVisibilityPublic}
	sqlboilertest.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
	assert.Equal(t, now, ch1.CreatedAt)

	newChatMessage := func(ch *models.Channel, user *testuser.User, content string) *models.ChatMessage {
		return &models.ChatMessage{ChannelID: ch.ID, UserID: null.Uint64From(user.Model.ID), UserName: user.Name, Content: content}
	}

	messages := []*models.ChatMessage{}
	for i := 0; i < 50; i++ {
		msgFooCh1 := newChatMessage(ch1, userFoo, fmt.Sprintf("ch1 msg from foo %02d", i+1))
		msgBarCh1 := newChatMessage(ch1, userBar, fmt.Sprintf("ch1 msg from bar %02d", i+1))
		msgFooCh2 := newChatMessage(ch2, userFoo, fmt.Sprintf("ch2 msg from foo %02d", i+1))
		msgBarCh2 := newChatMessage(ch2, userBar, fmt.Sprintf("ch2 msg from bar %02d", i+1))
		sqlboilertest.Insert(t, ctx, conn, boil.Infer(), msgFooCh1, msgBarCh1, msgFooCh2, msgBarCh2)
		messages = append(messages, msgFooCh1, msgBarCh1, msgFooCh2, msgBarCh2)
	}
	msgCount := len(messages)
	ch1Messages := collection.Filter[*models.ChatMessage](messages, func(m *models.ChatMessage) bool { return m.ChannelID == ch1.ID })

	type listPayload = chatmessages.ListPayload
	payload := func(funcs ...func(*listPayload)) *listPayload {
		r := &listPayload{SessionID: sessionID, Limit: 50}
		for _, f := range funcs {
			f(r)
		}
		return r
	}
	limit := func(limit int) func(*listPayload) { return func(p *listPayload) { p.Limit = limit } }
	after := func(after *uint64) func(*listPayload) { return func(p *listPayload) { p.After = after } }
	before := func(before *uint64) func(*listPayload) { return func(p *listPayload) { p.Before = before } }
	channelID := func(channelID *uint64) func(*listPayload) { return func(p *listPayload) { p.ChannelID = channelID } }

	t.Run("without options", func(t *testing.T) {
		res, err := srvc.List(ctx, payload())
		assert.NoError(t, err)
		assert.Len(t, res.Items, 50)
		assert.Equal(t, conv.ModelsToList(messages[msgCount-50:msgCount]), res)
	})

	t.Run("with limit", func(t *testing.T) {
		res, err := srvc.List(ctx, payload(limit(10)))
		assert.NoError(t, err)
		assert.Len(t, res.Items, 10)
		assert.Equal(t, conv.ModelsToList(messages[msgCount-10:msgCount]), res)
	})

	t.Run("with channel id", func(t *testing.T) {
		res, err := srvc.List(ctx, payload(channelID(&ch1.ID)))
		assert.NoError(t, err)
		assert.Len(t, res.Items, 50)
		assert.Equal(t, conv.ModelsToList(ch1Messages[len(ch1Messages)-50:]), res)
	})

	t.Run("with channel id and after", func(t *testing.T) {
		res, err := srvc.List(ctx, payload(channelID(&ch1.ID), after(&ch1Messages[25].ID)))
		assert.NoError(t, err)
		assert.Len(t, res.Items, 50)
		assert.Equal(t, conv.ModelsToList(ch1Messages[len(ch1Messages)-74:len(ch1Messages)-24]), res)
		assert.False(t, collection.Any[*chatmessages.ChatMessageListItem](res.Items, func(m *chatmessages.ChatMessageListItem) bool { return m.ID == ch1Messages[24].ID }))
		assert.False(t, collection.Any[*chatmessages.ChatMessageListItem](res.Items, func(m *chatmessages.ChatMessageListItem) bool { return m.ID == ch1Messages[25].ID }))
		assert.True(t, collection.Any[*chatmessages.ChatMessageListItem](res.Items, func(m *chatmessages.ChatMessageListItem) bool { return m.ID == ch1Messages[26].ID }))
	})
	t.Run("with channel id and before", func(t *testing.T) {
		assert.Len(t, ch1Messages, 100)
		res, err := srvc.List(ctx, payload(channelID(&ch1.ID), before(&ch1Messages[25].ID)))
		assert.NoError(t, err)
		assert.Len(t, res.Items, 25)
		assert.Equal(t, conv.ModelsToList(ch1Messages[:25]), res)
		assert.True(t, collection.Any[*chatmessages.ChatMessageListItem](res.Items, func(m *chatmessages.ChatMessageListItem) bool { return m.ID == ch1Messages[24].ID }))
		assert.False(t, collection.Any[*chatmessages.ChatMessageListItem](res.Items, func(m *chatmessages.ChatMessageListItem) bool { return m.ID == ch1Messages[25].ID }))
		assert.False(t, collection.Any[*chatmessages.ChatMessageListItem](res.Items, func(m *chatmessages.ChatMessageListItem) bool { return m.ID == ch1Messages[26].ID }))
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
