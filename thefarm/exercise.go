package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

// TODO: define the 'ValidateInputAndDivideFood' function

// TODO: define the 'ValidateNumberOfCows' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

type FodderCalculator interface {
	FodderAmount(cows int) (float64, error)
	FatteningFactor() (float64, error)
}

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amount, err1 := fc.FodderAmount(cows)
	factor, err2 := fc.FatteningFactor()

	if err1 != nil {
		return 0, err1
	}

	if err2 != nil {
		return 0, err2
	}

	return amount * factor / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fc, cows)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError interface {
	Error() string
}

func ValidateNumberOfCows(cows int) InvalidCowsError {
	if cows < 0 {
		return fmt.Errorf("%d cows are invalid: there are no negative cows", cows)
	}

	if cows == 0 {
		return fmt.Errorf("%d cows are invalid: no cows don't need food", cows)
	}

	return nil
}
