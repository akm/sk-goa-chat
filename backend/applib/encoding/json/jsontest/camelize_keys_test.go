package jsontest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelizeKeys(t *testing.T) {
	assert.Equal(t, map[string]interface{}{"FooBar": "baz"}, CamelizeKeys(map[string]interface{}{"foo_bar": "baz"}))
	t.Run("nested", func(t *testing.T) {
		assert.Equal(t,
			map[string]interface{}{
				"FooBar": map[string]interface{}{"BazQux": "quux"},
			},
			CamelizeKeys(map[string]interface{}{
				"foo_bar": map[string]interface{}{"baz_qux": "quux"},
			}),
		)
	})
}
