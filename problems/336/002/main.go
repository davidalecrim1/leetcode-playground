package main

// Time: O(n^2 * m)
// Strategy: Use Hashmap to trade off space for time speed.
func palindromePairs(words []string) [][]int {
	var res [][]int

	hm := make(map[int]string, len(words))
	for i, word := range words {
		hm[i] = word
	}

	for i := 0; i < len(hm); i++ {
		for j := 0; j < len(hm); j++ {
			if i == j {
				continue
			}

			// matched := strings.Join([]string{hm[i], hm[j]}, "")
			matched := hm[i] + hm[j]

			if isPalindrome(matched) {
				res = append(res, []int{i, j})
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
