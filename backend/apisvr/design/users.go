package design

import (
	"applib/time"
	"slices"

	. "goa.design/goa/v3/dsl"
)

var userActions = []string{
	"list",
	"show",
	"create",
	// "update",
	// "delete",
}

func userFields(action string) []string {
	if !slices.Contains[[]string, string](userActions, action) {
		panic("unknown user action: " + action)
	}

	r := []string{}

	if InRT() || action == "update" {
		r = append(r, field(1, "id", UInt64, "ID"))
	}

	if InRT() && action != "list" {
		r = append(r, field(2, "created_at", String, "CreatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
		r = append(r, field(3, "updated_at", String, "UpdatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
	}

	r = append(r, field(4, "name", String, "Name"))

	if InPayload() && action == "create" {
		r = append(r, field(5, "email", String, "Email"))
	}

	return r
}

var UserRT = ResultType("application/vnd.user", func() {
	Attributes(func() {
		Required(userFields("show")...)
	})
})

var UserListItemRT = ResultType("application/vnd.user-list-item", func() {
	Attributes(func() {
		Required(userFields("list")...)
	})
})
var UserListRT = ResultType("application/vnd.user-list", func() {
	Attributes(func() {
		Required(
			field(1, "items", CollectionOf(UserListItemRT), "Items"),
			field(2, "total", UInt64, "Total number of items", func() { Example(160) }),
			field(3, "offset", UInt64, "Offset", func() { Example(0) }),
		)
	})
})

var UserCreatePayload = Type("UserCreatePayload", func() {
	Required(userFields("create")...)
})

var _ = Service("users", func() {
	HTTP(func() {
		Path("/api/users")
	})

	Method("list", func() {
		Result(UserListRT)
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("create", func() {
		Result(UserRT)
		Payload(UserCreatePayload)
		httpInvalidPayload, grpcInvalidPayload := invalidPayload()

		HTTP(func() {
			POST("")
			Response(StatusCreated)
			httpInvalidPayload()
		})
		GRPC(func() {
			Response(CodeOK)
			grpcInvalidPayload()
		})
	})
})
