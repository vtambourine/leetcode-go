package last_stone_weight

import "testing"

type test struct {
	input  []int
	expect int
}

func TestLastStoneWeight(t *testing.T) {
	tests := []test{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1, 3}, 2},
	}

	for _, c := range tests {
		if result := lastStoneWeight(c.input); result != c.expect {
			t.Fatalf("lastStoneWeight(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
