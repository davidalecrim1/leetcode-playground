package main

func removeDuplicates(nums []int) int {
	idxToRemove := []int{}

	for i, num := range nums {
		if i+1 < len(nums) {
			if num == nums[i+1] {
				idxToRemove = append(idxToRemove, i)
			}
		} else {
			break
		}
	}

	return len(nums) - len(idxToRemove)
}
