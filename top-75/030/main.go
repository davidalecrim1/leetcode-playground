package main

// Time Complexity: O(m * n)
func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	res := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				dfs(grid, r, c)
				res++
			}
		}
	}

	return res
}

func dfs(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r == len(grid) || c == len(grid[0]) {
		return
	}

	if grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'

	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
