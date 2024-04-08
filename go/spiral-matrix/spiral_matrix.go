package spiralmatrix

type Square struct {
	startColumn, startRow, endColumn, endRow int
}

func SpiralMatrix(size int) [][]int {
	if size == 0 {
		return [][]int{}
	}
	if size == 1 {
		return [][]int{{1}}
	}

	matrix := make([][]int, 0, size*size)

	for i := 0; i < size; i++ {
		matrix = append(matrix, make([]int, size))
	}
	size--
	sq := Square{0, 0, size, size}
	var i int
	var value int = 1

	for sq.startRow <= sq.endRow && sq.startColumn <= sq.endColumn {
		for i = sq.startColumn; i <= sq.endColumn; i++ {
			matrix[sq.startRow][i] = value
			value++
		}
		sq.startRow++

		for i = sq.startRow; i <= sq.endRow; i++ {
			matrix[i][sq.endColumn] = value
			value++

		}
		sq.endColumn--

		for i = sq.endColumn; i >= sq.startColumn; i-- {
			matrix[sq.endRow][i] = value
			value++
		}
		sq.endRow--

		for i = sq.endRow; i > sq.startColumn; i-- {
			matrix[i][sq.startColumn] = value
			value++
		}
		sq.startColumn++
	}

	return matrix
}
