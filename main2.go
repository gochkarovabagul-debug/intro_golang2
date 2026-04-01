package main

import (
	"fmt"
)

func kw(x, n int) int {
	if n == 1 {
		return x
	} else {
		kwadrat := x
		for i := 1; i < n; i++ {
			kwadrat *= x
		}
		return kwadrat
	}
}

func main() {
	var n int
	var x int
	fmt.Printf("enter your numbers: ")
	fmt.Scan(&x)
	fmt.Scan(&n)
	fmt.Println("kwadrat: ", kw(x, n))
}
