package foodchain

import "strings"

var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var comments = map[string]string{
	"spider": "It wriggled and jiggled and tickled inside her.\n",
	"bird":   "How absurd to swallow a bird!\n",
	"cat":    "Imagine that, to swallow a cat!\n",
	"dog":    "What a hog, to swallow a dog!\n",
	"goat":   "Just opened her throat and swallowed a goat!\n",
	"cow":    "I don't know how she swallowed a cow!\n",
}

func Verse(v int) string {
	var b strings.Builder
	animal := animals[v-1]

	b.WriteString("I know an old lady who swallowed a ")
	b.WriteString(animal)
	b.WriteString(".\n")

	if animal == "horse" {
		b.WriteString("She's dead, of course!")
		return b.String()
	}

	if comment, ok := comments[animal]; ok {
		b.WriteString(comment)
	}

	for i := v - 1; i > 0; i-- {
		b.WriteString("She swallowed the ")
		b.WriteString(animals[i])
		b.WriteString(" to catch the ")
		b.WriteString(animals[i-1])

		if animals[i-1] == "spider" {
			b.WriteString(" that wriggled and jiggled and tickled inside her")
		}
		b.WriteString(".\n")
	}

	b.WriteString("I don't know why she swallowed the fly. Perhaps she'll die.")
	return b.String()
}

func Verses(start, end int) string {
	var b strings.Builder

	for i := start; i <= end; i++ {
		b.WriteString(Verse(i))
		if i != end {
			b.WriteString("\n\n")
		}
	}

	return b.String()
}

func Song() string {
	return Verses(1, len(animals))
}
