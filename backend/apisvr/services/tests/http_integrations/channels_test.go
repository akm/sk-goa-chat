package httpintegrations

import (
	chatapi "apisvr/services"
	"apisvr/services/gen/channels"
	"apisvr/services/gen/http/channels/server"
	"apisvr/services/gen/log"
	"applib/database/sql/sqltest"
	"applib/encoding/json/jsontest"
	"applib/firebase/auth/authtest"
	"applib/goa/goatest"
	"applib/log/logtest"
	"applib/sqlboiler/sqlboilertest"
	"applib/time"
	"applib/time/timetest"
	"biz/models"
	"biz/testuser"
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestChannels(t *testing.T) {
	logger := logtest.New(t)
	conn := sqltest.Setup(t, logger)
	defer conn.Close()

	now := time.Now()
	defer timetest.SetNow(now)

	ctx := context.Background()
	srvc := chatapi.NewChannels(&log.Logger{Logger: logger})
	conv := chatapi.NewChannelsConvertor()

	fbauth := authtest.Setup(t, ctx)

	userFoo := testuser.Foo()
	userFoo.Setup(t, ctx, fbauth, conn)
	idToken := userFoo.IDToken

	baseAuthSrvc := chatapi.NewBaseAuthService(&log.Logger{Logger: logger}, nil)
	apiKeyAuth := baseAuthSrvc.APIKeyAuth

	checker := goahttpcheck.New()
	checker.Mount(server.NewListHandler, server.MountListHandler, channels.NewListEndpoint(srvc, apiKeyAuth))
	checker.Mount(server.NewShowHandler, server.MountShowHandler, channels.NewShowEndpoint(srvc, apiKeyAuth))
	checker.Mount(server.NewCreateHandler, server.MountCreateHandler, channels.NewCreateEndpoint(srvc, apiKeyAuth))
	checker.Mount(server.NewUpdateHandler, server.MountUpdateHandler, channels.NewUpdateEndpoint(srvc, apiKeyAuth))
	checker.Mount(server.NewDeleteHandler, server.MountDeleteHandler, channels.NewDeleteEndpoint(srvc, apiKeyAuth))

	t.Run("no data", func(t *testing.T) {
		t.Run("list", func(t *testing.T) {
			checker.Test(t, http.MethodGet, "/api/channels").
				WithHeader("X-ID-TOKEN", idToken).
				Check().HasStatus(http.StatusOK).Cb(func(r *http.Response) {
				defer r.Body.Close()
				res := jsontest.UnmarshalFrom[server.ListResponseBody](t, r.Body)
				assert.Equal(t, &server.ListResponseBody{
					Total:  0,
					Offset: 0,
					Items:  server.ChannelListItemResponseBodyCollection{},
				}, res)
			})
		})
	})

	ch1 := &models.Channel{Name: "general", Visibility: models.ChannelsVisibilityPublic}
	ch2 := &models.Channel{Name: "random", Visibility: models.ChannelsVisibilityPublic}
	sqlboilertest.Insert(t, ctx, conn, boil.Infer(), ch1, ch2)
	assert.Equal(t, now, ch1.CreatedAt)

	t.Run("list", func(t *testing.T) {
		checker.Test(t, http.MethodGet, "/api/channels").
			WithHeader("X-ID-TOKEN", idToken).
			Check().HasStatus(http.StatusOK).Cb(func(r *http.Response) {
			defer r.Body.Close()
			res := jsontest.CamelizeJsonKeysAndUnmarshalFrom[channels.ChannelList](t, r.Body)
			assert.Equal(t, conv.ModelsToList([]*models.Channel{ch1, ch2}), res)
		})
	})

	t.Run("show", func(t *testing.T) {
		for _, ch := range []*models.Channel{ch1, ch2} {
			t.Run(ch.Name, func(t *testing.T) {
				checker.Test(t, http.MethodGet, fmt.Sprintf("/api/channels/%d", ch.ID)).
					WithHeader("X-ID-TOKEN", idToken).
					Check().HasStatus(http.StatusOK).Cb(func(r *http.Response) {
					defer r.Body.Close()
					res := jsontest.UnmarshalFrom[server.ShowResponseBody](t, r.Body)
					expected := jsontest.Unmarshal[server.ShowResponseBody](t, jsontest.MarshalAndSnakeizeJsonKeys(t, conv.ModelToResult(ch)))
					assert.Equal(t, expected, res)
				})
			})
		}
		t.Run("not found", func(t *testing.T) {
			checker.Test(t, http.MethodGet, fmt.Sprintf("/api/channels/%d", 999)).
				WithHeader("X-ID-TOKEN", idToken).
				Check().HasStatus(http.StatusNotFound).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "not_found",
					Message: "channel not found",
				})
			})
		})
	})

	t.Run("create", func(t *testing.T) {
		t.Run("valid name", func(t *testing.T) {
			name := "test1"
			checker.Test(t, http.MethodPost, "/api/channels").
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": name}).Check().HasStatus(http.StatusCreated).Cb(func(r *http.Response) {
				defer r.Body.Close()
				res := jsontest.UnmarshalFrom[server.CreateResponseBody](t, r.Body)
				expected := jsontest.Unmarshal[server.CreateResponseBody](t, jsontest.MarshalAndSnakeizeJsonKeys(t, conv.ModelToResult(&models.Channel{ID: res.ID, Name: name, CreatedAt: now, UpdatedAt: now})))
				assert.Equal(t, expected, res)
			})
		})
		t.Run("empty name", func(t *testing.T) {
			checker.Test(t, http.MethodPost, "/api/channels").
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": ""}).Check().HasStatus(http.StatusBadRequest).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "invalid_payload",
					Message: "name is required",
				})
			})
		})
		t.Run("too long name", func(t *testing.T) {
			checker.Test(t, http.MethodPost, "/api/channels").
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": strings.Repeat("a", 256)}).Check().HasStatus(http.StatusBadRequest).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "invalid_payload",
					Message: "name is too long",
				})
			})
		})
	})

	t.Run("update", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			checker.Test(t, http.MethodPut, fmt.Sprintf("/api/channels/%d", 999)).
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": "test"}).Check().HasStatus(http.StatusNotFound).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "not_found",
					Message: "channel not found",
				})
			})
		})
		t.Run("valid name", func(t *testing.T) {
			newName := ch1.Name + "-dash"
			checker.Test(t, http.MethodPut, fmt.Sprintf("/api/channels/%d", ch1.ID)).
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": newName}).Check().HasStatus(http.StatusOK).Cb(func(r *http.Response) {
				defer r.Body.Close()
				res := jsontest.UnmarshalFrom[server.UpdateResponseBody](t, r.Body)
				expected := jsontest.Unmarshal[server.UpdateResponseBody](t, jsontest.MarshalAndSnakeizeJsonKeys(t, conv.ModelToResult(&models.Channel{ID: res.ID, Name: newName, CreatedAt: now, UpdatedAt: now})))
				assert.Equal(t, expected, res)
			})
		})
		t.Run("empty name", func(t *testing.T) {
			checker.Test(t, http.MethodPut, fmt.Sprintf("/api/channels/%d", ch1.ID)).
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": ""}).Check().HasStatus(http.StatusBadRequest).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "invalid_payload",
					Message: "name is required",
				})
			})
		})
		t.Run("too long name", func(t *testing.T) {
			checker.Test(t, http.MethodPut, fmt.Sprintf("/api/channels/%d", ch1.ID)).
				WithHeader("X-ID-TOKEN", idToken).
				WithJSON(map[string]any{"name": strings.Repeat("a", 256)}).Check().HasStatus(http.StatusBadRequest).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "invalid_payload",
					Message: "name is too long",
				})
			})
		})
	})

	t.Run("delete", func(t *testing.T) {
		t.Run("invalid id", func(t *testing.T) {
			checker.Test(t, http.MethodDelete, fmt.Sprintf("/api/channels/%d", 999)).
				WithHeader("X-ID-TOKEN", idToken).
				Check().HasStatus(http.StatusNotFound).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "not_found",
					Message: "channel not found",
				})
			})
		})
		t.Run("valid id", func(t *testing.T) {
			ch1Loaded, err := models.FindChannel(ctx, conn, ch1.ID)
			assert.NoError(t, err)
			checker.Test(t, http.MethodDelete, fmt.Sprintf("/api/channels/%d", ch1.ID)).
				WithHeader("X-ID-TOKEN", idToken).
				Check().HasStatus(http.StatusOK).Cb(func(r *http.Response) {
				defer r.Body.Close()
				res := jsontest.UnmarshalFrom[server.UpdateResponseBody](t, r.Body)
				expected := jsontest.Unmarshal[server.UpdateResponseBody](t, jsontest.MarshalAndSnakeizeJsonKeys(t, conv.ModelToResult(ch1Loaded)))
				assert.Equal(t, expected, res)
			})
			//  削除後は 404 Not Found
			checker.Test(t, http.MethodGet, fmt.Sprintf("/api/channels/%d", ch1.ID)).
				WithHeader("X-ID-TOKEN", idToken).
				Check().HasStatus(http.StatusNotFound).Cb(func(r *http.Response) {
				defer r.Body.Close()
				goatest.ParseErrorBodyAndAssert(t, r.Body, &goatest.DefaultErrorResponseBody{
					Name:    "not_found",
					Message: "channel not found",
				})
			})
		})
	})
}
