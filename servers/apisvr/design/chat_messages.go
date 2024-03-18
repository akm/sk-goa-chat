package design

import (
	"applib/time"
	"slices"

	. "goa.design/goa/v3/dsl"
)

var chatMessageActions = []string{
	"list",
	"show",
	"create",
	"update",
	"delete",
}

func chatMessageFields(action string) []string {
	if !slices.Contains[[]string, string](chatMessageActions, action) {
		panic("unknown chat_messages action: " + action)
	}

	r := []string{}

	if InPayload() {
		r = append(r, idTokenApiKeyField(1))
	}

	if InRT() || action == "update" {
		r = append(r, field(2, "id", UInt64, "ID"))
	}
	if InRT() {
		r = append(r, field(3, "created_at", String, "CreatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
		r = append(r, field(4, "updated_at", String, "UpdatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
	}

	if InRT() || action == "create" {
		r = append(r, field(5, "channel_id", UInt64, "Channel ID"))
	}

	if InRT() {
		field(6, "user_id", UInt64, "User ID") // user_id は必須ではない
		r = append(r, field(7, "user_name", String, "User Name"))
	}

	r = append(r, field(8, "content", String, "Content"))

	return r
}

var ChatMessageRT = ResultType("application/vnd.chat-message", func() {
	Attributes(func() {
		Required(chatMessageFields("show")...)
	})
})
var ChatMessageListItemRT = ResultType("application/vnd.chat-message-list-item", func() {
	Attributes(func() {
		Required(chatMessageFields("list")...)
	})
})
var ChatMessageListRT = ResultType("application/vnd.chat-message-list", func() {
	Attributes(func() {
		Required(
			field(2, "items", CollectionOf(ChatMessageListItemRT), "Items"),
		)
	})
})

var ChatMessageCreatePayload = Type("ChatMessageCreatePayload", func() {
	Required(chatMessageFields("create")...)
})
var ChatMessageUpdatePayload = Type("ChatMessageUpdatePayload", func() {
	Required(chatMessageFields("update")...)
})

func chatMessageListPayloadAttrs() ([]string, []string) {
	requiredCookies := []string{
		idTokenApiKeyField(1),
	}
	requiredParams := []string{
		field(2, "limit", Int, "Limit", func() { Default(50) }),
	}
	optionalParams := []string{
		field(3, "channel_id", UInt64, "Channel ID"),
		field(4, "after", UInt64, "ChatMessage ID for query to get messages after this"),
		field(5, "before", UInt64, "ChatMessage ID for query to get messages before this"),
	}
	return append(requiredCookies, requiredParams...), append(requiredParams, optionalParams...)
}

var _ = Service("chat_messages", func() {
	httpIdToken, grpcIdToken := idTokenSecurity()

	httpUnautheticated, grpcUnauthenticated := unauthenticated()

	HTTP(func() {
		Path("/api/chat_messages")
		httpUnautheticated()
	})

	GRPC(func() {
		grpcUnauthenticated()
	})

	Method("list", func() {
		var paramsAttrs []string
		Payload(func() {
			var requiredAttrs []string
			requiredAttrs, paramsAttrs = chatMessageListPayloadAttrs()
			Required(requiredAttrs...)
		})

		Result(ChatMessageListRT)
		HTTP(func() {
			GET("")
			httpIdToken()
			Response(StatusOK)
			Params(func() {
				for _, attr := range paramsAttrs {
					Param(attr)
				}
			})
		})
		GRPC(func() {
			grpcIdToken()
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Payload(func() {
			Required(
				idTokenApiKeyField(1),
				field(2, "id", UInt64, "ID"),
			)
		})
		httpNotFound, grpcNotFound := notFound()
		Result(ChatMessageRT)

		HTTP(func() {
			GET("/{id}")
			httpIdToken()
			Response(StatusOK)
			httpNotFound()
		})
		GRPC(func() {
			grpcIdToken()
			Response(CodeOK)
			grpcNotFound()
		})
	})

	Method("create", func() {
		Payload(ChatMessageCreatePayload)
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()
		Result(ChatMessageRT)

		HTTP(func() {
			POST("")
			httpIdToken()
			Response(StatusCreated)
			httpInvalidPayload()
		})
		GRPC(func() {
			grpcIdToken()
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})

	Method("update", func() {
		Payload(ChatMessageUpdatePayload)
		httpNotFound, grpcNotFound := notFound()
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()
		Result(ChatMessageRT)

		HTTP(func() {
			PUT("/{id}")
			httpIdToken()
			Response(StatusOK)
			httpNotFound()
			httpInvalidPayload()
		})
		GRPC(func() {
			grpcIdToken()
			Response(CodeOK)
			grpcNotFound()
			grpcInvalidPayload()
		})
	})

	Method("delete", func() {
		Payload(func() {
			Required(
				idTokenApiKeyField(1),
				field(2, "id", UInt64, "ID"),
			)
		})
		httpNotFound, grpcNotFound := notFound()
		Result(ChatMessageRT)

		HTTP(func() {
			DELETE("/{id}")
			httpIdToken()
			Response(StatusOK)
			httpNotFound()
		})
		GRPC(func() {
			grpcIdToken()
			Response(CodeOK)
			grpcNotFound()
		})
	})
})
