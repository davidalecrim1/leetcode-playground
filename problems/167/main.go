package main

// Approach 1: Brute Force - 2 Loops - O(nˆ2)
// Approach 2: Hashmap to improve lookup - O(n) : Trade Off: Space O(n)
// Approach 3: Two Pointes and walk the array - O(n)
// Since it is sorted I can increase or decrease the sum based on the target
// Space Complexity: O(1)

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[right] + numbers[left]
		if target == sum {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			// decrease the sum
			right--
		} else {
			left++
		}
	}

	return []int{}
}
