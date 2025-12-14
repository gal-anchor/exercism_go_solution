// Package acronym provides a function to generate acronyms from strings.
// An acronym is formed by taking the first letter of each word in a string and capitalizing it.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns the acronym of a given string.
// It takes the first letter of each word (separated by spaces or hyphens)
// and returns the result in uppercase.
func Abbreviate(s string) string {
	var acronym []rune
	// Split the string by spaces and hyphens
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})

	for _, word := range words {
		for _, r := range word {
			if unicode.IsLetter(r) {
				acronym = append(acronym, unicode.ToUpper(r))
				break // only take the first letter of the word
			}
		}
	}

	return string(acronym)
}
