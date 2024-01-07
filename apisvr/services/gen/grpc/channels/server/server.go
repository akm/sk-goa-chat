// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels gRPC server
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	channels "apisvr/services/gen/channels"
	channelspb "apisvr/services/gen/grpc/channels/pb"
	"context"
	"errors"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the channelspb.ChannelsServer interface.
type Server struct {
	ListH   goagrpc.UnaryHandler
	ShowH   goagrpc.UnaryHandler
	CreateH goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
	DeleteH goagrpc.UnaryHandler
	channelspb.UnimplementedChannelsServer
}

// New instantiates the server struct with the channels service endpoints.
func New(e *channels.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH:   NewListHandler(e.List, uh),
		ShowH:   NewShowHandler(e.Show, uh),
		CreateH: NewCreateHandler(e.Create, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
		DeleteH: NewDeleteHandler(e.Delete, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "channels" service
// "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeListRequest, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in channelspb.ChannelsServer interface.
func (s *Server) List(ctx context.Context, message *channelspb.ListRequest) (*channelspb.ListResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "channels")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "unauthenticated":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*channelspb.ListResponse), nil
}

// NewShowHandler creates a gRPC handler which serves the "channels" service
// "show" endpoint.
func NewShowHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowRequest, EncodeShowResponse)
	}
	return h
}

// Show implements the "Show" method in channelspb.ChannelsServer interface.
func (s *Server) Show(ctx context.Context, message *channelspb.ShowRequest) (*channelspb.ShowResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "show")
	ctx = context.WithValue(ctx, goa.ServiceKey, "channels")
	resp, err := s.ShowH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			case "unauthenticated":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*channelspb.ShowResponse), nil
}

// NewCreateHandler creates a gRPC handler which serves the "channels" service
// "create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in channelspb.ChannelsServer interface.
func (s *Server) Create(ctx context.Context, message *channelspb.CreateRequest) (*channelspb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "channels")
	resp, err := s.CreateH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "invalid_payload":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "unauthenticated":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*channelspb.CreateResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "channels" service
// "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in channelspb.ChannelsServer interface.
func (s *Server) Update(ctx context.Context, message *channelspb.UpdateRequest) (*channelspb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "channels")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			case "invalid_payload":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "unauthenticated":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*channelspb.UpdateResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "channels" service
// "delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in channelspb.ChannelsServer interface.
func (s *Server) Delete(ctx context.Context, message *channelspb.DeleteRequest) (*channelspb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "channels")
	resp, err := s.DeleteH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			case "unauthenticated":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*channelspb.DeleteResponse), nil
}
