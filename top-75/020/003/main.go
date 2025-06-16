package main

func wordBreak(s string, wordDict []string) bool {
	// Strategy
	// Hashmap with a sliding window
	// Not working for all use cases.

	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	left := 0
	right := 0
	for ; right <= len(s); right++ {
		prefix := s[left:right]

		// fmt.Println("prefix", prefix)

		if ok := wordMap[prefix]; ok {
			left = right
		}
	}

	// fmt.Println(left)
	// fmt.Println(right)

	return left == right-1
}
