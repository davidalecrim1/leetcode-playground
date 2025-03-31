package main

func numDecodings(s string) int {
	// Approach: Make a decision (from a few options), and than make another decision.
	// This will look like a tree.
	// Therefore here we can apply a DFS strategy in that decision tree.
	// After a successful DFS, add memoization to store the res of subproblems.
	memo := make(map[int]int, len(s))
	return backtrack(0, s, memo)
}

func backtrack(i int, s string, memo map[int]int) int {
	if i == len(s) {
		return 1
	}

	if s[i] == '0' {
		return 0
	}

	if val, ok := memo[i]; ok {
		return val
	}

	res := backtrack(i+1, s, memo)

	if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] <= '6')) {
		res += backtrack(i+2, s, memo)
	}

	memo[i] = res
	return res
}
