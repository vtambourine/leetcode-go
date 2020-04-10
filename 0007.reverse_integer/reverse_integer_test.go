package reverse_integer

import "testing"

type test struct {
	input  int
	expect int
}

func TestReverse(t *testing.T) {
	tests := []test{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for _, c := range tests {
		if result := reverse(c.input); result != c.expect {
			t.Fatalf("reverse(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
