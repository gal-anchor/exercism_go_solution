package bob

import (
	"strings"
	"unicode"
)

// Hey responds to a remark according to Bob's rules.
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)

	switch {
	case trimmed == "":
		return "Fine. Be that way!"
	case isYelling(trimmed) && strings.HasSuffix(trimmed, "?"):
		return "Calm down, I know what I'm doing!"
	case isYelling(trimmed):
		return "Whoa, chill out!"
	case strings.HasSuffix(trimmed, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}

// isYelling returns true if the remark is all uppercase letters (and contains letters)
func isYelling(s string) bool {
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			if unicode.IsLower(r) {
				return false
			}
		}
	}
	return hasLetter
}
