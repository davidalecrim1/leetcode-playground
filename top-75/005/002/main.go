package main

func maxSubArray(nums []int) int {
	// Option 1: Brute Force with Two For Loops
	// Time: O(nˆ2)

	// Option 2: Some kind of sliding window
	// Time: O(n)

	res := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		// ignore the value in the sub array in the beginning
		if sum < 0 {
			sum = 0
		}

		sum += nums[i]
		res = max(res, sum, nums[i])

	}

	return res
}
