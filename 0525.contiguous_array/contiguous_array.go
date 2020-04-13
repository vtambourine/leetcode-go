package contiguous_array

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	balances := make(map[int]int)
	balances[0] = 0
	maxLength := 0
	b := 0
	for i, n := range nums {
		if n == 0 {
			b -= 1
		} else {
			b += 1
		}
		if balance, ok := balances[b]; ok {
			maxLength = max(maxLength, i-balance+1)
		} else {
			balances[b] = i + 1
		}

	}
	return maxLength
}
