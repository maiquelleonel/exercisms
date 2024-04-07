package stringset

import (
	"fmt"
	"slices"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	data []string
}

func New() Set {
	return Set{data: []string{}}
}

func NewFromSlice(l []string) Set {
	slices.Sort(l)
	return Set{data: slices.Compact(l)}
}

func (s Set) String() string {
	if s.IsEmpty() {
		return fmt.Sprintf("{%v}", "")
	}
	return fmt.Sprintf("{%v}", `"`+strings.Join(s.data, `", "`)+`"`)
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(elem string) bool {
	return slices.Index(s.data, elem) != -1
}

func (s *Set) Add(elem string) {
	if s.IsEmpty() || !s.Has(elem) {
		s.data = append(s.data, elem)
	}
}

func Subset(s1, s2 Set) bool {
	if len(s1.data) < len(s2.data) {
		return true
	} else if len(s1.data) > len(s2.data) {
		return false
	} else {
		for _, el := range s2.data {
			if slices.Index(s1.data, el) < 0 {
				return false
			}
		}
		return true
	}
}
func Disjoint(s1, s2 Set) bool {
	if len(s1.data) < len(s2.data) {
		return true
	}
	for _, el := range s2.data {
		if slices.Index(s1.data, el) >= 0 {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	for _, el := range s2.data {
		if slices.Index(s1.data, el) < 0 {
			return false
		}
	}
	return len(s1.data) == len(s2.data)
}

func Intersection(s1, s2 Set) Set {
	var inter Set
	for _, el := range s2.data {
		if s1.Has(el) {
			inter.Add(el)
		}
	}
	return inter
}

func Difference(s1, s2 Set) Set {
	var inter Set
	if s1.IsEmpty() {
		return inter
	}

	if s2.IsEmpty() {
		return s1
	}

	for _, el := range s1.data {
		if !s2.Has(el) {
			inter.Add(el)
		}
	}

	return inter

}

func Union(s1 Set, s2 Set) Set {
	s := slices.Concat(s1.data, s2.data)
	return NewFromSlice(s)
}
