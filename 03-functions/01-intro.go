package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	greetUser("Magesh", "Kuppan")
	fmt.Println(getGreetMsg("Magesh"))

	// > 1 return values
	// fmt.Println(divide(100, 7))

	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	// divide 100 by 7 (using divide(...) function) and print only the quotient
}

func sayHi() {
	fmt.Println("Hi there!")
}

// with 1 parameter
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

// with > 1 parameters
/*
func greetUser(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greetUser(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

// with return value
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a fantastic day!", userName)
}

// with > 1 return values
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
