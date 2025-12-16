package queenattack

import (
	"errors"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	// Validate positions
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid position")
	}

	whiteCol := whitePosition[0]
	whiteRow := whitePosition[1]
	blackCol := blackPosition[0]
	blackRow := blackPosition[1]

	// Check if columns and rows are within 'a'-'h' and '1'-'8'
	if whiteCol < 'a' || whiteCol > 'h' || blackCol < 'a' || blackCol > 'h' ||
		whiteRow < '1' || whiteRow > '8' || blackRow < '1' || blackRow > '8' {
		return false, errors.New("invalid position")
	}

	// Check if queens are on the same square
	if whitePosition == blackPosition {
		return false, errors.New("queens cannot occupy the same square")
	}

	// Convert to 0-based indexes
	wCol := int(whiteCol - 'a')
	wRow := int(whiteRow - '1')
	bCol := int(blackCol - 'a')
	bRow := int(blackRow - '1')

	// Same row or same column
	if wCol == bCol || wRow == bRow {
		return true, nil
	}

	// Same diagonal
	if abs(wCol-bCol) == abs(wRow-bRow) {
		return true, nil
	}

	return false, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
