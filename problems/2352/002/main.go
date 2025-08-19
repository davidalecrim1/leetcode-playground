package main

import "fmt"

// Time Complexity: O(nˆ2)
// Space Complexity: O(n)
func equalPairs(grid [][]int) int {
	n := len(grid)

	hm := make(map[string]int, n)

	// Store rows in the hashmap with the amount of rows for that key
	for r := 0; r < n; r++ {
		hm[fmt.Sprint(grid[r])]++
	}

	res := 0

	for c := 0; c < n; c++ {
		col := make([]int, n)
		for r := 0; r < n; r++ {
			col[r] = grid[r][c]
		}

		if val, ok := hm[fmt.Sprint(col)]; ok {
			res += val
		}
	}

	return res
}
