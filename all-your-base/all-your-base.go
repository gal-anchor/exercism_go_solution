package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	value := 0
	for _, d := range inputDigits {
		value = value*inputBase + d
	}

	if value == 0 {
		return []int{0}, nil
	}

	var result []int
	for value > 0 {
		remainder := value % outputBase
		result = append([]int{remainder}, result...)
		value /= outputBase
	}

	return result, nil
}
