package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")
	/*
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
	*/
}
