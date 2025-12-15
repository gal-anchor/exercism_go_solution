package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	var b strings.Builder
	count := 1

	for i := 1; i <= len(input); i++ {
		if i < len(input) && input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				b.WriteString(strconv.Itoa(count))
			}
			b.WriteByte(input[i-1])
			count = 1
		}
	}

	return b.String()
}

func RunLengthDecode(input string) string {
	var b strings.Builder
	count := 0

	for _, r := range input {
		if unicode.IsDigit(r) {
			digit, _ := strconv.Atoi(string(r))
			count = count*10 + digit
		} else {
			if count == 0 {
				count = 1
			}
			for i := 0; i < count; i++ {
				b.WriteRune(r)
			}
			count = 0
		}
	}

	return b.String()
}
