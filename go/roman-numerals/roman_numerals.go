package romanpattern

import (
	"errors"
	"math"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("Invalid")
	}
	var div, rem int
	var ret, suffix string

	switch {
	case input < 10:
		ret = pattern(input, "V", "I", "X")
	case input < 100:
		div = input / 10
		rem = int(math.Mod(float64(input), 10))
		suffix, _ = ToRomanNumeral(rem)
		ret = pattern(div, "L", "X", "C") + suffix
	case input < 1_000:
		div = input / 100
		rem = int(math.Mod(float64(input), 100))
		suffix, _ = ToRomanNumeral(rem)
		ret = pattern(div, "D", "C", "M") + suffix
	case input < 4_000:
		div = input / 1_000
		rem = int(math.Mod(float64(input), 1_000))
		suffix, _ = ToRomanNumeral(rem)
		ret = pattern(div, "C", "M", "") + suffix
	}

	return ret, nil
}

func pattern(num int, middle string, mod string, next string) string {
	switch {
	case num <= 3:
		return strings.Repeat(mod, num)
	case num == 4:
		return mod + middle
	case num <= 8:
		return middle + strings.Repeat(mod, num-5)
	case num == 9:
		return mod + next
	}
	return ""
}
