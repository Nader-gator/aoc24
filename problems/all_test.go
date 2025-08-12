package problems

import (
	"testing"
)

func TestCountConnections(t *testing.T) {
	all := []map[string][]int{
		// {"q": {1, 2, 3}, "a": {1}},
		// {"q": {1, 2, 3, 4, 5}, "a": {1}},
		// {"q": {1, 2, 4, 5}, "a": {1}},
		{"q": {10, 2, 20, 5, 1, 60, 134, 3, 59}, "a": {10}},
	}
	for _, val := range all {
		res := countConnections(val["q"])
		if res != val["a"][0] {
			t.Errorf(`had %v, wanted %d, got %d`, val["a"][9], val["a"][0], res)
		}
	}
}
