package main

// Use a sliding window
// Start at 0 and N (len).
// Calculate the max water based on the min of both lines.
// Move the keep the max and move the min between both of them.
// Do while the left < right

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	maxWater := 0
	for left < right {
		distance := right - left
		minHeight := min(height[right], height[left])
		maxWater = max(maxWater, minHeight*distance)

		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}

	return maxWater
}
