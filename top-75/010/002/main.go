package main

func maxArea(height []int) int {
	// Strategy
	// Apply sliding window
	// Use a strategy to move the lowest bar
	// Store the max

	left := 0
	right := len(height) - 1
	res := 0

	for left < right {
		// fmt.Println("right", right)
		// fmt.Println("left", left)

		res = max(res, (right-left)*min(height[right], height[left]))

		// fmt.Println("res", res)

		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}

	return res
}
