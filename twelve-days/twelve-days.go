package twelve

import "strings"

var gifts = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves,",
	"three French Hens,",
	"four Calling Birds,",
	"five Gold Rings,",
	"six Geese-a-Laying,",
	"seven Swans-a-Swimming,",
	"eight Maids-a-Milking,",
	"nine Ladies Dancing,",
	"ten Lords-a-Leaping,",
	"eleven Pipers Piping,",
	"twelve Drummers Drumming,",
}

var days = []string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("On the " + days[i-1] + " day of Christmas my true love gave to me: ")

	for j := i; j > 0; j-- {
		if j == 1 && i > 1 {
			sb.WriteString("and ")
		}
		sb.WriteString(gifts[j-1] + " ")
	}

	return strings.TrimSpace(sb.String())
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
