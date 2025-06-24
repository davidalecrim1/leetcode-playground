package main

func uniquePaths(m int, n int) int {
	// Strategy
	// This will result in a decision tree.
	// Walk this with DFS
	// Create a memo to change the complexity from O(2ˆn * m) to O(m * n)

	// Alternative:
	// Maybe a true dynamic programming
	// Need to understand the decision tree better for this.

	rows, cols := m, n
	memo := make(map[[2]int]int)

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r == rows || c == cols {
			return 0
		}

		if val, ok := memo[[2]int{r, c}]; ok {
			return val
		}

		if r == rows-1 && c == cols-1 {
			return 1
		}

		steps := dfs(r+1, c) + dfs(r, c+1)
		memo[[2]int{r, c}] = steps

		return steps
	}

	return dfs(0, 0)
}
