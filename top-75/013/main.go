package main

import "fmt"

// Time Complexity: O(n log n)
// Strategy: Brute force using the binary parser to string and count the runes.
// func countBits(n int) []int {
// 	res := make([]int, n+1)

// 	for i := range n + 1 {
// 		bin := strconv.FormatInt(int64(i), 2) // O(log n)
// 		res[i] = strings.Count(bin, "1")
// 	}

// 	return res
// }

// Time Complexity: O(n)
// Strategy: Use bit manipulation with dynamic programming.
func countBits(n int) []int {
	dp := make([]int, n+1)

	offset := 1 // the most significant bit to apply dinamic programming every 4 numbers because of the pattern.
	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i // update the offset to represent the most significant bit on the right every 4 calculations.
		}

		// when the offset is updated, the dp[i] is 1.
		dp[i] = 1 + dp[i-offset]
	}

	return dp
}

// Int to Binary:
// 0000 -> 0 (start) - pre setted, not used
// 0001 -> 1 - start offset as 1
// 0010 -> 2 - offset is updated to two
// 0011 -> 3 (end)
// 0100 -> 4 (start again the repeting of the right side - i.e. two bits) - update the offset
// 0101 -> 5
// 0110 -> 6
// 0111 -> 7
// 1000 -> 8

// To remember how to parse numbers to binary: 13 in Binary (Base 2)
// 13 / 2 -> 6 (rest 1)
// 6 / 2 -> 3 (rest 0)
// 3 / 2 -> 1 (rest 1)
// 1 / 2 -> 0 (rest 1)
// Bottom Up -> 1101 (13)

func main() {
	fmt.Printf("countBits(5): %v\n", countBits(5))
}
