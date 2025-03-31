package main

// Approach: DFS without memoization
// Time Complexity: O(2^N)

/*
This creates a binary recursion tree, where each function call branches into two more calls.
The recursion depth is at most O(n), and since each call spawns two new calls,
the number of recursive calls follows an exponential growth,
leading to a time complexity of O(2ⁿ).
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(robHouse(0, nums[1:]), robHouse(0, nums[:len(nums)-1]))
}

func robHouse(start int, nums []int) int {
	if start >= len(nums) {
		return 0
	}

	option1 := nums[start] + robHouse(start+2, nums)
	option2 := robHouse(start+1, nums)

	return max(option1, option2)
}
