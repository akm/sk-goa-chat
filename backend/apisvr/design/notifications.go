package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("notifications", func() {
	httpIdToken, grpcIdToken := idTokenSecurity()

	httpUnautheticated, grpcUnauthenticated := unauthenticated()

	HTTP(func() {
		Path("/ws/notifications")
		httpIdToken()
		httpUnautheticated()
	})

	GRPC(func() {
		grpcIdToken()
		grpcUnauthenticated()
	})

	Method("subscribe", func() {
		Description("Subscribe to events sent such new chat messages.")

		Payload(func() {
			Required(
				authApiKeyField(1),
			)
		})

		StreamingResult(NotificationEvent)

		HTTP(func() {
			GET("/subscribe")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var NotificationEvent = Type("NotificationEvent", func() {
	Required(
		field(1, "channel_ids", ArrayOf(UInt64), "IDs of channels which got new messages"),
	)
})
