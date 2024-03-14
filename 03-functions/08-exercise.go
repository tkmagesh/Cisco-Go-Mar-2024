/*
refactor the below solutions into functions
each function should have only one reason to change (SRP)
*/

package main

import "fmt"

func main() {

	var (
		userChoice, n1, n2, result int
		op                         func(int, int) int
	)

	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		op = getOperation(userChoice)
		if op == nil {
			fmt.Println("Invalid choice")
			continue
		}
		n1, n2 = getOperands()
		result = op(n1, n2)
		fmt.Println("Result :", result)
	}
}

func getOperation(userChoice int) func(int, int) int {
	switch userChoice {
	case 1:
		return add
	case 2:
		return subtract
	case 3:
		return multiply
	case 4:
		return divide
	default:
		return nil
	}
}

func add(n1, n2 int) int {
	return n1 + n2
}

func subtract(n1, n2 int) int {
	return n1 - n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}

func divide(n1, n2 int) int {
	return n1 / n2
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the operands :")
	fmt.Scanln(&n1, &n2)
	return
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice :")
	fmt.Scanln(&userChoice)
	return userChoice
}
