package main

// Approach: Store the max of left and right in memory
// Time Complexity: O(3n) = O(n)
// Space Complexity: O(n)
func trap(height []int) int {
	l := len(height)
	leftMax := make([]int, l)
	rightMax := make([]int, l)

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[l-1] = height[l-1]
	for i := l - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	totalWater := 0
	for i := range l {
		totalWater += max(0, min(leftMax[i], rightMax[i])-height[i])
	}

	return totalWater
}
