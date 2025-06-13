package main

func missingNumber(nums []int) int {
	// Strategy: XOR acts like a subtraction if no further operation is done.
	res := 0

	for i := 0; i <= len(nums); i++ {
		if i == len(nums) {
			res = res ^ i
		} else {
			res = res ^ (i ^ nums[i])
		}
	}

	return res
}
