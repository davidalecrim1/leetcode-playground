package main

func maxSubArray(nums []int) int {
	return maxSubArrayLikeSlidingWindow(nums)
}

// Time complexity: O(n)
func maxSubArrayLikeSlidingWindow(nums []int) int {
	maxSum, currentSum := nums[0], 0

	for i := 0; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = 0 // remove the previous elements if the sum is negative
		}

		currentSum += nums[i]
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// Time complexity: O(n^2)
func maxSubArrayBruteForce(nums []int) int {
	ans := nums[0]

	for i := 0; i < len(nums); i++ {
		sum := 0

		for j := i; j < len(nums); j++ {
			sum += nums[j]
			ans = max(ans, sum)
		}
	}
	return ans
}
