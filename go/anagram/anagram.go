package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var ret []string
	subjectChecksum := checksum(subject)
	for _, candidate := range candidates {
		if !strings.EqualFold(subject, candidate) {
			if is_anagram(subject, candidate) && subjectChecksum == checksum(candidate) {
				ret = append(ret, candidate)
			}
		}
	}
	return ret
}

func charMap(s string) map[rune]int {
	ret := make(map[rune]int, 0)
	for _, l := range strings.ToLower(s) {
		ret[l]++
	}
	return ret
}

func is_anagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aCharMap := charMap(a)
	bCharMap := charMap(b)
	for char := range aCharMap {
		if aCharMap[char] != bCharMap[char] {
			return false
		}
	}
	return true
}

var checksums map[string]int

func checksum(s string) int {
	if checksums[s] != -1 {
		return checksums[s]
	}
	var sum int
	for val := range s {
		sum += int(val)
	}
	checksums[s] = sum
	return sum
}
