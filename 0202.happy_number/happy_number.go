package happy_number

func getNext(n int) int {
	next := 0
	for n > 0 {
		a := n % 10
		next += a * a
		n /= 10
	}
	return next
}

func isHappyMemoized(n int, memo map[int]bool) bool {
	next := getNext(n)
	if next == 1 {
		return true
	}
	if ok := memo[next]; ok {
		return false
	}
	memo[next] = true
	return isHappyMemoized(next, memo)
}

func isHappy(n int) bool {
	return isHappyMemoized(n, make(map[int]bool))
}
