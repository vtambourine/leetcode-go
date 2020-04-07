package counting_elements

func countElements(arr []int) int {
	counts := make(map[int]bool)
	for _, n := range arr {
		counts[n] = true
	}

	total := 0
	for _, n := range arr {
		if counts[n+1] {
			total++
		}
	}

	return total
}
