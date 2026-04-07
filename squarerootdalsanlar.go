package main

import "fmt"

func check(n int) bool {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return false
		}
	}
	return true
}
func main() {
	var N int
	fmt.Scan(&N)
	slice := []int{}
	n := 1
	for len(slice) < N {
		if check(n) {
			slice = append(slice, n)
		}
		n++
	}
	fmt.Print(slice)
}
