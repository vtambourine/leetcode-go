package maximum_subarray

import "testing"

type test struct {
	input  []int
	expect int
}

func TestMaxSubArray(t *testing.T) {
	tests := []test{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-1, -2, -3, -4}, -1},
	}

	for _, c := range tests {
		if result := maxSubArray(c.input); result != c.expect {
			t.Fatalf("maxSubArray(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
