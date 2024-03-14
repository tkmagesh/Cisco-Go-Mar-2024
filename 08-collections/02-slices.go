package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// appending new items
	nos = append(nos, 9)
	nos = append(nos, 8, 6, 7)

	// appending another slice
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// len()
	fmt.Println("len(nos) :", len(nos))

	// slicing
	fmt.Println("nos[3:6] = ", nos[3:6])
	fmt.Println("nos[3:] = ", nos[3:])
	fmt.Println("nos[:6] = ", nos[:6])

	/*
		nos2 := nos // assignment creates a copy of the "pointer to the underlying array"
		nos2[0] = 9999
		fmt.Println("nos :", nos)
		fmt.Println("nos2 :", nos2)
	*/

	// creating a deep copy
	var nos2 []int = make([]int, len(nos)) // preallocate memory & initialize it
	copy(nos2, nos)
	nos2[0] = 9999
	fmt.Println("nos :", nos)
	fmt.Println("nos2 :", nos2)

	fmt.Println("Before sorting, nos :", nos)
	sort(nos)
	fmt.Println("After sorting, nos :", nos)
}

func sort(vals []int) {
	for i := 0; i < len(vals)-1; i++ {
		for j := i + 1; j < len(vals); j++ {
			if vals[i] > vals[j] {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
	}
}
