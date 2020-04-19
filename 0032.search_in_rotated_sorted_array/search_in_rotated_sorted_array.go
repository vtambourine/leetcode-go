package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var m int
	for l <= r {
		m = l + (r-l)/2
		if target == nums[m] {
			return m
		}
		if nums[m] < nums[r] {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
