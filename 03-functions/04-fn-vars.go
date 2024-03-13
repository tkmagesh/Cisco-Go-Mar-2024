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

	var greetUser func(string, string)
	greetUser = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}
	greetUser("Magesh", "Kuppan")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a fantastic day!", userName)
	}
	msg := getGreetMsg("Magesh")
	fmt.Println(msg)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Println(q, r)

}
