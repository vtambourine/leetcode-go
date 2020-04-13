package contiguous_array

import "testing"

type test struct {
	input  []int
	expect int
}

func TestTemplate(t *testing.T) {
	tests := []test{
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
		{[]int{0, 1, 0, 0, 0}, 2},
		{[]int{0, 1, 0, 1, 0}, 4},
		{[]int{1, 0, 0, 0, 1, 0}, 2},
		{[]int{1, 0, 0, 0, 0, 0, 1, 0}, 2},
		{[]int{1, 1, 1, 1, 0, 0, 0, 0}, 8},
	}

	for _, c := range tests {
		if result := findMaxLength(c.input); result != c.expect {
			t.Fatalf("findMaxLength((%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
