package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, it := range s {
		initial = fn(initial, it)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	s.Reverse()
	for _, it := range s {
		initial = fn(it, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if len(s) == 0 {
		return s
	}
	var r IntList
	for _, it := range s {
		if fn(it) {
			r = append(r, it)
		}
	}
	return r
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, it := range s {
		s[i] = fn(it)
	}
	return s
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	s = append(s, lst...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = append(s, list...)
	}
	return s
}
