package counting_elements

import "testing"

type test struct {
	input  []int
	expect int
}

func TestCountElements(t *testing.T) {
	tests := []test{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 1, 3, 3, 5, 5, 7, 7}, 0},
		{[]int{1, 3, 2, 3, 5, 0}, 3},
		{[]int{1, 1, 2, 2}, 2},
		{[]int{1, 1, 2}, 2},
	}

	for _, c := range tests {
		if result := countElements(c.input); result != c.expect {
			t.Fatalf("countElements(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
