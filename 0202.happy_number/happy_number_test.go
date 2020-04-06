package happy_number

import "testing"

type test struct {
	input  int
	expect bool
}

func TestIsHappy(t *testing.T) {
	tests := []test{
		{19, true},
		{82, true},
		{68, true},
		{100, true},
		{11, false},
		{2, false},
		{4, false},
		{16, false},
	}

	for _, c := range tests {
		if result := isHappy(c.input); result != c.expect {
			t.Fatalf("isHappy(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
