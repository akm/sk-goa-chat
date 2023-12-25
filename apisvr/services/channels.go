package chatapi

import (
	"apisvr/lib/sql"
	"apisvr/models"
	channels "apisvr/services/gen/channels"
	"context"
	"log"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// channels service example implementation.
// The example methods log the requests and return zero values.
type channelssrvc struct {
	logger *log.Logger
}

// NewChannels returns the channels service implementation.
func NewChannels(logger *log.Logger) channels.Service {
	return &channelssrvc{logger}
}

// List implements list.
func (s *channelssrvc) List(ctx context.Context) (res *channels.ChannelList, err error) {
	s.logger.Print("channels.list")
	ctx = SetupContext(ctx)
	db, err := sql.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := models.Channels().All(ctx, db)
	if err != nil {
		return nil, err
	}

	items := make(channels.ChannelListItemCollection, len(results))
	for i, result := range results {
		items[i] = &channels.ChannelListItem{
			ID:        result.ID,
			CreatedAt: result.CreatedAt.Format(time.RFC3339),
			UpdatedAt: result.UpdatedAt.Format(time.RFC3339),
			Name:      result.Name,
		}
	}
	res = &channels.ChannelList{
		Items:  items,
		Total:  uint64(len(items)),
		Offset: 0,
	}
	return
}

func (s *channelssrvc) ConvertModelsToListResult(models []*models.Channel) *channels.ChannelList {
	items := s.ConvertModelsToListItems(models)
	return &channels.ChannelList{
		Items:  items,
		Total:  uint64(len(items)),
		Offset: 0,
	}
}

func (*channelssrvc) ConvertModelsToListItems(models []*models.Channel) channels.ChannelListItemCollection {
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

// Show implements show.
func (s *channelssrvc) Show(ctx context.Context, p *channels.ShowPayload) (res *channels.Channel, err error) {
	s.logger.Print("channels.show")
	ctx = SetupContext(ctx)
	db, err := sql.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m, err := models.FindChannel(ctx, db, p.ID)
	if err != nil {
		return nil, err
	}

	res = s.ConvertModelToResult(m)
	return
}

// Create implements create.
func (s *channelssrvc) Create(ctx context.Context, p *channels.ChannelCreatePayload) (res *channels.Channel, err error) {
	s.logger.Print("channels.create")
	ctx = SetupContext(ctx)
	db, err := sql.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m := &models.Channel{
		Name:       p.Name,
		Visibility: models.ChannelsVisibilityPublic,
	}
	if err := m.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	res = s.ConvertModelToResult(m)
	return
}

// Update implements update.
func (s *channelssrvc) Update(ctx context.Context, p *channels.ChannelUpdatePayload) (res *channels.Channel, err error) {
	s.logger.Print("channels.update")
	ctx = SetupContext(ctx)
	db, err := sql.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m, err := models.FindChannel(ctx, db, p.ID)
	if err != nil {
		return nil, err
	}

	m.Name = p.Name
	if _, err := m.Update(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	res = s.ConvertModelToResult(m)
	return
}

// Delete implements delete.
func (s *channelssrvc) Delete(ctx context.Context, p *channels.DeletePayload) (res *channels.Channel, err error) {
	s.logger.Print("channels.delete")
	ctx = SetupContext(ctx)
	db, err := sql.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m, err := models.FindChannel(ctx, db, p.ID)
	if err != nil {
		return nil, err
	}

	if _, err := m.Delete(ctx, db); err != nil {
		return nil, err
	}

	res = s.ConvertModelToResult(m)
	return
}

func (*channelssrvc) ConvertModelToResult(model *models.Channel) *channels.Channel {
	return &channels.Channel{
		ID:        model.ID,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
		Name:      model.Name,
	}
}
