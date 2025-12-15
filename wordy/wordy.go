package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	// Remove trailing "?" and leading "What is "
	question = strings.TrimSuffix(question, "?")
	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSpace(question)
	if question == "" {
		return 0, false
	}

	tokens := strings.Fields(question)
	if len(tokens) == 1 {
		// Single number case: "What is 5?"
		num, err := strconv.Atoi(tokens[0])
		if err != nil {
			return 0, false
		}
		return num, true
	}

	if len(tokens) < 3 {
		return 0, false
	}

	// Convert first number
	num, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	i := 1
	for i < len(tokens) {
		if i+1 >= len(tokens) {
			return 0, false // incomplete operation
		}

		op := tokens[i]
		var nextNum int

		// Handle two-word operators like "multiplied by" or "divided by"
		if op == "plus" || op == "minus" {
			nextNum, err = strconv.Atoi(tokens[i+1])
			if err != nil {
				return 0, false
			}
		} else if op == "multiplied" && tokens[i+1] == "by" {
			if i+2 >= len(tokens) {
				return 0, false
			}
			nextNum, err = strconv.Atoi(tokens[i+2])
			if err != nil {
				return 0, false
			}
			op = "multiplied by"
			i++ // skip "by"
		} else if op == "divided" && tokens[i+1] == "by" {
			if i+2 >= len(tokens) {
				return 0, false
			}
			nextNum, err = strconv.Atoi(tokens[i+2])
			if err != nil {
				return 0, false
			}
			op = "divided by"
			i++ // skip "by"
		} else {
			return 0, false
		}

		// Apply operation
		switch op {
		case "plus":
			num += nextNum
		case "minus":
			num -= nextNum
		case "multiplied by":
			num *= nextNum
		case "divided by":
			num /= nextNum
		}

		i += 2
	}

	return num, true
}
