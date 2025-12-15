package forth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Forth(input []string) ([]int, error) {
	stack := []int{}
	dict := map[string][]string{}

	// Flatten input
	tokens := []string{}
	for _, part := range input {
		tokens = append(tokens, strings.Fields(part)...)
	}

	// Execute a list of tokens
	var exec func([]string) error
	exec = func(tok []string) error {
		i := 0
		for i < len(tok) {
			token := strings.ToLower(tok[i])

			// Word definition
			if token == ":" {
				i++
				if i >= len(tok) {
					return errors.New("expected word name after ':'")
				}
				word := strings.ToLower(tok[i])

				// Cannot redefine any number (positive or negative)
				if _, err := strconv.Atoi(word); err == nil {
					return errors.New("cannot redefine numbers")
				}

				i++
				def := []string{}
				for i < len(tok) && tok[i] != ";" {
					t := strings.ToLower(tok[i])
					// Snapshot current definition
					if d, ok := dict[t]; ok {
						def = append(def, d...)
					} else {
						def = append(def, t)
					}
					i++
				}
				if i >= len(tok) || tok[i] != ";" {
					return errors.New("expected ';' to end definition")
				}
				dict[word] = def
				i++
				continue
			}

			// Execute user-defined word
			if def, ok := dict[token]; ok {
				if err := exec(def); err != nil {
					return err
				}
				i++
				continue
			}

			// Built-in operations
			switch token {
			case "+":
				if len(stack) < 2 {
					return errors.New("stack underflow")
				}
				stack[len(stack)-2] += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			case "-":
				if len(stack) < 2 {
					return errors.New("stack underflow")
				}
				stack[len(stack)-2] -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			case "*":
				if len(stack) < 2 {
					return errors.New("stack underflow")
				}
				stack[len(stack)-2] *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			case "/":
				if len(stack) < 2 {
					return errors.New("stack underflow")
				}
				if stack[len(stack)-1] == 0 {
					return errors.New("division by zero")
				}
				stack[len(stack)-2] /= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			case "dup":
				if len(stack) < 1 {
					return errors.New("stack underflow")
				}
				stack = append(stack, stack[len(stack)-1])
			case "drop":
				if len(stack) < 1 {
					return errors.New("stack underflow")
				}
				stack = stack[:len(stack)-1]
			case "swap":
				if len(stack) < 2 {
					return errors.New("stack underflow")
				}
				stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
			case "over":
				if len(stack) < 2 {
					return errors.New("stack underflow")
				}
				stack = append(stack, stack[len(stack)-2])
			default:
				num, err := strconv.Atoi(token)
				if err != nil {
					return fmt.Errorf("unknown word: %s", token)
				}
				stack = append(stack, num)
			}
			i++
		}
		return nil
	}

	if err := exec(tokens); err != nil {
		return nil, err
	}
	return stack, nil
}
