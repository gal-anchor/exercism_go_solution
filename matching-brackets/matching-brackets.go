package brackets

func Bracket(input string) bool {
	stack := []rune{}

	// mapping of closing brackets to opening brackets
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range input {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char) // push opening bracket
		case ')', ']', '}':
			if len(stack) == 0 {
				return false // no matching opening bracket
			}
			top := stack[len(stack)-1]
			if top != pairs[char] {
				return false // mismatch
			}
			stack = stack[:len(stack)-1] // pop
		}
	}

	return len(stack) == 0 // balanced if stack is empty
}
