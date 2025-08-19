package main

func minCostClimbingStairs(cost []int) int {
	// Approach 2:
	// Dynamic Programming
	// Understand the recursion pattern

	// Decide between index 1 and 2
	// Sum = 10 + 20 = 30
	// Sum = 15

	// Memo
	// Keep the lowest value in the current index
	// Choose between one or two index
	// Store the minimum

	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]
	// Recursion pattern
	// dp[1] = min(dp[i]+dp[i-1], dp[i]+dp[i-2])
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-2], dp[i-1])
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}
