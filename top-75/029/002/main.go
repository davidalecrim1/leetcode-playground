package main

// Time Complexity: O(m * n)
func pacificAtlantic(heights [][]int) [][]int {
	// Strategy:
	// Apply bottom up DFS to store until where I can reach
	// Then look up to see what I've found.

	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := range rows {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	var dfs func(r, c int, ocean [][]bool)
	dfs = func(r, c int, ocean [][]bool) {
		if ocean[r][c] {
			return
		}

		ocean[r][c] = true // is reachable

		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]

			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue
			}

			if heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, ocean)
			}
		}
	}

	// top and left borders
	for r := range rows {
		dfs(r, 0, pacific)
	}
	for c := range cols {
		dfs(0, c, pacific)
	}

	// bottom and right borders
	for r := range rows {
		dfs(r, cols-1, atlantic)
	}
	for c := range cols {
		dfs(rows-1, c, atlantic)
	}

	var res [][]int
	for r := range rows {
		for c := range cols {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
