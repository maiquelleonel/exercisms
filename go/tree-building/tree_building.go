package tree

import (
	"cmp"
	"errors"
	"slices"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func validate(r []Record) error {
	if r[0].ID > 0 || r[0].Parent > 0 {
		return errors.New("root error")
	}

	for i := 0; i < len(r)-1; i++ {
		if r[i].ID == r[i+1].ID {
			return errors.New("duplicate registry")
		}
		if r[i+1].ID-r[i].ID > 1 {
			return errors.New("non contigous registry")
		}

		if r[i].ID > 0 && r[i].ID == r[i].Parent {
			return errors.New("circular reference")
		}
		if r[i].Parent > r[i].ID {
			return errors.New("high parent id or indirect cycle")
		}
	}

	return nil
}

func NewNode(ID int) *Node {
	return &Node{ID: ID}
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	slices.SortFunc(records, func(a, b Record) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if err := validate(records); err != nil {
		return nil, err
	}

	var Root *Node
	Root, records = NewNode(records[0].ID), records[1:]

	if len(records) == 0 {
		return Root, nil
	}

	return addNodes(Root, records), nil

}

func addNodes(root *Node, records []Record) *Node {
	nodes := make(map[int]*Node, len(records))
	nodes[root.ID] = root
	for _, record := range records {
		nodes[record.ID] = NewNode(record.ID)
		nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
	}
	return nodes[0]
}
