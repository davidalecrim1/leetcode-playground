package main

import "slices"

// Brute force this with three nested loops - Time: O(n^3)

// Target is 0
// Sort the array
// Take on element at a time
// Use a sliding window to increase/decrease to 0
// If the match is found, save it.
// Time: O(nˆ2 * n log n)
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	//fmt.Println(nums)

	res := [][]int{}

	for n := range nums {
		curr := n
		left := n + 1
		right := len(nums) - 1

		if n > 0 && nums[curr] == nums[n-1] {
			continue
		}

		for left < right {
			total := nums[curr] + nums[right] + nums[left]

			if total < 0 {
				left++
			}

			if total > 0 {
				right--
			}

			if total == 0 {
				res = append(res, []int{nums[curr], nums[left], nums[right]})

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right+1] == nums[right] {
					right--
				}
			}

		}
	}

	return res
}
