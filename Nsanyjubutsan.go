package main

import "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
func main() {
	var N int
	fmt.Scan(&N)
	evens := []int{}
	j := 1
	for len(evens) < N {
		if isEven(j){
			evens = append(evens, j)
		}
		j++
	}

	fmt.Println(evens)
}
