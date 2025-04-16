package main

// Time Complexity: O(m * n)
func pacificAtlantic(heights [][]int) [][]int {
	// Aproach: DFS with optimized visited notes.
	// Also, start bottom up on each side.

	// If this was a pure DFS, it would be O(m * n)^2.
	// But for it to stop, one needs to keep track of the visited nodes.
	rows := len(heights)
	cols := len(heights[0])

	atVisited := make([][]bool, rows)
	pacVisited := make([][]bool, rows)
	for i := range rows {
		atVisited[i] = make([]bool, cols)
		pacVisited[i] = make([]bool, cols)

	}

	var dfs func(r, c int, visited [][]bool, prevHeight int)
	dfs = func(r, c int, visited [][]bool, prevHeight int) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}
		if visited[r][c] || heights[r][c] < prevHeight {
			return
		}

		visited[r][c] = true
		height := heights[r][c]

		dfs(r+1, c, visited, height) // down
		dfs(r, c+1, visited, height) // right
		dfs(r-1, c, visited, height) // up
		dfs(r, c-1, visited, height) // left
	}

	for c := range cols {
		// walk the first row (pacific)
		dfs(0, c, pacVisited, heights[0][c])

		// walk the last row (atlantic)
		dfs(rows-1, c, atVisited, heights[rows-1][c])
	}

	for r := range rows {
		// walk the first column (pacific)
		dfs(r, 0, pacVisited, heights[r][0])

		// walk the last column (atlantic)
		dfs(r, cols-1, atVisited, heights[r][cols-1])
	}

	res := [][]int{}
	for r := range rows {
		for c := range cols {
			if pacVisited[r][c] && atVisited[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
