package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Time Complexity: O(n²)
	// Space Complexity: O(n)
	output := 0

	i := 0
	for i < len(s) {
		colisionAt := lookup(s, i, &output)
		i = colisionAt + 1
	}

	return output
}

func lookup(s string, startAt int, output *int) int {
	res := 0
	seen := map[byte]int{}
	for i := startAt; i < len(s); i++ {
		if val, ok := seen[s[i]]; ok {
			return val
		}
		res++
		*output = max(res, *output)
		seen[s[i]] = i
	}

	return len(s)
}

func main() {
	fmt.Printf("lengthOfLongestSubstring(\" \"): %v\n", lengthOfLongestSubstring(" "))
}
