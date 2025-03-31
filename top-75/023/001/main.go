package main

import "fmt"

// Time Complexity: O(n)
// Space Complexity: O(n)
func rob(nums []int) int {
	// Approach 1: Apply DP

	if len(nums) == 1 {
		return nums[0]
	}

	return max(robHouse(nums[1:]), robHouse(nums[:len(nums)-1]))
}

func robHouse(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]
}

func main() {
	fmt.Printf("rob([]int{1, 2, 3, 1}): %v\n", rob([]int{1, 2, 3, 1}))
}
