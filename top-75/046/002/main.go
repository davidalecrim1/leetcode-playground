package main

// Time Complexity: O(m * n)
// Space Complexity: O(m + n)
func setZeroes(matrix [][]int) {
	zeroRows := make(map[int]bool)
	zeroCols := make(map[int]bool)

	for r := range len(matrix) {
		for c := range len(matrix[0]) {
			if matrix[r][c] == 0 {
				zeroRows[r] = true
				zeroCols[c] = true
			}
		}
	}

	for r := range zeroRows {
		for c := range len(matrix[0]) {
			matrix[r][c] = 0
		}
	}

	for c := range zeroCols {
		for r := range len(matrix) {
			matrix[r][c] = 0
		}
	}
}
