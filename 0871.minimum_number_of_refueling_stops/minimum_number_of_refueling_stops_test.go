package minimum_number_of_refueling_stops

import "testing"

type test struct {
	target    int
	startFuel int
	stations  [][]int
	expect    int
}

func TestMinRefuelStops(t *testing.T) {
	tests := []test{
		{
			1,
			1,
			[][]int{},
			0,
		},
		{
			100,
			1,
			[][]int{
				{10, 100},
			},
			-1,
		},
		{
			100,
			10,
			[][]int{
				{10, 60},
				{20, 30},
				{30, 30},
				{60, 40},
			},
			2,
		},
	}

	for _, c := range tests {
		if result := minRefuelStops(c.target, c.startFuel, c.stations); result != c.expect {
			t.Fatalf("minRefuelStops(%v, %v, %v) fails.\nExpected %v\nReceived %v", c.target, c.startFuel, c.stations, c.expect, result)
		}
	}
}
