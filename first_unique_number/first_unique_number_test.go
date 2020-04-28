package main

import "testing"

type test struct {
	input  []int
	expect int
}

func TestTemplate(t *testing.T) {
	tests := []test{}

	for _, c := range tests {
		if result := solution(c.input); result != c.expect {
			t.Fatalf("solution(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
