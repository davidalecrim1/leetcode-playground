package main

import "strconv"

// add numbers to the stack
// on symbol, pop two from the stack, calculate and append
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for i := range tokens {
		if isOperator(tokens[i]) {
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			//fmt.Printf("calculating %v %v %v\n", num1, num2, tokens[i])
			res := calculate(num1, num2, tokens[i])
			//fmt.Printf("adding res to stack %v\n", res)
			stack = append(stack, res)
		} else {
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
			//fmt.Printf("adding val to stack %v\n", val)
		}
	}

	return stack[0]
}

func isOperator(charStr string) bool {
	if len(charStr) != 1 {
		return false
	}

	char := byte(charStr[0])
	if char == '/' || char == '+' || char == '-' || char == '*' {
		return true
	}

	return false
}

func calculate(num1 int, num2 int, opStr string) int {
	op := byte(opStr[0])
	switch op {
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	case '-':
		return num1 - num2
	case '+':
		return num1 + num2
	default:
		panic("invalid operation")
	}
}
