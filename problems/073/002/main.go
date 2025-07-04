package main

func setZeroes(matrix [][]int) {
	// Strategy
	// Do it in place
	// Time: O(n * m)
	// O(1)
	firstRowZero := false
	firstColZero := false

	rows, cols := len(matrix), len(matrix[0])

	for r := range rows {
		if matrix[r][0] == 0 {
			firstRowZero = true
			break
		}
	}

	for c := range cols {
		if matrix[0][c] == 0 {
			firstColZero = true
			break
		}
	}

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	for r := 1; r < rows; r++ {
		if matrix[r][0] == 0 {
			for c := range cols {
				matrix[r][c] = 0
			}
		}
	}

	for c := range cols {
		if matrix[0][c] == 0 {
			for r := range rows {
				matrix[r][c] = 0
			}
		}
	}

	if firstRowZero {
		for c := range cols {
			matrix[0][c] = 0
		}
	}

	if firstColZero {
		for r := range rows {
			matrix[r][0] = 0
		}
	}
}
