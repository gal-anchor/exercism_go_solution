package palindrome

import (
	"errors"
	"fmt"
)

// Product holds the value of the palindrome and its factor pairs.
type Product struct {
	Value          int
	Factorizations [][2]int // Renamed from Factors to match your test suite
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	var minP, maxP Product
	found := false

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPalindrome(p) {
				if !found {
					minP = Product{Value: p, Factorizations: [][2]int{{i, j}}}
					maxP = Product{Value: p, Factorizations: [][2]int{{i, j}}}
					found = true
				} else {
					// Handle Minimum
					if p < minP.Value {
						minP = Product{Value: p, Factorizations: [][2]int{{i, j}}}
					} else if p == minP.Value {
						minP.Factorizations = append(minP.Factorizations, [2]int{i, j})
					}

					// Handle Maximum
					if p > maxP.Value {
						maxP = Product{Value: p, Factorizations: [][2]int{{i, j}}}
					} else if p == maxP.Value {
						maxP.Factorizations = append(maxP.Factorizations, [2]int{i, j})
					}
				}
			}
		}
	}

	if !found {
		return Product{}, Product{}, fmt.Errorf("no palindromes products found in range [%d, %d]", fmin, fmax)
	}

	return minP, maxP, nil
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	reversed := 0
	temp := n
	for temp > 0 {
		reversed = reversed*10 + temp%10
		temp /= 10
	}
	return n == reversed
}
