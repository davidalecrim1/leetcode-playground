package main

import "fmt"

// Time Complexity: O(m * n * 3^L) -> Close to O(n^2)
// Strategy: Use DFS with another seen matrix for each search. Add and remove current visited cell.
func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	seen := make([][]bool, len(board))

	for i := range board {
		seen[i] = make([]bool, len(board[0]))
	}

	for row := range rows {
		for col := range cols {
			if backtrack(row, col, board, seen, word) {
				return true
			}
		}
	}

	return false
}

type direction struct {
	row int
	col int
}

func backtrack(row int, col int, board [][]byte, seen [][]bool, suffix string) bool {
	if len(suffix) == 0 {
		return true
	}

	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
		return false
	}

	if seen[row][col] || board[row][col] != suffix[0] {
		return false
	}

	seen[row][col] = true

	directions := map[string]direction{
		"up": {
			row: 0,
			col: -1,
		},
		"down": {
			row: 0,
			col: 1,
		},
		"left": {
			row: -1,
			col: 0,
		},
		"right": {
			row: 1,
			col: 0,
		},
	}

	for _, dir := range directions {
		nextRow := row + dir.row
		nextCol := col + dir.col

		if backtrack(nextRow, nextCol, board, seen, suffix[1:]) {
			return true
		}
	}

	seen[row][col] = false
	return false
}

func main() {
	res := exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")

	fmt.Printf("res: %v\n", res)
}
