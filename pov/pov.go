package pov

type Tree struct {
	value    string
	children []*Tree
	parent   *Tree // Helpful for navigating back up during re-rooting
}

// New creates and returns a new Tree.
func New(value string, children ...*Tree) *Tree {
	tr := &Tree{value: value, children: children}
	for _, child := range children {
		child.parent = tr
	}
	return tr
}

func (tr *Tree) Value() string {
	return tr.value
}

func (tr *Tree) Children() []*Tree {
	return tr.children
}

// FromPov returns the tree re-rooted at the 'from' node.
func (tr *Tree) FromPov(from string) *Tree {
	target := tr.find(from)
	if target == nil {
		return nil
	}

	return target.reroot()
}

// find performs a DFS to locate the node with the given value.
func (tr *Tree) find(val string) *Tree {
	if tr.value == val {
		return tr
	}
	for _, child := range tr.children {
		if found := child.find(val); found != nil {
			return found
		}
	}
	return nil
}

// reroot is a recursive helper that flips the parent-child relationship.
func (tr *Tree) reroot() *Tree {
	if tr.parent == nil {
		return tr
	}

	parent := tr.parent
	// Remove 'tr' from its parent's children
	newChildren := make([]*Tree, 0)
	for _, child := range parent.children {
		if child != tr {
			newChildren = append(newChildren, child)
		}
	}
	parent.children = newChildren

	// Move parent to be a child of this node
	tr.children = append(tr.children, parent.reroot())
	tr.parent = nil
	parent.parent = tr

	return tr
}

// PathTo returns the shortest path between two nodes.
func (tr *Tree) PathTo(from, to string) []string {
	rooted := tr.FromPov(from)
	if rooted == nil {
		return nil
	}

	target := rooted.find(to)
	if target == nil {
		return nil
	}

	// Trace back from target to the new root using the parent pointers
	var path []string
	for curr := target; curr != nil; curr = curr.parent {
		path = append([]string{curr.value}, path...)
	}
	return path
}
