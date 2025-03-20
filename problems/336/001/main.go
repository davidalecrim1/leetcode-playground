package main

// Time: O(n^2 * m)
// Strategy: Brute Force with Two For Loops
func palindromePairs(words []string) [][]int {
	var res [][]int

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}

			matched := words[i] + words[j]

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
