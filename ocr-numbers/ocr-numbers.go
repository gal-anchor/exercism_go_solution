package ocr

import (
	"strings"
)

const (
	pipe       = iota
	underscore = iota
	space      = iota
	junk       = iota
	last       = iota
)

var zero = " _ | ||_|   "
var one = "     |  |   "
var two = " _  _||_    "
var three = " _  _| _|   "
var four = "   |_|  |   "
var five = " _ |_  _|   "
var six = " _ |_ |_|   "
var seven = " _   |  |   "
var eight = " _ |_||_|   "
var nine = " _ |_| _|   "

var digitMap = make(map[int64]string)

func init() {
	digitMap[generateInitialHash(zero)] = "0"
	digitMap[generateInitialHash(one)] = "1"
	digitMap[generateInitialHash(two)] = "2"
	digitMap[generateInitialHash(three)] = "3"
	digitMap[generateInitialHash(four)] = "4"
	digitMap[generateInitialHash(five)] = "5"
	digitMap[generateInitialHash(six)] = "6"
	digitMap[generateInitialHash(seven)] = "7"
	digitMap[generateInitialHash(eight)] = "8"
	digitMap[generateInitialHash(nine)] = "9"
}

// Recognize attempts to recognize OCR numeric data
func Recognize(input string) []string {
	lines := strings.Split(input, "\n")

	jj := 0
	res := make([]string, ((len(lines)-1)+3)/4)
	for ii := 1; ii < len(lines); ii += 4 {
		res[jj] = recognizeLines(lines[ii : ii+4])
		jj++
	}

	return res
}

func recognizeLines(input []string) string {
	for len(input) < 4 {
		input = append(input, "")
	}

	max := 0
	for _, v := range input {
		if len(v) > max {
			max = len(v)
		}
	}

	for max%3 != 0 {
		max++
	}

	for k := range input {
		for len(input[k]) < max {
			input[k] += " "
		}
	}

	res := ""
	lines := make([]string, 4)
	for ii := 0; ii < max; ii += 3 {
		lines[0] = input[0][ii : ii+3]
		lines[1] = input[1][ii : ii+3]
		lines[2] = input[2][ii : ii+3]
		lines[3] = input[3][ii : ii+3]
		res += recognizeDigit(lines)
	}

	return res
}

func recognizeDigit(lines []string) string {
	hash := int64(0)
	for _, line := range lines {
		for _, v := range line {
			hash *= int64(last)
			hash += recognizePart(v)
		}
	}

	res, ok := digitMap[hash]
	if !ok {
		return "?"
	}

	return res
}

func recognizePart(ch rune) int64 {
	switch ch {
	case '|':
		return pipe
	case '_':
		return underscore
	case ' ':
		return space
	default:
		return junk
	}
}

func generateInitialHash(input string) int64 {
	hash := int64(0)
	for _, v := range input {
		hash *= int64(last)
		hash += recognizePart(v)
	}

	return hash
}
