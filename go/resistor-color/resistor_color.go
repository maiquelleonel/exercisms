package resistorcolor

import "slices"

// Colors should return the list of all colors.
var colors = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func Colors() []string {
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return slices.Index(Colors(), color)
}
