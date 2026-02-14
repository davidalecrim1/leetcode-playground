package main

// Use a stack
// Store in the stack each day that is unresolved
// Pop from the stack to compare with the actual
// If resolved, update the answer
func dailyTemperatures(temperatures []int) []int {
	t := temperatures
	n := len(t)
	stack := make([]int, 0, n)
	ans := make([]int, n)

	for i := range t {
		for len(stack) > 0 && t[i] > t[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			ans[j] = i - j
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}
