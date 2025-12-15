package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	result := FreqMap{}
	ch := make(chan FreqMap)

	// Start a goroutine for each string
	for _, text := range texts {
		go func(s string) {
			ch <- Frequency(s)
		}(text)
	}

	// Collect results from all goroutines and merge
	for range texts {
		fm := <-ch
		for r, count := range fm {
			result[r] += count
		}
	}

	return result
}
