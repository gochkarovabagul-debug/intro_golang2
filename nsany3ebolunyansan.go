package main

import "fmt"

func check(n int) bool {
	if n%3 == 0 {
		return true
	}
	return false
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
	fmt.Println(slice)
}
