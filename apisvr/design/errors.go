package design

import (
	"goa.design/goa/v3/dsl"
)

func userError(key string, httpStatusCode int, grpcCode int) func() (func(), func()) {
	return func() (func(), func()) {
		dsl.Error(key)
		httpFunc := func() {
			dsl.Response(key, httpStatusCode)
		}
		grpcFunc := func() {
			dsl.Response(key, grpcCode)
		}
		return httpFunc, grpcFunc
	}
}

var (
	unauthenticated = userError("unauthenticated", dsl.StatusUnauthorized, dsl.CodeUnauthenticated)
	unauthorized    = userError("unauthorized", dsl.StatusForbidden, dsl.CodePermissionDenied)
	notFound        = userError("not_found", dsl.StatusNotFound, dsl.CodeNotFound)
	invalidQuery    = userError("invalid_query", dsl.StatusBadRequest, dsl.CodeInvalidArgument)
	invalidPayload  = userError("invalid_payload", dsl.StatusBadRequest, dsl.CodeInvalidArgument)
	conflict        = userError("conflict", dsl.StatusConflict, dsl.CodeAlreadyExists)
	paymentRequired = userError("payment_required", dsl.StatusPaymentRequired, dsl.CodeFailedPrecondition)
)
