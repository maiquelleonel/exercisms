package wordsearch

import (
	"errors"
	"slices"
	"strings"
)

type predicate func([]string, string) ([2][2]int, error)

var preds = []predicate{
	leftToRight,
	rightToLeft,
	topToBottom,
	bottomToTop,
	topLeftToBottomRight,
	bottomRightToTopLeft,
	bottomLeftToTopRight,
	topRightToBottomLeft,
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	ret := make(map[string][2][2]int, 0)

	for _, word := range words {
		for _, pred := range preds {
			if response, err := pred(puzzle, word); err == nil {
				if _, exists := ret[word]; !exists {
					ret[word] = response
					break
				}
			}
		}
		if _, exits := ret[word]; !exits {
			return nil, errors.New("word not in puzzle")
		}
	}
	if len(ret) == 0 {
		return nil, errors.New("word not found")
	}
	return ret, nil
}

func leftToRight(puzzle []string, word string) ([2][2]int, error) {
	firstCharPosition, lastCharPosition, atLine := -1, -1, -1
	for l, line := range puzzle {
		if strings.Contains(line, word) {
			firstCharPosition = strings.IndexByte(line, word[0])
			lastCharPosition = strings.LastIndexByte(line, word[len(word)-1])
			if firstCharPosition != -1 && lastCharPosition != -1 {
				atLine = l
				break
			}
		}
	}
	if firstCharPosition != -1 && lastCharPosition != -1 {
		return [2][2]int{{firstCharPosition, atLine}, {lastCharPosition, atLine}}, nil
	}
	return [2][2]int{}, errors.New("not found from left to right")
}

func reverse(st string) (rs string) {
	r := strings.Split(st, "")
	slices.Reverse(r)
	rs = strings.Join(r, "")
	return
}

func rightToLeft(puzzle []string, word string) ([2][2]int, error) {
	firstCharPosition, lastCharPosition, atLine := -1, -1, -1
	for l, line := range puzzle {
		if strings.Contains(line, reverse(word)) {
			start, end := word[0], word[len(word)-1]
			for i := len(line) - 1; i >= 0; i-- {
				if line[i] == start {
					firstCharPosition = i
				}
				if line[i] == end {
					lastCharPosition = i
					if firstCharPosition != -1 {
						atLine = l
						break
					}
				}

			}
		}
	}
	if firstCharPosition != -1 && lastCharPosition != -1 {
		return [2][2]int{{firstCharPosition, atLine}, {lastCharPosition, atLine}}, nil
	}
	return [2][2]int{}, errors.New("not found from right to left")
}

func topToBottom(puzzle []string, word string) ([2][2]int, error) {
	firstCharPosition, lastCharPosition, atLine := -1, -1, -1
	var localPuzzle []string = puzzle
	lines := traversePuzzle(localPuzzle)

	for l, line := range lines {
		if len(line) < len(word) {
			return [2][2]int{}, errors.New("invalid lenght")
		}
		if strings.Contains(line, word) {
			firstCharPosition = strings.IndexByte(line, word[0])
			lastCharPosition = strings.LastIndexByte(line, word[len(word)-1])
			if firstCharPosition != -1 && lastCharPosition != -1 {
				atLine = l
				break
			}
		}
	}

	if firstCharPosition != -1 && lastCharPosition != -1 {
		return [2][2]int{{atLine, firstCharPosition}, {atLine, lastCharPosition}}, nil
	}
	return [2][2]int{}, errors.New("not found from top to bottom")
}

func bottomToTop(puzzle []string, word string) ([2][2]int, error) {
	firstCharPosition, lastCharPosition, atLine := -1, -1, -1
	var localPuzzle []string = puzzle
	lines := traversePuzzle(localPuzzle)
	for l, line := range lines {
		if strings.Contains(reverse(line), word) {
			start, end := word[0], word[len(word)-1]
			for i := len(line) - 1; i >= 0; i-- {
				if line[i] == start {
					firstCharPosition = i
				}
				if line[i] == end {
					lastCharPosition = i
					if firstCharPosition != -1 {
						atLine = l
						break
					}
				}
			}
		}
	}

	if firstCharPosition != -1 && lastCharPosition != -1 {
		return [2][2]int{{atLine, firstCharPosition}, {atLine, lastCharPosition}}, nil
	}
	return [2][2]int{}, errors.New("not found from bottom to top")
}

