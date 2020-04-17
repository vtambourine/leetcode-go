package number_of_islands

func walkIsland(grid *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	walkIsland(grid, i-1, j)
	walkIsland(grid, i+1, j)
	walkIsland(grid, i, j-1)
	walkIsland(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				result++
				walkIsland(&grid, i, j)
			}
		}
	}
	return result
}
