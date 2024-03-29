// Code generated by goa v3.14.1, DO NOT EDIT.
//
// notifications HTTP server encoders and decoders
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	"context"
	"errors"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// DecodeSubscribeRequest returns a decoder for requests sent to the
// notifications subscribe endpoint.
func DecodeSubscribeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			idToken string
			err     error
		)
		idToken = r.URL.Query().Get("token")
		if idToken == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("id_token", "query string"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewSubscribePayload(idToken)

		return payload, nil
	}
}

// EncodeSubscribeError returns an encoder for errors returned by the subscribe
// notifications endpoint.
func EncodeSubscribeError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewSubscribeUnauthenticatedResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
