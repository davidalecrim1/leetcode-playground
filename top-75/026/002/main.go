package main

func canJump(nums []int) bool {
	// Approach 3: Apply a greedy algorithm to optimize in a better way than DP.
	// Bottom Up.
	// Time Complexity: O(n)

	lastGood := len(nums) - 1
	// Last good saves the index of the "good" posotion.
	// Then I just do a sum instead of walking the array.

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastGood {
			lastGood = i
		}
	}

	return lastGood == 0
}
