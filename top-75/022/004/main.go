package main

func rob(nums []int) int {
	// Problem:
	// Rob or skip the house (2 options)

	// Strategy:
	// DFS (walk the tree) + with memo: Time and Space O(n)

	// DP
	// Time and Space O(n)

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}
