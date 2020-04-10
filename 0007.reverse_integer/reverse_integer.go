package reverse_integer

import "math"

func reverse(x int) int {
	out := 0
	for x != 0 {
		out = out*10 + x%10
		x = x / 10
		if out > math.MaxInt32 || out < math.MinInt32 {
			return 0
		}
	}
	return out
}
