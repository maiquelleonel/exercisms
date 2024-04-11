package linkedlist

import (
	"errors"
	"fmt"
)

// Define List and Node types here.

type Node struct {
	Value      any
	next, prev *Node
}

type List struct {
	head, tail *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...interface{}) *List {
	list := List{}
	for _, el := range elements {
		list.Push(el)
	}
	return &list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (l *List) String() string {
	if l.head == nil {
		return "(empty)"
	}
	node := l.head
	str := node.String()
	for node := l.head.next; node != nil; node = node.next {
		str += "->" + node.String()
	}
	return str + " | "
}

func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v, next: l.head}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.prev = node
		l.head = node
	}
}

func (l *List) Push(v interface{}) {
	node := &Node{Value: v, prev: l.tail}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return int(0), errors.New("empty list")
	}
	v := l.head.Value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return v, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.head == nil {
		return int(0), errors.New("empty list")
	}
	v := l.tail.Value
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return v, nil
}

func (l *List) Reverse() *List {
	for n := l.head; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	l.head, l.tail = l.tail, l.head
	return l
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
