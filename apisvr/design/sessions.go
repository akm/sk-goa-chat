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
			Required(field(1, "session_id", String, "Session ID"))
		})

		HTTP(func() {
			POST("")
			Response(StatusCreated, func() {
				Cookie("session_id") // Return session ID in "SID" cookie
				CookieMaxAge(3600)   // Sessions last one hour
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
			Field(1, "session_id", String, "Session ID")
			Required("session_id")
		})

		Result(func() {})
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()

		HTTP(func() {
			DELETE("")
			Cookie("session_id")
			Response(StatusOK)
			httpInvalidPayload()
		})

		GRPC(func() {
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})

})
