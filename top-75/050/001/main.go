package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Approach 1:
	// Some kind of greedy

	// Approach 2:
	// Try some kind of sliding window
	res := 0
	seen := map[rune]int{}
	curr := 0
	for i, char := range s {
		if _, ok := seen[char]; ok {
			// Start again
			curr = 1
			seen = map[rune]int{} // Confirm I can do this.
			seen[char] = i
			continue

			// Also start from the next of the last seen.
		}

		curr++
		res = max(res, curr)
		seen[char] = i
	}

	return res
}

func main() {
	fmt.Printf("lengthOfLongestSubstring(\"abcabcbb\"): %v\n", lengthOfLongestSubstring("abcabcbb"))
}
