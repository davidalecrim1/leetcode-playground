package main

import "strings"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}

		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		if strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}
