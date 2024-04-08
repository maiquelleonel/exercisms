package resistorcolortrio

import (
	"fmt"
	"strconv"
	"strings"
)

var COLORS = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var UNITS = []string{
	"giga",
	"mega",
	"kilo",
}
var BASES = []int{
	1_000_000_000,
	1_000_000,
	1_000,
}

// Label describes the resistance value given the c of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(c []string) string {
	strNumber := fmt.Sprintf(
		"%d%s",
		COLORS[c[0]]*10+COLORS[c[1]],
		strings.Repeat("0", COLORS[c[2]]),
	)

	intVal, _ := strconv.Atoi(strNumber)

	scale := ""

	if intVal > 0 {
		for i, unit := range UNITS {
			if intVal%BASES[i] == 0 {
				intVal /= BASES[i]
				scale = unit
				break
			}
		}
	}

	return fmt.Sprintf("%d %sohms", intVal, scale)
}
