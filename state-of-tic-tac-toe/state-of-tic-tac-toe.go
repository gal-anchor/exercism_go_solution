package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	if len(board) != 3 {
		return "", errors.New("board must have 3 rows")
	}
	for _, row := range board {
		if len(row) != 3 {
			return "", errors.New("each row must have 3 columns")
		}
	}

	// Count X and O
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 'X' {
				xCount++
			} else if cell == 'O' {
				oCount++
			} else if cell != ' ' {
				return "", errors.New("invalid character on board")
			}
		}
	}

	// Validate turn counts
	if xCount != oCount && xCount != oCount+1 {
		return "", errors.New("invalid board: wrong number of Xs and Os")
	}

	// Check all possible win conditions
	lines := [8][3][2]int{
		// Rows
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		// Columns
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		// Diagonals
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	winX, winO := false, false
	for _, line := range lines {
		a, b, c := line[0], line[1], line[2]
		if board[a[0]][a[1]] != ' ' &&
			board[a[0]][a[1]] == board[b[0]][b[1]] &&
			board[a[0]][a[1]] == board[c[0]][c[1]] {
			if board[a[0]][a[1]] == 'X' {
				winX = true
			} else {
				winO = true
			}
		}
	}

	// Validate wins (both cannot win at the same time)
	if winX && winO {
		return "", errors.New("invalid board: both players cannot win")
	}

	if winX || winO {
		return Win, nil
	}

	// Check if there are empty cells
	for _, row := range board {
		for _, cell := range row {
			if cell == ' ' {
				return Ongoing, nil
			}
		}
	}

	return Draw, nil
}
