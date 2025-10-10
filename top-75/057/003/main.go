package main

func longestPalindrome(s string) string {
	// Strategy 2
	// To check if something is a palindrome
	// Start at the middle character and expand outwards
	// Check if it is odd or even to handle edge cases

	var res string
	resLen := 0

	var storePalindrome func(i, j int)
	storePalindrome = func(i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			if j-i+1 > resLen {
				resLen = max(resLen, j-i+1)
				res = s[i : j+1]
			}

			i--
			j++
		}
	}

	for i := range s {

		// odd check
		storePalindrome(i, i)

		// even check
		storePalindrome(i, i+1)
	}

	return res

}
