package main

// Time: O(n * m)
// Strategy: Use Hashmap to trade off space for time speed.
func palindromePairs(words []string) [][]int {
	var res [][]int

	hm := make(map[string]int, len(words))
	for i, word := range words {
		hm[reverseString(word)] = i
	}

	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			prefix := word[:j] // Left part of the split
			suffix := word[j:] // Right part of the split

			// Case 1: Check if the prefix has a palindrome suffix
			if k, exists := hm[prefix]; exists && i != k && isPalindrome(suffix) {
				res = append(res, []int{i, k})
			}

			// Case 2: Check if the suffix has a palindrome prefix
			if j > 0 { // Avoid duplicate checks for empty prefix
				if k, exists := hm[suffix]; exists && i != k && isPalindrome(prefix) {
					res = append(res, []int{k, i})
				}
			}
		}
	}

	return res
}

// Time: O(m)
func isPalindrome(matched string) bool {
	left := 0
	right := len(matched) - 1

	for left < right {
		if matched[left] != matched[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func reverseString(word string) string {
	runes := []rune(word)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
