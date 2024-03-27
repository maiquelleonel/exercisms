package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`\b\w{2}n't\b|\b\w{2}u'r\w\b|\b\w+\b`)
	words := re.FindAllString(strings.ToLower(phrase), -1)
	freq := make(Frequency, 0)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
