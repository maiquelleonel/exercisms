package stringset

import (
	"fmt"
	"reflect"
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
	return Set{[]string{}}
}

func NewFromSlice(l []string) Set {
	slices.Sort(l)
	return Set{slices.Compact(l)}
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

func (s Set) Len() int {
	return len(s.data)
}

func (s *Set) Add(elem string) {
	if s.IsEmpty() || !s.Has(elem) {
		s.data = append(s.data, elem)
		slices.Sort(s.data)
	}
}

func Subset(s1, s2 Set) bool {
	return Equal(s1, Intersection(s1, s2))
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).Len() == 0

}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	inter := New()
	for _, el := range s2.data {
		if s1.Has(el) {
			inter.Add(el)
		}
	}
	return inter
}

func Difference(s1, s2 Set) Set {
	inter := New()
	for _, el := range s1.data {
		if !s2.Has(el) {
			inter.Add(el)
		}
	}
	return inter
}

func Union(s1 Set, s2 Set) Set {
	return NewFromSlice(append(s1.data, s2.data...))
}
