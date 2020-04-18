package minimum_path_sum

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	height, width := len(grid), len(grid[0])
	var row, column int
	for row < height || column < width {
		if row < height && column < width && row > 0 && column > 0 {
			grid[row][column] += min(grid[row-1][column], grid[row][column-1])
		}

		if row < height {
			for c := column + 1; c < width; c++ {
				if row == 0 {
					grid[row][c] += grid[row][c-1]
				} else {
					grid[row][c] += min(grid[row-1][c], grid[row][c-1])
				}
			}
		}

		if column < width {
			for r := row + 1; r < height; r++ {
				if column == 0 {
					grid[r][column] += grid[r-1][column]
				} else {
					grid[r][column] += min(grid[r-1][column], grid[r][column-1])
				}
			}
		}

		row++
		column++
	}

	return grid[height-1][width-1]
}
