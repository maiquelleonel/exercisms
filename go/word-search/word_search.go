package wordsearch

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type predicate func([]string, string) ([2][2]int, error)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	ret := make(map[string][2][2]int, 0)
	preds := make([]predicate, 0)
	preds = append(preds,
		leftToRight, rightToLeft, topToBottom, bottomToTop,
		topLeftToBottomRight, bottomRightToTopLeft, bottomLeftToTopRight,
		topRightToBottomLeft)

	for _, word := range words {
		for _, pred := range preds {
			if response, err := pred(puzzle, word); err == nil {
				if _, exists := ret[word]; !exists {
					ret[word] = response
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
			i := len(line) - 1
			for {
				if i < 0 {
					break
				}
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
				i--
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
			i := len(line) - 1
			for {
				if i < 0 {
					break
				}
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
				i--
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
	var c, l, d, w, firstInc, lastInc int
	var debug bool = false
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		c = 0
		l = len(puzzle) - 1
		d = l
		firstInc = 2
		lastInc = 1
		/*
			0.0 1.0 2.0 3.0 4.0 5.0 6.0 7.0 8.0 9.0
			0.1 1.1 2.1 3.1 4.1 5.1 6.1 7.1 8.1 9.1
			0.2 1.2 2.2 3.2 4.2 5.2 6.2 7.2 8.2 9.2
			0.3 1.3 2.3 3.3 4.3 5.3 6.3 7.3 8.3 9.3
			0.4 1.4 2.4 3.4 4.4 5.4 6.4 7.4 8.4 9.4
			0.5 1.5 2.5 3.5 4.5 5.5 6.5 7.5 8.5 9.5
			0.6 1.6 2.6 3.6 4.6 5.6 6.6 7.6 8.6 9.6
			0.7 1.7 2.7 3.7 4.7 5.7 6.7 7.7 8.7 9.7
			0.8 1.8 2.8 3.8 4.8 5.8 6.8 7.8 8.8 9.8
			0.9 1.9 2.9 3.9 4.9 5.9 6.9 7.9 8.9 9.9
		*/
		for {
			// time.Sleep(time.Second / 1000)
			// println(fmt.Sprintf("(%d.%d) %s %s", c, d, finded.String(), word))

			if puzzle[c][d] == word[w] {
				finded.WriteByte(word[w])
				if w == 0 {
					atLineFirst, atColFirst = c, d
				}

				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = c, d
					if debug {
						println(fmt.Sprintf("(%d.%d)(%d.%d) %s %s TLBR", atLineFirst, atColFirst, l, d, finded.String(), word))
						debug = false
					}
					finded.Reset()
					return
				}
				w++

			}
			if c == len(puzzle)-1 && d == 0 {
				break
			}
			if d == len(puzzle)-lastInc {
				c = 0
				d = len(puzzle) - firstInc
				l = d
				finded.Reset()
				w = 0
				if d == 0 {
					l = 0
					c = c + lastInc - 1
					lastInc++
				} else {
					firstInc++
				}
			} else {
				c++
				d++
			}
		}

	}
	return
}

func diagonalTraverseBottomRightToTopLeft(puzzle []string, word string) (atColFirst int, atLineFirst int, atColLast int, atLineLast int) {
	atColFirst, atLineFirst, atLineFirst, atLineLast = -1, -1, -1, -1
	var l, c, d, w, lastInc int
	var debug bool = false
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		l = len(puzzle) - 1
		d = 0
		lastInc = 2
		/*
			0.0 1.0 2.0 3.0 4.0 5.0 6.0 7.0 8.0 9.0
			0.1 1.1 2.1 3.1 4.1 5.1 6.1 7.1 8.1 9.1
			0.2 1.2 2.2 3.2 4.2 5.2 6.2 7.2 8.2 9.2
			0.3 1.3 2.3 3.3 4.3 5.3 6.3 7.3 8.3 9.3
			0.4 1.4 2.4 3.4 4.4 5.4 6.4 7.4 8.4 9.4
			0.5 1.5 2.5 3.5 4.5 5.5 6.5 7.5 8.5 9.5
			0.6 1.6 2.6 3.6 4.6 5.6 6.6 7.6 8.6 9.6
			0.7 1.7 2.7 3.7 4.7 5.7 6.7 7.7 8.7 9.7
			0.8 1.8 2.8 3.8 4.8 5.8 6.8 7.8 8.8 9.8
			0.9 1.9 2.9 3.9 4.9 5.9 6.9 7.9 8.9 9.9
		*/

		for {
			//time.Sleep(time.Second / 10)
			//println(fmt.Sprintf("(%d.%d)(%d.%d) %s %s", atLineFirst, atColFirst, d, l, finded.String(), word))
			if puzzle[l][d] == word[w] {
				finded.WriteByte(word[w])
				if w == 0 {
					atLineFirst, atColFirst = l, d
				}

				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = l, d
					if debug {
						println(fmt.Sprintf("(%d.%d)(%d.%d) %s %s BRTL", atLineFirst, atColFirst, d, l, finded.String(), word))
						debug = false
					}
					finded.Reset()
					return
				}
				w++
			}

			if c == len(puzzle)-1 && d == 0 {
				break
			}

			if d == 0 {
				if c < len(puzzle)-1 {
					c++
					l = len(puzzle) - 1
				} else {
					l = len(puzzle) - lastInc
					lastInc++
				}
				finded.Reset()
				w = 0
				d = c
			} else {
				l--
				d--
			}
		}

	}
	return
}

func diagonalTraverseBottomLeftTopRight(puzzle []string, word string) (atColFirst int, atLineFirst int, atColLast int, atLineLast int) {
	atColFirst, atLineFirst, atLineLast, atColLast = -1, -1, -1, -1
	var c, l, w, firstInc int
	var debug bool = false
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		puzzleLimit := len(puzzle) - 1
		lastLimit := 0
		firstInc = 1
		/*
			0.0 1.0 2.0 3.0 4.0 5.0 6.0 7.0 8.0 9.0
			0.1 1.1 2.1 3.1 4.1 5.1 6.1 7.1 8.1 9.1
			0.2 1.2 2.2 3.2 4.2 5.2 6.2 7.2 8.2 9.2
			0.3 1.3 2.3 3.3 4.3 5.3 6.3 7.3 8.3 9.3
			0.4 1.4 2.4 3.4 4.4 5.4 6.4 7.4 8.4 9.4
			0.5 1.5 2.5 3.5 4.5 5.5 6.5 7.5 8.5 9.5
			0.6 1.6 2.6 3.6 4.6 5.6 6.6 7.6 8.6 9.6
			0.7 1.7 2.7 3.7 4.7 5.7 6.7 7.7 8.7 9.7
			0.8 1.8 2.8 3.8 4.8 5.8 6.8 7.8 8.8 9.8
			0.9 1.9 2.9 3.9 4.9 5.9 6.9 7.9 8.9 9.9
		*/
		for {
			if c == puzzleLimit && l == puzzleLimit {
				break
			}
			if puzzle[l][c] == word[w] {
				finded.WriteByte(word[w])

				if w == 0 {
					atColFirst, atLineFirst = c, l
				}

				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = l, c
					if debug {
						println(fmt.Sprintf("(%d.%d)(%d.%d) %s %s BLTR", atColFirst, atLineFirst, atColLast, atLineLast, finded.String(), word))
						debug = false
					}
					finded.Reset()
					return
				}
				w++
			}

			if l == lastLimit {
				c = lastLimit
				l = firstInc
				if l > puzzleLimit {
					l = puzzleLimit
					c++
				}
				if firstInc <= puzzleLimit {
					firstInc++
				} else {
					lastLimit++
				}
				finded.Reset()
				w = 0
			} else {
				l--
				c++
			}
		}

	}
	return
}

