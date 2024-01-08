package chatapi

import (
	"apisvr/lib/sql"
	"apisvr/lib/time"
	"apisvr/models"
	chatmessages "apisvr/services/gen/chat_messages"
	log "apisvr/services/gen/log"
	"context"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// chat_messages service example implementation.
// The example methods log the requests and return zero values.
type chatMessagessrvc struct {
	baseAuthService
	*ChatMessageConvertor
}

// NewChatMessages returns the chat_messages service implementation.
func NewChatMessages(logger *log.Logger) chatmessages.Service {
	return &chatMessagessrvc{baseAuthService: newBaseAuthService(logger), ChatMessageConvertor: NewChatMessageConvertor()}
}

// List implements list.
func (s *chatMessagessrvc) List(ctx context.Context, p *chatmessages.ListPayload) (res *chatmessages.ChatMessageList, err error) {
	err = s.actionWithAuth(ctx, "chatMessages.list", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		results, err := models.ChatMessages().All(ctx, db)
		if err != nil {
			return err
		}
		res = s.ModelsToList(results)
		return nil
	})
	return
}

// Show implements show.
func (s *chatMessagessrvc) Show(ctx context.Context, p *chatmessages.ShowPayload) (res *chatmessages.ChatMessage, err error) {
	err = s.actionWithAuth(ctx, "chatMessages.show", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		m, err := models.FindChatMessage(ctx, db, p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return chatmessages.MakeNotFound(fmt.Errorf("chat message not found"))
			}
			return err
		}
		res = s.ModelToResult(m)
		return nil
	})
	return
}

// Create implements create.
const maxContentLength = 16_777_215

func (s *chatMessagessrvc) Create(ctx context.Context, p *chatmessages.ChatMessageCreatePayload) (res *chatmessages.ChatMessage, err error) {
	err = s.actionWithAuth(ctx, "chatMessages.create", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		if p.Content == "" {
			return chatmessages.MakeInvalidPayload(fmt.Errorf("content is required"))
		} else {
			runes := []rune(p.Content)
			if len(runes) > maxContentLength {
				return chatmessages.MakeInvalidPayload(fmt.Errorf("content is too long"))
			}
		}
		ch, err := models.FindChannel(ctx, db, p.ChannelID)
		if err != nil {
			if err == sql.ErrNoRows {
				return chatmessages.MakeInvalidPayload(fmt.Errorf("channel not found"))
			}
		}
		m := &models.ChatMessage{
			ChannelID: ch.ID,
			UserID:    null.Uint64From(user.ID),
			UserName:  user.Name,
			Content:   p.Content,
		}
		if err := m.Insert(ctx, db, boil.Infer()); err != nil {
			return err
		}
		res = s.ModelToResult(m)
		return nil
	})
	return
}

// Update implements update.
func (s *chatMessagessrvc) Update(ctx context.Context, p *chatmessages.ChatMessageUpdatePayload) (res *chatmessages.ChatMessage, err error) {
	err = s.actionWithAuth(ctx, "chatMessages.update", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		if p.Content == "" {
			return chatmessages.MakeInvalidPayload(fmt.Errorf("content is required"))
		} else {
			runes := []rune(p.Content)
			if len(runes) > maxContentLength {
				return chatmessages.MakeInvalidPayload(fmt.Errorf("content is too long"))
			}
		}
		m, err := models.FindChatMessage(ctx, db, p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return chatmessages.MakeNotFound(fmt.Errorf("chat message not found"))
			}
			return err
		}
		m.Content = p.Content
		if _, err := m.Update(ctx, db, boil.Infer()); err != nil {
			return err
		}
		res = s.ModelToResult(m)
		return nil
	})
	return
}

// Delete implements delete.
func (s *chatMessagessrvc) Delete(ctx context.Context, p *chatmessages.DeletePayload) (res *chatmessages.ChatMessage, err error) {
	err = s.actionWithAuth(ctx, "chatMessages.delete", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		m, err := models.FindChatMessage(ctx, db, p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return chatmessages.MakeNotFound(fmt.Errorf("chat message not found"))
			}
			return err
		}
		if _, err := m.Delete(ctx, db); err != nil {
			return err
		}
		res = s.ModelToResult(m)
		return nil
	})
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
