package chatapi

import (
	"apisvr/lib/time"
	"apisvr/models"
	chatmessages "apisvr/services/gen/chat_messages"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

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
