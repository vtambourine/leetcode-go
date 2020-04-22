package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				sum++
			}
		}
	}
	return sum
}
