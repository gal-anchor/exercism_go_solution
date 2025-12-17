package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Entry struct represents input data type
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type localeFormats struct {
	dateFormat        string
	currencyPostfix   string
	colDate           string
	colDescription    string
	colValue          string
	negPrefix         rune
	negPostfix        rune
	thousandSeparator rune
	fracSeparator     rune
}

var locales = map[string]localeFormats{
	"en-US": {
		dateFormat:        "01/02/2006",
		colDate:           "Date",
		colDescription:    "Description",
		colValue:          "Change",
		currencyPostfix:   "",
		negPrefix:         '(',
		negPostfix:        ')',
		thousandSeparator: ',',
		fracSeparator:     '.',
	},
	"nl-NL": {
		dateFormat:        "02-01-2006",
		colDate:           "Datum",
		colDescription:    "Omschrijving",
		colValue:          "Verandering",
		currencyPostfix:   " ",
		negPrefix:         ' ',
		negPostfix:        '-',
		thousandSeparator: '.',
		fracSeparator:     ',',
	},
}

var currencies = map[string]rune{
	"USD": '$',
	"EUR": '€',
}

type alignment = int

const (
	alignLeft alignment = iota
	alignRight
)

type parsedEntry struct {
	date        time.Time
	description string
	change      int
}

const dateLayout = "2006-01-02"
const columnSeparator = " | "

func printValue(value int, currency rune, locale string) string {
	nPref := ' '
	nPost := ' '
	cPost := locales[locale].currencyPostfix
	fSep := locales[locale].fracSeparator
	isNegative := value < 0
	if isNegative {
		value = -value
		nPref = locales[locale].negPrefix
		nPost = locales[locale].negPostfix
	}
	cents := value % 100
	value /= 100
	parts := []string{}
	if value == 0 {
		parts = []string{"0"}
	}
	for value > 0 {
		part := value % 1000
		value /= 1000
		format := "%d"
		if value > 0 {
			format = "%03d"
		}
		parts = append([]string{fmt.Sprintf(format, part)}, parts...)
	}
	v := strings.Join(parts, string(locales[locale].thousandSeparator))
	return fmt.Sprintf("%c%c%s%s%c%02d%c", nPref, currency, cPost, v, fSep, cents, nPost)
}

func formatField(value string, width int, align alignment) string {
	runes := []rune(value)
	if len(runes) > width {
		return string(runes[:width-3]) + "..."
	}
	rep := strings.Repeat(" ", width-len(runes))
	switch align {
	case alignLeft:
		return value + rep
	case alignRight:
		return rep + value
	}
	return "" // should not run
}

// FormatLedger return pretty formatted ledger using given locale and currency
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	currSign, ok := currencies[currency]
	if !ok {
		return "", errors.New("unknown currency")
	}
	locData, ok := locales[locale]
	if !ok {
		return "", errors.New("unknown locale")
	}

	var parsedEntries []parsedEntry
	for _, e := range entries {
		d, err := time.Parse(dateLayout, e.Date)
		if err != nil {
			return "", err
		}
		parsedEntries = append(parsedEntries, parsedEntry{
			date:        d,
			description: e.Description,
			change:      e.Change,
		})
	}
	sort.Slice(parsedEntries, func(i, j int) bool {
		switch {
		case parsedEntries[i].date.Before(parsedEntries[j].date):
			return true
		case parsedEntries[j].date.Before(parsedEntries[i].date):
			return false
		}
		switch {
		case parsedEntries[i].description < parsedEntries[j].description:
			return true
		case parsedEntries[j].description < parsedEntries[i].description:
			return false
		}
		return parsedEntries[i].change < parsedEntries[j].change
	})

	var buf strings.Builder
	// Header
	buf.WriteString(fmt.Sprintf("%s%s%s%s%s\n", formatField(locData.colDate, 10, alignLeft),
		columnSeparator, formatField(locData.colDescription, 25, alignLeft),
		columnSeparator, locData.colValue))

	for _, e := range parsedEntries {
		buf.WriteString(fmt.Sprintf("%s%s%s%s%s\n",
			formatField(e.date.Format(locData.dateFormat), 10, alignLeft),
			columnSeparator, formatField(e.description, 25, alignLeft),
			columnSeparator, formatField(printValue(e.change, currSign, locale), 13, alignRight)))
	}

	return buf.String(), nil
}
