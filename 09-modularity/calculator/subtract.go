package calculator

import "fmt"

func init() {
	fmt.Println("[calculator - subtract.go] - init")
}

func Subtract(x, y int) int {
	opCount["subtract"]++
	return x - y
}
