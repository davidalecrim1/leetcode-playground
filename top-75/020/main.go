package main

import "fmt"

/*
Approach1: Perform a Brute Force with all possible combinations.
Time: O(2^N)
Space: O(N)

Start with a single chat (l) and keep appending until a word is found.
if true, I keep trying the rest until both is true to return true.

Otherwise, I return false.
*/

func workBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})

	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}

	return backtrack(s, wordMap)
}

func backtrack(s string, wordMap map[string]struct{}) bool {
	if s == "" {
		return true
	}

	for i := 1; i <= len(s); i++ {
		prefix := s[:i]

		if _, exists := wordMap[prefix]; exists && backtrack(s[i:], wordMap) {
			return true
		}
	}

	return false
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(workBreak(s, wordDict)) // true
}
