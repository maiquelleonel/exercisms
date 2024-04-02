package prime

func Factors(n int64) []int64 {
	var ret []int64
	var i int64 = 2
	for n >= i {
		if n%i == 0 {
			ret = append(ret, i)
			n /= i
		} else {
			i++
		}
	}
	return ret
}
