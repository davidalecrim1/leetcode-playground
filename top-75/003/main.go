package main

func containsDuplicate(nums []int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		_, ok := seen[num]

		if ok {
			return true
		}

		seen[num] = i
	}

	return false
}
