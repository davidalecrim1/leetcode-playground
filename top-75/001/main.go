package main

func twoSum(nums []int, target int) []int {
	return twoSum_OnePass(nums, target)
}

// Approach: One Pass Hash Table
// Time Complexity: O(n)
// Space Complexity: O(n)
func twoSum_OnePass(nums []int, target int) []int {
	hm := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if j, ok := hm[diff]; ok {
			return []int{j, i}
		} else {
			hm[num] = i
		}
	}

	return []int{}
}

// Approach: Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func twoSum_BruteFoce(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1+num2 == target && i != j {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