func topLeftToBottomRight(puzzle []string, word string) ([2][2]int, error) {
	atCol, firstChar, atLine, lastChar := -1, -1, -1, -1
	var localPuzzle []string = puzzle
	atCol, firstChar, atLine, lastChar = diagonalTraverseTopLeftBottomRight(localPuzzle, word)

	if atCol != -1 && firstChar != -1 && atLine != -1 && lastChar != -1 {
		return [2][2]int{{atCol, firstChar}, {atLine, lastChar}}, nil
	}
	return [2][2]int{}, errors.New("not found from bottom to top")
}

func bottomRightToTopLeft(puzzle []string, word string) ([2][2]int, error) {
	atCol, firstChar, atLine, lastChar := -1, -1, -1, -1
	var localPuzzle []string = puzzle
	atCol, firstChar, atLine, lastChar = diagonalTraverseBottomRightToTopLeft(localPuzzle, word)

	if atCol != -1 && firstChar != -1 && atLine != -1 && lastChar != -1 {
		return [2][2]int{{atCol, firstChar}, {atLine, lastChar}}, nil
	}
	return [2][2]int{}, errors.New("not found from bottom to top")
}

func bottomLeftToTopRight(puzzle []string, word string) ([2][2]int, error) {
	atCol, firstChar, atLine, lastChar := -1, -1, -1, -1
	var localPuzzle []string = puzzle
	atCol, firstChar, atLine, lastChar = diagonalTraverseBottomLeftTopRight(localPuzzle, word)

	if atCol != -1 && firstChar != -1 && atLine != -1 && lastChar != -1 {
		return [2][2]int{{atCol, firstChar}, {atLine, lastChar}}, nil
	}
	return [2][2]int{}, errors.New("not found from bottom to top")
}

func topRightToBottomLeft(puzzle []string, word string) ([2][2]int, error) {
	atCol, firstChar, atLine, lastChar := -1, -1, -1, -1
	var localPuzzle []string = puzzle
	atCol, firstChar, atLine, lastChar = diagonalTraverseTopRightBottomLeft(localPuzzle, word)

	if atCol != -1 && firstChar != -1 && atLine != -1 && lastChar != -1 {
		return [2][2]int{{atCol, firstChar}, {atLine, lastChar}}, nil
	}
	return [2][2]int{}, errors.New("not found from bottom to top")
}

func traversePuzzle(puzzle []string) []string {
	ret := make([]string, 0, len(puzzle))
	var col strings.Builder
	for i := 0; i < len(puzzle); i++ {
		for _, line := range puzzle {
			if len(line) > 1 {
				col.WriteByte(line[i])
			}
		}
		ret = append(ret, col.String())
		col.Reset()
	}
	return slices.Clip(ret)
}

func diagonalTraverseTopLeftBottomRight(puzzle []string, word string) (atColFirst int, atLineFirst int, atLineLast int, atColLast int) {
	atColFirst, atLineFirst, atLineLast, atColLast = -1, -1, -1, -1
	var column, line, diagonalIndex, wordIndex, firstInc, lastInc int
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		column = 0
		line = len(puzzle) - 1
		diagonalIndex = line
		firstInc = 2
		lastInc = 1
		for {
			if puzzle[column][diagonalIndex] == word[wordIndex] {
				finded.WriteByte(word[wordIndex])
				if wordIndex == 0 {
					atLineFirst, atColFirst = column, diagonalIndex
				}
				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = column, diagonalIndex
					finded.Reset()
					return
				}
				wordIndex++

			}
			if column == len(puzzle)-1 && diagonalIndex == 0 {
				break
			}
			if diagonalIndex == len(puzzle)-lastInc {
				column = 0
				diagonalIndex = len(puzzle) - firstInc
				line = diagonalIndex
				finded.Reset()
				wordIndex = 0
				if diagonalIndex == 0 {
					line = 0
					column = column + lastInc - 1
					lastInc++
				} else {
					firstInc++
				}
			} else {
				column++
				diagonalIndex++
			}
		}

	}
	return
}

