package diamond

import (
	"bytes"
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	// Validate input
	if char < 'A' || char > 'Z' {
		return "", errors.New("char out of range")
	}

	row_lenght := (2 * (char - 'A')) + 1
	rows := make([]string, int(row_lenght))

	for letter := byte('A'); letter <= char; letter++ {
		row := bytes.Repeat([]byte{' '}, int(row_lenght))

		row[char-letter], row[row_lenght-1-char+letter] = letter, letter

		rows[letter-'A'], rows[row_lenght-1-letter+'A'] = string(row), string(row)
	}

	return strings.Join(rows, "\n"), nil
}
