package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	// If the number is not divisible by 3, 5, or 7
	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}
