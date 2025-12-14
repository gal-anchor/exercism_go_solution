package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidNumber = errors.New("invalid phone number")

// Number returns the cleaned 10-digit phone number
func Number(phoneNumber string) (string, error) {
	// Remove all non-digit characters
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(phoneNumber, "")

	// Handle 11-digit numbers starting with 1
	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	}

	if len(digits) != 10 {
		return "", ErrInvalidNumber
	}

	// Validate area code and exchange code
	if digits[0] < '2' || digits[3] < '2' {
		return "", ErrInvalidNumber
	}

	return digits, nil
}

// AreaCode returns the area code from a phone number
func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format returns the phone number in (XXX) XXX-XXXX format
func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}
