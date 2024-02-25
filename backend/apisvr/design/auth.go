package design

import "goa.design/goa/v3/dsl"

//nolint:unused
const authApiKeyScheme = "api_key"
const authApiKeyName = "auth_api_key"

//nolint:unused
var authApiKeySecurity = dsl.APIKeySecurity(authApiKeyScheme, func() {
})

func authApiKeyField(tag any) string {
	dsl.APIKeyField(tag, authApiKeyScheme, authApiKeyName, dsl.String, "Auth API Key")
	return authApiKeyName
}
