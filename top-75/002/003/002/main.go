package main

func containsDuplicate(nums []int) bool {
	// Option 1:
	// sort the array -> O(n log n) complexity
	// compare i-1 with i
	// O(n log n * n)

	// Option 2:
	// Create a hashmap
	// lookup the value
	// Time Complexity: O(n)
	// Space Complexity O(n)

	hm := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, exists := hm[nums[i]]; exists {
			return true
		}

		hm[nums[i]] = i
	}

	return false
}
