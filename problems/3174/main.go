package main

func clearDigits(s string) string {
	stack := make([]byte, 0, len(s))
	input := []byte(s)

	for i := range input {
		if isDigit(input[i]) {
			if len(stack) < 1 {
				return string(input)
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, input[i])
		}
	}

	return string(stack)
}

func isDigit(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}

	return false
}
