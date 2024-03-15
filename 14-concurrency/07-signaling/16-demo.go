package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := time.After(5 * time.Second)
	fibCh := genFib(stopCh)
	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
}

func genFib(stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-stopCh:
				fmt.Println("timeout occurred... exiting!")
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
