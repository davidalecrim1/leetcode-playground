package main

func longestPalindrome(s string) string {
	res := ""
	resLen := 0

	storePalindrome := func(l, r int) {
		for (l >= 0 && r < len(s)) && s[l] == s[r] {
			if (r - l + 1) > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			l -= 1
			r += 1
		}
	}

	for i := range len(s) {
		// odd check
		l, r := i, i
		storePalindrome(l, r)

		// even check
		l, r = i, i+1
		storePalindrome(l, r)
	}

	return res
}
