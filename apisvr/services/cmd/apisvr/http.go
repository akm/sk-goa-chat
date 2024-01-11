package main

import (
	channels "apisvr/services/gen/channels"
	chatmessages "apisvr/services/gen/chat_messages"
	channelssvr "apisvr/services/gen/http/channels/server"
	chatmessagessvr "apisvr/services/gen/http/chat_messages/server"
	notificationssvr "apisvr/services/gen/http/notifications/server"
	sessionssvr "apisvr/services/gen/http/sessions/server"
	userssvr "apisvr/services/gen/http/users/server"
	log "apisvr/services/gen/log"
	notifications "apisvr/services/gen/notifications"
	sessions "apisvr/services/gen/sessions"
	users "apisvr/services/gen/users"
	"context"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, channelsEndpoints *channels.Endpoints, chatMessagesEndpoints *chatmessages.Endpoints, notificationsEndpoints *notifications.Endpoints, sessionsEndpoints *sessions.Endpoints, usersEndpoints *users.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = logger
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		channelsServer      *channelssvr.Server
		chatMessagesServer  *chatmessagessvr.Server
		notificationsServer *notificationssvr.Server
		sessionsServer      *sessionssvr.Server
		usersServer         *userssvr.Server
	)
	{
		eh := errorHandler(logger)
		upgrader := &websocket.Upgrader{}
		channelsServer = channelssvr.New(channelsEndpoints, mux, dec, enc, eh, nil)
		chatMessagesServer = chatmessagessvr.New(chatMessagesEndpoints, mux, dec, enc, eh, nil)
		notificationsServer = notificationssvr.New(notificationsEndpoints, mux, dec, enc, eh, nil, upgrader, nil)
		sessionsServer = sessionssvr.New(sessionsEndpoints, mux, dec, enc, eh, nil)
		usersServer = userssvr.New(usersEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				channelsServer,
				chatMessagesServer,
				notificationsServer,
				sessionsServer,
				usersServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	channelssvr.Mount(mux, channelsServer)
	chatmessagessvr.Mount(mux, chatMessagesServer)
	notificationssvr.Mount(mux, notificationsServer)
	sessionssvr.Mount(mux, sessionsServer)
	userssvr.Mount(mux, usersServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range channelsServer.Mounts {
		logger.Info().Msgf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range chatMessagesServer.Mounts {
		logger.Info().Msgf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range notificationsServer.Mounts {
		logger.Info().Msgf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range sessionsServer.Mounts {
		logger.Info().Msgf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range usersServer.Mounts {
		logger.Info().Msgf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Info().Msgf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Info().Msgf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Info().Msgf("failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Error().Str("id", id).Err(err).Send()
	}
}
