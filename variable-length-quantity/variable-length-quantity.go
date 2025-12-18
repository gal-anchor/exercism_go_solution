package variablelengthquantity

import "errors"

// EncodeVarint encodes a slice of uint32 into VLQ bytes.
func EncodeVarint(input []uint32) []byte {
	var result []byte

	for _, n := range input {
		// Special case: zero
		if n == 0 {
			result = append(result, 0)
			continue
		}

		var chunks []byte
		for n > 0 {
			chunks = append(chunks, byte(n&0x7F))
			n >>= 7
		}

		// Reverse and set continuation bits
		for i := len(chunks) - 1; i >= 0; i-- {
			b := chunks[i]
			if i != 0 {
				b |= 0x80
			}
			result = append(result, b)
		}
	}

	return result
}

// DecodeVarint decodes VLQ bytes into a slice of uint32.
func DecodeVarint(input []byte) ([]uint32, error) {
	var result []uint32
	var value uint32

	for i, b := range input {
		value = (value << 7) | uint32(b&0x7F)

		// Last byte in this number
		if b&0x80 == 0 {
			result = append(result, value)
			value = 0
		} else if i == len(input)-1 {
			// Continuation bit set but no more bytes
			return nil, errors.New("incomplete byte sequence")
		}
	}

	return result, nil
}
