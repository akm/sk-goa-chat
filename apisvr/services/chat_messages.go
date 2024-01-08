package chatapi

import (
	"apisvr/lib/time"
	"apisvr/models"
	chatmessages "apisvr/services/gen/chat_messages"
	log "apisvr/services/gen/log"
	"context"
)

// chat_messages service example implementation.
// The example methods log the requests and return zero values.
type chatMessagessrvc struct {
	logger *log.Logger
}

// NewChatMessages returns the chat_messages service implementation.
func NewChatMessages(logger *log.Logger) chatmessages.Service {
	return &chatMessagessrvc{logger}
}

// List implements list.
func (s *chatMessagessrvc) List(ctx context.Context, p *chatmessages.ListPayload) (res *chatmessages.ChatMessageList, err error) {
	res = &chatmessages.ChatMessageList{}
	s.logger.Info().Msg("chatMessages.list")
	return
}

// Show implements show.
func (s *chatMessagessrvc) Show(ctx context.Context, p *chatmessages.ShowPayload) (res *chatmessages.ChatMessage, err error) {
	res = &chatmessages.ChatMessage{}
	s.logger.Info().Msg("chatMessages.show")
	return
}

// Create implements create.
func (s *chatMessagessrvc) Create(ctx context.Context, p *chatmessages.ChatMessageCreatePayload) (res *chatmessages.ChatMessage, err error) {
	res = &chatmessages.ChatMessage{}
	s.logger.Info().Msg("chatMessages.create")
	return
}

// Update implements update.
func (s *chatMessagessrvc) Update(ctx context.Context, p *chatmessages.ChatMessageUpdatePayload) (res *chatmessages.ChatMessage, err error) {
	res = &chatmessages.ChatMessage{}
	s.logger.Info().Msg("chatMessages.update")
	return
}

// Delete implements delete.
func (s *chatMessagessrvc) Delete(ctx context.Context, p *chatmessages.DeletePayload) (res *chatmessages.ChatMessage, err error) {
	res = &chatmessages.ChatMessage{}
	s.logger.Info().Msg("chatMessages.delete")
	return
}

type ChatMessageConvertor struct{}

func NewChatMessageConvertor() *ChatMessageConvertor {
	return &ChatMessageConvertor{}
}

func (s *ChatMessageConvertor) ModelsToList(models []*models.ChatMessage) *chatmessages.ChatMessageList {
	items := s.ModelsToListItems(models)
	return &chatmessages.ChatMessageList{
		Items: items,
	}
}

func (*ChatMessageConvertor) ModelsToListItems(models []*models.ChatMessage) chatmessages.ChatMessageListItemCollection {
	items := make(chatmessages.ChatMessageListItemCollection, len(models))
	for i, model := range models {
		items[i] = &chatmessages.ChatMessageListItem{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Format(time.RFC3339),
			UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
			ChannelID: model.ChannelID,
			UserID:    model.UserID.Ptr(),
			UserName:  model.UserName,
			Content:   model.Content,
		}
	}
	return items
}

func (*ChatMessageConvertor) ModelToResult(model *models.ChatMessage) *chatmessages.ChatMessage {
	return &chatmessages.ChatMessage{
		ID:        model.ID,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
		ChannelID: model.ChannelID,
		UserID:    model.UserID.Ptr(),
		UserName:  model.UserName,
		Content:   model.Content,
	}
}