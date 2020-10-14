package detect_capital

import "testing"

type test struct {
	input  string
	expect bool
}

func TestTemplate(t *testing.T) {
	tests := []test{
		{ "USA", true},
		{ "leetcode", true},
		{ "Google", true},
		{ "FlaG", false},
	}

	for _, c := range tests {
		if result := detectCapitalUse(c.input); result != c.expect {
			t.Fatalf("detectCapitalUse(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
