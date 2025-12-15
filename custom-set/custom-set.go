package stringset

import (
	"sort"
	"strconv"
	"strings"
)

// Set represents a collection of unique string values.
type Set struct {
	elements map[string]struct{}
}

// New creates and returns a new empty set.
func New() Set {
	return Set{elements: make(map[string]struct{})}
}

// NewFromSlice creates a set from a slice of strings.
func NewFromSlice(l []string) Set {
	s := New()
	for _, e := range l {
		s.Add(e)
	}
	return s
}

// String returns the string representation of the set.
func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}

	elems := make([]string, 0, len(s.elements))
	for e := range s.elements {
		elems = append(elems, strconv.Quote(e))
	}

	// Sort for consistent output
	sort.Strings(elems)
	return "{" + strings.Join(elems, ", ") + "}"
}

// IsEmpty returns true if the set has no elements.
func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

// Has checks whether an element exists in the set.
func (s Set) Has(elem string) bool {
	_, exists := s.elements[elem]
	return exists
}

// Add inserts an element into the set.
func (s Set) Add(elem string) {
	s.elements[elem] = struct{}{}
}

// Subset returns true if s1 is a subset of s2.
func Subset(s1, s2 Set) bool {
	for e := range s1.elements {
		if !s2.Has(e) {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common.
func Disjoint(s1, s2 Set) bool {
	for e := range s1.elements {
		if s2.Has(e) {
			return false
		}
	}
	return true
}

// Equal returns true if s1 and s2 contain the same elements.
func Equal(s1, s2 Set) bool {
	if len(s1.elements) != len(s2.elements) {
		return false
	}
	return Subset(s1, s2)
}

// Intersection returns a new set containing elements present in both s1 and s2.
func Intersection(s1, s2 Set) Set {
	result := New()
	for e := range s1.elements {
		if s2.Has(e) {
			result.Add(e)
		}
	}
	return result
}

// Difference returns a new set containing elements in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	result := New()
	for e := range s1.elements {
		if !s2.Has(e) {
			result.Add(e)
		}
	}
	return result
}

// Union returns a new set containing all elements from s1 and s2.
func Union(s1, s2 Set) Set {
	result := New()
	for e := range s1.elements {
		result.Add(e)
	}
	for e := range s2.elements {
		result.Add(e)
	}
	return result
}
