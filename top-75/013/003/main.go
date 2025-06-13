package main

func countBits(n int) []int {
	dp := make([]int, n+1)

	dp[0] = 0
	offset := 1

	for i := 1; i <= n; i++ {
		// The offset is updated based on the power of two
		// to improve the repetition pattern in bit representation.
		if offset*2 == i {
			offset = i
		}

		dp[i] = 1 + dp[i-offset] // add 1 bit as count and used the pre calculated amount of other bits for numbers

	}

	return dp
}
