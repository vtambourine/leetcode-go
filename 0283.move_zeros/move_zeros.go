package move_zeros

func moveZeroes(nums []int) {
	var z int
	for i, n := range nums {
		if n != 0 {
			nums[i], nums[z] = nums[z], nums[i]
			z++
		}
	}
}
