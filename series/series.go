package series

// All returns all consecutive substrings of length n.
func All(n int, s string) []string {
	if n > len(s) {
		return []string{}
	}

	result := make([]string, 0, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

// UnsafeFirst returns the first substring of length n.
// Panics if n is larger than the string length.
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		panic("substring length is greater than string length")
	}
	return s[0:n]
}
