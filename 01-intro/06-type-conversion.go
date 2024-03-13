package main

import "fmt"

func main() {
	var i int8 = 100
	var f float64
	// syntax for type conversion => use the type name like a function
	f = float64(i)
	fmt.Println(f)
}
