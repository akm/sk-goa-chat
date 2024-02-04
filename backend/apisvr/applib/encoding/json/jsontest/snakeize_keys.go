package jsontest

import (
	"encoding/json"
	"testing"

	"github.com/iancoleman/strcase"
)

func SnakeizeJsonKeys(t *testing.T, b []byte) []byte {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	r, err := json.Marshal(SnakeizeKeys(m))
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func SnakeizeKeys(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		r := make(map[string]interface{})
		for k, v := range v {
			newKey := strcase.ToSnake(k)
			r[newKey] = SnakeizeKeys(v)
		}
		return r
	case []interface{}:
		r := make([]interface{}, len(v))
		for i, val := range v {
			r[i] = SnakeizeKeys(val)
		}
		return r
	default:
		return v
	}
}

func MarshalAndSnakeizeJsonKeys(t *testing.T, v interface{}) []byte {
	return SnakeizeJsonKeys(t, Marshal(t, v))
}
