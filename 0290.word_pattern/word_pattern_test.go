package word_pattern

import "testing"

type test struct {
	pattern string
	input   string
	expect  bool
}

func TestWordPattern(t *testing.T) {
	tests := []test{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"aa", "aa aa aa aa", false},
	}

	for _, c := range tests {
		if result := wordPattern(c.pattern, c.input); result != c.expect {
			t.Fatalf("wordPattern(%v, %v) fails.\nExpected %v\nReceived %v", c.pattern, c.input, c.expect, result)
		}
	}
}
