package say

import (
	"fmt"
	"strings"
)

var ones = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	return strings.TrimSpace(recursiveSay(n)), true
}

func recursiveSay(n int64) string {
	switch {
	case n >= 1_000_000_000:
		return fmt.Sprintf("%s billion %s", recursiveSay(n/1_000_000_000), recursiveSay(n%1_000_000_000))
	case n >= 1_000_000:
		return fmt.Sprintf("%s million %s", recursiveSay(n/1_000_000), recursiveSay(n%1_000_000))
	case n >= 1_000:
		return fmt.Sprintf("%s thousand %s", recursiveSay(n/1_000), recursiveSay(n%1_000))
	case n >= 100:
		return fmt.Sprintf("%s hundred %s", recursiveSay(n/100), recursiveSay(n%100))
	case n >= 20:
		hyphen := ""
		if n%10 != 0 {
			hyphen = "-" + ones[n%10]
		}
		return tens[n/10] + hyphen
	case n > 0:
		return ones[n]
	default:
		return ""
	}
}
