package main

// Time Complexity: O(n^2
// Space Complexity: O(n)
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true

	for i := n - 1; i >= 0; i-- {
		for w := range wordSet {
			if i+len(w) <= n {
				prefix := s[i : i+len(w)]
				// fmt.Println("prefix", prefix)

				if prefix == w {
					dp[i] = dp[i+len(w)]
				}
			}

			// fmt.Println("dp", i, dp[i])

			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}
