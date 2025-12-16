package minesweeper

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
		return []string{}
	}

	rows := len(board)
	cols := len(board[0])
	result := make([]string, rows)

	// Directions for 8 neighboring cells
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for r := 0; r < rows; r++ {
		row := []rune(board[r])
		for c := 0; c < cols; c++ {
			if row[c] == '*' {
				continue
			}
			count := 0
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && board[nr][nc] == '*' {
					count++
				}
			}
			if count > 0 {
				row[c] = rune('0' + count)
			}
		}
		result[r] = string(row)
	}
	return result
}
