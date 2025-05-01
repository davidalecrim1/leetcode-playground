package main

import "fmt"

// Time Complexity: O(m² * n + m * n²)
// Space Complexity: O(m * n)
func setZeroes(matrix [][]int) {
	// Approach: Walk on every cell
	// Store the ones that have zero.
	// After walking the whole matrix, apply the changes where a zero was found.
	// Why? I want to avoid doing in place update while im walking every cell.
	// If this is not a good option, store the ones that were originally zero.

	rows := len(matrix)
	cols := len(matrix[0])

	zeros := make([][2]int, 0)

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				zeros = append(zeros, [2]int{r, c})
			}
		}
	}

	for _, zero := range zeros {
		r := zero[0]
		c := zero[1]

		for i := range cols {
			matrix[r][i] = 0
		}

		for j := range rows {
			matrix[j][c] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1},
	}

	setZeroes(matrix)
	fmt.Println(matrix)
}
