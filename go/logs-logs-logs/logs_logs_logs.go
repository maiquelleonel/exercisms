package logs

import (
	"strings"
	"unicode/utf8"
)

var exclamation rune = '❗'
var search rune = '🔍'
var weather rune = '☀'

// Application identifies the application emitting the given log.
func Application(log string) string {

	e := strings.IndexRune(log, exclamation)
	s := strings.IndexRune(log, search)
	w := strings.IndexRune(log, weather)

	if e == 0 || e > s && e > w && 0 > s && 0 > w {
		return "recommendation"
	}

	if s == 0 || s > e && s > w && 0 > e && 0 > w {
		return "search"
	}

	if w == 0 || w > s && w > e && 0 > s && 0 > e {
		return "weather"
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune rune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if r, size := utf8.DecodeLastRuneInString(log); r != utf8.RuneError {
		return limit > len(log)-size
	}
	return false

}
