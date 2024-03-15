package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	/*
		wg.Add(1)
		wg.Wait()
	*/
	wg.Add(15)
	for i := 0; i < 10; i++ {
		go fn(wg)
	}
	wg.Wait()
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn invoked")
}
