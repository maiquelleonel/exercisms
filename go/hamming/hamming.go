package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("sequence lenght msut be equal")
	}
	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil

}
