package main

// Time Complexity: O(m * n)
// Space Complexity: O(1)
func rotate(matrix [][]int) {
	// Aproach: Transpose the matrix.
	// Than mirror it.

	// Transpose
	n := len(matrix)
	// Time Complexity: O(m * n)
	for i := range n {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reserve
	// Time Complexity: O((m * n)/2)
	for i := range n {
		for j := range n / 2 {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
