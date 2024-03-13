package main

import "fmt"

func main() {
	counter := getCounter()
	fmt.Println(counter()) //=> 1
	fmt.Println(counter()) //=> 2
	fmt.Println(counter()) //=> 3
	fmt.Println(counter()) //=> 4
}

func getCounter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}
