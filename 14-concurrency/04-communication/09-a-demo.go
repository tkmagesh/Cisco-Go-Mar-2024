package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// share memory by communicating
	go add(wg, ch, 100, 200)
	result := <-ch
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(wg *sync.WaitGroup, ch chan int, x, y int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
