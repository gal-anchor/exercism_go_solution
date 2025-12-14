package reverse

func Reverse(input string) string {
	// Convert the string to a slice of runes to handle Unicode correctly
	runes := []rune(input)
	n := len(runes)

	// Reverse the slice of runes in place
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
