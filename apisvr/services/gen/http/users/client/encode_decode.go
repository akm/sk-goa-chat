// Code generated by goa v3.14.1, DO NOT EDIT.
//
// users HTTP client encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	users "apisvr/services/gen/users"
	usersviews "apisvr/services/gen/users/views"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "users" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListUsersPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("users", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the users
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("users", "list", err)
			}
			p := NewListUserListOK(&body)
			view := "default"
			vres := &usersviews.UserList{Projected: p, View: view}
			if err = usersviews.ValidateUserList(vres); err != nil {
				return nil, goahttp.ErrValidationError("users", "list", err)
			}
			res := users.NewUserList(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("users", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "users" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUsersPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("users", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the users create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*users.UserCreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("users", "create", "*users.UserCreatePayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("users", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the users
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "invalid_payload" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("users", "create", err)
			}
			p := NewCreateUserCreated(&body)
			view := "default"
			vres := &usersviews.User{Projected: p, View: view}
			if err = usersviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("users", "create", err)
			}
			res := users.NewUser(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateInvalidPayloadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("users", "create", err)
			}
			err = ValidateCreateInvalidPayloadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("users", "create", err)
			}
			return nil, NewCreateInvalidPayload(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("users", "create", resp.StatusCode, string(body))
		}
	}
}

// unmarshalUserListItemResponseBodyToUsersviewsUserListItemView builds a value
// of type *usersviews.UserListItemView from a value of type
// *UserListItemResponseBody.
func unmarshalUserListItemResponseBodyToUsersviewsUserListItemView(v *UserListItemResponseBody) *usersviews.UserListItemView {
	res := &usersviews.UserListItemView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}