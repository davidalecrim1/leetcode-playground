package main

func lengthOfLongestSubstring(s string) int {
	// Strategy 1
	// Brute force this
	// Start at every character and check a substring
	// O(nˆ2)

	// Strategy 2
	// Sliding window
	// Hashmap to store the seen chars
	// Move the left pointer when found something that exists
	// Store the max

	// abcabcbb
	// l = 0, r = 0
	// Store each char in the hashmap
	// ab
	// abc
	// abca -> found in the hashmap
	// move the left to the val from the hashmap
	// always have the window valid
	// O(n)

	hm := make(map[byte]int, len(s))
	l := 0
	res := 0

	for r := range s {
		if val, ok := hm[s[r]]; ok {
			l = max(l, val+1)
		}

		res = max(res, r-l+1)
		hm[s[r]] = r
	}

	return res
}
