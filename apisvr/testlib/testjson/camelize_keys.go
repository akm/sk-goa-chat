package testjson

import (
	"encoding/json"
	"testing"

	"github.com/iancoleman/strcase"
)

func CamelizeJsonKeys(t *testing.T, b []byte) []byte {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	r, err := json.Marshal(CamelizeKeys(m))
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func CamelizeKeys(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		r := make(map[string]interface{})
		for k, v := range v {
			newKey := strcase.ToCamel(k)
			r[newKey] = CamelizeKeys(v)
		}
		return r
	case []interface{}:
		r := make([]interface{}, len(v))
		for i, val := range v {
			r[i] = CamelizeKeys(val)
		}
		return r
	default:
		return v
	}
}
