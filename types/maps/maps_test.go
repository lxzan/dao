package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var m = map[string]any{
			"a": 1,
			"b": 1,
			"c": 1,
		}
		assert.ElementsMatch(t, Keys(m), []string{"a", "b", "c"})
	})

	t.Run("", func(t *testing.T) {
		type Map map[string]any
		var m = Map{
			"a": 1,
			"b": 1,
			"c": 1,
		}
		assert.ElementsMatch(t, Keys(m), []string{"a", "b", "c"})
	})
}

func TestValues(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var m = map[string]any{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		assert.ElementsMatch(t, Values(m), []int{1, 2, 3})
	})
}
