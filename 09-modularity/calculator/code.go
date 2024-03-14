package calculator

import "fmt"

// private
var opCount map[string]int

func init() {
	fmt.Println("[calculator - code.go] - init")
	opCount = make(map[string]int)
}

// public
func OpCount() map[string]int {
	return opCount
}
