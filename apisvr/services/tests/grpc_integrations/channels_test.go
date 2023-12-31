package grpcintegrations

import (
	"apisvr/lib/time"
	"apisvr/models"
	chatapi "apisvr/services"
	"apisvr/services/gen/channels"
	channelspb "apisvr/services/gen/grpc/channels/pb"
	channelssvr "apisvr/services/gen/grpc/channels/server"
	"apisvr/testlib/testjson"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"apisvr/testlib/testsqlboiler"
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestChannels(t *testing.T) {
	ctx := context.Background()
	logger := testlog.New(t)

	client, closer := setupChannelsServer(ctx, logger)
	defer closer()

	conn := testsql.Setup(t)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	conv := NewChannelsConvertor()

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			out, err := client.List(ctx, &channelspb.ListRequest{})
			assert.NoError(t, err)
			assert.Equal(t, &channelspb.ListResponse{
				Total:  uint64(0),
				Offset: uint64(0),
				Items:  &channelspb.ChannelListItemCollection{},
			}, testjson.Reassign(t, out))
		})
	})

	ch1 := &models.Channel{Name: "general", Visibility: models.ChannelsVisibilityPublic}
	ch2 := &models.Channel{Name: "random", Visibility: models.ChannelsVisibilityPublic}
	testsqlboiler.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
	assert.Equal(t, now, ch1.CreatedAt)

	t.Run("list", func(t *testing.T) {
		out, err := client.List(ctx, &channelspb.ListRequest{})
		assert.NoError(t, err)
		assert.Equal(t, conv.ModelsToListResponse([]*models.Channel{ch1, ch2}), testjson.Reassign(t, out))
	})
}

func setupChannelsServer(ctx context.Context, logger *log.Logger) (channelspb.ChannelsClient, func()) {
	buffer := 101024 * 1024
	listener := bufconn.Listen(buffer)

	adapter := middleware.NewLogger(logger)
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)
	channelsSvc := chatapi.NewChannels(logger)
	channelsEndpoints := channels.NewEndpoints(channelsSvc)
	channelsServer := channelssvr.New(channelsEndpoints, nil)
	channelspb.RegisterChannelsServer(srv, channelsServer)

	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	gRpcDialOptions := []grpc.DialOption{
		grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) { return listener.Dial() },
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, "", gRpcDialOptions...)
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := listener.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		srv.Stop()
	}

	client := channelspb.NewChannelsClient(conn)

	return client, closer
}

type channelsConvertor struct{}

func NewChannelsConvertor() *channelsConvertor {
	return &channelsConvertor{}
}

func (c *channelsConvertor) ModelsToListResponse(models []*models.Channel) *channelspb.ListResponse {
	items := make([]*channelspb.ChannelListItem, len(models))
	for i, m := range models {
		items[i] = c.ModelToListItem(m)
	}
	return &channelspb.ListResponse{
		Total:  uint64(len(models)),
		Offset: uint64(0),
		Items: &channelspb.ChannelListItemCollection{
			Field: items,
		},
	}
}

func (c *channelsConvertor) ModelToListItem(m *models.Channel) *channelspb.ChannelListItem {
	return &channelspb.ChannelListItem{
		Id:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
	}
}
