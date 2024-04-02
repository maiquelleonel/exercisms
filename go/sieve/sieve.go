package sieve

func Sieve(limit int) []int {

	if limit < 2 {
		return []int{}
	}

	primes := make([]bool, limit+1)

	for i := range primes {
		primes[i] = true
	}

	var prime int = 2
	for prime*prime < limit {
		if primes[prime] {
			for i := prime * prime; i <= limit; i += prime {
				primes[i] = false
			}
		}
		prime += 1
	}

	var ret []int
	for i, prime := range primes {
		if i < 2 {
			continue
		}
		if prime {
			ret = append(ret, i)
		}
	}

	return ret

}
