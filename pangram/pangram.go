package pangram

import "unicode"

func IsPangram(input string) bool {
	seen := make(map[rune]bool)

	for _, r := range input {
		if unicode.IsLetter(r) {
			seen[unicode.ToLower(r)] = true
		}
	}

	return len(seen) == 26
}
