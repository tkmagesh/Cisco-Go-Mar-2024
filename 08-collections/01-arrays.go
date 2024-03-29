package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// use :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating an array (using indexer)
	fmt.Println("iterating an array (using indexer)")
	for idx := 0; idx < 5; idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	// iterating an array (using range)
	fmt.Println("iterating an array (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	var nos2 = nos // a copy of the array is created
	nos2[0] = 9999

	fmt.Println("nos2 :", nos2)
	fmt.Println("nos :", nos)

	// No need to deference to access the members of a "pointer to an array"
	nosPtr := &nos
	// fmt.Println((*nosPtr)[0])
	fmt.Println(nosPtr[0])

	fmt.Println("Before sorting, nos :", nos)
	sort(&nos)
	fmt.Println("After sorting, nos :", nos)

	/*
		multi := [5][5]int{}
		fmt.Println(multi)
	*/
}

func sort(vals *[5]int) /* no return values */ {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if vals[i] > vals[j] {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
	}
}
