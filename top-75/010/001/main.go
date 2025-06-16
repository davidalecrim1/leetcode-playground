package main

import "fmt"

// Approach 1: Brute Force O(n^2) (Time) with two for loops.
// Approach Chosen: Two pointes with decrement. O(n) (Time)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		res = max(res, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func main() {
	fmt.Printf("maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}): %v\n", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
