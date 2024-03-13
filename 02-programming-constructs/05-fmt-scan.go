package main

import "fmt"

func main() {
	var username string
	fmt.Println("Enter your name :")
	// get the input
	fmt.Scanln(&username)
	fmt.Printf("Hi %s, Have a nice day!\n", username)

	var x, y int
	fmt.Println("Enter the numbers")
	fmt.Scanf("%d %d\n", &x, &y)
	fmt.Println(x + y)
}
