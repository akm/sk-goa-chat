package main

import (
	chatapi "apisvr/services"
	channels "apisvr/services/gen/channels"
	chatmessages "apisvr/services/gen/chat_messages"
	log "apisvr/services/gen/log"
	notifications "apisvr/services/gen/notifications"
	users "apisvr/services/gen/users"
	goaendpoints "applib/goa/endpoints"
	"applib/goa/goaext"
	"applib/goa/goasql"
	"context"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		grpcPortF = flag.String("grpc-port", "", "gRPC port (overrides host gRPC port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New("chatapi", dbgF != nil && *dbgF)
	}

	// Initialize the services.
	var (
		channelsSvc      channels.Service
		chatMessagesSvc  chatmessages.Service
		notificationsSvc notifications.Service
		usersSvc         users.Service
	)
	{
		channelsSvc = chatapi.NewChannels(logger)
		chatMessagesSvc = chatapi.NewChatMessages(logger)
		notificationsSvc = chatapi.NewNotifications(logger)
		usersSvc = chatapi.NewUsers(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		channelsEndpoints      *channels.Endpoints
		chatMessagesEndpoints  *chatmessages.Endpoints
		notificationsEndpoints *notifications.Endpoints
		usersEndpoints         *users.Endpoints
	)
	{
		wappers := goaendpoints.Wrappers{
			goaendpoints.ErrorHandler(goaext.DefaultErrorHandler(logger)),
			goasql.ConnectionEndpointWrapper(logger.Logger),
		}
		channelsEndpoints = goaendpoints.Wrap[*channels.Endpoints](channels.NewEndpoints(channelsSvc), wappers.Wrap)
		chatMessagesEndpoints = goaendpoints.Wrap[*chatmessages.Endpoints](chatmessages.NewEndpoints(chatMessagesSvc), wappers.Wrap)
		notificationsEndpoints = goaendpoints.Wrap[*notifications.Endpoints](notifications.NewEndpoints(notificationsSvc), wappers.Wrap)
		usersEndpoints = goaendpoints.Wrap[*users.Endpoints](users.NewEndpoints(usersSvc), wappers.Wrap)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8000"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatal().Msgf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatal().Msgf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, channelsEndpoints, chatMessagesEndpoints, notificationsEndpoints, usersEndpoints, &wg, errc, logger, *dbgF)
		}

		{
			addr := "grpc://localhost:8080"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatal().Msgf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatal().Msgf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *grpcPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "8080")
			}
			handleGRPCServer(ctx, u, channelsEndpoints, chatMessagesEndpoints, notificationsEndpoints, usersEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Fatal().Msgf("invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Info().Msgf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info().Msg("exited")
}
