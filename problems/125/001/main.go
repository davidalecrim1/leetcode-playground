package main

import "strings"

// Skip non alphanumerical characters (symbols and space)
// Make case insensitive
// Use two pointers in the string, one at the left and one at the right

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if !isLowChar(s[l]) {
			l++
			continue
		}

		if !isLowChar(s[r]) {
			r--
			continue
		}

		if strings.EqualFold(string(s[l]), string(s[r])) {
			l++
			r--
			continue
		} else {
			// not a palindrome
			return false
		}
	}

	return true
}

func isLowChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
