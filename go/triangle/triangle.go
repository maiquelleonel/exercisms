// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if (a <= 0 || b <= 0 || c <= 0) ||
		(a+b < c || b+c < a || a+c < b) {
		return Kind(NaT)
	}

	if a == b && a == c {
		return Kind(Equ)
	}

	if a == b || c == b || c == a {
		return Kind(Iso)
	}

	return Kind(Sca)

}
