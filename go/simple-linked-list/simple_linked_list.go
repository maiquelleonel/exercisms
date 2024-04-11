package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type Element struct {
	value int
	next  *Element
}

type List struct {
	head *Element
	size int
}

// Define the List and Element types here.

func (e *Element) String() string {
	return fmt.Sprintf("%v", e.value)
}

func (l *List) String() string {
	if l.head == nil {
		return "(empty)"
	}
	var data []string
	for el := l.head; el.next != nil; el = el.next {
		data = append(data, el.String())
	}
	return strings.Join(data, "->") + " | "
}

func New(elements []int) *List {
	l := &List{}
	for _, el := range elements {
		l.Push(el)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	node := &Element{value: element}
	if l.head == nil {
		l.head = node
		l.size++
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
		l.size++
	}
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("empty list")
	}
	var v int
	if l.head.next == nil {
		v = l.head.value
		l.head = nil
	} else {
		current := l.head
		for current.next.next != nil {
			current = current.next
		}
		v = current.next.value
		current.next = nil
	}
	l.size--
	return v, nil
}

func (l *List) Array() []int {
	var ret []int
	if l.head == nil {
		return ret
	}
	current := l.head
	for current.next != nil {
		ret = append(ret, current.value)
		current = current.next
	}
	ret = append(ret, current.value)
	return ret
}

func (l *List) Reverse() *List {
	current := l.head
	var prev *Element
	for current != nil {
		current.next, prev, current = prev, current, current.next
	}
	l.head = prev
	return l
}
