package reverse

import (
	"slices"
	"strings"
)

func Reverse(input string) string {
	str := strings.Split(input, "")
	slices.Reverse(str)
	return strings.Join(str, "")

}
