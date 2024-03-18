package diffsquares

func SquareOfSum(n int) int {
	var number int
	for i := n; i > 0; i-- {
		number += i
	}

	return number * number
}

func SumOfSquares(n int) int {
	var number int
	for i := n; i > 0; i-- {
		number += i * i
	}
	return number
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
