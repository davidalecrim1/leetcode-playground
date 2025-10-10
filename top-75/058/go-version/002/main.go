package main

// Time: O(Nˆ3)
// Because there is N * N substrings and isPalindrome is N.
// O(Nˆ2 * N) = O(Nˆ3)
func countSubstrings(s string) int {
	// Stragegy: DFS + memo
	// function to see if its a palindrome
	// dfs(start int)
	// increase count
	// end the dfs after going over the len
	// start at each index
	// think about a memo

	count := 0
	seen := map[[2]int]bool{}

	var dfs func(start int, end int)
	dfs = func(start int, end int) {
		if start > len(s) || end > len(s) || start >= end {
			return
		}

		prefix := s[start:end]
		//fmt.Println(prefix)
		if seen[[2]int{start, end}] {
			return
		}

		if isPalindrome(prefix) {
			count++
		}

		seen[[2]int{start, end}] = true

		dfs(start, end+1)
		dfs(start+1, end+1)
	}

	dfs(0, 1)

	return count
}

func isPalindrome(s string) bool {
	// if len(substring) == 1 {
	//     return true
	// }

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
