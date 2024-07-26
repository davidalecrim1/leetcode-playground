package main

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(current string, open, close, max int, result *[]string) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(current+"(", open+1, close, max, result)
	}
	if close < open {
		backtrack(current+")", open, close+1, max, result)
	}
}
