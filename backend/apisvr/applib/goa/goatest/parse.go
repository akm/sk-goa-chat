package goatest

import (
	"apisvr/applib/encoding/json/jsontest"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DefaultErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

func ParseErrorBody(t *testing.T, b []byte) *DefaultErrorResponseBody {
	t.Helper()
	return jsontest.Unmarshal[DefaultErrorResponseBody](t, b)
}

func ParseErrorBodyFrom(t *testing.T, reader io.Reader) *DefaultErrorResponseBody {
	t.Helper()
	return jsontest.UnmarshalFrom[DefaultErrorResponseBody](t, reader)
}

func AssertErrorWithoutID(t *testing.T, expected, actual *DefaultErrorResponseBody) {
	practial := *actual
	practial.ID = ""
	assert.Equal(t, expected, &practial)
}

func ParseErrorBodyAndAssert(t *testing.T, reader io.Reader, expected *DefaultErrorResponseBody) {
	t.Helper()
	actual := ParseErrorBodyFrom(t, reader)
	AssertErrorWithoutID(t, expected, actual)
}
