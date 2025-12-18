package alphametics

import (
	"errors"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	parts := strings.FieldsFunc(puzzle, func(r rune) bool {
		return r == '+' || r == '=' || r == ' '
	})

	resultWord := parts[len(parts)-1]
	summands := parts[:len(parts)-1]

	// Pre-calculate letters for each column (right to left)
	maxLen := 0
	for _, w := range parts {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	// Leading letters cannot be zero
	isLeading := make(map[byte]bool)
	for _, w := range parts {
		if len(w) > 1 {
			isLeading[w[0]] = true
		}
	}

	mapping := make(map[byte]int)
	for i := range mapping {
		mapping[uint8(i)] = -1
	}
	usedDigits := [10]bool{}

	if backtrack(0, 0, 0, maxLen, summands, resultWord, isLeading, &usedDigits, mapping) {
		// Convert byte map to string map for return
		res := make(map[string]int)
		for k, v := range mapping {
			res[string(k)] = v
		}
		return res, nil
	}

	return nil, errors.New("no solution found")
}

func backtrack(col, row, carry, maxLen int, summands []string, result string,
	isLeading map[byte]bool, used *[10]bool, mapping map[byte]int) bool {

	// If we've processed all columns
	if col == maxLen {
		return carry == 0
	}

	// Process summands for the current column
	if row < len(summands) {
		word := summands[row]
		if col >= len(word) {
			return backtrack(col, row+1, carry, maxLen, summands, result, isLeading, used, mapping)
		}

		char := word[len(word)-1-col]
		if val, ok := mapping[char]; ok {
			return backtrack(col, row+1, carry+val, maxLen, summands, result, isLeading, used, mapping)
		}

		for d := 0; d <= 9; d++ {
			if used[d] || (d == 0 && isLeading[char]) {
				continue
			}
			used[d] = true
			mapping[char] = d
			if backtrack(col, row+1, carry+d, maxLen, summands, result, isLeading, used, mapping) {
				return true
			}
			delete(mapping, char)
			used[d] = false
		}
		return false
	}

	// All summands for this column processed; now check against result word
	if col >= len(result) {
		return false // Result word is too short
	}

	resChar := result[len(result)-1-col]
	targetDigit := carry % 10
	newCarry := carry / 10

	if val, ok := mapping[resChar]; ok {
		if val == targetDigit {
			return backtrack(col+1, 0, newCarry, maxLen, summands, result, isLeading, used, mapping)
		}
		return false
	}

	if !used[targetDigit] && !(targetDigit == 0 && isLeading[resChar]) {
		used[targetDigit] = true
		mapping[resChar] = targetDigit
		if backtrack(col+1, 0, newCarry, maxLen, summands, result, isLeading, used, mapping) {
			return true
		}
		delete(mapping, resChar)
		used[targetDigit] = false
	}

	return false
}
