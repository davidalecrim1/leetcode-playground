package main

func longestOnes(nums []int, k int) int {
	// Strategy
	// Create a zero tolerance
	// Use a sliding window
	// Keep track of the max possible len
	// When I go over my tolerance

	// Op1 (Go with this one)
	// Move left until the tolerance is equal to K

	// Op2:
	// Move left until tolerance is valid again

	// track of 1s and 0s
	// when move the left, see if its one or zero to change the count

	res := 0
	left := 0
	flips := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			flips++
		}

		for flips > k {
			if nums[left] == 0 {
				flips--
			}

			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
