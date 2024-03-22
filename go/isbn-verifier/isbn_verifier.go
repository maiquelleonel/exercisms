package isbn

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {
	re := regexp.MustCompile(`^\d-?\d{3}-?\d{5}-?[0-9X]$`)
	if !re.MatchString(isbn) {
		return false
	}

	isbn = strings.Replace(isbn, "-", "", 3)
	var sum int
	for i, n := range isbn {
		if n == 'X' {
			sum += 10
		} else {
			sum += int(n-'0') * (10 - i)
		}
	}

	return sum%11 == 0

}
