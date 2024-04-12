package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var a, b, c int
	var ret []Triplet

	for a = min; a <= max; a++ {
		for b = a; b <= max; b++ {
			for c = b; c <= max; c++ {
				if a*a+b*b == c*c {
					ret = append(ret, Triplet{a, b, c})
					break
				}
				if a*a+b*b < c*c {
					break
				}
			}
		}
	}

	return ret
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var ret []Triplet
	// a, b always div by 2, 3, 4, 5
	// a < b < c
	// a = p / 3
	// b = (p-a) / 2
	// c = p - a - b
	for a := 1; a <= (p / 3); a++ {
		for b := a; b <= (p-a)/2; b++ {
			c := p - a - b
			if a*a+b*b == c*c {
				ret = append(ret, Triplet{a, b, c})
			}
		}
	}

	return ret
}
