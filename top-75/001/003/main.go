package main

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		// see if the diff is saved in the hashmap
		if j, ok := hm[diff]; ok {
			return []int{i, j}
		}

		// save in the hashmap if the value is not found to later use.
		hm[num] = i
	}

	return []int{}
}
