package main

import (
	"fmt"
	"time"
)

func main() {
	result := go add(100, 200)
	fmt.Println("Result :", result)
}

func add(x, y int) int {
	time.Sleep(2 * time.Second)
	return x + y
}
