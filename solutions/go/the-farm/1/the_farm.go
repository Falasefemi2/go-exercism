package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cow     int
	message string
}

// TODO: define the 'DivideFood' function
var errCalc = errors.New("something went wrong")

func DivideFood(fc FodderCalculator, cow int) (float64, error) {
    fodder, err := fc.FodderAmount(cow)
    if err != nil {
        return 0, err
    }
    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    amount := (fodder * factor) / float64(cow)
    return amount, nil
}


// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cow int) (float64, error) {
	if cow < 0 {
		return 0, errors.New("invalid number of cows")
	}
	if cow == 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cow)
}


// TODO: define the 'ValidateNumberOfCows' function
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cow, e.message)
}

func ValidateNumberOfCows(cow int) error {
	if cow < 0 {
		return &InvalidCowsError{
			cow:     cow,
			message: "there are no negative cows",
		}
	}
	if cow == 0 {
		return &InvalidCowsError{
			cow:     cow,
			message: "no cows don't need food",
		}
	}
	return nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
