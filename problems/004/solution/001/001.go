package main

import "strings"

var brackets = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func IsValid(s string) bool {
	isValid := true

	for i := 0; i < len(s); i++ {
		closingBracket := brackets[s[i]]

		// is not a closing bracket
		if closingBracket == 0 {
			continue
		}

		if !strings.Contains(s, string(closingBracket)) {
			isValid = false
			break
		}

	}

	return isValid
}
