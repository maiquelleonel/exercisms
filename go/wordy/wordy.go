package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	if match, _ := regexp.MatchString(`What\sis\s\-?\d+(?:\s(?:plus|minus|divided\sby|multiplied\sby)\s\-?\d+)*\?`, question); !match {
		return 0, false
	}

	reop := regexp.MustCompile(`(plus|minus|multiplied|divided)`)
	ops := reop.FindAllString(question, -1)

	renum := regexp.MustCompile(`-?\d+`)
	nums := renum.FindAllString(question, -1)

	result, _ := strconv.Atoi(nums[0])

	for i, op := range ops {
		next, _ := strconv.Atoi(nums[i+1])
		switch op {
		case "plus":
			result += next
		case "minus":
			result -= next
		case "multiplied":
			result *= next
		case "divided":
			result /= next
		}

	}

	return result, true
}
