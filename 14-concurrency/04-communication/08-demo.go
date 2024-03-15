package main

import (
	"fmt"
	"sync"
	"time"
)

// communicate by sharing memory
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200)
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(wg *sync.WaitGroup, x, y int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	result = x + y
}
