package main

import "fmt"

func main() {
	ch := make(chan int)
	/*
		go func() {
			ch <- 100 // (NB)
		}()
		data := <-ch // (B, UB)
		fmt.Println(data)
	*/

	go func() {
		data := <-ch // (NB)
		fmt.Println(data)
	}()
	ch <- 100 // (B, UB)
}
