package main

import (
	"fmt"
	"strconv"
)

func order(n int) string {
	if n == 0 {
		return ""
	}
	return order(n-1) + strconv.Itoa(n) + " "
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println("order: ", order(n))
}
