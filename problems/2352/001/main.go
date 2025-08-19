package main

// Time Complexity: O(nˆ3)
func equalPairs(grid [][]int) int {
	n := len(grid)
	res := 0

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if checkEqual(grid, r, c) {
				res++
			}
		}
	}
	return res
}

func checkEqual(grid [][]int, row, col int) bool {
	n := len(grid)

	for i := 0; i < n; i++ {
		if grid[row][i] != grid[i][col] {
			return false
		}
	}

	return true
}
