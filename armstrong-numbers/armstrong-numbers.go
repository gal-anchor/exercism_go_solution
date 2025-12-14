package armstrong

import "math"

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	num := n
	sum := 0
	digits := 0

	// Count the number of digits
	for temp := n; temp > 0; temp /= 10 {
		digits++
	}

	// Sum the digits raised to the power of 'digits'
	for temp := n; temp > 0; temp /= 10 {
		d := temp % 10
		sum += int(math.Pow(float64(d), float64(digits)))
	}

	return sum == num
}
