package contiguous_array

func findMaxLength(nums []int) int {
	balances := make(map[int]int, len(nums))
	balances[0] = 0
	var balance, length, maxLength int
	for i, n := range nums {
		if n == 0 {
			balance -= 1
		} else {
			balance += 1
		}
		if b, ok := balances[balance]; ok {
			length = i - b + 1
			if length > maxLength {
				maxLength = length
			}
		} else {
			balances[balance] = i + 1
		}

	}
	return maxLength
}
