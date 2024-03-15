package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go getNos(ch)
	for val := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(val)
	}
}

func getNos(ch chan int) {
	count := rand.Intn(20)
	for i := 1; i <= count; i++ {
		ch <- 10 * i
	}
	fmt.Println("All the data are produced. Closing the channel!")
	close(ch)
}
