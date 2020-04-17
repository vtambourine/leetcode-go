package number_of_islands

import "testing"

type test struct {
	input  [][]byte
	expect int
}

func TestNumIslands(t *testing.T) {
	tests := []test{
		{[][]byte{{'0'}}, 0},
		{[][]byte{{'1'}}, 1},
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, 3},
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
	}

	for _, c := range tests {
		if result := numIslands(c.input); result != c.expect {
			t.Fatalf("numIslands(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
