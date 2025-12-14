package isogram

import "strings"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, char := range word {
		lowerChar := strings.ToLower(string(char))
		r := []rune(lowerChar)[0]
		if r == ' ' || r == '-' {
			continue
		}
		if _, exists := seen[r]; exists {
			return false
		}

		seen[r] = true
	}
	return true
}
