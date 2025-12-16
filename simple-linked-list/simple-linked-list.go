package linkedlist

import "errors"

// Element represents a node in the linked list
type Element struct {
	Value int
	Next  *Element
}

// List represents the linked list
type List struct {
	head *Element
	size int
}

// New creates a new linked list from a slice of integers
func New(elements []int) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

// Size returns the number of elements in the list
func (l *List) Size() int {
	return l.size
}

// Push adds an element to the end of the list
func (l *List) Push(value int) {
	newElem := &Element{Value: value}
	if l.head == nil {
		l.head = newElem
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newElem
	}
	l.size++
}

// Pop removes the last element and returns its value
func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}

	var prev *Element
	current := l.head
	for current.Next != nil {
		prev = current
		current = current.Next
	}

	if prev == nil {
		// Only one element
		l.head = nil
	} else {
		prev.Next = nil
	}
	l.size--
	return current.Value, nil
}

// Array converts the linked list into a slice of integers
func (l *List) Array() []int {
	arr := make([]int, 0, l.size)
	current := l.head
	for current != nil {
		arr = append(arr, current.Value)
		current = current.Next
	}
	return arr
}

// Reverse returns a new list with elements in reversed order
func (l *List) Reverse() *List {
	newList := &List{}
	current := l.head
	for current != nil {
		newElem := &Element{Value: current.Value, Next: newList.head}
		newList.head = newElem
		newList.size++
		current = current.Next
	}
	return newList
}
