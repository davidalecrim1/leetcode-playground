package main

func missingNumber(nums []int) int {
	// Strategy:
	// Sum the numbers and return the diff between expected and actual.

	// Time: O(n)
	// Space: O(1)

	expected := 0
	actual := 0

	for i := 0; i <= len(nums); i++ {
		expected += i

		if i != len(nums) {
			actual += nums[i]
		}
	}

	return expected - actual
}
