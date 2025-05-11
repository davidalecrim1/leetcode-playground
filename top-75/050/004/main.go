package main

import "fmt"

// Time Complexity: O(n)
// Space Complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	start := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		char := s[end]

		if lastIdx, ok := seen[char]; ok && lastIdx >= start {
			start = lastIdx + 1
		}

		seen[char] = end
		windowLen := end - start + 1 // I add +1 because the idx starts at zero.
		maxLen = max(maxLen, windowLen)
	}

	return maxLen
}

func main() {
	fmt.Printf("lengthOfLongestSubstring(\"abcabcbb\"): %v\n", lengthOfLongestSubstring("abcabcbb"))
}
