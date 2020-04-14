package perform_string_shifts

func shiftOnce(arr *[]byte, shift []int) {
	direction, amount := shift[0], shift[1]
	length := len(*arr)
	amount %= length
	if direction == 0 {
		*arr = append((*arr)[amount:], (*arr)[:amount]...)
	} else if direction == 1 {
		*arr = append((*arr)[length-amount:], (*arr)[:length-amount]...)
	}
}

func stringShift(s string, shift [][]int) string {
	arr := []byte(s)
	for _, v := range shift {
		shiftOnce(&arr, v)
	}
	return string(arr)
}
