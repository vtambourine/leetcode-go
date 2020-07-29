package remove_duplicates_from_sorted_array

import (
	"reflect"
	"testing"
)

type test struct {
	input  []int
	length int
	expect []int
}

func TestTemplate(t *testing.T) {
	tests := []test{
		{
			[]int{1, 1, 2},
			2,
			[]int{1, 2},
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, c := range tests {
		if length := removeDuplicates(c.input); !reflect.DeepEqual(c.input[:length], c.expect) {
			t.Fatalf("removeDuplicates(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, c.input[:length])
		}
	}
}
