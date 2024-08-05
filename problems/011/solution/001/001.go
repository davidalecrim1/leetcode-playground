package main

func searchInsert(nums []int, target int) int {
	if target >= nums[len(nums)-1] {
		return len(nums)
	}

	if target <= nums[0] {
		return 0
	}

	for i, num := range nums {
		// return existing index
		if num == target {
			return i
		}

		// tell where it should be if index is not found
		if target > num && target <= nums[i+1] {
			return i + 1
		}
	}

	return -1
}
