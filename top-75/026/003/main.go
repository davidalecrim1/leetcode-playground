package main

// Time Complexity: O(n^2)
// Space Complexty: O(n)
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)

	dp[0] = true

	for i := range n {
		if dp[i] {
			steps := nums[i]

			for j := 1; j <= steps && i+j < n; j++ {
				dp[i+j] = true
			}
		}
	}

	return dp[n-1]
}
