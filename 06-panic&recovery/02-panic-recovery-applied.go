package main

import "fmt"

func main() {
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideClient(100, divisor)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

// facade that converts a panic into an error (so that a different course of action can be pursued)
func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error) // converting the "any" value returned by recover() to an "error" type
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api
func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
