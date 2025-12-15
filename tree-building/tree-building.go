package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Sort records by ID to ensure proper tree construction
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))

	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("non-continuous or missing record ID")
		}
		if r.ID == 0 && r.Parent != 0 {
			return nil, errors.New("root node cannot have a parent")
		}
		if r.ID != 0 && r.ID <= r.Parent {
			return nil, errors.New("invalid parent ID")
		}

		nodes[i] = &Node{ID: r.ID}
		if r.ID != 0 {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
