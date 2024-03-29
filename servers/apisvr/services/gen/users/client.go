// Code generated by goa v3.14.1, DO NOT EDIT.
//
// users client
//
// Command:
// $ goa gen apisvr/design -o ./services

package users

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "users" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
}

// NewClient initializes a "users" service client given the endpoints.
func NewClient(list, create goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		CreateEndpoint: create,
	}
}

// List calls the "list" endpoint of the "users" service.
func (c *Client) List(ctx context.Context) (res *UserList, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*UserList), nil
}

// Create calls the "create" endpoint of the "users" service.
// Create may return the following errors:
//   - "invalid_payload" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Create(ctx context.Context, p *UserCreatePayload) (res *User, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}
