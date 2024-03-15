package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fibCh := genFib(stopCh)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
}

func genFib(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-stopCh:
				fmt.Println("Stop instruction received... exiting!")
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}
