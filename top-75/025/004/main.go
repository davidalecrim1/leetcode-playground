package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Recursion pattern
	// dp[i][j] = dp[i+1][j] + dp[i][j+1]

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = 1
			}

			if i+1 < m {
				dp[i][j] += dp[i+1][j]
			}

			if j+1 < n {
				dp[i][j] += dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
