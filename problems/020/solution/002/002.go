package main

func IsValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// Stack to keep track of opening parentheses
	stack := []rune{}

	for _, r := range s {
		if _, isOpening := pairs[r]; isOpening {
			stack = append(stack, r)
		} else {
			// If it's a closing parenthesis, check for a match
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
				return false
			}

			// Pop the matched opening parenthesis from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, all parentheses were matched
	return len(stack) == 0
}
