package resistorcolortrio

import (
	"fmt"
	"strconv"
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

var UNIT = map[int]string{
	0: "",
	1: "",
	2: "",
	3: "kilo",
	4: "kilo",
	5: "kilo",
	6: "mega",
	7: "mega",
	8: "mega",
	9: "giga",
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(c []string) string {
	strNumber := fmt.Sprintf("%v%v", COLORS[c[0]], COLORS[c[1]])
	intVal, _ := strconv.Atoi(strNumber)
	switch c[2] {
	case "brown", "yellow", "green", "violet", "grey":
		strNumber += "0"
		intVal *= 10
	case "red":
		if c[1] == "black" {
			intVal /= 10
			COLORS[c[2]]++
		}
	}

	return fmt.Sprintf("%s %sohms", strconv.Itoa(intVal), UNIT[COLORS[c[2]]])
}
