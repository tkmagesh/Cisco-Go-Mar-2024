package main

import "fmt"

func main() {
	var no int
	no = 10

	// address of a value (referencing)
	var noPtr *int
	noPtr = &no
	fmt.Println(noPtr)

	// accessing the value in the underlying address (dereferencing)
	var val int
	val = *noPtr
	fmt.Println(val)

	fmt.Println(no == *(&no))
}
