package design

import (
	"goa.design/goa/v3/dsl"
)

func field(tag any, name string, args ...any) string {
	dsl.Field(tag, name, args...)
	return name
}
