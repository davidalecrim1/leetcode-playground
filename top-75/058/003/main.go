package main

func countSubstrings(s string) int {
	// Strategy
	// Check every odd and even pairs if they are a palindrome
	// Remove the repeated work of DFS
	// Time: O(Nˆ2)

	res := 0

	for i := range s {
		l, r := i, i
		res += countPalindrome(l, r, s)

		l, r = i, i+1
		res += countPalindrome(l, r, s)
	}

	return res
}

func countPalindrome(l, r int, s string) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
