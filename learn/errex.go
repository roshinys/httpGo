package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type divideError struct {
	dividend float64
}

// ?
func (de divideError) Error() string {
	return errors.New(fmt.Sprintf("cannot divide %.2f by zero", de.dividend)).Error()
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
