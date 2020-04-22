package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	var answer, sum int
	hash := make(map[int]int)
	hash[0] = 1
	for _, n := range nums {
		sum += n
		answer += hash[sum-k]
		hash[sum]++
	}
	return answer
}
