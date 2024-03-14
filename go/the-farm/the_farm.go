package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	var totalFoodNeeded float64
	var factor float64
	var err error

	if totalFoodNeeded, err = calc.FodderAmount(cows); err != nil {
		return 0, err
	}

	if factor, err = calc.FatteningFactor(); err != nil {
		return 0, err
	}

	return (totalFoodNeeded * factor) / float64(cows), nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}
	return nil
}

type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
