package main

// Approach: Brute force approach
// O(nˆ2)
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	totalWater := 0
	for i := 0; i < len(height); i++ {
		leftMax, rightMax := 0, 0

		// find left max
		for j := 0; j <= i; j++ {
			leftMax = max(leftMax, height[j])
		}

		// find left max
		for j := i; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}

		totalWater += max(0, min(leftMax, rightMax)-height[i])
	}

	return totalWater
}
