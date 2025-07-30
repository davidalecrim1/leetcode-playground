package main

func characterReplacement(s string, k int) int {
	hm := map[byte]int{}
	maxChar := 0
	res := 0
	l := 0

	for r := 0; r < len(s); r++ {
		char := s[r]
		hm[char]++

		maxChar = max(maxChar, hm[char])

		diff := r - l + 1
		if diff-maxChar > k {
			hm[s[l]]--
			l++
		}

		diff = r - l + 1
		res = max(res, diff)
	}

	return res
}
