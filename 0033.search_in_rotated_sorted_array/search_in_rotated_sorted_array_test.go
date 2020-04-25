package search_in_rotated_sorted_array

import "testing"

type test struct {
	numbers []int
	target  int
	expect  int
}

func TestTemplate(t *testing.T) {
	tests := []test{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 3, -1},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 4, 3},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 7, 6},
	}

	for _, c := range tests {
		if result := search(c.numbers, c.target); result != c.expect {
			t.Fatalf("solution(%v, %v) fails.\nExpected %v\nReceived %v", c.numbers, c.target, c.expect, result)
		}
	}
}
