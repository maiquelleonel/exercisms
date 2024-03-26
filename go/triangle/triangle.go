// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if (a <= 0 || b <= 0 || c <= 0) || (a+b < c || b+c < a || a+c < b) {
		return Kind(NaT)
	}

	if a == b && a == c || b == c && b == a || c == a && c == b {
		return Kind(Equ)
	}

	if a == c && a != b || b == a && b != c || c == b && c != a {
		return Kind(Iso)
	}

	if a*2 > b+c || b*2 > a+c || c*2 > a+b {
		return Kind(Sca)
	}

	return Kind(NaT)

}
