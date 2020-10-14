package minimum_number_of_refueling_stops

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	longestRoad := make([]int, len(stations) + 1)
	longestRoad[0] = startFuel

	for i := range stations {
		for t := i; t >= 0; t-- {
			if longestRoad[t] >= stations[i][0] {
				longestRoad[t+1] = max(longestRoad[t+1], longestRoad[t] + stations[i][1])
			}
		}
	}

	for i := range longestRoad {
		if longestRoad[i] >= target {
			return i
		}
	}

	return -1
}