package design

import "goa.design/goa/v3/dsl"

//nolint:unused
const sessionIdScheme = "api_key"
const sessionIdKey = "session_id"

//nolint:unused
var sessionAuth = dsl.APIKeySecurity(sessionIdScheme, func() {
})

func fieldSessionID(tag any) string {
	// dsl.APIKeyField(tag, sessionIdScheme, sessionIdKey, dsl.String, "Session ID")
	dsl.Field(tag, sessionIdKey, dsl.String, "Session ID")
	return sessionIdKey
}
