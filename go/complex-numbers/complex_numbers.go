package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		(n1.Real() + n2.Real()),
		(n1.Imaginary() + n2.Imaginary()),
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		(n1.Real() - n2.Real()),
		(n1.Imaginary() - n2.Imaginary()),
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		(n1.Real() * n2.Real()) - (n1.Imaginary() * n2.Imaginary()),
		(n1.Imaginary() * n2.Real()) + (n1.Real() * n2.Imaginary()),
	}
}

func (n Number) Times(factor float64) Number {
	return Number{n.Real() * factor, n.Imaginary() * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	var divd float64 = ((n2.Real() * n2.Real()) + (n2.Imaginary() * n2.Imaginary()))
	return Number{
		((n1.Real() * n2.Real()) + (n1.Imaginary() * n2.Imaginary())) / divd,
		((n1.Imaginary() * n2.Real()) - (n1.Real() * n2.Imaginary())) / divd,
	}
}

func (n Number) Conjugate() Number {
	return Number{n.Real(), -n.Imaginary()}
}

func (n Number) Abs() float64 {
	return math.Abs(
		math.Sqrt(
			(n.Real() * n.Real()) + (n.Imaginary() * n.Imaginary()),
		),
	)
}

func (n Number) Exp() Number {
	var e float64 = math.Exp(n.Real())
	return Number{
		e * math.Cos(n.Imaginary()),
		e * math.Sin(n.Imaginary()),
	}
}
