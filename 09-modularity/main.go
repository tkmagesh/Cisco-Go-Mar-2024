package main

import (
	"fmt"

	"github.com/tkmagesh/cisco-go-mar-2024/09-modularity/calculator"
	"github.com/tkmagesh/cisco-go-mar-2024/09-modularity/calculator/utils"
)

func main() {
	fmt.Println("application executed")
	fmt.Println(greet("Magesh"))
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Operation Count :", calculator.OpCount())

	fmt.Println("IsEven(21) :", utils.IsEven(21))
	fmt.Println("IsOdd(21) :", utils.IsOdd(21))

}
