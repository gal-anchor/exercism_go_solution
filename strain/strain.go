package strain

// Keep returns a new slice containing all elements for which
// the predicate function returns true.
func Keep[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range collection {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Discard returns a new slice containing all elements for which
// the predicate function returns false.
func Discard[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range collection {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
