package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var output []int
	if inputBase < 2 {
		return output, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return output, errors.New("output base must be >= 2")
	}

	var num float64
	var j int

	for i := len(inputDigits) - 1; i >= 0; i-- {
		if inputDigits[j] < 0 || inputDigits[j] >= inputBase {
			return output, errors.New("all digits must satisfy 0 <= d < input base")
		}
		num += float64(inputDigits[j]) * math.Pow(float64(inputBase), float64(i))
		j++
	}

	for num >= 1 {
		rm := int(num) % outputBase
		output = append([]int{rm}, output...)
		num /= float64(outputBase)
	}

	if output == nil {
		output = []int{0}
	}

	return output, nil
}
