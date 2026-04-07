package main

import "fmt"

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}
func main() {
	var N int
	fmt.Scan(&N)
	odds := []int{}
	j := 1
	for len(odds) < N {
		if isOdd(j) {
			odds = append(odds, j)
		}
		j++
	}
	fmt.Println(odds)
}
