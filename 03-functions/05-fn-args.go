package main

import "fmt"

func main() {
	exec(f1) // for f1
	exec(f2) // for f2
}

func exec(fn func()) {
	// execute either f1 or f2 depending on the argument
	fn()
}
func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
