// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications endpoints
//
// Command:
// $ goa gen apisvr/design -o ./services

package notifications

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "notifications" service endpoints.
type Endpoints struct {
	Subscribe goa.Endpoint
}

// SubscribeEndpointInput holds both the payload and the server stream of the
// "subscribe" method.
type SubscribeEndpointInput struct {
	// Payload is the method payload.
	Payload *SubscribePayload
	// Stream is the server stream used by the "subscribe" method to send data.
	Stream SubscribeServerStream
}

// NewEndpoints wraps the methods of the "notifications" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Subscribe: NewSubscribeEndpoint(s, a.APIKeyAuth),
	}
}

// Use applies the given middleware to all the "notifications" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Subscribe = m(e.Subscribe)
}

// NewSubscribeEndpoint returns an endpoint function that calls the method
// "subscribe" of service "notifications".
func NewSubscribeEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		ep := req.(*SubscribeEndpointInput)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, ep.Payload.IDToken, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Subscribe(ctx, ep.Payload, ep.Stream)
	}
}
