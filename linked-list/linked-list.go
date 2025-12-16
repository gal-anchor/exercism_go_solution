package linkedlist

import "errors"

// Node represents a node in a doubly linked list.
type Node struct {
	Value any
	prev  *Node
	next  *Node
}

// List represents a doubly linked list.
type List struct {
	head *Node
	tail *Node
	len  int
}

// NewList creates a new list optionally initialized with elements.
func NewList(elements ...any) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

// Next returns the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift inserts a value at the beginning of the list.
func (l *List) Unshift(v any) {
	node := &Node{Value: v}
	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.len++
}

// Push inserts a value at the end of the list.
func (l *List) Push(v any) {
	node := &Node{Value: v}
	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

// Shift removes and returns the first element of the list.
func (l *List) Shift() (any, error) {
	if l.len == 0 {
		return nil, errors.New("list is empty")
	}
	node := l.head
	if l.len == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = node.next
		l.head.prev = nil
		node.next = nil
	}
	l.len--
	return node.Value, nil
}

// Pop removes and returns the last element of the list.
func (l *List) Pop() (any, error) {
	if l.len == 0 {
		return nil, errors.New("list is empty")
	}
	node := l.tail
	if l.len == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = node.prev
		l.tail.next = nil
		node.prev = nil
	}
	l.len--
	return node.Value, nil
}

// Reverse reverses the list in-place.
func (l *List) Reverse() {
	current := l.head
	var prev *Node
	l.tail = l.head
	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	l.head = prev
}

// First returns the first node in the list.
func (l *List) First() *Node {
	return l.head
}

// Last returns the last node in the list.
func (l *List) Last() *Node {
	return l.tail
}
