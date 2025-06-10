package main

func maxProduct(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i]*dp[i-1])
	}

	return dp[len(nums)-1]
}
