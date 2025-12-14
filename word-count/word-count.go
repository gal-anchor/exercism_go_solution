package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	output := Frequency{}
	var wordMatcher = regexp.MustCompile(`\b([a-zA-Z]+('+[a-zA-Z]+)*|[0-9]*)`)
	words := wordMatcher.FindAllString(input, -1)

	for _, word := range words {
		output[strings.ToLower(word)] += 1
	}
	return output
}
