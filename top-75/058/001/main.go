package main

func countSubstrings(s string) int {
	res := 0

	checkPalidrome := func(l, r int) {
		for (l >= 0 && r < len(s)) && s[l] == s[r] {
			res++

			l--
			r++
		}
	}

	for i := range len(s) {
		// odd
		checkPalidrome(i, i)

		// even
		checkPalidrome(i, i+1)
	}

	return res
}
