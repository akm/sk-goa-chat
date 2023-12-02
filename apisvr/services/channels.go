package chatapi

import (
	channels "apisvr/services/gen/channels"
	"context"
	"log"
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
	res = &channels.ChannelList{}
	s.logger.Print("channels.list")
	return
}

// Show implements show.
func (s *channelssrvc) Show(ctx context.Context, p *channels.ShowPayload) (res *channels.Channel, err error) {
	res = &channels.Channel{}
	s.logger.Print("channels.show")
	return
}

// Create implements create.
func (s *channelssrvc) Create(ctx context.Context, p *channels.ChannelCreatePayload) (res *channels.Channel, err error) {
	res = &channels.Channel{}
	s.logger.Print("channels.create")
	return
}

// Update implements update.
func (s *channelssrvc) Update(ctx context.Context, p *channels.ChannelUpdatePayload) (res *channels.Channel, err error) {
	res = &channels.Channel{}
	s.logger.Print("channels.update")
	return
}

// Delete implements delete.
func (s *channelssrvc) Delete(ctx context.Context, p *channels.DeletePayload) (res *channels.Channel, err error) {
	res = &channels.Channel{}
	s.logger.Print("channels.delete")
	return
}
