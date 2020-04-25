package jump_game

import "testing"

type test struct {
	input  []int
	expect bool
}

func TestCanJump(t *testing.T) {
	tests := []test{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for _, c := range tests {
		if result := canJump(c.input); result != c.expect {
			t.Fatalf("canJump((%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
