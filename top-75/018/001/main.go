package main

import (
	"fmt"
	"slices"
)

// Time Complexity: O(nˆ2)
func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))

	for i := range lis {
		lis[i] = 1
	}

	// walk backwards, and when I do, i compare the previous stored
	// LIS from the other loops.
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			// only store LIS of valid subsequence
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], 1+lis[j]) // add one is like considering [2, 3]. The two positions.
			}
		}
	}

	return slices.Max(lis)
}

func main() {
	res := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Printf("Expected 4, got: %v\n", res)
}
