package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	mapping := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return r
	}

	trimmed := strings.Map(mapping, s)
	trimmed = strings.ToLower(trimmed)

	i := 0
	j := len(trimmed) - 1

	for i < j {
		if trimmed[i] != trimmed[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// Chain of Thought
// Remove all that is not letter or number
// Use two pointers to walk from the beggning and the end making sure its equal
