package main

import "fmt"

func main() {
	var productRanks map[string]int
	/*
	   productRanks = make(map[string]int)
	   	productRanks["pen"] = 5
	   	productRanks["pencil"] = 1
	   	productRanks["marker"] = 2
	*/
	// productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 2}
	productRanks = map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 2,
	}
	fmt.Println(productRanks)

	// accessing the value using the key
	fmt.Println("Rank of pen is :", productRanks["pen"])

	// iterating a map
	fmt.Println("iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of a key
	keyToCheck := "scribble-pad"
	// fmt.Printf("Rank of %q is %d\n", keyToCheck, productRanks[keyToCheck])

	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key [%q] does not exist\n", keyToCheck)
	}

	// removing a key/value pair (using delete() function)
	keyToRemove := "scribble-pad"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing %q, productRanks = %v\n", keyToRemove, productRanks)

}
