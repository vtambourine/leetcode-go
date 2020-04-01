package two_sum

import (
	"testing"
)

type test struct {
	numbers []int
	target  int
	expect  []int
}

func TestTwoSums(t *testing.T) {
	tests := []test{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3,2,4},
			6,
			[]int{1, 2},
		},
		{
			[]int{0,1},
			1,
			[]int{0,1},
		},
	}

	for _, c := range tests {
		if result := twoSum(c.numbers, c.target); result[0] != c.expect[0] || result[1] != c.expect[1] {
			t.Fatalf("twoSum(%v, %d) fails.\nExpected %v\nReceived %v", c.numbers, c.target, c.expect, result)
		}
	}
}
