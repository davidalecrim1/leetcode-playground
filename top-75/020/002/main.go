package main

import "fmt"

/*
Approach 2: Apply Bottom Up Dynamic Programming

Time: O(n^2)
Space: O(n)
*/

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)

	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true // This considers the first DP as true

	for i := 1; i <= len(s); i++ {
		// if `i` is in 4, and the result is leet. the `j` will look from `l to t` substring.
		for j := 0; j < i; j++ {
			prefix := s[j:i]
			// the dp[j] makes sure the rest of the string, if `leet` was divided in `le` and `et`, that the dp[2] that representes `le`
			// is saved as true to make sure this prefix was found
			if dp[j] && wordMap[prefix] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict)) // true
}
