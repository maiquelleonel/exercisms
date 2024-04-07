package matrix

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	m := Matrix{}
	cell := []int{}
	var lastColsLen int
	for i, row := range rows {

		if row == "" {
			return Matrix{}, errors.New("invalid pattern")
		}

		cells := strings.Split(row, " ")

		cells = slices.DeleteFunc(cells, func(it string) bool {
			return it == ""
		})

		if i > 0 && len(cells) > len(rows) && len(cells) > lastColsLen {
			return Matrix{}, errors.New("uneven rows lenght")
		}

		lastColsLen = len(cells)

		for _, value := range cells {
			if vl, err := strconv.Atoi(value); err == nil {
				cell = append(cell, vl)
			} else {
				return Matrix{}, errors.New("invalid integer")
			}
		}

		m = append(m, cell)
		cell = []int{}

	}

	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	ret := make([][]int, 0)
	col := make([]int, 0)

	for column := range m[0] {
		for row := range m {
			col = append(col, m[row][column])
		}
		ret = append(ret, col)
		col = []int{}
	}

	return ret
}

func (m Matrix) Rows() [][]int {
	ret := [][]int{}

	for _, row := range m {
		ret = append(ret, slices.Clone(row))
	}

	return ret
}

func (m Matrix) Set(row, col, val int) bool {
	if (row >= 0 && row < len(m.Rows())) && (col >= 0 && col < len(m.Cols())) {
		m[row][col] = val
		return true
	}
	return false
}
