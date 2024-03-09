// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels endpoints
//
// Command:
// $ goa gen apisvr/design -o ./services

package channels

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "channels" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Show   goa.Endpoint
	Create goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "channels" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:   NewListEndpoint(s, a.APIKeyAuth),
		Show:   NewShowEndpoint(s, a.APIKeyAuth),
		Create: NewCreateEndpoint(s, a.APIKeyAuth),
		Update: NewUpdateEndpoint(s, a.APIKeyAuth),
		Delete: NewDeleteEndpoint(s, a.APIKeyAuth),
	}
}

// Use applies the given middleware to all the "channels" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "channels".
func NewListEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.IDToken, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedChannelList(res, "default")
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "channels".
func NewShowEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ShowPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.IDToken, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedChannel(res, "default")
		return vres, nil
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "channels".
func NewCreateEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ChannelCreatePayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.IDToken, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedChannel(res, "default")
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "channels".
func NewUpdateEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ChannelUpdatePayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.IDToken, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Update(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedChannel(res, "default")
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "channels".
func NewDeleteEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "api_key",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.IDToken, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Delete(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedChannel(res, "default")
		return vres, nil
	}
}