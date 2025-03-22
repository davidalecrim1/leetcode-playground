package main

// Approach 1: Brute Force
// Approach 2: Use a HashMap for Look Ups. Store the diff between numbers.
// Approach 3: Sort the nums and use two pointers

// Chosen: Hashmap for LookUps.
// Time Complexity: O(n^2) -- Can be optimized for O(n)
func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, num := range nums {
		hm[num] = i
	}

	for idx, num := range nums {
		diff := target - num

		if hmIdx, exists := hm[diff]; exists && idx != hmIdx {
			return []int{idx, hmIdx}
		}
	}

	return []int{}
}

func main() {
	twoSum([]int{3, 2, 4}, 6)
}
