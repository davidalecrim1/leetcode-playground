package main

func canJump(nums []int) bool {
	// Approach 1: DFS to try every possibility (brute force).
	// A DFS because this will be a decision tree.
	// Time Complexity: O(2^n)

	// Approach 2: Apply dynamic programming with memoization.
	// This would work well.
	// Bottom Up.
	// Time Complexity: O(n^2)

	// Approach 3: Apply a greedy algorithm to optimize in a better way than DP.
	// Bottom Up.

	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		for j := 1; j <= nums[i]; j++ {
			k := i + j

			if k == len(nums) {
				break
			}

			if dp[k] {
				dp[i] = dp[k]
				break
			}
		}
	}

	return dp[0]
}
