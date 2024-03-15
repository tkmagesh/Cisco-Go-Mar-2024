/*
	 Refactor the below program as follows
		- genPrimes() will generate the prime numbers concurrently
		- main() will print the generated prime numbers one at a time as and when a prime number is generated
*/
package main

import "fmt"

func main() {
	primeCh := genPrimes(2, 100)
	for primeNo := range primeCh {
		fmt.Printf("prime : %d\n", primeNo)
	}
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
