package main

func lengthOfLongestSubstring(s string) int {
	// Approach 1:
	// Brute force with every possible start
	// Time Complexity: O(nˆ2)
	// Space Complexity: O(n)

	res := 0
	for i := 0; i < len(s); i++ {
		res = max(res, lookup(s, i))
	}

	return res
}

func lookup(s string, startAt int) int {
	res := 0
	seen := map[byte]int{}
	for i := startAt; i < len(s); i++ {
		if _, ok := seen[s[i]]; ok {
			return res
		}
		res++
		seen[s[i]] = i
	}

	return res
}
