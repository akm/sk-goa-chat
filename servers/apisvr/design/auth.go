package design

import "goa.design/goa/v3/dsl"

// ID Token による認証
const idTokenApiKeyScheme = "api_key"
const idTokenApiKeyName = "id_token"

var idTokenApiKeySecurity = dsl.APIKeySecurity(idTokenApiKeyScheme, func() {
})

func idTokenApiKeyField(tag any) string {
	dsl.APIKeyField(tag, idTokenApiKeyScheme, idTokenApiKeyName, dsl.String, "X-ID-TOKEN", func() { dsl.Example("abcdef12345") })
	return idTokenApiKeyName
}

func idTokenSecurity() (func(), func()) {
	dsl.Security(idTokenApiKeySecurity)
	return func() {
			dsl.Header(idTokenApiKeyName + ":X-ID-TOKEN")
		},
		func() {
			dsl.Message(func() {
				dsl.Attribute(idTokenApiKeyName)
			})
		}
}
