package main

func canJump(nums []int) bool {
	maxReach := 0

	for i := range nums {
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, nums[i]+i)
	}

	return true
}
