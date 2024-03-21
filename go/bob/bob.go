// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	phrase := strings.TrimSpace(remark)
	is_upper := isUpper(phrase)
	is_question := strings.HasSuffix(phrase, "?")
	re := regexp.MustCompile(`\w+`)
	if len(phrase) > 0 {
		if is_upper && is_question {
			return "Calm down, I know what I'm doing!"
		} else if is_question {
			return "Sure."
		} else if is_upper {
			return "Whoa, chill out!"
		} else if re.FindAllString(phrase, -1) != nil {
			return "Whatever."
		}
	}
	return "Fine. Be that way!"
}

func isUpper(s string) bool {
	re := regexp.MustCompile(`(?i)[a-z]`)
	// has no letter
	if re.FindAllString(s, -1) == nil {
		return false
	}

	return strings.ToUpper(s) == s
}
