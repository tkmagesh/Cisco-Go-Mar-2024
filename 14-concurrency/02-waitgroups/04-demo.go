package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines to spin")
	flag.Parse()

	fmt.Printf("Starting %d goroutines... hit ENTER to start\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()

	fmt.Println("All done.. hit ENTER to shutdown")
	fmt.Scanln()

}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
