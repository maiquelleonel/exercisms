package resistorcolorduo

import (
	"slices"
	"strconv"
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var encodedColors = []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}

	var totalValue string

	for _, color := range colors[:2] {
		totalValue += strconv.Itoa(slices.Index(encodedColors, color))
	}
	intval, _ := strconv.Atoi(totalValue)
	return intval
}
