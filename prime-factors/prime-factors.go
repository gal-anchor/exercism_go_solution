package prime

// Factors returns the prime factors of n in ascending order.
func Factors(n int64) []int64 {
	var result []int64
	if n < 2 {
		return result
	}

	// Factor out 2s first
	for n%2 == 0 {
		result = append(result, 2)
		n /= 2
	}

	// Factor out odd numbers starting from 3
	var i int64 = 3
	for i*i <= n {
		for n%i == 0 {
			result = append(result, i)
			n /= i
		}
		i += 2
	}

	// If n is still greater than 1, it is prime
	if n > 1 {
		result = append(result, n)
	}

	return result
}
