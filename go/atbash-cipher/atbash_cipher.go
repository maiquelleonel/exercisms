package atbash

import (
	"strings"
	"unicode"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyz"
var cipher string = "zyxwvutsrqponmlkjihgfedcba"

func Atbash(s string) string {
	var wordSize int
	var ret string
	for _, letter := range strings.ToLower(s) {
		if unicode.IsPunct(letter) || unicode.IsSpace(letter) {
			continue
		}

		if wordSize == 5 {
			ret += " "
			wordSize = 0
		}

		if unicode.IsNumber(letter) {
			ret += string(letter)
		} else {
			ret += string(cipher[strings.IndexRune(alphabet, letter)])
		}
		wordSize++
	}

	return ret
}
