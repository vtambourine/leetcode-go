package perform_string_shifts

import "testing"

type test struct {
	string string
	shift  [][]int
	expect string
}

func TestStringShift(t *testing.T) {
	tests := []test{
		{"abc", [][]int{{0, 1}}, "bca"},
		{"abc", [][]int{{1, 1}}, "cab"},
		{"abc", [][]int{{0, 1}, {1, 2}}, "cab"},
		{"abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}, "efgabcd"},
	}

	for _, c := range tests {
		if result := stringShift(c.string, c.shift); result != c.expect {
			t.Fatalf("stringShift(%v, %v) fails.\nExpected %v\nReceived %v", c.string, c.shift, c.expect, result)
		}
	}
}
