// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re := regexp.MustCompile(`\s_?|\s-?\s?|-`)
	words := re.Split(s, 15)
	var acr string
	println(fmt.Sprintf("%s", words))
	for _, word := range words {
		if len(word) == 0 || unicode.IsSpace(rune(word[0])) {
			continue
		}

		if len(word) > 1 && unicode.IsPunct(rune(word[0])) {
			acr += string(word[1])
			continue
		}

		acr += string(word[0])
	}
	return strings.ToUpper(acr)
}
