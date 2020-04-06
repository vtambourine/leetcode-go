package best_time_to_buy_and_sell_stock_ii

import "testing"

type test struct {
	input  []int
	expect int
}

func TestMaxProfit(t *testing.T) {
	tests := []test{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, c := range tests {
		if result := maxProfit(c.input); result != c.expect {
			t.Fatalf("maxProfit(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
