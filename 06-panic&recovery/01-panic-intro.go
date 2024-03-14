package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("application exiting because of a panic : ", x)
			return
		}
		fmt.Println("application executed successfully")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("f1 exiting because of a panic : ", x)
			return
		}
		fmt.Println("[f1] - deferred")
	}()
	fmt.Println("f1 started")
	panic("dummy panic")
	fmt.Println("f1 completed")
}
