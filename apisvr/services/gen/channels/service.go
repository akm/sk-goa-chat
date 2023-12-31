// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels service
//
// Command:
// $ goa gen apisvr/design -o ./services

package channels

import (
	channelsviews "apisvr/services/gen/channels/views"
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the channels service interface.
type Service interface {
	// List implements list.
	List(context.Context) (res *ChannelList, err error)
	// Show implements show.
	Show(context.Context, *ShowPayload) (res *Channel, err error)
	// Create implements create.
	Create(context.Context, *ChannelCreatePayload) (res *Channel, err error)
	// Update implements update.
	Update(context.Context, *ChannelUpdatePayload) (res *Channel, err error)
	// Delete implements delete.
	Delete(context.Context, *DeletePayload) (res *Channel, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "channels"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "show", "create", "update", "delete"}

// Channel is the result type of the channels service show method.
type Channel struct {
	// ID
	ID uint64
	// CreatedAt
	CreatedAt string
	// UpdatedAt
	UpdatedAt string
	// Name
	Name string
}

// ChannelCreatePayload is the payload type of the channels service create
// method.
type ChannelCreatePayload struct {
	// Name
	Name string
}

// ChannelList is the result type of the channels service list method.
type ChannelList struct {
	// Items
	Items ChannelListItemCollection
	// Total number of items
	Total uint64
	// Offset
	Offset uint64
}

type ChannelListItem struct {
	// ID
	ID uint64
	// CreatedAt
	CreatedAt string
	// UpdatedAt
	UpdatedAt string
	// Name
	Name string
}

type ChannelListItemCollection []*ChannelListItem

// ChannelUpdatePayload is the payload type of the channels service update
// method.
type ChannelUpdatePayload struct {
	// ID
	ID uint64
	// Name
	Name string
}

// DeletePayload is the payload type of the channels service delete method.
type DeletePayload struct {
	// ID
	ID uint64
}

// ShowPayload is the payload type of the channels service show method.
type ShowPayload struct {
	// ID
	ID uint64
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_found", false, false, false)
}

// MakeInvalidPayload builds a goa.ServiceError from an error.
func MakeInvalidPayload(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_payload", false, false, false)
}

// NewChannelList initializes result type ChannelList from viewed result type
// ChannelList.
func NewChannelList(vres *channelsviews.ChannelList) *ChannelList {
	return newChannelList(vres.Projected)
}

// NewViewedChannelList initializes viewed result type ChannelList from result
// type ChannelList using the given view.
func NewViewedChannelList(res *ChannelList, view string) *channelsviews.ChannelList {
	p := newChannelListView(res)
	return &channelsviews.ChannelList{Projected: p, View: "default"}
}

// NewChannel initializes result type Channel from viewed result type Channel.
func NewChannel(vres *channelsviews.Channel) *Channel {
	return newChannel(vres.Projected)
}

// NewViewedChannel initializes viewed result type Channel from result type
// Channel using the given view.
func NewViewedChannel(res *Channel, view string) *channelsviews.Channel {
	p := newChannelView(res)
	return &channelsviews.Channel{Projected: p, View: "default"}
}

// newChannelList converts projected type ChannelList to service type
// ChannelList.
func newChannelList(vres *channelsviews.ChannelListView) *ChannelList {
	res := &ChannelList{}
	if vres.Total != nil {
		res.Total = *vres.Total
	}
	if vres.Offset != nil {
		res.Offset = *vres.Offset
	}
	if vres.Items != nil {
		res.Items = newChannelListItemCollection(vres.Items)
	}
	return res
}

// newChannelListView projects result type ChannelList to projected type
// ChannelListView using the "default" view.
func newChannelListView(res *ChannelList) *channelsviews.ChannelListView {
	vres := &channelsviews.ChannelListView{
		Total:  &res.Total,
		Offset: &res.Offset,
	}
	if res.Items != nil {
		vres.Items = newChannelListItemCollectionView(res.Items)
	}
	return vres
}

// newChannelListItemCollection converts projected type
// ChannelListItemCollection to service type ChannelListItemCollection.
func newChannelListItemCollection(vres channelsviews.ChannelListItemCollectionView) ChannelListItemCollection {
	res := make(ChannelListItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newChannelListItem(n)
	}
	return res
}

// newChannelListItemCollectionView projects result type
// ChannelListItemCollection to projected type ChannelListItemCollectionView
// using the "default" view.
func newChannelListItemCollectionView(res ChannelListItemCollection) channelsviews.ChannelListItemCollectionView {
	vres := make(channelsviews.ChannelListItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newChannelListItemView(n)
	}
	return vres
}

// newChannelListItem converts projected type ChannelListItem to service type
// ChannelListItem.
func newChannelListItem(vres *channelsviews.ChannelListItemView) *ChannelListItem {
	res := &ChannelListItem{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newChannelListItemView projects result type ChannelListItem to projected
// type ChannelListItemView using the "default" view.
func newChannelListItemView(res *ChannelListItem) *channelsviews.ChannelListItemView {
	vres := &channelsviews.ChannelListItemView{
		ID:        &res.ID,
		CreatedAt: &res.CreatedAt,
		UpdatedAt: &res.UpdatedAt,
		Name:      &res.Name,
	}
	return vres
}

// newChannel converts projected type Channel to service type Channel.
func newChannel(vres *channelsviews.ChannelView) *Channel {
	res := &Channel{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newChannelView projects result type Channel to projected type ChannelView
// using the "default" view.
func newChannelView(res *Channel) *channelsviews.ChannelView {
	vres := &channelsviews.ChannelView{
		ID:        &res.ID,
		CreatedAt: &res.CreatedAt,
		UpdatedAt: &res.UpdatedAt,
		Name:      &res.Name,
	}
	return vres
}
