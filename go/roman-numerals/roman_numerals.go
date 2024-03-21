package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("Invalid")
	}

	ret := strings.Repeat("I", input)
	ret = strings.ReplaceAll(ret, "IIIII", "V")
	ret = strings.ReplaceAll(ret, "IIII", "IV")
	ret = strings.ReplaceAll(ret, "VV", "X")
	ret = strings.ReplaceAll(ret, "VIV", "IX")
	ret = strings.ReplaceAll(ret, "XXXXX", "L")
	ret = strings.ReplaceAll(ret, "XXXX", "XL")
	ret = strings.ReplaceAll(ret, "LL", "C")
	ret = strings.ReplaceAll(ret, "LXL", "XC")
	ret = strings.ReplaceAll(ret, "CCCCC", "D")
	ret = strings.ReplaceAll(ret, "CCCC", "CD")
	ret = strings.ReplaceAll(ret, "DD", "M")
	ret = strings.ReplaceAll(ret, "DCD", "CM")

	return ret, nil
}
