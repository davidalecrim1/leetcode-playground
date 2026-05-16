package main

// Create a function to validate the sub board
// Use a bitfield map to see if the number already exists
// If so, return false
// Else add and continue

func isValidSudoku(board [][]byte) bool {
	subBoardSeen := make(map[byte]bool, 9)
	rowSeen := make(map[int]map[byte]bool, 9)
	colSeen := make(map[int]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowSeen[i] = make(map[byte]bool)
		colSeen[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			if !isValidSubBoard(r, c, board, subBoardSeen, rowSeen, colSeen) {
				return false
			}
		}
	}

	return true
}

func isValidSubBoard(row, col int, board [][]byte, subBoardSeen map[byte]bool, rowSeen,
	colSeen map[int]map[byte]bool) bool {
	clear(subBoardSeen)

	for r := row; r < row+3; r++ {
		for c := col; c < col+3; c++ {
			num := board[r][c]
			if num == '.' {
				continue
			}

			if subBoardSeen[num] {
				return false
			}
			subBoardSeen[num] = true

			if rowSeen[r][num] {
				return false
			}
			rowSeen[r][num] = true

			if colSeen[c][num] {
				return false
			}
			colSeen[c][num] = true
		}
	}

	return true
}
