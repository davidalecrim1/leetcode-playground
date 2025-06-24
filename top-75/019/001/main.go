package main

func longestCommonSubsequence(text1 string, text2 string) int {
	// Strategy
	// Apply a DFS in the decision tree.
	// Create a memo for the repeated sub problems using a matrix.

	dp := make([][]int, len(text1)+1)
	for i := range len(dp) {
		dp[i] = make([]int, len(text2)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		if text1[i] == text2[j] {
			dp[i][j] = 1 + dfs(i+1, j+1)
		} else {
			dp[i][j] = max(dfs(i+1, j), dfs(i, j+1))
		}

		return dp[i][j]
	}

	return dfs(0, 0)
}
