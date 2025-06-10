package main

import "slices"

// My Default Solution with some help from Youtube
// Doesnt work in all tests
// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProduct_FirstSolution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	curr := nums[0]
	result := 0

	for i := 1; i < len(nums); i++ {
		curr *= nums[i]

		if curr == 0 {
			curr = nums[i] // make sure to preserve the current number as a subarray
		}

		result = max(result, curr)
	}

	return result
}

// Best solution learning from Youtube
// Time Complexity: O(n)
// Space Complexity: O(1)
func maxProduct(nums []int) int {
	// consider the product is the largest number if there is not a good subproduct
	res := slices.Max(nums)

	// just to make sure the minimum and maximum are considered itself (e.g. 1 * N)
	curMin, curMax := 1, 1

	for _, n := range nums {
		if n == 0 {
			curMax, curMin = 1, 1
			continue
		}

		tmp := curMax * n

		curMax = max(n*curMax, n*curMin, n)
		// here we use the minimum to keep track of the negative numbers minumum product
		// cuz since the array can end up having negative maximum product given how even
		// the positive and negative numbers are
		curMin = min(tmp, n*curMin, n)

		res = max(res, curMax)
	}

	return res
}
