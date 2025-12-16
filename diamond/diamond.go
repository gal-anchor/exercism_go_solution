package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("input must be an uppercase letter A-Z")
	}

	n := int(char - 'A')
	lines := make([]string, 0, 2*n+1)

	for i := 0; i <= n; i++ {
		letter := string('A' + byte(i))
		if i == 0 {
			line := strings.Repeat(" ", n) + letter + strings.Repeat(" ", n)
			lines = append(lines, line)
		} else {
			midSpaces := strings.Repeat(" ", 2*i-1)
			line := strings.Repeat(" ", n-i) + letter + midSpaces + letter + strings.Repeat(" ", n-i)
			lines = append(lines, line)
		}
	}

	// Bottom half (mirror top half, excluding middle row)
	for i := n - 1; i >= 0; i-- {
		lines = append(lines, lines[i])
	}

	return strings.Join(lines, "\n"), nil
}
