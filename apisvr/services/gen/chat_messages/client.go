// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages client
//
// Command:
// $ goa gen apisvr/design -o ./services

package chatmessages

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "chat_messages" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "chat_messages" service client given the endpoints.
func NewClient(list, show, create, update, delete_ goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		CreateEndpoint: create,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
	}
}

// List calls the "list" endpoint of the "chat_messages" service.
// List may return the following errors:
//   - "unauthenticated" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) List(ctx context.Context, p *ListPayload) (res *ChatMessageList, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatMessageList), nil
}

// Show calls the "show" endpoint of the "chat_messages" service.
// Show may return the following errors:
//   - "not_found" (type *goa.ServiceError)
//   - "unauthenticated" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *ChatMessage, err error) {
	var ires any
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatMessage), nil
}

// Create calls the "create" endpoint of the "chat_messages" service.
// Create may return the following errors:
//   - "invalid_payload" (type *goa.ServiceError)
//   - "unauthenticated" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Create(ctx context.Context, p *ChatMessageCreatePayload) (res *ChatMessage, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatMessage), nil
}

// Update calls the "update" endpoint of the "chat_messages" service.
// Update may return the following errors:
//   - "not_found" (type *goa.ServiceError)
//   - "invalid_payload" (type *goa.ServiceError)
//   - "unauthenticated" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Update(ctx context.Context, p *ChatMessageUpdatePayload) (res *ChatMessage, err error) {
	var ires any
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatMessage), nil
}

// Delete calls the "delete" endpoint of the "chat_messages" service.
// Delete may return the following errors:
//   - "not_found" (type *goa.ServiceError)
//   - "unauthenticated" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (res *ChatMessage, err error) {
	var ires any
	ires, err = c.DeleteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatMessage), nil
}
