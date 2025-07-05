package main

func pivotIndex(nums []int) int {
	// Strategy
	// Option 1:
	// Brute force this with two for loops

	// Option 2:
	// Try some dynamic programming approach
	// Pre calculate:
	// Input: [1,7,3,6,5,6]
	// DP: [1,8,11,17,22,28]
	// e.g. index 0
	// left 0
	// dp[len-1]-nums[i] = 26
	// e.g. index 3
	// left dp[i-1] = 11
	// right dp[len-1]=28 - dp[i]=17 = 11

	dp := make([]int, len(nums))

	for i := range nums {
		if i-1 < 0 {
			dp[i] = nums[i]
			continue
		}

		dp[i] = dp[i-1] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		var left int
		var right int

		if i-1 < 0 {
			left = 0
		} else {
			left = dp[i-1]
		}

		right = dp[len(nums)-1] - dp[i]
		if left == right {
			return i
		}

	}

	return -1
}