func diagonalTraverseBottomRightToTopLeft(puzzle []string, word string) (atColFirst int, atLineFirst int, atColLast int, atLineLast int) {
	atColFirst, atLineFirst, atLineFirst, atLineLast = -1, -1, -1, -1
	var line, column, diagonalIndex, wordIndex, lastInc int
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		line = len(puzzle) - 1
		diagonalIndex = 0
		lastInc = 2

		for {
			if puzzle[line][diagonalIndex] == word[wordIndex] {
				finded.WriteByte(word[wordIndex])
				if wordIndex == 0 {
					atLineFirst, atColFirst = line, diagonalIndex
				}

				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = line, diagonalIndex
					finded.Reset()
					return
				}
				wordIndex++
			}

			if column == len(puzzle)-1 && diagonalIndex == 0 {
				break
			}

			if diagonalIndex == 0 {
				if column < len(puzzle)-1 {
					column++
					line = len(puzzle) - 1
				} else {
					line = len(puzzle) - lastInc
					lastInc++
				}
				finded.Reset()
				wordIndex = 0
				diagonalIndex = column
			} else {
				line--
				diagonalIndex--
			}
		}

	}
	return
}

func diagonalTraverseBottomLeftTopRight(puzzle []string, word string) (atColFirst int, atLineFirst int, atColLast int, atLineLast int) {
	atColFirst, atLineFirst, atLineLast, atColLast = -1, -1, -1, -1
	var column, line, wordIndex, firstInc int
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		puzzleLimit := len(puzzle) - 1
		lastLimit := 0
		firstInc = 1
		for {
			if column == puzzleLimit && line == puzzleLimit {
				break
			}
			if puzzle[line][column] == word[wordIndex] {
				finded.WriteByte(word[wordIndex])

				if wordIndex == 0 {
					atColFirst, atLineFirst = column, line
				}

				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = line, column
					finded.Reset()
					return
				}
				wordIndex++
			}

			if line == lastLimit {
				column = lastLimit
				line = firstInc
				if line > puzzleLimit {
					line = puzzleLimit
					column++
				}
				if firstInc <= puzzleLimit {
					firstInc++
				} else {
					lastLimit++
				}
				finded.Reset()
				wordIndex = 0
			} else {
				line--
				column++
			}
		}

	}
	return
}

func diagonalTraverseTopRightBottomLeft(puzzle []string, word string) (atColFirst int, atLineFirst int, atColLast int, atLineLast int) {
	atColFirst, atLineFirst, atLineLast, atColLast = -1, -1, -1, -1
	var column, line, wordIndex, firstInc, lastInc int
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		puzzleLimit := len(puzzle) - 1
		firstInc = 1
		lastInc = 0
		line = puzzleLimit
		column = puzzleLimit
		for {
			if column == 0 && line == 0 {
				break
			}
			if puzzle[line][column] == word[0] {
				atLineFirst, atColFirst = line, column
			}

			if puzzle[line][column] == word[wordIndex] {
				finded.WriteByte(word[wordIndex])
				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = line, column
					finded.Reset()
					return
				}
				wordIndex++
			}

			if line == puzzleLimit-lastInc {
				column = puzzleLimit - lastInc
				line = puzzleLimit - firstInc
				if line >= 0 {
					firstInc++
				} else {
					lastInc++
					column = puzzleLimit - lastInc
					line = 0
				}
				finded.Reset()
				wordIndex = 0
			} else {
				column--
				line++
			}
		}

	}
	return
}
