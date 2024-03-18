package design

import "goa.design/goa/v3/dsl"

// Firebase UID による認証(プロキシで認証済みのユーザーを取得)
const uidApiKeyScheme = "uid_api_key"
const uidApiKeyName = "uid"

var uidApiKeySecurity = dsl.APIKeySecurity(uidApiKeyScheme, func() {
})

func uidApiKeyField(tag any) string {
	dsl.APIKeyField(tag, uidApiKeyScheme, uidApiKeyName, dsl.String, "X-UID", func() { dsl.Example("abcdef12345") })
	return uidApiKeyName
}

func uidSecurity() (func(), func()) {
	dsl.Security(uidApiKeySecurity)
	return func() {
			dsl.Header(uidApiKeyName + ":X-UID")
		},
		func() {
			dsl.Message(func() {
				dsl.Attribute(uidApiKeyName)
			})
		}
}

// ID Token による認証
const idTokenApiKeyScheme = "id_token_api_key"
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
