// Eger N san girizilse,

// primes=[2]
// for i=1 den N cenli
// for j=primes[:len(primes)-1]+1; true; j++
// If isPrime(j)
// primes append
// break
package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var N int
	fmt.Scan(&N)
	primes := []int{2}
	for len(primes) < N {
		j := primes[len(primes)-1] + 1
		for {
			if isPrime(j) {
				primes = append(primes, j)
				break
			}
			j++
		}
	}

	fmt.Println(primes)
}
