package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	// Strategy
	// Create two bool matrix to record if a cell is
	// reachable by the sea to the values.

	// Do the reverse DFS from the edges to the cells.
	// Given the memory optimization, it's going to be O(m * n)

	rows, cols := len(heights), len(heights[0])
	var dfs func(r, c int, ocean [][]bool)
	dfs = func(r, c int, ocean [][]bool) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}

		ocean[r][c] = true

		// right - left - up -- down
		for _, dir := range [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			nr, nc := r+dir[0], c+dir[1]

			if !(nr < 0) && !(nc < 0) && !(nr >= rows) && !(nc >= cols) {
				if heights[r][c] <= heights[nr][nc] && ocean[nr][nc] != true {
					dfs(nr, nc, ocean)
				}
			}
		}
	}

	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)

	for i := range rows {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	// first column and top row
	for r := range rows {
		pacific[r][0] = true
		dfs(r, 0, pacific)

		fmt.Println("done top row")
	}
	for c := range cols {
		pacific[0][c] = true
		dfs(0, c, pacific)

		fmt.Println("done left col")
	}

	// right column and bottom row
	for r := range rows {
		atlantic[r][cols-1] = true
		dfs(r, cols-1, atlantic)

		fmt.Println("done bottom row")
	}

	for c := range cols {
		atlantic[rows-1][c] = true
		dfs(rows-1, c, atlantic)

		fmt.Println("done right col")
	}

	res := [][]int{}
	for r := range rows {
		for c := range cols {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
