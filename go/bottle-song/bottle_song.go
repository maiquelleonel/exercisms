package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	var sNumbers [][]string = [][]string{
		{"no", "bottles"},
		{"One", "bottle"},
		{"Two", "bottles"},
		{"Three", "bottles"},
		{"Four", "bottles"},
		{"Five", "bottles"},
		{"Six", "bottles"},
		{"Seven", "bottles"},
		{"Eight", "bottles"},
		{"Nine", "bottles"},
		{"Ten", "bottles"},
	}

	var ret []string
	var verses int = 1
	for {
		ret = append(ret,
			fmt.Sprintf("%s green %s hanging on the wall,",
				sNumbers[startBottles-verses+1][0], sNumbers[startBottles-verses+1][1]))
		ret = append(ret,
			fmt.Sprintf("%s green %s hanging on the wall,",
				sNumbers[startBottles-verses+1][0], sNumbers[startBottles-verses+1][1]))
		ret = append(ret, "And if one green bottle should accidentally fall,")

		ret = append(ret,
			fmt.Sprintf("There'll be %s green %s hanging on the wall.",
				strings.ToLower(sNumbers[startBottles-verses][0]), sNumbers[startBottles-verses][1]))

		if verses == takeDown {
			break
		}

		verses++
		ret = append(ret, "")

	}
	return ret
}
