package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var fn func()
	fn = getFn()
	fn()
	time.Sleep(time.Duration(rand.Intn(int(time.Now().Unix()))))
	fn = getFn()
	fn()
}

func getFn() func() {
	r := rand.Intn(int(time.Now().Unix()))
	if r%2 == 0 {
		return f1
	}
	return f2
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
