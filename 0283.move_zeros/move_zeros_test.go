package move_zeros

import (
	"reflect"
	"testing"
)

type test struct {
	input  []int
	expect []int
}

func TestTemplate(t *testing.T) {
	tests := []test{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 1, 0, 3, 0, 12}, []int{1, 3, 12, 0, 0, 0}},
		{[]int{0, 0, 1, 3, 12, 0, 0}, []int{1, 3, 12, 0, 0, 0, 0}},
		{[]int{1, 3, 12}, []int{1, 3, 12}},
		{[]int{1, 3, 12, 0, 0}, []int{1, 3, 12, 0, 0}},
	}

	for _, c := range tests {
		result := make([]int, len(c.input))
		copy(result, c.input)
		moveZeroes(result)
		if !reflect.DeepEqual(result, c.expect) {
			t.Fatalf("moveZeroes(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
