package main

import "slices"

func missingNumber(nums []int) int {
	return missingNumber_SolutionUsingSum((nums))
}

// Time: O(n)
// Space: O(1)
// Strategy: Find the missing number using the diff
// between two arrays (complete vs incomplete)
// with a sum strategy to save space.
func missingNumber_SolutionUsingSum(nums []int) int {
	result := len(nums)

	for i, num := range nums {
		result += (i - num)
	}

	return result
}

// Time: O(n log n) worst scenario and O(n) in best scenario
// Space: O(1)
func missingNumber_MyFirstSolution(nums []int) int {
	slices.Sort(nums) // O (n log n) in worst scenario

	for i, num := range nums {
		if i != num {
			return i
		}
	}

	return len(nums)
}
