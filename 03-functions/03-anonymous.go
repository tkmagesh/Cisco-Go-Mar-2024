package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	// with 1 parameter
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	msg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a fantastic day!", userName)
	}("Magesh")
	fmt.Println(msg)

	q, r := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Println(q, r)
}

// with > 1 parameters
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

// with return value

// with > 1 return values
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// using named results
