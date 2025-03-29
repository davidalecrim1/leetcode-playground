package main

import "fmt"

// Approach 1: Brute Force
// Time Complexity: O(n^2)
func rob(nums []int) int {
	return robHouse(nums, 0)
}

func robHouse(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}

	// here I have two options:
	// rob house the current house
	// skip the current and rob the next
	// i want to calculate the max of this two approaches
	option1 := nums[index] + robHouse(nums, index+2)
	option2 := robHouse(nums, index+1)
	return max(option1, option2)
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums)) // Output: 12
}
