package main

func rotate(matrix [][]int) {
	// Strategy
	// To do this in place
	// Transpose the image (flip rows with cols)
	// Mirror the image (rows)

	rows, cols := len(matrix), len(matrix[0])

	// Transpose
	// Walk every row
	// Start the col on the same ID+1 of the row
	for i := range rows {
		for j := i + 1; j < rows; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// fmt.Println(matrix)
	// [[1 4 7] [2 5 8] [3 6 9]]

	// Reserve
	// Walk until the half
	for i := range rows {
		for j := range cols / 2 {
			matrix[i][j], matrix[i][cols-1-j] = matrix[i][cols-1-j], matrix[i][j]
		}
	}

	// fmt.Println(matrix)
	// [[7,4,1],[8,5,2],[9,6,3]]
}
