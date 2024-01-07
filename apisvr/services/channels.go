package chatapi

import (
	"apisvr/lib/sql"
	"apisvr/models"
	channels "apisvr/services/gen/channels"
	log "apisvr/services/gen/log"
	"context"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"goa.design/goa/v3/security"
)

// channels service example implementation.
// The example methods log the requests and return zero values.
type channelssrvc struct {
	logger *log.Logger
	*ChannelsConvertor
}

// NewChannels returns the channels service implementation.
func NewChannels(logger *log.Logger) channels.Service {
	return &channelssrvc{logger: logger, ChannelsConvertor: NewChannelsConvertor()}
}

func (s *channelssrvc) sqlOpen() (*sql.DB, error) {
	return sql.Open(s.logger.Logger)
}

// APIKeyAuth implements the authorization logic for service "channels" for the
// "api_key" security scheme.
func (s *channelssrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// List implements list.
func (s *channelssrvc) List(ctx context.Context, p *channels.ListPayload) (res *channels.ChannelList, err error) {
	s.logger.Info().Msg("channels.list")
	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := models.Channels().All(ctx, db)
	if err != nil {
		return nil, err
	}

	res = s.ModelsToList(results)
	return
}

// Show implements show.
func (s *channelssrvc) Show(ctx context.Context, p *channels.ShowPayload) (res *channels.Channel, err error) {
	s.logger.Info().Msg("channels.show")
	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m, err := models.FindChannel(ctx, db, p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, channels.MakeNotFound(fmt.Errorf("channel not found"))
		}
		return nil, err
	}

	res = s.ModelToResult(m)
	return
}

// Create implements create.
func (s *channelssrvc) Create(ctx context.Context, p *channels.ChannelCreatePayload) (res *channels.Channel, err error) {
	s.logger.Info().Msg("channels.create")

	if p.Name == "" {
		return nil, channels.MakeInvalidPayload(fmt.Errorf("name is required"))
	} else {
		runes := []rune(p.Name)
		if len(runes) > 255 {
			return nil, channels.MakeInvalidPayload(fmt.Errorf("name is too long"))
		}
	}

	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
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

	res = s.ModelToResult(m)
	return
}

// Update implements update.
func (s *channelssrvc) Update(ctx context.Context, p *channels.ChannelUpdatePayload) (res *channels.Channel, err error) {
	s.logger.Info().Msg("channels.update")

	if p.Name == "" {
		return nil, channels.MakeInvalidPayload(fmt.Errorf("name is required"))
	} else {
		runes := []rune(p.Name)
		if len(runes) > 255 {
			return nil, channels.MakeInvalidPayload(fmt.Errorf("name is too long"))
		}
	}

	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m, err := models.FindChannel(ctx, db, p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, channels.MakeNotFound(fmt.Errorf("channel not found"))
		}
		return nil, err
	}

	m.Name = p.Name
	if _, err := m.Update(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	res = s.ModelToResult(m)
	return
}

// Delete implements delete.
func (s *channelssrvc) Delete(ctx context.Context, p *channels.DeletePayload) (res *channels.Channel, err error) {
	s.logger.Info().Msg("channels.delete")
	ctx = SetupContext(ctx)
	db, err := s.sqlOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m, err := models.FindChannel(ctx, db, p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, channels.MakeNotFound(fmt.Errorf("channel not found"))
		}
		return nil, err
	}

	if _, err := m.Delete(ctx, db); err != nil {
		return nil, err
	}

	res = s.ModelToResult(m)
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
