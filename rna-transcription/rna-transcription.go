package strand

import "strings"

func ToRNA(dna string) string {
	var b strings.Builder
	b.Grow(len(dna))

	for _, r := range dna {
		switch r {
		case 'G':
			b.WriteByte('C')
		case 'C':
			b.WriteByte('G')
		case 'T':
			b.WriteByte('A')
		case 'A':
			b.WriteByte('U')
		default:
			// invalid nucleotide — return empty string (matches common exercise behavior)
			return ""
		}
	}

	return b.String()
}
