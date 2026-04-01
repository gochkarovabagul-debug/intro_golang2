package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("san girizin: ")
	fmt.Scan(&n)
	var faktorial = n
	for i := 1; i < n; i++ {
		faktorial *= (n - i)
	}
	fmt.Printf("sanyn faktorialy: %v\n", faktorial)
}
