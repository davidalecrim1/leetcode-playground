package main

// Time: O(n)
// Space: O(n)
func lengthOfLongestSubstring(s string) int {
	/*
	   Strategy:

	   Using a hashmap to keep track of seen
	   If a new caracter has been seen, move the left pointer

	   sliding window
	   update the res on each interation
	*/

	left := 0
	seen := map[byte]int{}
	res := 0

	for right := 0; right < len(s); right++ {
		if val, ok := seen[s[right]]; ok {
			left = max(left, val+1)
		}

		res = max(res, right-left+1)
		seen[s[right]] = right
	}

	return res
}
