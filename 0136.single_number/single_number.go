package single_number

func singleNumber(nums []int) int {
	mask := 0

	for _, n := range nums {
		mask ^= n
	}

	return mask
}
