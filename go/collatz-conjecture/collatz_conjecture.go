package collatzconjecture

import "fmt"

var steps = 0

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("number %v needs to be positive and greather than 0", n)
	}
	if n > 1 {
		steps++
		if n%2 == 0 {
			return CollatzConjecture(n / 2)
		} else {
			return CollatzConjecture(3*n + 1)
		}
	}
	total := steps
	steps = 0
	return total, nil
}
