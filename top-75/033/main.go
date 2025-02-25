package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	// could use for loop on the array but it's O(n)
	// Input: nums = [4,5,6,7,0,1,2], target = 0
	// Output: 4

	first, last := 0, len(nums)-1

	for first <= last {
		mid := first + (last-first)/2 // == (last + first) / 2

		if target == nums[mid] {
			return mid
		}

		// decide if i walk left or right

		if nums[first] <= nums[mid] {
			if target > nums[mid] || target < nums[first] {
				// means the target is in the right
				first = mid + 1
			} else {
				last = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[last] {
				// means the target is in the left
				last = mid - 1
			} else {
				first = mid + 1
			}
		}

	}

	return -1
}

func main() {
	fmt.Printf("search([]int{3, 1}, 1): %v\n", search([]int{3, 1}, 1))
}
