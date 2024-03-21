package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("Invalid")
	}

	roman_numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman_letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	current := input
	ret := ""
	for i, num := range roman_numbers {
		for current >= num {
			ret += roman_letters[i]
			current -= num
		}
	}

	return ret, nil
}
