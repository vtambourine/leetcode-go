package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	length := len(nums)
	output := make([]int, length)
	output[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		output[i] = nums[i+1] * output[i+1]
	}

	product := 1
	for i := 0; i < length; i++ {
		output[i] = product * output[i]
		product *= nums[i]
	}
	return output
}
