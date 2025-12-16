package wordsearch

import "errors"
import "fmt"

// Solve solves word search puzzles. Puzzles and words are assumed to be composed of only ASCII characters
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	solution := make(map[string][2][2]int, len(words))
	rows := len(puzzle)

	// Validate words
	for _, w := range words {
		if len(w) == 0 {
			return nil, errors.New("empty string found in word list")
		}
	}

	// Validate puzzle
	if len(puzzle) == 0 {
		return nil, errors.New("empty puzzle")
	}

	cols := len(puzzle[0])
	for i := 1; i < rows; i++ {
		if len(puzzle[i]) != cols {
			return nil, errors.New("puzzle is not rectangular")
		}
	}

	// Search for each word
WordCheck:
	for _, w := range words {
		// Loop over each index in the puzzle
		for i := 0; i < rows; i++ {
		PuzzleCell:
			for j := 0; j < cols; j++ {
				// Test current index
				if w[0] != puzzle[i][j] {
					continue
				} else if len(w) == 1 {
					// Add solution
					solution[w] = [2][2]int{{j, i}, {j, i}}
					continue WordCheck
				}

				// Test left-to-right
				if j+len(w) <= cols {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i][j+k] {
							goto RightToLeft
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j + len(w) - 1, i}}
					continue WordCheck
				}

				// Test right-to-left
			RightToLeft:
				if j-len(w) >= -1 {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i][j-k] {
							goto UpDown
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j - len(w) + 1, i}}
					continue WordCheck
				}

				// Test up-down
			UpDown:
				if i+len(w) <= rows {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i+k][j] {
							goto DownUp
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j, i + len(w) - 1}}
					continue WordCheck
				}

				// Test down-up
			DownUp:
				if i-len(w) >= -1 {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i-k][j] {
							goto DiagonalUpAndLeft
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j, i - len(w) + 1}}
					continue WordCheck
				}

			DiagonalUpAndLeft:
				// Test up-and-left
				if j-len(w) >= -1 && i-len(w) >= -1 {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i-k][j-k] {
							goto DiagonalUpAndRight
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j - len(w) + 1, i - len(w) + 1}}
					continue WordCheck
				}

			DiagonalUpAndRight:
				// Test up-and-right
				if j+len(w) <= cols && i-len(w) >= -1 {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i-k][j+k] {
							goto DiagonalDownAndLeft
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j + len(w) - 1, i - len(w) + 1}}
					continue WordCheck
				}

			DiagonalDownAndLeft:
				// Test down-and-left
				if j-len(w) >= -1 && i+len(w) <= rows {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i+k][j-k] {
							goto DiagonalDownAndRight
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j - len(w) + 1, i + len(w) - 1}}
					continue WordCheck
				}

			DiagonalDownAndRight:
				// Test down-and-right
				if j+len(w) <= cols && i+len(w) <= rows {
					for k := 1; k < len(w); k++ {
						if w[k] != puzzle[i+k][j+k] {
							continue PuzzleCell
						}
					}

					// Add solution
					solution[w] = [2][2]int{{j, i}, {j + len(w) - 1, i + len(w) - 1}}
					continue WordCheck
				}
			}
		}

		return solution, fmt.Errorf("could not find word in puzzle: %v", w)
	}

	return solution, nil
}
