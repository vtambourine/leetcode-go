package two_sum

func twoSum(nums []int, target int) []int  {
	store := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if j, ok := store[diff]; ok {
			return []int{j, i}
		} else {
			store[n] = i
		}
	}

	return nil
}
