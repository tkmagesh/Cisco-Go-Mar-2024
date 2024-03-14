/*
	 Refactor the below program as follows
		- genPrimes() will generate the prime number within the given range
		- main() will print the generated prime numbers one after another
*/
package main

import "fmt"

func main() {
OUTER_LOOP:
	for no := 2; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue OUTER_LOOP
			}
		}
		fmt.Printf("prime : %d\n", no)
	}
}
