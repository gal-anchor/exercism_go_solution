package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA []rune

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, n := range d {
		switch n {
		case 'A', 'C', 'G', 'T':
			h[n]++
		default:
			return nil, fmt.Errorf("invalid nucleotide: %q", n)
		}
	}

	return h, nil
}
