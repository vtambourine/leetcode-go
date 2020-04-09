package backspace_string_compare

import "testing"

type test struct {
	input  [2]string
	expect bool
}

func TestBackspaceCompare(t *testing.T) {
	tests := []test{
		{[2]string{"ab#c", "ad#c"}, true},
		{[2]string{"ab##", "c#d#"}, true},
		{[2]string{"a##c", "#a#c"}, true},
		{[2]string{"a#c", "b"}, false},
		{[2]string{"bxj##tw", "bxo#j##tw"}, true},
	}

	for _, c := range tests {
		if result := backspaceCompare(c.input[0], c.input[1]); result != c.expect {
			t.Fatalf("backspaceCompare(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
