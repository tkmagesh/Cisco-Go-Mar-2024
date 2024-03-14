package main

import "fmt"

func main() {
	var no int
	no = 100
	fmt.Println("Before incrementing, no = ", no)
	fmt.Println("[main] &no :", &no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 100, 200
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)

}

func increment(val *int) /* no return values */ {
	fmt.Println("[increment] &val :", val)
	*val++
}

func swap(n1, n2 *int) /* no return values */ {
	*n1, *n2 = *n2, *n1
}
