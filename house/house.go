package house

import "strings"

var parts = []struct {
	subject string
	verb    string
}{
	{"the house that Jack built.", ""},
	{"the malt", "that lay in"},
	{"the rat", "that ate"},
	{"the cat", "that killed"},
	{"the dog", "that worried"},
	{"the cow with the crumpled horn", "that tossed"},
	{"the maiden all forlorn", "that milked"},
	{"the man all tattered and torn", "that kissed"},
	{"the priest all shaven and shorn", "that married"},
	{"the rooster that crowed in the morn", "that woke"},
	{"the farmer sowing his corn", "that kept"},
	{"the horse and the hound and the horn", "that belonged to"},
}

func Verse(v int) string {
	var b strings.Builder
	b.WriteString("This is ")

	for i := v - 1; i >= 0; i-- {
		b.WriteString(parts[i].subject)
		if i > 0 {
			b.WriteString("\n")
			b.WriteString(parts[i].verb)
			b.WriteString(" ")
		}
	}

	return b.String()
}

func Song() string {
	var verses []string
	for i := 1; i <= len(parts); i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}
