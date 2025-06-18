package main

func numDecodings(s string) int {
	memo := make(map[int]int)
	n := len(s)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 1
		}

		if s[i] == '0' {
			return 0
		}

		if val, ok := memo[i]; ok {
			return val
		}

		res := dfs(i + 1)
		if i+1 < n && (s[i] == '1' || (s[i] == '2' && s[i+1] <= '6')) {
			res += dfs(i + 2)
		}

		memo[i] = res
		return res
	}

	return dfs(0)
}
