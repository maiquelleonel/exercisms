package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	var ret bool = true
	re := regexp.MustCompile(`\W`)
	word = strings.ToLower(word)
	word = re.ReplaceAllString(word, "")
	var splitWord []string = strings.Split(word, "")
	for _, letter := range splitWord {
		if strings.Count(word, letter) > 1 {
			ret = false
			break
		}
	}
	return ret
}
