package prime

import (
	"errors"
)

// Nth returns the nth prime number. An error is returned if n <= 0.
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input: n must be greater than 0")
	}

	count := 0
	num := 1

	for {
		num++
		if isPrime(num) {
			count++
		}
		if count == n {
			return num, nil
		}
	}
}

// isPrime checks if a number is prime.
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
