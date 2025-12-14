package lsproduct

import "errors"

// LargestSeriesProduct returns the largest product of 'span' consecutive digits in 'digits'.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	// Validation checks
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return 0, errors.New("span must not be greater than the length of the digits string")
	}
	if len(digits) == 0 && span == 0 {
		return 1, nil // edge case: empty string with span 0
	}

	// Convert string digits to a slice of integers
	nums := make([]int64, len(digits))
	for i, d := range digits {
		if d < '0' || d > '9' {
			return 0, errors.New("digits string contains non-numeric characters")
		}
		nums[i] = int64(d - '0')
	}

	var maxProduct int64 = 0

	// Edge case: span of 0 always returns 1
	if span == 0 {
		return 1, nil
	}

	// Compute products for all consecutive slices of length 'span'
	for i := 0; i <= len(nums)-span; i++ {
		product := int64(1)
		for j := i; j < i+span; j++ {
			product *= nums[j]
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
