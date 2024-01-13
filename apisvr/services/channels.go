package chatapi

import (
	"apisvr/lib/sql"
	"apisvr/lib/time"
	"apisvr/models"
	channels "apisvr/services/gen/channels"
	log "apisvr/services/gen/log"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// channels service example implementation.
// The example methods log the requests and return zero values.
type channelssrvc struct {
	baseAuthService
	*ChannelsConvertor
}

// NewChannels returns the channels service implementation.
func NewChannels(logger *log.Logger) channels.Service {
	return &channelssrvc{
		baseAuthService: newBaseAuthService(
			logger,
			channels.MakeUnauthenticated,
		),
		ChannelsConvertor: NewChannelsConvertor(),
	}
}

// List implements list.
func (s *channelssrvc) List(ctx context.Context, p *channels.ListPayload) (res *channels.ChannelList, err error) {
	err = s.actionWithAuth(ctx, "channels.list", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		results, err := models.Channels().All(ctx, db)
		if err != nil {
			return err
		}
		res = s.ModelsToList(results)
		return nil
	})
	return
}

// Show implements show.
func (s *channelssrvc) Show(ctx context.Context, p *channels.ShowPayload) (res *channels.Channel, err error) {
	err = s.actionWithAuth(ctx, "channels.show", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		m, err := models.FindChannel(ctx, db, p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return channels.MakeNotFound(fmt.Errorf("channel not found"))
			}
			return err
		}

		res = s.ModelToResult(m)
		return nil
	})
	return
}

// Create implements create.
func (s *channelssrvc) Create(ctx context.Context, p *channels.ChannelCreatePayload) (res *channels.Channel, err error) {
	err = s.actionWithAuth(ctx, "channels.create", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		if p.Name == "" {
			return channels.MakeInvalidPayload(fmt.Errorf("name is required"))
		} else {
			runes := []rune(p.Name)
			if len(runes) > 255 {
				return channels.MakeInvalidPayload(fmt.Errorf("name is too long"))
			}
		}
		m := &models.Channel{
			Name:       p.Name,
			Visibility: models.ChannelsVisibilityPublic,
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
func (s *channelssrvc) Update(ctx context.Context, p *channels.ChannelUpdatePayload) (res *channels.Channel, err error) {
	err = s.actionWithAuth(ctx, "channels.update", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		if p.Name == "" {
			return channels.MakeInvalidPayload(fmt.Errorf("name is required"))
		} else {
			runes := []rune(p.Name)
			if len(runes) > 255 {
				return channels.MakeInvalidPayload(fmt.Errorf("name is too long"))
			}
		}
		m, err := models.FindChannel(ctx, db, p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return channels.MakeNotFound(fmt.Errorf("channel not found"))
			}
			return err
		}
		m.Name = p.Name
		if _, err := m.Update(ctx, db, boil.Infer()); err != nil {
			return err
		}
		res = s.ModelToResult(m)
		return nil
	})
	return
}

// Delete implements delete.
func (s *channelssrvc) Delete(ctx context.Context, p *channels.DeletePayload) (res *channels.Channel, err error) {
	err = s.actionWithAuth(ctx, "channels.delete", p.SessionID, func(ctx context.Context, db *sql.DB, user *models.User) error {
		m, err := models.FindChannel(ctx, db, p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return channels.MakeNotFound(fmt.Errorf("channel not found"))
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

type ChannelsConvertor struct{}

func NewChannelsConvertor() *ChannelsConvertor {
	return &ChannelsConvertor{}
}

func (s *ChannelsConvertor) ModelsToList(models []*models.Channel) *channels.ChannelList {
	items := s.ModelsToListItems(models)
	return &channels.ChannelList{
		Items:  items,
		Total:  uint64(len(items)),
		Offset: 0,
	}
}

func (*ChannelsConvertor) ModelsToListItems(models []*models.Channel) channels.ChannelListItemCollection {
	items := make(channels.ChannelListItemCollection, len(models))
	for i, result := range models {
		items[i] = &channels.ChannelListItem{
			ID:        result.ID,
			CreatedAt: result.CreatedAt.Format(time.RFC3339),
			UpdatedAt: result.UpdatedAt.Format(time.RFC3339),
			Name:      result.Name,
		}
	}
	return items
}

func (*ChannelsConvertor) ModelToResult(model *models.Channel) *channels.Channel {
	return &channels.Channel{
		ID:        model.ID,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
		Name:      model.Name,
	}
}
