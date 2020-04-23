package bitwise_and_of_numbers_range

func rangeBitwiseAnd(m int, n int) int {
	if n >= m<<1 {
		return 0
	}
	for n > m {
		n &= n - 1
	}
	return n
}
