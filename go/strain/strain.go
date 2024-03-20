package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](s []T, predicate func(v T) bool) []T {
	ret := make([]T, 0)
	for _, v := range s {
		if predicate(v) {
			ret = append(ret, v)
		}
	}

	return ret

}

func Discard[T any](s []T, predicate func(v T) bool) []T {
	ret := make([]T, 0)
	for _, v := range s {
		if !predicate(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
