package pangram

import "strings"

func IsPangram(input string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	input = strings.ToLower(input)
	for _, letter := range alphabet {
		if !strings.ContainsRune(input, letter) {
			return false
		}
	}
	return true
}
