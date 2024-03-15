package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // schedule the execution of f1 through the scheduler
	f2()
	// block the execution of this function (main) so that the scheduler can look for other goroutines that scheduled and execute them
	// DO NOT DO THIS IN PRODUCTION
	time.Sleep(1 * time.Second)
}

/*
func f1() {
	fmt.Println("f1 invoked")
}
*/

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
