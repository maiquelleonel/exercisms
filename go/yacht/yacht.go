package yacht

import (
	"slices"
)

const (
	YACHT           = "yacht"
	ONES            = "ones"
	TWOS            = "twos"
	THREES          = "threes"
	FOURS           = "fours"
	FIVES           = "fives"
	SIXES           = "sixes"
	FULL_HOUSE      = "full house"
	FOUR_OF_A_KIND  = "four of a kind"
	LITTLE_STRAIGHT = "little straight"
	BIG_STRAIGHT    = "big straight"
	CHOICE          = "choice"
)

var scores map[string]int = map[string]int{
	ONES:   1,
	TWOS:   2,
	THREES: 3,
	FOURS:  4,
	FIVES:  5,
	SIXES:  6,
}
var categories []string = []string{
	ONES,
	TWOS,
	THREES,
	FOURS,
	FIVES,
	SIXES,
}

func Score(dice []int, category string) int {
	fns := make(map[string]func([]int) int, 6)

	fns[YACHT] = func(dice []int) int {
		total := totalSides(dice)
		if total[5] > 0 {
			return 50
		}
		return 0
	}

	fns[FULL_HOUSE] = func(dice []int) int {
		if isFullHouse(dice) {
			return sum(dice)
		}
		return 0
	}

	fns[FOUR_OF_A_KIND] = func(dice []int) int {
		if isFourAKind(dice) {
			return slices.Max(dice) * 4
		}
		return 0
	}

	fns[LITTLE_STRAIGHT] = func(dice []int) int {
		slices.Sort(dice)
		if slices.Compare(dice, []int{1, 2, 3, 4, 5}) == 0 {
			return 30
		}
		return 0
	}

	fns[BIG_STRAIGHT] = func(dice []int) int {
		slices.Sort(dice)
		if slices.Compare(dice, []int{2, 3, 4, 5, 6}) == 0 {
			return 30
		}
		return 0
	}

	fns[CHOICE] = func(dice []int) int {
		return sum(dice)
	}

	if slices.Index(categories, category) != -1 {
		return calcScore(dice, scores[category], scores[category])
	}

	return fns[category](dice)
}

func calcScore(dice []int, side int, score int) int {
	total := countSides(dice)
	return total[side] * score
}

func sum(dice []int) int {
	var r int
	for _, side := range dice {
		r += side
	}
	return r
}

func countSides(dice []int) map[int]int {
	var r map[int]int = map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, side := range dice {
		r[side] += 1
	}
	return r
}

func totalSides(dice []int) map[int]int {
	totals := make(map[int]int, 0)
	for side, count := range countSides(dice) {
		totals[count] = side
	}
	return totals
}

func isFourAKind(dice []int) bool {
	total := totalSides(dice)
	return total[4] > 0 || total[5] > 0
}

func isFullHouse(dice []int) bool {
	total := totalSides(dice)
	return total[2] > 0 && total[3] > 0
}
