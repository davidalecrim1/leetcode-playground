package main

// Follow the rules
// 1. validate each row
// 2. validate each col
// 3. use some kind of sliding window to validate the 3x3

// Time: O(n)
// Space: O(n)
func isValidSudoku(board [][]byte) bool {
	// 1: Validate each row
	seen := make(map[byte]bool, len(board))
	for r := range board {
		for c := range board[r] {
			if !valid(board, seen, board[r][c]) {
				return false
			}
		}

		clear(seen)
	}

	// 2. Validate each col
	for c := range board {
		for r := range board {
			if !valid(board, seen, board[r][c]) {
				return false
			}
		}

		clear(seen)
	}

	// 3. Validate the 3x3
	for r := 0; r < len(board); r += 3 {
		for c := 0; c < len(board); c += 3 {
			if !validSubBox(board, seen, r, r+3, c, c+3) {
				return false
			}

			clear(seen)
		}
	}

	return true
}

func validSubBox(board [][]byte, seen map[byte]bool, fromR, toR, fromC, toC int) bool {
	for r := fromR; r < toR; r++ {
		for c := fromC; c < toC; c++ {
			if !valid(board, seen, board[r][c]) {
				return false
			}
		}
	}

	return true
}

func valid(board [][]byte, seen map[byte]bool, cell byte) bool {
	if cell != '.' && seen[cell] {
		return false
	} else {
		seen[cell] = true
	}

	return true
}
