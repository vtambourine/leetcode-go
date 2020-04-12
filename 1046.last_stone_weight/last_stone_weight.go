package last_stone_weight

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lastStoneWeight(stones []int) int {
	var x, y, l int
	for len(stones) > 1 {
		sort.Ints(stones)
		l = len(stones)
		x = stones[l-1]
		y = stones[l-2]
		stones = stones[:l-2]
		if x != y {
			stones = append(stones[:l-2], abs(x-y))
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
