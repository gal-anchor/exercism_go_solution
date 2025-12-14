package luhn

func Valid(id string) bool {
	digits := make([]byte, 0, len(id))

	for i := 0; i < len(id); i++ {
		if id[i] == ' ' {
			continue
		}

		if !('0' <= id[i] && id[i] <= '9') {
			return false
		}

		digits = append(digits, byte((id[i] - '0')))
	}

	length := len(digits)

	if length <= 1 {
		return false
	}

	parity := (length - 2) % 2

	var product int

	for i := 0; i < length; i++ {
		num := int(digits[i])

		if (i % 2) == parity {
			num *= 2

			if num > 9 {
				num -= 9
			}
		}

		product += num
	}

	return (product % 10) == 0
}
