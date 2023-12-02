// Code generated by goa v3.14.0, DO NOT EDIT.
//
// channels HTTP client types
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	channels "apisvr/services/gen/channels"
	channelsviews "apisvr/services/gen/channels/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "channels" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateRequestBody is the type of the "channels" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// ListResponseBody is the type of the "channels" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Items
	Items ChannelListItemCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// Total number of items
	Total *uint64 `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// Offset
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty" xml:"offset,omitempty"`
}

// ShowResponseBody is the type of the "channels" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CreateResponseBody is the type of the "channels" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateResponseBody is the type of the "channels" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// DeleteResponseBody is the type of the "channels" service "delete" endpoint
// HTTP response body.
type DeleteResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ChannelListItemCollectionResponseBody is used to define fields on response
// body types.
type ChannelListItemCollectionResponseBody []*ChannelListItemResponseBody

// ChannelListItemResponseBody is used to define fields on response body types.
type ChannelListItemResponseBody struct {
	// ID
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// CreatedAt
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "channels" service.
func NewCreateRequestBody(p *channels.ChannelCreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Name: p.Name,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "channels" service.
func NewUpdateRequestBody(p *channels.ChannelUpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name: p.Name,
	}
	return body
}

// NewListChannelListOK builds a "channels" service "list" endpoint result from
// a HTTP "OK" response.
func NewListChannelListOK(body *ListResponseBody) *channelsviews.ChannelListView {
	v := &channelsviews.ChannelListView{
		Total:  body.Total,
		Offset: body.Offset,
	}
	v.Items = make([]*channelsviews.ChannelListItemView, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalChannelListItemResponseBodyToChannelsviewsChannelListItemView(val)
	}

	return v
}

// NewShowChannelOK builds a "channels" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowChannelOK(body *ShowResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewCreateChannelCreated builds a "channels" service "create" endpoint result
// from a HTTP "Created" response.
func NewCreateChannelCreated(body *CreateResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewUpdateChannelOK builds a "channels" service "update" endpoint result from
// a HTTP "OK" response.
func NewUpdateChannelOK(body *UpdateResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// NewDeleteChannelOK builds a "channels" service "delete" endpoint result from
// a HTTP "OK" response.
func NewDeleteChannelOK(body *DeleteResponseBody) *channelsviews.ChannelView {
	v := &channelsviews.ChannelView{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
	}

	return v
}

// ValidateChannelListItemCollectionResponseBody runs the validations defined
// on Channel-List-ItemCollectionResponseBody
func ValidateChannelListItemCollectionResponseBody(body ChannelListItemCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateChannelListItemResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChannelListItemResponseBody runs the validations defined on
// Channel-List-ItemResponseBody
func ValidateChannelListItemResponseBody(body *ChannelListItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created_at", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updated_at", *body.UpdatedAt, goa.FormatDateTime))
	}
	return
}
