package design

import "goa.design/goa/v3/dsl"

//nolint:unused
const authApiKeyScheme = "api_key"
const authApiKeyName = "session_id"

//nolint:unused
var authApiKeySecurity = dsl.APIKeySecurity(authApiKeyScheme, func() {
})

func authApiKeyField(tag any) string {
	// dsl.APIKeyField(tag, sessionIdScheme, sessionIdKey, dsl.String, "Session ID")
	dsl.Field(tag, authApiKeyName, dsl.String, "Session ID")
	return authApiKeyName
}
