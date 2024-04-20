package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	result := make([]map[string]int, 0)
	db := make(map[string]int, 0)
	for i := 0; i < len(input); i++ {
		if _, e := db[string(input[i])]; !e && i > 0 {
			result = append(result, db)
			db = map[string]int{}
		}
		db[string(input[i])]++
	}
	result = append(result, db)
	str := ""
	for _, m := range result {
		for k, v := range m {
			if v == 1 {
				str += k
			} else {
				str += fmt.Sprintf("%d%s", v, k)
			}
		}
	}

	return str
}

func RunLengthDecode(input string) string {
	var result, letter, num string
	re := regexp.MustCompile(`(\d{1,2}\w{1}|\d?\s)|(\w)`)
	ret := re.FindAllString(input, -1)
	for _, data := range ret {
		if len(data) == 1 {
			result += data
		} else {
			letter = string(data[len(data)-1])
			num = strings.Replace(data, letter, "", -1)
			times, _ := strconv.Atoi(num)
			result += strings.Repeat(letter, times)
		}
	}

	return result
}
