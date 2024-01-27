package design

import (
	. "goa.design/goa/v3/dsl"
)

func sessionIdCookie() {
	Cookie(sessionIdKey) // Return session ID in "SID" cookie
	CookieMaxAge(3600)   // Sessions last one hour
	CookiePath("/")
	CookieHTTPOnly()
}

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
			Response(StatusCreated, sessionIdCookie)
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
		Result(func() {
			Required(field(1, sessionIdKey, String, "Session ID"))
		})

		HTTP(func() {
			DELETE("")
			Cookie(sessionIdKey)
			Response(StatusOK, sessionIdCookie)
			httpInvalidPayload()
		})

		GRPC(func() {
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})
})
