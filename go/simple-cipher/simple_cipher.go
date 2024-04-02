package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	distance int
}

type vigenere struct {
	distance []int
}

func NewCaesar() Cipher {
	return &shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	if distance < 0 {
		distance += 26
	}
	return &shift{distance: distance}
}

func (c shift) Encode(input string) string {
	var ret strings.Builder

	for _, letter := range strings.ToLower(input) {
		if unicode.IsLetter(letter) {
			r := calcNextRune(letter, c.distance)
			ret.WriteRune(rune('a' + r))
		}
	}

	return ret.String()
}

func (c shift) Decode(input string) string {
	var ret strings.Builder

	for _, letter := range strings.ToLower(input) {
		if unicode.IsLetter(letter) {
			r := calcNextRune(letter, 26-c.distance)
			ret.WriteRune(rune('a' + r))
		}
	}
	return ret.String()
}

func NewVigenere(key string) Cipher {
	if key == "" || strings.Count(key, "a") == len(key) {
		return nil
	}

	var distance []int
	for _, letter := range key {

		if letter < 'a' || letter > 'z' {
			return nil
		}

		distance = append(distance, int(letter-'a'))
	}

	return &vigenere{distance}
}

func (v vigenere) Encode(input string) string {
	var ret strings.Builder
	var index int
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			r := calcNextRune(letter, v.distance[index%len(v.distance)])
			ret.WriteRune(rune('a' + r))
			index++
		}
	}
	return ret.String()
}

func (v vigenere) Decode(input string) string {
	var ret strings.Builder
	var index int
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			r := calcNextRune(letter, 26-v.distance[index%len(v.distance)])
			ret.WriteRune(rune('a' + r))
			index++
		}
	}
	return ret.String()
}

func calcNextRune(l rune, n int) int {
	return (int(l) - 'a' + n) % 26
}
