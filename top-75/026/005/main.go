package main

func canJump(nums []int) bool {
	// Strategy:
	// Brute Force
	// Time: (n * m where m is the largest value in nums)
	// Space: O(n)

	// Greedy
	// Version with O(n) memory and O(1).

	maxReach := 0

	for i := range nums {
		if maxReach < i {
			return false
		}

		maxReach = max(maxReach, i+nums[i])
	}

	return true
}
