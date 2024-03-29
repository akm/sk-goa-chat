package main

import (
	channels "apisvr/services/gen/channels"
	chatmessages "apisvr/services/gen/chat_messages"
	channelspb "apisvr/services/gen/grpc/channels/pb"
	channelssvr "apisvr/services/gen/grpc/channels/server"
	chat_messagespb "apisvr/services/gen/grpc/chat_messages/pb"
	chatmessagessvr "apisvr/services/gen/grpc/chat_messages/server"
	notificationspb "apisvr/services/gen/grpc/notifications/pb"
	notificationssvr "apisvr/services/gen/grpc/notifications/server"
	userspb "apisvr/services/gen/grpc/users/pb"
	userssvr "apisvr/services/gen/grpc/users/server"
	log "apisvr/services/gen/log"
	notifications "apisvr/services/gen/notifications"
	users "apisvr/services/gen/users"
	"context"
	"net"
	"net/url"
	"sync"

	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, channelsEndpoints *channels.Endpoints, chatMessagesEndpoints *chatmessages.Endpoints, notificationsEndpoints *notifications.Endpoints, usersEndpoints *users.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = logger
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		channelsServer      *channelssvr.Server
		chatMessagesServer  *chatmessagessvr.Server
		notificationsServer *notificationssvr.Server
		usersServer         *userssvr.Server
	)
	{
		channelsServer = channelssvr.New(channelsEndpoints, nil)
		chatMessagesServer = chatmessagessvr.New(chatMessagesEndpoints, nil)
		notificationsServer = notificationssvr.New(notificationsEndpoints, nil)
		usersServer = userssvr.New(usersEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
		grpc.ChainStreamInterceptor(
			grpcmdlwr.StreamRequestID(),
			grpcmdlwr.StreamServerLog(adapter),
		),
	)

	// Register the servers.
	channelspb.RegisterChannelsServer(srv, channelsServer)
	chat_messagespb.RegisterChatMessagesServer(srv, chatMessagesServer)
	notificationspb.RegisterNotificationsServer(srv, notificationsServer)
	userspb.RegisterUsersServer(srv, usersServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Info().Msgf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Info().Msgf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Info().Msgf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
