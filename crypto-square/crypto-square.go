package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	// Step 1: Normalize
	var normalized []rune
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normalized = append(normalized, unicode.ToLower(r))
		}
	}

	n := len(normalized)
	if n == 0 {
		return ""
	}

	// Step 2: Determine columns and rows
	cols := int(math.Ceil(math.Sqrt(float64(n))))
	rows := int(math.Ceil(float64(n) / float64(cols)))

	// Step 3: Build grid
	grid := make([]string, rows)
	for i := 0; i < n; i++ {
		row := i / cols
		grid[row] += string(normalized[i])
	}

	// Step 4: Read columns vertically with padding
	var result []string
	for c := 0; c < cols; c++ {
		var word strings.Builder
		for r := 0; r < rows; r++ {
			if c < len(grid[r]) {
				word.WriteByte(grid[r][c])
			} else {
				word.WriteByte(' ') // pad shorter columns
			}
		}
		result = append(result, word.String())
	}

	// Step 5: Join with space
	return strings.Join(result, " ")
}
