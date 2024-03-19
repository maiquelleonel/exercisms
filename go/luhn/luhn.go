package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.TrimSpace(id)

	if id == "0" ||
		doesntHaveValidLenght(id) ||
		doesntHaveValidChars(id) {
		return false
	}

	return calculate(id)

}

func onlyNumbers(id string) []string {
	re := regexp.MustCompile(`\D`)
	return strings.Split(re.ReplaceAllString(id, ""), "")
}

func doesntHaveValidLenght(id string) bool {
	return len(id) <= 1
}

func doesntHaveValidChars(id string) bool {
	re := regexp.MustCompile(`[^\d\s]`)
	return re.FindAllString(id, -1) != nil
}

func calculate(id string) bool {
	var positions, forCalc []int
	var total int
	card := onlyNumbers(id)
	positions = getPositions(card)

	for _, number := range card {
		number, _ := strconv.Atoi(number)
		forCalc = append(forCalc, number)
	}

	for _, position := range positions {
		forCalc[position] *= 2
		if forCalc[position] > 9 {
			forCalc[position] -= 9
		}
	}

	for _, number := range forCalc {
		total += number
	}

	return total%10 == 0
}

func getPositions(card []string) []int {
	var positions []int
	i := len(card) - 2
	for {
		if i <= -1 {
			break
		}
		positions = append(positions, i)
		i -= 2
	}
	return positions
}
