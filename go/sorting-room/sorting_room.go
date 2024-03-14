package sorting

import (
	"fmt"
	"regexp"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", float64(f))
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		if i, er := strconv.Atoi(fnb.Value()); er == nil {
			return i
		}
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf(
		"This is a fancy box containing the number %.1f",
		float64(ExtractFancyNumber(fnb)))
}

type NewNumberBox struct {
	n int
}

func (i NewNumberBox) Number() int {
	return i.n
}

func onlyDigits(s string) string {
	re := regexp.MustCompile(`[^\d\.?\d?]`)
	return re.ReplaceAllLiteralString(s, "")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	data := onlyDigits(fmt.Sprintf("%v", i))
	switch i.(type) {
	case int:
		if num, err := strconv.Atoi(data); err == nil {
			return DescribeNumber(float64(num))
		}
	case float64:
		if num, err := strconv.ParseFloat(data, 64); err == nil {
			return DescribeNumber(num)
		}
	case NumberBox:
		if num, e := strconv.Atoi(data); e == nil {
			return DescribeNumberBox(NewNumberBox{n: int(num)})
		}
	case FancyNumberBox:
		return DescribeFancyNumberBox(FancyNumber{data})
	}
	return "Return to sender"
}
