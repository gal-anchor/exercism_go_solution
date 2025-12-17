package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree with a single value.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree following BST rules.
func (bst *BinarySearchTree) Insert(i int) {
	if bst == nil {
		return
	}

	if i <= bst.data {
		if bst.left == nil {
			bst.left = &BinarySearchTree{data: i}
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = &BinarySearchTree{data: i}
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as a slice of ints.
func (bst *BinarySearchTree) SortedData() []int {
	if bst == nil {
		return []int{}
	}

	result := []int{}

	// Helper function for in-order traversal
	var inOrder func(node *BinarySearchTree)
	inOrder = func(node *BinarySearchTree) {
		if node == nil {
			return
		}
		inOrder(node.left)
		result = append(result, node.data)
		inOrder(node.right)
	}

	inOrder(bst)
	return result
}
