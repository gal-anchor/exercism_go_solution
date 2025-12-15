package robotname

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type Robot struct {
	name string
}

var (
	mu        sync.Mutex
	usedNames = make(map[string]struct{})
	rng       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func (r *Robot) Name() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if r.name != "" {
		return r.name, nil
	}

	for attempts := 0; attempts < 26*26*1000; attempts++ {
		name := randomName()
		if _, exists := usedNames[name]; !exists {
			usedNames[name] = struct{}{}
			r.name = name
			return name, nil
		}
	}

	return "", errors.New("no more unique robot names available")
}

func (r *Robot) Reset() {
	mu.Lock()
	defer mu.Unlock()

	if r.name != "" {
		delete(usedNames, r.name)
		r.name = ""
	}
}

func randomName() string {
	letters := []byte{
		byte('A' + rng.Intn(26)),
		byte('A' + rng.Intn(26)),
	}

	numbers := rng.Intn(1000)

	return string(letters) + formatNumber(numbers)
}

func formatNumber(n int) string {
	if n < 10 {
		return "00" + itoa(n)
	}
	if n < 100 {
		return "0" + itoa(n)
	}
	return itoa(n)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var digits [3]byte
	i := 3
	for n > 0 {
		i--
		digits[i] = byte('0' + n%10)
		n /= 10
	}
	return string(digits[i:])
}
