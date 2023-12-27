package design

import "goa.design/goa/v3/dsl"

func userError(key string, httpStatusCode int) func() func() {
	return func() func() {
		dsl.Error(key)
		return func() {
			dsl.Response(key, httpStatusCode)
		}
	}
}

var (
	unauthenticated = userError("unauthenticated", dsl.StatusUnauthorized)
	unauthorized    = userError("unauthorized", dsl.StatusForbidden)
	notFound        = userError("not_found", dsl.StatusNotFound)
	invalidQuery    = userError("invalid_query", dsl.StatusBadRequest)
	invalidPayload  = userError("invalid_payload", dsl.StatusBadRequest)
	conflict        = userError("conflict", dsl.StatusConflict)
	paymentRequired = userError("payment_required", dsl.StatusPaymentRequired)
)
