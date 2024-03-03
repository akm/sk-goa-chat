package design

import "goa.design/goa/v3/dsl"

//nolint:unused
const authApiKeyScheme = "api_key"
const authApiKeyName = "id_token"

//nolint:unused
var authApiKeySecurity = dsl.APIKeySecurity(authApiKeyScheme, func() {
})

func authApiKeyField(tag any) string {
	dsl.APIKeyField(tag, authApiKeyScheme, authApiKeyName, dsl.String, "X-ID-TOKEN", func() { dsl.Example("abcdef12345") })
	return authApiKeyName
}

func idTokenSecurity() (func(), func()) {
	dsl.Security(authApiKeySecurity)
	return func() {
			dsl.Header(authApiKeyName + ":X-ID-TOKEN")
		},
		func() {
			dsl.Message(func() {
				dsl.Attribute(authApiKeyName)
			})
		}
}
