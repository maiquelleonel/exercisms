package raindrops

import "fmt"

func Convert(number int) string {

	var sounds = []string{"Pling", "Plang", "Plong"}
	var ret string
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
