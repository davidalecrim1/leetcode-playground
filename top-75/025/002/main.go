package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] += dp[i-1][j] // left to right
			}

			if j > 0 {
				dp[i][j] += dp[i][j-1] // top to bottom
			}
		}
	}

	// Since the movement is restricted from left to right and top to bottom.
	// The know the dp[i][j] = dp[i-1][j] + dp[i][j-1]

	return dp[m-1][n-1]
}
