package sieve

import (
	"errors"
)

var ErrNegativeNumber = errors.New("Cannot use a negative number")

func Sieve(n int) (bool, error) {
	if n % 2 == 0 {
		// even numbers are not prime
		return false, nil
	}	

	if n < 0 {
		return false, ErrNegativeNumber
	}

	return false, nil
}
