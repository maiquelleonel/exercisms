package raindrops

import "fmt"

func Convert(number int) string {
	var ret string
	var sounds = []string{"Pling", "Plang", "Plong"}
	for i, divisor := range []int{3, 5, 7} {
		if number%divisor == 0 {
			ret += sounds[i]
		}
	}

	if len(ret) > 0 {
		return ret
	}

	return fmt.Sprintf("%v", number)
}
