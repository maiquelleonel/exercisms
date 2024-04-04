// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\s_?|-`)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var words []string = strings.Fields(re.ReplaceAllString(s, " "))
	var acr string
	for _, word := range words {
		acr += string(word[0])
	}
	return strings.ToUpper(acr)
}
