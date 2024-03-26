package resistorcolorduo

import "slices"

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

	return slices.Index(encodedColors, colors[0])*10 +
		slices.Index(encodedColors, colors[1])
}
