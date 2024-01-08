package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("sessions", func() {
	HTTP(func() {
		Path("/api/session")
	})

	Method("create", func() {
		Payload(func() {
			Required(field(1, "id_token", String, "ID Token"))
		})
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()
		Result(func() {
			Required(field(1, sessionIdKey, String, "Session ID"))
		})

		HTTP(func() {
			POST("")
			Response(StatusCreated, func() {
				Cookie(sessionIdKey) // Return session ID in "SID" cookie
				CookieMaxAge(3600)   // Sessions last one hour
				CookiePath("/")
				CookieHTTPOnly()
			})
			httpInvalidPayload()
		})
		GRPC(func() {
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})

	Method("delete", func() {
		Payload(func() {
			Field(1, sessionIdKey, String, "Session ID")
			Required(sessionIdKey)
		})
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()
		Result(func() {})

		HTTP(func() {
			DELETE("")
			Cookie(sessionIdKey)
			Response(StatusOK)
			httpInvalidPayload()
		})

		GRPC(func() {
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})
})
