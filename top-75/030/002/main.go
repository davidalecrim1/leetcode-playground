package main

// O(r * c)
func numIslands(grid [][]byte) int {
	// Strategy:
	// DFS in the matrix
	// Count the island as I remove them with 0's

	rows := len(grid)
	cols := len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r == rows || c == cols || grid[r][c] == '0' {
			return
		}

		// fmt.Println("setting to zero", r, c)
		grid[r][c] = '0'

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	island := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				island++

				dfs(r, c)
			}
		}
	}

	return island
}
