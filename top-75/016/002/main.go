package main

func climbStairs(n int) int {
	// n position
	// n = n-1 || n-2

	// n position
	// n - 1 = 1 step
	// n - 2 = 1+1 or 2

	dp := make([]int, n+1)

	dp[0] = 1 // 1 because I am already at the "top"
	dp[1] = 1 // 1 because I only take one step from the "top" to n
	// dp[2] = dp[1]

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
