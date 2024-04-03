package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var sN string = strconv.Itoa(n)
	var lenght int = len(sN)
	var result float64

	for _, number := range sN {
		result += math.Pow(float64(number-'0'), float64(lenght))
	}

	return int(result) == n
}
