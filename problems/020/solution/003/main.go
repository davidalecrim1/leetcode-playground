package main

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	/*
	   Use a stack to keep track of the open ones.
	   Map the closing to the opening brackets with a hashmap.
	*/

	brackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := range s {
		if gotOpening, ok := brackets[s[i]]; ok {
			if len(stack) == 0 {
				return false
			}

			expectedClosing := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if expectedClosing != gotOpening {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
