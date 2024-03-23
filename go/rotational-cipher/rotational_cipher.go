package rotationalcipher

import (
	"slices"
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var alpha, cipher []string
	var ret strings.Builder

	alpha = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	cipher = append(cipher, alpha[shiftKey:]...)
	cipher = append(cipher, alpha[:shiftKey]...)

	for _, ch := range plain {
		if unicode.IsLetter(ch) {
			ret.WriteString(
				capitalize(
					cipher[slices.Index(alpha, strings.ToLower(string(ch)))],
					unicode.IsUpper(ch)))
			continue
		}
		ret.WriteRune(ch)
	}
	return ret.String()

}

func capitalize(t string, isUpper bool) string {
	if isUpper {
		return strings.ToUpper(t)
	}
	return t
}
