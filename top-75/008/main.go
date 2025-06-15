package main

import "fmt"

func search(nums []int, target int) int {
	// Strategy: Apply Binary Search
	// TIme: O(log n)

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// fmt.Printf("mid: %d\n", mid)
		// fmt.Printf("right: %d\n", right)
		// fmt.Printf("left: %d\n\n", left)

		// left sorted portion
		if nums[left] <= nums[mid] {
			if target > nums[mid] || target < nums[left] {
				// go to the right sorted portion
				left = mid + 1
			} else {
				right = mid - 1
			}

			// right sorted portion
		} else {
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Printf("search([]int{4, 5, 6, 7, 0, 1, 2}, 0): %v\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
