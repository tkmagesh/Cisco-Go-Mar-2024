package main

import (
	"fmt"

	"github.com/fatih/color"
	calc "github.com/tkmagesh/cisco-go-mar-2024/09-modularity/calculator"
	"github.com/tkmagesh/cisco-go-mar-2024/09-modularity/calculator/utils"
	"github.com/tkmagesh/cisco-go-mar-2024/09-modularity/models"
)

func main() {
	color.Red("application executed")
	fmt.Println(greet("Magesh"))

	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Operation Count :", calculator.OpCount())
	*/

	// using the alias
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Operation Count :", calc.OpCount())

	fmt.Println("IsEven(21) :", utils.IsEven(21))
	fmt.Println("IsOdd(21) :", utils.IsOdd(21))

	pen := models.Product{}
	fmt.Printf("%#v\n", pen)
	fmt.Println(pen.Id)

}
