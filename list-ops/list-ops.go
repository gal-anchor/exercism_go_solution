package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Foldl applies a function from the left (start) of the list
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial
	for _, v := range s {
		result = fn(result, v)
	}
	return result
}

// Foldr applies a function from the right (end) of the list
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial
	for i := len(s) - 1; i >= 0; i-- {
		result = fn(s[i], result)
	}
	return result
}

// Filter returns a new list containing only elements that satisfy the predicate
func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0, len(s)) // always non-nil
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Length returns the number of elements in the list
func (s IntList) Length() int {
	return len(s)
}

// Map returns a new list with fn applied to each element
func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

// Reverse returns a new list with elements in reverse order
func (s IntList) Reverse() IntList {
	result := make(IntList, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

// Append returns a new list with lst appended to s
func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, len(s)+len(lst))
	copy(result, s)
	copy(result[len(s):], lst)
	return result
}

// Concat concatenates a slice of IntLists to s
func (s IntList) Concat(lists []IntList) IntList {
	totalLength := len(s)
	for _, lst := range lists {
		totalLength += len(lst)
	}
	result := make(IntList, 0, totalLength)
	result = append(result, s...)
	for _, lst := range lists {
		result = append(result, lst...)
	}
	return result
}
