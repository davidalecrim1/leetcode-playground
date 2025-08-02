package main

// Time: O(n)
// Space: O(1)
func characterReplacement(s string, k int) int {
	// Strategy
	// Use k as a tolerance
	// Apply a sliding window algorithm
	// When the window is invalid, move the left pointer.
	// Use a variable to keep track of the longest char
	// Use a hashmap to count the chars in place

	count := make(map[byte]int, len(s))
	left := 0
	lrcc := 0
	res := 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		lrcc = max(lrcc, count[s[right]])

		tolerance := lrcc + k
		for right-left+1 > tolerance {
			count[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
