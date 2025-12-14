package cipher

import "unicode"

/* -------------------- Shift / Caesar -------------------- */

type shift struct {
	distance int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	return shiftString(input, c.distance)
}

func (c shift) Decode(input string) string {
	return shiftString(input, -c.distance)
}

/* -------------------- Vigenere -------------------- */

type vigenere struct {
	key []int
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	var shifts []int
	for _, r := range key {
		if !unicode.IsLower(r) {
			return nil
		}
		if r != 'a' {
			shifts = append(shifts, int(r-'a'))
		} else {
			shifts = append(shifts, 0)
		}
	}

	// key must not be all 'a'
	allZero := true
	for _, s := range shifts {
		if s != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return nil
	}

	return vigenere{key: shifts}
}

func (v vigenere) Encode(input string) string {
	return v.apply(input, true)
}

func (v vigenere) Decode(input string) string {
	return v.apply(input, false)
}

func (v vigenere) apply(input string, encode bool) string {
	var result []rune
	keyIndex := 0

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}

		r = unicode.ToLower(r)
		if r < 'a' || r > 'z' {
			continue
		}

		shift := v.key[keyIndex%len(v.key)]
		if !encode {
			shift = -shift
		}

		result = append(result, rotate(r, shift))
		keyIndex++
	}

	return string(result)
}

/* -------------------- Helpers -------------------- */

func shiftString(input string, distance int) string {
	var result []rune

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}

		r = unicode.ToLower(r)
		if r < 'a' || r > 'z' {
			continue
		}

		result = append(result, rotate(r, distance))
	}

	return string(result)
}

func rotate(r rune, distance int) rune {
	offset := int(r-'a') + distance
	offset = (offset%26 + 26) % 26
	return rune('a' + offset)
}
