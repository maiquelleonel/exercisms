package summultiples

import (
	"slices"
)

func SumMultiples(limit int, divisors ...int) int {
	var mults []int
	var res int
	for _, div := range divisors {
		for i := 0; i < limit; i += div {
			if div == 0 {
				i++
				continue
			}
			if i%div == 0 {
				mults = append(mults, i)
			}
		}
	}

	if mults == nil {
		return 0
	}

	slices.Sort(mults)
	mults = slices.Compact(mults)

	for _, v := range mults {
		res += v
	}

	return res
}
