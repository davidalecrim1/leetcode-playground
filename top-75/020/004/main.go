package main

func wordBreak(s string, wordDict []string) bool {
	// Strategy
	// Hashmap with a sliding window
	// Not working for all use cases.
	// With a suffix control.

	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	left := 0
	right := 0
	for ; right <= len(s); right++ {
		prefix := s[left:right]
		if ok := wordMap[prefix]; ok {
			left = right
		}

		// fmt.Println("prefix", prefix)
		suffix := s[left:]
		if ok := wordMap[suffix]; ok {
			return true
		}
	}

	// fmt.Println(left)
	// fmt.Println(right)

	return left == right-1
}
