package bitwise_and_of_numbers_range

import "testing"

type test struct {
	input  [2]int
	expect int
}

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []test{
		{[2]int{5, 7}, 4},
		{[2]int{0, 1}, 0},
		{[2]int{5000000, 2147483642}, 0},
	}

	for _, c := range tests {
		if result := rangeBitwiseAnd(c.input[0], c.input[1]); result != c.expect {
			t.Fatalf("rangeBitwiseAnd(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
