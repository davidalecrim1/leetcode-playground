package main

// Time Complexity: O(n)
// Space Complexity: O(1)
func trap(height []int) int {
	l := len(height)

	if l <= 1 {
		return 0
	}

	left, right := 0, l-1
	leftMax, rightMax := height[left], height[right]
	res := 0

	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			res += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			res += rightMax - height[right]

		}
	}

	return res
}
