package main

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	res := 0
	for i := range dp {
		res = max(res, dp[i])
	}

	return res
}
