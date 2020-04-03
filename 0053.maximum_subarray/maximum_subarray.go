package maximum_subarray

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := math.MinInt32

	for _, n := range nums {
		sum = max(n, sum+n)
		maxSum = max(maxSum, sum)
	}

	return maxSum
}
