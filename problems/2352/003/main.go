package main

import "fmt"

func equalPairs(grid [][]int) int {
	// Walk every column
	// Store the arrays in a hashmap O(1) lookup instead of O(n) if its an array
	// Walk every row and see if they exist in the hashmap, update the count

	n := len(grid)
	hm := make(map[string]int)
	values := make([]int, n)

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			values[i] = grid[i][j]
		}

		hm[fmt.Sprintf("%v", values)] += 1
	}

	//fmt.Println(hm)

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			values[j] = grid[i][j]
		}

		if val, ok := hm[fmt.Sprintf("%v", values)]; ok {
			res += val
		}
	}

	//fmt.Println(hm)

	return res
}
