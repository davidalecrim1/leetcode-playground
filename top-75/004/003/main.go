// Time Complexity: O(n)
// Space Complexity: (1)
package main

func maxProduct(nums []int) int {
	resMax := nums[0]
	currMin, currMax := 1, 1

	for i := range nums {
		newMax := nums[i] * currMax
		newMin := nums[i] * currMin

		currMax = max(newMax, newMin, nums[i])
		currMin = min(newMax, newMin, nums[i])

		resMax = max(currMax, currMin, resMax)
	}

	return resMax
}
