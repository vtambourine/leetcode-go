package product_of_array_except_self

import (
	"reflect"
	"testing"
)

type test struct {
	input  []int
	expect []int
}

func TestProductExceptSelf(t *testing.T) {
	tests := []test{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			[]int{1, 2, 3, 4, 0},
			[]int{0, 0, 0, 0, 24},
		},
		{
			[]int{1, 2, 3, 4, 0, 0},
			[]int{0, 0, 0, 0, 0, 0},
		},
	}

	for _, c := range tests {
		if result := productExceptSelf(c.input); !reflect.DeepEqual(result, c.expect) {
			t.Fatalf("productExceptSelf(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
