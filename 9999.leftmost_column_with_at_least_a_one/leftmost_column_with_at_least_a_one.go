package leftmost_column_with_at_least_a_one

type BinaryMatrix interface {
	Get(int, int) int
	Dimensions() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()
	height, width := dimensions[0], dimensions[1]
	x, y := width-1, 0
	answer := width
	for x >= 0 && y < height {
		value := binaryMatrix.Get(y, x)
		if value == 0 {
			y++
		} else {
			if x < answer {
				answer = x
			}
			x--
		}
	}

	if answer == width {
		answer = -1
	}
	return answer
}
