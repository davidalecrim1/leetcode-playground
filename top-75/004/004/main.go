package main

// find the repetition pattern
// if recompute n*n
// pre calculation
// post calculation
// walk the array, if index 0, consider 1
// multiply previous precalculation[i-1] with nums[i-i] and place this on precalculation[i]
// [1, 1, 2, 6]
// do the same walking from the end
// [24,12,4,1]
// res
// [24,12,8,6]
// Time: O(n)
// Space: O(n)
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pos := make([]int, len(nums))

	for i := range nums {
		if i == 0 {
			pre[i] = 1
			continue
		}

		pre[i] = pre[i-1] * nums[i-1]
	}

	//fmt.Println(pre)

	pos[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		pos[i] = pos[i+1] * nums[i+1]
	}

	//fmt.Println(pos)

	for i := range nums {
		nums[i] = pre[i] * pos[i]
	}

	return nums
}
