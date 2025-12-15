package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a 2D integer matrix.
type Matrix struct {
	data [][]int
}

// New creates a new Matrix from a string.
func New(s string) (Matrix, error) {
	if strings.TrimSpace(s) == "" {
		return Matrix{}, errors.New("empty matrix")
	}

	rows := strings.Split(s, "\n")
	matrix := make([][]int, len(rows))

	var cols int
	for i, row := range rows {
		fields := strings.Fields(row)
		if i == 0 {
			cols = len(fields)
			if cols == 0 {
				return Matrix{}, errors.New("invalid matrix")
			}
		} else if len(fields) != cols {
			return Matrix{}, errors.New("uneven rows")
		}

		matrix[i] = make([]int, cols)
		for j, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				return Matrix{}, err
			}
			matrix[i][j] = n
		}
	}

	return Matrix{data: matrix}, nil
}

// Cols returns a copy of the matrix columns.
func (m Matrix) Cols() [][]int {
	if len(m.data) == 0 {
		return [][]int{}
	}

	rows, cols := len(m.data), len(m.data[0])
	result := make([][]int, cols)

	for c := 0; c < cols; c++ {
		result[c] = make([]int, rows)
		for r := 0; r < rows; r++ {
			result[c][r] = m.data[r][c]
		}
	}

	return result
}

// Rows returns a copy of the matrix rows.
func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m.data))
	for i := range m.data {
		result[i] = make([]int, len(m.data[i]))
		copy(result[i], m.data[i])
	}
	return result
}

// Set sets the value at (row, col).
// Returns false if indices are out of bounds.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m.data) || col >= len(m.data[0]) {
		return false
	}
	m.data[row][col] = val
	return true
}
