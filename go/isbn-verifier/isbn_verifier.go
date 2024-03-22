package isbn

import (
	"regexp"
	"strings"
)

func IsValidISBN(isbn string) bool {

	if len(isbn) == 0 {
		return false
	}
	re := regexp.MustCompile(`^\d-?\d{3}-?\d{5}-?[0-9X]$`)
	if !re.MatchString(isbn) {
		return false
	}

	isbn = strings.Replace(isbn, "-", "", 3)

	if len(isbn) != 10 {
		return false
	}
	factors := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var sum int
	for i, num := range isbn {
		if num == 'X' {
			sum += 10
		} else {
			sum += int(num-'0') * factors[i]
		}
	}

	return sum%11 == 0

}
