package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	i = i - 1
	verses := []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, ",
		"three French Hens, ",
		"four Calling Birds, ",
		"five Gold Rings, ",
		"six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ",
		"eight Maids-a-Milking, ",
		"nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ",
		"eleven Pipers Piping, ",
		"twelve Drummers Drumming, ",
	}

	keys := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}

	var ret string = ""
	ret += fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", keys[i])
	if i > 0 {
		for i > 0 {
			ret += verses[i]
			i--
		}
		ret += "and " + verses[0]
	} else {
		ret += verses[0]
	}
	return ret

}

func Song() string {
	return strings.Join([]string{
		Verse(1),
		Verse(2),
		Verse(3),
		Verse(4),
		Verse(5),
		Verse(6),
		Verse(7),
		Verse(8),
		Verse(9),
		Verse(10),
		Verse(11),
		Verse(12),
	}, "\n")

}
