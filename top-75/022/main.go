package main

func rob(nums []int) int {
	return rob_BestSolution(nums)
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func rob_BestSolution(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	var current int

	for i := 2; i < len(nums); i++ {
		current = max(prev1, prev2+nums[i])
		prev2, prev1 = prev1, current
	}

	return current
}

// Time Complexity: O(n)
// Space Complexity: O(n)
// Not so great solution because it won't solve all cases.
func rob_MyFirstSolution(nums []int) int {
	bestRob := 0
	secondBestRob := 0

	// first section
	for i := 0; i <= len(nums)-1; i += 2 {
		bestRob += nums[i]
	}

	// second section
	for i := 1; i <= len(nums)-1; i += 2 {
		secondBestRob += nums[i]
	}

	return max(bestRob, secondBestRob)
}
