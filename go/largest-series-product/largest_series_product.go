package lsproduct

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("span needs to be greather then string")
	}

	if match, _ := regexp.MatchString(`\D`, digits); match {
		return 0, errors.New("only digits can be parse")
	}

	series := findSeries(strings.Split(digits, ""), span)

	var products []int64

	for _, serie := range series {
		products = append(products, serie[span])
	}

	return slices.Max(products), nil
}

func calcProduct(numbers []int64) int64 {
	var product int64 = 1
	for _, number := range numbers {
		product *= number
	}
	return product
}

func findSeries(numbers []string, size int) [][]int64 {
	var number int64
	var slot []int64
	var ret [][]int64
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < size; j++ {
			number = 0
			if i+j < len(numbers) {
				number, _ = strconv.ParseInt(numbers[i+j], 10, 64)
			}
			slot = append(slot, number)
		}
		if len(ret) < len(numbers)-1 {
			slot = append(slot, calcProduct(slot))
			ret = append(ret, slot)
			slot = make([]int64, 0)
		}
	}
	return ret
}
