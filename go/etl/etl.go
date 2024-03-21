package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	ret := make(map[string]int, 0)
	for point, letters := range in {
		for _, letter := range letters {
			ret[strings.ToLower(letter)] = point
		}
	}
	return ret
}
