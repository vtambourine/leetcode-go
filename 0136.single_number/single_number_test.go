package single_number

import "testing"

type test struct {
	input  []int
	expect int
}

func TestSingleNumber(t *testing.T) {
	tests := []test{
		{[]int{1}, 1},
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, c := range tests {
		if result := singleNumber(c.input); result != c.expect {
			t.Fatalf("singleNumber(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
