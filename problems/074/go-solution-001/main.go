package main

// Check the first and last row, if target is between the beginning and the ending of the column
// If it is, split in the middle and keep going until is found.
// Basically a binary search on the rows then the columns.
func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bottom := len(matrix) - 1
	lastCol := len(matrix[0]) - 1

	// search the row
	for top < bottom {
		mid := top + ((bottom - top) / 2)
		if target == matrix[mid][lastCol] {
			return true
		}

		if target < matrix[mid][lastCol] {
			bottom = mid
		} else {
			top = mid + 1
		}
	}

	// search the column
	left := 0
	right := len(matrix[bottom])
	for left < right {
		mid := left + ((right - left) / 2)
		if matrix[bottom][mid] == target {
			return true
		}

		if target < matrix[bottom][mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return false
}
