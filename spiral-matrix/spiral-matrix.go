package spiralmatrix

func SpiralMatrix(size int) [][]int {
	if size == 0 {
		return [][]int{}
	}

	// Initialize an empty matrix
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// Define boundaries
	top, bottom := 0, size-1
	left, right := 0, size-1
	num := 1

	for top <= bottom && left <= right {
		// Fill top row
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// Fill right column
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Fill bottom row
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
		}

		// Fill left column
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}
