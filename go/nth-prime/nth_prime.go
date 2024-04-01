package prime

import (
	"errors"
	"math"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n should be greather then 1")
	}

	if n < 3 {
		return n + 1, nil
	}

	numberOfPrimes := 2
	nextPrime := 3
	for {
		nextPrime++
		if isPrime(nextPrime) {
			numberOfPrimes++
		}
		if numberOfPrimes == n {
			return nextPrime, nil
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
