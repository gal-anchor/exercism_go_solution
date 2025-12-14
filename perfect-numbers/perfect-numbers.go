package perfect

import "errors"

// Classification represents the type of number
type Classification string

const (
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
	ClassificationDeficient Classification = "deficient"
)

// ErrOnlyPositive is returned when n is not a positive integer
var ErrOnlyPositive = errors.New("only positive integers are allowed")

// Classify classifies a number as perfect, abundant, or deficient.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return "", ErrOnlyPositive
	}

	sum := int64(0)

	// Sum of proper divisors
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