func diagonalTraverseTopRightBottomLeft(puzzle []string, word string) (atColFirst int, atLineFirst int, atColLast int, atLineLast int) {
	atColFirst, atLineFirst, atLineLast, atColLast = -1, -1, -1, -1
	var c, l, w, firstInc, lastInc int
	var debug bool = false
	var finded strings.Builder
	if len(puzzle) == len(puzzle[0]) {
		puzzleLimit := len(puzzle) - 1
		firstInc = 1
		lastInc = 0
		l = puzzleLimit
		c = puzzleLimit
		/*
			0.0 1.0 2.0 3.0 4.0 5.0 6.0 7.0 8.0 9.0
			0.1 1.1 2.1 3.1 4.1 5.1 6.1 7.1 8.1 9.1
			0.2 1.2 2.2 3.2 4.2 5.2 6.2 7.2 8.2 9.2
			0.3 1.3 2.3 3.3 4.3 5.3 6.3 7.3 8.3 9.3
			0.4 1.4 2.4 3.4 4.4 5.4 6.4 7.4 8.4 9.4
			0.5 1.5 2.5 3.5 4.5 5.5 6.5 7.5 8.5 9.5
			0.6 1.6 2.6 3.6 4.6 5.6 6.6 7.6 8.6 9.6
			0.7 1.7 2.7 3.7 4.7 5.7 6.7 7.7 8.7 9.7
			0.8 1.8 2.8 3.8 4.8 5.8 6.8 7.8 8.8 9.8
			0.9 1.9 2.9 3.9 4.9 5.9 6.9 7.9 8.9 9.9
		*/
		for {
			if c == 0 && l == 0 {
				break
			}
			if puzzle[l][c] == word[0] {
				atLineFirst, atColFirst = l, c
			}

			if puzzle[l][c] == word[w] {
				finded.WriteByte(word[w])

				if strings.Contains(finded.String(), word) {
					atLineLast, atColLast = l, c
					//time.Sleep(time.Second / 3)
					if debug {
						println(fmt.Sprintf("(%d.%d)(%d.%d) %s %s TRBL", atColFirst, atLineFirst, atColLast, atLineLast, finded.String(), word))
						debug = false
					}
					finded.Reset()
					return
				}
				w++
			}

			if l == puzzleLimit-lastInc {
				c = puzzleLimit - lastInc
				l = puzzleLimit - firstInc
				if l >= 0 {
					firstInc++
				} else {
					lastInc++
					c = puzzleLimit - lastInc
					l = 0
				}
				finded.Reset()
				w = 0
			} else {
				c--
				l++
			}
		}

	}
	return
}
