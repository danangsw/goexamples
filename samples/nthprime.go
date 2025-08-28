package samples

import (
	"fmt"
)

// NthPrime is a simple code for counting what is the n-th prime number 
func NthPrime(nth int, nthprnt bool) int  {
	if nth <= 0 {
		return 0
	}

	nthprime := 1

	for i := 1; i <= nth; {
		nthprime++
		if isPrime(nthprime) {
			i++;

			if nthprnt {
				fmt.Printf("%v ", nthprime)
			}
		}
	}

	if nthprnt {
		fmt.Println()
	}

	return nthprime
}

func isPrime(n int) bool {
	// Corner case
	if (n <= 3) {
		return !(n <= 1)
	}

	// This is checked so that we  
    // can skip middle five numbers 
	// in below loop
	if n % 2 == 0 || n % 3 == 0 {
		return false
	} 

	for i := 5; i * i <= n; i += 6 {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
	}

	return true
}