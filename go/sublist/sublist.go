package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if contains(l1, l2) {
		if len(l1) == len(l2) {
			return RelationEqual
		}
		return RelationSuperlist
	}

	if contains(l2, l1) {
		return RelationSublist
	}

	return RelationUnequal
}

func contains(maj, min []int) bool {
	var i, j int
	for ; i < len(maj) && j < len(min); i++ {
		if maj[i] == min[j] {
			j++
		} else {
			i, j = i-j, 0
		}
	}
	return j == len(min)
}
