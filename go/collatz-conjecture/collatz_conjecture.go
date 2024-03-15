package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	var steps int
	if n <= 0 {
		return 0, fmt.Errorf("number %v needs to be positive and greather than 0", n)
	}
	for steps = 0; n > 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}

	return steps, nil
}
