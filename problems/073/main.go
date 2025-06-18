package main

import (
	"fmt"
	"math"
)

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for r := range rows {
		for c := range cols {
			fmt.Println("rc", r, c, matrix[r][c])
			if matrix[r][c] == 0 {
				// rows
				for i := 0; i < rows; i++ {
					if matrix[i][c] != 0 {
						matrix[i][c] = math.MaxInt64
					}
				}

				// cols
				for j := 0; j < cols; j++ {
					if matrix[r][j] != 0 {
						matrix[r][j] = math.MaxInt64
					}
				}
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == math.MaxInt64 {
				matrix[r][c] = 0
			}
		}
	}
}
