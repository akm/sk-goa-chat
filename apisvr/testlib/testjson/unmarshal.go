package testjson

import (
	"encoding/json"
	"io"
	"testing"
)

func Unmarshal[T any](t *testing.T, b []byte) *T {
	t.Helper()
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	return &v
}

func UnmarshalFrom[T any](t *testing.T, reader io.Reader) *T {
	t.Helper()
	b, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	return Unmarshal[T](t, b)
}
