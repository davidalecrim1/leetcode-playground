package main

func spiralOrder(matrix [][]int) []int {
	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix[0])

	res := []int{}
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[left][i])
		}
		top++

		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right--

		// We check if there are still rows to process
		// If we don't, we will end up walking the same row.
		if !(left < right && top < bottom) {
			break
		}

		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom--

		for i := bottom - 1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
