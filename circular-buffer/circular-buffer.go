package circular

import "errors"

var ErrEmptyBuffer = errors.New("buffer is empty")
var ErrFullBuffer = errors.New("buffer is full")

type Buffer struct {
	data       []byte
	size       int
	readIndex  int
	writeIndex int
	count      int
}

// NewBuffer creates a new circular buffer of the given size
func NewBuffer(size int) *Buffer {
	return &Buffer{
		data: make([]byte, size),
		size: size,
	}
}

// ReadByte reads a byte from the buffer. Returns error if empty.
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, ErrEmptyBuffer
	}
	c := b.data[b.readIndex]
	b.readIndex = (b.readIndex + 1) % b.size
	b.count--
	return c, nil
}

// WriteByte writes a byte to the buffer. Returns error if full.
func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
		return ErrFullBuffer
	}
	b.data[b.writeIndex] = c
	b.writeIndex = (b.writeIndex + 1) % b.size
	b.count++
	return nil
}

// Overwrite writes a byte, overwriting the oldest data if buffer is full
func (b *Buffer) Overwrite(c byte) {
	if b.count == b.size {
		// buffer full: advance readIndex to drop oldest byte
		b.readIndex = (b.readIndex + 1) % b.size
		b.count--
	}
	b.WriteByte(c) // WriteByte cannot fail now
}

// Reset clears the buffer
func (b *Buffer) Reset() {
	b.readIndex = 0
	b.writeIndex = 0
	b.count = 0
}
