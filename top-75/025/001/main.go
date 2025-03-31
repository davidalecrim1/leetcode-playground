package main

type direction struct {
	row int
	col int
}

// Time Complexity: O(m + n)
// Space Complexity: O(m + n)
func uniquePaths(m int, n int) int {
	// Approach 1: Apply DFS given this will result in a decision tree with two posibilities.
	// Edge case: Don't go over the matrix (<= m or <= n)
	// Apply memoization to improve on sub problems?

	// Brute Force DFS without Memoization
	// Time Complexity: O(2^m+n)

	memo := make(map[direction]int)

	return backtrack(0, 0, m, n, memo)
}

func backtrack(row, col, m, n int, memo map[direction]int) int {
	if row == m-1 && col == n-1 {
		return 1
	}

	if val, ok := memo[direction{row: row, col: col}]; ok {
		return val
	}

	if row > m || col > n {
		return 0
	}

	res := backtrack(row+1, col, m, n, memo) + backtrack(row, col+1, m, n, memo)
	memo[direction{row, col}] = res

	return res
}
