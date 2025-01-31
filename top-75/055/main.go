package main

// Time: O(n)
// Space: O(n)
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	var stack []rune

	hm := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		// means its an opening bracket
		if _, ok := hm[char]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			poped := stack[len(stack)-1] // pop last one

			if hm[poped] == char {
				stack = append(stack[:len(stack)-1]) // remove last value from the stack
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
