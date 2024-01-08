package design

import (
	"time"

	. "goa.design/goa/v3/dsl"
)

func channelFields(dt DataType, action string) []string {
	r := []string{}

	r = append(r, fieldSessionID(1))

	if dt.IsRT() || action == "update" {
		r = append(r, field(2, "id", UInt64, "ID"))
	}
	if dt.IsRT() {
		r = append(r, field(3, "created_at", String, "CreatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
		r = append(r, field(4, "updated_at", String, "UpdatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
	}

	r = append(r, field(5, "name", String, "Name"))

	return r
}

var ChannelRT = ResultType("application/vnd.channel", func() {
	Attributes(func() {
		Required(channelFields(DataTypeResultType, "show")...)
	})
})
var ChannelListItemRT = ResultType("application/vnd.channel-list-item", func() {
	Attributes(func() {
		Required(channelFields(DataTypeResultType, "list")...)
	})
})
var ChannelListRT = ResultType("application/vnd.channel-list", func() {
	Attributes(func() {
		Required(
			field(1, "items", CollectionOf(ChannelListItemRT), "Items"),
			field(2, "total", UInt64, "Total number of items", func() { Example(160) }),
			field(3, "offset", UInt64, "Offset", func() { Example(0) }),
		)
	})
})

var ChannelCreatePayload = Type("ChannelCreatePayload", func() {
	Required(channelFields(DataTypePayload, "create")...)
})
var ChannelUpdatePayload = Type("ChannelUpdatePayload", func() {
	Required(channelFields(DataTypePayload, "update")...)
})

var _ = Service("channels", func() {
	// Security(sessionAuth)

	httpUnautheticated, grpcUnauthenticated := unauthenticated()

	HTTP(func() {
		Path("/api/channels")
		httpUnautheticated()
	})

	GRPC(func() {
		grpcUnauthenticated()
	})

	Method("list", func() {
		Payload(func() { Required(fieldSessionID(1)) })
		Result(ChannelListRT)
		HTTP(func() {
			GET("")
			Cookie(sessionIdKey)
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Payload(func() {
			Required(
				fieldSessionID(1),
				field(2, "id", UInt64, "ID"),
			)
		})
		httpNotFound, grpcNotFound := notFound()
		Result(ChannelRT)

		HTTP(func() {
			GET("/{id}")
			Cookie(sessionIdKey)
			Response(StatusOK)
			httpNotFound()
		})
		GRPC(func() {
			Response(CodeOK)
			grpcNotFound()
		})
	})

	Method("create", func() {
		Payload(ChannelCreatePayload)
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()
		Result(ChannelRT)

		HTTP(func() {
			POST("")
			Cookie(sessionIdKey)
			Response(StatusCreated)
			httpInvalidPayload()
		})
		GRPC(func() {
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})

	Method("update", func() {
		Payload(ChannelUpdatePayload)
		httpNotFound, grpcNotFound := notFound()
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()
		Result(ChannelRT)

		HTTP(func() {
			PUT("/{id}")
			Cookie(sessionIdKey)
			Response(StatusOK)
			httpNotFound()
			httpInvalidPayload()
		})
		GRPC(func() {
			Response(CodeOK)
			grpcNotFound()
			grpcInvalidPayload()
		})
	})

	Method("delete", func() {
		Payload(func() {
			Required(
				fieldSessionID(1),
				field(2, "id", UInt64, "ID"),
			)
		})
		httpNotFound, grpcNotFound := notFound()
		Result(ChannelRT)

		HTTP(func() {
			DELETE("/{id}")
			Cookie(sessionIdKey)
			Response(StatusOK)
			httpNotFound()
		})
		GRPC(func() {
			Response(CodeOK)
			grpcNotFound()
		})
	})
})
