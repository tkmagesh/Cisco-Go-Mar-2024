/*
write a divide function that returns the quotient and remainder
the function should return an error if the divisor == 0
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r, err := divide(100, 0)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println(q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divide by zero error")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

// using named results
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divide by zero error")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
