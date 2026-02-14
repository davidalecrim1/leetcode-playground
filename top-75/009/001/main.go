package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums) // O(n log n)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			// The goal here is to avoid duplicates since the array is sorted
			continue
		}

		// Approach: use two pointers to solve the two sum
		left, right := i+1, len(nums)-1

		for left < right {
			threeSummed := num + nums[left] + nums[right]

			if threeSummed > 0 {
				// Decrease the sum by moving the right pointer to lower value (since the nums is sorted)
				right--
			}

			if threeSummed < 0 {
				left++
			}

			if threeSummed == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				left++

				// Make sure the skip same numbers
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Printf("threeSum([]int{0, 0, 0, 0}): %v\n", threeSum([]int{0, 0, 0, 0}))
}
