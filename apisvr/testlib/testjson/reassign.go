package testjson

import (
	"encoding/json"
	"testing"
)

func Reassign[T any](t *testing.T, v *T) *T {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var r T
	if err := json.Unmarshal(b, &r); err != nil {
		t.Fatal(err)
	}
	return &r
}
