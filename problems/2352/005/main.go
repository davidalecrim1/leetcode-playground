package main

// I like this solution a lot
// Because it uses the restriction of the grid being 200x200
// So instead of using fmt.Sprint, we can use the array which is way more efficient

// Time Complexity: O(nˆ2)
// Space Complexity: O(n)
func equalPairs(grid [][]int) int {
	// Walk every column
	// Store the arrays in a hashmap O(1) lookup instead of O(n) if its an array
	// Walk every row and see if they exist in the hashmap, update the count

	n := len(grid)
	hm := make(map[[200]int]int)
	var values [200]int

	for j := 0; j < n; j++ {
		values = [200]int{} // reset
		for i := 0; i < n; i++ {
			values[i] = grid[i][j]
		}

		hm[values] += 1

	}

	//fmt.Println(hm)

	var res int
	for i := 0; i < n; i++ {
		values = [200]int{} // reset
		for j := 0; j < n; j++ {
			values[j] = grid[i][j]
		}

		if val, ok := hm[values]; ok {
			res += val
		}
	}

	//fmt.Println(hm)

	return res
}
