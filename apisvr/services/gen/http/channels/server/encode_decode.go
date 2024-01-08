// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels HTTP server encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	channelsviews "apisvr/services/gen/channels/views"
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the channels
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*channelsviews.ChannelList)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the channels list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			sessionID string
			err       error
			c         *http.Cookie
		)
		c, err = r.Cookie("session_id")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_id", "cookie"))
		} else {
			sessionID = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(sessionID)

		return payload, nil
	}
}

// EncodeListError returns an encoder for errors returned by the list channels
// endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthenticated":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListUnauthenticatedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowResponse returns an encoder for responses returned by the channels
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*channelsviews.Channel)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the channels show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id        uint64
			sessionID string
			err       error
			c         *http.Cookie

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		c, err = r.Cookie("session_id")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_id", "cookie"))
		} else {
			sessionID = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id, sessionID)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show channels
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthenticated":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewShowUnauthenticatedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateResponse returns an encoder for responses returned by the
// channels create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*channelsviews.Channel)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the channels
// create endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			sessionID string
			c         *http.Cookie
		)
		c, err = r.Cookie("session_id")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_id", "cookie"))
		} else {
			sessionID = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateChannelCreatePayload(&body, sessionID)

		return payload, nil
	}
}

// EncodeCreateError returns an encoder for errors returned by the create
// channels endpoint.
func EncodeCreateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_payload":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateInvalidPayloadResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthenticated":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateUnauthenticatedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the
// channels update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*channelsviews.Channel)
		enc := encoder(ctx, w)
		body := NewUpdateResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the channels
// update endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id        uint64
			sessionID string
			c         *http.Cookie

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		c, err = r.Cookie("session_id")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_id", "cookie"))
		} else {
			sessionID = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateChannelUpdatePayload(&body, id, sessionID)

		return payload, nil
	}
}

// EncodeUpdateError returns an encoder for errors returned by the update
// channels endpoint.
func EncodeUpdateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "invalid_payload":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateInvalidPayloadResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthenticated":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateUnauthenticatedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the
// channels delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*channelsviews.Channel)
		enc := encoder(ctx, w)
		body := NewDeleteResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the channels
// delete endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id        uint64
			sessionID string
			err       error
			c         *http.Cookie

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		c, err = r.Cookie("session_id")
		if err == http.ErrNoCookie {
			err = goa.MergeErrors(err, goa.MissingFieldError("session_id", "cookie"))
		} else {
			sessionID = c.Value
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(id, sessionID)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete
// channels endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not_found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthenticated":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteUnauthenticatedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalChannelsviewsChannelListItemViewToChannelListItemResponseBody builds
// a value of type *ChannelListItemResponseBody from a value of type
// *channelsviews.ChannelListItemView.
func marshalChannelsviewsChannelListItemViewToChannelListItemResponseBody(v *channelsviews.ChannelListItemView) *ChannelListItemResponseBody {
	res := &ChannelListItemResponseBody{
		ID:        *v.ID,
		CreatedAt: *v.CreatedAt,
		UpdatedAt: *v.UpdatedAt,
		Name:      *v.Name,
	}

	return res
}
