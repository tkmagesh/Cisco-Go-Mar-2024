/*
	 Refactor the below program as follows
		- genPrimes() will generate the prime numbers concurrently
		- main() will print the generated prime numbers one at a time as and when a prime number is generated
*/
package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	for _, primeNo := range primes {
		fmt.Printf("prime : %d\n", primeNo)
	}
}

func genPrimes(start, end int) []int {
	// primes :=[]int{}
	primes := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
