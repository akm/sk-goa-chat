package grpcintegrations

import (
	"apisvr/lib/firebase"
	"apisvr/lib/firebase/auth"
	"apisvr/lib/time"
	"apisvr/models"
	chatapi "apisvr/services"
	"apisvr/services/gen/channels"
	channelspb "apisvr/services/gen/grpc/channels/pb"
	channelssvr "apisvr/services/gen/grpc/channels/server"
	"apisvr/services/gen/log"
	"apisvr/testlib/testfirebase/testauth"
	"apisvr/testlib/testjson"
	"apisvr/testlib/testlog"
	"apisvr/testlib/testsql"
	"apisvr/testlib/testsqlboiler"
	"context"
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func TestChannels(t *testing.T) {
	ctx := context.Background()
	logger := testlog.New(t)

	client, closer := setupChannelsServer(ctx, &log.Logger{Logger: logger})
	defer closer()

	conn := testsql.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer time.SetTime(now)

	conv := NewChannelsConvertor()

	fbapp, err := firebase.NewApp(ctx, nil)
	require.NoError(t, err)
	fbauth, err := auth.NewClientRaw(ctx, fbapp)
	require.NoError(t, err)

	t.Run("delete all of users before test", func(t *testing.T) {
		testauth.DeleteUsers(t, ctx, fbauth)
	})

	fooEmail := "foo@example.com"
	fooName := "Foo"
	fooPassword := "Passw0rd!"

	t.Run("create foo", func(t *testing.T) {
		args := &auth.UserToCreate{}
		args.Email(fooEmail)
		args.DisplayName(fooName)
		args.Password(fooPassword)
		res, err := fbauth.CreateUser(ctx, args)
		require.NoError(t, err)
		// t.Logf("result: %+v", res)
		require.NotEmpty(t, res.UID)

		t.Run("insert user to db", func(t *testing.T) {
			user := &models.User{FbauthUID: res.UID, Email: fooEmail, Name: fooName}
			testsqlboiler.Insert(t, ctx, conn, boil.Infer(), user)
		})
	})

	sessionID := testauth.GetSessionCookie(t, ctx, fbauth, fooEmail, fooPassword)

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			out, err := client.List(ctx, &channelspb.ListRequest{SessionId: sessionID})
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
		out, err := client.List(ctx, &channelspb.ListRequest{SessionId: sessionID})
		assert.NoError(t, err)
		assert.Equal(t, conv.ModelsToListResponse([]*models.Channel{ch1, ch2}), testjson.Reassign(t, out))
	})

	t.Run("show", func(t *testing.T) {
		for _, ch := range []*models.Channel{ch1, ch2} {
			t.Run(ch.Name, func(t *testing.T) {
				out, err := client.Show(ctx, &channelspb.ShowRequest{SessionId: sessionID, Id: ch.ID})
				assert.NoError(t, err)
				assert.Equal(t, conv.ModelToShowResponse(ch), testjson.Reassign(t, out))
			})
		}
		t.Run("not found", func(t *testing.T) {
			out, err := client.Show(ctx, &channelspb.ShowRequest{SessionId: sessionID, Id: 999})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.NotFound, "channel not found").Error(), err.Error())
		})
	})

	t.Run("create", func(t *testing.T) {
		t.Run("valid name", func(t *testing.T) {
			name := "test1"
			out, err := client.Create(ctx, &channelspb.CreateRequest{SessionId: sessionID, Name: name})
			assert.NoError(t, err)
			require.NotNil(t, out)
			ch := &models.Channel{ID: out.Id, Name: name, CreatedAt: now, UpdatedAt: now}
			assert.Equal(t, conv.ModelToCreateResponse(t, ch), testjson.Reassign(t, out))
		})
		t.Run("empty name", func(t *testing.T) {
			out, err := client.Create(ctx, &channelspb.CreateRequest{SessionId: sessionID, Name: ""})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.InvalidArgument, "name is required").Error(), err.Error())
		})
		t.Run("too long name", func(t *testing.T) {
			out, err := client.Create(ctx, &channelspb.CreateRequest{SessionId: sessionID, Name: strings.Repeat("a", 256)})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.InvalidArgument, "name is too long").Error(), err.Error())
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			out, err := client.Update(ctx, &channelspb.UpdateRequest{SessionId: sessionID, Id: 999, Name: "test"})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.NotFound, "channel not found").Error(), err.Error())
		})
		t.Run("valid name", func(t *testing.T) {
			newName := ch1.Name + "-dash"
			out, err := client.Update(ctx, &channelspb.UpdateRequest{SessionId: sessionID, Id: ch1.ID, Name: newName})
			assert.NoError(t, err)
			require.NotNil(t, out)
			ch := &models.Channel{ID: ch1.ID, Name: newName, CreatedAt: now, UpdatedAt: now}
			assert.Equal(t, conv.ModelToUpdateResponse(t, ch), testjson.Reassign(t, out))

		})
		t.Run("empty name", func(t *testing.T) {
			out, err := client.Update(ctx, &channelspb.UpdateRequest{SessionId: sessionID, Id: ch1.ID, Name: ""})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.InvalidArgument, "name is required").Error(), err.Error())
		})
		t.Run("too long name", func(t *testing.T) {
			out, err := client.Update(ctx, &channelspb.UpdateRequest{SessionId: sessionID, Id: ch1.ID, Name: strings.Repeat("a", 256)})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.InvalidArgument, "name is too long").Error(), err.Error())
		})
	})

	t.Run("delete", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			out, err := client.Delete(ctx, &channelspb.DeleteRequest{SessionId: sessionID, Id: 999})
			assert.Nil(t, out)
			assert.Error(t, err)
			assert.Equal(t, status.Error(codes.NotFound, "channel not found").Error(), err.Error())
		})
		t.Run("valid id", func(t *testing.T) {
			out, err := client.Delete(ctx, &channelspb.DeleteRequest{SessionId: sessionID, Id: ch2.ID})
			assert.NoError(t, err)
			require.NotNil(t, out)
			assert.Equal(t, conv.ModelToDeleteResponse(t, ch2), testjson.Reassign(t, out))
		})
	})
}

func setupChannelsServer(ctx context.Context, logger *log.Logger) (channelspb.ChannelsClient, func()) {
	buffer := 101024 * 1024
	listener := bufconn.Listen(buffer)

	adapter := logger // middleware.NewLogger(logger)
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
			logger.Error().Msgf("error serving server: %v", err)
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
		logger.Error().Msgf("error connecting to server: %v", err)
	}

	closer := func() {
		err := listener.Close()
		if err != nil {
			logger.Error().Msgf("error closing listener: %v", err)
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

func (c *channelsConvertor) ModelToShowResponse(m *models.Channel) *channelspb.ShowResponse {
	return &channelspb.ShowResponse{
		Id:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
	}
}
func (c *channelsConvertor) ModelToCreateResponse(t *testing.T, m *models.Channel) *channelspb.CreateResponse {
	return testjson.ReassignAs[channelspb.ShowResponse, channelspb.CreateResponse](t, c.ModelToShowResponse(m))
}
func (c *channelsConvertor) ModelToUpdateResponse(t *testing.T, m *models.Channel) *channelspb.UpdateResponse {
	return testjson.ReassignAs[channelspb.ShowResponse, channelspb.UpdateResponse](t, c.ModelToShowResponse(m))
}
func (c *channelsConvertor) ModelToDeleteResponse(t *testing.T, m *models.Channel) *channelspb.DeleteResponse {
	return testjson.ReassignAs[channelspb.ShowResponse, channelspb.DeleteResponse](t, c.ModelToShowResponse(m))
}
