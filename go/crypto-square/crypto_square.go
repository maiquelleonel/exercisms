package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	re := regexp.MustCompile(`\W`)
	clean := re.ReplaceAllString(strings.ToLower(pt), "")
	if len(clean) == 0 {
		return ""
	}
	c, r := dimensions(len(clean))
	sq := squarefy(clean, c, r)
	crypt := crypt(sq, c)
	return strings.Join(crypt, " ")
}

func dimensions(l int) (int, int) {
	c := math.Sqrt(float64(l))
	r := (2 * c) / 2
	return int(math.Ceil(c)), int(math.Floor(r))
}

func squarefy(s string, c int, r int) []string {
	var ret []string
	for i := 0; i < r; i++ {
		if i < r-1 {
			ret = append(ret, s[i*c:i*c+c])
		} else {
			cand := s[i*c:]
			if len(cand) > c {
				ret = append(ret, cand[:c])
				cand = cand[c:]
			}
			if len(cand) < c {
				cand += strings.Repeat(" ", c-len(cand))
			}
			ret = append(ret, cand)
		}
	}
	return ret
}

func crypt(sq []string, c int) []string {
	ret := make([]string, c)
	for i := 0; i < c; i++ {
		for _, row := range sq {
			ret[i] += string(row[i])
		}
	}
	return ret
}
