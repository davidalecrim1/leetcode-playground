package main

func countSubstrings(s string) int {
	// Strategy
	// Do a DFS to try every possibility
	// Have a function to check if it is a palindrome
	// Start at every character and try to match it
	// If it is not a palindrome, drop it

	// Complexity: palindrome check (N)
	// DFS = N * N
	// O(Nˆ3)

	res := 0
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i == len(s) || j > len(s) {
			return
		}

		if isPalindrome(s[i:j]) {
			res++
		}

		dfs(i, j+1)
	}

	for i := range s {
		dfs(i, i+1)
	}

	return res
}

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
