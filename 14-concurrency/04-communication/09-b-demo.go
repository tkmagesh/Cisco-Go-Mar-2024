package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// share memory by communicating
	go add(ch, 100, 200)
	result := <-ch
	fmt.Println("Result :", result)
}

func add(ch chan int, x, y int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
}
