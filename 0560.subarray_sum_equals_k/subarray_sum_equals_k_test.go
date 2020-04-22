package subarray_sum_equals_k

import "testing"

type test struct {
	nums   []int
	sum    int
	expect int
}

func TestSubarraySum(t *testing.T) {
	tests := []test{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 1, 1, 1}, 2, 3},
		{[]int{1, -1, 1, -1, 1}, 0, 6},
	}

	for _, c := range tests {
		if result := subarraySum(c.nums, c.sum); result != c.expect {
			t.Fatalf("subarraySum(%v, %v) fails.\nExpected %v\nReceived %v", c.nums, c.sum, c.expect, result)
		}
	}
}
