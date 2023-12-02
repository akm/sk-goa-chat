package design

import (
	"time"

	. "goa.design/goa/v3/dsl"
)

func channelFields(dt string, action string) []string {
	r := []string{}

	if dt == "RT" || action == "update" {
		r = append(r, field(1, "id", UInt64, "ID"))
	}
	if dt == "RT" {
		r = append(r, field(2, "created_at", String, "CreatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
		r = append(r, field(3, "updated_at", String, "UpdatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
	}

	r = append(r, field(4, "name", String, "Name"))

	return r
}

var ChannelRT = ResultType("application/vnd.channel", func() {
	Attributes(func() {
		Required(channelFields("RT", "show")...)
	})
})
var ChannelListItemRT = ResultType("application/vnd.channel-list-item", func() {
	Attributes(func() {
		Required(channelFields("RT", "list")...)
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
	Required(channelFields("Payload", "create")...)
})
var ChannelUpdatePayload = Type("ChannelUpdatePayload", func() {
	Required(channelFields("Payload", "update")...)
})

var _ = Service("channels", func() {
	HTTP(func() {
		Path("/api/channels")
	})

	Method("list", func() {
		Result(ChannelListRT)
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
		gRPC()
	})

	Method("show", func() {
		Result(ChannelRT)
		Payload(func() { Required(field(1, "id", UInt64, "ID")) })

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
		gRPC()
	})

	Method("create", func() {
		Result(ChannelRT)
		Payload(ChannelCreatePayload)

		HTTP(func() {
			POST("")
			Response(StatusCreated)
		})
		gRPC()
	})

	Method("update", func() {
		Result(ChannelRT)
		Payload(ChannelUpdatePayload)

		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
		gRPC()
	})

	Method("delete", func() {
		Result(ChannelRT)
		Payload(func() { Required(field(1, "id", UInt64, "ID")) })

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
		})
		gRPC()
	})
})
