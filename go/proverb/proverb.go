// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var ret []string
	for i := range rhyme {
		if i+1 == len(rhyme) {
			ret = append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			break
		}
		ret = append(ret, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	return ret
}
